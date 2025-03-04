package request

type CreateBookRequest struct {
	Title    string `json:"title" validate:"required"`
	Author   string `json:"author" validate:"required"`
	Category uint   `json:"category_id" validate:"required"`
}
type UpdateBookRequest struct {
	ID        uint   `json:"id" validate:"required"`
	Title     string `json:"title" validate:"required"`
	Author    string `json:"author" validate:"required"`
	Category  uint   `json:"category_id" validate:"required"`
	Available bool   `json:"available" validate:"required"`
}
type DeleteBookRequest struct {
	ID uint `json:"id" validate:"required"`
}
