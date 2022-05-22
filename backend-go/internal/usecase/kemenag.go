package usecase

import (
	"context"
	"log"
	"strings"
	"time"

	"github.com/siroj100/hikarie-islamy/internal/constant/kemenag"
	"github.com/siroj100/hikarie-islamy/internal/model"
	"github.com/siroj100/hikarie-islamy/internal/model/db"
	"github.com/siroj100/hikarie-islamy/pkg/errorx"
)

func (u IslamyUseCase) KemenagScrapeSurat(ctx context.Context) ([]model.QuranKemenagSurat, error) {
	result, err := u.kemenag.ScrapeListSurat(ctx)
	if err != nil {
		log.Println(errorx.PrintTrace(err))
		return result, err
	}
	log.Println("surat:", len(result))
	return result, nil
}

func (u IslamyUseCase) KemenagScrapeAyat(ctx context.Context, listSurat []model.QuranKemenagSurat) (map[int][]model.QuranKemenagAyat, error) {
	result := make(map[int][]model.QuranKemenagAyat)
	for _, surat := range listSurat {
		time.Sleep(1 * time.Second)
		listAyat, err := u.kemenag.ScrapeListAyat(ctx, surat.ID, surat.CountAyat)
		if err != nil {
			log.Println(errorx.PrintTrace(err))
			continue
		}
		log.Println("surat:", surat.ID, len(listAyat))
		result[surat.ID] = listAyat
	}
	return result, nil
}

func (u IslamyUseCase) KemenagConvertToQuranData(kemenagSurat []model.QuranKemenagSurat, kemenagAyat map[int][]model.QuranKemenagAyat) model.QuranData {
	result := model.QuranData{}
	result.LangData = db.QuranL10N{
		LangCode:   "id",
		SourceID:   "kemenag",
		SourceName: "Quran Kemenag",
		SourceDesc: "Quran Kemenag",
	}
	listSurat := make([]db.QuranSurat, 0, len(kemenagSurat))
	listSuratL10N := make([]db.QuranSuratL10N, 0, len(kemenagSurat))
	listAllAyat := make(map[int][]db.QuranAyat)
	listAllAyatL10N := make(map[int][]db.QuranAyatL10N)
	for _, dataSurat := range kemenagSurat {
		surat := db.QuranSurat{
			SuratID:   dataSurat.ID,
			Name:      strings.TrimSpace(dataSurat.SuratText),
			AyatCount: dataSurat.CountAyat,
		}
		suratL10N := db.QuranSuratL10N{
			SuratID:   dataSurat.ID,
			Translit:  dataSurat.SuratName,
			Translate: dataSurat.SuratTerjemahan,
		}
		listAyat := make([]db.QuranAyat, 0, len(kemenagAyat[dataSurat.ID]))
		listAyatL10N := make([]db.QuranAyatL10N, 0, len(kemenagAyat[dataSurat.ID]))
		for _, dataAyat := range kemenagAyat[dataSurat.ID] {
			ayat := db.QuranAyat{
				SuratID:    surat.SuratID,
				AyatNumber: dataAyat.AyaNumber,
				AyatText:   dataAyat.AyaText,
				JuzID:      dataAyat.JuzID,
				PageNumber: dataAyat.PageNumber,
			}
			ayatL10N := db.QuranAyatL10N{
				Translate:  dataAyat.TranslationAyaText,
			}
			listAyat = append(listAyat, ayat)
			listAyatL10N = append(listAyatL10N, ayatL10N)
		}
		listAllAyat[surat.SuratID] = listAyat
		listAllAyatL10N[surat.SuratID] = listAyatL10N
		listSurat = append(listSurat, surat)
		listSuratL10N = append(listSuratL10N, suratL10N)
	}
	result.ListSurat = listSurat
	result.ListSuratL10N = listSuratL10N
	result.ListAyat = listAllAyat
	result.ListAyatL10N = listAllAyatL10N
	return result
}

func (u IslamyUseCase) KemenagV1ListSurat(ctx context.Context, startSurat, count int) ([]model.QuranKemenagSurat, error) {
	var result []model.QuranKemenagSurat
	listSurat, err := u.quran.ListSuratL10N(ctx, kemenag.LangID, startSurat, count)
	if err != nil {
		log.Println(errorx.PrintTrace(err))
		return result, errorx.ErrServerError
	}
	result = u.kemenag.QuranSuratL10NToKemenag(listSurat)
	return result, nil
}

func (u IslamyUseCase) KemenagV1ListAyat(ctx context.Context, suratID, startAyat, count int) ([]model.QuranKemenagAyat, error) {
	var result []model.QuranKemenagAyat
	listAyat, err := u.quran.ListAyatL10N(ctx, kemenag.LangID, suratID, startAyat, count)
	if err != nil {
		log.Println(errorx.PrintTrace(err))
		return result, errorx.ErrServerError
	}
	result = u.kemenag.QuranAyatL10NToKemenag(listAyat)
	return result, nil
}
