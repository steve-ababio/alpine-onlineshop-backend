package createcartcontroller

import (
	"config/auth_token"
	"config/token"
	"fmt"
	"model/cart"
	"net/http"
	"response"
	"github.com/gin-gonic/gin"
)

type (
	//cart items json structure
	cartitems struct {
		ProductID string `json:"productID"`
		Title     string `json:"Title"`
		Price     string `json:"price"`
	}
	// database schema
	usercartinfo struct {
		User_ID string      `json:"user_id" bson:"user_id"`
		Cart    []cartitems `json:"cartitems" bson:"Cart"`
	}
)

var cartitemsreq struct {
	Usercart []cartitems `json:"cartitems"`
}

func Addusercart(c *gin.Context) {
	tokench := make(chan string)        //token channel
	payloadch := make(chan interface{}) //token payload channel

	go auth_token.Getauthheaders(c, tokench) // get authentication token from authorization bearer header scheme
	logintoken := <-tokench

	fmt.Println("authorization cart: ", logintoken)
	err := c.BindJSON(&cartitemsreq) // parses json in cartitems datastructure

	if err != nil {
		c.JSON(http.StatusInternalServerError, response.Internalservererrorresponse())
		return
	}
	go token.Getcustompayloadclaims(logintoken, payloadch) //retrieves payload claims from user valid token
	userid_token_payload := <-payloadch

	dbuser_cart := usercartinfo{
		User_ID: userid_token_payload.(string),
		Cart:    cartitemsreq.Usercart,
	}
	_, err = cart.Insertusercartitems(dbuser_cart) //inserts user cart

	if err != nil {
		c.JSON(http.StatusInternalServerError, response.Internalservererrorresponse())
		return
	}
	c.JSON(http.StatusOK, struct {
		Insertstatus bool `json:"insertstatus"`
	}{Insertstatus: true}) //sends json response to user
}
