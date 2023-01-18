package auth

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/kataras/jwt"
	"os"
	"time"
)

type JWT struct {
	Secret []byte
}

func NewJWT() *JWT {
	return &JWT{Secret: []byte(os.Getenv("SECRET_JWT"))}
}

func (j *JWT) GenerateToken(username string) (string, error) {
	claims := jwt.Claims{
		Subject: username,
		Expiry:  time.Now().Add(72 * time.Hour).Unix(),
	}
	token, err := jwt.Sign(jwt.HS256, j.Secret, claims)
	if err != nil {
		return "", err
	}
	return string(token), nil
}
func (j *JWT) Authenticate() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString, err := c.Request.Cookie("token")
		if err != nil {
			c.AbortWithStatusJSON(500, fmt.Sprintf("Error getting jwt from cookie, %v", err))
		}
		fmt.Println(tokenString)
		verified, err := jwt.Verify(jwt.HS256, j.Secret, []byte(tokenString.String()))
		if err != nil {
			c.AbortWithStatusJSON(500, fmt.Sprintf("invalid token, %v", err))
			return
		}
		if time.Now().Unix() > verified.StandardClaims.Expiry {
			c.AbortWithStatusJSON(500, fmt.Sprintf("Token expired"))
			return
		}
	}
}

//func (j *JWT) GenerateToken(username string) (string, error) {
//	token := jwt.New(jwt.SigningMethodHS256)
//	claims := token.Claims.(jwt.MapClaims)
//	claims["sub"] = username
//	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()
//	fmt.Println(j.Secret)
//	tokenString, err := token.SignedString(j.Secret)
//	fmt.Println(tokenString)
//	if err != nil {
//		return "", err
//	}
//	return tokenString, nil
//}

//func (j *JWT) ValidateToken(tokenString string) (*Claims, error) {
//	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
//		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
//			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
//		}
//		return []byte(j.Secret), nil
//	})
//	if err != nil {
//		return nil, err
//	}
//	if claims, ok := token.Claims.(*Claims); ok && token.Valid {
//		return claims, nil
//	}
//	return nil, errors.New("Invalid token")
//}

//if err != nil {
//	c.JSON(500, fmt.Sprintf("missing token,%v", err))
//	return
//}
//fmt.Println(tokenString)
//if tokenString == "" {
//	c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Missing or invalid token"})
//	return
//}
//token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
//	if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
//		return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
//	}
//	return []byte(j.Secret), nil
//})
//if err != nil {
//	c.AbortWithStatusJSON(http.StatusUnauthorized, fmt.Sprintf("Unexpected signing method: %v", err))
//}
//token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
//	if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
//		return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
//	}
//	return token, nil
//})
//if err != nil {
//	c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
//	return
//}
//claims, ok := token.Claims.(jwt.MapClaims)
//if ok && token.Valid {
//	expireTime := int64(claims["exp"].(float64))
//	fmt.Print("exp:")
//	fmt.Println(expireTime)
//	if time.Now().Unix() > expireTime {
//		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Token expired"})
//		return
//	}
//} else {
//	c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
//	return
//}
