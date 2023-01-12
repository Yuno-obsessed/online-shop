package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"time"
	"zusammen/application"
	"zusammen/internal/domain/entity"
	"zusammen/internal/infrastructure/interaction"
	"zusammen/internal/infrastructure/security"
)

type User struct {
	userApp application.UserAppInterface
	userInt interaction.UserInteraction
}

func NewUser(uapp application.UserAppInterface, uint interaction.UserInteraction) *User {
	return &User{
		userInt: uint,
		userApp: uapp,
	}
}

// This route is dedicated to take info from frontend, validate it and call
// web-pages rendering

// For now is a GodObject, so gotta move taking info from frontend to interaction pkg
func (us *User) SaveUser(c *gin.Context) {
	if c.Request.Method == "GET" {
		us.userInt.Form(c, "register")
		return
	}
	var nickname string
	if c.PostForm("nickname") == "" {
		nickname = c.PostForm("firstname") + "_" + c.PostForm("lastname")
	}
	// Need to find some string to generate salt from
	newUser := entity.User{
		UUID:      uuid.New(),
		FirstName: c.PostForm("firstname"),
		LastName:  c.PostForm("lastname"),
		Nickname:  nickname,
		Age:       c.PostForm("age"),
		Email:     c.PostForm("email"),
		Phone:     c.PostForm("phone"),
		Password:  c.PostForm("password"),
		Salt:      security.GenerateSalt(nickname),
		Purchases: 0,
		CreatedAt: time.Now().Format("2006-01-02 15:04:05"),
		UpdatedAt: time.Now().Format("2006-01-02 15:04:05"),
	}
	secondPassword := c.PostForm("repeatpassword")
	saveUserErrors := newUser.Validate("register", secondPassword)
	if len(saveUserErrors) > 0 {
		us.userInt.FormErrors(c, "register", saveUserErrors)
		return
	}
	_, saveErr := us.userApp.SaveUser(&newUser)
	if saveErr != nil {
		c.JSON(500, saveErr)
		return
	}
	us.userInt.Success(c, "register")
}

func (us *User) GetUser(c *gin.Context) {
	// router.GET("/user/:uuid", GetUser)
	urlUuid := c.Param("uuid")
	userUuid, err := uuid.Parse(urlUuid)
	if err != nil {
		// To check if uuid.Parse works correctly
		c.JSON(500, nil)
		return
	}
	_, getUserErr := us.userApp.GetUser(userUuid)
	if getUserErr != nil {
		us.userInt.NotFound(c)
		return
	}
	us.userInt.Success(c, "register")
}

// Am not sure yet where/how to use it
func (us *User) GetUsers(c *gin.Context) {

}

func (us *User) GetUserByEmailAndPassword(c *gin.Context) {
	if c.Request.Method == "GET" {
		us.userInt.Form(c, "login")
		return
	}
	loginUser := entity.User{
		Email:    c.PostForm("email"),
		Password: c.PostForm("password"),
	}
	_, logErrors := us.userApp.GetUserByEmailAndPassword(&loginUser)
	if len(logErrors) > 0 {
		us.userInt.FormErrors(c, "login", logErrors)
		return
	}

	us.userInt.Success(c, "login")
}
