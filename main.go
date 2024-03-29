package main

import (
	"customers-management/initializers"
	"customers-management/routes"

	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main() {
	r := gin.Default()
	
	routes.SetupRoutes(r)

	r.Run()
}