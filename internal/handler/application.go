package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"web3-recruitment-admin/internal/service"
)

type ApplicationHandler struct {
	service *service.ApplicationService
}

func NewApplicationHandler(s *service.ApplicationService) *ApplicationHandler {
	return &ApplicationHandler{service: s}
}

func (h *ApplicationHandler) RegisterRoutes(r *gin.RouterGroup) {
	r.GET("/applications", h.getAll)
	r.GET("/applications/:jobId", h.getByJob)
}

func (h *ApplicationHandler) getAll(c *gin.Context) {
	apps, err := h.service.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, apps)
}

func (h *ApplicationHandler) getByJob(c *gin.Context) {
	jobID, err := strconv.Atoi(c.Param("jobId"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid job ID"})
		return
	}

	apps, err := h.service.GetByJobID(jobID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, apps)
}
