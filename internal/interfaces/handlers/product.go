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
	if c.Request.Method == "GET" {
		pr.productInt.Form(c)
		return
	}
	price, _ := strconv.Atoi(c.PostForm("price"))
	quantity, _ := strconv.Atoi(c.PostForm("quantity"))
	//category, _ := strconv.Atoi(c.PostForm("category"))
	emptyProduct := entity.Product{
		Name:        c.PostForm("name"),
		Description: c.PostForm("description"),
		Image:       "img",
		Category:    c.PostForm("category"),
		Seller:      "seller",
		Price:       price,
		Quantity:    quantity,
		Likes:       0,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	postProductErrors := emptyProduct.Validate()
	if len(postProductErrors) > 0 {
		pr.productInt.FormErrors(c, postProductErrors)
		return
	}
	_, postErr := pr.productApp.PostProduct(&emptyProduct)
	if postErr != nil {
		c.JSON(500, postErr)
		return
	}
	pr.productInt.Success(c)
}

func (pr *Product) GetProduct(c *gin.Context) {
	pr.productInt.Form(c)
	urlUuid := c.Param("product_uuid")
	productUuid, err := uuid.Parse(urlUuid)
	if err != nil {
		c.JSON(500, "invalid request")
		return
	}
	getProduct, err := pr.productApp.GetProduct(productUuid)
	if err != nil {
		c.JSON(500, err.Error())
		return
	}
	pr.productInt.Page(c, getProduct)
}

func (pr *Product) GetProducts(c *gin.Context) {
	var limit, offset int64 = 5, 0
	products, err := pr.productApp.GetProducts(limit, offset)
	if err != nil {
		c.JSON(500, err.Error())
		return
	}
	// change to uploading some page
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
	price, _ := strconv.Atoi(c.PostForm("new_price"))
	quantity, _ := strconv.Atoi(c.PostForm("new_quantity"))
	//category, _ := strconv.Atoi(c.PostForm("new_category"))
	emptyProduct := &entity.Product{
		UUID:        productUuid,
		Name:        c.PostForm("new_name"),
		Description: c.PostForm("new_description"),
		Category:    c.PostForm("new_category"),
		Image:       "Image",
		Seller:      "seller",
		Price:       price,
		Quantity:    quantity,
		UpdatedAt:   time.Now(),
	}

	getProductErrors = emptyProduct.Validate()
	if len(getProductErrors) > 0 {
		c.JSON(422, getProductErrors)
		return
	}
	postProduct, updateErr := pr.productApp.EditProduct(emptyProduct, productUuid)
	if updateErr != nil {
		c.JSON(500, updateErr)
		return
	}
	pr.productInt.Page(c, postProduct)
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
	pr.productInt.Delete(c)
}
