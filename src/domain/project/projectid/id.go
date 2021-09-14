package projectid

import (
	"github.com/google/uuid"
)

type ID uuid.UUID

func New() ID {
	return ID(uuid.New())
}

func FromString(str string) (ID, error) {
	oid, err := uuid.Parse(str)
	if err != nil {
		return ID{}, err
	}

	return ID(oid), nil
}

func (id ID) String() string {
	return uuid.UUID(id).String()
}
