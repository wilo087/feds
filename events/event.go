package events

import (
	"context"

	"github.com/wilo087/feeds/models"
)

type EventStore interface {
	Close()
	PublishCreadtedFeed(ctx context.Context, feed *models.Feed) error
	SubscribeCreatedFeed(ctx context.Context) (<-chan *CreatedFeedMessage, error)
	OnCreatedFeed(handler func(CreatedFeedMessage)) error
}

var eventStore EventStore

func Close() {
	eventStore.Close()
}

func PublishCreadtedFeed(ctx context.Context, feed *models.Feed) error {
	return eventStore.PublishCreadtedFeed(ctx, feed)
}

func SubscribeCreatedFeed(ctx context.Context) (<-chan *CreatedFeedMessage, error) {
	return eventStore.SubscribeCreatedFeed(ctx)
}

func OnCreatedFeed(ctx context.Context, handler func(CreatedFeedMessage)) error {
	return eventStore.OnCreatedFeed(handler)
}
