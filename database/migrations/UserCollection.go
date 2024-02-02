package migrations

import (
	"context"
	"github.com/arangodb/go-driver"
	"github.com/google/uuid"
	"github.com/noldwidjaja/slate/arango"
	"integration-test/app/pkg/user/domain/model/aggregate"
	"integration-test/database/database"
	"integration-test/utils/arangodb"
)

func CreateUserCollection(client arangodb.ArangoDB) {
	var collect driver.Collection
	ctx := context.Background()

	// Specify the collection name for users
	collectionName := "users"

	databaseData := database.CreateDatabase(client)

	exists, err := databaseData.CollectionExists(ctx, collectionName)
	if err != nil {
		return
	}

	if exists == true {
		collect, err = databaseData.Collection(ctx, collectionName)
		if err != nil {
			return
		}
	} else {
		collect, err = databaseData.CreateCollection(ctx, collectionName, nil)
		if err != nil {
			return
		}
	}

	err = collect.Truncate(ctx)
	if err != nil {
		return
	}
	parse, err := uuid.Parse("ed914adc-180e-4a74-abc2-82f2ca14cb8f")
	if err != nil {
		return
	}
	_, err = collect.CreateDocument(ctx, aggregate.User{
		DocumentModel: arango.DocumentModel{},
		Id:            parse,
		Name:          "test",
		Address:       "test",
		Email:         "test@test.com",
		PhoneNumber:   "019231231",
	})
	if err != nil {
		return
	}
}
