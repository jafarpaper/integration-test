package arangodb

import (
	"context"
	"integration-test/utils/errors"

	"github.com/arangodb/go-driver"
	"github.com/sirupsen/logrus"
)

type ArangoDBRepositoryInterface interface {
	Find(ID string, request ArangoInterface) error
	Create(request ArangoInterface) error
	Update(request ArangoInterface) error
	Delete(request ArangoInterface) error
}

type ArangoDBRepository struct {
	ArangoDB   ArangoDB
	Collection string
}

func NewArangoBaseRepository(arangoDB ArangoDB) ArangoDBRepositoryInterface {
	return &ArangoDBRepository{
		ArangoDB: arangoDB,
	}
}

func (r *ArangoDBRepository) Find(ID string, request ArangoInterface) error {
	query := `
		FOR data IN ` + r.Collection + `
		FILTER data._id == @data_id
		RETURN data
	`

	ctx := driver.WithQueryCount(context.Background())
	condition := map[string]interface{}{
		"data_id": ID,
	}

	data, err := r.ArangoDB.DB().Query(ctx, query, condition)
	if err != nil {
		logrus.Error(err)
		return err
	}

	if data.Count() > 0 {
		data.ReadDocument(ctx, &request)
		return nil
	}

	err = errors.New("Not found")
	return err
}

func (r *ArangoDBRepository) Create(request ArangoInterface) error {
	ctx := context.Background()
	collection, err := r.ArangoDB.DB().Collection(ctx, r.Collection)
	if err != nil {
		logrus.Error(err)
		return err
	}

	request.InitializeTimestamp()
	insert, err := collection.CreateDocument(ctx, request)
	if err != nil {
		logrus.Error(err)
		return err
	}

	request.Set(insert.ID.String(), insert.Key, insert.Rev)

	return nil
}

func (r *ArangoDBRepository) Update(request ArangoInterface) error {
	ctx := context.Background()
	collection, err := r.ArangoDB.DB().Collection(ctx, r.Collection)
	if err != nil {
		logrus.Error(err)
		return err
	}

	request.UpdateTimestamp()
	_, err = collection.UpdateDocument(ctx, request.GetKey(), request)
	if err != nil {
		logrus.Error(err)
		return err
	}

	return nil
}

func (r *ArangoDBRepository) Delete(request ArangoInterface) error {
	ctx := context.Background()
	collection, err := r.ArangoDB.DB().Collection(ctx, r.Collection)
	if err != nil {
		logrus.Error(err)
		return err
	}

	_, err = collection.RemoveDocument(ctx, request.GetKey())
	if err != nil {
		logrus.Error(err)
		return err
	}

	return nil
}
