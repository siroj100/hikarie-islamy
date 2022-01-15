package repository

import (
	"context"

	"github.com/siroj100/hikarie-islamy/internal/model/db"
)

type (
	QuranRepo struct {
		db GormDb
	}
)

func NewQuran(db GormDb) QuranRepo {
	return QuranRepo{db: db}
}

func (r QuranRepo) InsertQuranL10N(ctx context.Context, data *db.QuranL10N) error {
	err := r.db.GetTx(ctx).Create(data).Error
	return err
}

func (r QuranRepo) InsertQuranSurat(ctx context.Context, data *db.QuranSurat) error {
	err := r.db.GetTx(ctx).Create(data).Error
	return err
}

func (r QuranRepo) InsertQuranSuratL10N(ctx context.Context, data *db.QuranSuratL10N) error {
	err := r.db.GetTx(ctx).Create(data).Error
	return err
}

func (r QuranRepo) InsertQuranAyat(ctx context.Context, data *db.QuranAyat) error {
	err := r.db.GetTx(ctx).Create(data).Error
	return err
}

func (r QuranRepo) InsertQuranAyatL10N(ctx context.Context, data *db.QuranAyatL10N) error {
	err := r.db.GetTx(ctx).Create(data).Error
	return err
}
