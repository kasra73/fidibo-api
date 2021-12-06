package search_service

import (
	"github.com/kasra73/fidibo-api/models"
)

type BookSearcher interface {
	Search(q string) ([]models.Book, error)
}

func GetBookSearcher() BookSearcher {
	return &FidiboBookSearch{}
}
