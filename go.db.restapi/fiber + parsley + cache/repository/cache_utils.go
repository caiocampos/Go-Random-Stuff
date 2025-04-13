package repository

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func getKey(prefix string, id string) string {
	return prefix + "-" + id
}

func getOIDKey(prefix string, id primitive.ObjectID) string {
	return getKey(prefix, id.Hex())
}
