package order

import (
	"context"
	"net/http"

	"github.com/tinchourteaga/go-grpc-order-svc/internal/models"
	pb "github.com/tinchourteaga/go-grpc-order-svc/internal/pb/order"
)

type Order struct {
	orderSvc Service
}

func NewOrder(svc Service) *Order {
	return &Order{
		orderSvc: svc,
	}
}

func (o *Order) Create(ctx context.Context, req *pb.CreateRequest) (*pb.CreateResponse, error) {
	orderReqInfo := models.OrderReqInfo{
		ProductId: req.ProductId,
		Quantity:  req.Quantity,
		UserId:    req.UserId,
	}

	id, err := o.orderSvc.Create(ctx, orderReqInfo)
	if err != nil {
		return &pb.CreateResponse{
			Status: http.StatusConflict,
			Error:  err.Error(),
		}, nil
	}

	return &pb.CreateResponse{
		Status: http.StatusCreated,
		Id:     id,
	}, nil
}
