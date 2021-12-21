package checkout

import (
	"context"
	"databases"
	"log"

	"go.mongodb.org/mongo-driver/bson"
)

type itemdbinfo struct{
	ProductID string `bson:"productID"`;
	Title string `bson:"Title"`
	Price int64 `bson:"price"`;
	Image string `bson:"image"`;
}

type Itemdbinfo = itemdbinfo;

func Getproducts(prod_id []string)[]*itemdbinfo{
	client := databases.Getmongoclient();
	productscollection := client.Database("products_db").Collection("products_coll");
	
	filter := bson.D{{Key:"productID",Value:bson.D{{Key:"$in",Value:prod_id}}}};

	results := []*itemdbinfo{};
	productsdata,err := productscollection.Find(context.TODO(),filter);

	if err != nil {
	 log.Fatal(err)
	}

	for productsdata.Next(context.TODO()){
	  var item itemdbinfo;
	  err := productsdata.Decode(&item);
	  if err != nil{
		log.Fatal(err);
	  }
	  results = append(results,&item);
	}
	return results;
} 