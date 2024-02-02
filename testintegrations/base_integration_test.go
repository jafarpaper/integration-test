package testintegrations

import (
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"integration-test/app/config/database"
	"integration-test/database/migrations"
	"integration-test/utils/arangodb"
	"integration-test/utils/constant"
	"integration-test/utils/middleware"
	"io"
	"net/http"
	"os"
)

func InitIntegration() (*gin.Engine, arangodb.ArangoDB) {
	if err := godotenv.Load(".env"); err != nil {
		panic(err)
	}

	arangoDB, err := database.InitArango()
	if err != nil {
		logrus.Error("Error init connection to Arango: ", err)
		os.Exit(1)
	}

	gin.SetMode(os.Getenv(constant.EnvGinMode))

	migrations.CreateUserCollection(arangoDB)

	g := gin.Default()

	g.Use(
		middleware.CORSMiddleware(),
		middleware.JSONMiddleware(),
		middleware.RequestId(),
	)

	return g, arangoDB
}

func createGetRequest(url string) *http.Request {
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		logrus.Fatal(err)
	}

	return req
}

func createPostRequest(url string, body io.Reader) *http.Request {
	req, err := http.NewRequest(http.MethodPost, url, body)
	if err != nil {
		logrus.Fatal(err)
	}

	return req
}

func createPutRequest(url string, body io.Reader) *http.Request {
	req, err := http.NewRequest(http.MethodPut, url, body)
	if err != nil {
		logrus.Fatal(err)
	}

	return req
}

func createPatchRequest(url string, body io.Reader) *http.Request {
	req, err := http.NewRequest(http.MethodPatch, url, body)
	if err != nil {
		logrus.Fatal(err)
	}

	return req
}

func createDeleteRequest(url string) *http.Request {
	req, err := http.NewRequest(http.MethodDelete, url, nil)
	if err != nil {
		logrus.Fatal(err)
	}

	return req
}
