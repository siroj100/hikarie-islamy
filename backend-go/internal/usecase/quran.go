package usecase

import (
	"context"
	"log"

	"github.com/siroj100/hikarie-islamy/internal/model"
	"github.com/siroj100/hikarie-islamy/pkg/errorx"
)

func (u IslamyUseCase) QuranInsertQuranData(ctx context.Context, quranData model.QuranData) error {
	quranL10N, err := u.quran.SaveQuranL10N(ctx, quranData.LangData)
	if err != nil {
		log.Println(errorx.PrintTrace(err))
		return err
	}
	for i, surat := range quranData.ListSurat {
		_, err = u.quran.SaveQuranSurat(ctx, surat)
		if err != nil {
			log.Println(errorx.PrintTrace(err))
			return err
		}
		suratL10N := quranData.ListSuratL10N[i]
		suratL10N.LangID = quranL10N.LangID
		_, err = u.quran.SaveQuranSuratL10N(ctx, suratL10N)
		if err != nil {
			log.Println(errorx.PrintTrace(err))
			return err
		}
		for j, ayat := range quranData.ListAyat[surat.SuratID] {
			dbAyat, err := u.quran.SaveQuranAyat(ctx, ayat)
			if err != nil {
				log.Println(errorx.PrintTrace(err))
				return err
			}
			ayatL10N := quranData.ListAyatL10N[surat.SuratID][j]
			ayatL10N.LangID = quranL10N.LangID
			ayatL10N.AyatID = dbAyat.AyatID
			_, err = u.quran.SaveQuranAyatL10N(ctx, ayatL10N)
			if err != nil {
				log.Println(errorx.PrintTrace(err))
				return err
			}
		}
	}
	return nil
}
