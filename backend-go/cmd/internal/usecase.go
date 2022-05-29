package internal

import (
	"net/http"

	"github.com/siroj100/hikarie-islamy/internal/config"
	"github.com/siroj100/hikarie-islamy/internal/repository"
	"github.com/siroj100/hikarie-islamy/internal/service"
	"github.com/siroj100/hikarie-islamy/internal/usecase"
)

func InitUseCase(cfg config.Config, db map[string]repository.GormDb) usecase.IslamyUseCase {
	kemenagRepo := repository.NewKemenag(http.Client{})
	kemenagSvc := service.NewKemenag(kemenagRepo)

	quranRepo := repository.NewQuran(db[config.DbIslamy])
	quranSvc := service.NewQuran(quranRepo)

	v1QuranSvc := service.NewV1Quran(quranRepo)

	return usecase.NewIslamyUseCase(usecase.IslamySvc{
		Kemenag: kemenagSvc,
		Quran:   quranSvc,
		V1Quran: v1QuranSvc,
	})
}
