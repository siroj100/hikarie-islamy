package main

import (
	"context"
	"log"
	"net/http"

	"github.com/siroj100/hikarie-islamy/internal/repository"
	"github.com/siroj100/hikarie-islamy/internal/service"
	"github.com/siroj100/hikarie-islamy/internal/usecase"
	"github.com/siroj100/hikarie-islamy/pkg/errorx"
)

func main() {
	repo := repository.NewKemenagRepo(http.Client{})
	svc := service.NewKemenag(repo)
	ucase := usecase.NewIslamyUseCase(usecase.IslamySvc{Kemenag: svc})

	ctx := context.TODO()
	listSurat, err := ucase.ScrapeKemenagSurat(ctx)
	if err != nil {
		log.Fatalln(errorx.PrintTrace(err))
	}
	mapSuratAyat, err := ucase.ScrapeKemenagAyat(ctx, listSurat)
	if err != nil {
		log.Fatalln(errorx.PrintTrace(err))
	}
	log.Println(len(mapSuratAyat))
}
