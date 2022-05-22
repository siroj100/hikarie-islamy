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
	db0 := dbs[config.DbIslamy]
	err := db0.Pri().AutoMigrate(
		db.QuranL10N{}, db.QuranSurat{}, db.QuranSuratL10N{},
		db.QuranAyat{}, db.QuranAyatL10N{},
	)
	log.Println(err)
	ctx := context.TODO()
	var listAyat []db.QuranAyatL10N
	err = db0.GetTxRead(ctx).Model(db.QuranAyatL10N{}).Joins("Ayat").Order("ayat_id").Find(&listAyat).Error
	if err != nil {
		log.Fatalln(errorx.PrintTrace(err))
	}
	//db0Tx := repository.NewGormDbTx(db0)
	//err = db0Tx.Transaction(ctx, func(txCtx context.Context) error {
	//	for i := range listAyat {
	//		ayatL10n := listAyat[i]
	//		ayat := ayatL10n.Ayat
	//		log.Printf("%d: %d - %d, Page %d %d", ayatL10n.AyatID, ayat.SuratID, ayat.AyatNumber, ayatL10n.PageNumber, ayat.PageNumber)
	//		if ayat.PageNumber == 0 {
	//			ayat.PageNumber = ayatL10n.PageNumber
	//			txErr := db0.GetTx(txCtx).Model(ayat).Updates(map[string]interface{}{
	//				"page_number": ayat.PageNumber,
	//			}).Error
	//			if txErr != nil {
	//				log.Println(errorx.PrintTrace(txErr))
	//				return txErr
	//			}
	//		}
	//	}
	//	return nil
	//})
	//if err != nil {
	//	log.Fatalln(errorx.PrintTrace(err))
	//}
}
