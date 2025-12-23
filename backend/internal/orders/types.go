package orders

import (
	"context"

	repo "github.com/danielsonng/ecomgo/internal/adapters/postgresql/sqlc"
)

type orderItem struct {
	ProductID int64 `json:"productId"`
	Quantity  int32 `json:"Quantity"`
}

type createOrderParams struct {
	CustomerID int64       `json:"customerId`
	Items      []orderItem `json:"items"`
}

type Service interface {
	PlaceOrder(ctx context.Context, tempOrder createOrderParams) (repo.Order, error)
}
