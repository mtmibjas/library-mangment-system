package controller

import (
	"library-mngmt/app/config"
	"library-mngmt/app/container"
	"library-mngmt/app/http/request"
	"library-mngmt/app/http/response"
	"library-mngmt/domain/usecases"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type UserController struct {
	Config       *config.Config
	Adapters     *container.Container
	UserServices *usecases.UserService
}

func NewUserController(cfg *config.Config, ctr *container.Container) *UserController {
	return &UserController{
		Config:       cfg,
		Adapters:     ctr,
		UserServices: usecases.NewUserService(ctr),
	}
}

// CreateUser godoc
// @Summary Create a new user
// @Description This API creates a new user.
// @Tags Users
// @Security BearerAuth
// @Security ApiKeyAuth
// @Accept json
// @Produce json
// @Param user body request.CreateUserRequest true "User data"
// @Success 201 {object} map[string]interface{} "User created successfully"
// @Failure 400 {object} map[string]interface{} "Invalid request"
// @Failure 500 {object} map[string]interface{} "Internal server error"
// @Router /api/v1/user [post]
func (uc *UserController) CreateUser(c echo.Context) error {

	req := request.CreateUserRequest{}
	if err := c.Bind(&req); err != nil {
		return response.Error(c, http.StatusBadRequest, err)
	}

	if err := c.Validate(req); err != nil {
		return response.Error(c, http.StatusBadRequest, err)
	}
	if err := uc.UserServices.CreateUser(req, uc.Config.Service.EncrytionKey); err != nil {
		return response.Error(c, http.StatusInternalServerError, err)
	}
	return response.Send(c, http.StatusCreated, map[string]any{
		"message": "User created successfully",
	})
}

// GetUser godoc
// @Summary  Get user details
// @Description This API to Get user details.
// @Tags Users
// @Security BearerAuth
// @Security ApiKeyAuth
// @Accept json
// @Produce json
// @Param id path int true "User ID"
// @Success 200 {object} map[string]interface{} "User details"
// @Failure 400 {object} map[string]interface{} "Invalid request"
// @Failure 500 {object} map[string]interface{} "Internal server error"
// @Router /api/v1/user/{id} [get]
func (uc *UserController) GetUser(c echo.Context) error {
	id := c.Param("id")
	userID, err := strconv.Atoi(id)
	if err != nil {
		return response.Error(c, http.StatusBadRequest, err)
	}
	result, err := uc.UserServices.GetUser(uint(userID))
	if err != nil {
		return response.Error(c, http.StatusInternalServerError, err)
	}
	return response.Send(c, http.StatusOK, map[string]any{
		"data": result,
	})
}

// UpdateUser godoc
// @Summary Update user details
// @Description This API updates user details.
// @Tags Users
// @Security BearerAuth
// @Security ApiKeyAuth
// @Accept json
// @Produce json
// @Param user body request.UpdateUserRequest true "User data"
// @Success 200 {object} map[string]interface{} "User updated successfully"
// @Failure 400 {object} map[string]interface{} "Invalid request"
// @Failure 500 {object} map[string]interface{} "Internal server error"
// @Router /api/v1/user [put]
func (uc *UserController) UpdateUser(c echo.Context) error {
	req := request.UpdateUserRequest{}
	if err := c.Bind(&req); err != nil {
		return response.Error(c, http.StatusBadRequest, err)
	}
	if err := c.Validate(req); err != nil {
		return response.Error(c, http.StatusBadRequest, err)
	}
	if err := uc.UserServices.UpdateUser(req); err != nil {
		return response.Error(c, http.StatusInternalServerError, err)
	}
	return response.Send(c, http.StatusOK, map[string]any{
		"message": "User updated successfully",
	})
}

// DeleteUser godoc
// @Summary Delete a user
// @Description This API deletes a user.
// @Tags Users
// @Security BearerAuth
// @Security ApiKeyAuth
// @Accept json
// @Produce json
// @Param id path int true "User ID"
// @Success 200 {object} map[string]interface{} "User deleted"
// @Failure 400 {object} map[string]interface{} "Invalid request"
// @Failure 500 {object} map[string]interface{} "Internal server error"
// @Router /api/v1/user/{id} [delete]
func (uc *UserController) DeleteUser(c echo.Context) error {
	id := c.Param("id")
	userID, err := strconv.Atoi(id)
	if err != nil {
		return response.Error(c, http.StatusBadRequest, err)
	}
	if err := uc.UserServices.DeleteUser(uint(userID)); err != nil {
		return response.Error(c, http.StatusInternalServerError, err)
	}
	return response.Send(c, http.StatusOK, map[string]any{
		"message": "User deleted",
	})
}

// GetUserList godoc
// @Summary Get list of users
// @Description This API gets list of users.
// @Tags Users
// @Security BearerAuth
// @Security ApiKeyAuth
// @Accept json
// @Produce json
// @Success 200 {object} map[string]interface{} "List of users"
// @Failure 500 {object} map[string]interface{} "Internal server error"
// @Router /api/v1/user [get]
func (uc *UserController) GetUserList(c echo.Context) error {
	result, err := uc.UserServices.GetUserList()
	if err != nil {
		return response.Error(c, http.StatusInternalServerError, err)
	}
	return response.Send(c, http.StatusOK, map[string]any{
		"data": result,
	})
}

// UpdatePassword godoc
// @Summary Update user password
// @Description This API updates user password.
// @Tags Users
// @Security BearerAuth
// @Security ApiKeyAuth
// @Accept json
// @Produce json
// @Param user body request.UpdatePasswordRequest true "User data"
// @Success 200 {object} map[string]interface{} "Password updated successfully"
// @Failure 400 {object} map[string]interface{} "Invalid request"
// @Failure 500 {object} map[string]interface{} "Internal server error"
// @Router /api/v1/user/password [put]
func (uc *UserController) UpdatePassword(c echo.Context) error {
	req := request.UpdatePasswordRequest{}
	if err := c.Bind(&req); err != nil {
		return response.Error(c, http.StatusBadRequest, err)
	}
	if err := c.Validate(req); err != nil {
		return response.Error(c, http.StatusBadRequest, err)
	}
	if err := uc.UserServices.UpdatePassword(req); err != nil {
		return response.Error(c, http.StatusInternalServerError, err)
	}
	return response.Send(c, http.StatusOK, map[string]any{
		"message": "Password updated successfully",
	})
}

// GetUserByEmail godoc
// @Summary Get user by email
// @Description This API fetches user by email.
// @Tags Users
// @Security BearerAuth
// @Security ApiKeyAuth
// @Accept json
// @Produce json
// @Param email query string true "User email"
// @Success 200 {object} map[string]interface{} "User details"
// @Failure 500 {object} map[string]interface{} "Internal server error"
// @Router /api/v1/user/email [get]
func (uc *UserController) GetUserByEmail(c echo.Context) error {
	email := c.QueryParam("email")
	result, err := uc.UserServices.GetUserByEmail(email)
	if err != nil {
		return response.Error(c, http.StatusInternalServerError, err)
	}
	return response.Send(c, http.StatusOK, map[string]any{
		"data": result,
	})
}

// GetBorrowedHistoryByUserID godoc
// @Summary Get borrowed history by user ID
// @Description This API fetches borrowed history by user ID.
// @Tags Users
// @Security BearerAuth
// @Security ApiKeyAuth
// @Accept json
// @Produce json
// @Param id path int true "User ID"
// @Success 200 {object} map[string]interface{} "Borrowed history"
// @Failure 500 {object} map[string]interface{} "Internal server error"
// @Router /api/v1/user/{id}/history [get]
func (uc *UserController) GetBorrowedHistoryByUserID(c echo.Context) error {
	id := c.Param("id")
	userID, err := strconv.Atoi(id)
	if err != nil {
		return response.Error(c, http.StatusBadRequest, err)
	}
	result, err := uc.UserServices.GetBorrowedHistoryByUserID(uint(userID))
	if err != nil {
		return response.Error(c, http.StatusInternalServerError, err)
	}
	return response.Send(c, http.StatusOK, map[string]any{
		"data": result,
	})
}
