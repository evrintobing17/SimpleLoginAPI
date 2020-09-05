package main

import (
	"log"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

var (
	router = gin.Default()
)

var jwtKey = []byte("chrombit2020")

func main() {
	router.POST("/login", Login)
	log.Fatal(router.Run(":1111"))
}

type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type Claims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

var user = User{
	Username: "admin",
	Password: "root",
}

func Login(c *gin.Context) {
	var u User
	if err := c.ShouldBindJSON(&u); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": "can't bind struct",
		})
		return
	}
	//compare the user from the request, with the one we defined:
	if user.Username != u.Username || user.Password != u.Password {
		c.JSON(http.StatusUnauthorized, gin.H{
			"status":  http.StatusUnauthorized,
			"message": "wrong username or password",
		})
		return
	}

	claims := &Claims{
		Username: u.Username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Minute * 15).Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "login success",
		"token":   tokenString,
	})
}

// func CreateToken(username string) (uint64, error) {
// 	var err error

// 	atClaims := jwt.MapClaims{}
// 	atClaims["authorized"] = true
// 	atClaims["username"] = username
// 	atClaims["exp"] = time.Now().Add(time.Minute * 15).Unix()
// 	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
// 	token, err := at.SignedString([]byte(os.Getenv("chrombit2020")))
// 	if err != nil {
// 		return "", err
// 	}
// 	return token, nil
// }