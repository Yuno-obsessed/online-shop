package interaction

import (
	"github.com/gin-gonic/gin"
	"os"
	"zusammen/internal/domain/entity"
)

type UserInt struct {
	Errors   map[string]string
	TmplPath string
}

func NewUserInt() *UserInt {
	return &UserInt{
		Errors: map[string]string{},
		//TmplPath: "../../assets/templates/user/",
		TmplPath: os.Getenv("TMPL_PATH") + "templates/user/",
	}
}

var _ UserInteraction = &UserInt{}

// Prolly I can take away map from struct

func (ui *UserInt) Form(c *gin.Context, action string) {
	if action == "login" {
		RenderTemplate(c, ui.TmplPath+"account_login.html", nil)
	}
	if action == "register" {
		RenderTemplate(c, ui.TmplPath+"account_register.html", nil)
	}
}

func (ui *UserInt) FormErrors(c *gin.Context, action string, errors map[string]string) {
	ui.Errors = errors
	if action == "register" {
		RenderTemplate(c, ui.TmplPath+"account_register.html", ui.Errors)
		return
	}
	if action == "login" {
		RenderTemplate(c, ui.TmplPath+"account_login.html", ui.Errors)
		return
	}
}

func (ui *UserInt) Success(c *gin.Context, action string) {
	if action == "register" {
		RenderTemplate(c, ui.TmplPath+"success_register.html", nil)
		return
	}
	if action == "login" {
		RenderTemplate(c, ui.TmplPath+"success_login.html", nil)
		return
	}
}

func (ui *UserInt) NotFound(c *gin.Context) {
	RenderTemplate(c, ui.TmplPath+"user_not_found.html", nil)
}

func (ui *UserInt) Page(c *gin.Context, user *entity.User) {
	RenderTemplate(c, ui.TmplPath+"user_page.html", user)
}
