package interaction

import (
	"github.com/gin-gonic/gin"
)

type ProductInt struct {
	Errors map[string]string
	TmplPath string
}

func NewProductInt() *ProductInt {
	return &ProductInt{
		Errors: map[string]string{},
		TmplPath: "../../assets/templates/product/",
	}
}

var _ ProductInteraction = &ProductInt{}

// I guess I can take gin.Context out to ProductInt struct as well as UserInt
func (pi *ProductInt) Form(c *gin.Context) {
RenderTemplate(c, pi.TmplPath + "create_product.html", nil)
}

func (pi *ProductInt) FormErrors(c *gin.Context, errors map[string]string) {
	pi.Errors = errors
	RenderTemplate(c, pi.TmplPath + "create_product.html", pi.Errors)
}

func (pi *ProductInt) Success(c *gin.Context) {
	RenderTemplate(c, pi.TmplPath + "success_creating_product.html", nil)
}

func (pi *ProductInt) NotFound(c *gin.Context){
	RenderTemplate(c, pi.TmplPath + "product_not_found",nil)
}

