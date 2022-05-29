package http

import (
	"log"
	"net/http"

	"github.com/labstack/echo/v4"

	"github.com/siroj100/hikarie-islamy/pkg/errorx"
)

func (h EchoHttpHandler) V1ListSura(c echo.Context) error {
	resp, err := h.ucase.V1QuranListSura(ContextEcho(c))
	if err != nil {
		log.Println(errorx.PrintTrace(err))
		return err
	}
	return c.JSON(http.StatusOK, resp)
}
