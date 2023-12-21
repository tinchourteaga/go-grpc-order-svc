package order

import (
	"context"

	"github.com/tinchourteaga/go-grpc-order-svc/internal/models"
	"github.com/tinchourteaga/go-grpc-order-svc/pkg/db"
)

type Repository interface {
	Create(context.Context, models.Order) (int64, error)
	Delete(context.Context, int64) error
}

type repository struct {
	con db.Connector
}

func NewRepository(con db.Connector) Repository {
	return &repository{
		con: con,
	}
}

func (r *repository) Create(ctx context.Context, order models.Order) (int64, error) {
	r.con.DB.Create(&order)
	return order.Id, nil
}

func (r *repository) Delete(ctx context.Context, orderId int64) error {
	result := r.con.DB.Delete(&models.Order{}, orderId)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
