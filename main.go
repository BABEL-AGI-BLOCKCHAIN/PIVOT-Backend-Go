package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"github.com/BABEL-AGI-BLOCKCHAIN/PIVOT-Backend-Go/common/config"
	"github.com/BABEL-AGI-BLOCKCHAIN/PIVOT-Backend-Go/common/database"
	"github.com/BABEL-AGI-BLOCKCHAIN/PIVOT-Backend-Go/listener"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

const (
	cfgFile = "config.json"
)

func main() {
	logrus.Info("start the server")
	var err error
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	cfg, err := config.NewConfig(cfgFile)
	if err != nil {
		logrus.Fatal("failed to load config", "err", err)
	}
	logrus.Info("load config success")
	var db *gorm.DB
	db, err = database.InitDB(&cfg.DB)
	if err != nil {
		logrus.Fatal("failed to init db", "err", err)
	}
	logrus.Info("init db success")

	Client, err := ethclient.Dial(cfg.Listener.Endpoint)
	if err != nil {
		logrus.Fatal("failed to connect to geth", "endpoint", cfg.Listener.Endpoint, "err", err)
	}
	logrus.Info("connect to geth success")

	listener, err := listener.NewListener(ctx, cfg.Listener, Client, db)
	if err != nil {
		logrus.Fatal("init client failed: ", err)
	}
	logrus.Info("init listener success")
	go listener.Start()
	logrus.Info("start listener success")
	// Create a channel to listen for OS signals
	sigint := make(chan os.Signal, 1)
	signal.Notify(sigint, syscall.SIGINT, syscall.SIGTERM)

	// Block until a signal is received
	select {
	case <-sigint:
		logrus.Info("stop the server")
		cancel()         // Cancel the context to stop the listener
		listener.Close() // Assuming you have a Stop method to gracefully stop the listener
	}
}
