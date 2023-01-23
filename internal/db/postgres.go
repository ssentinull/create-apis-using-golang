package db

import (
	"context"
	"errors"
	"fmt"
	"regexp"
	"time"

	"github.com/jpillora/backoff"
	"github.com/sirupsen/logrus"
	"github.com/ssentinull/create-apis-using-golang/internal/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	gormLogger "gorm.io/gorm/logger"
)

var (
	PostgresDB     *gorm.DB
	StopTickerChan chan bool
	sqlRegexp      = regexp.MustCompile(`(\$\d+)|\?`)
)

func InitializePostgresConn() {
	conn, err := openPostgresConn(config.DatabaseDSN())
	if err != nil {
		logrus.WithField("databaseDSN", config.DatabaseDSN()).
			Fatal("failed to connect cockroach database: ", err)
	}

	PostgresDB = conn
	StopTickerChan = make(chan bool)

	go checkConnection(time.NewTicker(config.PostgresPingInterval()))

	switch config.LogLevel() {
	case "error":
		PostgresDB.Logger = PostgresDB.Logger.LogMode(gormLogger.Error)
	case "warn":
		PostgresDB.Logger = PostgresDB.Logger.LogMode(gormLogger.Warn)
	default:
		PostgresDB.Logger = PostgresDB.Logger.LogMode(gormLogger.Info)

	}
	logrus.Info("Connection to Cockroach Server success...")
}

func openPostgresConn(dsn string) (*gorm.DB, error) {
	dialector := postgres.Open(dsn)
	db, err := gorm.Open(dialector, &gorm.Config{})
	if err != nil {
		return nil, err
	}

	conn, err := db.DB()
	if err != nil {
		return nil, err
	}

	conn.SetMaxIdleConns(config.PostgresMaxIdleConns())
	conn.SetMaxOpenConns(config.PostgresMaxOpenConns())
	conn.SetConnMaxLifetime(config.PostgresConnMaxLifetime())

	return db, nil
}

func checkConnection(ticker *time.Ticker) {
	for {
		select {
		case <-StopTickerChan:
			ticker.Stop()
			return
		case <-ticker.C:
			if _, err := PostgresDB.DB(); err != nil {
				reconnectPostgresConn()
			}
		}
	}
}

func reconnectPostgresConn() {
	b := backoff.Backoff{
		Factor: 2,
		Jitter: true,
		Min:    100 * time.Millisecond,
		Max:    1 * time.Second,
	}

	postgresRetryAttempts := float64(config.PostgresRetryAttempts())
	for b.Attempt() < postgresRetryAttempts {
		conn, err := openPostgresConn(config.DatabaseDSN())
		if err != nil {
			logrus.WithField("databaseDSN", config.DatabaseDSN()).
				Fatal("failed to connect cockroach database: ", err)
		}

		if conn != nil {
			PostgresDB = conn
			break
		}

		time.Sleep(b.Duration())
	}

	if b.Attempt() >= postgresRetryAttempts {
		logrus.Fatal("have retried max num of times to connect to db")
	}

	b.Reset()
}

// GormCustomLogger override gorm logger
type GormCustomLogger struct {
	gormLogger.Config
}

// NewGormCustomLogger :nodoc:
func NewGormCustomLogger() *GormCustomLogger {
	return &GormCustomLogger{
		Config: gormLogger.Config{
			LogLevel: gormLogger.Info,
		},
	}
}

// LogMode :nodoc:
func (g *GormCustomLogger) LogMode(level gormLogger.LogLevel) gormLogger.Interface {
	g.LogLevel = level
	return g
}

// Info :nodoc:
func (g *GormCustomLogger) Info(ctx context.Context, message string, values ...interface{}) {
	if g.LogLevel >= gormLogger.Info {
		logrus.WithField("data", values).Error(message)
	}
}

// Warn :nodoc:
func (g *GormCustomLogger) Warn(ctx context.Context, message string, values ...interface{}) {
	if g.LogLevel >= gormLogger.Warn {
		logrus.WithField("data", values).Warn(message)
	}

}

// Error :nodoc:
func (g *GormCustomLogger) Error(ctx context.Context, message string, values ...interface{}) {
	if g.LogLevel >= gormLogger.Error {
		logrus.WithField("data", values).Error(message)
	}
}

// Trace :nodoc:
func (g *GormCustomLogger) Trace(ctx context.Context, begin time.Time, fc func() (string, int64), err error) {
	sql, rows := fc()
	if g.LogLevel <= 0 {
		return
	}

	elapsed := time.Since(begin)
	logger := logrus.WithField("took", elapsed)
	sqlLog := sqlRegexp.ReplaceAllString(sql, "%v")
	if rows >= 0 {
		logger.WithField("rows", rows)
	} else {
		logger.WithField("rows", "-")
	}

	switch {
	case err != nil && !errors.Is(err, gorm.ErrRecordNotFound) && g.LogLevel >= gormLogger.Error:
		logger.WithField("sql", sqlLog).Error(err)
	case elapsed > g.SlowThreshold && g.SlowThreshold != 0 && g.LogLevel >= gormLogger.Warn:
		slowLog := fmt.Sprintf("SLOW SQL >= %v", g.SlowThreshold)
		logger.WithField("sql", sqlLog).Warn(slowLog)
	case g.LogLevel >= gormLogger.Info:
		logger.Info(sqlLog)

	}
}
