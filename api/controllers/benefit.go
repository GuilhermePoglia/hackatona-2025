package controllers

import (
	"net/http"
	"strconv"

	"hacka/core/models"
	"hacka/core/services"

	"github.com/gin-gonic/gin"
	"github.com/volatiletech/null/v8"
)

type BenefitController struct {
	service *services.BenefitService
}

func NewBenefitController(service *services.BenefitService) *BenefitController {
	return &BenefitController{service: service}
}

func (c *BenefitController) GetAll(ctx *gin.Context) {
	benefits, err := c.service.GetAll(ctx.Request.Context())
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, benefits)
}

func (c *BenefitController) GetByID(ctx *gin.Context) {
	id := ctx.Param("id")
	benefit, err := c.service.GetByID(ctx.Request.Context(), id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, benefit)
}

func (c *BenefitController) Create(ctx *gin.Context) {
	var req struct {
		Name        string  `json:"name" binding:"required"`
		Description string  `json:"description"`
		Price       float64 `json:"price" binding:"required"`
		Image       string  `json:"image"`
	}

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	benefit := &models.Benefit{
		Name:        req.Name,
		Description: null.StringFromPtr(&req.Description),
		Price:       req.Price,
		Image:       null.StringFromPtr(&req.Image),
	}

	err := c.service.Create(ctx.Request.Context(), benefit)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, benefit)
}

func (c *BenefitController) Update(ctx *gin.Context) {
	id := ctx.Param("id")

	var req struct {
		Name        string  `json:"name" binding:"required"`
		Description string  `json:"description"`
		Price       float64 `json:"price" binding:"required"`
		Image       string  `json:"image"`
	}

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	benefit := &models.Benefit{
		Name:        req.Name,
		Description: null.StringFromPtr(&req.Description),
		Price:       req.Price,
		Image:       null.StringFromPtr(&req.Image),
	}

	err := c.service.Update(ctx.Request.Context(), id, benefit)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, benefit)
}

func (c *BenefitController) Delete(ctx *gin.Context) {
	id := ctx.Param("id")

	err := c.service.Delete(ctx.Request.Context(), id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Benefit deleted successfully"})
}

func (c *BenefitController) GetByPriceRange(ctx *gin.Context) {
	minPriceStr := ctx.Query("min_price")
	maxPriceStr := ctx.Query("max_price")

	if minPriceStr == "" || maxPriceStr == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "min_price and max_price query parameters are required"})
		return
	}

	minPrice, err := strconv.ParseFloat(minPriceStr, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid min_price value"})
		return
	}

	maxPrice, err := strconv.ParseFloat(maxPriceStr, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid max_price value"})
		return
	}

	benefits, err := c.service.GetByPriceRange(ctx.Request.Context(), minPrice, maxPrice)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, benefits)
}

func SetupBenefitRoutes(v1 *gin.RouterGroup, service *services.BenefitService) {
	controller := NewBenefitController(service)

	benefits := v1.Group("/benefits")
	{
		benefits.GET("", controller.GetAll)
		benefits.GET("/:id", controller.GetByID)
		benefits.POST("", controller.Create)
		benefits.PUT("/:id", controller.Update)
		benefits.DELETE("/:id", controller.Delete)
		benefits.GET("/price-range", controller.GetByPriceRange)
	}
}
