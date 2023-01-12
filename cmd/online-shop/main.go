package main

import (
	"crypto/sha512"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"os"
	"zusammen/internal/infrastructure/auth"
	"zusammen/internal/infrastructure/config"
	"zusammen/internal/infrastructure/interaction"
	"zusammen/internal/infrastructure/persistence"
	"zusammen/internal/interfaces"
	"zusammen/internal/interfaces/handlers"
)

func main() {
	pepper := sha512.New().Sum([]byte("pepper"))
	os.Setenv("PEPPER", string(pepper))
	dbConf := config.NewConfig().MySqlConfig()
	router := gin.Default()
	services, err := persistence.NewRepositories(
		config.NewMysqlConn(dbConf))
	if err != nil {
		log.Fatal(err)
	}
	//frontend,err := interaction.
	defer services.Close()

	//userService := handlers.NewUsers()
	productService := handlers.NewProduct(services.Product, interaction.NewProductInt(),
		interfaces.NewFileUpload("product"), auth.JWT{Secret: os.Getenv("SECRET_JWT")})
	//router.Use(middleware.CORS())

	router.GET("/", interaction.HomeTemplate)
	//router.StaticFS("/assets/static/", gin.Dir("../../assets/static/", false))
	router.StaticFS("/assets/static/", gin.Dir(os.Getenv("TMPL_PATH")+"static/", false))
	router.GET("/create_product", productService.PostProduct)
	router.POST("/create_product", productService.PostProduct)
	router.GET("/products", productService.GetProducts)
	router.GET("/products/product:id", productService.GetProduct)

	log.Fatal(router.Run(":8080"))
}

// TODO
// Add a dependency-injection to handlers to check if user is authorised, using jwt
// Find out how to implement caching
