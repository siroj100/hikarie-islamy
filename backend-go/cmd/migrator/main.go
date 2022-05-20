package main

import (
	"log"

	"github.com/siroj100/hikarie-islamy/cmd/internal"
	"github.com/siroj100/hikarie-islamy/internal/config"
	"github.com/siroj100/hikarie-islamy/internal/model/db"
)

func main() {
	conf := config.Init()
	dbs := internal.InitGormDb(conf.Database)
	db0 := dbs[config.DbIslamy]
	err := db0.Pri().AutoMigrate(
		db.QuranL10N{}, db.QuranSurat{}, db.QuranSuratL10N{},
		db.QuranAyat{}, db.QuranAyatL10N{},
	)
	log.Println(err)
}
