package repositories

import (
	"library-mngmt/domain/entities"
)

func (ar *AdminRepository) CreatePermission(permission entities.Permission) error {
	query := `INSERT INTO permissions (action) VALUES ($1)`
	_, err := ar.Database.Exec(query, permission.Action)
	if err != nil {
		return err
	}
	return nil
}
func (ar *AdminRepository) CreateRolePermission(rolePermission entities.RolePermission) error {
	query := `INSERT INTO role_permissions (role_id, permission_id) VALUES ($1, $2)`
	_, err := ar.Database.Exec(query, rolePermission.RoleID, rolePermission.PermissionID)
	if err != nil {
		return err
	}
	return nil
}
func (ar *AdminRepository) UpdatePermission(permission entities.Permission) error {
	query := `UPDATE permissions SET action = $1 WHERE id = $2`
	_, err := ar.Database.Exec(query, permission.Action, permission.ID)
	if err != nil {
		return err
	}
	return nil
}
func (ar *AdminRepository) DeletePermission(permission entities.Permission) error {
	query := `DELETE FROM permissions WHERE id = $1`
	_, err := ar.Database.Exec(query, permission.ID)
	if err != nil {
		return err
	}
	return nil
}
func (ar *AdminRepository) DeleteRolePermission(rolePermission entities.RolePermission) error {
	query := `DELETE FROM role_permissions WHERE role_id = $1 AND permission_id = $2`
	_, err := ar.Database.Exec(query, rolePermission.RoleID, rolePermission.PermissionID)
	if err != nil {
		return err
	}
	return nil
}
func (ar *AdminRepository) GetPermissions() ([]entities.Permission, error) {
	var permissions []entities.Permission
	query := `SELECT * FROM permissions`
	rows, err := ar.Database.Query(query)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var permission entities.Permission
		err := rows.Scan(&permission.ID, &permission.Action)
		if err != nil {
			return nil, err
		}
		permissions = append(permissions, permission)
	}
	return permissions, nil
}
func (ar *AdminRepository) GetPermissionsRoleID(roleID uint) ([]entities.RolePermission, error) {
	var permissions []entities.RolePermission
	query := `SELECT rp.id, rp.role_id, rp.permission_id, p.action
	 FROM role_permissions as rp inner join permissions as p on p.id = rp.permission_id  WHERE rp.role_id = $1 `
	rows, err := ar.Database.Query(query, roleID)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var permission entities.RolePermission
		err := rows.Scan(&permission.ID, &permission.RoleID, &permission.PermissionID, &permission.Action)
		if err != nil {
			return nil, err
		}
		permissions = append(permissions, permission)
	}
	return permissions, nil
}
func (ar *AdminRepository) GetRoles() ([]entities.Role, error) {
	var roles []entities.Role
	query := `SELECT * FROM roles`
	rows, err := ar.Database.Query(query)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var role entities.Role
		err := rows.Scan(&role.ID, &role.Name)
		if err != nil {
			return nil, err
		}
		roles = append(roles, role)
	}
	return roles, nil
}
