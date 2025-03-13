package main

import (
	"github.com/alexandrejuniorc/gopportunities/config"
	"github.com/alexandrejuniorc/gopportunities/router"
)

var (
	logger *config.Logger
)

func main() {
	logger = config.GetLogger("main")

	// Initialize Configs
	err := config.Init()
	if err != nil {
		logger.Errorf("Config initialization error: %v", err)
		return
	}

	// Initialize Router
	router.Initialize()
}
