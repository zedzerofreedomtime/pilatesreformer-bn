package main

import (
	"context"
	"log"

	"github.com/zedzerofreedomtime/pilatesreformer/api/internal/app"
)

func main() {
	ctx := context.Background()

	application, err := app.New(ctx)
	if err != nil {
		log.Fatal(err)
	}

	if err := application.Run(); err != nil {
		log.Fatal(err)
	}
}
