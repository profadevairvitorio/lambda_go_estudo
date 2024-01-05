package main

import (
	"context"
	"fmt"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func Handler(ctx context.Context, event events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	nome := event.QueryStringParameters["Test"]
	mensagem := fmt.Sprintf("teste, recebi o %s.", nome)

	return events.APIGatewayProxyResponse{
		StatusCode: 200,
		Body: mensagem,
	}, nil
}

func main() {
	lambda.Start(Handler)
}
