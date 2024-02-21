package main

import (
	"fmt"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	var greeting string
	sourceIP := request.RequestContext.Identity.SourceIP

	if sourceIP == "" {
		fmt.Println("tadhSSSSidddddsd is sanssssseaawassasdass from the Usdzaasdasdsdfsssdfsddser lambdassdfsdfsddasda!asdasd!!!!SSSSSasdfa!")

		greeting = "Hadasello, world!\n"
	} else {
		greeting = fmt.Sprintf("He2adsdasdassccsdasssdasdllo, %s!\n", sourceIP)
	}

	return events.APIGatewayProxyResponse{
		Body:       greeting,
		StatusCode: 200,
	}, nil
}

func main() {
	lambda.Start(handler)
}
