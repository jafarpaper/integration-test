package database

import (
	"context"
	"github.com/arangodb/go-driver"
	"integration-test/utils/arangodb"
	"log"
)

func CreateDatabase(client arangodb.ArangoDB) driver.Database {
	ctx := context.Background()

	// Specify the collection name for users
	databaseName := "data"

	// Create a new collection
	col, err := client.Client().Database(ctx, databaseName)
	if err != nil {
		log.Fatal(err)
	}

	return col
}
