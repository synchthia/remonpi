package main

import (
	"os"

	"github.com/sirupsen/logrus"
	"github.com/synchthia/remonpi/logger"
	"github.com/synchthia/remonpi/remote"
	"github.com/synchthia/remonpi/server"
)

func startHTTPServer(port string, remote *remote.Remote) error {
	return server.NewHTTPServer(remote).Run(":" + port)
}

func main() {
	logger.Init()
	logrus.Infof("[RemonPi] Initialize...")

	// Remote
	vendor := os.Getenv("REMONPI_VENDOR")
	model := os.Getenv("REMONPI_MODEL")
	dbPath := os.Getenv("REMONPI_DATABASE_PATH")
	r := remote.NewRemote(vendor, model, dbPath)

	// Initialize HTTP Server...
	httpPort := os.Getenv("HTTP_PORT")
	if len(httpPort) == 0 {
		httpPort = "8080"
	}
	err := startHTTPServer(httpPort, r)
	if err != nil {
		panic(err)
	}
}
