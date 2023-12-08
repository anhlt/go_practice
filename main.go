package main

import (
	"context"
	"fmt"
	"goroutine_practice/helper"
	"net/http"
	"time"

	"golang.org/x/sync/errgroup"
)

func longRunning(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("`longRunning` has been cancelled")
			return
		case <-time.After(1 * time.Second):
			helper.Help()
		}
	}
}

func main() {

	// Call the Help() function from the helper package

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)

	defer cancel()

	go longRunning(ctx)

	time.Sleep(10 * time.Second)

	// create err errgroup
	var g errgroup.Group

	var urls = []string{
		"https://www.google.com",
		"https://golang.org",
	}

	for _, url := range urls {
		// Launch a goroutine to fetch the URL.
		url := url // https://golang.org/doc/faq#closures_and_goroutines
		g.Go(func() error {
			// Fetch the URL.
			resp, err := http.Get(url)
			if err == nil {
				resp.Body.Close()
				fmt.Println("Successfully fetched URL:", url)
			}
			return err
		})
	}

	if err := g.Wait(); err == nil {
		fmt.Println("Successfully fetched all URLs.")
	} else {
		fmt.Println("Failed to fetch all URLs.")
	}

}
