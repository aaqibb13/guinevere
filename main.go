package main

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"guinevere/pkg/database/arangodb"
	"guinevere/pkg/lib/setting"
)

func main() {
	fmt.Println("This is the main function...")
	setting.Setup()
	_, err := arangodb.InitializeArangoDb()
	if err != nil {
		logrus.Error("error initializing arangodb")
	}
}
