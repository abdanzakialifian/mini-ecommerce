package main

import (
	"context"
	"log"
	"mini-ecommerce/internal/database"
)

func main() {
	ctx := context.Background()
	conn, err := database.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close(ctx)
}
