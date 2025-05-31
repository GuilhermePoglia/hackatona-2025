package api

import (
	"database/sql"
	"hacka/api/controllers"
	"hacka/core/services"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func Setup(router *gin.Engine, db *sql.DB) {
	jwtKey := []byte(os.Getenv("JWT_SECRET"))
	if len(jwtKey) == 0 {
		jwtKey = []byte("my_secret_key")
	}

	employeeService := services.NewEmployeeService(db)
	feedbackService := services.NewFeedbackService(db, employeeService)
	resourceService := services.NewResourceService(db)
	activityService := services.NewActivityService(db)

	v1 := router.Group("/api/v1")

	controllers.SetupEmployeeRoutes(v1, employeeService)
	controllers.SetupFeedbackRoutes(v1, feedbackService)
	controllers.SetupResourceRoutes(v1, resourceService)
	controllers.SetupActivityRoutes(v1, activityService)
	router.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"message": "pong"})
	})
}
