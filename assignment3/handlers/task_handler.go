package handlers

import (
	"assignment3/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func (h *Handler) CreateTask(c *gin.Context) {
	var task models.Task
	if err := c.ShouldBindJSON(&task); err != nil {
		h.logger.Error("Error binding json", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.service.CreateTask(&task); err != nil {
		h.logger.Error("Error creating task", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	h.logger.Info("Task created", task)
	c.JSON(http.StatusCreated, task)
}

func (h *Handler) GetTask(c *gin.Context) {
	id := c.Param("id")
	strc, err := strconv.Atoi(id)
	if err != nil {
		h.logger.Error("Error converting id to int", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	task, err := h.service.GetTask(strc)
	if err != nil {
		h.logger.Error("Error getting task", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	h.logger.Info("Task retrieved", task)
	c.JSON(http.StatusOK, task)
}

func (h *Handler) ClearCache(c *gin.Context) {
	err := h.service.ClearCache()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Cache cleared"})
}
