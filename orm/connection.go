package orm

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/ethereum/go-ethereum/log"

	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/utils"
)

type gormLogger struct {
	gethLogger log.Logger
}

func (g *gormLogger) LogMode(level logger.LogLevel) logger.Interface {
	return g
}

func (g *gormLogger) Info(_ context.Context, msg string, data ...interface{}) {
	infoMsg := fmt.Sprintf(msg, data...)
	g.gethLogger.Info("gorm", "info message", infoMsg)
}

func (g *gormLogger) Warn(_ context.Context, msg string, data ...interface{}) {
	warnMsg := fmt.Sprintf(msg, data...)
	g.gethLogger.Warn("gorm", "warn message", warnMsg)
}

func (g *gormLogger) Error(_ context.Context, msg string, data ...interface{}) {
	errMsg := fmt.Sprintf(msg, data...)
	g.gethLogger.Error("gorm", "err message", errMsg)
}

func (g *gormLogger) Trace(_ context.Context, begin time.Time, fc func() (string, int64), err error) {
	elapsed := time.Since(begin)
	sql, rowsAffected := fc()
	g.gethLogger.Debug("gorm", "line", utils.FileWithLineNum(), "cost", elapsed, "sql", sql, "rowsAffected", rowsAffected, "err", err)
}

// InitDB init the db handler
func InitDB() (*gorm.DB, error) {
	viper.SetConfigFile("config.toml")
	if err := viper.ReadInConfig(); err != nil {
		fmt.Printf("Error reading config file, %s", err)
		return nil, err
	}

	dsn := viper.GetString("bridge_db_config.dsn")
	maxOpenNum := viper.GetInt("bridge_db_config.maxOpenNum")
	maxIdleNum := viper.GetInt("bridge_db_config.maxIdleNum")

	tmpGormLogger := gormLogger{
		gethLogger: log.Root(),
	}
	var dialector gorm.Dialector

	dialector = mysql.Open(dsn)

	db, err := gorm.Open(dialector, &gorm.Config{
		Logger: &tmpGormLogger,
		NowFunc: func() time.Time {
			return NowUTC()
		},
	})
	if err != nil {
		return nil, err
	}

	sqlDB, pingErr := Ping(db)
	if pingErr != nil {
		return nil, pingErr
	}

	sqlDB.SetConnMaxLifetime(time.Minute * 10)
	sqlDB.SetConnMaxIdleTime(time.Minute * 5)

	sqlDB.SetMaxOpenConns(maxOpenNum)
	sqlDB.SetMaxIdleConns(maxIdleNum)

	return db, nil
}

// CloseDB close the db handler. notice the db handler only can close when then program exit.
func CloseDB(db *gorm.DB) error {
	sqlDB, err := db.DB()
	if err != nil {
		return err
	}
	if err := sqlDB.Close(); err != nil {
		return err
	}
	return nil
}

// Ping check db status
func Ping(db *gorm.DB) (*sql.DB, error) {
	sqlDB, err := db.DB()
	if err != nil {
		return nil, err
	}

	if err = sqlDB.Ping(); err != nil {
		return nil, err
	}
	return sqlDB, nil
}

func NowUTC() time.Time {
	utc, _ := time.LoadLocation("")
	return time.Now().In(utc)
}
