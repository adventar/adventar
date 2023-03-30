package usecase

import (
	"github.com/adventar/adventar/backend/pkg/gen/sqlc/adventar_db"
	"github.com/adventar/adventar/backend/pkg/infra"
)

type Usecase struct {
	clients *infra.Clients
	queries *adventar_db.Queries
}

func New(clients *infra.Clients) *Usecase {
	return &Usecase{
		clients: clients,
		queries: clients.DB().Queries(),
	}
}
