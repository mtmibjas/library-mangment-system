package repositories

import "library-mngmt/domain/entities"

type BookRepositoriesInterface interface {
	CreateBook(book entities.Book) error
	UpdateBook(book entities.Book) error
	DeleteBook(book entities.Book) error
	GetBookByID(id uint) (entities.Book, error)
	GetBookList() ([]entities.Book, error)
	GetBorrowedHistoryByBookID(id uint) ([]entities.BorrowRecord, error)
}
