package usecase

import (
	"context"
	"log"
	"time"

	"github.com/siroj100/hikarie-islamy/internal/model"
	"github.com/siroj100/hikarie-islamy/pkg/errorx"
)

type (
	IslamyUseCase struct {
		kemenag KemenagSvc
	}

	IslamySvc struct {
		Kemenag KemenagSvc
	}

	KemenagSvc interface {
		ScrapeListSurat(ctx context.Context) ([]model.QuranKemenagSurat, error)
		ScrapeListAyat(ctx context.Context, suratID, ayatCount int) ([]model.QuranKemenagAyat, error)
	}
)

func NewIslamyUseCase(svc IslamySvc) IslamyUseCase {
	return IslamyUseCase{
		kemenag: svc.Kemenag,
	}
}

func (u IslamyUseCase) ScrapeKemenagSurat(ctx context.Context) ([]model.QuranKemenagSurat, error) {
	result, err := u.kemenag.ScrapeListSurat(ctx)
	if err != nil {
		log.Println(errorx.PrintTrace(err))
		return result, err
	}
	log.Println("surat:", len(result))
	return result, nil
}

func (u IslamyUseCase) ScrapeKemenagAyat(ctx context.Context, listSurat []model.QuranKemenagSurat) (map[int][]model.QuranKemenagAyat, error) {
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
