package service

import (
	"context"
	"log"

	"github.com/siroj100/hikarie-islamy/internal/model/db"
	"github.com/siroj100/hikarie-islamy/pkg/errorx"
)

type (
	QuranService struct {
		repo QuranRepo
	}

	QuranRepo interface {
		InsertQuranL10N(ctx context.Context, data *db.QuranL10N) error
		InsertQuranSurat(ctx context.Context, data *db.QuranSurat) error
		InsertQuranSuratL10N(ctx context.Context, data *db.QuranSuratL10N) error
		InsertQuranAyat(ctx context.Context, data *db.QuranAyat) error
		InsertQuranAyatL10N(ctx context.Context, data *db.QuranAyatL10N) error
	}
)

func NewQuran(repo QuranRepo) QuranService {
	return QuranService{repo: repo}
}

func (s QuranService) SaveQuranL10N(ctx context.Context, data db.QuranL10N) (db.QuranL10N, error) {
	var result db.QuranL10N
	err := s.repo.InsertQuranL10N(ctx, &data)
	if err != nil {
		log.Println(errorx.PrintTrace(err))
		return result, errorx.ErrServerError
	}
	result = data
	return result, nil
}

func (s QuranService) SaveQuranSurat(ctx context.Context, data db.QuranSurat) (db.QuranSurat, error) {
	var result db.QuranSurat
	err := s.repo.InsertQuranSurat(ctx, &data)
	if err != nil {
		log.Println(errorx.PrintTrace(err))
		return result, errorx.ErrServerError
	}
	result = data
	return result, nil
}

func (s QuranService) SaveQuranSuratL10N(ctx context.Context, data db.QuranSuratL10N) (db.QuranSuratL10N, error) {
	var result db.QuranSuratL10N
	err := s.repo.InsertQuranSuratL10N(ctx, &data)
	if err != nil {
		log.Println(errorx.PrintTrace(err))
		return result, errorx.ErrServerError
	}
	result = data
	return result, nil
}

func (s QuranService) SaveQuranAyat(ctx context.Context, data db.QuranAyat) (db.QuranAyat, error) {
	var result db.QuranAyat
	err := s.repo.InsertQuranAyat(ctx, &data)
	if err != nil {
		log.Println(errorx.PrintTrace(err))
		return result, errorx.ErrServerError
	}
	result = data
	return result, nil
}

func (s QuranService) SaveQuranAyatL10N(ctx context.Context, data db.QuranAyatL10N) (db.QuranAyatL10N, error) {
	var result db.QuranAyatL10N
	err := s.repo.InsertQuranAyatL10N(ctx, &data)
	if err != nil {
		log.Println(errorx.PrintTrace(err))
		return result, errorx.ErrServerError
	}
	result = data
	return result, nil
}
