package interaction

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"regexp"
	"zusammen/internal/domain/entity"
)

type UserInt struct {
	Errors map[string]string
}

func NewUserInt() *UserInt {
	return &UserInt{
		Errors: map[string]string{},
	}
}

var _ UserInteraction = &UserInt{}

func (ui *UserInt) Validate(user entity.User) bool {
	ui.Errors = make(map[string]string)
	reg := regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")
	match := reg.Match([]byte(user.Email))
	if !match {
		ui.Errors["Email"] = "Please provide a valid email address"
	}
	if user.Password == "" {
		ui.Errors["Password"] = "Please enter a password"
	}
	return len(ui.Errors) == 0
}

func (ui *UserInt) Send(c *gin.Context) {
	resp := entity.User{
		Email:    c.PostForm("email"),
		Password: c.PostForm("password"),
	}

	if !ui.Validate(resp) {
		log.Println(ui.Errors)
		RenderTemplate(c, "../../assets/templates/account_login", resp)
		return
	}
	http.Redirect(c.Writer, c.Request, "/account/page", 303)
}

func (ui *UserInt) Info(c *gin.Context) {
	RenderTemplate(c, "../../assets/templates/account_login", nil)
}

//func (ui *UserInteraction) GetInfo() any {
//	return &entity.User{}
//}
