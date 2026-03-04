package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"web3-recruitment-admin/internal/model"
	"web3-recruitment-admin/internal/service"
)

type JobViewHandler struct {
	service *service.JobViewService
}

func NewJobViewHandler(s *service.JobViewService) *JobViewHandler {
	return &JobViewHandler{service: s}
}

func (h *JobViewHandler) RegisterRoutes(r *gin.RouterGroup) {
	r.GET("/stats/views", h.getAllStats)
	r.GET("/stats/views/:jobId", h.getStatsByJob)
}

// TrackView - public endpoint for tracking job views
func (h *JobViewHandler) TrackView(c *gin.Context) {
	var req model.TrackViewRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ip := c.ClientIP()
	userAgent := c.Request.UserAgent()

	if err := h.service.TrackView(req.JobID, ip, userAgent); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to track view"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"success": true})
}

func (h *JobViewHandler) getAllStats(c *gin.Context) {
	stats, err := h.service.GetAllStats()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	
	c.JSON(http.StatusOK, stats)
}

func (h *JobViewHandler) getStatsByJob(c *gin.Context) {
	jobID, err := strconv.Atoi(c.Param("jobId"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid job ID"})
		return
	}
	
	stats, err := h.service.GetStatsByJobID(jobID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Job not found"})
		return
	}

	c.JSON(http.StatusOK, stats)
}
