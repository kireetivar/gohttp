package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"os"

	"github.com/kireetivar/gohttp"
	"github.com/kireetivar/gohttp/internal/config"
)

func main() {
	cfg := config.Load()
	if cfg.Url == "" {
		fmt.Fprintln(os.Stderr, "error url cannot be empty")
		os.Exit(1)
	}

	ctx := context.Background()
	res, err := gohttp.Get(ctx, cfg.Url)
	if err != nil {
		fmt.Println("error get call failed")
	}
	defer res.Body.Close()

	bodyBytes, err := io.ReadAll(res.Body)
	if err != nil {
		log.Fatalf("Failed to read response body: %v", err)
	}
	fmt.Println(string(bodyBytes))
}
