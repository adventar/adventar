package usecase

import (
	"github.com/adventar/adventar/backend/pkg/gen/sqlc/adventar_db"
	"github.com/adventar/adventar/backend/pkg/infra"
	"github.com/adventar/adventar/backend/pkg/util"
)

type metaFetcher interface {
	Fetch(string) (*util.SiteMeta, error)
}

type Usecase struct {
	clients     *infra.Clients
	queries     *adventar_db.Queries
	metaFetcher metaFetcher
}

func New(clients *infra.Clients, metaFetcher metaFetcher) *Usecase {
	return &Usecase{
		clients:     clients,
		queries:     clients.DB().Queries(),
		metaFetcher: metaFetcher,
	}
}
