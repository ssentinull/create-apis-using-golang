package main

import (
	"net/http"
	"os"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
	"github.com/ssentinull/create-apis-using-golang/pkg/config"
	_bookHTTPHndlr "github.com/ssentinull/create-apis-using-golang/pkg/delivery"
	_bookRepo "github.com/ssentinull/create-apis-using-golang/pkg/repository"
	_bookUcase "github.com/ssentinull/create-apis-using-golang/pkg/usecase"
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

// run initLogger() before running main()
func init() {
	config.GetConf()
	initLogger()
}

func main() {
	e := echo.New()

	bookRepo := _bookRepo.NewBookRepository()
	bookUsecase := _bookUcase.NewBookUsecase(bookRepo)
	_bookHTTPHndlr.NewBookHTTPHandler(e, bookUsecase)

	logrus.Info("di sini!!!")
	logrus.Info(config.ServerPort())
	s := &http.Server{
		Addr:         ":" + config.ServerPort(),
		ReadTimeout:  2 * time.Minute,
		WriteTimeout: 2 * time.Minute,
	}

	logrus.Fatal(e.StartServer(s))
}
