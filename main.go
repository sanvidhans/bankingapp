package main

import (
	"github.com/sanvidhans/bankingapp/app"
	"github.com/sanvidhans/bankingapp/logger"
)

func main() {
	logger.Info("Application getting starting...")
	app.Start()
}
