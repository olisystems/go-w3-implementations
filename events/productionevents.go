package main

import (
	"context"
	"fmt"
	"go-w3-implementations/energy-browser/contracts/energy-production"
	_ethereum "go-w3-implementations/ethereum"
	"log"
	"math/big"
	"os"
	"os/signal"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

func main() {
	eth := _ethereum.Connect()
	prepareShutdown(eth)

	subscribeToEvent(eth.Client)

	eth.DisConnect()
}

func subscribeToEvent(_client *ethclient.Client) {

	latestBlock, err := _client.BlockByNumber(context.Background(), nil)
	fmt.Println(latestBlock.Number)
	if err != nil {
		log.Fatal(err)
	}

	query := ethereum.FilterQuery{
		Addresses: []common.Address{common.HexToAddress("0x12c6b301ed57ceb5e4f185eeee1b27b2a3da0393")},
		FromBlock: latestBlock.Number(),
		ToBlock:   nil,
	}

	logs := make(chan types.Log)

	contractAbi, err := abi.JSON(strings.NewReader(string(energyproduction.EnergyproductionABI)))

	sub, err := _client.SubscribeFilterLogs(context.Background(), query, logs)
	if err != nil {
		log.Fatal(err)
	}

	for {
		select {
		case err := <-sub.Err():
			log.Fatal(err)

		case vLog := <-logs:
			fmt.Println("Receiving Event:")
			writeout := fmt.Sprintf("Block Hash: " + vLog.BlockHash.String())
			fmt.Println(writeout)
			writeout = fmt.Sprintf("Block Number: %d", vLog.BlockNumber)
			fmt.Println(writeout)
			//fmt.Println(vLog.TxHash.Hex()) // Might not work

			event := struct {
				OliAddr    common.Address
				ETime      *big.Int
				EnerAmount uint32
			}{}

			err := contractAbi.Unpack(&event, "EnerProductionEvent", vLog.Data)
			if err != nil {
				log.Fatal(err)
			}

			writeout = fmt.Sprintf("Oli Address: " + event.OliAddr.String())
			fmt.Println(writeout)
			writeout = fmt.Sprintf("Energy Timestamp: %d", event.ETime)
			fmt.Println(writeout)
			writeout = fmt.Sprintf("Energy Amount: %d", event.EnerAmount)
			fmt.Println(writeout)

			fmt.Println("________________________________________")
		}
	}
}

func prepareShutdown(eth *_ethereum.Ethereum) {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		for sig := range c {
			log.Printf("captured %v, stopping profiler and exiting ...", sig)
			if eth != nil {
				eth.DisConnect()
			}
			os.Exit(1)
		}
	}()
}
