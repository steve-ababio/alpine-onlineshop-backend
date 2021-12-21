package databases

import (
	"config/env"
	"context"
	"log"
	"os"
	"sync"
	"time"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type mongoclient struct {
	client *mongo.Client
}

var mclient mongoclient

//Getmongoclient returns mongo client
func Getmongoclient() *mongo.Client {
	return mclient.client
}

//ConnectDB connect application to,our database
func ConnectDB(wg *sync.WaitGroup) {

	var err error
	env.Loadenvfile()
	databaseuri := os.Getenv("DATABASE_URI")
	var c mongoclient
	c.client, err = mongo.NewClient(options.Client().ApplyURI(string(databaseuri)))

	mclient = c
	if err != nil {
		log.Fatal(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10 * time.Second)
	err = c.client.Connect(ctx)

	if err != nil {
		log.Fatal(err)
	}
	defer cancel()
	//defer client.Disconnect(ctx)
	defer wg.Done()

}
