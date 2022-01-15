package internal

import (
	"log"

	"github.com/siroj100/hikarie-islamy/internal/config"
	"github.com/siroj100/hikarie-islamy/internal/repository"
)

func InitGormDb(cfg map[string]config.DatabaseConfig) map[string]repository.GormDb {
	result := make(map[string]repository.GormDb)

	for k, v := range cfg {
		db, err := repository.NewGormDb(v)
		if err != nil {
			log.Fatalln(err)
		}
		result[k] = db
	}

	return result
}
