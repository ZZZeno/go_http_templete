package middlewares

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/octopart/go-itsdangerous"
	"log"
	"net/http"
)

const (
	// this should be the same as flask app.secret_key 
	secretKey = "123"
)

type UserInCAS struct {
	Username string `json:"CAS_USERNAME"`
	Ref      string `json:"ref,omitempty"`
	Service  string `json:"service,omitempty"`
	Temp     string `json:"temp,omitempty"`
}

// get user info from session
func UserFromCAS(session string) UserInCAS {
	timeStampSinger := itsdangerous.NewTimestampSignature(secretKey,
		"cookie-session",
		"",
		"hmac",
		nil,
		nil,
	)
	u, err := timeStampSinger.UnsignB64([]byte(session), 0)
	userInfo := UserInCAS{}
	if err != nil {
		log.Println("从session中获取CAS_USERNAME失败")
	}
	err = json.Unmarshal(u, &userInfo)
	if err != nil {
		log.Println("SSO中的session信息序列化成json失败")
	}
	//fmt.Printf(userInfo.Username)
	return userInfo
}

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		username := c.GetString("CAS_USERNAME")
		if username == "" {
			c.AbortWithStatus(http.StatusForbidden)
		}
		c.Next()
	}
}

// get infomation from session, and pass them to handlers later on. other handlers can get CAS_USERNAME by c.GetString("CAS_USERNAME")
func FillSession() gin.HandlerFunc {
	return func(c *gin.Context) {
		cookie, _ := c.Cookie("session")
		userInfo := UserFromCAS(cookie)
		c.Set("CAS_USERNAME", userInfo.Username)
		c.Next()
	}
}
