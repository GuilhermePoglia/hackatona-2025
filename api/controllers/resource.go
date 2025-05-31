package controllers

import (
	"hacka/core/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ResourceController struct {
	service *services.ResourceService
}

func NewResourceController(service *services.ResourceService) *ResourceController {
	return &ResourceController{
		service: service,
	}
}

func SetupResourceRoutes(router *gin.RouterGroup, service *services.ResourceService) {
	controller := NewResourceController(service)

	resources := router.Group("/resources")
	{
		resources.GET("", controller.GetAllResources)
		resources.GET("/:id", controller.GetResourceByID)
		resources.POST("", controller.CreateResource)
		resources.PUT("/:id", controller.UpdateResource)
		resources.DELETE("/:id", controller.DeleteResource)
		resources.GET("/type/:type", controller.GetResourcesByType)
	}
}

func (ctrl *ResourceController) GetAllResources(c *gin.Context) {
	resources, err := ctrl.service.GetAllResources(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   "Erro interno do servidor",
			"details": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"count": len(resources),
		"data":  resources,
	})
}

func (ctrl *ResourceController) GetResourceByID(c *gin.Context) {
	id := c.Param("id")

	resource, err := ctrl.service.GetResourceByID(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error":   "Recurso não encontrado",
			"details": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": resource,
	})
}

func (ctrl *ResourceController) CreateResource(c *gin.Context) {
	var request struct {
		Name  string `json:"name" binding:"required"`
		Type  string `json:"type" binding:"required"`
		Midia string `json:"midia"`
	}

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Dados inválidos",
			"details": err.Error(),
		})
		return
	}

	resource, err := ctrl.service.CreateResource(c.Request.Context(), request.Name, request.Type, request.Midia)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   "Erro ao criar recurso",
			"details": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Recurso criado com sucesso",
		"data":    resource,
	})
}

func (ctrl *ResourceController) UpdateResource(c *gin.Context) {
	id := c.Param("id")

	var request struct {
		Name  string `json:"name"`
		Type  string `json:"type"`
		Midia string `json:"midia"`
	}

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Dados inválidos",
			"details": err.Error(),
		})
		return
	}

	resource, err := ctrl.service.UpdateResource(c.Request.Context(), id, request.Name, request.Type, request.Midia)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   "Erro ao atualizar recurso",
			"details": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Recurso atualizado com sucesso",
		"data":    resource,
	})
}

func (ctrl *ResourceController) DeleteResource(c *gin.Context) {
	id := c.Param("id")

	err := ctrl.service.DeleteResource(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   "Erro ao deletar recurso",
			"details": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Recurso deletado com sucesso",
	})
}

func (ctrl *ResourceController) GetResourcesByType(c *gin.Context) {
	resourceType := c.Param("type")

	resources, err := ctrl.service.GetResourcesByType(c.Request.Context(), resourceType)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   "Erro interno do servidor",
			"details": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"count": len(resources),
		"data":  resources,
		"type":  resourceType,
	})
}
