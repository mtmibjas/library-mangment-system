package entities

type Permission struct {
	ID     uint   `json:"id"`
	Action string `json:"action"`
}

type RolePermission struct {
	ID           uint   `json:"id"`
	RoleID       uint   `json:"role_id"`
	PermissionID uint   `json:"permission_id"`
	Action       string `json:"action"`
}
