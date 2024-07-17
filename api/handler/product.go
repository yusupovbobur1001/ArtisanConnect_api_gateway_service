package handler

import (
	pb "api_service/genproto/products"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)


// @Summary Create a product
// @Description Create a new product
// @Tags products
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param product body products.CreateProductRequest true "Product Create Info"
// @Success 201 {object} products.ProductInfo
// @Failure 400 {object} string
// @Failure 500 {object} string
// @Router /products [post]
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

// UpdateProduct godoc
// @Summary Update a product
// @Description Update product information by ID
// @Tags products
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param id path string true "Product ID"
// @Param product body products.UpdateProductRequest true "Product update data"
// @Success 200 {object} products.ProductInfo
// @Failure 400 {object} string
// @Failure 500 {object} string
// @Router /products/{id} [put]
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


// DeleteProduct godoc
// @Summary Delete a product
// @Description Delete a product by ID
// @Tags products
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param id path string true "Product ID"
// @Success 200 {object} products.DeleteProductResponse
// @Failure 400 {object} string
// @Failure 500 {object} string
// @Router /products/{id} [delete]
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


// ListProducts godoc
// @Summary List products
// @Description Get a paginated list of products
// @Tags products
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param page query int true "Page number"
// @Param limit query int true "Number of items per page"
// @Success 200 {object} products.ListProductsResponse
// @Failure 400 {object} string
// @Failure 502 {object} string
// @Failure 500 {object} string
// @Router /profiles/aa [get]
func (h *Handler) ListProducts(c *gin.Context) {
	p := c.Query("page")

	page, err := strconv.Atoi(p)
	if err != nil {
		c.JSON(http.StatusBadGateway, gin.H{
			"error": "gape invalid",
		})
		return
	}

	l := c.Query("limit")

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

// GetProduct godoc
// @Summary Get a product
// @Description Get a product by ID
// @Tags products
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param id path string true "Product ID"
// @Success 200 {object} products.ProductInfo
// @Failure 400 {object} string
// @Failure 500 {object} string
// @Router /profiles/{id} [get]
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


// SearchProducts godoc
// @Summary Search products
// @Description Search products by query and price range with pagination
// @Tags products
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param page query int true "Page number"
// @Param limit query int true "Number of items per page"
// @Param query query string false "Search query"
// @Param min_price query string false "Minimum price"
// @Param max_price query string false "Maximum price"
// @Success 200 {object} products.SearchProductsResponse
// @Failure 400 {object} string
// @Failure 502 {object} string
// @Failure 500 {object} string
// @Router /profiles/search [get]
func (h *Handler) SearchProducts(c *gin.Context) {
	p := c.Query("page")

	page, err := strconv.Atoi(p)
	if err != nil {
		c.JSON(http.StatusBadGateway, gin.H{
			"error": "gape invalid",
		})
		return
	}

	l := c.Query("limit")

	limit, err := strconv.Atoi(l)
	if err != nil {
		c.JSON(http.StatusBadGateway, gin.H{
			"error": "limit invalid",
		})
		return
	}

	q := c.Query("query")

	minp := c.Query("min_price")

	minPrice, err := strconv.ParseFloat(minp, 32)
	if err != nil {
		c.JSON(http.StatusBadGateway, gin.H{
			"error": "min_price invalid",
		})
		return
	}

	maxp := c.Query("max_price")

	maxPrice, err := strconv.ParseFloat(maxp, 32)
	if err != nil {
		c.JSON(http.StatusBadGateway, gin.H{
			"error": "max_price invalid",
		})
		return
	}

	if maxPrice < minPrice {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "min price katta max price dan",
		})
		return
	}

	p1 := pb.SearchProductsRequest{
		Page:     int32(page),
		Limit:    int32(limit),
		MinPrice: float32(maxPrice),
		MaxPrice: float32(minPrice),
		Query:    q,
	}

	resp, err := h.ClientProduct.SearchProducts(c, &p1)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{
			"Error":   err.Error(),
			"Message": "Error while SearchProducts invitation",
		})
		log.Println("Error while SearchProducts invitation ", err)
		return
	}

	c.JSON(http.StatusOK, resp)
}

// AddProductRating godoc
// @Summary Add product rating
// @Description Add a rating for a product by product ID
// @Tags products
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param product_id path string true "Product ID"
// @Param rating body products.AddProductRatingRequest true "Product rating data"
// @Success 200 {object} products.RatingInfo
// @Failure 400 {object}  string
// @Failure 500 {object}  string
// @Router /profiles/rating/{product_id} [post]
func (h *Handler) AddProductRating(c *gin.Context) {
	id := c.Param("product_id")

	_, err := uuid.Parse(id)
	if err != nil {
		h.Logger.Info("Error id invalid", "err: ", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "StatusBadRequest",
			"message": fmt.Sprintf("Error with getting Id from URL: %s", err.Error()),
		})
		return
	}

	p := pb.AddProductRatingRequest{}

	err = json.NewDecoder(c.Request.Body).Decode(&p)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "StatusBadRequest",
			"message": fmt.Sprintf("Error NewDecoder or Decoder: %s", err.Error()),
		})
		return
	}

	p.ProductId = id

	resp, err := h.ClientProduct.AddProductRating(c, &p)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{
			"Error":   err.Error(),
			"Message": "Error while AddProductRating invitation",
		})
		log.Println("Error while AddProductRating invitation ", err)
		return
	}

	c.JSON(http.StatusOK, resp)
}


// GetProductRatings godoc
// @Summary Get product ratings
// @Description Get ratings for a product by product ID
// @Tags products
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param product_id path string true "Product ID"
// @Success 200 {object} products.GetProductRatingsResponse
// @Failure 400 {object} string
// @Failure 500 {object} string
// @Router /profiles/ratings/{product_id} [get]
func (h *Handler) GetProductRatings(c *gin.Context) {
	id := c.Param("product_id")

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

	resp, err := h.ClientProduct.GetProductRatings(c, &p)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{
			"Error":   err.Error(),
			"Message": "Error while AddProductRating invitation",
		})
		log.Println("Error while AddProductRating invitation ", err)
		return
	}

	c.JSON(http.StatusOK, resp)
}
