package energybrowser

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"go-w3-implementations/energy-browser/contracts/energy-consumption"
	"go-w3-implementations/energy-browser/contracts/energy-production"
	"go-w3-implementations/ethereum"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

// Type holding all necessary data to communicate with the EnergyBrowser contracts
type EnergyBrowserCom struct {
	eth                 *ethereum.Ethereum
	from                common.Address
	privateKey          *ecdsa.PrivateKey
	productionContract  *energyproduction.Energyproduction
	consumptionContract *energyconsumption.Energyconsumption
}

func NewEnergyBrowser(eth *ethereum.Ethereum, from common.Address, privateKey *ecdsa.PrivateKey) *EnergyBrowserCom {

	productionInstance, err := energyproduction.NewEnergyproduction(common.HexToAddress("0x12c6b301ed57ceb5e4f185eeee1b27b2a3da0393"), eth.Client)
	if err != nil {
		log.Fatal(err)
	}
	consumptionInstance, err := energyconsumption.NewEnergyconsumption(common.HexToAddress("0x8f93f0bef6e23b112a179702c9fbfea06042530d"), eth.Client)
	if err != nil {
		log.Fatal(err)
	}

	return &EnergyBrowserCom{
		eth:                 eth,
		from:                from,
		privateKey:          privateKey,
		productionContract:  productionInstance,
		consumptionContract: consumptionInstance,
	}
}

func (browser EnergyBrowserCom) WriteEnergyValues(consumption uint32, production uint32) {
	auth := browser.getAuthObject(browser.eth.Client)
	fmt.Println(auth.Nonce)
	tx, err := browser.consumptionContract.SetEnerConsumption(auth, consumption)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("tx sent to ConsumptionContract: %s \n", tx.Hash().Hex())
	auth = browser.getAuthObject(browser.eth.Client)
	auth.Nonce = auth.Nonce.Add(auth.Nonce, big.NewInt(1))
	fmt.Println(auth.Nonce)
	tx, err = browser.productionContract.SetEnerProduction(auth, production)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("tx sent to ProductionContract: %s \n", tx.Hash().Hex())
}

func (browser EnergyBrowserCom) getAuthObject(_client *ethclient.Client) *bind.TransactOpts {
	nonce, err := _client.PendingNonceAt(context.Background(), browser.from)
	if err != nil {
		fmt.Println(err)
	}

	auth := bind.NewKeyedTransactor(browser.privateKey)
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)
	auth.GasLimit = uint64(1500000)
	auth.GasPrice = big.NewInt(50000000000)

	return auth
}
