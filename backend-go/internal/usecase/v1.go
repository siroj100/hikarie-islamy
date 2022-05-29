package usecase

import (
	"context"
	"log"

	"github.com/siroj100/hikarie-islamy/internal/constant/kemenag"
	v1 "github.com/siroj100/hikarie-islamy/internal/model/v1"
	"github.com/siroj100/hikarie-islamy/pkg/errorx"
)

func (u IslamyUseCase) V1QuranListSura(ctx context.Context) (v1.ListSuraResp, error) {
	var result v1.ListSuraResp

	listSura, err := u.quran.ListSuratL10N(ctx, kemenag.LangID, 0, 114)
	if err != nil {
		log.Println(errorx.PrintTrace(err))
		return result, errorx.ErrServerError
	}
	listAya, err := u.quran.ListFirstAyat(ctx, 0, 114)
	if err != nil {
		log.Println(errorx.PrintTrace(err))
		return result, errorx.ErrServerError
	}
	result = v1.ListSuraResp{}
	result.ListSura, _ = u.v1Quran.BuildListSura(listSura, listAya)

	return result, nil
}
