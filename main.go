package main

import (
	"fmt"
	"go-w3-implementations/creds"
	"go-w3-implementations/database"
	"go-w3-implementations/energy-browser"
	"go-w3-implementations/ethereum"
	"log"
	"os"
	"os/signal"
	"time"

	"github.com/ethereum/go-ethereum/common"
)

func main() {
	eth := ethereum.Connect()
	db := database.Connect()
	prepareShutdown(db, eth)
	from := common.HexToAddress("0x22563388B5f57fe921661B721B87E490c18A6392")
	browser := energybrowser.NewEnergyBrowser(eth, from, keysandcredentials.GetPrivateKey())
	for {
		sendEnergyValue(db, eth, browser)
		time.Sleep(time.Second * 60)
	}

	eth.DisConnect()
	db.Disconnect()
}

func sendEnergyValue(db *database.Database, eth *ethereum.Ethereum, browser *energybrowser.EnergyBrowserCom) {
	result := db.ReadFromMeter("ORDER BY timestamp DESC LIMIT 1")

	for _, element := range result {
		writeout := fmt.Sprintf("Current database value: %d", element)
		fmt.Println(writeout)
		browser.WriteEnergyValues(uint32(element.Consumption), uint32(element.Production))
	}
}

func prepareShutdown(db *database.Database, eth *ethereum.Ethereum) {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		for sig := range c {
			log.Printf("captured %v, stopping profiler and exiting ...", sig)
			if db != nil {
				db.Disconnect()
			}
			if eth != nil {
				eth.DisConnect()
			}
			os.Exit(1)
		}
	}()
}
