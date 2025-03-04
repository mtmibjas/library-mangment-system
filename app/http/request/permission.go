package request

type CreatePermissionRequest struct {
	Action string `json:"action" validate:"required"`
}
type UpdatePermissionRequest struct {
	ID     uint   `json:"id" validate:"required"`
	Action string `json:"action" validate:"required"`
}
type CreateRolePermissionRequest struct {
	RoleID       uint `json:"role_id" validate:"required"`
	PermissionID uint `json:"permission_id" validate:"required"`
}
type RemoveRolePermissionRequest struct {
	RoleID       uint `json:"role_id" validate:"required"`
	PermissionID uint `json:"permission_id" validate:"required"`
}
