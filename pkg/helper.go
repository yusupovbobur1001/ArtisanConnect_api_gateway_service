package pkg

import (
	"api_service/config"
	pbuAuthservice "api_service/genproto/auth"
	pbuOrderservice "api_service/genproto/proto"
	pbuProduct "api_service/genproto/product"


	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func NewAuthenticationClient(cfg *config.Config) pbuAuthservice.AuthClient {
	conn, err := grpc.NewClient("localhost"+cfg.AUTH_SERVICE_PORT,
		grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Println("error while connecting authentication service ", err)
	}

	return pbuAuthservice.NewAuthClient(conn)
}

func NewOrderClient(cfg *config.Config) pbuProduct. {
	conn, err := grpc.NewClient("localhost"+cfg.RESERVATION_SERVICE,
		grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Println("erro while connecting order service ", err)
	}

	return pbuProduct.NewOrderServiceClient(conn)
}








