package controller

import (
	"library-mngmt/app/container"
	"library-mngmt/app/http/request"
	"library-mngmt/app/http/response"
	"library-mngmt/domain/usecases"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type AdminController struct {
	Adapters      *container.Container
	AdminServices *usecases.AdminService
}

func NewAdminController(ctr *container.Container) *AdminController {
	return &AdminController{
		Adapters:      ctr,
		AdminServices: usecases.NewAdminService(ctr),
	}
}

// CreatePermission godoc
// @Summary Create a new permission
// @Description This API creates a new permission.
// @Tags admin
// @Accept json
// @Produce json
// @Param permission body request.CreatePermissionRequest true "Permission data"
// @Success 201 {object} map[string]interface{} "Permission created successfully"
// @Failure 400 {object} map[string]interface{} "Invalid request"
// @Failure 500 {object} map[string]interface{} "Internal server error"
// @Router /admin/v1/permission [post]
func (ac *AdminController) CreatePermission(c echo.Context) error {
	req := request.CreatePermissionRequest{}
	if err := c.Bind(&req); err != nil {
		return response.Error(c, http.StatusBadRequest, err)
	}

	if err := c.Validate(req); err != nil {
		return response.Error(c, http.StatusBadRequest, err)
	}
	if err := ac.AdminServices.CreatePermission(req); err != nil {
		return response.Error(c, http.StatusInternalServerError, err)
	}
	return response.Send(c, http.StatusCreated, map[string]any{
		"message": "Permission created successfully",
	})
}

// GetPermissions godoc
// @Summary Get all permissions
// @Description This API fetches all permissions.
// @Tags admin
// @Accept json
// @Produce json
// @Success 200 {object} map[string]interface{} "List of permissions"
// @Failure 500 {object} map[string]interface{} "Internal server error"
// @Router /admin/v1/permissions [get]
func (a *AdminController) GetPermissions(c echo.Context) error {

	result, err := a.AdminServices.GetPermissions()
	if err != nil {
		return response.Error(c, http.StatusInternalServerError, err)
	}
	return response.Send(c, http.StatusOK, map[string]any{
		"data": result,
	})
}

// UpdatePermission godoc
// @Summary Update a permission
// @Description This API updates a permission.
// @Tags admin
// @Accept json
// @Produce json
// @Param permission body request.UpdatePermissionRequest true "Permission data"
// @Success 200 {object} map[string]interface{} "Permission updated successfully"
// @Failure 400 {object} map[string]interface{} "Invalid request"
// @Failure 500 {object} map[string]interface{} "Internal server error"
// @Router /admin/v1/permission [put]
func (a *AdminController) UpdatePermission(c echo.Context) error {
	req := request.UpdatePermissionRequest{}
	if err := c.Bind(&req); err != nil {
		return response.Error(c, http.StatusBadRequest, err)
	}
	if err := c.Validate(req); err != nil {
		return response.Error(c, http.StatusBadRequest, err)
	}
	if err := a.AdminServices.UpdatePermission(req); err != nil {
		return response.Error(c, http.StatusInternalServerError, err)
	}
	return response.Send(c, http.StatusOK, map[string]any{
		"message": "Permission updated successfully",
	})
}

// DeletePermission godoc
// @Summary Delete a permission
// @Description This API deletes a permission.
// @Tags admin
// @Accept json
// @Produce json
// @Param id path int true "Permission ID"
// @Success 200 {object} map[string]interface{} "Permission deleted"
// @Failure 400 {object} map[string]interface{} "Invalid request"
// @Failure 500 {object} map[string]interface{} "Internal server error"
// @Router /admin/v1/permission/{id} [delete]
func (a *AdminController) DeletePermission(c echo.Context) error {
	id := c.Param("id")
	permissionID, err := strconv.Atoi(id)
	if err != nil {
		return response.Error(c, http.StatusBadRequest, err)
	}
	if err := a.AdminServices.DeletePermission(uint(permissionID)); err != nil {
		return response.Error(c, http.StatusInternalServerError, err)
	}
	return response.Send(c, http.StatusOK, map[string]any{
		"message": "Permission deleted",
	})
}

// CreateRole godoc
// @Summary Create a new role
// @Description This API creates a new role.
// @Tags admin
// @Accept json
// @Produce json
// @Param role body request.CreateRolePermissionRequest true "Role data"
// @Success 201 {object} map[string]interface{} "Role created successfully"
// @Failure 400 {object} map[string]interface{} "Invalid request"
// @Failure 500 {object} map[string]interface{} "Internal server error"
// @Router /admin/v1/add [post]
func (a *AdminController) AddPermissionToRole(c echo.Context) error {
	req := request.CreateRolePermissionRequest{}
	if err := c.Bind(&req); err != nil {
		return response.Error(c, http.StatusBadRequest, err)
	}
	if err := c.Validate(req); err != nil {
		return response.Error(c, http.StatusBadRequest, err)
	}
	return response.Send(c, http.StatusOK, map[string]any{
		"message": "Permission added to role",
	})
}

// RemovePermissionFromRole godoc
// @Summary Remove a permission from a role
// @Description This API removes a permission from a role.
// @Tags admin
// @Accept json
// @Produce json
// @Param role body request.RemoveRolePermissionRequest true "Role data"
// @Success 200 {object} map[string]interface{} "Permission removed from role"
// @Failure 400 {object} map[string]interface{} "Invalid request"
// @Failure 500 {object} map[string]interface{} "Internal server error"
// @Router /admin/v1/remove [patch]
func (a *AdminController) RemovePermissionFromRole(c echo.Context) error {
	req := request.RemoveRolePermissionRequest{}
	if err := c.Bind(&req); err != nil {
		return response.Error(c, http.StatusBadRequest, err)
	}
	if err := c.Validate(req); err != nil {
		return response.Error(c, http.StatusBadRequest, err)
	}
	return response.Send(c, http.StatusOK, map[string]any{
		"message": "Permission removed from role",
	})
}

// GetRoles godoc
// @Summary Get all roles
// @Description This API fetches all roles.
// @Tags admin
// @Accept json
// @Produce json
// @Success 200 {object} map[string]interface{} "List of roles"
// @Failure 500 {object} map[string]interface{} "Internal server error"
// @Router /admin/v1/roles [get]
func (a *AdminController) GetRoles(c echo.Context) error {

	result, err := a.AdminServices.GetRoles()
	if err != nil {
		return response.Error(c, http.StatusInternalServerError, err)
	}
	return response.Send(c, http.StatusOK, map[string]any{
		"data": result,
	})
}
