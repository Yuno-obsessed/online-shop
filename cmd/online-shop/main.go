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
	defer services.Close()

	jwtService := auth.NewJWT()
	userService := handlers.NewUser(services.User, interaction.NewUserInt(),
		interfaces.NewFileUpload("user"), *jwtService)
	productService := handlers.NewProduct(services.Product, interaction.NewProductInt(),
		interfaces.NewFileUpload("product"), *jwtService)
	//router.Use(middleware.CORS())

	router.GET("/", interaction.HomeTemplate)
	//router.StaticFS("/assets/static/", gin.Dir("../../assets/static/", false))
	router.StaticFS("/assets/static/", gin.Dir(os.Getenv("TMPL_PATH")+"static/", false))
	//prods := router.Group("account/:/products/")
	prods := router.Group("/products/")
	prods.Use(jwtService.Authenticate())
	prods.GET("create", productService.PostProduct)
	prods.POST("create", productService.PostProduct)
	prods.GET("products", productService.GetProducts)
	prods.GET(":id", productService.GetProduct)

	router.GET("/account/register", userService.SaveUser)
	router.POST("/account/register", userService.SaveUser)
	router.GET("/account/login", userService.GetUserByEmailAndPassword)
	log.Fatal(router.Run(":8080"))
}
