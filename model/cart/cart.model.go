package cart

import (
	"context"
	"databases"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func getcollection() *mongo.Collection {
	client := databases.Getmongoclient()
	return client.Database("products_db").Collection("user_cart")
}
func Insertusercartitems(cartitems interface{}) (*mongo.InsertOneResult, error) {
	cart_coll := getcollection()
	result, err := cart_coll.InsertOne(context.TODO(), cartitems)
	return result, err
}
func Findusercart(user_id string) *mongo.SingleResult {
	collection := getcollection()
	result := collection.FindOne(context.TODO(), bson.M{"user_id": user_id},options.FindOne().SetProjection(bson.D{{"user_id",0},{"_id",0}}));
	return result;
}
func Findusercartitem(user_id string,product_id string) *mongo.SingleResult{
	collection := getcollection();
	result := collection.FindOne(context.TODO(),bson.M{"user_id":user_id,"Cart.productid":product_id});
	fmt.Println("Error:",result.Err())
	return result;
}
func Updateusercart(userID string, cartitem interface{}) (interface{}, error) {
	fmt.Println("userID",userID)
	collection := getcollection();

	result, err := collection.UpdateOne(context.TODO(),
		bson.M{"user_id": userID},
		bson.M{"$push":bson.M{"Cart":cartitem}},
	);
	return result.MatchedCount, err
}
