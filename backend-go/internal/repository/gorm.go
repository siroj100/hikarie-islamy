package repository

import (
	"context"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/siroj100/hikarie-islamy/internal/config"
	"github.com/siroj100/hikarie-islamy/internal/ctxs"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

const (
	MyTestDbDriver = "mytxdb"
)

type (
	GormDb interface {
		Pri() *gorm.DB
		Sec() *gorm.DB
		GetTx(ctx context.Context) *gorm.DB
		GetTxRead(ctx context.Context) *gorm.DB
	}

	gormDb struct {
		primary, secondary *gorm.DB
	}
)

func (r *gormDb) GetTxRead(ctx context.Context) *gorm.DB {
	if ctxTx, ok := ctx.Value(ctxs.Transaction).(*gorm.DB); ok {
		//log.Println("tx:", ctxTx)
		return ctxTx
	} else {
		//log.Println("tx not found:", ctxTx)
		return r.Sec()
	}
}

func (r *gormDb) GetTx(ctx context.Context) *gorm.DB {
	if ctxTx, ok := ctx.Value(ctxs.Transaction).(*gorm.DB); ok {
		//log.Println("tx:", ctxTx)
		return ctxTx
	} else {
		//log.Println("tx not found:", ctxTx)
		return r.Pri()
	}
}

func (r *gormDb) Pri() *gorm.DB {
	return r.primary
}

func (r *gormDb) Sec() *gorm.DB {
	return r.secondary
}

func NewGormDb(cfg config.DatabaseConfig) (GormDb, error) {
	var (
		primary, secondary *gorm.DB
		err                error
	)
	host := cfg.Host
	port := cfg.Port
	username := cfg.User
	password := cfg.Password
	dbName := cfg.DbName
	schema := cfg.Schema
	maxOpenConns := cfg.MaxOpenConns
	maxIdleConns := cfg.MaxIdleConns
	primary, err = GormOpenDb(cfg.Driver, host, port, username, password, dbName, schema, cfg.Debug, maxOpenConns, maxIdleConns)
	if err != nil {
		return nil, err
	}

	host = cfg.SecHost
	if len(host) < 1 {
		host = cfg.Host
	}
	port = cfg.SecPort
	if port < 1 {
		port = cfg.Port
	}
	username = cfg.SecUser
	if len(username) < 1 {
		username = cfg.User
	}
	password = cfg.SecPassword
	if len(password) < 1 {
		password = cfg.Password
	}
	dbName = cfg.SecDbname
	if len(dbName) < 1 {
		dbName = cfg.DbName
	}
	schema = cfg.SecSchema
	if len(schema) < 1 {
		schema = cfg.Schema
	}
	maxOpenConns = cfg.SecMaxOpenConns
	if maxOpenConns == 0 {
		maxOpenConns = cfg.MaxOpenConns
	}
	maxIdleConns = cfg.SecMaxIdleConns
	if maxIdleConns == 0 {
		maxIdleConns = cfg.MaxIdleConns
	}
	if host == cfg.Host &&
		port == cfg.Port &&
		username == cfg.User &&
		password == cfg.Password &&
		dbName == cfg.DbName {
		secondary = primary
	} else {
		secondary, err = GormOpenDb(cfg.Driver, host, port, username, password, dbName, schema, cfg.Debug, maxOpenConns, maxIdleConns)
		if err != nil {
			return nil, err
		}
	}

	return &gormDb{
		primary:   primary,
		secondary: secondary,
	}, nil
}

func GormOpenDb(driver, host string, port int, username, password, dbName, dbSchema string, debug bool, maxOpenConns, maxIdleConns int) (*gorm.DB, error) {
	logCfg := logger.Config{
		SlowThreshold: 100 * time.Millisecond,
		LogLevel:      logger.Silent,
	}
	if debug {
		logCfg.LogLevel = logger.Info
	}
	gormLog := logger.New(log.New(os.Stdout, "\r\n", log.LstdFlags), logCfg)
	gormCfg := gorm.Config{
		Logger: gormLog,
	}
	switch strings.ToLower(driver) {
	case "postgres":
		if port == 0 {
			port = 5432
		}
		if len(dbSchema) == 0 {
			dbSchema = "public"
		}
		dsn := fmt.Sprintf("user=%s password=%s host=%s port=%d dbname=%s search_path=%s", username, password, host, port, dbName, dbSchema)
		//log.Println("dsn:", dsn)
		theDb, err := gorm.Open(postgres.New(postgres.Config{
			DSN: dsn,
		}), &gormCfg)
		if err != nil {
			return nil, fmt.Errorf("error creating db connection to %s@%s/%s, %v\n", username, host, dbName, err)
		}
		db, _ := theDb.DB()
		if maxOpenConns > 0 {
			db.SetMaxOpenConns(maxOpenConns)
		}
		if maxIdleConns > 0 {
			db.SetMaxIdleConns(maxIdleConns)
		}
		return theDb, nil

	case MyTestDbDriver:
		dsn := dbName
		theDb, err := gorm.Open(postgres.New(postgres.Config{DriverName: MyTestDbDriver, DSN: dsn}), &gormCfg)
		if err != nil {
			return nil, fmt.Errorf("error creating db connection using %s, %v\n", MyTestDbDriver, err)
		}
		db, _ := theDb.DB()
		if maxOpenConns > 0 {
			db.SetMaxOpenConns(maxOpenConns)
		}
		if maxIdleConns > 0 {
			db.SetMaxIdleConns(maxIdleConns)
		}
		return theDb, nil
	}
	return nil, fmt.Errorf("unknown db driver")
}
