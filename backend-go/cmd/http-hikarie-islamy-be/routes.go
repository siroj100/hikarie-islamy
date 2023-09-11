package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"github.com/siroj100/hikarie-islamy/internal/config"
	handler "github.com/siroj100/hikarie-islamy/internal/http"
	"github.com/siroj100/hikarie-islamy/internal/usecase"
)

func initRoutes(cfg config.Config, ucase usecase.IslamyUseCase) http.Handler {
	e := echo.New()
	e.Use(middleware.GzipWithConfig(middleware.GzipConfig{Level: 5}))
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: cfg.Server.CORSOrigins,
		AllowHeaders: cfg.Server.CORSHeaders,
	}))

	httpHandler := handler.NewEchoHttpHandler(ucase)

	kemenagEp := e.Group("/api/kemenag")
	{
		kemenagEp.GET("/v1/surat/:startSurat/:count", httpHandler.KemenagV1Surat)
		kemenagEp.GET("/v1/ayatweb/:suratID/0/:startAyat/:count", httpHandler.KemenagV1Ayat)
	}

	v1Ep := e.Group("/api/v1")
	{
		v1Ep.GET("/sura", httpHandler.V1ListSura)
	}

	quranEp := e.Group("/api/quran")
	{
		quranEp.GET("/v1/sura", httpHandler.QuranV1ListSura)
		quranEp.GET("/v1/page/:layoutID/:pageNumber", httpHandler.QuranV1Page)
	}

	return e
}
