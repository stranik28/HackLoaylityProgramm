package main

import (
	server "github.com/stranik28/HackLoaylityProgramm/internal"
	"github.com/stranik28/HackLoaylityProgramm/internal/handlers"
	"github.com/stranik28/HackLoaylityProgramm/internal/helper"
	"github.com/stranik28/HackLoaylityProgramm/internal/logger"
	"github.com/stranik28/HackLoaylityProgramm/internal/storage"
	"go.uber.org/zap"
)

func main() {
	err := logger.Init("info")
	server.ParsFlags()
	helper.Init()
	storage.SetupDatabase()
	r := handlers.Routers()
	if err != nil {
		logger.Log.Fatal("Failed to initialize logger", zap.Error(err))
		panic(err)
	}
	logger.Log.Info("Running server", zap.String("address", server.FlagRunAddr))
	err = r.Run(server.FlagRunAddr)
	if err != nil {
		logger.Log.Fatal("Failed to start server", zap.Error(err))
		panic(err)
	}
}
