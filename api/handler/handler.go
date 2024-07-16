package handler

import (
	"api_service/config"
	pbuAuth "api_service/genproto/auth"
	pbuOrders "api_service/genproto/order"
	pbuProduct "api_service/genproto/product"
	"api_service/pkg"
	"api_service/pkg/logger"
	"log"
	"log/slog"
)

type Handler struct {
	ClientAuthentication pbuAuth.AuthClient
	ClientProduct        pbuProduct.ProductServiceClient
	ClientOrder          pbuOrders.OrderServiceClient
	Logger               *slog.Logger
}

func NewHandler(cfg *config.Config) *Handler {

	l, err := logger.New()
	if err != nil {
		log.Fatal("error: ", err)
	}
	return &Handler{
		ClientAuthentication: pkg.NewAuthenticationClient(cfg),
		ClientProduct:        pkg.NewPtoductClient(cfg),
		ClientOrder:          pkg.NewOrderClient(cfg),
		Logger:               l,
	}
}
