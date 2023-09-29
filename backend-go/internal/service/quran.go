package service

import (
	"context"
	"errors"
	"fmt"
	"log"

	"github.com/siroj100/hikarie-islamy/internal/model/db"
	"github.com/siroj100/hikarie-islamy/internal/model/quran"
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
		ListSuratL10N(ctx context.Context, langID, startID, count int) ([]db.QuranSuratL10N, error)
		ListAyatL10N(ctx context.Context, langID, suratID, startID, count int) ([]db.QuranAyatL10N, error)

		ListFirstAyat(ctx context.Context, startID, count int) ([]db.QuranAyat, error)

		GetPageDetailFromFile(ctx context.Context, layoutID, pageNumber int) (quran.V1PageResp, error)
		ListAyatByListSuratAyat(ctx context.Context, surat, ayat []int) ([]db.QuranAyat, error)
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

func (s QuranService) ListSuratL10N(ctx context.Context, langID, startID, count int) ([]db.QuranSuratL10N, error) {
	result, err := s.repo.ListSuratL10N(ctx, langID, startID, count)
	if err != nil {
		log.Println(errorx.PrintTrace(err))
		return result, errorx.ErrServerError
	}
	return result, nil
}

func (s QuranService) ListAyatL10N(ctx context.Context, langID, suratID, startID, count int) ([]db.QuranAyatL10N, error) {
	result, err := s.repo.ListAyatL10N(ctx, langID, suratID, startID, count)
	if err != nil {
		log.Println(errorx.PrintTrace(err))
		return result, errorx.ErrServerError
	}
	return result, nil
}

func (s QuranService) ListFirstAyat(ctx context.Context, startID, count int) ([]db.QuranAyat, error) {
	result, err := s.repo.ListFirstAyat(ctx, startID, count)
	if err != nil {
		log.Println(errorx.PrintTrace(err))
		return result, errorx.ErrServerError
	}
	return result, nil
}

func (s QuranService) GetV1Page(ctx context.Context, layoutID, pageNumber int) (quran.V1PageResp, error) {
	result, err := s.repo.GetPageDetailFromFile(ctx, layoutID, pageNumber)
	if err != nil {
		log.Println(errorx.PrintTrace(err))
		return result, errorx.ErrServerError
	}
	var suratList, ayatList []int
	lastSurat := 0
	lastAyat := 0
	for _, line := range result.ListLine {
		for _, ayat := range line.ListAyat {
			if ayat.SuratID != lastSurat {
				suratList = append(suratList, ayat.SuratID)
				lastSurat = ayat.SuratID
			}
			if ayat.AyatID != lastAyat {
				ayatList = append(ayatList, ayat.AyatID)
				lastAyat = ayat.AyatID
			}
		}
	}
	dbAyatList, err := s.repo.ListAyatByListSuratAyat(ctx, suratList, ayatList)
	if err != nil {
		log.Println(errorx.PrintTrace(err))
		return result, errorx.ErrServerError
	}
	dbAyatMap := make(map[string]db.QuranAyat)
	for i := range dbAyatList {
		key := fmt.Sprintf("%d-%d", dbAyatList[i].SuratID, dbAyatList[i].AyatNumber)
		dbAyatMap[key] = dbAyatList[i]
	}
	charStart := 0
	for i := range result.ListLine {
		for j := range result.ListLine[i].ListAyat {
			ayatResp := &result.ListLine[i].ListAyat[j]
			key := fmt.Sprintf("%d-%d", ayatResp.SuratID, ayatResp.AyatID)
			ayat, found := dbAyatMap[key]
			if !found {
				err = errors.New("Ayat not found: " + key)
				log.Println(errorx.PrintTrace(err))
				continue
			}
			ayatResp.CharStart = charStart
			if ayatResp.TotalChar == 0 {
				ayatResp.TotalChar = len(ayat.AyatText) - ayatResp.CharStart
			}
			charStart += ayatResp.TotalChar
			if charStart >= len(ayat.AyatText) {
				charStart = 0
			}
			log.Printf("#%d %d:%d (%d), %d-%d\n", i+1, ayat.SuratID, ayat.AyatNumber, len(ayat.AyatText), ayatResp.CharStart, ayatResp.CharStart+ayatResp.TotalChar)
			ayatResp.Text = ayat.AyatText[ayatResp.CharStart : ayatResp.CharStart+ayatResp.TotalChar]
			if ayatResp.CharStart+ayatResp.TotalChar >= len(ayat.AyatText) {
				ayatResp.AyatEnd = true
			}
		}
	}
	return result, nil
}
