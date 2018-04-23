package main

import (
	"context"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-lambda-go/lambdacontext"
)

type SubmissionData struct {
	Path string
	IP   string
}

type Submission struct {
	Number int
	Data   SubmissionData
}

func handler(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	lc, ok := lambdacontext.FromContext(ctx)
	if !ok {
		return events.APIGatewayProxyResponse{
			StatusCode: 503,
			Body:       "Something went wrong :(",
		}, nil
	}

	cc := lc.ClientContext
	key := cc.Env["API_KEY"]
	path := request.QueryStringParameters["path"]

	req, _ := http.NewRequest("GET", "https://api.netlify.com/api/v1/forms/5ade3140e4708575eb7932d4/submissions", nil)
	req.Header.Set("Authorization", "Authorization: Bearer "+key)

	client := http.Client{}
	resp, _ := client.Do(req)

	submissions := []Submission{}

	_ = json.NewDecoder(resp.Body).Decode(submissions)

	numberOfLikes := 0
	for _, submission := range submissions {
		if submission.Data.Path == path {
			numberOfLikes = numberOfLikes + 1
		}
	}

	return events.APIGatewayProxyResponse{
		StatusCode: 200,
		Body:       strconv.Itoa(numberOfLikes),
	}, nil
}

func main() {
	// Make the handler available for Remote Procedure Call by AWS Lambda
	lambda.Start(handler)
}
