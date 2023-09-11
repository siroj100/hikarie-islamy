package repository

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/siroj100/hikarie-islamy/internal/model/db"
	"github.com/siroj100/hikarie-islamy/internal/model/quran"
	"github.com/siroj100/hikarie-islamy/pkg/errorx"

	"gopkg.in/yaml.v2"
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

func (r QuranRepo) ListSuratL10N(ctx context.Context, langID, startID, count int) ([]db.QuranSuratL10N, error) {
	var result []db.QuranSuratL10N
	err := r.db.GetTxRead(ctx).Model(db.QuranSuratL10N{}).
		Joins("Surat").
		Where("lang_id = ? AND quran_surat_l10n.surat_id > ?", langID, startID).
		Limit(count).
		Find(&result).Error
	return result, err
}

func (r QuranRepo) ListAyatL10N(ctx context.Context, langID, suratID, startID, count int) ([]db.QuranAyatL10N, error) {
	var result []db.QuranAyatL10N
	err := r.db.GetTxRead(ctx).Model(db.QuranAyatL10N{}).
		Joins("Ayat").
		Where("lang_id = ? AND \"Ayat\".surat_id = ? AND \"Ayat\".ayat_number > ?", langID, suratID, startID).
		Limit(count).
		Find(&result).Error
	return result, err
}

func (r QuranRepo) ListFirstAyat(ctx context.Context, startID, count int) ([]db.QuranAyat, error) {
	var result []db.QuranAyat
	err := r.db.GetTxRead(ctx).Model(db.QuranAyat{}).
		Where("ayat_number = 1 AND surat_id > ?", startID).
		Limit(count).
		Find(&result).Error
	return result, err
}

func (r QuranRepo) GetPageDetailFromFile(ctx context.Context, layoutID, pageNumber int) (quran.V1PageResp, error) {
	var result quran.V1PageResp
	raw, err := os.ReadFile(fmt.Sprintf("static/quran/v1Page%02d%03d.yaml", layoutID, pageNumber))
	if err != nil {
		log.Println(errorx.PrintTrace(err))
		return result, err
	}
	if err = yaml.Unmarshal(raw, &result); err != nil {
		log.Println(errorx.PrintTrace(err))
		return result, err
	}
	return result, nil
}

func (r QuranRepo) ListAyatByListSuratAyat(ctx context.Context, surat, ayat []int) ([]db.QuranAyat, error) {
	var result []db.QuranAyat
	err := r.db.GetTxRead(ctx).Model(db.QuranAyat{}).
		Where("surat_id IN (?) AND ayat_number IN (?)", surat, ayat).
		Find(&result).Error
	return result, err
}
