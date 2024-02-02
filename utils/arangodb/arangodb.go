package arangodb

import (
	"context"
	"fmt"
	"integration-test/utils/models"
	"sync"

	"github.com/arangodb/go-driver"
	"github.com/arangodb/go-driver/http"
)

type arangoDB struct {
	client driver.Client
	db     driver.Database
}

type ArangoDB interface {
	Client() driver.Client
	DB() driver.Database
	ParseFilterQuery(request models.FilterRequest, searchColumns []string) (string, map[string]interface{})
	Count(query string, args map[string]interface{}) (int64, error)
	Get(query string, args map[string]interface{}) ([]interface{}, error)
	Find(query string, args map[string]interface{}, output interface{}) error
	InitCollections(cols []string) error
}

func NewArangoDB(host, database, username, password string) (ArangoDB, error) {
	conn, err := http.NewConnection(http.ConnectionConfig{
		Endpoints: []string{host},
		ConnLimit: 32,
	})
	if err != nil {
		return nil, fmt.Errorf("Error when creating new connection to Arango: %s", err.Error())
	}

	client, err := driver.NewClient(driver.ClientConfig{
		Connection:     conn,
		Authentication: driver.BasicAuthentication(username, password),
	})
	if err != nil {
		return nil, fmt.Errorf("Error when creating new Arango client: %s", err.Error())
	}

	ctx := driver.WithQueryCount(context.Background())
	db, err := client.Database(ctx, database)
	if err != nil {
		return nil, fmt.Errorf("Error when opening database connection: %s", err.Error())
	}

	return &arangoDB{
		client: client,
		db:     db,
	}, nil
}

func (a *arangoDB) Client() driver.Client {
	return a.client
}

func (a *arangoDB) DB() driver.Database {
	return a.db
}

func (a *arangoDB) ParseFilterQuery(request models.FilterRequest, searchColumns []string) (string, map[string]interface{}) {
	filter := ""
	args := map[string]interface{}{}

	// Loop through the search column that we want.
	for _, col := range searchColumns {
		for key, el := range request.Filters {
			// Compare if the key in request is exists in search column.
			if key != "global" && col == key {
				element := el.(map[string]interface{})
				value := element["value"]
				mode := element["mode"]

				if value != nil && value != "" {
					if mode == "like" {
						filter += fmt.Sprintf(`FILTER row["%s"] LIKE "%%%s%%"`, key, value)
					} else {
						filter += fmt.Sprintf(`FILTER row["%s"] == @_%s `, key, key)
						args["_"+key] = value
					}
				}

				break
			}
		}
	}

	if global := request.Filters["global"]; global != nil {
		global := global.(map[string]interface{})
		value := ""

		if global["value"] != nil {
			value = global["value"].(string)
		}

		if value != "" {
			n := 0
			for _, col := range searchColumns {
				if n == 0 {
					filter += fmt.Sprintf(`FILTER row["%s"] LIKE "%%%s%%" `, col, value)
				} else {
					filter += fmt.Sprintf(`OR row["%s"] LIKE "%%%s%%" `, col, value)
				}

				n++
			}
		}
	}

	return filter, args
}

func (a *arangoDB) Count(query string, args map[string]interface{}) (int64, error) {
	var count int64

	context := driver.WithQueryCount(context.Background())
	data, err := a.DB().Query(context, query, args)
	if err != nil {
		return count, err
	}

	if data.Count() > 0 {
		_, err := data.ReadDocument(context, &count)

		if err != nil {
			return count, err
		}
	}

	return count, nil
}

func (a *arangoDB) Get(query string, args map[string]interface{}) ([]interface{}, error) {
	var rows []interface{}

	context := driver.WithQueryCount(context.Background())
	data, err := a.DB().Query(context, query, args)
	if err != nil {
		return rows, err
	}

	for data.HasMore() {
		var row interface{}

		data.ReadDocument(context, &row)
		rows = append(rows, row)
	}

	return rows, nil
}

func (a *arangoDB) Find(query string, args map[string]interface{}, output interface{}) error {
	context := driver.WithQueryCount(context.Background())
	data, err := a.DB().Query(context, query, args)
	if err != nil {
		return err
	}

	if data.Count() > 0 {
		data.ReadDocument(context, &output)
	}

	return nil
}

func (a *arangoDB) InitCollections(cols []string) error {
	wg := sync.WaitGroup{}
	for _, col := range cols {
		wg.Add(1)
		go func() {
			defer wg.Done()
			ctx := context.Background()
			//options := &driver.CreateCollectionOptions{ Type: }
			colExist, err := a.db.CollectionExists(ctx, col)
			if err == nil && !colExist {
				a.db.CreateCollection(ctx, col, nil)
			}
		}()
	}
	wg.Wait()
	return nil
}
