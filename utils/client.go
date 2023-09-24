package utils

import (
	"crypto"
	"errors"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/gin-contrib/requestid"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/golang-jwt/jwt/v5/request"
	logger "github.com/sirupsen/logrus"
)

type AccessClaims struct {
	UID       int64
	TokenType string
	jwt.RegisteredClaims
}

type RefreshClaims struct {
	UID       int64
	TokenType string
	Salt      int64
	jwt.RegisteredClaims
}

type ClientInfo struct {
	Ip        string
	Platform  string
	Os        string
	UserAgent string
	SessionId string
	UserId    int64
	RequestId string
}

func (c *ClientInfo) LogFormatShort() string {
	return fmt.Sprintf("%s/%d/%s", c.RequestId, c.UserId, c.SessionId)
}

func (c *ClientInfo) LogFormatShortAndIp() string {
	return fmt.Sprintf("%s/%d/%s/%s", c.RequestId, c.UserId, c.SessionId, c.Ip)
}

func (c *ClientInfo) LogFormatLong() string {
	return fmt.Sprintf("%s/%d/%s/%s/%s/%s/%s", c.RequestId, c.UserId, c.SessionId, c.Ip, c.Platform, c.Os, c.UserAgent)
}

func TrailClient(domain string, maxAge int) gin.HandlerFunc {
	return func(c *gin.Context) {
		client := ClientInfo{}
		client.Ip = c.ClientIP()
		client.Os = c.GetHeader("os")
		client.Platform = c.GetHeader("platform")
		client.UserAgent = c.GetHeader("User-Agent")
		if sessionId, err := c.Cookie(SESSION_COOKIE_NAME); err == nil {
			client.SessionId = sessionId
		} else {
			sessionId = GenerateID()
			client.SessionId = sessionId
			c.SetSameSite(http.SameSiteNoneMode)
			c.SetCookie(SESSION_COOKIE_NAME, sessionId, maxAge, "/",
				domain, true, false)
		}
		c.Set(CLIENT_KEY, client)
		c.Next()
	}
}

func GetClientInfo(c *gin.Context) (*ClientInfo, error) {
	if cinfo, exists := c.Get(CLIENT_KEY); exists {
		if client, ok := cinfo.(ClientInfo); ok {
			client.RequestId = requestid.Get(c)
			client.UserId = c.GetInt64(UID_KEY)
			logger.Debugf("get client: %s", client.LogFormatLong())
			return &client, nil
		} else {
			logger.Error("get error client")
			return nil, errors.New("invalid client")
		}
	} else {
		logger.Error("get empty client")
		return nil, errors.New("empty client")
	}
}

func GetUser(verifyKey crypto.PublicKey) gin.HandlerFunc {
	return func(c *gin.Context) {
		if id, err := verifyToken(c, verifyKey); err == nil {
			c.Set(UID_KEY, id)
			AddExtraInfoToContext(c, LOGINED, true)
		} else {
			c.Set(UID_KEY, 0)
			AddExtraInfoToContext(c, LOGINED, false)
		}
		c.Next()
	}
}

func UserAuthen(verifyKey crypto.PublicKey) gin.HandlerFunc {
	return func(c *gin.Context) {
		if id, err := verifyToken(c, verifyKey); err == nil {
			c.Set(UID_KEY, id)
			AddExtraInfoToContext(c, LOGINED, true)
			c.Next()
		} else {
			AddExtraInfoToContext(c, LOGINED, false)
			ResponseFailedJson(c, ERRCODE_NOT_LOGINED, ERRMSG_NOT_LOGINED, nil, http.StatusForbidden)
		}
	}
}

func verifyToken(c *gin.Context, verifyKey crypto.PublicKey) (int64, error) {
	token, err := request.ParseFromRequest(c.Request, request.OAuth2Extractor, func(token *jwt.Token) (interface{}, error) {
		// since we only use the one private key to sign the tokens,
		// we also only use its public counter part to verify
		return verifyKey, nil
	}, request.WithClaims(&AccessClaims{}))

	// If the token is missing or invalid, return error
	if err != nil {
		if errors.Is(err, jwt.ErrTokenExpired) {
			if token.Claims.(*AccessClaims).IssuedAt.Unix() < time.Now().Add(-TOKEN_EXPIRE*2).Unix() {
				c.Header("Auth-Refresh", "refresh")
				return token.Claims.(*AccessClaims).UID, nil
			}
		}
		logger.Errorf("Invalid token: %s", err)
		return 0, err
	}
	return token.Claims.(*AccessClaims).UID, nil
}

func VerifyRefleshToken(t string, verifyKey *crypto.PublicKey) (*jwt.Token, error) {
	return jwt.ParseWithClaims(t, &RefreshClaims{}, func(token *jwt.Token) (interface{}, error) {
		return verifyKey, nil
	})
}

func LoadEdPrivateKeyFromDisk(location string) crypto.PrivateKey {
	keyData, e := os.ReadFile(location)
	if e != nil {
		panic(e.Error())
	}
	key, e := jwt.ParseEdPrivateKeyFromPEM(keyData)
	if e != nil {
		panic(e.Error())
	}
	return key
}

func LoadEdPublicKeyFromDisk(location string) crypto.PublicKey {
	keyData, e := os.ReadFile(location)
	if e != nil {
		panic(e.Error())
	}
	key, e := jwt.ParseEdPublicKeyFromPEM(keyData)
	if e != nil {
		panic(e.Error())
	}
	return key
}
