package mongo

import (
	"context"
	"log"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const queryConnectionPattern = "mongodb://%s:%s@%s:%s"

type configMongoConnection struct {
	Collection string
	Database   string
	Url        string
}

func GetCollection() *mongo.Collection {
	cd := getConnectionData()
	clientOptions := options.Client().ApplyURI(cd.Url)
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	}

	collection := client.Database(cd.Database).Collection(cd.Collection)
	return collection
}

func getConnectionData() configMongoConnection {
	return configMongoConnection{
		Collection: os.Getenv("DB_COLLECTION"),
		Database:   os.Getenv("DB_NAME"),
		Url:        os.Getenv("MONGO_URL"),
	}
}
