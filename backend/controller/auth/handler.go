package auth

import (
	"log"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	Service Service
}

func NewHandler(service Service) *Handler {
	return &Handler{service}
}

func (h *Handler) GetNasabah(c *gin.Context) {
	nasabah, status, err := h.Service.GetNasabah()
	if err != nil {
		log.Println("Error handler Get : ", err)
		c.JSON(status, gin.H{
			"message": err.Error(),
		})
		return
	}
	c.JSON(status, gin.H{
		"message": "success",
		"data":    nasabah,
	})
}
