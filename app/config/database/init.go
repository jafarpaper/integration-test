package database

import (
	"integration-test/utils/arangodb"
	"integration-test/utils/constant"
	"os"

	"github.com/sirupsen/logrus"
)

func InitArango() (arangodb.ArangoDB, error) {
	arangoDB, err := arangodb.NewArangoDB(
		os.Getenv(constant.EnvArangoDBHost),
		os.Getenv(constant.EnvArangoDBDatabase),
		os.Getenv(constant.EnvArangoDBUser),
		os.Getenv(constant.EnvArangoDBPassword),
	)

	if err != nil {
		logrus.Error("Error init connection to Arango: ", err)
		return nil, err
	}

	return arangoDB, nil
}
