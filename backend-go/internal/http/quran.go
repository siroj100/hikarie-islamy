package http

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/siroj100/hikarie-islamy/internal/model/quran"
	"github.com/siroj100/hikarie-islamy/pkg/errorx"

	"github.com/labstack/echo/v4"
	"gopkg.in/yaml.v2"
)

func (h EchoHttpHandler) QuranV1ListSura(c echo.Context) error {
	var resp quran.V1ListSuratResp
	raw, err := os.ReadFile("static/quran/v1ListSurat.yaml")
	if err != nil {
		log.Println(errorx.PrintTrace(err))
		return err
	}
	if err = yaml.Unmarshal(raw, &resp); err != nil {
		log.Println(errorx.PrintTrace(err))
		return err
	}
	fmt.Printf("resp: %s, %+v\n", raw, resp)
	return c.JSON(http.StatusOK, resp)
}

func (h EchoHttpHandler) QuranV1Page(c echo.Context) error {
	var (
		req  quran.V1PageReq
		resp quran.V1PageResp
	)
	err := c.Bind(&req)
	if err != nil {
		return errorx.ErrInvalidParam
	}

	//raw, err := os.ReadFile(fmt.Sprintf("static/quran/v1Page%02d%03d.yaml", req.LayoutID, req.PageNumber))
	//if err != nil {
	//	log.Println(errorx.PrintTrace(err))
	//	return err
	//}
	//if err = yaml.Unmarshal(raw, &resp); err != nil {
	//	log.Println(errorx.PrintTrace(err))
	//	return err
	//}
	//fmt.Printf("resp: %s, %+v\n", raw, resp)
	resp, err = h.ucase.QuranV1Page(ContextEcho(c), req.LayoutID, req.PageNumber)
	if err != nil {
		log.Println(errorx.PrintTrace(err))
		return err
	}
	return c.JSON(http.StatusOK, resp)
}
