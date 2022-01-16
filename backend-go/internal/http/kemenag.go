package http

import (
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
	
	"github.com/siroj100/hikarie-islamy/internal/constant/quran"
	"github.com/siroj100/hikarie-islamy/internal/model"
	"github.com/siroj100/hikarie-islamy/pkg/errorx"
)

func (h EchoHttpHandler) KemenagV1Surat(c echo.Context) error {
	var (
		resp struct {
			Msg  string                    `json:"msg"`
			Data []model.QuranKemenagSurat `json:"data"`
		}
		req model.QuranKemenagV1ListSuratReq
	)
	resp.Msg = "All Surat Data"
	err := c.Bind(&req)
	if err != nil {
		return errorx.ErrInvalidParam
	}
	if req.StartSurat < quran.SuratMin-1 || req.StartSurat > quran.SuratMax-1 {
		resp.Data = make([]model.QuranKemenagSurat, 0)
		return c.JSON(http.StatusOK, resp)
	}
	resp.Data, err = h.ucase.KemenagV1ListSurat(ContextEcho(c), req.StartSurat, req.Count)
	if err != nil {
		log.Println(errorx.PrintTrace(err))
		return err
	}
	return c.JSON(http.StatusOK, resp)
}

func (h EchoHttpHandler) KemenagV1Ayat(c echo.Context) error {
	var (
		resp struct {
			Data []model.QuranKemenagAyat `json:"data"`
		}
		req model.QuranKemenagV1ListAyatReq
	)
	err := c.Bind(&req)
	if err != nil {
		return errorx.ErrInvalidParam
	}
	if req.SuratID < quran.SuratMin || req.SuratID > quran.SuratMax {
		resp.Data = make([]model.QuranKemenagAyat, 0)
		return c.JSON(http.StatusOK, resp)
	}
	resp.Data, err = h.ucase.KemenagV1ListAyat(ContextEcho(c), req.SuratID, req.StartAyat, req.Count)
	if err != nil {
		log.Println(errorx.PrintTrace(err))
		return err
	}
	return c.JSON(http.StatusOK, resp)
}
