package repositories

import "library-mngmt/domain/entities"

func (r *BookRepository) CreateBook(book entities.Book) error {
	_, err := r.Database.Exec("INSERT INTO books (title, author, category_id, available) VALUES ($1, $2, $3, $4)",
		book.Title, book.Author, book.Category, book.Available)
	if err != nil {
		return err
	}
	return nil
}

// --- need to check updated values
func (r *BookRepository) UpdateBook(book entities.Book) error {
	_, err := r.Database.Exec("UPDATE books SET title = $1, author = $2, year = $3 WHERE id = $4",
		book.Title, book.Author, book.Category, book.Available, book.ID)
	if err != nil {
		return err
	}
	return nil
}
func (r *BookRepository) DeleteBook(book entities.Book) error {
	_, err := r.Database.Exec("DELETE FROM books WHERE id = $1", book.ID)
	if err != nil {
		return err
	}
	return nil
}
func (r *BookRepository) GetBookByID(id uint) (entities.Book, error) {
	var book entities.Book
	row := r.Database.QueryRow("SELECT * FROM books WHERE id = $1", id)
	err := row.Scan(&book.ID, &book.Title, &book.Author, &book.Category, &book.Available)
	if err != nil {
		return book, err
	}
	return book, nil
}
func (r *BookRepository) GetBookList() ([]entities.Book, error) {
	var books []entities.Book
	rows, err := r.Database.Query("SELECT * FROM books")
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var book entities.Book
		err := rows.Scan(&book.ID, &book.Title, &book.Author, &book.Category, &book.Available)
		if err != nil {
			return nil, err
		}
		books = append(books, book)
	}
	return books, nil
}
func (r *BookRepository) GetBorrowedHistoryByBookID(id uint) ([]entities.BorrowRecord, error) {
	var borrowedHistories []entities.BorrowRecord
	rows, err := r.Database.Query("SELECT * FROM borrowed_histories WHERE book_id = $1", id)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var borrowedHistory entities.BorrowRecord
		err := rows.Scan(&borrowedHistory.ID, &borrowedHistory.BookID, &borrowedHistory.UserID, &borrowedHistory.BorrowedAt, &borrowedHistory.ReturnedAt)
		if err != nil {
			return nil, err
		}
		borrowedHistories = append(borrowedHistories, borrowedHistory)
	}
	return borrowedHistories, nil
}
