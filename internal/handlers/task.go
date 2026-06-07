package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/yourorg/tasks-api/internal/models"
	"github.com/yourorg/tasks-api/internal/service"
)

type TaskHandler struct{ Svc *service.TaskService }

func (h *TaskHandler) List(c *gin.Context) {
	ownerID := c.GetUint("user_id")
	tasks, err := h.Svc.ListByOwner(c.Request.Context(), ownerID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, tasks)
}

func (h *TaskHandler) Create(c *gin.Context) {
	var dto models.CreateTaskDTO
	if err := c.ShouldBindJSON(&dto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	task, err := h.Svc.Create(c.Request.Context(), c.GetUint("user_id"), dto)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, task)
}

func (h *TaskHandler) Get(c *gin.Context) {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 64)
	task, err := h.Svc.Get(c.Request.Context(), uint(id), c.GetUint("user_id"))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "task not found"})
		return
	}
	c.JSON(http.StatusOK, task)
}
