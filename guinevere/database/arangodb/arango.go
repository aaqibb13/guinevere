package arangodb

import (
	"context"
	"crypto/tls"
	"github.com/arangodb/go-driver"
	arangoHttp "github.com/arangodb/go-driver/http"
	"github.com/sirupsen/logrus"
	"guinevere/constants"
	"guinevere/lib/setting"
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
		rootPassword: db.Password,
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
	config := newArangoDbConfig()
	ctx := context.Background()
	conn, err := arangoHttp.NewConnection(arangoHttp.ConnectionConfig{
		Endpoints: config.endpointUrls,
		TLSConfig: &tls.Config{InsecureSkipVerify: true},
	})
	if err != nil {
		return nil, err
	}

	client, err := driver.NewClient(driver.ClientConfig{
		Connection:     conn,
		Authentication: driver.BasicAuthentication(config.user, config.password),
	})
	if err != nil {
		return nil, err
	}

	exists, err := client.DatabaseExists(ctx, config.databaseName)
	if !exists {
		db, err := initDatabase(ctx, conn)
		if err != nil {
			log.Println("error in initializing database", err)
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

func initDatabase(ctx context.Context, conn driver.Connection) (driver.Database, error) {
	log.Println("Initializing database...")
	dbConfig := setting.DatabaseSetting
	client, err := driver.NewClient(driver.ClientConfig{
		Connection: conn,
		Authentication: driver.
			BasicAuthentication(dbConfig.Root, dbConfig.Password),
	})

	if err != nil {
		log.Println("error in connection using root password", err)
		return nil, err
	}
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
		log.Println("error creating database...", err)
		return nil, err
	}
	log.Println("Database created...")
	log.Println("New User created...")
	return db, nil
}

// This method is used to create collections
func CreateCollections(db driver.Database, ctx context.Context) error {
	logrus.Info("checking collections")
	for _, collection := range DBCollections {
		if exists, err := db.CollectionExists(ctx, collection); err != nil {
			log.Println("error fetching collection...", err)
			return err
		} else {
			log.Println("collection exists: ", collection)
			if !exists {
				logrus.Info("creating collection: ", collection)
				col, err := db.CreateCollection(ctx, collection, nil)
				if err != nil {
					logrus.Println("error creating collection: ", err)
					return err
				}
				logrus.Println("collection created: ", col)
			}
		}
	}
	log.Println("Collections initialized...")
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
