package auth_token

import (
	"fmt"
	"strings"

	"github.com/gin-gonic/gin"
)


func Getauthheaders(c *gin.Context,ch chan string){
	authtoken_str := c.Request.Header.Get("authorization")
	authtoken := strings.Split(authtoken_str, " ")
	fmt.Println(authtoken_str)
	ch <- authtoken[1];
}