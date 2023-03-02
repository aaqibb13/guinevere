package setting

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"log"
)

type Database struct {
	Type         	string
	User         	string
	UserPassword 	string
	Host         	string
	Name         	string
	Root         	string
	RootPassword    string
}

type Collection struct {
	Movies    string
	Actors    string
	Directors string
	Albums    string
	Jury      string
}

type CollectionView struct {
	MoviesView    string
	ActorsView    string
	DirectorsView string
	AlbumsView    string
	JuryView      string
}

var DatabaseSetting = &Database{}
var Collections = &Collection{}
var CollectionViews = &CollectionView{}

// Setup initialize the configuration instance
func Setup() {
	var err error
	viper.SetConfigFile(".env")
	err = viper.ReadInConfig()
	if err != nil {
		log.Fatalf("cannot read the env file, please check : %s", err)
		return
	}
	setupDB()
	setupCollections()
	PORT := viper.GetString("PORT")
	logrus.Info("Port is: ", PORT)
}

func setupDB() {
	db := Database{
		User:         	viper.GetString("DATABASE_USER"),
		UserPassword: 	viper.GetString("DATABASE_PASSWORD"),
		Root:         	viper.GetString("ARANGO_ROOT"),
		RootPassword:   viper.GetString("ARANGO_ROOT_PASSWORD"),
		Host:         	viper.GetString("DATABASE_HOST"),
		Name:         	viper.GetString("DATABASE_NAME"),
	}
	log.Println("database set...")
	DatabaseSetting = &db
}

func setupCollections() {
	c := Collection{
		Actors:    viper.GetString("ACTORS_COLLECTION"),
		Movies:    viper.GetString("MOVIES_COLLECTION"),
		Directors: viper.GetString("DIRECTORS_COLLECTION"),
		Albums:    viper.GetString("ALBUM_COLLECTION"),
		Jury:      viper.GetString("JURY_COLLECTION"),
	}
	Collections = &c

	v := CollectionView{
		ActorsView:    viper.GetString("ACTORS_VIEW"),
		MoviesView:    viper.GetString("MOVIES_VIEW"),
		DirectorsView: viper.GetString("DIRECTORS_VIEW"),
		AlbumsView:    viper.GetString("ALBUMS_VIEW"),
		JuryView:      viper.GetString("JURY_VIEW"),
	}

	CollectionViews = &v
}
