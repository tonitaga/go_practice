package main

import (
	"fmt"
	"iter"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Book struct {
	ID     uint   `gorm:"primaryKey"`
	Name   string `gorm:"size:255;not null"`
	Author string `gorm:"size:255;not null"`
	Theme  string `gorm:"size:255"`
}

func DoQuery[T any](database *gorm.DB, query string) iter.Seq[T] {
	return func(yield func(T) bool) {
		storage := []T{}
		database.Raw(query).Scan(&storage)

		for _, value := range storage {
			if !yield(value) {
				return
			}
		}
	}
}

func main() {
	database, err := gorm.Open(sqlite.Open("books.db"), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	database.AutoMigrate(&Book{})

	books := []*Book{
		{Name: "Story #1", Author: "Author #1", Theme: "Story"},
		{Name: "Story #2", Author: "Author #3", Theme: "Story"},
		{Name: "Story #5", Author: "Author #1", Theme: "Story"},
		{Name: "True Detective", Author: "Author #6", Theme: "Detective"},
	}
	for _, book := range books {
		database.Create(book)
	}

	for book := range DoQuery[Book](database, `SELECT b.* FROM books b WHERE theme=="Story"`) {
		fmt.Println(book)
	}

}
