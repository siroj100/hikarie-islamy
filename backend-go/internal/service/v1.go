package service

import (
	"github.com/siroj100/hikarie-islamy/internal/model/db"
	v1 "github.com/siroj100/hikarie-islamy/internal/model/v1"
)

type (
	V1QuranService struct {
		repo QuranRepo
	}
)

func NewV1Quran(repo QuranRepo) V1QuranService {
	return V1QuranService{repo: repo}
}

func (s V1QuranService) BuildListSura(listSura []db.QuranSuratL10N, listFirstAya []db.QuranAyat) ([]v1.QuranSura, error) {
	result := make([]v1.QuranSura, len(listSura))
	for i := range listSura {
		sura := listSura[i]
		aya := listFirstAya[i]
		result[i] = v1.QuranSura{
			ID:          sura.SuratID,
			Name:        sura.Surat.Name,
			AyaCount:    sura.Surat.AyatCount,
			PageStart:   aya.PageNumber,
			Translation: sura.Translate,
		}
	}
	return result, nil
}