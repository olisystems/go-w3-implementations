package ethereum

import (
	"log"

	"github.com/ethereum/go-ethereum/ethclient"
)

type Ethereum struct {
	Client *ethclient.Client
}

func Connect() *Ethereum {
	client, err := ethclient.Dial("ws://localhost:8546")
	if err != nil {
		log.Fatal(err)
	}
	return &Ethereum{Client: client}
}

func (eth Ethereum) DisConnect() {
	eth.Client.Close()
}
