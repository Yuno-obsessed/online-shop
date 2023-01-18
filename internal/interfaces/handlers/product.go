package handlers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"strconv"
	"time"
	"zusammen/application"
	"zusammen/internal/domain/entity"
	"zusammen/internal/infrastructure/auth"
	"zusammen/internal/infrastructure/interaction"
	"zusammen/internal/interfaces"
)

type Product struct {
	productApp application.ProductAppInterface
	productInt interaction.ProductInteraction
	fileUpload interfaces.FileUploadInterface
	jwtToken   auth.JWT
}

func NewProduct(pApp application.ProductAppInterface, pInt interaction.ProductInteraction,
	fu interfaces.FileUploadInterface, token auth.JWT) *Product {
	return &Product{
		productApp: pApp,
		productInt: pInt,
		fileUpload: fu,
		jwtToken:   token,
	}
}

func (pr *Product) PostProduct(c *gin.Context) {
	//tokendata, err := pr.jwtToken.ValidateToken(c.Request.Header.Get("Authorization"))
	//if err != nil {
	//	c.JSON(401, "unauthorized")
	//	return
	//}
	if c.Request.Method == "GET" {
		pr.productInt.Form(c)
		return
	}
	price, err := strconv.Atoi(c.PostForm("price"))
	if err != nil {
		fmt.Println(err)
	}
	quantity, err := strconv.Atoi(c.PostForm("quantity"))
	if err != nil {
		fmt.Println(err)
	}
	imgName, err := c.FormFile("image")
	if err != nil {
		c.JSON(500, fmt.Sprintf("invalid image, %v", err))
		return
	}
	username, ok := c.Get("username")
	if !ok {
		c.JSON(500, fmt.Sprintf("invalid token metadata, %v", err))
	}
	image, err := pr.fileUpload.UploadFile(imgName, username.(string))
	if err != nil {
		c.JSON(500, fmt.Sprintf("can't upload image, %v", err))
		return
	}
	//category, _ := strconv.Atoi(c.PostForm("category"))
	emptyProduct := &entity.Product{
		UUID:        uuid.New(),
		Name:        c.PostForm("name"),
		Description: c.PostForm("description"),
		Category:    c.PostForm("category"),
		Seller:      "seller",
		Image:       image,
		Price:       price,
		Quantity:    quantity,
		Likes:       0,
		CreatedAt:   time.Now().Format("2006-01-02 15:04:05"),
		UpdatedAt:   time.Now().Format("2006-01-02 15:04:05"),
	}

	postProductErrors := emptyProduct.Validate()
	if len(postProductErrors) > 0 {
		pr.productInt.FormErrors(c, postProductErrors)
		return
	}
	_, postErr := pr.productApp.PostProduct(emptyProduct)
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
	//tokendata, err := pr.jwtToken.ValidateToken(c.Request.Header.Get("Authorization"))
	//if err != nil {
	//	c.JSON(401, "unauthorized")
	//	return
	//}
	getProductErrors := make(map[string]string)
	urlUuid := c.Param("product_uuid")
	productUuid, err := uuid.Parse(urlUuid)
	if err != nil {
		c.JSON(500, "invalid request")
		return
	}
	price, _ := strconv.Atoi(c.PostForm("new_price"))
	quantity, _ := strconv.Atoi(c.PostForm("new_quantity"))
	imgName, err := c.FormFile("new_image")
	if err != nil {
		c.JSON(500, "invalid image")
		return
	}
	username, ok := c.Get("username")
	if !ok {
		c.JSON(500, "Error getting username")
		return
	}
	image, err := pr.fileUpload.UploadFile(imgName, username.(string))
	if err != nil {
		c.JSON(500, "invalid image")
		return
	}
	//category, _ := strconv.Atoi(c.PostForm("new_category"))
	emptyProduct := &entity.Product{
		UUID:        productUuid,
		Name:        c.PostForm("new_name"),
		Description: c.PostForm("new_description"),
		Category:    c.PostForm("new_category"),
		Image:       image,
		Seller:      "seller",
		Price:       price,
		Quantity:    quantity,
		UpdatedAt:   time.Now().Format("2006-01-02 15:04:05"),
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
