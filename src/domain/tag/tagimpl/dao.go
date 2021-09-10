package tagimpl

import "github.com/aws/aws-sdk-go/service/dynamodb/dynamodbiface"

func NewDAO(dbApi dynamodbiface.DynamoDBAPI) *dao {
	return &dao{}
}

type dao struct {
	db dynamodbiface.DynamoDBAPI
}

func (dao *dao) InsertTags() {

}
