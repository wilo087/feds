package events

import (
	"bytes"
	"context"
	"encoding/gob"

	"github.com/nats-io/nats.go"
	"github.com/wilo087/feeds/models"
)

type NatsEventStore struct {
	conn           *nats.Conn
	feedCreatedSub *nats.Subscription
	feedCreatedCh  chan CreatedFeedMessage
}

func NewNatsEventStore(url string) (*NatsEventStore, error) {
	conn, err := nats.Connect(url)
	if err != nil {
		return nil, err
	}

	return &NatsEventStore{conn: conn}, nil
}

func (n *NatsEventStore) Close() {
	if n.conn != nil {
		n.conn.Close()
	}

	if n.feedCreatedSub != nil {
		n.feedCreatedSub.Unsubscribe()
	}
}

func (n *NatsEventStore) encodeMessage(msg Message) ([]byte, error) {
	buffer := bytes.Buffer{}
	err := gob.NewEncoder(&buffer).Encode(msg)

	if err != nil {
		return nil, err
	}

	return buffer.Bytes(), nil
}

func (n *NatsEventStore) PublishCreadtedFeed(ctx context.Context, feed *models.Feed) error {
	msg := CreatedFeedMessage{
		ID:          feed.ID,
		Title:       feed.Title,
		Description: feed.Description,
		CreatedAt:   feed.CreatedAt,
		UpdatedAt:   feed.UpdatedAt,
	}

	data, err := n.encodeMessage(msg)
	if err != nil {
		return err
	}

	return n.conn.Publish(msg.Type(), data)
}
