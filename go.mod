module github.com/Sovianum/figma-search-app

go 1.16

replace github.com/Sovianum/figma-search-app => ./

require (
	github.com/aws/aws-lambda-go v1.26.0
	github.com/aws/aws-sdk-go v1.40.41
	github.com/davecgh/go-spew v1.1.1
	github.com/davyzhang/agw v0.0.0-20200908000401-cd09bf93af20
	github.com/go-zoo/bone v1.3.0 // indirect
	github.com/google/uuid v1.3.0
	github.com/google/wire v0.5.0
	github.com/gorilla/mux v1.8.0
	github.com/joomcode/errorx v1.0.3
	github.com/json-iterator/go v1.1.11 // indirect
	github.com/justinas/alice v1.2.0 // indirect
	github.com/stretchr/testify v1.7.0
	go.mongodb.org/mongo-driver v1.7.2
)
