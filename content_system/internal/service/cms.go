package service

import (
	"content_system/internal/process"
	"context"

	"github.com/redis/go-redis/v9"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	goflow "github.com/s8sg/goflow/v1"
)

type CmsApp struct {
	db          *gorm.DB
	rdb         *redis.Client
	flowService *goflow.FlowService
}

func NewCmsApp() *CmsApp {
	app := &CmsApp{}
	connDB(app)
	connRdb(app)
	flowService(app)
	go func() {
		process.ExceContentFlow(app.db)
	}()
	return app
}

func connDB(app *CmsApp) {
	mysqlDB, err := gorm.Open(mysql.Open("user:password@tcp..."))
	if err != nil {
		panic(err)
	}
	db, err := mysqlDB.DB()
	if err != nil {
		panic(err)
	}
	db.SetMaxOpenConns(4)
	db.SetMaxIdleConns(2)
	mysqlDB = mysqlDB.Debug()

	app.db = mysqlDB
}

func flowService(app *CmsApp) {
	app.flowService = &goflow.FlowService{
		RedisURL: "localhost:6379",
	}
}

func connRdb(app *CmsApp) {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})
	_, err := rdb.Ping(context.Background()).Result()
	if err != nil {
		panic(err)
	}
	app.rdb = rdb
}
