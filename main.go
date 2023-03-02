package main

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"guinevere/pkg/database/arangodb"
	"guinevere/pkg/lib/setting"
)

func main() {
	fmt.Println("This is the main function...")

	// Read config and load it
	setting.Setup()

	// Initialize ArangoDB
	_, err := arangodb.InitializeArangoDb()
	if err != nil {
		logrus.Error("error initializing arangodb")
	}
}
