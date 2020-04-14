package helpers

import (
	"fmt"
	"golang_training-master-Assignments/GolangTrianingAssignments/TodoMiddleware/api/models"

	jwt "github.com/dgrijalva/jwt-go"

	"errors"
	"time"

	"github.com/gin-gonic/gin"
)

var GlobalJWTKey string

func init() {
	GlobalJWTKey = "TODO"
}

type jwtCustomClaim struct {
	UserID   string   `json:"userId"`
	Password string   `json:"password"`
	Groups   []string `json:"groups"`
	Token    string   `json:"token"`
	jwt.StandardClaims
}

func GenerateToken(loginID string, password string, groups []string, expirationTime time.Duration) (string, error) {
	claims := jwtCustomClaim{
		UserID:   loginID,
		Password: password,
		Groups:   groups,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(expirationTime).Unix(),
		},
	}
	// Create token with claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// Generate encoded token and send it as response.
	t, err := token.SignedString([]byte(GlobalJWTKey))
	if err != nil {
		fmt.Println("Error during token generation", err)
	}
	return t, nil
}

// GetLoginFromToken login object from JWT Token
func GetLoginFromToken(c *gin.Context) (models.Login, error) {

	login := models.Login{}
	decodedToken, err := DecodeToken(c.GetHeader("Authorization"), GlobalJWTKey)
	fmt.Println("GetLoginFromToken -- decodedToken", decodedToken)
	if err != nil {
		return login, errors.New("GetLoginFromToken - unable to decode token")
	}
	// login ID is the compulsary field, so haven't added check for nil
	if decodedToken["userId"] == nil || decodedToken["userId"] == "" {
		return login, errors.New("GetLoginFromToken - login id not found")
	}
	login.LoginID = decodedToken["userId"].(string)
	login.Password = decodedToken["password"].(string)
	return login, nil
}

func DecodeToken(tokenFromRequest, jwtKey string) (jwt.MapClaims, error) {

	// get data i.e.Claims from token
	token, err := jwt.Parse(tokenFromRequest, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(jwtKey), nil
	})

	if err != nil {
		fmt.Println("Error while parsing JWT Token: ", err)
		return nil, err
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, errors.New("Error while getting claims")
	}
	return claims, nil
}
