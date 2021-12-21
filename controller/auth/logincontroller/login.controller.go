package logincontroller

import (
	// "cookieconfig"
	"config/token"
	"fmt"
	"model/auth/loginmodel"
	"model/cart"
	"net/http"
	"response"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	// "golang.org/x/crypto/bcrypt"
)

type (
	user struct {
		Email string `json:"email"`
		Password string `json:"password"`
	}
	loginsuccessresponse struct {
		Messageresponse interface{} `json:"message"`
		Token  string `json:"token"`
		Cartexist bool `json:"cartexist"`
	}
	tokeninfo struct{
		ID primitive.ObjectID;
		Email string;
	}
	dbuserinfo struct{
	 	ID	primitive.ObjectID `bson:"_id"`;
		Email string `bson:"email"`;
		Password string `bson:"password"`; 
	}
)
func authenticate(dbuser *dbuserinfo,clientuser *user,context *gin.Context) bool{
	failedauthresp := response.Getauthresponse("Login attempt failed, invalid login credentials", false)
	db_user_data := loginmodel.Findoneuser(clientuser.Email);
	err := db_user_data.Decode(&dbuser)
	if err != nil {
		fmt.Println(failedauthresp);
		context.JSON(http.StatusBadRequest, failedauthresp)
		return false;
	}
	err = bcrypt.CompareHashAndPassword([]byte(req_user.Password), []byte(dbuser.password))
	if err != nil {
		context.JSON(http.StatusBadRequest, failedauthresp)
		return
	}	
	return true;
} 
//Login handler
func Login(context *gin.Context) {
	var req_user user;
	var dbuser dbuserinfo;
	okauthresp := response.Getauthresponse("Login successful", true)
	err := context.BindJSON(&req_user)
	if err != nil {
		interservererrorresponse := response.Internalservererrorresponse()
		context.JSON(http.StatusInternalServerError, interservererrorresponse)
		return;
	}
	if authenticate(&dbuser,&req_user,context){
		token, err := token.CreateToken(&tokeninfo{ID:dbuser.ID,Email: dbuser.Email});
		if err != nil {
			context.JSON(http.StatusUnprocessableEntity,response.Internalservererrorresponse())
		}
		fmt.Println(token);
		r := cart.Findusercart(dbuser.ID.Hex())

		if(r.Err() == mongo.ErrNoDocuments){
			okres := loginsuccessresponse{okauthresp, token,false};
			context.JSON(http.StatusOK, okres);
		}else{
			okres := loginsuccessresponse{okauthresp, token,true};
			context.JSON(http.StatusOK, okres);
		}
	}

}
