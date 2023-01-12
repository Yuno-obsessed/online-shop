package interaction

import (
	"github.com/gin-gonic/gin"
	"html/template"
	"log"
	"net/http"
	"os"
	"zusammen/internal/domain/entity"
)

type ProductInteraction interface {
	Form(c *gin.Context)
	FormErrors(c *gin.Context, errors map[string]string)
	Success(c *gin.Context)
	NotFound(c *gin.Context)
	Page(c *gin.Context, product *entity.Product)
	Delete(c *gin.Context)
}

type UserInteraction interface {
	Form(c *gin.Context, action string)
	FormErrors(c *gin.Context, action string, errors map[string]string)
	Success(c *gin.Context, action string)
	NotFound(c *gin.Context)
	Page(c *gin.Context, user *entity.User)
}

func RenderTemplate(c *gin.Context, filename string, data interface{}) {
	tmpl, err := template.ParseFiles(filename,
		//"../../assets/templates/base_layout.tmpl",
		//"../../assets/templates/base_footer_bot.tmpl",
		os.Getenv("TMPL_PATH")+"templates/base_layout.tmpl",
		os.Getenv("TMPL_PATH")+"templates/base_footer_bot.tmpl",
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
	//RenderTemplate(c, "../../assets/templates/home_page.html", nil)
	RenderTemplate(c, os.Getenv("TMPL_PATH")+"templates/home_page.html", nil)
}
