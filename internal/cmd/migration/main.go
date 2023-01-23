package main

import (
	"flag"
	"os"

	migrate "github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/sirupsen/logrus"
	"github.com/ssentinull/create-apis-using-golang/internal/config"
	"github.com/ssentinull/create-apis-using-golang/internal/db"
	"github.com/ssentinull/create-apis-using-golang/internal/utils"
)

// initialize logger configurations
func initLogger() {
	logLevel := logrus.ErrorLevel
	switch config.Env() {
	case "dev", "development":
		logLevel = logrus.InfoLevel
	}

	logrus.SetFormatter(&logrus.TextFormatter{
		ForceColors:     true,
		DisableSorting:  true,
		DisableColors:   false,
		FullTimestamp:   true,
		TimestampFormat: "15:04:05 02-01-2006",
	})

	logrus.SetOutput(os.Stdout)
	logrus.SetReportCaller(true)
	logrus.SetLevel(logLevel)
}

func init() {
	config.GetConf()
	initLogger()
}

func main() {
	direction := flag.String("direction", "up", "migration direction")
	step := flag.Int("step", 0, "migration step")

	flag.Parse()
	db.InitializePostgresConn()

	sqlDB, err := db.PostgresDB.DB()
	if err != nil {
		logrus.WithField("DatabaseDSN", config.DatabaseDSN()).Fatal("Failed to connect database: ", err)
	}

	driver, err := postgres.WithInstance(sqlDB, &postgres.Config{})
	if err != nil {
		logrus.WithField("sqlDB", utils.Dump(sqlDB)).Fatal("Failed to create driver: ", err)
	}

	migrations, err := migrate.NewWithDatabaseInstance("file://db/migration", "postgres", driver)
	if err != nil {
		logrus.WithField("driver", utils.Dump(driver)).Fatal("Failed to create migration instance: ", err)
	}

	migrations.Steps(*step)
	switch *direction {
	case "up":
		err = migrations.Up()
	case "down":
		err = migrations.Down()
	default:
		logrus.WithField("direction", *direction).Error("invalid direction: ", *direction)
		return
	}

	if err != nil {
		logrus.WithFields(logrus.Fields{
			"migrations": utils.Dump(migrations),
			"direction":  direction,
		}).Fatal("Failed to migrate database: ", err)
	}

	logrus.Infof("Applied migrations from step %d!\n", *step)
}
