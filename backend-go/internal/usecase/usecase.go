package usecase

import (
	"context"

	"github.com/siroj100/hikarie-islamy/internal/model"
	"github.com/siroj100/hikarie-islamy/internal/model/db"
)

type (
	IslamyUseCase struct {
		kemenag KemenagSvc
		quran   QuranSvc
	}

	IslamySvc struct {
		Kemenag KemenagSvc
		Quran   QuranSvc
	}

	KemenagSvc interface {
		ScrapeListSurat(ctx context.Context) ([]model.QuranKemenagSurat, error)
		ScrapeListAyat(ctx context.Context, suratID, ayatCount int) ([]model.QuranKemenagAyat, error)
		QuranSuratL10NToKemenag(listSurat []db.QuranSuratL10N) []model.QuranKemenagSurat
		QuranAyatL10NToKemenag(listAyat []db.QuranAyatL10N) []model.QuranKemenagAyat
	}

	QuranSvc interface {
		SaveQuranL10N(ctx context.Context, data db.QuranL10N) (db.QuranL10N, error)
		SaveQuranSurat(ctx context.Context, data db.QuranSurat) (db.QuranSurat, error)
		SaveQuranSuratL10N(ctx context.Context, data db.QuranSuratL10N) (db.QuranSuratL10N, error)
		SaveQuranAyat(ctx context.Context, data db.QuranAyat) (db.QuranAyat, error)
		SaveQuranAyatL10N(ctx context.Context, data db.QuranAyatL10N) (db.QuranAyatL10N, error)
		ListSuratL10N(ctx context.Context, langID, startID, count int) ([]db.QuranSuratL10N, error)
		ListAyatL10N(ctx context.Context, langID, suratID, startID, count int) ([]db.QuranAyatL10N, error)
	}
)

func NewIslamyUseCase(svc IslamySvc) IslamyUseCase {
	return IslamyUseCase{
		kemenag: svc.Kemenag,
		quran:   svc.Quran,
	}
}
