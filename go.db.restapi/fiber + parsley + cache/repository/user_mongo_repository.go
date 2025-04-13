package repository

import (
	"context"
	"errors"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"

	"go.db.restapi/database"
	"go.db.restapi/model"
)

// UserMongoRepository defines the repository for the User entity
type UserMongoRepository struct {
	collectionName string
	started        bool
	database       database.Database[database.MongoInstance]
}

func NewUserMongoRepository(database database.Database[database.MongoInstance]) UserRepository[model.User] {
	return &UserMongoRepository{
		collectionName: "user",
		database:       database,
		started:        false,
	}
}

func (u *UserMongoRepository) connect() {
	if !u.started {
		if err := u.database.Connect(); err != nil {
			log.Fatal(err)
		}
		u.started = true
	}
}

func (u *UserMongoRepository) disconnect() {
	if u.started {
		if err := u.database.Disconnect(); err != nil {
			log.Fatal(err)
		}
		u.started = false
	}
}

func (u *UserMongoRepository) collection() *mongo.Collection {
	if !u.started {
		u.connect()
	}
	return u.database.Get().Db.Collection(u.collectionName)
}

func (u *UserMongoRepository) GetDBType() string {
	return "mongo"
}

// FindAll method returns all users in database
func (u *UserMongoRepository) FindAll(ctx context.Context) ([]model.User, error) {
	cursor, err := u.collection().Find(ctx, bson.M{})
	defer u.disconnect()
	if err != nil {
		return nil, err
	}
	users := []model.User{}
	if err := cursor.All(ctx, &users); err != nil {
		return nil, err
	}
	return users, err
}

// FindByName method returns a user in database
func (u *UserMongoRepository) FindByName(ctx context.Context, name string) (model.User, error) {
	user := model.User{}
	err := u.collection().FindOne(ctx, bson.M{"name": name}).Decode(&user)
	defer u.disconnect()
	return user, err
}

// FindByID method returns a user in database
func (u *UserMongoRepository) FindByID(ctx context.Context, id string) (model.User, error) {
	user := model.User{}
	_id, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return user, err
	}
	err = u.collection().FindOne(ctx, bson.M{"_id": _id}).Decode(&user)
	defer u.disconnect()
	return user, err
}

// Insert method inserts a user in database
func (u *UserMongoRepository) Insert(ctx context.Context, user model.User) (model.User, error) {
	user.ID = primitive.NilObjectID
	collection := u.collection()
	defer u.disconnect()
	result, err := collection.InsertOne(ctx, user)
	if err != nil {
		return user, err
	}
	if _id, ok := result.InsertedID.(string); ok {
		return u.FindByID(ctx, _id)
	}
	return user, err
}

// deleteByObjectID method deletes a user from database
func (u *UserMongoRepository) deleteByObjectID(ctx context.Context, id primitive.ObjectID) error {
	result, err := u.collection().DeleteOne(ctx, bson.M{"_id": id})
	defer u.disconnect()
	if err != nil {
		return err
	}
	// the user might not exist
	if result.DeletedCount < 1 {
		return errors.New("user not found")
	}
	return nil
}

// DeleteByID method deletes a user in database
func (u *UserMongoRepository) DeleteByID(ctx context.Context, id string) error {
	_id, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}
	return u.deleteByObjectID(ctx, _id)
}

// Delete method deletes a user from database
func (u *UserMongoRepository) Delete(ctx context.Context, user model.User) error {
	return u.deleteByObjectID(ctx, user.ID)
}

// Update method updates a user in database
func (u *UserMongoRepository) Update(ctx context.Context, user model.User) error {
	result := u.collection().FindOneAndUpdate(ctx, bson.M{"_id": user.ID}, user).Err()
	defer u.disconnect()
	return result
}
