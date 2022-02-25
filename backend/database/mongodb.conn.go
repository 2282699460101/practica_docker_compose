package database

import (
	"context"
	"fmt"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	usr      = "mongoadmin"
	pwd      = "So1pass1S_2022"
	host     = "localhost"
	port     = 27017
	database = "practica"
)

func GetCollection(collection string) *mongo.Collection {
	hostEnv := os.Getenv("DB_HOST")
	if hostEnv == "" {
		hostEnv = host
	}
	uri := fmt.Sprintf("mongodb://%s:%s@%s:%d", usr, pwd, hostEnv, port)

	client, err := mongo.NewClient(options.Client().ApplyURI(uri))

	if err != nil {
		panic(err.Error())
	}

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)

	if err != nil {
		panic(err.Error())
	}

	return client.Database(database).Collection(collection)
}
