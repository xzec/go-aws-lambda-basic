package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/aws/aws-lambda-go/lambda"
	"io"
	"log"
	"net/http"
)

type Quote struct {
	QuoteText   string `json:"quoteText"`
	QuoteAuthor string `json:"quoteAuthor"`
	SenderName  string `json:"senderName"`
	SenderLink  string `json:"senderLink"`
	QuoteLink   string `json:"quoteLink"`
}

type MyEvent struct {
	Name string `json:"name"`
}

func (e MyEvent) String() string {
	return fmt.Sprintf("Event name: %v", e.Name)
}

func HandleLambdaEvent(ctx context.Context, event *MyEvent) (*Quote, error) {
	if event == nil {
		return nil, fmt.Errorf("received nil event")
	}
	fmt.Printf("Received an event %v\n", event)

	const quoteApiEndpoint = "https://api.forismatic.com/api/1.0/?method=getQuote&key=111111&format=json&lang=en"
	res, err := http.Get(quoteApiEndpoint)
	if err != nil {
		log.Fatalf("Failed to fetch Quote API: %v.", err)
	}

	defer func(Body io.ReadCloser) {
		if err = Body.Close(); err != nil {
			log.Fatalf("Failed to close the response body: %v.", err)
		}
	}(res.Body)

	raw, err := io.ReadAll(res.Body)
	if err != nil {
		log.Fatalf("Failed to read response body: %v", err)
	}

	sanitized := bytes.ReplaceAll(raw, []byte(`\'`), []byte(`'`))

	var quote Quote
	if err = json.Unmarshal(sanitized, &quote); err != nil {
		log.Fatalf("Failed to unmarshal res body: %v.", err)
	}

	fmt.Println(quote)

	return &quote, nil
}

func main() {
	lambda.Start(HandleLambdaEvent)
}
