package db

import (
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"os"
	"context"
	"time"
	"io/ioutil"
	log "github.com/sirupsen/logrus"
)

func Connect() *mongo.Client {
	
	var DB_CONNECTION string
	
	if os.Getenv("DB_CONNECTION")  == "" {
		data,_ := ioutil.ReadFile("config.txt")
		DB_CONNECTION = string(data)
	} else {
		DB_CONNECTION = os.Getenv("DB_CONNECTION")
	}
	
	log.Info("Starting DB connection")
	ctx, cancel := context.WithTimeout(context.TODO(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(DB_CONNECTION))
	if err != nil { log.Fatal(err) }

	log.Info("Successfully connected")
	log.Info(client)

	return client
}
