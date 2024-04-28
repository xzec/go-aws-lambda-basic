package main

import (
	"bytes"
	"encoding/json"
	"fmt"
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

func main() {
	const quoteApiEndpoint = "https://api.forismatic.com/api/1.0/?method=getQuote&key=111111&format=json&lang=en"
	res, err := http.Get(quoteApiEndpoint)
	if err != nil {
		log.Fatalf("Failed to fetch Quote API: %v.", err)
	}

	var buf bytes.Buffer
	if _, err := io.Copy(&buf, res.Body); err != nil {
		log.Fatalf("Failed to read response body: %v", err)
	}

	sanitizedJSON := buf.Bytes()
	sanitizedJSON = bytes.ReplaceAll(sanitizedJSON, []byte(`\'`), []byte(`'`))

	defer func(Body io.ReadCloser) {
		if cErr := Body.Close(); cErr != nil {
			log.Fatalf("Failed to close the response body: %v.", cErr)
		}
	}(res.Body)

	var q Quote
	fmt.Printf("%s\n", sanitizedJSON)
	if qErr := json.Unmarshal(sanitizedJSON, &q); qErr != nil {
		log.Fatalf("Failed to unmarshal res body: %v.", qErr)
	}

	fmt.Println(q)
}
