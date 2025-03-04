package entities

type Book struct {
	ID        uint   `json:"id"`
	Title     string `json:"title"`
	Author    string `json:"author"`
	Category  uint   `json:"category_id"`
	Available bool   `json:"available"`
}

type BorrowRecord struct {
	ID         uint   `json:"id"`
	UserID     uint   `json:"user_id"`
	BookID     uint   `json:"book_id"`
	BorrowedAt string `json:"borrowed_at"`
	ReturnedAt string `json:"returned_at"`
}
