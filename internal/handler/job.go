package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"web3-recruitment-admin/internal/service"
)

type JobHandler struct {
	service *service.JobService
}

func NewJobHandler(s *service.JobService) *JobHandler {
	return &JobHandler{service: s}
}

func (h *JobHandler) RegisterRoutes(r *gin.RouterGroup) {
	r.GET("/jobs", h.getAll)
	r.GET("/jobs/:id", h.getByID)
}

func (h *JobHandler) getAll(c *gin.Context) {
	jobs, err := h.service.GetAllWithStats()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, jobs)
}

func (h *JobHandler) getByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid job ID"})
		return
	}

	job, err := h.service.GetByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Job not found"})
		return
	}

	c.JSON(http.StatusOK, job)
}
