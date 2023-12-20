package handler

import (
	"clinics/config"
	"clinics/storage"
	"log"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	cfg  *config.Config
	strg storage.StorageI
}

type Response struct {
	Status      int         `json:"status"`
	Description string      `json:"description"`
	Data        interface{} `json:"data"`
}

func NewHandler(cfg *config.Config, strg storage.StorageI) *Handler {
	return &Handler{
		cfg:  cfg,
		strg: strg,
	}
}

func handleResponse(c *gin.Context, status int, data interface{}) {
	var description string
	switch code := status; {
	case code < 400:
		description = "success"
	default:
		description = "error"
		log.Println(config.Error, "error while:", Response{
			Status:      status,
			Description: description,
			Data:        data,
		})

		if code == 500 {
			data = "Internal Server Error"
		}
	}

	c.JSON(status, Response{
		Status:      status,
		Description: description,
		Data:        data,
	})
}
