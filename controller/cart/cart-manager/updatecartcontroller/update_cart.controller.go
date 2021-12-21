package updatecartcontroller

import (
	"config/auth_token"
	"config/token"
	"fmt"
	"model/cart"
	"net/http"
	"response"
	"github.com/gin-gonic/gin"
)

type(
	//cart item json structure
	cartitems struct {
		ProductID string `json:"productID"`;
		Title string `json:"Title"`;
		Price string `json:"price"`;
	}
	usercart struct{
		User_id string `json:"user_id" bson:"user_id"`;
		Cart []cartitems  `json:"cartitems" bson:"Cart"`;
	}
)
var reqcartitem struct{
	Item cartitems `json:"cartitem"`;
}
func Addcartitem(c *gin.Context) {
	
	//channels for token and payload claims
	tokench := make(chan string);  
	payloadch := make(chan interface{});

	//retrieves user token from authorization bearer scheme header 
	go auth_token.Getauthheaders(c,tokench);
	logintoken := <-tokench;

	fmt.Println("authorization: ",logintoken);

	//parses cartitem json's body  
	err := c.BindJSON(&reqcartitem)
	fmt.Println(reqcartitem);

	if err != nil {
		c.JSON(http.StatusInternalServerError,response.Internalservererrorresponse());
		return
	}
	//get payload private claims from user login token
	go token.Getcustompayloadclaims(logintoken,payloadch);
	userid_token_payload := <-payloadch;	

	//checks whether cart item exists   
	find_cartitem_results := cart.Findusercartitem(userid_token_payload.(string),reqcartitem.Item.ProductID);
	if find_cartitem_results.Err() == nil{
		return;
	}

	//updates cart
	matched_count,upd_err := cart.Updateusercart(userid_token_payload.(string), reqcartitem.Item);

	if upd_err != nil{
		c.JSON(http.StatusInternalServerError,response.Internalservererrorresponse());
		return;
	}

	//creates new cart if carts doesn't exist during update
	if(matched_count.(int64) < 1){
		cart_ := &usercart{User_id:userid_token_payload.(string)};
		cart_.Cart = []cartitems{reqcartitem.Item};
		cart.Insertusercartitems(cart_);
	}

}
