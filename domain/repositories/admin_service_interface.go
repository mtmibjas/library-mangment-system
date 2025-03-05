package repositories

import "library-mngmt/domain/entities"

type AdminRepositoriesInterface interface {
	CreatePermission(permission entities.Permission) error
	CreateRolePermission(rolePermission entities.RolePermission) error
	UpdatePermission(permission entities.Permission) error
	DeletePermission(permission entities.Permission) error
	DeleteRolePermission(rolePermission entities.RolePermission) error
	GetPermissions() ([]entities.Permission, error)
	GetRoles() ([]entities.Role, error)
	GetPermissionsRoleID(roleID uint) ([]entities.RolePermission, error)
}

type AuthRepositoriesInterface interface {
}
