package controller

import (
	"library-mngmt/app/container"
	"library-mngmt/app/http/request"
	"library-mngmt/app/http/response"
	"library-mngmt/domain/usecases"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type BookController struct {
	Adapters    *container.Container
	BookService *usecases.BookService
}

func NewBookController(ctr *container.Container) *BookController {
	return &BookController{
		Adapters:    ctr,
		BookService: usecases.NewBookService(ctr),
	}
}

// CreateBook godoc
// @Summary Create a new book
// @Description This API creates a new book.
// @Tags Books
// @Security BearerAuth
// @Security ApiKeyAuth
// @Accept json
// @Produce json
// @Param book body request.CreateBookRequest true "Book data"
// @Success 201 {object} map[string]interface{} "Book created successfully"
// @Failure 400 {object} map[string]interface{} "Invalid request"
// @Failure 500 {object} map[string]interface{} "Internal server error"
// @Router /api/v1/book [post]
func (uc *BookController) CreateBook(c echo.Context) error {
	req := request.CreateBookRequest{}
	if err := c.Bind(&req); err != nil {
		return response.Error(c, http.StatusBadRequest, err)
	}
	if err := c.Validate(req); err != nil {
		return response.Error(c, http.StatusBadRequest, err)
	}
	if err := uc.BookService.CreateBook(req); err != nil {
		return response.Error(c, http.StatusInternalServerError, err)
	}
	return response.Send(c, http.StatusCreated, map[string]any{
		"message": "Book created successfully",
	})
}

// GetBook godoc
// @Summary  Get book details
// @Description This API to Get book details.
// @Tags Books
// @Security BearerAuth
// @Security ApiKeyAuth
// @Accept json
// @Produce json
// @Param id path int true "Book ID"
// @Success 200 {object} map[string]interface{} "Book details"
// @Failure 400 {object} map[string]interface{} "Invalid request"
// @Failure 500 {object} map[string]interface{} "Internal server error"
// @Router /api/v1/book/{id} [get]
func (uc *BookController) GetBook(c echo.Context) error {
	id := c.Param("id")
	bookID, err := strconv.Atoi(id)
	if err != nil {
		return response.Error(c, http.StatusBadRequest, err)
	}
	result, err := uc.BookService.GetBook(uint(bookID))
	if err != nil {
		return response.Error(c, http.StatusInternalServerError, err)
	}
	return response.Send(c, http.StatusOK, map[string]any{
		"data": result,
	})
}

// GetBookList godoc
// @Summary  Get list of books
// @Description This API to Get list of books.
// @Tags Books
// @Security BearerAuth
// @Security ApiKeyAuth
// @Accept json
// @Produce json
// @Success 200 {object} map[string]interface{} "List of books"
// @Failure 500 {object} map[string]interface{} "Internal server error"
// @Router /api/v1/book [get]
func (uc *BookController) GetBookList(c echo.Context) error {
	result, err := uc.BookService.GetBookList()
	if err != nil {
		return response.Error(c, http.StatusInternalServerError, err)
	}
	return response.Send(c, http.StatusOK, map[string]any{
		"data": result,
	})
}

// UpdateBook godoc
// @Summary  Update book
// @Description This API to update book.
// @Tags Books
// @Security BearerAuth
// @Security ApiKeyAuth
// @Accept json
// @Produce json
// @Param id path int true "Book ID"
// @Param book body request.UpdateBookRequest true "Book data"
// @Success 200 {object} map[string]interface{} "Book updated successfully"
// @Failure 400 {object} map[string]interface{} "Invalid request"
// @Failure 500 {object} map[string]interface{} "Internal server error"
// @Router /api/v1/book/{id} [put]
func (uc *BookController) UpdateBook(c echo.Context) error {
	req := request.UpdateBookRequest{}
	if err := c.Bind(&req); err != nil {
		return response.Error(c, http.StatusBadRequest, err)
	}
	if err := c.Validate(req); err != nil {
		return response.Error(c, http.StatusBadRequest, err)
	}
	if err := uc.BookService.UpdateBook(req); err != nil {
		return response.Error(c, http.StatusInternalServerError, err)
	}
	return response.Send(c, http.StatusOK, map[string]any{
		"message": "Book updated successfully",
	})
}

// DeleteBook godoc
// @Summary  Delete book
// @Description This API to delete book.
// @Tags Books
// @Security BearerAuth
// @Security ApiKeyAuth
// @Accept json
// @Produce json
// @Param id path int true "Book ID"
// @Success 200 {object} map[string]interface{} "Book deleted successfully"
// @Failure 400 {object} map[string]interface{} "Invalid request"
// @Failure 500 {object} map[string]interface{} "Internal server error"
// @Router /api/v1/book/{id} [delete]
func (uc *BookController) DeleteBook(c echo.Context) error {
	id := c.Param("id")
	bookID, err := strconv.Atoi(id)
	if err != nil {
		return response.Error(c, http.StatusBadRequest, err)
	}
	if err := uc.BookService.DeleteBook(request.DeleteBookRequest{ID: uint(bookID)}); err != nil {
		return response.Error(c, http.StatusInternalServerError, err)
	}
	return response.Send(c, http.StatusOK, map[string]any{
		"message": "Book deleted successfully",
	})
}

// GetBorrowedHistoryByBookID godoc
// @Summary  Get borrowed history by book ID
// @Description This API to Get borrowed history by book ID.
// @Tags Books
// @Security BearerAuth
// @Security ApiKeyAuth
// @Accept json
// @Produce json
// @Param id path int true "Book ID"
// @Success 200 {object} map[string]interface{} "Borrowed history"
// @Failure 400 {object} map[string]interface{} "Invalid request"
// @Failure 500 {object} map[string]interface{} "Internal server error"
// @Router /api/v1/book/{id}/history [get]
func (uc *BookController) GetBorrowedHistoryByBookID(c echo.Context) error {
	id := c.Param("id")
	bookID, err := strconv.Atoi(id)
	if err != nil {
		return response.Error(c, http.StatusBadRequest, err)
	}
	result, err := uc.BookService.GetBorrowedHistoryByBookID(uint(bookID))
	if err != nil {
		return response.Error(c, http.StatusInternalServerError, err)
	}
	return response.Send(c, http.StatusOK, map[string]any{
		"data": result,
	})
}
