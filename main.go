package main

import (
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"

	"restaurant_management/database"
	"restaurant_management/middleware"
	"restaurant_management/routes"
)

var foodCollection *mongo.Collection = database.OpenCollection(database.Client, "food")

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		fmt.Println("no port attached, port will be set as 8080 immediately...")
		port = "8080"
	}

	router := gin.New()
	router.Use(gin.Logger())

	routes.UserRoutes(router)

	router.Use(middleware.Authentication())

	routes.FoodRoutes(router)
	routes.MenuRoutes(router)
	routes.TableRoutes(router)
	routes.OrderRoutes(router)
	routes.OrderItemRoutes(router)
	routes.InvoiceRoutes(router)

	router.Run(":" + port)
}
