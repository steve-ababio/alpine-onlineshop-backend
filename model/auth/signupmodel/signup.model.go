package signupmodel

import (
	"context"
	"databases"
	"go.mongodb.org/mongo-driver/mongo"
)

//Saveoneuser saves use data
func Saveoneuser(userdata interface{})(*mongo.InsertOneResult,error){
	client := databases.Getmongoclient();
	usercollection := client.Database("userinfodb").Collection("userdata");
	insertresult,err := usercollection.InsertOne(context.TODO(),userdata);
	return insertresult,err;

}