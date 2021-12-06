package search_service

import (
	"errors"

	"github.com/kasra73/fidibo-api/models"
)

type FidiboBookSearch struct {
	URL string
}

func (f *FidiboBookSearch) Search(q string) ([]models.Book, error) {
	// @todo
	// resp, err := http.Get("http://gobyexample.com")

	return nil, errors.New("not implemented")
}
