package arangodb

import (
	"context"
	"crypto/tls"
	"github.com/arangodb/go-driver"
	arangoHttp "github.com/arangodb/go-driver/http"
	"github.com/sirupsen/logrus"
	"guinevere/pkg/constants"
	"guinevere/pkg/lib/setting"
	"log"
)

type arangoDbConfig struct {
	endpointUrls []string
	root         string
	rootPassword string
	user         string
	password     string
	databaseName string
}

func newArangoDbConfig() arangoDbConfig {
	db := setting.DatabaseSetting
	var config arangoDbConfig

	config = arangoDbConfig{
		endpointUrls: []string{db.Host},
		root:         db.Root,
		rootPassword: db.RootPassword,
		user:         db.User,
		password:     db.UserPassword,
		databaseName: db.Name,
	}
	return config
}

var DBCollections = []string{
	"Movies",
	"Actors",
	"Directors",
	"Albums",
	"Jury",
}

func InitializeArangoDb() (driver.Database, error) {

	// Load ArangoDB config
	config := newArangoDbConfig()
	ctx := context.Background()

	// A new connection is established to the url where our database is running
	conn, err := arangoHttp.NewConnection(
		arangoHttp.ConnectionConfig{
			Endpoints: config.endpointUrls,
			TLSConfig: &tls.Config{InsecureSkipVerify: true},
	})
	if err != nil {
		logrus.Error("error establishing connection with db: ", err)
		return nil, err
	}
	
	// A new client (root) is created, which can perform the dedicated actions then
	client, err := driver.NewClient(
		driver.ClientConfig{
			Connection:     conn,
			Authentication: driver.BasicAuthentication(config.root, config.rootPassword),
		})
	if err != nil {
		logrus.Error("error creating root client: ", err)
		return nil, err
	}

	// If the database exists, skip creating it, else create a new database using specified name
	exists, err := client.DatabaseExists(ctx, config.databaseName)
	if !exists {
		db, err := initDatabase(ctx, client)
		if err != nil {
			logrus.Error("error initializing database", err)
			return nil, err
		}
		return db, nil

	} else {
		db, err := client.Database(context.Background(), config.databaseName)
		if err != nil {
			return nil, err
		}
		return db, nil
	}
}

func initDatabase(ctx context.Context, client driver.Client) (driver.Database, error) {
	dbConfig := setting.DatabaseSetting
	active := true
	db, err := client.CreateDatabase(ctx, dbConfig.Name, &driver.CreateDatabaseOptions{
		Users: []driver.CreateDatabaseUserOptions{
			{
				UserName: dbConfig.User,
				Password: dbConfig.UserPassword,
				Active:   &active,
			},
		},
	})

	if err != nil {
		logrus.Error("error creating database:", err)
		return nil, err
	}
	logrus.Info("Database created...")


	// Setup database skeletons here (Collections, Analyzers and Views)
	err = SetupDatabaseSkeleton(db, ctx)
	if err != nil {
		logrus.Error("error setting up database skeleton: ", err)
		return nil, err
	}
	return db, nil
}

// Checks whether the collection exists, if a collection does not exist
// this method creates the specified collection from DBCollections
func CreateCollections(db driver.Database, ctx context.Context) error {
	logrus.Info("checking collections...")
	for _, collection := range DBCollections {
		if exists, err := db.CollectionExists(ctx, collection); err != nil {
			logrus.Error("error fetching collection: ", err)
			return err
		} else {
			logrus.Info("collection exists: ", collection)
			if !exists {
				logrus.Info("creating collection: ", collection)
				col, err := db.CreateCollection(ctx, collection, nil)
				if err != nil {
					logrus.Error("error creating collection: ", err)
					return err
				}
				logrus.Info("collection created: ", col)
			}
		}
	}
	logrus.Println("Collections initialized...")
	return nil
}

// All the views are created using this method
func CreateSearchViews(db driver.Database, ctx context.Context) error {
	orgView, err := db.CreateArangoSearchView(ctx, constants.Actors, &driver.ArangoSearchViewProperties{
		Links: driver.ArangoSearchLinks{
			constants.ActorsView: driver.ArangoSearchElementProperties{
				Fields: map[string]driver.ArangoSearchElementProperties{
					"firstName": {
						Analyzers: []string{
							"custom_analyzer",
							"text_en",
						},
					},
					"lastName": {
						Analyzers: []string{
							"custom_analyzer",
							"text_en",
						},
					},
					"stage_name": {
						Analyzers: []string{
							"custom_analyzer",
							"text_en",
						},
					},
				},
			},
		},
	})
	if err != nil {
		log.Println("error creating view: ", err)
	}
	logrus.Println("created view: ", orgView)

	// Used Vendor View here
	directorsView, err := db.CreateArangoSearchView(ctx, constants.Directors, &driver.ArangoSearchViewProperties{
		Links: driver.ArangoSearchLinks{
			constants.DirectorsView: driver.ArangoSearchElementProperties{
				Fields: map[string]driver.ArangoSearchElementProperties{
					"name": {
						Analyzers: []string{
							"custom_analyzer",
							"text_en",
						},
					},
					"genre": {
						Analyzers: []string{
							"custom_analyzer",
							"text_en",
						},
					},
				},
			},
		},
	})
	if err != nil {
		log.Println("error creating view: ", err)
	}
	logrus.Println("created view: ", directorsView)

	albumsView, err := db.CreateArangoSearchView(ctx, constants.Albums, &driver.ArangoSearchViewProperties{
		Links: driver.ArangoSearchLinks{
			constants.AlbumsView: driver.ArangoSearchElementProperties{
				Fields: map[string]driver.ArangoSearchElementProperties{
					"name": {
						Analyzers: []string{
							"custom_analyzer",
							"text_en",
						},
					},
					"genre": {
						Analyzers: []string{
							"custom_analyzer",
							"text_en",
						},
					},
				},
			},
		},
	})
	if err != nil {
		log.Println("error creating view: ", err)
	}

	logrus.Println("created view: ", albumsView)

	logrus.Println("Views set...")
	logrus.Println("Database initialization process complete.")
	return nil
}

func CreateACustomAnalyzer(db driver.Database) error {
	// Setting up a new analyzer
	ctx := context.Background()
	var min int64 = 3
	var max int64 = 3
	var t = true
	var utf8 = driver.ArangoSearchNGramStreamUTF8

	// Example 1: Creating a custom analyzer here
	customAnalyzerDefinition := driver.ArangoSearchAnalyzerDefinition{
		Name: "custom_analyzer",
		Type: driver.ArangoSearchAnalyzerTypePipeline,
		Properties: driver.ArangoSearchAnalyzerProperties{
			Pipeline: []driver.ArangoSearchAnalyzerPipeline{
				{
					Type: driver.ArangoSearchAnalyzerTypeNGram,
					Properties: driver.ArangoSearchAnalyzerProperties{
						Min:              &min,
						Max:              &max,
						PreserveOriginal: &t,
						StreamType:       &utf8,
					},
				},
				{
					Type: driver.ArangoSearchAnalyzerTypeNorm,
					Properties: driver.ArangoSearchAnalyzerProperties{
						Locale: "en",
						Case:   driver.ArangoSearchCaseLower,
					},
				},
			},
		},
		Features: []driver.ArangoSearchAnalyzerFeature{
			driver.ArangoSearchAnalyzerFeatureFrequency,
			driver.ArangoSearchAnalyzerFeaturePosition,
			driver.ArangoSearchAnalyzerFeatureNorm,
		},
	}
	_, customAnalyzer, err := db.EnsureAnalyzer(ctx, customAnalyzerDefinition)
	if err != nil {
		logrus.Error("error ensuring analyzer exists: ", err)
		return err
	}
	logrus.Info("existed: ", customAnalyzer)
	return nil
}

func InitializeDBTransaction(readCols, writeCols []string) (driver.Database, context.Context, context.Context, driver.TransactionID, error) {
	ctx := context.Background()

	// Creating db connection here
	db, err := InitializeArangoDb()
	if err != nil {
		logrus.Error("error initializing database connection: ", err)
		return nil, nil, nil, "", err
	}

	// Specifying which transactions are accessed by a transaction
	cols := driver.TransactionCollections{
		Read:  readCols,
		Write: writeCols,
	}

	// Begin database transaction
	txnId, err := db.BeginTransaction(ctx, cols, nil)
	if err != nil {
		logrus.Error("error: ", err)
		return nil, nil, nil, "", err
	}

	// Create transaction context
	txnCtx := driver.WithTransactionID(ctx, txnId)

	return db, ctx, txnCtx, txnId, nil
}
