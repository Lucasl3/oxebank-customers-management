// package main

// import (
// 	"log"
// )

// func main() {
// 	store, err := NewPostgresStore()
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	if err := store.Init(); err != nil {
// 		log.Fatal(err)
// 	}

// 	server := NewAPIServer(":8080", store)
// 	server.Run()
// }

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