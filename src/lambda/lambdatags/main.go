package main

import (
	"github.com/Sovianum/figma-search-app/src/client"
	"github.com/aws/aws-lambda-go/lambda"
)

func show() (*client.Node, error) {
	return &client.Node{
		ID: "id_1",
	}, nil
}

func main() {
	lambda.Start(show)
}
