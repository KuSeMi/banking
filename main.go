package main

import (
	"github.com/KuSeMi/banking/app"
	"github.com/KuSeMi/banking/logger"
)

func main() {
	logger.Info("Starting the application")
	app.Start()
}
