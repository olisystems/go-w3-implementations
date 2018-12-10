# go-w3-implementations

Some go implementations of smart contracts and interaction with a parity PoA network over the go-ethereum client library.

## Structure

### Database structure

The database (meter.db) is structured in a shema which accepts the following input:

table name: meter

| timestamp | consumption | production |
| --------- | ----------- | ---------- |
| int       | int         | int        |

Before running the application you need to create a database (meter.db) with this shema.

## Compilation

How to compile:

### Prerequisits

Create a file ./creds/keysandcredentials.go where you supply all your keys and credentials for the used ethereum accounts.

The file coud look like the following example:

```go
package keysandcredentials

import (
    "crypto/ecdsa"
    "log"

    "github.com/ethereum/go-ethereum/crypto"

)

func GetPrivateKey() \*ecdsa.PrivateKey {
    privateKey, err := crypto.HexToECDSA("yourprivatekey")
    if err != nil {
        log.Fatal(err)
    }
    return privateKey
}

func GetPassword() string {
    return "apassword"
}
```

### Dependencies

To get missing dependencies run the `go get`command. Current external dependencies:

```bash
$ go get github.com/ethereum/go-ethereum
$ go get github.com/mattn/go-sqlite3
```

### Build for host system

Just run _go build_ in the root directory to build the main package in _main.go_

Run _go build productionsevents.go_ in the ./events directory to build the events package.

### Cross-compile for any target system

Just use [xgo](https://github.com/karalabe/xgo) and the go toolchain which will make your live easy.
Install xgo with:

```bash
$ docker pull karalabe/xgo-latest
$ go get github.com/karalabe/xgo
```

It will download a docker image which has all the necessary flags and environments preinstalled in order to cross-compile go-ethereum

Run xgo in the root folder with:

```bash
$ xgo --targets=linux/arm-7 ./
```

or

```
$ xgo --targets=linux/arm-7 ./events/.
```

#### IMPORTANT notes for compiling the events package for Parity Clients

If running a Parity Client one must first change the gen*log_json.go type. Therefore change the following code in \_go-ethereum/core/types/gen_log_json.go*:

From:

```go
if dec.TxHash == nil {
    return errors.New("missing required field 'transactionHash' for Log")
}
l.TxHash = *dec.TxHash
```

To:

```go
if dec.TxHash != nil {
    l.TxHash = \*dec.TxHash
}
```

A Github issue and pull request will be provided if existent!

## Documentation

A documenation on how to develop with the go-ethereum client library:
https://goethereumbook.org

Github Repository for XGO
https://github.com/karalabe/xgo

Golang tutorial/documentation book
https://astaxie.gitbooks.io/build-web-application-with-golang/en/

Sqlite3
https://www.sqlite.org/index.html

## TODO

- [ ] Error handling
- [ ] Setup script for database
