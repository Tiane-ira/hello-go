package db

import (
	"fmt"
	"hello-go/configs"
	"hello-go/model"
	"hello-go/zlog"
	"log"
	"os"
	"strings"
	"time"

	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

func NewMysqlDb() *gorm.DB {
	mysqlConf := configs.Get().Mysql

	dns := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?%s", mysqlConf.Username, mysqlConf.Password, mysqlConf.Host, mysqlConf.Port, mysqlConf.Db, mysqlConf.Params)
	mysqlConfig := mysql.Config{
		DSN:                       dns,   // DSN data source name
		DefaultStringSize:         191,   // string 类型字段的默认长度
		SkipInitializeWithVersion: false, // 根据版本自动配置
	}
	db, err := gorm.Open(mysql.New(mysqlConfig), getGormConfig())
	if err != nil {
		zlog.Logger.Error("mysql connect error", zap.Error(err))
		panic(err)
	}
	db.InstanceSet("gorm:table_options", "ENGINE=InnoDB")
	sqlDB, _ := db.DB()
	sqlDB.SetMaxIdleConns(mysqlConf.MaxIdleConns)
	sqlDB.SetMaxOpenConns(mysqlConf.MaxOpenConns)
	zlog.Logger.Info("mysql connected")
	return db
}

func getGormConfig() *gorm.Config {
	mysqlConf := configs.Get().Mysql
	gormConfig := &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   mysqlConf.TablePrefix,
			SingularTable: mysqlConf.SingularTable,
		},
		DisableForeignKeyConstraintWhenMigrating: true,
	}
	_default := logger.New(log.New(os.Stdout, "\r\n", log.LstdFlags), logger.Config{
		SlowThreshold: 200 * time.Millisecond,
		LogLevel:      logger.Warn,
		Colorful:      true,
	})
	switch strings.ToLower(mysqlConf.LogMode) {
	case "silent":
		gormConfig.Logger = _default.LogMode(logger.Silent)
	case "error":
		gormConfig.Logger = _default.LogMode(logger.Error)
	case "warn":
		gormConfig.Logger = _default.LogMode(logger.Warn)
	case "info":
		gormConfig.Logger = _default.LogMode(logger.Info)
	default:
		gormConfig.Logger = _default.LogMode(logger.Info)
	}
	return gormConfig
}

func RegisterTables(db *gorm.DB) {
	err := db.AutoMigrate(
		model.CsUser{},
	)
	if err != nil {
		zlog.Logger.Error("register table failed", zap.Error(err))
		os.Exit(0)
	}
	zlog.Logger.Info("register table success")
}
