package handler

import (
	pb "api_service/genproto/orders"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)


// @Summary Update an order
// @Description Update the status of an order
// @Tags orders
// @Accept json
// @Produce json
// @Param id path string true "Order ID"
// @Security ApiKeyAuth
// @Param order body orders.UpdateOrderStatusRequest true "Order Update Info"
// @Success 200 {object} orders.UpdateOrderRespons
// @Failure 400 {object} string
// @Failure 500 {object} string
// @Router /orders/{id} [put]
func (h *Handler) UpdateOrder(c *gin.Context) {
	id := c.Param("id")

	if _, err := uuid.Parse(id); err != nil {
		h.Logger.Info("Error id invalid", "err: ", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "StatusBadRequest",
			"message": fmt.Sprintf("Error with getting Id from URL: %s", err.Error()),
		})
		return
	}
	order := pb.UpdateOrderStatusRequest{}

	err := json.NewDecoder(c.Request.Body).Decode(&order)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "StatusBadRequest",
			"message": fmt.Sprintf("Error NewDecoder or Decoder: %s", err.Error()),
		})
		return
	}

	order.OrderId = id
	if len(order.Status) == 0 {
		h.Logger.Info("Error Status not fount", "err: ", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "StatusBadRequest",
			"message": fmt.Sprintf("Status not fount, err: %v", err),
		})
		log.Fatal("Status not fount")
		return
	}

	resq, err := h.ClientOrder.UpdateOrderStatus(c, &order)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{
			"Error":   err.Error(),
			"Message": "Error while updating invitation",
		})
		log.Println("Error while creating invitation ", err)
		return
	}

	c.JSON(http.StatusOK, resq)
}


// @Summary Delete an order
// @Description Cancel an order
// @Tags orders
// @Accept json
// @Produce json
// @Param id path string true "Order ID"
// @Security ApiKeyAuth
// @Success 200 {object} orders.CancelOrder1
// @Failure 400 {object} string
// @Failure 500 {object} string
// @Router /orders/{id} [delete]
func (h *Handler) DeleteOrder(c *gin.Context) {
	id := c.Param("id")

	if _, err := uuid.Parse(id); err != nil {
		h.Logger.Info("Error id invalid", "err: ", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "StatusBadRequest",
			"message": fmt.Sprintf("Error with getting Id from URL: %s", err.Error()),
		})
		return
	}

	order := pb.Id{}
	order.OrderId = id

	resp, err := h.ClientOrder.CancelOrder(c, &order)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{
			"Error":   err.Error(),
			"Message": "Error while creating invitation",
		})
		log.Println("Error while creating invitation ", err)
		return
	}

	c.JSON(http.StatusOK, resp)
}


// @Summary Create an order
// @Description Create a new order
// @Tags orders
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param order body orders.CreateOrderRequest true "Order Create Info"
// @Success 201 {object} orders.OrderInfo
// @Failure 400 {object} string
// @Failure 500 {object} string
// @Router /orders [post]
func (h *Handler) CreateOrder(c *gin.Context) {
	order := pb.CreateOrderRequest{}

	err := json.NewDecoder(c.Request.Body).Decode(&order)
	if err != nil {
		h.Logger.Info("jsondan malumot xato o`qildi")
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "StatusBadRequest",
			"message": fmt.Sprintf("jsondan malumot xato o`qildi: %s", err.Error()),
		})
		return
	}

	resp, err := h.ClientOrder.CreateOrder(c, &order)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{
			"Error":   err.Error(),
			"Message": "Error while creating invitation",
		})
		log.Println("Error while creating invitation ", err)
		return
	}

	c.JSON(http.StatusCreated, resp)
}

// @Summary List orders
// @Description Get a list of orders with pagination
// @Tags orders
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param page path int true "Page number"
// @Param limit path int true "Limit number"
// @Success 200 {object} orders.ListOrdersResponse
// @Failure 400 {object} string
// @Failure 500 {object} string
// @Router /orders/{page}/{limit} [get]
func (h *Handler) ListOrders(c *gin.Context) {
	p := c.Query("page")

	page, err := strconv.Atoi(p)
	if err != nil {
		h.Logger.Info("page not fount")
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "page not fount",
		})
		return
	}

	l := c.Query("limit")

	limit, err := strconv.Atoi(l)
	if err != nil {
		h.Logger.Info("limit not fount")
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "limit not fount",
		})
		return
	}

	order := pb.ListOrdersRequest{
		Page:  int32(page),
		Limit: int32(limit),
	}

	resp, err := h.ClientOrder.ListOrders(c, &order)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{
			"Error":   err.Error(),
			"Message": "Error while lisOrder invitation",
		})
		log.Println("Error while lisOrer invitation ", err)
		return
	}

	c.JSON(http.StatusOK, resp)
}

// @Summary Get an order
// @Description Get the details of an order
// @Tags orders
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param id path string true "Order ID"
// @Success 200 {object} orders.OrderInfo
// @Failure 400 {object} string
// @Failure 500 {object} string
// @Router /orders/{id} [get]
func (h *Handler) GetOrder(c *gin.Context) {
	id := c.Param("id")

	_, err := uuid.Parse(id)
	if err != nil {
		h.Logger.Info("Error id invalid", "err: ", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "StatusBadRequest",
			"message": fmt.Sprintf("Error with getting Id from URL: %s", err.Error()),
		})
		return
	}

	order := pb.Id{
		OrderId: id,
	}

	resp, err := h.ClientOrder.GetOrder(c, &order)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{
			"Error":   err.Error(),
			"Message": "Error while getOrder invitation",
		})
		log.Println("Error while getOrder invitation ", err)
		return
	}

	c.JSON(http.StatusOK, resp)
}


// @Summary Pay for an order
// @Description Make a payment for an order
// @Tags orders
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param order_id path string true "Order ID"
// @Param order body orders.PayOrderRequest true "Payment Info"
// @Success 200 {object} orders.PaymentInfo
// @Failure 400 {object} string
// @Failure 500 {object} string
// @Router /orders/{order_id} [post]
func (h *Handler) PayOrder(c *gin.Context) {
	id := c.Param("order_id")

	_, err := uuid.Parse(id)
	if err != nil {
		h.Logger.Info("Error id invalid", "err: ", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "StatusBadRequest",
			"message": fmt.Sprintf("Error with getting Id from URL: %s", err.Error()),
		})
		return
	}
	order := pb.PayOrderRequest{}

	err = json.NewDecoder(c.Request.Body).Decode(&order)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "StatusBadRequest",
			"message": fmt.Sprintf("Error NewDecoder or Decoder: %s", err.Error()),
		})
		return
	}

	order.OrderId = id

	if len(order.Cvv) != 3 {
		h.Logger.Info("card cvv invalid")
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "StatusBadRequest",
			"message": fmt.Sprintf("Error Cvv: %s", err),
		})
		return
	}

	if len(order.CardNumber) != 16 {
		h.Logger.Info("card number does not 16")
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "StatusBadRequest",
			"message": fmt.Sprintf("Error Card number invalid: %s", err),
		})
		return
	}

	resp, err := h.ClientOrder.PayOrder(c, &order)
	if err !=  nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{
			"Error":   err.Error(),
			"Message": "Error while payOrder invitation",
		})
		log.Println("Error while payOrder invitation ", err)
		return
	}

	c.JSON(http.StatusOK, resp)
}

// @Summary Get payment status
// @Description Get the payment status of an order
// @Tags orders
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param order_id path string true "Order ID"
// @Success 200 {object} orders.PaymentInfo
// @Failure 400 {object} string
// @Failure 500 {object} string
// @Router /orders/status/{order_id} [get]
func (h *Handler) GetPaymentStatus(c *gin.Context) {
	id := c.Param("order_id")

	_, err := uuid.Parse(id)
	if err != nil {
		h.Logger.Info("Error id invalid", "err: ", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "StatusBadRequest",
			"message": fmt.Sprintf("Error with getting Id from URL: %s", err.Error()),
		})
		return
	}

	order := pb.GetPaymentStatusRequest{
		OrderId: id,
	}

	resp, err := h.ClientOrder.GetPaymentStatus(c, &order)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{
			"Error":   err.Error(),
			"Message": "Error while GetPaymentStatusRequest invitation",
		})
		log.Println("Error while GetPaymentStatusRequest invitation ", err)
		return
	}

	c.JSON(http.StatusOK, resp)
}

