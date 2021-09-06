package fileid

import "go.mongodb.org/mongo-driver/bson/primitive"

type ID primitive.ObjectID

func New() ID {
	return ID(primitive.NewObjectID())
}

func FromString(str string) (ID, error) {
	oid, err := primitive.ObjectIDFromHex(str)
	if err != nil {
		return ID{}, err
	}

	return ID(oid), nil
}

func (id ID) String() string {
	return primitive.ObjectID(id).String()
}
