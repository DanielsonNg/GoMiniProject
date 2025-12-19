package products

import (
	"context"

	tutorial "github.com/danielsonng/ecomgo/internal/adapters/postgresql/sqlc"
)

type Service interface {
	ListProducts(ctx context.Context) ([]tutorial.Product, error)
	GetProductById(ctx context.Context, id int64) (tutorial.Product, error)
}

type svc struct {
	// repo
	repo tutorial.Querier
}

func NewService(repo tutorial.Querier) Service {
	return &svc{repo: repo}
}

func (s *svc) ListProducts(ctx context.Context) ([]tutorial.Product, error) {
	return s.repo.ListProducts(ctx)
}

func (s *svc) GetProductById(ctx context.Context, id int64) (tutorial.Product, error) {
	return s.repo.FindProductByID(ctx, id)
}
