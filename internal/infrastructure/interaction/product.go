package interaction

import (
	"github.com/gin-gonic/gin"
	"os"
	"zusammen/internal/domain/entity"
)

type ProductInt struct {
	Errors   map[string]string
	TmplPath string
}

func NewProductInt() *ProductInt {
	return &ProductInt{
		Errors: map[string]string{},
		//TmplPath: "../../assets/templates/product/",
		TmplPath: os.Getenv("TMPL_PATH") + "templates/product/",
	}
}

var _ ProductInteraction = &ProductInt{}

// I guess I can take gin.Context out to ProductInt struct as well as UserInt
func (pi *ProductInt) Form(c *gin.Context) {
	RenderTemplate(c, pi.TmplPath+"create_product.html", nil)
}

func (pi *ProductInt) FormErrors(c *gin.Context, errors map[string]string) {
	pi.Errors = errors
	RenderTemplate(c, pi.TmplPath+"create_product.html", pi.Errors)
}

func (pi *ProductInt) Success(c *gin.Context) {
	RenderTemplate(c, pi.TmplPath+"success_creating_product.html", nil)
}

func (pi *ProductInt) NotFound(c *gin.Context) {
	RenderTemplate(c, pi.TmplPath+"product_not_found.html", nil)
}

func (pi *ProductInt) Page(c *gin.Context, product *entity.Product) {
	RenderTemplate(c, pi.TmplPath+"product_page.html", product)
}

func (pi *ProductInt) Delete(c *gin.Context) {
	RenderTemplate(c, pi.TmplPath+"success_deleting_product.html", nil)
}
