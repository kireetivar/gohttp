package main

import (
	"context"
	"fmt"
	"io"
	"log"

	"github.com/kireetivar/gohttp"
)

func main() {
	ctx := context.Background()
	res, err := gohttp.Get(ctx, "http://info.cern.ch:80/")
	if err != nil {
		fmt.Println("ERROR DEATH")
	}
	defer res.Body.Close()

	// 2. Read all contents from the stream into a byte slice
	bodyBytes, err := io.ReadAll(res.Body)
	if err != nil {
		log.Fatalf("Failed to read response body: %v", err)
	}
	fmt.Println(string(bodyBytes))
}
