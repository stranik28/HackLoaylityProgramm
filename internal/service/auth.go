package service

import (
	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/stranik28/HackLoaylityProgramm/internal/helper"
	"github.com/stranik28/HackLoaylityProgramm/internal/logger"
	"github.com/stranik28/HackLoaylityProgramm/internal/models"
	"github.com/stranik28/HackLoaylityProgramm/internal/storage"
	"go.uber.org/zap"
	"time"
)

type ResponseToken struct {
	Token string `json:"token"`
}

func AuthHandler(c *gin.Context) {
	var requestBody models.User

	if err := c.BindJSON(&requestBody); err != nil {
		logger.Log.Fatal("NOT VALID JSON")
		panic(err)
	}

	logger.Log.Info("GET DATA", zap.Any("body", requestBody))

	if exists, _ := storage.UserId(requestBody.Username); !exists {
		id, err := storage.CreateUser(&requestBody)
		if err != nil {
			logger.Log.Fatal("Cannot create user", zap.Error(err))
			panic(err)
		}
		claims := &helper.SignedDetails{
			Id: id,
			StandardClaims: jwt.StandardClaims{
				ExpiresAt: time.Now().Local().Add(time.Hour * time.Duration(24)).Unix(),
			},
		}
		token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte(helper.SecretKey))
		if err != nil {
			logger.Log.Fatal("JWT ERROR", zap.Error(err))
			panic(err)
		}
		respToken := ResponseToken{Token: token}
		c.JSON(200, respToken)
		return
	}

	id, exc := c.Get("id")
	if !exc {
		logger.Log.Info("NO ID")
		panic("NO ID")
	}

	claims := &helper.SignedDetails{
		Id: id.(int),
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Local().Add(time.Hour * time.Duration(24)).Unix(),
		},
	}

	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte(helper.SecretKey))
	if err != nil {
		logger.Log.Fatal("JWT ERROR", zap.Error(err))
		panic(err)
	}
	respToken := ResponseToken{Token: token}
	c.JSON(200, respToken)
}

func CheckJWT(c *gin.Context) {
	id, _ := c.Get("id")
	logger.Log.Info("Here")
	c.String(200, string(rune(id.(int))))
}
