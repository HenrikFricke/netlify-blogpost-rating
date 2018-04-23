package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strconv"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

type SubmissionData struct {
	Path string
	IP   string
}

type Submission struct {
	Number int
	Data   SubmissionData
}

func handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	key := os.Getenv("API_KEY")
	fmt.Println(key)
	path := request.QueryStringParameters["path"]

	req, _ := http.NewRequest("GET", "https://api.netlify.com/api/v1/forms/5ade3140e4708575eb7932d4/submissions", nil)
	req.Header.Set("Authorization", "Authorization: Bearer "+key)

	client := http.Client{}
	resp, _ := client.Do(req)

	submissions := []Submission{}

	_ = json.NewDecoder(resp.Body).Decode(submissions)

	numberOfLikes := 0
	fmt.Println(submissions)
	for _, submission := range submissions {
		fmt.Println(submission)
		fmt.Println(path)
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
