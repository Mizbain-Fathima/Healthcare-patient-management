package routes

import (
	"health-api/controllers"

	"github.com/gin-gonic/gin"
)

func PatientRoutes(r *gin.Engine) {
	r.POST("/patients", controllers.CreatePatient)
	r.GET("/patients", controllers.GetPatients)
	r.GET("/patients/:id", controllers.GetPatient)
	r.PUT("/patients/:id", controllers.UpdatePatient)
	r.DELETE("/patients/:id", controllers.DeletePatient)
}
