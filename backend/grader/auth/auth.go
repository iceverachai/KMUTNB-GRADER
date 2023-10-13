package auth

import (
	"encoding/json"
	"errors"
	"fmt"
	"grader/database"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

var SecretKey = []byte("SecretKey")

type AuthenticateReq struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type Userdata struct {
	Username  string
	Password  string
	Firstname string
	Lastname  string
	Fullname  string
	Groups    []string
	IsAdmin   bool
	Work      gin.H
}

func Authenticate(ctx *gin.Context) {
	authenticateData := AuthenticateReq{}

	ctx.BindJSON(&authenticateData)

	if len(authenticateData.Username) <= 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":   "DataError",
			"response": "Username is blank or undefined",
		})
		return
	}
	if len(authenticateData.Password) <= 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":   "DataError",
			"response": "Password is blank or undefined",
		})
		return
	}

	result := database.Connection.QueryRow(
		"SELECT * FROM account WHERE Username = ? AND Password = ?",
		authenticateData.Username,
		authenticateData.Password,
	)

	userd := Userdata{}
	var unparseddata string
	var unparseddata2 string
	err := result.Scan(&userd.Username, &userd.Password, &userd.Firstname, &userd.Lastname, &unparseddata, &userd.IsAdmin, &unparseddata2)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"status":   "InternalError",
			"response": "Error while processing user data",
		})
		fmt.Println(err.Error())
		return
	}

	err = json.Unmarshal([]byte(unparseddata), &userd.Groups)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"status":   "InternalError",
			"response": "Error while parsing data",
		})
		fmt.Println(err.Error())
		return
	}

	err = json.Unmarshal([]byte(unparseddata2), &userd.Work)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"status":   "InternalError",
			"response": "Error while parsing data",
		})
		fmt.Println(err.Error())
		return
	}

	userd.Fullname = userd.Firstname + " " + userd.Lastname

	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["username"] = userd.Username
	claims["isadmin"] = userd.IsAdmin
	claims["firstname"] = userd.Firstname
	claims["lastname"] = userd.Lastname
	claims["fullname"] = userd.Fullname
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix()

	tokenString, err := token.SignedString(SecretKey)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"status":   "InternalError",
			"response": "Error generating token",
		})
		fmt.Println(err.Error())
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status":   "success",
		"response": tokenString,
	})

}

func AccessGate(ctx *gin.Context) {
	authHeader := ctx.Request.Header.Get("Authorization")

	tokenString, err := ProcessHeaderAuthorizeField(authHeader)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":   "BadData",
			"response": err.Error(),
		})
		return
	}

	token, err := IsAuthorized(tokenString)

	if err != nil || !token.Valid {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":   "InvalidData",
			"response": "Invalid token",
		})
		return
	}

	claims, ok := token.Claims.(jwt.MapClaims)

	if ok {
		response := gin.H{}
		for key, value := range claims {
			response[key] = value
		}

		ctx.JSON(http.StatusOK, gin.H{
			"status":   "success",
			"response": response,
		})
	} else {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":   "InvalidData",
			"response": "Invalid token to claim",
		})
	}

}

func ProcessHeaderAuthorizeField(header string) (string, error) {
	if header == "" {
		return "", errors.New("empty header field")
	}

	if len(header) > 7 && strings.ToUpper(header[0:6]) == "BEARER" {
		return header[7:], nil
	}

	return "", errors.New("bad token")
}

func IsAuthorized(tkn string) (*jwt.Token, error) {
	token, err := jwt.Parse(tkn, func(token *jwt.Token) (any, error) {
		return SecretKey, nil
	})

	if err != nil {
		return nil, err
	}

	return token, nil
}
