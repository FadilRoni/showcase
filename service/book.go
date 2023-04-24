package service

import (
	"challange-8_2/model"
)

type BookService interface {
	AddBook(in model.Book) (res model.Book, err error)
	GetAllBook() (res []model.Book, err error)
	GetBookId(id int64) (res model.Book, err error)
	UpdateBook(in model.Book) (res model.Book, err error)
	DeleteBook(id int64) (err error)
}

func (s *Service) AddBook(in model.Book) (res model.Book, err error) {
	res, err = s.repo.AddBook(in)
	if err != nil {
		return res, err
	}

	return res, nil
}

func (s *Service) GetBookId(id int64) (res model.Book, err error) {
	return s.repo.GetBookId(id)
}

func (s *Service) GetAllBook() (res []model.Book, err error) {
	return s.repo.GetAllBook()
}

func (s *Service) UpdateBook(in model.Book) (res model.Book, err error) {
	return s.repo.UpdateBook(in)
}

func (s *Service) DeleteBook(id int64) (err error) {
	return s.repo.DeleteBook(id)
}
