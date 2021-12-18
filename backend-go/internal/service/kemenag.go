package service

import (
	"context"
	"log"

	"github.com/siroj100/hikarie-islamy/internal/model"
	"github.com/siroj100/hikarie-islamy/pkg/errorx"
)

type (
	KemenagService struct {
		repo KemenagRepo
	}

	KemenagRepo interface {
		GetListSurat(ctx context.Context) ([]model.QuranKemenagSurat, error)
		GetListAyat(ctx context.Context, suratID, ayatStart, ayatEnd int) ([]model.QuranKemenagAyat, error)
	}
)

func NewKemenag(repo KemenagRepo) KemenagService {
	return KemenagService{repo: repo}
}

func (s KemenagService) ScrapeListSurat(ctx context.Context) ([]model.QuranKemenagSurat, error) {
	result, err := s.repo.GetListSurat(ctx)
	if err != nil {
		log.Println(errorx.PrintTrace(err))
		return result, err
	}
	return result, nil
}

func (s KemenagService) ScrapeListAyat(ctx context.Context, suratID, ayatCount int) ([]model.QuranKemenagAyat, error) {
	result, err := s.repo.GetListAyat(ctx, suratID, 0, ayatCount)
	if err != nil {
		log.Println(errorx.PrintTrace(err))
		return result, err
	}
	return result, nil
}
