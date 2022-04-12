package models

import "errors"

type Book struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Author string `json:"author"`
	Qty    int    `json:"qty"`
}

var bookDatas = []Book{
	{ID: 1, Title: "Akuntansi Pengantar 1", Author: "Supardi", Qty: 2},
	{ID: 2, Title: "A-Z Psikologi : Berbagai kumpulan topik Psikologi", Author: "Zainul Anwar", Qty: 4},
	{ID: 3, Title: "Kesehjateraan Sosial", Author: "Isbandi Rukminto Adi", Qty: 1},
}

func GetAllBooks() []Book {
	return bookDatas
}

func AddNewBook(newBook *Book) {
	data := &bookDatas
	newId := (*data)[len(*data)-1].ID + 1
	newBook.ID = newId
	bookDatas = append(bookDatas, *newBook)
}

func GetBookById(id int) (*Book, error) {
	for i, data := range bookDatas {
		if data.ID == id {
			return &bookDatas[i], nil
		}
	}

	return nil, errors.New("Data not found")
}

func RemoveBookById(id int) (*Book, error) {
	for i, data := range bookDatas {
		if data.ID == id {
			newData := &bookDatas

			deletedData := (*newData)[i]

			(*newData)[i] = (*newData)[len(*newData)-1]
			*newData = (*newData)[:len(*newData)-1]

			return &deletedData, nil
		}
	}

	return nil, errors.New("Data cannot be deleted")
}

func UpdateBookById(id int, updateBook *Book) (*Book, error) {
	for i, data := range bookDatas {
		if data.ID == id {
			newData := &bookDatas

			oldData := (*newData)[i]
			updateBook.ID = oldData.ID
			(*newData)[i] = *updateBook

			return &oldData, nil
		}
	}

	return nil, errors.New("Data not found")
}
