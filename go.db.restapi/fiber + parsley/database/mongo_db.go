package database

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"go.db.restapi/config"
)

type MongoDataBase struct {
	mongo  *MongoInstance
	config config.ConfigLoader
}

// mongoInstance contains the Mongo client and database objects
type MongoInstance struct {
	client *mongo.Client
	Db     *mongo.Database
}

func NewMongoDataBase(config config.ConfigLoader) Database[MongoInstance] {
	return &MongoDataBase{
		config: config,
	}
}

// connect function establish a connection to database
func (m *MongoDataBase) Connect() error {
	if m.mongo == nil {
		m.config.Load()
		config := m.config.Get()
		client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(config.DB.Server))
		if err != nil {
			return err
		}
		db := client.Database(config.DB.Database)
		m.mongo = &MongoInstance{
			client,
			db,
		}
	}
	return nil
}

// connect function establish a connection to database
func (m *MongoDataBase) Disconnect() error {
	result := m.mongo.client.Disconnect(context.TODO())
	m.mongo = nil
	return result
}

func (m *MongoDataBase) Get() *MongoInstance {
	m.Connect()
	return m.mongo
}
