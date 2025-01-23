package main

import (
	"context"
	"os"
	"os/signal"
	"sync"
	"syscall"

	"github.com/BABEL-AGI-BLOCKCHAIN/PIVOT-Backend-Go/listener"
	"github.com/BABEL-AGI-BLOCKCHAIN/PIVOT-Backend-Go/orm"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

const (
	nodeURL = ""
)

var (
	ContractABI *abi.ABI

	EventSig common.Hash
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	var db *gorm.DB
	var err error
	db, err = orm.InitDB()
	if err != nil {
		logrus.Fatal("failed to init db", "err", err)
	}

	Client, err := ethclient.Dial(nodeURL)
	if err != nil {
		logrus.Fatal("failed to connect to L1 geth", "endpoint", nodeURL, "err", err)
	}
	l1Watcher, err := listener.NewListener(ctx, Client, db)
	if err != nil {
		logrus.Fatal("init L1 client failed: ", err)
	}

	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		defer wg.Done()
		err = l1Watcher.Run(ctx)
		if err != nil {
			logrus.Fatal("l1 client run failed: ", err)
		}
	}()

	// Handle OS signals to gracefully shutdown
	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, syscall.SIGINT, syscall.SIGTERM)
	<-sigCh

	// Cancel the context to stop the listener
	cancel()
	wg.Wait()
	logrus.Println("Program exited gracefully")
}
