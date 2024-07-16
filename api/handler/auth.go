package handler

import (
	pb "api_service/genproto/auth"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func (h *Handler) UpdateProfile(c *gin.Context) {
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

	u := pb.UserUpdate{}

	err = json.NewDecoder(c.Request.Body).Decode(&u)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "StatusBadRequest",
			"message": fmt.Sprintf("Error NewDecoder or Decoder: %s", err.Error()),
		})
		return
	}

	u.Id = id

	resp, err := h.ClientAuthentication.UpdateProfile(c, &u)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{
			"Error":   err.Error(),
			"Message": "Error while UpdateProfile invitation",
		})
		log.Println("Error while UpdateProfile invitation ", err)
		return
	}

	c.JSON(http.StatusOK, resp)
}

func (h *Handler) DeleteProfile(c *gin.Context) {
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

	u := pb.Id{Id: id}

	resp, err := h.ClientAuthentication.DeleteProfile(c, &u)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{
			"Error":   err.Error(),
			"Message": "Error while DeleteProfile invitation",
		})
		log.Println("Error while DeleteProfile invitation ", err)
		return
	}

	c.JSON(http.StatusOK, resp)
}

func (h *Handler) GetByIdProfile(c *gin.Context) {
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

	u := pb.Id{Id: id}

	resp, err := h.ClientAuthentication.GetByIdProfile(c, &u)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{
			"Error":   err.Error(),
			"Message": "Error while GetByIdProfile invitation",
		})
		log.Println("Error while GetByIdProfile invitation ", err)
		return
	}

	c.JSON(http.StatusOK, resp)
}

func (h *Handler) GetAllProfil(c *gin.Context) {
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

	u := pb.Filter{
		Page: int32(page),
		Limit: int32(limit),
	}

	resp, err := h.ClientAuthentication.GetAllProfil(c, &u)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{
			"Error":   err.Error(),
			"Message": "Error while GetAllProfil invitation",
		})
		log.Println("Error while GetAllProfil invitation ", err)
		return
	}

	c.JSON(http.StatusOK, resp)
}


