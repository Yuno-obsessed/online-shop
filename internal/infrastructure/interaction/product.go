package interaction

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"regexp"
	"strconv"
	"strings"
	"zusammen/internal/domain/entity"
)

type ProductInt struct {
	Errors map[string]string
}

func NewProductInt() *ProductInt {
	return &ProductInt{
		Errors: map[string]string{},
	}
}

var _ ProductInteraction = &ProductInt{}

func (pi *ProductInt) Validate(product entity.Product) bool {
	pi.Errors = make(map[string]string)
	reg := regexp.MustCompile(`[a-zA-Z]*.`)
	match := reg.Match([]byte(product.Name))
	if !match {
		pi.Errors["Name"] = "Please enter a valid product name"
	}
	if strings.TrimSpace(product.Description) == "" {
		pi.Errors["Description"] = "Please enter a description"
	}
	reg = regexp.MustCompile(`^\d*\.\d{2}$`)
	match = reg.Match([]byte(fmt.Sprintf("%v", product.Price)))
	if !match {
		pi.Errors["Price"] = "Please enter a valid price"
	}
	reg = regexp.MustCompile(`\d+`)
	match = reg.Match([]byte(fmt.Sprintf("%v", product.Quantity)))
	if !match {
		pi.Errors["Quantity"] = "Please enter a valid quantity"
	}
	log.Println(pi.Errors)
	return len(pi.Errors) == 0
}

func (pi *ProductInt) Send(c *gin.Context) {
	price, _ := strconv.ParseFloat(c.Request.PostFormValue("price"), 32)
	quantity, _ := strconv.Atoi(c.Request.PostFormValue("quantity"))
	resp := entity.Product{
		Name:        c.Request.PostFormValue("name"),
		Description: c.Request.PostFormValue("description"),
		Price:       price,
		Quantity:    quantity,
	}

	if !pi.Validate(resp) {
		log.Println(pi.Errors)
		RenderTemplate(c, "../../assets/templates/create_product.html", resp)
		return
	}
	http.Redirect(c.Writer, c.Request, "create_product/success", 303)
}

func (pi *ProductInt) Info(c *gin.Context) {
	RenderTemplate(c, "../../assets/templates/create_product.html", nil)
}

//func (pi *ProductInteraction) GetInfo() any {
//	return &entity.Product{
//		Name:        product.Name,
//		Description: product.Description,
//		Seller:      "some guy",
//		Price:       pi.Product.Price,
//		Quantity:    pi.Product.Quantity,
//		Likes:       0,
//		CreatedAt:   time.Now(),
//	}
//}
