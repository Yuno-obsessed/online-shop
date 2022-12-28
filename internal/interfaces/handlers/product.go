package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"strconv"
	"time"
	"zusammen/application"
	"zusammen/internal/domain/entity"
	"zusammen/internal/infrastructure/interaction"
)

type Product struct {
	productApp application.ProductAppInterface
	productInt interaction.ProductInteraction
	//token       auth.TokenInterface
}

func NewProduct(pApp application.ProductAppInterface, pInt interaction.ProductInteraction) *Product {
	return &Product{
		productApp: pApp,
		productInt: pInt,
		//token:      token,
	}
}

func (pr *Product) PostProduct(c *gin.Context) {
	//tokendata, err := pr.token.ExtractTokenMetadata(c.Request)
	//if err != nil {
	//	c.JSON(401, "unauthorized")
	//	return
	//}
	postProductErrors := make(map[string]string)
	if c.Request.Method == "GET" {
		pr.productInt.Info(c)
		return
	}
	price, _ := strconv.ParseFloat(c.PostForm("price"), 32)
	quantity, _ := strconv.Atoi(c.PostForm("quantity"))
	emptyProduct := entity.Product{
		Name:        c.PostForm("name"),
		Description: c.PostForm("description"),
		Image:       "img",
		Seller:      "some guy",
		Price:       price,
		Quantity:    quantity,
		Likes:       0,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	postProductErrors = emptyProduct.Validate("")
	if len(postProductErrors) > 0 {
		c.JSON(422, postProductErrors)
		return
	}
	postProduct, postErr := pr.productApp.PostProduct(&emptyProduct)
	if postErr != nil {
		c.JSON(500, postErr)
		return
	}
	c.JSON(201, postProduct)
}

func (pr *Product) GetProduct(c *gin.Context) {
	pr.productInt.Info(c)
	urlUuid := c.Param("product_uuid")
	productUuid, err := uuid.Parse(urlUuid)
	if err != nil {
		c.JSON(500, "invalid request")
		return
	}
	product, err := pr.productApp.GetProduct(productUuid)
	if err != nil {
		c.JSON(500, err.Error())
		return
	}
	c.JSON(200, product)
}

func (pr *Product) GetProducts(c *gin.Context) {
	var limit, offset int64 = 5, 0
	products, err := pr.productApp.GetProducts(limit, offset)
	if err != nil {
		c.JSON(500, err.Error())
		return
	}
	c.JSON(200, products)
}

func (pr *Product) UpdateProduct(c *gin.Context) {
	getProductErrors := make(map[string]string)
	urlUuid := c.Param("product_uuid")
	productUuid, err := uuid.Parse(urlUuid)
	if err != nil {
		c.JSON(500, "invalid request")
		return
	}
	price, _ := strconv.ParseFloat(c.PostForm("new_price"), 32)
	quantity, _ := strconv.Atoi(c.PostForm("new_quantity"))
	emptyProduct := &entity.Product{
		UUID:        productUuid,
		Name:        c.PostForm("new_name"),
		Description: c.PostForm("new_description"),
		Image:       "Image",
		Seller:      "some guy",
		Price:       price,
		Quantity:    quantity,
		UpdatedAt:   time.Now(),
	}

	getProductErrors = emptyProduct.Validate("")
	if len(getProductErrors) > 0 {
		c.JSON(422, getProductErrors)
		return
	}
	postProduct, updateErr := pr.productApp.EditProduct(emptyProduct, productUuid)
	if updateErr != nil {
		c.JSON(500, updateErr)
		return
	}
	c.JSON(201, postProduct)
}

func (pr *Product) DeleteProduct(c *gin.Context) {
	urlUuid := c.Param("product_uuid")
	productUuid, err := uuid.Parse(urlUuid)
	if err != nil {
		c.JSON(500, "invalid request")
		return
	}
	_, deleteErr := pr.productApp.DeleteProduct(productUuid)
	if err != nil {
		c.JSON(500, deleteErr)
		return
	}
	c.JSON(201, "success")
}
