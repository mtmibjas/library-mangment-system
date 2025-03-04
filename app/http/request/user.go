package request

type CreateUserRequest struct {
	Name     string `json:"name" validate:"required"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
	RoleID   uint   `json:"role_id" validate:"required"`
}

type UpdateUserRequest struct {
	ID     uint   `json:"id" validate:"required"`
	Name   string `json:"name" validate:"required"`
	Email  string `json:"email" validate:"required,email"`
	RoleID uint   `json:"role_id" validate:"required"`
}
type UpdatePasswordRequest struct {
	ID       uint   `json:"id" validate:"required"`
	Password string `json:"password" validate:"required"`
}
