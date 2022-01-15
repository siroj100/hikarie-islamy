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
	}

	QuranSvc interface {
		SaveQuranL10N(ctx context.Context, data db.QuranL10N) (db.QuranL10N, error)
		SaveQuranSurat(ctx context.Context, data db.QuranSurat) (db.QuranSurat, error)
		SaveQuranSuratL10N(ctx context.Context, data db.QuranSuratL10N) (db.QuranSuratL10N, error)
		SaveQuranAyat(ctx context.Context, data db.QuranAyat) (db.QuranAyat, error)
		SaveQuranAyatL10N(ctx context.Context, data db.QuranAyatL10N) (db.QuranAyatL10N, error)
	}
)

func NewIslamyUseCase(svc IslamySvc) IslamyUseCase {
	return IslamyUseCase{
		kemenag: svc.Kemenag,
		quran:   svc.Quran,
	}
}
