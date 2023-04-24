package repository

import (
	"challange-8_2/model"
)

type BookRepo interface {
	AddBook(in model.Book) (res model.Book, err error)
	GetAllBook() (res []model.Book, err error)
	GetBookId(id int64) (res model.Book, err error)
	UpdateBook(in model.Book) (res model.Book, err error)
	DeleteBook(id int64) (err error)
}

func (r Repo) AddBook(in model.Book) (res model.Book, err error) {

	err = r.gorm.Create(&in).Scan(&res).Error
	if err != nil {
		return res, err
	}

	return res, nil
}

func (r Repo) GetBookId(id int64) (res model.Book, err error) {

	err = r.gorm.First(&res, id).Scan(&res).Error
	if err != nil {
		return res, err
	}

	return res, nil
}

func (r Repo) GetAllBook() (res []model.Book, err error) {

	err = r.gorm.Find(&res).Scan(&res).Error
	if err != nil {
		return res, err
	}

	return res, nil
}

func (r Repo) UpdateBook(in model.Book) (res model.Book, err error) {
	err = r.gorm.Model(&res).Where("id = ?", in.ID).Updates(model.Book{
		NameBook: in.NameBook,
		Author:   in.Author,
	}).Scan(&res).Error

	if err != nil {
		return res, nil
	}

	return res, nil

}

func (r Repo) DeleteBook(id int64) (err error) {

	in := model.Book{}

	err = r.gorm.Where("id = ?", id).Delete(&in).Error
	if err != nil {
		return err
	}

	return nil
}
