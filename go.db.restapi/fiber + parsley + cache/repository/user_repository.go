package repository

import "go.mongodb.org/mongo-driver/bson/primitive"

type UserRepository[Model any] interface {
	Repository
	FindAll[Model]
	FindByName[Model]
	FindByID[Model]
	Insert[Model]
	DeleteByID[Model]
	Delete[Model]
	Update[Model]
}

type UserCacheRepository[Model any] interface {
	Repository
	FindByID[Model]
	Insert[Model]
	DeleteByID[Model]
	Delete[Model]
	Update[Model]
}

const userPF = "user"

func getUserKey(id string) string {
	return getKey(userPF, id)
}

func getUserOIDKey(id primitive.ObjectID) string {
	return getKey(userPF, id.Hex())
}
