package controllers

import (
	"hacka/core/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ActivityController struct {
	service *services.ActivityService
}

func NewActivityController(service *services.ActivityService) *ActivityController {
	return &ActivityController{
		service: service,
	}
}

func SetupActivityRoutes(router *gin.RouterGroup, service *services.ActivityService) {
	controller := NewActivityController(service)

	activities := router.Group("/activities")
	{
		activities.GET("", controller.GetAllActivities)
		activities.GET("/:id", controller.GetActivityByID)
		activities.POST("", controller.CreateActivity)
		activities.PUT("/:id", controller.UpdateActivity)
		activities.DELETE("/:id", controller.DeleteActivity)
		activities.GET("/type/:type", controller.GetActivitiesByType)
	}
}

func (ctrl *ActivityController) GetAllActivities(c *gin.Context) {
	activities, err := ctrl.service.GetAllActivities(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   "Erro interno do servidor",
			"details": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"count": len(activities),
		"data":  activities,
	})
}

func (ctrl *ActivityController) GetActivityByID(c *gin.Context) {
	id := c.Param("id")

	activity, err := ctrl.service.GetActivityByID(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error":   "Atividade não encontrada",
			"details": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": activity,
	})
}

func (ctrl *ActivityController) CreateActivity(c *gin.Context) {
	var request struct {
		Name        string `json:"name" binding:"required"`
		Description string `json:"description"`
		Type        string `json:"type" binding:"required"`
	}

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Dados inválidos",
			"details": err.Error(),
		})
		return
	}

	activity, err := ctrl.service.CreateActivity(c.Request.Context(), request.Name, request.Description, request.Type)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   "Erro ao criar atividade",
			"details": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Atividade criada com sucesso",
		"data":    activity,
	})
}

func (ctrl *ActivityController) UpdateActivity(c *gin.Context) {
	id := c.Param("id")

	var request struct {
		Name        string `json:"name"`
		Description string `json:"description"`
		Type        string `json:"type"`
	}

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Dados inválidos",
			"details": err.Error(),
		})
		return
	}

	activity, err := ctrl.service.UpdateActivity(c.Request.Context(), id, request.Name, request.Description, request.Type)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   "Erro ao atualizar atividade",
			"details": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Atividade atualizada com sucesso",
		"data":    activity,
	})
}

func (ctrl *ActivityController) DeleteActivity(c *gin.Context) {
	id := c.Param("id")

	err := ctrl.service.DeleteActivity(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   "Erro ao deletar atividade",
			"details": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Atividade deletada com sucesso",
	})
}

func (ctrl *ActivityController) GetActivitiesByType(c *gin.Context) {
	activityType := c.Param("type")

	activities, err := ctrl.service.GetActivitiesByType(c.Request.Context(), activityType)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   "Erro interno do servidor",
			"details": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"count": len(activities),
		"data":  activities,
		"type":  activityType,
	})
}
