package search

import (
	"context"

	"github.com/wilo087/feeds/models"
)

type SearchRepository interface {
	Close()
	InsertFeed(ctx context.Context, feed *models.Feed) error
	SearchFeed(ctx context.Context, query string) ([]*models.Feed, error)
}
