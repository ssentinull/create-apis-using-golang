package main

import (
	"context"
	"flag"
	"os"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/sirupsen/logrus"
	"github.com/ssentinull/create-apis-using-golang/internal/config"
	"github.com/ssentinull/create-apis-using-golang/internal/db"
	"github.com/ssentinull/create-apis-using-golang/internal/model"
	"github.com/ssentinull/create-apis-using-golang/internal/repository"
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
	seed := flag.Int("seed", 1, "number of seed")
	flag.Parse()

	db.InitializePostgresConn()
	db.InitializeRedisConn()

	cacheRepo := repository.NewCacheRepository(db.RedisClient)
	bookRepo := repository.NewBookRepository(db.PostgresDB, cacheRepo)

	logrus.Infof("Running %d seeds!", *seed)

	for i := 0; i < *seed; i++ {
		book := &model.Book{
			ID:            utils.GenerateID(),
			Title:         gofakeit.BookTitle(),
			Author:        gofakeit.BookAuthor(),
			Description:   gofakeit.Paragraph(1, 3, 10, "."),
			PublishedDate: gofakeit.Date().Format("2006-01-02"),
		}

		if err := bookRepo.Create(context.TODO(), book); err != nil {
			logrus.WithField("book", utils.Dump(book)).Error(err)
		}
	}

	logrus.Info("Finished running seeder!")
}
