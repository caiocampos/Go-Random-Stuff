package database

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"go.db.restapi/config"
)

// MongoInstance contains the Mongo client and database objects
type MongoInstance struct {
	Client *mongo.Client
	Db     *mongo.Database
}

var MongoDB *MongoInstance

// connect function establish a connection to database
func MongoConnect() error {
	if MongoDB == nil {
		config.ReadTOML()
		client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(config.TOMLConfig.DB.Server))
		if err != nil {
			return err
		}
		defer func() {
			if err := client.Disconnect(context.TODO()); err != nil {
				panic(err)
			}
		}()
		db := client.Database(config.TOMLConfig.DB.Database)
		MongoDB = &MongoInstance{
			Client: client,
			Db:     db,
		}
	}
	
	return nil
}