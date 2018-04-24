package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

type SubmissionData struct {
	ID string `json:"id"`
	IP string `json:"ip"`
}

type Submission struct {
	Number int            `json:"number"`
	Data   SubmissionData `json:"data"`
}

func handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	key := os.Getenv("API_KEY")
	formID := os.Getenv("LIKES_FORM_ID")
	id := request.QueryStringParameters["id"]

	req, _ := http.NewRequest("GET", "https://api.netlify.com/api/v1/forms/"+formID+"/submissions", nil)
	req.Header.Set("Authorization", "Bearer "+key)

	client := http.Client{}
	resp, _ := client.Do(req)
	body, _ := ioutil.ReadAll(resp.Body)

	submissions := make([]Submission, 0)
	json.Unmarshal(body, &submissions)

	fmt.Println(submissions)
	numberOfLikes := 0
	for _, submission := range submissions {
		fmt.Println(submission)
		fmt.Println(id)
		if submission.Data.ID == id {
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
