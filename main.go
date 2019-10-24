package main

import (
	"flag"
	"os"
	"strconv"

	"github.com/ditcraft/indexer/database"
	"github.com/ditcraft/indexer/v2demo"
	"github.com/ditcraft/indexer/v2live"
	"github.com/golang/glog"
	"github.com/joho/godotenv"
)

func main() {
	flag.Set("alsologtostderr", "true")
	flag.Set("log_dir", "./log")
	flag.Set("v", "0")
	flag.Parse()
	glog.Info("Starting ditCraft api...")

	err := godotenv.Load()
	if err != nil {
		glog.Fatal("Error loading .env file")
	}

	err = database.InitDatabase()
	if err != nil {
		glog.Fatal(err)
	}

	v2live.Init()
	v2demo.Init()

	intLastBlock, err := strconv.Atoi(os.Getenv("LAST_BLOCK"))
	if err != nil {
		glog.Fatal(err)
	}

	go v2demo.WatchEvents()
	go v2live.WatchEvents()

	currentBlock, err := v2demo.GetCurrentBlock()
	if err != nil {
		glog.Fatal(err)
	}

	go v2demo.GetOldEvents(int64(intLastBlock), currentBlock)
	go v2live.GetOldEvents(int64(intLastBlock), currentBlock)

	// go api.Start("server.ditcraft.io:3004", 10)

	select {}
}
