package main

import (
	"log"

	"github.com/GeekCAMPvol2/QuizAPI/api"
	"github.com/GeekCAMPvol2/QuizAPI/db"
	"github.com/GeekCAMPvol2/QuizAPI/util"
)

func main() {
	config, err := util.Loadenv(".")

	if err != nil {
		log.Fatal("app.envファイルが読み込めませんでした")
	}
	coll := db.NewClient(config.MongoTestUri)

	if err != nil {
		log.Fatal("Mongo Connection Error :", err)
	}

	if err := api.NewServer(coll).Start(config.ServerAddress); err != nil {
		log.Fatal(err)
	}
}
