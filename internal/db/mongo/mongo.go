package db

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

/**
 * This is package helper for connection MongoDB
 */

type dbMongo struct {
	Conn *mongo.Database
	Err  error
}

func (_d *dbMongo) Init() {
	_d.Conn, _d.Err = _d.Connect()
}

// Connect Execute Connect to Mongo server
func (_d *dbMongo) Connect() (*mongo.Database, error) {
	connStr, dbname := _d.ConnStr()

	// Set client options
	clientOptions := options.Client().ApplyURI(connStr)

	// Connect to MongoDB
	dbTimeout := 10 * time.Second // Default

	ctx, _ := context.WithTimeout(context.Background(), dbTimeout)
	client, err := mongo.Connect(ctx, clientOptions)

	if err != nil {
		log.Fatal(err)
	}

	// Check the connection
	connTimeout := 2 * time.Second // Default

	ctx, _ = context.WithTimeout(context.Background(), connTimeout)
	err = client.Ping(ctx, readpref.Primary())

	if err != nil {
		if os.Getenv("GO_ENV") == "development" {
			log.Fatalf("Failed when connecting to: %s, error: %v", connStr, err)
		} else {
			log.Fatalf("Failed when connecting to mongo db. error: %v", err)
		}
	}

	if os.Getenv("GO_ENV") == "development" {
		log.Println("Connected to MongoDB: ", dbname)
	} else {
		log.Println("Connected to MongoDB: ", os.Getenv("MONGO_HOST"))
	}

	return client.Database(dbname), nil
}

// ConnStr Generate connection string for mongo from Env vars
func (*dbMongo) ConnStr() (string, string) {
	var (
		host = os.Getenv("MONGO_HOST")
		/* user      = os.Getenv("MONGO_USERNAME")*/
		dbname = os.Getenv("MONGO_DBNAME")
		/*password  = url.QueryEscape(os.Getenv("MONGO_PASSWORD"))
		mongoAuth = os.Getenv("MONGO_AUTH") */
	)

	connStr := fmt.Sprintf("mongodb://%s/?readPreference=primary&directConnection=true&ssl=false", host)

	return connStr, dbname
}

var Mongo = &dbMongo{}
