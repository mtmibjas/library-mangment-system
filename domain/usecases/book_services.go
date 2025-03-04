package usecases

import (
	"library-mngmt/app/http/request"
	"library-mngmt/domain/entities"
)

func (s *BookService) CreateBook(req request.CreateBookRequest) error {

	book := entities.Book{
		Title:     req.Title,
		Author:    req.Author,
		Category:  req.Category,
		Available: true,
	}
	if err := s.BookRepository.CreateBook(book); err != nil {
		return err
	}
	return nil
}
func (s *BookService) UpdateBook(req request.UpdateBookRequest) error {
	book := entities.Book{
		ID:        req.ID,
		Title:     req.Title,
		Author:    req.Author,
		Available: req.Available,
	}
	if err := s.BookRepository.UpdateBook(book); err != nil {
		return err
	}
	return nil
}
func (s *BookService) DeleteBook(req request.DeleteBookRequest) error {
	book := entities.Book{
		ID: req.ID,
	}
	if err := s.BookRepository.DeleteBook(book); err != nil {
		return err
	}
	return nil
}
func (s *BookService) GetBook(id uint) (entities.Book, error) {
	return s.BookRepository.GetBookByID(id)
}
func (s *BookService) GetBookList() ([]entities.Book, error) {
	return s.BookRepository.GetBookList()
}
func (s *BookService) GetBorrowedHistoryByBookID(id uint) ([]entities.BorrowRecord, error) {
	return s.BookRepository.GetBorrowedHistoryByBookID(id)
}
