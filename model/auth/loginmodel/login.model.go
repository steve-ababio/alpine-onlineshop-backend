package loginmodel

import (
	"context"
	"databases"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

//Findoneuser retrieves user info from database
func Findoneuser(email string) *mongo.SingleResult{
	client := databases.Getmongoclient();
    usercollection := client.Database("products_db").Collection("user_info");
	filter := bson.D{primitive.E{Key:"email",Value:email}};
	userinfo := usercollection.FindOne(context.TODO(),filter,options.FindOne().SetProjection(bson.M{"_id":1}));
	return userinfo;
} 
