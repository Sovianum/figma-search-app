package tagid

import "go.mongodb.org/mongo-driver/bson/primitive"

type ID string

func New() ID {
	return ID(primitive.NewObjectID().Hex())
}
