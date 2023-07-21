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

func (n *NatsEventStore) decodeMessage(data []byte, m interface{}) error {
	buffer := bytes.NewBuffer(data)
	return gob.NewDecoder(buffer).Decode(m)
}

func (n NatsEventStore) OnCreatedFeed(f func(CreatedFeedMessage)) (err error) {
	msg := CreatedFeedMessage{}
	n.feedCreatedSub, err = n.conn.Subscribe(msg.Type(), func(m *nats.Msg) {
		n.decodeMessage(m.Data, &msg)
		f(msg)
	})

	return
}

func (n *NatsEventStore) SubscribeCreatedFeed(ctx context.Context) (<-chan CreatedFeedMessage, error) {
	msg := CreatedFeedMessage{}
	n.feedCreatedCh = make(chan CreatedFeedMessage, 64)
	ch := make(chan *nats.Msg, 64)

	var err error
	n.feedCreatedSub, err = n.conn.ChanSubscribe(msg.Type(), ch)
	if err != nil {
		return nil, err
	}

	go func() {
		for {
			select {
			case m := <-ch:
				n.decodeMessage(m.Data, &msg)
				n.feedCreatedCh <- msg
			}
		}
	}()

	return (<-chan CreatedFeedMessage)(n.feedCreatedCh), nil
	// return n.feedCreatedCh, nil
}
