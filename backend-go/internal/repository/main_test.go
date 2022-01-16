package repository

import (
	"log"
	"os"
	"testing"

	"github.com/siroj100/hikarie-islamy/internal/config"
	"github.com/siroj100/hikarie-islamy/pkg/errorx"
)

var (
	theDb GormDb
)

func TestMain(m *testing.M) {
	var err error
	cfg := config.Init()
	theDb, err = NewGormDb(cfg.Database[config.DbIslamy])
	if err != nil {
		log.Fatalln(errorx.PrintTrace(err))
	}

	os.Exit(m.Run())
}
