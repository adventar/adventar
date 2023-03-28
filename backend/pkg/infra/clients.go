package infra

import "github.com/adventar/adventar/backend/pkg/infra/db"

type Clients struct {
	db db.Client
}

func New(options ...Option) *Clients {
	clients := &Clients{}

	for _, opt := range options {
		opt(clients)
	}

	return clients
}

func (x *Clients) DB() db.Client {
	if x.db == nil {
		panic("DB client is not configured, but called")
	}
	return x.db
}

// Option provides functional option pattern
type Option func(c *Clients)

func WithDB(dbClient db.Client) Option {
	return func(c *Clients) {
		c.db = dbClient
	}
}
