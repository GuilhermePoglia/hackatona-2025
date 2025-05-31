package controllers

import (
	"hacka/core/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type FeedbackController struct {
	service *services.FeedbackService
}

func NewFeedbackController(service *services.FeedbackService) *FeedbackController {
	return &FeedbackController{
		service: service,
	}
}

func SetupFeedbackRoutes(router *gin.RouterGroup, service *services.FeedbackService) {
	controller := NewFeedbackController(service)

	feedbacks := router.Group("/feedbacks")
	{
		feedbacks.GET("", controller.GetAllFeedbacks)
		feedbacks.GET("/:id", controller.GetFeedbackByID)
		feedbacks.POST("", controller.CreateFeedback)
		feedbacks.GET("/receiver/:id", controller.GetFeedbacksByReceiver)
		feedbacks.GET("/sender/:id", controller.GetFeedbacksBySender)
		feedbacks.GET("/stats/:id", controller.GetEmployeeStats)
	}
}

func (ctrl *FeedbackController) CreateFeedback(c *gin.Context) {
	var request struct {
		SenderID    string `json:"sender_id" binding:"required"`
		ReceiverID  string `json:"receiver_id" binding:"required"`
		Stars       int    `json:"stars" binding:"required,min=1,max=5"`
		Description string `json:"description"`
	}

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Dados inválidos",
			"details": err.Error(),
		})
		return
	}

	if request.SenderID == request.ReceiverID {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Não é possível enviar feedback para si mesmo",
		})
		return
	}

	feedback, err := ctrl.service.CreateFeedback(
		c.Request.Context(),
		request.SenderID,
		request.ReceiverID,
		request.Stars,
		request.Description,
	)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   "Erro ao criar feedback",
			"details": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"data":    feedback,
		"message": "Feedback criado com sucesso",
	})
}

func (ctrl *FeedbackController) GetAllFeedbacks(c *gin.Context) {
	feedbacks, err := ctrl.service.GetAllFeedbacks(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   "Erro ao buscar feedbacks",
			"details": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data":  feedbacks,
		"count": len(feedbacks),
	})
}

func (ctrl *FeedbackController) GetFeedbackByID(c *gin.Context) {
	id := c.Param("id")

	feedback, err := ctrl.service.GetFeedbackByID(c.Request.Context(), id)
	if err != nil {
		if err.Error() == "sql: no rows in result set" {
			c.JSON(http.StatusNotFound, gin.H{
				"error": "Feedback não encontrado",
			})
			return
		}

		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   "Erro ao buscar feedback",
			"details": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": feedback,
	})
}

func (ctrl *FeedbackController) GetFeedbacksByReceiver(c *gin.Context) {
	receiverID := c.Param("id")

	feedbacks, err := ctrl.service.GetFeedbacksByReceiver(c.Request.Context(), receiverID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   "Erro ao buscar feedbacks recebidos",
			"details": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data":        feedbacks,
		"count":       len(feedbacks),
		"receiver_id": receiverID,
	})
}

func (ctrl *FeedbackController) GetFeedbacksBySender(c *gin.Context) {
	senderID := c.Param("id")

	feedbacks, err := ctrl.service.GetFeedbacksBySender(c.Request.Context(), senderID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   "Erro ao buscar feedbacks enviados",
			"details": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data":      feedbacks,
		"count":     len(feedbacks),
		"sender_id": senderID,
	})
}

func (ctrl *FeedbackController) GetEmployeeStats(c *gin.Context) {
	employeeID := c.Param("id")

	stats, err := ctrl.service.GetEmployeeStats(c.Request.Context(), employeeID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   "Erro ao buscar estatísticas",
			"details": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data":        stats,
		"employee_id": employeeID,
	})
}
