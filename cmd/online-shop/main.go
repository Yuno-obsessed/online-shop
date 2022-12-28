package main

import (
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"zusammen/internal/infrastructure/config"
	"zusammen/internal/infrastructure/interaction"
	"zusammen/internal/infrastructure/persistence"
	"zusammen/internal/interfaces/handlers"
)

func main() {
	dbConf := config.NewConfig().MySqlConfig()
	router := gin.Default()
	services, err := persistence.NewRepositories(dbConf)
	if err != nil {
		log.Fatal(err)
	}
	//frontend,err := interaction.
	defer services.Close()

	//userService := handlers.NewUsers()
	productService := handlers.NewProduct(services.Product, interaction.NewProductInt())
	//router.Use(middleware.CORS())

	router.GET("/", interaction.HomeTemplate)
	router.StaticFS("/assets/static/", gin.Dir("../../assets/static/", false))
	router.GET("/create_product", productService.PostProduct)
	router.POST("/create_product", productService.PostProduct)
	router.GET("/products", productService.GetProducts)
	router.GET("/products/product:id", productService.GetProduct)

	log.Fatal(router.Run(":8080"))
}
