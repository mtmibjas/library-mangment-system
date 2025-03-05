package router

import (
	"fmt"
	"library-mngmt/app/config"
	"library-mngmt/app/container"
	"library-mngmt/app/http/controller"
	middle "library-mngmt/app/http/middleware"
	_ "library-mngmt/docs"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	echoSwagger "github.com/swaggo/echo-swagger"
)

type Validator struct {
	validator *validator.Validate
}

func Init(cfg *config.Config, ctr *container.Container) *echo.Echo {

	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Validator = &Validator{validator: validator.New()}

	requestLimit := middle.NewRequestLimiter(cfg.Service.RequestPerSecond)

	// Apply rate limiting middleware for Echo routes
	e.Use(requestLimit.RequestLimitMiddleware)

	e.GET("/healthcheck", func(c echo.Context) error {
		return c.String(http.StatusOK, " I'm breathing..")
	})

	anc := controller.NewAdminController(ctr)
	ac := controller.NewAuthController(cfg, ctr)
	bc := controller.NewBookController(ctr)
	uc := controller.NewUserController(cfg, ctr)

	authMiddlewares := []echo.MiddlewareFunc{
		middle.AuthMiddleware(cfg.Service.JWTSecret),
		ac.ValidateAPIKey,
	}
	// Auth routes
	authGroup := e.Group("/auth")
	authGroup.POST("/login", ac.Login)
	authGroup.POST("/logout", ac.ValidateAPIKey(ac.Logout), middle.AuthMiddleware(cfg.Service.JWTSecret))
	authGroup.POST("/refresh-token", ac.ValidateAPIKey(ac.RefreshToken))

	// Admin routes
	adminGroup := e.Group("/api/v1/admin")
	adminGroup.Group("", authMiddlewares...)
	adminManagement := adminGroup.Group("", ac.ValidateRolePermission("permission"))
	{
		adminManagement.POST("/permission", anc.CreatePermission)
		adminManagement.GET("/permissions", anc.GetPermissions)
		adminManagement.GET("/roles", anc.GetRoles)
		adminManagement.PUT("/permission", anc.UpdatePermission)
		adminManagement.DELETE("/permission/{id}", anc.DeletePermission)
		adminManagement.GET("/permission/{id}", anc.GetPermissions)
		adminManagement.PATCH("/permission/add", anc.AddPermissionToRole)
		adminManagement.PATCH("/permission/remove", anc.RemovePermissionFromRole)
	}

	bookGroup := e.Group("/api/v1/book")
	bookManagement := bookGroup.Group("", authMiddlewares...)
	{
		bookManagement.POST("", bc.CreateBook, ac.ValidateRolePermission("book"))
		bookManagement.PUT("/{id}", bc.UpdateBook, ac.ValidateRolePermission("book"))
		bookManagement.DELETE("/{id}", bc.DeleteBook, ac.ValidateRolePermission("book"))
		bookManagement.GET("/{id}", bc.GetBook)
		bookManagement.GET("", bc.GetBookList)
		bookManagement.GET("/{id}/history", bc.GetBorrowedHistoryByBookID)
	}

	userGroup := e.Group("/api/v1/user")
	userManagement := userGroup.Group("", authMiddlewares...)
	{
		userManagement.POST("", uc.CreateUser, ac.ValidateRolePermission("user"))
		userManagement.GET("/{id}", uc.GetUser)
		userManagement.PUT("/{id}", uc.UpdateUser)
		userManagement.DELETE("/{id}", uc.DeleteUser)
		userManagement.GET("", uc.GetUserList)
		userManagement.GET("/email", uc.GetUserByEmail)
		userManagement.GET("/{id}/history", uc.GetBorrowedHistoryByUserID)
	}

	e.GET("/swagger/*", echoSwagger.WrapHandler)

	go func() {
		OpenBrowser(fmt.Sprintf("%s:%d%s", cfg.Service.BaseURL, cfg.Service.Port, "/swagger/index.html"))
	}()
	return e
}
func (cv *Validator) Validate(i interface{}) error {
	return cv.validator.Struct(i)
}
