package client

import (
	"context"

	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
	"github.com/tinchourteaga/go-grpc-order-svc/internal/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type ProductServiceClient struct {
	Client pb.ProductServiceClient
}

func NewProductServiceClient() pb.ProductServiceClient {
	connection, err := grpc.Dial(viper.GetString("PRODUCT_SVC_URL"), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Error().Msg("could not establish connection: " + err.Error())
		return nil
	}

	return pb.NewProductServiceClient(connection)
}

func (svc *ProductServiceClient) FindOne(productId int64) (*pb.FindOneResponse, error) {
	req := &pb.FindOneRequest{
		Id: productId,
	}

	return svc.Client.FindOne(context.Background(), req)
}

func (svc *ProductServiceClient) DecreaseStock(productId int64, orderId int64) (*pb.DecreaseStockResponse, error) {
	req := &pb.DecreaseStockRequest{
		Id:      productId,
		OrderId: orderId,
	}

	return svc.Client.DecreaseStock(context.Background(), req)
}
