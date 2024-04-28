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
}
