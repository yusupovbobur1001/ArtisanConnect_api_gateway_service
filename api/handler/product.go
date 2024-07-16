package handler

import (
	pb "api_service/genproto/product"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func (h *Handler) CreateProduct(c *gin.Context) {
	p := pb.CreateProductRequest{}

	err := json.NewDecoder(c.Request.Body).Decode(&p)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "StatusBadRequest",
			"message": fmt.Sprintf("Error NewDecoder or Decoder: %s", err.Error()),
		})
		return
	}

	resp, err := h.ClientProduct.CreateProduct(c, &p)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{
			"Error":   err.Error(),
			"Message": "Error while StatusInternalServerError invitation",
		})
		log.Println("Error while StatusInternalServerError invitation ", err)
		return
	}

	c.JSON(http.StatusCreated, resp)

}

func (h *Handler) UpdateProduct(c *gin.Context) {
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

	p := pb.UpdateProductRequest{}
	err = json.NewDecoder(c.Request.Body).Decode(&p)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "StatusBadRequest",
			"message": fmt.Sprintf("Error NewDecoder or Decoder: %s", err.Error()),
		})
		return
	}

	resp, err := h.ClientProduct.UpdateProduct(c, &p)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{
			"Error":   err.Error(),
			"Message": "Error while UpdateProduct invitation",
		})
		log.Println("Error while UpdateProduct invitation ", err)
		return
	}

	c.JSON(http.StatusOK, resp)

}

func (h *Handler) DeleteProduct(c *gin.Context) {
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

	p := pb.Id{ProductId: id}

	resp, err := h.ClientProduct.DeleteProduct(c, &p)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{
			"Error":   err.Error(),
			"Message": "Error while DeleteProduct invitation",
		})
		log.Println("Error while DeleteProduct invitation ", err)
		return
	}

	c.JSON(http.StatusOK, resp)
}

func (h *Handler) ListProducts(c *gin.Context) {
	p := c.Param("page")

	page, err := strconv.Atoi(p)
	if err != nil {
		c.JSON(http.StatusBadGateway, gin.H{
			"error": "gape invalid",
		})
		return
	}

	l := c.Param("limit")

	limit, err := strconv.Atoi(l)
	if err != nil {
		c.JSON(http.StatusBadGateway, gin.H{
			"error": "limit invalid",
		})
		return
	}

	pr := pb.ListProductsRequest{
		Page:  int32(page),
		Limit: int32(limit),
	}

	resp, err := h.ClientProduct.ListProducts(c, &pr)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{
			"Error":   err.Error(),
			"Message": "Error while ListProducts invitation",
		})
		log.Println("Error while ListProducts invitation ", err)
		return
	}

	c.JSON(http.StatusOK, resp)
}

func (h *Handler) GetProduct(c *gin.Context) {
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

	p1 := pb.Id{ProductId: id}

	resp, err := h.ClientProduct.GetProduct(c, &p1)

	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{
			"Error":   err.Error(),
			"Message": "Error while GetProduct invitation",
		})
		log.Println("Error while GetProduct invitation ", err)
		return
	}

	c.JSON(http.StatusOK, resp)
}

func (h *Handler) SearchProducts(c *gin.Context) {

	

}
