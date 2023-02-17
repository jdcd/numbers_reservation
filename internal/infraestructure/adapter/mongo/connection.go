package mongo

import (
	"context"
	"github.com/jdcd/numbers_reservation/pkg"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

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
		pkg.ErrorLogger().Fatal(err)
	}

	err = client.Ping(context.TODO(), nil)
	if err != nil {
		pkg.ErrorLogger().Fatal(err)
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
