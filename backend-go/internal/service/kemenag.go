package service

import (
	"context"
	"log"

	"github.com/siroj100/hikarie-islamy/internal/model"
	"github.com/siroj100/hikarie-islamy/internal/model/db"
	"github.com/siroj100/hikarie-islamy/pkg/errorx"
)

type (
	KemenagService struct {
		repo KemenagRepo
	}

	KemenagRepo interface {
		GetListKemenagSurat(ctx context.Context) ([]model.QuranKemenagSurat, error)
		GetListKemenagAyat(ctx context.Context, suratID, ayatStart, ayatEnd int) ([]model.QuranKemenagAyat, error)
	}
)

func NewKemenag(repo KemenagRepo) KemenagService {
	return KemenagService{repo: repo}
}

func (s KemenagService) ScrapeListSurat(ctx context.Context) ([]model.QuranKemenagSurat, error) {
	result, err := s.repo.GetListKemenagSurat(ctx)
	if err != nil {
		log.Println(errorx.PrintTrace(err))
		return result, err
	}
	return result, nil
}

func (s KemenagService) ScrapeListAyat(ctx context.Context, suratID, ayatCount int) ([]model.QuranKemenagAyat, error) {
	result, err := s.repo.GetListKemenagAyat(ctx, suratID, 0, ayatCount)
	if err != nil {
		log.Println(errorx.PrintTrace(err))
		return result, err
	}
	return result, nil
}

func (s KemenagService) QuranSuratL10NToKemenag(listSurat []db.QuranSuratL10N) []model.QuranKemenagSurat {
	result := make([]model.QuranKemenagSurat, 0, len(listSurat))
	for _, data := range listSurat {
		result = append(result, model.QuranKemenagSurat{
			ID:              data.SuratID,
			SuratName:       data.Translit,
			SuratText:       data.Surat.Name,
			SuratTerjemahan: data.Translate,
			CountAyat:       data.Surat.AyatCount,
		})
	}
	return result
}

func (s KemenagService) QuranAyatL10NToKemenag(listAyat []db.QuranAyatL10N) []model.QuranKemenagAyat {
	result := make([]model.QuranKemenagAyat, 0, len(listAyat))
	for _, data := range listAyat {
		result = append(result, model.QuranKemenagAyat{
			AyaID:              data.AyatID,
			AyaNumber:          data.Ayat.AyatNumber,
			AyaText:            data.Ayat.AyatText,
			SuraID:             data.Ayat.SuratID,
			JuzID:              data.Ayat.JuzID,
			PageNumber:         data.PageNumber,
			TranslationAyaText: data.Translate,
		})
	}
	return result
}
