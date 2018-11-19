package etc

import (
	"context"
	"crypto/ecdsa"
	"database/sql"
	"fmt"
	"go-w3-implementations/etc/validator"
	"io/ioutil"
	"log"
	"math/big"
	"os"
	"os/signal"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

func Etc() {
	client, err := ethclient.Dial("ws://localhost:8546")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("we have a connection")

	accountAdress := common.HexToAddress("0x36e4182c362e1f6c0e517b3c6e77500ddcee082e")
	balance, err := client.BalanceAt(context.Background(), accountAdress, nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(balance)

	//account, _, privateKey := getKAccount(client)
	//fmt.Println(account.Address.Hex())
	//to := common.HexToAddress("0x36e4182c362e1f6c0e517b3c6e77500ddcee082e")

	//sendEther(client, ks, account, to)

	//whisper.Whisper()

	//deployContract(client, account, privateKey)

	//validator.NewValidator(common.HexToAddress("0xf5a55d202a8597859576c20fe36e6bb82efdf51b"), client)

	// if err != nil {
	// 	log.Fatal(err)
	// }

	//privateKeys := keysandcredentials.GetPrivateKeys()

	// common.HexToAddress("0x22563388B5f57fe921661B721B87E490c18A6392")

	//readContract(instance, client)
	//writeToContract(instance, client, from, privateKey)

	//database.Database()

	var db *sql.DB

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		for sig := range c {
			log.Printf("captured %v, stopping profiler and exiting ...", sig)
			if db != nil {
				db.Close()
			}
			os.Exit(1)
		}
	}()

	db, err = sql.Open("sqlite3", "./meter.db")
	if err != nil {
		fmt.Print("tet1")
		log.Fatal(err)
	}

	subscribeToEvent(client, db)
}

func getKAccount(_client *ethclient.Client) (accounts.Account, *keystore.KeyStore, *keystore.Key) {

	file := "../res/UTC--2018-07-25T14-41-32Z--432418bd-8a25-8180-00d7-e6375847f139.htm"
	ks := keystore.NewKeyStore("../tmp", keystore.StandardScryptN, keystore.StandardScryptP)
	jsonBytes, err := ioutil.ReadFile(file)
	if err != nil {
		log.Fatal(err)
	}

	password := "AL45hgf"
	account, err := ks.Import(jsonBytes, password, password)
	if err != nil {
		log.Fatal(err)
	}
	ks.Unlock(account, password)
	if err := os.RemoveAll("../tmp"); err != nil {
		log.Fatal(err)
	}

	privateKey, err := keystore.DecryptKey(jsonBytes, password)
	if err != nil {
		log.Fatal(err)
	}
	return account, ks, privateKey
}

func deployContract(_client *ethclient.Client, from accounts.Account, privateKey *ecdsa.PrivateKey) *validator.Validator {

	nonce, err := _client.PendingNonceAt(context.Background(), from.Address)
	if err != nil {
		log.Fatal(err)
	}

	auth := bind.NewKeyedTransactor(privateKey)
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)
	auth.GasLimit = uint64(1500000)
	auth.GasPrice = big.NewInt(50000000000)

	validators := []common.Address{from.Address}

	contractAddress, tx, instance, err := validator.DeployValidator(auth, _client, validators)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(contractAddress.Hex())
	fmt.Println(tx.Hash().Hex())

	return instance
}

func writeToContract(instance *validator.Validator, _client *ethclient.Client, from common.Address, privateKey *ecdsa.PrivateKey) {
	nonce, err := _client.PendingNonceAt(context.Background(), from)
	if err != nil {
		log.Fatal(err)
	}

	auth := bind.NewKeyedTransactor(privateKey)
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)
	auth.GasLimit = uint64(1500000)
	auth.GasPrice = big.NewInt(50000000000)

	tx, err := instance.SetBackup(auth, big.NewInt(12), big.NewInt(41), big.NewInt(50), "More Awesomeness")

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("tx sent: %s", tx.Hash().Hex())
}

func readContract(validatorContract *validator.Validator, _client *ethclient.Client) {

	instance := validatorContract

	firstVerifier, err := instance.Verifiers(nil, big.NewInt(0))

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("First Verifier: %s \n", firstVerifier.Hex())

	verifiedBackup, err := instance.VerifiedBackup(nil)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Verified Backup Hash: %s \n", verifiedBackup.BlockHash)
	fmt.Printf("BlockNumberEnd: %s \n", verifiedBackup.BlockNumberEnd)
	fmt.Printf("BlockNumberStart: %s \n", verifiedBackup.BlockNumberStart)
	fmt.Printf("Timestamp: %s \n", verifiedBackup.Timestamp)

}

func subscribeToEvent(_client *ethclient.Client, db *sql.DB) {

	latestBlock, err := _client.BlockByNumber(context.Background(), nil)

	if err != nil {
		log.Fatal(err)
	}

	query := ethereum.FilterQuery{
		Addresses: []common.Address{common.HexToAddress("0xf5a55d202a8597859576c20fe36e6bb82efdf51b")},
		FromBlock: latestBlock.Number(),
		ToBlock:   nil,
	}

	logs := make(chan types.Log)

	contractAbi, err := abi.JSON(strings.NewReader(string(validator.ValidatorABI)))

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
			fmt.Println(vLog.BlockHash.Hex())
			fmt.Println(vLog.BlockNumber)
			fmt.Println(vLog.TxHash.Hex()) // Might not work

			event := struct {
				Timestamp        *big.Int
				BlockNumberStart *big.Int
				BlockNumberEnd   *big.Int
				BlockHash        string
			}{}

			err := contractAbi.Unpack(&event, "VerifiedBackup", vLog.Data)
			if err != nil {
				log.Fatal(err)
			}

			fmt.Println(event.Timestamp)
			fmt.Println(event.BlockNumberStart)
			fmt.Println(event.BlockNumberEnd)
			fmt.Println(event.BlockHash)

			stmt, err := db.Prepare("insert into backup(timestamp, start, end, hash) values(?,?,?,?)")
			if err != nil {
				log.Fatal(err)
			}

			res, err := stmt.Exec(event.Timestamp.Int64(), event.BlockNumberStart.Int64(), event.BlockNumberEnd.Int64(), event.BlockHash)
			if err != nil {
				log.Fatal(err)
			}

			fmt.Println(res.RowsAffected())
			fmt.Println("________________________________________")
		}
	}
}

func sendEther(_client *ethclient.Client, ks *keystore.KeyStore, from accounts.Account, to common.Address) {
	nonce, err := _client.PendingNonceAt(context.Background(), from.Address)
	if err != nil {
		log.Fatal(err)
	}

	value := big.NewInt(1000000000000000000)
	gasLimit := uint64(21000)
	gasPrice := big.NewInt(50000000000)

	tx := types.NewTransaction(nonce, to, value, gasLimit, gasPrice, nil)

	chainID, err := _client.NetworkID(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	signedTx, err := ks.SignTx(from, tx, chainID)
	if err != nil {
		log.Fatal(err)
	}
	err = _client.SendTransaction(context.Background(), signedTx)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("tx sent: %s", signedTx.Hash().Hex())
}
