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

// UserRepository defines the repository for the User entity
type UserRepository struct {
	collectionName string
	started bool
}

func (u *UserRepository) connect() {
	if (!u.started) {
		if err := database.MongoConnect(); err != nil {
			log.Fatal(err)
		}
		u.collectionName = "user"
		u.started = true
	}
}

func (u *UserRepository) collection() *mongo.Collection {
	if (!u.started) {
		u.connect()
	}
	return database.MongoDB.Db.Collection(u.collectionName)
}

// FindAll method returns all users in database
func (u *UserRepository) FindAll(ctx context.Context) ([]model.User, error) {
	cursor, err := u.collection().Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	var users []model.User = make([]model.User, 0)
	if err := cursor.All(ctx, &users); err != nil {
		return nil, err
	}
	return users, err
}

// FindByName method returns a user in database
func (u *UserRepository) FindByName(ctx context.Context, name string) (model.User, error) {
	var user model.User
	err := u.collection().FindOne(ctx, bson.M{"name": name}).Decode(&user)
	return user, err
}

// FindByID method returns a user in database
func (u *UserRepository) FindByID(ctx context.Context, id string) (model.User, error) {
	var user model.User
	_id, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return user, err
	}
	err = u.collection().FindOne(ctx, bson.M{"_id": _id}).Decode(&user)
	return user, err
}

// Insert method inserts a user in database
func (u *UserRepository) Insert(ctx context.Context, user model.User) (model.User, error) {
	user.ID = primitive.NilObjectID
	collection :=  u.collection()
	result, err := collection.InsertOne(ctx, user)
	if err != nil {
		return user, err
	}
	if _id, ok := result.InsertedID.(string); ok {
		return u.FindByID(ctx, _id)
	}
	return user, err
}

// Delete method deletes a user in database
func (u *UserRepository) DeleteByObjectID(ctx context.Context, id primitive.ObjectID) error {
	result, err := u.collection().DeleteOne(ctx, bson.M{"_id": id})
	if err != nil {
		return err
	}
	// the employee might not exist
	if result.DeletedCount < 1 {
		return  errors.New("user not found")
	}
	return nil
}

// Delete method deletes a user in database
func (u *UserRepository) DeleteByID(ctx context.Context, id string) error {
	_id, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}
	return u.DeleteByObjectID(ctx, _id)
}

// Delete method deletes a user in database
func (u *UserRepository) Delete(ctx context.Context, user model.User) error {
	return u.DeleteByObjectID(ctx, user.ID)
}

// Update method updates a user in database
func (u *UserRepository) Update(ctx context.Context, user model.User) error {
	return u.collection().FindOneAndUpdate(ctx, bson.M{"_id": user.ID}, user).Err()
}
