package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/PullRequestInc/go-gpt3"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()

	if err != nil {
		log.Fatalln("Error loading .env file" + err.Error())
	}
	apikey := os.Getenv("API_KEY")
	if apikey == "" {
		log.Fatalln("missing api key" + apikey)
	}

	ctx := context.Background()

	client := gpt3.NewClient(apikey)

	resp, err := client.Completion(ctx, gpt3.CompletionRequest{

		Prompt:    []string{"First thing you should know about Golang is"},
		MaxTokens: gpt3.IntPtr(30),
		Stop:      []string{"."},
		Echo:      true,
	})

	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(resp.Choices[0].Text)
}
