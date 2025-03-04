package usecases

import (
	"library-mngmt/app/http/request"
	"library-mngmt/domain/entities"
)

func (a *AdminService) CreatePermission(req request.CreatePermissionRequest) error {

	permission := entities.Permission{
		Action: req.Action,
	}
	if err := a.AdminRepository.CreatePermission(permission); err != nil {
		return err
	}
	return nil
}

func (a *AdminService) GetPermissions() ([]entities.Permission, error) {
	return a.AdminRepository.GetPermissions()
}
func (a *AdminService) UpdatePermission(req request.UpdatePermissionRequest) error {
	permission := entities.Permission{
		ID:     req.ID,
		Action: req.Action,
	}
	if err := a.AdminRepository.UpdatePermission(permission); err != nil {
		return err
	}
	return nil
}
func (a *AdminService) DeletePermission(id uint) error {
	permission := entities.Permission{
		ID: id,
	}
	if err := a.AdminRepository.DeletePermission(permission); err != nil {
		return err
	}
	return nil
}
func (a *AdminService) CreateRolePermission(req request.CreateRolePermissionRequest) error {
	rolePermission := entities.RolePermission{
		RoleID:       req.RoleID,
		PermissionID: req.PermissionID,
	}
	if err := a.AdminRepository.CreateRolePermission(rolePermission); err != nil {
		return err
	}
	return nil
}
func (a *AdminService) DeleteRolePermission(req request.RemoveRolePermissionRequest) error {
	rolePermission := entities.RolePermission{
		RoleID:       req.RoleID,
		PermissionID: req.PermissionID,
	}
	if err := a.AdminRepository.DeleteRolePermission(rolePermission); err != nil {
		return err
	}
	return nil
}
func (a *AdminService) GetRoles() ([]entities.Role, error) {
	return a.AdminRepository.GetRoles()
}
