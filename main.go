package main

import (
	"health-api/database"
	"health-api/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.SetTrustedProxies(nil)
	database.ConnectDB()
	routes.PatientRoutes(r)

	r.Run(":8080") // http://localhost:8080
}
