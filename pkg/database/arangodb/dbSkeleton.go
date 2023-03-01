package arangodb

import (
	"context"
	"github.com/arangodb/go-driver"
	"github.com/sirupsen/logrus"
)

func SetupDatabaseSkeleton(db driver.Database, ctx context.Context) error {
	// Create Database Collections
	if err := CreateCollections(db, ctx); err != nil {
		logrus.Error("error creating collections: ", err)
		return err
	}

	// Create Custom Analyzers if they are created for use in Views
	if err := CreateACustomAnalyzer(db); err != nil {
		logrus.Error("error creating a custom analyzer")
		return err
	}

	// Create Search Views for your Collections
	if err := CreateSearchViews(db, ctx); err != nil {
		logrus.Error("error creating database views: ", err)
		return err
	}

	return nil
}

func EnsureBaseEntities() {
	//Check whether entities exists, or else create them
}
