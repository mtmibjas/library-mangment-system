package entities

type Permission struct {
	ID     uint   `json:"primaryKey"`
	Action string `json:"not null"`
}

type RolePermission struct {
	ID           uint `json:"primaryKey"`
	RoleID       uint `json:"not null"`
	PermissionID uint `json:"not null"`
}
