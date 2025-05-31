package controllers

import (
	"hacka/core/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (ctrl *EmployeeController) GetEmployeeByID(c *gin.Context) {
	id := c.Param("id")

	employee, err := ctrl.service.GetEmployeeByID(c.Request.Context(), id)
	if err != nil {
		if err.Error() == "sql: no rows in result set" {
			c.JSON(http.StatusNotFound, gin.H{
				"error": "Funcionário não encontrado",
			})
			return
		}

		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   "Erro ao buscar funcionário",
			"details": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": employee,
	})
}

type EmployeeController struct {
	service *services.EmployeeService
}

func NewEmployeeController(service *services.EmployeeService) *EmployeeController {
	return &EmployeeController{
		service: service,
	}
}

func SetupEmployeeRoutes(router *gin.RouterGroup, service *services.EmployeeService) {
	controller := NewEmployeeController(service)

	employees := router.Group("/employees")
	{
		employees.GET("", controller.GetAllEmployees)
		employees.GET("/ranking", controller.GetEmployeesByAverageRanking)
		employees.GET("/:id", controller.GetEmployeeByID)
		employees.POST("", controller.CreateEmployee)
		employees.PUT("/:id", controller.UpdateEmployee)
		employees.DELETE("/:id", controller.DeleteEmployee)
		employees.GET("/position/:position", controller.GetEmployeesByPosition)
	}
}

func (ctrl *EmployeeController) GetAllEmployees(c *gin.Context) {
	employees, err := ctrl.service.GetAllEmployees(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   "Erro ao buscar funcionários",
			"details": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data":  employees,
		"count": len(employees),
	})
}

func (ctrl *EmployeeController) CreateEmployee(c *gin.Context) {
	var request struct {
		Name         string  `json:"name" binding:"required"`
		Email        string  `json:"email" binding:"required,email"`
		Position     string  `json:"position" binding:"required"`
		Balance      float64 `json:"balance"`
		Average      float64 `json:"average"`
		Qrcode       string  `json:"qrcode"`
		PasswordHash string  `json:"password_hash"`
		Midia        string  `json:"midia"`
	}

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Dados inválidos",
			"details": err.Error(),
		})
		return
	}

	employee, err := ctrl.service.CreateEmployee(
		c.Request.Context(),
		request.Name,
		request.Email,
		request.Position,
		request.Balance,
		request.Average,
		request.Qrcode,
		request.PasswordHash,
		request.Midia,
	)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   "Erro ao criar funcionário",
			"details": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"data":    employee,
		"message": "Funcionário criado com sucesso",
	})
}

func (ctrl *EmployeeController) UpdateEmployee(c *gin.Context) {
	id := c.Param("id")
	c.JSON(http.StatusOK, gin.H{
		"message": "Update employee - TODO: implement after generating models",
		"id":      id,
	})
}

func (ctrl *EmployeeController) DeleteEmployee(c *gin.Context) {
	id := c.Param("id")
	c.JSON(http.StatusOK, gin.H{
		"message": "Delete employee - TODO: implement after generating models",
		"id":      id,
	})
}

func (ctrl *EmployeeController) GetEmployeesByPosition(c *gin.Context) {
	position := c.Param("position")

	employees, err := ctrl.service.GetEmployeesByPosition(c.Request.Context(), position)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   "Erro ao buscar funcionários por cargo",
			"details": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data":     employees,
		"count":    len(employees),
		"position": position,
	})
}

func (ctrl *EmployeeController) GetEmployeesByAverageRanking(c *gin.Context) {
	limitStr := c.DefaultQuery("limit", "0")
	limit := 0

	if limitStr != "0" {
		if parsedLimit, err := strconv.Atoi(limitStr); err == nil && parsedLimit > 0 {
			limit = parsedLimit
		}
	}

	employees, err := ctrl.service.GetEmployeesByAverageRanking(c.Request.Context(), limit)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   "Erro ao buscar ranking de funcionários",
			"details": err.Error(),
		})
		return
	}

	response := gin.H{
		"data":         employees,
		"count":        len(employees),
		"ranking_type": "average",
	}

	if limit > 0 {
		response["limit"] = limit
	}

	c.JSON(http.StatusOK, response)
}
