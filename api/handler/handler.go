package handler

import (
	"api_service/config"
	pbuAuth "api_service/genproto/auth"
)

type Handler struct {
	ClientAuthentication pbuAuth.AuthClient

	// ClientAuthentication pbAuthservice.AuthClient
	// ClientOrder          pbOrderservice.OrderServiceClient
	// ClientReservation    pbReservation.ReservationServiceClient
	// Menu                 menu.MenuServiceClient
	// Payments             payments.PaymentsClient
	// Restaurant           restaurant.RestaurantClient
	// Logger               *slog.Logger
}

func NewHandler(cfg *config.Config) *Handler {
	return nil
}
