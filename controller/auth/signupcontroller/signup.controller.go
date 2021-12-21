package signupcontroller

import (
	"fmt"
	"model/auth/signupmodel"
	"net/http"
	"response"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
	"go.mongodb.org/mongo-driver/mongo"
	"github.com/gin-gonic/gin"
)
type(
	clientdata struct{
		id primitive.ObjectID
		Email string `json:"email"`;
		Firstname string `json:"firstname"`;
		Surname string `json:"surname"`;
		Password string `json:"password"`;
		Phone string `json:"phone"`;
	}
)

const cost = 10;
//Signup registers a new user
func Signup(context  *gin.Context){

	var userdata clientdata;
	var err error;
	var insertresult *mongo.InsertOneResult
	var hashedpassword []byte;
	// err = json.NewDecoder(req.Body).Decode(&userdata);
	context.BindJSON(&userdata);

	if err != nil{
		interserverresponse := response.Internalservererrorresponse();
		context.JSON(http.StatusInternalServerError,interserverresponse);
		return;
	}
	//verify if username and email dont already exist
	//hash password
	hashedpassword,err = bcrypt.GenerateFromPassword([]byte(userdata.Password),cost);
	if err != nil{
		interserverresponse := response.Internalservererrorresponse();
		context.JSON(http.StatusInternalServerError,interserverresponse);
		return;
	}
	//store hashed password;
	userdata.Password = string(hashedpassword);
	insertresult,err = signupmodel.Saveoneuser(userdata);
	signup_ok := response.Getauthresponse("Thanks,account created successfully",true);
	context.JSON(http.StatusOK,signup_ok);
	fmt.Println(insertresult);

}