package interaction

import (
	"github.com/gin-gonic/gin"
	"html/template"
	"log"
	"net/http"
	"zusammen/internal/domain/entity"
)

type ProductInteraction interface {
	Validate(product entity.Product) bool
	Send(c *gin.Context)
	Info(c *gin.Context)
}

type UserInteraction interface {
	Validate(user entity.User) bool
	Send(c *gin.Context)
	Info(c *gin.Context)
}

func RenderTemplate(c *gin.Context, filename string, data interface{}) {
	tmpl, err := template.ParseFiles(filename,
		"../../assets/templates/base_layout.tmpl",
		"../../assets/templates/base_footer_bot.tmpl",
	)
	if err != nil {
		log.Println(err.Error())
		http.Error(c.Writer, "Something went wrong", 500)
		return
	}
	err = tmpl.Execute(c.Writer, data)
	if err != nil {
		log.Println(err.Error())
		http.Error(c.Writer, "Something went wrong", 500)
	}
}

// Should add some logic (get categories list) here

func HomeTemplate(c *gin.Context) {
	RenderTemplate(c, "../../assets/templates/home_page.html", nil)
}
