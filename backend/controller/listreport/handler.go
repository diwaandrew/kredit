package listreport

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	Service Service
}

func NewHandler(service Service) *Handler {
	return &Handler{service}
}

func (h *Handler) GetListReport(c *gin.Context) {
	statustrx := c.Query("statustrx")
	req := GetNasabah{StatusTrx: statustrx}
	listReport, status, err := h.Service.GetListReport(req)
	if err != nil {
		log.Println("Error handler Get : ", err)
		c.JSON(status, gin.H{
			"message": err.Error(),
		})
		return
	}
	c.JSON(status, gin.H{
		"message": "success",
		"data":    listReport,
	})
}

func (h *Handler) GetBranch(c *gin.Context) {
	Branch, status, err := h.Service.GetBranch()
	if err != nil {
		log.Println("Error handler Get : ", err)
		c.JSON(status, gin.H{
			"message": err.Error(),
		})
		return
	}
	c.JSON(status, gin.H{
		"message": "success",
		"data":    Branch,
	})
}

func (h *Handler) GetCompany(c *gin.Context) {
	Company, status, err := h.Service.GetCompany()
	if err != nil {
		log.Println("Error handler Get : ", err)
		c.JSON(status, gin.H{
			"message": err.Error(),
		})
		return
	}
	c.JSON(status, gin.H{
		"message": "success",
		"data":    Company,
	})
}

func (h *Handler) SearchListReport(c *gin.Context) {
	branch := c.Query("branch")
	company := c.Query("company")
	startdate := c.Query("startdate")
	enddate := c.Query("enddate")
	statustrx := c.Query("statustrx")
	req := GetSearchRequest{Branch: branch, Company: company, StartDate: startdate, EndDate: enddate, StatusTrx: statustrx}
	listReport, status, err := h.Service.SearchListReport(req)
	if err != nil {
		log.Println("Error handler Get : ", err)
		c.JSON(status, gin.H{
			"message": err.Error(),
		})
		return
	}
	c.JSON(status, gin.H{
		"message": "success",
		"data":    listReport,
	})
}

func (h *Handler) UpdateCustomer(c *gin.Context) {

	var req []requestbody
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Input data not suitable"})
		return
	}

	status, err := h.Service.UpdateCustomer(req)
	if err != nil {
		c.JSON(status, gin.H{
			"message": "Error bad request",
			"code":    "99",
		})
	}

	c.JSON(status, gin.H{
		"message": "success",
		"code":    "00",
	})
}
