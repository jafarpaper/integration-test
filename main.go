// main.go
package main

import (
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"integration-test/app/config/database"
	"os"
	"sync"
)

// Item represents a simple data structure
type Item struct {
	ID    string  `json:"id,omitempty"`
	Name  string  `json:"name,omitempty"`
	Price float64 `json:"price,omitempty"`
}

var items map[string]Item

func init() {
	items = make(map[string]Item)
}

func main() {
	if err := godotenv.Load(".env"); err != nil {
		panic(err)
	}
	wg := sync.WaitGroup{}

	arangoDB, err := database.InitArango()
	if err != nil {
		logrus.Error("Error init connection to Arango: ", err)
		os.Exit(1)
	}

	wg.Add(1)
	go func() {
		defer wg.Done()
		logrus.Info("Starting Paper Fintech HTTP handler")
		MainHttpHandler(arangoDB)
	}()

	wg.Wait()
}
