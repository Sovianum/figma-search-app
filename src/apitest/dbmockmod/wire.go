package dbmockmod

import (
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbiface"
	"github.com/google/wire"
	"github.com/gusaul/go-dynamock"
)

var M = wire.NewSet(
	NewDB,
	NewMock,
	NewDBWithMock,
)

func NewDB(dbWithMock *DBWithMock) dynamodbiface.DynamoDBAPI {
	return dbWithMock.DB
}

func NewMock(dbWithMock *DBWithMock) *dynamock.DynaMock {
	return dbWithMock.Mock
}

type DBWithMock struct {
	DB   dynamodbiface.DynamoDBAPI
	Mock *dynamock.DynaMock
}

func NewDBWithMock() *DBWithMock {
	db, mock := dynamock.New()

	return &DBWithMock{
		DB:   db,
		Mock: mock,
	}

}
