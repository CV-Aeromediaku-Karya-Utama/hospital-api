package middleware

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"net/http"
	"strings"
)

type UnsignedResponse struct {
	Message interface{} `json:"message"`
}

type SignedResponse struct {
	Token   string `json:"token"`
	Message string `json:"message"`
}

func extractBearerToken(header string) (string, error) {
	if header == "" {
		return "", errors.New("bad header value given")
	}

	jwtToken := strings.Split(header, " ")
	if len(jwtToken) != 2 {
		return "", errors.New("incorrectly formatted authorization header")
	}

	return jwtToken[1], nil
}

func parseToken(jwtToken string) (*jwt.Token, error) {
	token, err := jwt.Parse(jwtToken, func(token *jwt.Token) (interface{}, error) {
		if _, OK := token.Method.(*jwt.SigningMethodHMAC); !OK {
			return nil, errors.New("bad signed method received")
		}
		return []byte("secret"), nil
	})

	if err != nil {
		return nil, errors.New("bad jwt token")
	}

	return token, nil
}

// JwtTokenCheck is to check jwt token
func JwtTokenCheck(c *gin.Context) {
	jwtToken, err := extractBearerToken(c.GetHeader("Authorization"))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, UnsignedResponse{
			Message: err.Error(),
		})
		return
	}

	token, err := parseToken(jwtToken)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, UnsignedResponse{
			Message: "bad jwt token",
		})
		return
	}

	_, OK := token.Claims.(jwt.MapClaims)

	if !OK {
		c.AbortWithStatusJSON(http.StatusInternalServerError, UnsignedResponse{
			Message: "unable to parse claims",
		})
		return
	}
	c.Next()
}

// PrivateACLCheck is to check jwt token and check params id equals to token.Claims id
func PrivateACLCheck(c *gin.Context) {
	jwtToken, err := extractBearerToken(c.GetHeader("Authorization"))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, UnsignedResponse{
			Message: err.Error(),
		})
		return
	}

	token, err := parseToken(jwtToken)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, UnsignedResponse{
			Message: "bad jwt token",
		})
		return
	}

	claims, OK := token.Claims.(jwt.MapClaims)
	fmt.Println(claims)
	if !OK {
		c.AbortWithStatusJSON(http.StatusInternalServerError, UnsignedResponse{
			Message: "unable to parse claims",
		})
		return
	}

	claimedUID, OK := claims["user_id"]
	if !OK {
		c.AbortWithStatusJSON(http.StatusBadRequest, UnsignedResponse{
			Message: "no user property in claims",
		})
		return
	}

	uid := c.Param("uid")
	if claimedUID != uid {
		c.AbortWithStatusJSON(http.StatusBadRequest, UnsignedResponse{
			Message: "token uid does not match resource uid",
		})
		return
	}

	c.Next()
}