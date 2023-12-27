package order

import (
	"context"

	"github.com/tinchourteaga/go-grpc-order-svc/internal/client"
	"github.com/tinchourteaga/go-grpc-order-svc/internal/models"
)

type Service interface {
	Create(context.Context, models.OrderReqInfo) (int64, error)
}

type service struct {
	repository    Repository
	productClient client.ProductServiceClient
}

func NewService(repo Repository, prodClient client.ProductServiceClient) Service {
	return &service{
		repository:    repo,
		productClient: prodClient,
	}
}

func (s *service) Create(ctx context.Context, orderReqInfo models.OrderReqInfo) (int64, error) {
	product, err := s.productClient.FindOne(orderReqInfo.ProductId)
	if err != nil {
		return product.Status, err
	}

	order := models.Order{
		ProductId: product.Data.Id,
		Price:     product.Data.Price,
		UserId:    orderReqInfo.UserId,
	}

	id, err := s.repository.Create(ctx, order)
	if err != nil {
		return 0, err
	}

	res, err := s.productClient.DecreaseStock(orderReqInfo.ProductId, id)
	if err != nil {
		s.repository.Delete(ctx, id)
		return res.Status, err
	}

	return id, nil
}
