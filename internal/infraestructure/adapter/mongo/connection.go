package mongo

import (
	"context"
	"fmt"
	"log"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const queryConnectionPattern = "mongodb://%s:%s@%s:%s"

type configMongoConnection struct {
	Collection string
	Database   string
	Host       string
	Pass       string
	Port       string
	User       string
}

func GetCollection() *mongo.Collection {
	connectionData := getConnectionData()
	clientOptions := options.Client().ApplyURI(getStringConnection(connectionData))
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	}

	collection := client.Database(connectionData.Database).Collection(connectionData.Collection)
	return collection
}

func getConnectionData() configMongoConnection {
	return configMongoConnection{
		Collection: os.Getenv("DB_COLLECTION"),
		Database:   os.Getenv("DB_NAME"),
		Host:       os.Getenv("DB_HOST"),
		Pass:       os.Getenv("DB_PASS"),
		Port:       os.Getenv("DB_PORT"),
		User:       os.Getenv("DB_USER"),
	}
}

func getStringConnection(c configMongoConnection) string {
	return fmt.Sprintf(queryConnectionPattern, c.User, c.Pass, c.Host, c.Port)
}
