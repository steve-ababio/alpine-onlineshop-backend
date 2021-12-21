package getcartcontroller

import (
	"config/auth_token"
	"config/token"
	"fmt"
	"model/cart"
	"net/http"
	"github.com/gin-gonic/gin"
)

type cartiteminfo struct {
	ProductID string `json:"productID" bson:"productid"`
	Title     string `json:"Title" bson:"title"`
	Price     string `json:"price" bson:"price"`
}

var cartitem struct {
	Item []cartiteminfo `json:"Cart" bson:"cart"`
}

func Getcartdata(context *gin.Context) {

	authtokench := make(chan string)
	payloadch := make(chan interface{})

	go auth_token.Getauthheaders(context, authtokench)
	authtoken := <-authtokench
	
	go token.Getcustompayloadclaims(authtoken, payloadch)
	userid_token_payload := <-payloadch
	cartresult := cart.Findusercart(userid_token_payload.(string))
	cartresult.Decode(&cartitem)
	fmt.Println(cartitem.Item)
	context.JSON(http.StatusOK, cartitem.Item)

}
