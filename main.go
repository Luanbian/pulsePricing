package main

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/Luanbian/pulsePricing/config"
)

func main() {
	dbConnection := config.SetupDatabase()
	defer func() {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		if err := dbConnection.Disconnect(ctx); err != nil {
			log.Fatal("Error disconnecting from MongoDB:", err)
		}
	}()

	log.Fatal(http.ListenAndServe(":8080", nil))
}