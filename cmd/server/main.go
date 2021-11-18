package main

import (
	"net/http"
	"os"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
	"github.com/ssentinull/create-apis-using-golang/config"
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
	initLogger()
}

func main() {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	s := &http.Server{
		Addr:         ":" + config.ServerPort(),
		ReadTimeout:  2 * time.Minute,
		WriteTimeout: 2 * time.Minute,
	}

	logrus.Fatal(e.StartServer(s))
}
