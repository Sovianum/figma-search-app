package tagid

import "go.mongodb.org/mongo-driver/bson/primitive"

type ID primitive.ObjectID

func New() ID {
	return ID(primitive.NewObjectID())
}
