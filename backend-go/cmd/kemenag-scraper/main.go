package main

import (
	"context"
	"log"

	"github.com/siroj100/hikarie-islamy/cmd/internal"
	"github.com/siroj100/hikarie-islamy/internal/config"
	"github.com/siroj100/hikarie-islamy/internal/model/db"
	"github.com/siroj100/hikarie-islamy/pkg/errorx"
)

func main() {
	conf := config.Init()
	dbs := internal.InitGormDb(conf.Database)
	ucase := internal.InitUseCase(conf, dbs)

	db0 := dbs[config.DbIslamy]
	err := db0.Pri().AutoMigrate(
		db.QuranL10N{}, db.QuranSurat{}, db.QuranSuratL10N{},
		db.QuranAyat{}, db.QuranAyatL10N{},
	)
	if err != nil {
		log.Fatalln(errorx.PrintTrace(err))
	}

	ctx := context.TODO()
	listSurat, err := ucase.ScrapeKemenagSurat(ctx)
	if err != nil {
		log.Fatalln(errorx.PrintTrace(err))
	}
	mapSuratAyat, err := ucase.ScrapeKemenagAyat(ctx, listSurat)
	if err != nil {
		log.Fatalln(errorx.PrintTrace(err))
	}
	quranData := ucase.ConvertKemenagToDb(listSurat, mapSuratAyat)
	log.Println(len(mapSuratAyat))
	err = ucase.InsertQuranData(ctx, quranData)
	if err != nil {
		log.Fatalln(errorx.PrintTrace(err))
	}
}
