package auth

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/google/uuid"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

//type JwtWrapper struct {
//	Secret          string
//	Issuer          string
//	ExpirationHours int64
//}

//type JwtClaim struct {
//	UserId string
//	jwt.StandardClaims
//}

type Token struct{}

func NewToken() *Token {
	return &Token{}
}

type TokenInterface interface {
	GenerateToken(userid uuid.UUID) (*TokenDetails, error)
	ExtractTokenMetadata(*http.Request) (*AccessDetails, error)
}

// Token implements the TokenInterface
var _ TokenInterface = &Token{}

func (t *Token) GenerateToken(userId uuid.UUID) (*TokenDetails, error) {

	td := &TokenDetails{}
	td.AtExpires = time.Now().Add(time.Minute * 60).Unix()
	td.TokenUuid = uuid.NewV4().String()

	td.RtExpires = time.Now().Add(time.Hour * 24 * 7).Unix()
	td.RefreshUuid = td.TokenUuid + "++" + id

	atClaims := jwt.MapClaims{}
	atClaims["authorized"] = true
	atClaims["access_uuid"] = td.TokenUuid
	atClaims["user_id"] = userId
	atClaims["exp"] = td.AtExpires
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	var err error
	td.AccessToken, err = at.SignedString([]byte(os.Getenv("ACCESS_SECRET")))
	if err != nil {
		return nil, err
	}

	rtClaims := jwt.MapClaims{}
	rtClaims["refresh_uuid"] = td.RefreshUuid
	rtClaims["user_id"] = userId
	rtClaims["exp"] = td.AtExpires
	rt := jwt.NewWithClaims(jwt.SigningMethodHS256, rtClaims)
	td.RefreshToken, err = rt.SignedString([]byte(os.Getenv("REFRESH_SECRET")))
	if err != nil {
		return nil, err
	}
	//claims := &JwtClaim{
	//	UserId: id,
	//	StandardClaims: jwt.StandardClaims{
	//		ExpiresAt: time.Now().Local().Add(time.Hour * time.Duration(j.ExpirationHours)).Unix(),
	//		Issuer:    j.Issuer,
	//	},
	//}

	//token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	//
	//signedToken, err := token.SignedString([]byte(j.Secret))
	//if err != nil {
	//	return "", err
	//}
	//
	//return signedToken, nil
	return td, nil
}

func TokenValid(r *http.Request) error {
	token, err := VerifyToken(r)
	if err != nil {
		return err
	}
	if _, ok := token.Claims.(jwt.Claims); !ok && !token.Valid {
		return err
	}
	return nil
}

func VerifyToken(r *http.Request) (*jwt.Token, error) {
	tokenString := ExtractToken(r)
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		//Make sure that the token method conform to "SigningMethodHMAC"
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("ACCESS_SECRET")), nil
	})
	if err != nil {
		return nil, err
	}
	return token, nil
}

func ExtractToken(r *http.Request) string {
	bearToken := r.Header.Get("Authorization")
	strArr := strings.Split(bearToken, " ")
	if len(strArr) == 2 {
		return strArr[1]
	}
	return ""
}

func (t *Token) ExtractTokenMetadata(r *http.Request) (*AccessDetails, error) {
	fmt.Println("WE ENTERED METADATA")
	token, err := VerifyToken(r)
	if err != nil {
		return nil, err
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if ok && token.Valid {
		accessUuid, ok := claims["access_uuid"].(string)
		if !ok {
			return nil, err
		}
		userId, err := strconv.ParseUint(fmt.Sprintf("%.f", claims["user_id"]), 10, 64)
		if err != nil {
			return nil, err
		}
		return &AccessDetails{
			TokenUuid: accessUuid,
			UserId:    userId,
		}, nil
	}
	return nil, err
}

//func (j *JwtWrapper) ValidateToken(signedToken string) (*JwtClaim, error) {
//	token, err := jwt.ParseWithClaims(
//		signedToken,
//		&JwtClaim{},
//		func(token *jwt.Token) (interface{}, error) {
//			return []byte(j.Secret), nil
//		},
//	)
//
//	if err != nil {
//		return nil, err
//	}
//
//	claims, ok := token.Claims.(*JwtClaim)
//	if !ok {
//		err = errors.New("Couldn't parse claims")
//		return nil, err
//	}
//
//	if claims.ExpiresAt < time.Now().Local().Unix() {
//		err = errors.New("JWT is expired")
//		return nil, err
//	}
//
//	return claims, nil
//}
