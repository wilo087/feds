package events

type Message interface {
	Type() string
}

type CreatedFeedMessage struct {
	ID          string `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	CreatedAt   string `json:"created_at"`
	UpdatedAt   string `json:"updated_at"`
}

func (m CreatedFeedMessage) Type() string {
	return "feed.created"
}
