package repository

import (
	"context"
	"log"

	"github.com/siroj100/hikarie-islamy/internal/ctxs"
)

type (
	GormDbTx struct {
		GormDb
	}
)

func (r GormDbTx) Transaction(ctx context.Context, theFunc func(txCtx context.Context) error) error {
	tx := r.Pri().Begin()
	txCtx := context.WithValue(ctx, ctxs.Transaction, tx)
	err := theFunc(txCtx)
	if err != nil {
		log.Println("tx failed, rollback")
		errx := tx.Rollback().Error
		if errx != nil {
			log.Println("rollback failed")
			return errx
		}
		return err
	}
	err = tx.Commit().Error
	if err != nil {
		log.Println("commit failed, rollback")
		errx := tx.Rollback().Error
		if errx != nil {
			log.Println("rollback failed")
			return errx
		}
		return err
	}
	return nil
}

func NewGormDbTx(db GormDb) GormDbTx {
	return GormDbTx{GormDb: db}
}
