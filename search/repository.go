package search

import (
	"context"

	"github.com/wilo087/feeds/models"
)

type SearchRepository interface {
	Close()
	InsertFeed(ctx context.Context, feed *models.Feed) error
	SearchFeed(ctx context.Context, query string) ([]models.Feed, error)
}

var repo SearchRepository

func SetRepository(r SearchRepository) {
	repo = r
}

func Close() {
	repo.Close()
}

func InsertFeed(ctx context.Context, feed *models.Feed) error {
	return repo.InsertFeed(ctx, feed)
}

func SearchFeed(ctx context.Context, query string) ([]models.Feed, error) {
	return repo.SearchFeed(ctx, query)
}
