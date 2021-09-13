package projectid

import "go.mongodb.org/mongo-driver/bson/primitive"

type ID string

func New() ID {
	return ID(primitive.NewObjectID().Hex())
}

func FromString(str string) (ID, error) {
	oid, err := primitive.ObjectIDFromHex(str)
	if err != nil {
		return "", err
	}

	return ID(oid.Hex()), nil
}

func (id ID) String() string {
	return string(id)
}
