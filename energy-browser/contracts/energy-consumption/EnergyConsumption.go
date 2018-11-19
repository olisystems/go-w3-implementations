// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package energyconsumption

import (
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = abi.U256
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// EnergyconsumptionABI is the input ABI used to generate the binding from.
const EnergyconsumptionABI = "[{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"consAccntList\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getConsAccntsList\",\"outputs\":[{\"name\":\"\",\"type\":\"address[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getEnerConsumption\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"},{\"name\":\"\",\"type\":\"uint256\"},{\"name\":\"\",\"type\":\"uint32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"consumerAddr\",\"type\":\"address\"}],\"name\":\"consAccntsArr\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getConsumer\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"},{\"name\":\"\",\"type\":\"string\"},{\"name\":\"\",\"type\":\"string\"},{\"name\":\"\",\"type\":\"uint32\"},{\"name\":\"\",\"type\":\"uint32\"},{\"name\":\"\",\"type\":\"uint32\"},{\"name\":\"\",\"type\":\"uint32\"},{\"name\":\"\",\"type\":\"uint32\"},{\"name\":\"\",\"type\":\"string\"},{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"countConsumers\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_consAccntAddr\",\"type\":\"address\"}],\"name\":\"getConsBalance\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_enerValue\",\"type\":\"uint32\"}],\"name\":\"setEnerConsumption\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_owner\",\"type\":\"string\"},{\"name\":\"_deviceType\",\"type\":\"string\"},{\"name\":\"_peakPowerPos\",\"type\":\"uint32\"},{\"name\":\"_peakPowerNeg\",\"type\":\"uint32\"},{\"name\":\"_latitude\",\"type\":\"uint32\"},{\"name\":\"_longitude\",\"type\":\"uint32\"},{\"name\":\"_voltageLevel\",\"type\":\"uint32\"},{\"name\":\"_location\",\"type\":\"string\"},{\"name\":\"_installDate\",\"type\":\"string\"}],\"name\":\"setConsumer\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_consAccntAddr\",\"type\":\"address\"}],\"name\":\"getConsAccntDetails\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"},{\"name\":\"\",\"type\":\"string\"},{\"name\":\"\",\"type\":\"uint32\"},{\"name\":\"\",\"type\":\"string\"},{\"name\":\"\",\"type\":\"uint32\"},{\"name\":\"\",\"type\":\"uint32\"},{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_consAccntAddr\",\"type\":\"address\"}],\"name\":\"getConsEnerConsumption\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"},{\"name\":\"\",\"type\":\"uint256[]\"},{\"name\":\"\",\"type\":\"uint32[]\"},{\"name\":\"\",\"type\":\"uint256[]\"},{\"name\":\"\",\"type\":\"bytes32[]\"},{\"name\":\"\",\"type\":\"uint256[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"pvAddr\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"owner\",\"type\":\"string\"},{\"indexed\":false,\"name\":\"deviceType\",\"type\":\"string\"},{\"indexed\":false,\"name\":\"peakPowerPos\",\"type\":\"uint32\"},{\"indexed\":false,\"name\":\"peakPowerNeg\",\"type\":\"uint32\"},{\"indexed\":false,\"name\":\"latitude\",\"type\":\"uint32\"},{\"indexed\":false,\"name\":\"longitude\",\"type\":\"uint32\"},{\"indexed\":false,\"name\":\"voltageLevel\",\"type\":\"uint32\"},{\"indexed\":false,\"name\":\"location\",\"type\":\"string\"},{\"indexed\":false,\"name\":\"installDate\",\"type\":\"string\"}],\"name\":\"ConsumerRegs\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"oliAddr\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"eTime\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"enerAmount\",\"type\":\"uint32\"}],\"name\":\"EnerConsumptionEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"oliAddr\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"txTime\",\"type\":\"uint256[]\"},{\"indexed\":false,\"name\":\"txValue\",\"type\":\"uint32[]\"},{\"indexed\":false,\"name\":\"blockNumber\",\"type\":\"uint256[]\"},{\"indexed\":false,\"name\":\"blockHash\",\"type\":\"bytes32[]\"},{\"indexed\":false,\"name\":\"txGasPrice\",\"type\":\"uint256[]\"}],\"name\":\"ConsTransactionEvent\",\"type\":\"event\"}]"

// Energyconsumption is an auto generated Go binding around an Ethereum contract.
type Energyconsumption struct {
	EnergyconsumptionCaller     // Read-only binding to the contract
	EnergyconsumptionTransactor // Write-only binding to the contract
	EnergyconsumptionFilterer   // Log filterer for contract events
}

// EnergyconsumptionCaller is an auto generated read-only Go binding around an Ethereum contract.
type EnergyconsumptionCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EnergyconsumptionTransactor is an auto generated write-only Go binding around an Ethereum contract.
type EnergyconsumptionTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EnergyconsumptionFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type EnergyconsumptionFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EnergyconsumptionSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type EnergyconsumptionSession struct {
	Contract     *Energyconsumption // Generic contract binding to set the session for
	CallOpts     bind.CallOpts      // Call options to use throughout this session
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// EnergyconsumptionCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type EnergyconsumptionCallerSession struct {
	Contract *EnergyconsumptionCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts            // Call options to use throughout this session
}

// EnergyconsumptionTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type EnergyconsumptionTransactorSession struct {
	Contract     *EnergyconsumptionTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts            // Transaction auth options to use throughout this session
}

// EnergyconsumptionRaw is an auto generated low-level Go binding around an Ethereum contract.
type EnergyconsumptionRaw struct {
	Contract *Energyconsumption // Generic contract binding to access the raw methods on
}

// EnergyconsumptionCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type EnergyconsumptionCallerRaw struct {
	Contract *EnergyconsumptionCaller // Generic read-only contract binding to access the raw methods on
}

// EnergyconsumptionTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type EnergyconsumptionTransactorRaw struct {
	Contract *EnergyconsumptionTransactor // Generic write-only contract binding to access the raw methods on
}

// NewEnergyconsumption creates a new instance of Energyconsumption, bound to a specific deployed contract.
func NewEnergyconsumption(address common.Address, backend bind.ContractBackend) (*Energyconsumption, error) {
	contract, err := bindEnergyconsumption(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Energyconsumption{EnergyconsumptionCaller: EnergyconsumptionCaller{contract: contract}, EnergyconsumptionTransactor: EnergyconsumptionTransactor{contract: contract}, EnergyconsumptionFilterer: EnergyconsumptionFilterer{contract: contract}}, nil
}

// NewEnergyconsumptionCaller creates a new read-only instance of Energyconsumption, bound to a specific deployed contract.
func NewEnergyconsumptionCaller(address common.Address, caller bind.ContractCaller) (*EnergyconsumptionCaller, error) {
	contract, err := bindEnergyconsumption(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &EnergyconsumptionCaller{contract: contract}, nil
}

// NewEnergyconsumptionTransactor creates a new write-only instance of Energyconsumption, bound to a specific deployed contract.
func NewEnergyconsumptionTransactor(address common.Address, transactor bind.ContractTransactor) (*EnergyconsumptionTransactor, error) {
	contract, err := bindEnergyconsumption(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &EnergyconsumptionTransactor{contract: contract}, nil
}

// NewEnergyconsumptionFilterer creates a new log filterer instance of Energyconsumption, bound to a specific deployed contract.
func NewEnergyconsumptionFilterer(address common.Address, filterer bind.ContractFilterer) (*EnergyconsumptionFilterer, error) {
	contract, err := bindEnergyconsumption(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &EnergyconsumptionFilterer{contract: contract}, nil
}

// bindEnergyconsumption binds a generic wrapper to an already deployed contract.
func bindEnergyconsumption(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(EnergyconsumptionABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Energyconsumption *EnergyconsumptionRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Energyconsumption.Contract.EnergyconsumptionCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Energyconsumption *EnergyconsumptionRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Energyconsumption.Contract.EnergyconsumptionTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Energyconsumption *EnergyconsumptionRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Energyconsumption.Contract.EnergyconsumptionTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Energyconsumption *EnergyconsumptionCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Energyconsumption.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Energyconsumption *EnergyconsumptionTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Energyconsumption.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Energyconsumption *EnergyconsumptionTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Energyconsumption.Contract.contract.Transact(opts, method, params...)
}

// ConsAccntList is a free data retrieval call binding the contract method 0x18cbdce3.
//
// Solidity: function consAccntList( uint256) constant returns(address)
func (_Energyconsumption *EnergyconsumptionCaller) ConsAccntList(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Energyconsumption.contract.Call(opts, out, "consAccntList", arg0)
	return *ret0, err
}

// ConsAccntList is a free data retrieval call binding the contract method 0x18cbdce3.
//
// Solidity: function consAccntList( uint256) constant returns(address)
func (_Energyconsumption *EnergyconsumptionSession) ConsAccntList(arg0 *big.Int) (common.Address, error) {
	return _Energyconsumption.Contract.ConsAccntList(&_Energyconsumption.CallOpts, arg0)
}

// ConsAccntList is a free data retrieval call binding the contract method 0x18cbdce3.
//
// Solidity: function consAccntList( uint256) constant returns(address)
func (_Energyconsumption *EnergyconsumptionCallerSession) ConsAccntList(arg0 *big.Int) (common.Address, error) {
	return _Energyconsumption.Contract.ConsAccntList(&_Energyconsumption.CallOpts, arg0)
}

// ConsAccntsArr is a free data retrieval call binding the contract method 0x24a0d6bc.
//
// Solidity: function consAccntsArr(consumerAddr address) constant returns(bool)
func (_Energyconsumption *EnergyconsumptionCaller) ConsAccntsArr(opts *bind.CallOpts, consumerAddr common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Energyconsumption.contract.Call(opts, out, "consAccntsArr", consumerAddr)
	return *ret0, err
}

// ConsAccntsArr is a free data retrieval call binding the contract method 0x24a0d6bc.
//
// Solidity: function consAccntsArr(consumerAddr address) constant returns(bool)
func (_Energyconsumption *EnergyconsumptionSession) ConsAccntsArr(consumerAddr common.Address) (bool, error) {
	return _Energyconsumption.Contract.ConsAccntsArr(&_Energyconsumption.CallOpts, consumerAddr)
}

// ConsAccntsArr is a free data retrieval call binding the contract method 0x24a0d6bc.
//
// Solidity: function consAccntsArr(consumerAddr address) constant returns(bool)
func (_Energyconsumption *EnergyconsumptionCallerSession) ConsAccntsArr(consumerAddr common.Address) (bool, error) {
	return _Energyconsumption.Contract.ConsAccntsArr(&_Energyconsumption.CallOpts, consumerAddr)
}

// CountConsumers is a free data retrieval call binding the contract method 0x739ae785.
//
// Solidity: function countConsumers() constant returns(uint256)
func (_Energyconsumption *EnergyconsumptionCaller) CountConsumers(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Energyconsumption.contract.Call(opts, out, "countConsumers")
	return *ret0, err
}

// CountConsumers is a free data retrieval call binding the contract method 0x739ae785.
//
// Solidity: function countConsumers() constant returns(uint256)
func (_Energyconsumption *EnergyconsumptionSession) CountConsumers() (*big.Int, error) {
	return _Energyconsumption.Contract.CountConsumers(&_Energyconsumption.CallOpts)
}

// CountConsumers is a free data retrieval call binding the contract method 0x739ae785.
//
// Solidity: function countConsumers() constant returns(uint256)
func (_Energyconsumption *EnergyconsumptionCallerSession) CountConsumers() (*big.Int, error) {
	return _Energyconsumption.Contract.CountConsumers(&_Energyconsumption.CallOpts)
}

// GetConsAccntDetails is a free data retrieval call binding the contract method 0xb1b79dbf.
//
// Solidity: function getConsAccntDetails(_consAccntAddr address) constant returns(string, string, uint32, string, uint32, uint32, string)
func (_Energyconsumption *EnergyconsumptionCaller) GetConsAccntDetails(opts *bind.CallOpts, _consAccntAddr common.Address) (string, string, uint32, string, uint32, uint32, string, error) {
	var (
		ret0 = new(string)
		ret1 = new(string)
		ret2 = new(uint32)
		ret3 = new(string)
		ret4 = new(uint32)
		ret5 = new(uint32)
		ret6 = new(string)
	)
	out := &[]interface{}{
		ret0,
		ret1,
		ret2,
		ret3,
		ret4,
		ret5,
		ret6,
	}
	err := _Energyconsumption.contract.Call(opts, out, "getConsAccntDetails", _consAccntAddr)
	return *ret0, *ret1, *ret2, *ret3, *ret4, *ret5, *ret6, err
}

// GetConsAccntDetails is a free data retrieval call binding the contract method 0xb1b79dbf.
//
// Solidity: function getConsAccntDetails(_consAccntAddr address) constant returns(string, string, uint32, string, uint32, uint32, string)
func (_Energyconsumption *EnergyconsumptionSession) GetConsAccntDetails(_consAccntAddr common.Address) (string, string, uint32, string, uint32, uint32, string, error) {
	return _Energyconsumption.Contract.GetConsAccntDetails(&_Energyconsumption.CallOpts, _consAccntAddr)
}

// GetConsAccntDetails is a free data retrieval call binding the contract method 0xb1b79dbf.
//
// Solidity: function getConsAccntDetails(_consAccntAddr address) constant returns(string, string, uint32, string, uint32, uint32, string)
func (_Energyconsumption *EnergyconsumptionCallerSession) GetConsAccntDetails(_consAccntAddr common.Address) (string, string, uint32, string, uint32, uint32, string, error) {
	return _Energyconsumption.Contract.GetConsAccntDetails(&_Energyconsumption.CallOpts, _consAccntAddr)
}

// GetConsAccntsList is a free data retrieval call binding the contract method 0x19bccb9c.
//
// Solidity: function getConsAccntsList() constant returns(address[])
func (_Energyconsumption *EnergyconsumptionCaller) GetConsAccntsList(opts *bind.CallOpts) ([]common.Address, error) {
	var (
		ret0 = new([]common.Address)
	)
	out := ret0
	err := _Energyconsumption.contract.Call(opts, out, "getConsAccntsList")
	return *ret0, err
}

// GetConsAccntsList is a free data retrieval call binding the contract method 0x19bccb9c.
//
// Solidity: function getConsAccntsList() constant returns(address[])
func (_Energyconsumption *EnergyconsumptionSession) GetConsAccntsList() ([]common.Address, error) {
	return _Energyconsumption.Contract.GetConsAccntsList(&_Energyconsumption.CallOpts)
}

// GetConsAccntsList is a free data retrieval call binding the contract method 0x19bccb9c.
//
// Solidity: function getConsAccntsList() constant returns(address[])
func (_Energyconsumption *EnergyconsumptionCallerSession) GetConsAccntsList() ([]common.Address, error) {
	return _Energyconsumption.Contract.GetConsAccntsList(&_Energyconsumption.CallOpts)
}

// GetConsBalance is a free data retrieval call binding the contract method 0x7ed8917f.
//
// Solidity: function getConsBalance(_consAccntAddr address) constant returns(uint256)
func (_Energyconsumption *EnergyconsumptionCaller) GetConsBalance(opts *bind.CallOpts, _consAccntAddr common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Energyconsumption.contract.Call(opts, out, "getConsBalance", _consAccntAddr)
	return *ret0, err
}

// GetConsBalance is a free data retrieval call binding the contract method 0x7ed8917f.
//
// Solidity: function getConsBalance(_consAccntAddr address) constant returns(uint256)
func (_Energyconsumption *EnergyconsumptionSession) GetConsBalance(_consAccntAddr common.Address) (*big.Int, error) {
	return _Energyconsumption.Contract.GetConsBalance(&_Energyconsumption.CallOpts, _consAccntAddr)
}

// GetConsBalance is a free data retrieval call binding the contract method 0x7ed8917f.
//
// Solidity: function getConsBalance(_consAccntAddr address) constant returns(uint256)
func (_Energyconsumption *EnergyconsumptionCallerSession) GetConsBalance(_consAccntAddr common.Address) (*big.Int, error) {
	return _Energyconsumption.Contract.GetConsBalance(&_Energyconsumption.CallOpts, _consAccntAddr)
}

// GetConsEnerConsumption is a free data retrieval call binding the contract method 0xfd3c13ab.
//
// Solidity: function getConsEnerConsumption(_consAccntAddr address) constant returns(address, uint256[], uint32[], uint256[], bytes32[], uint256[])
func (_Energyconsumption *EnergyconsumptionCaller) GetConsEnerConsumption(opts *bind.CallOpts, _consAccntAddr common.Address) (common.Address, []*big.Int, []uint32, []*big.Int, [][32]byte, []*big.Int, error) {
	var (
		ret0 = new(common.Address)
		ret1 = new([]*big.Int)
		ret2 = new([]uint32)
		ret3 = new([]*big.Int)
		ret4 = new([][32]byte)
		ret5 = new([]*big.Int)
	)
	out := &[]interface{}{
		ret0,
		ret1,
		ret2,
		ret3,
		ret4,
		ret5,
	}
	err := _Energyconsumption.contract.Call(opts, out, "getConsEnerConsumption", _consAccntAddr)
	return *ret0, *ret1, *ret2, *ret3, *ret4, *ret5, err
}

// GetConsEnerConsumption is a free data retrieval call binding the contract method 0xfd3c13ab.
//
// Solidity: function getConsEnerConsumption(_consAccntAddr address) constant returns(address, uint256[], uint32[], uint256[], bytes32[], uint256[])
func (_Energyconsumption *EnergyconsumptionSession) GetConsEnerConsumption(_consAccntAddr common.Address) (common.Address, []*big.Int, []uint32, []*big.Int, [][32]byte, []*big.Int, error) {
	return _Energyconsumption.Contract.GetConsEnerConsumption(&_Energyconsumption.CallOpts, _consAccntAddr)
}

// GetConsEnerConsumption is a free data retrieval call binding the contract method 0xfd3c13ab.
//
// Solidity: function getConsEnerConsumption(_consAccntAddr address) constant returns(address, uint256[], uint32[], uint256[], bytes32[], uint256[])
func (_Energyconsumption *EnergyconsumptionCallerSession) GetConsEnerConsumption(_consAccntAddr common.Address) (common.Address, []*big.Int, []uint32, []*big.Int, [][32]byte, []*big.Int, error) {
	return _Energyconsumption.Contract.GetConsEnerConsumption(&_Energyconsumption.CallOpts, _consAccntAddr)
}

// GetConsumer is a free data retrieval call binding the contract method 0x6c07681f.
//
// Solidity: function getConsumer() constant returns(address, string, string, uint32, uint32, uint32, uint32, uint32, string, string)
func (_Energyconsumption *EnergyconsumptionCaller) GetConsumer(opts *bind.CallOpts) (common.Address, string, string, uint32, uint32, uint32, uint32, uint32, string, string, error) {
	var (
		ret0 = new(common.Address)
		ret1 = new(string)
		ret2 = new(string)
		ret3 = new(uint32)
		ret4 = new(uint32)
		ret5 = new(uint32)
		ret6 = new(uint32)
		ret7 = new(uint32)
		ret8 = new(string)
		ret9 = new(string)
	)
	out := &[]interface{}{
		ret0,
		ret1,
		ret2,
		ret3,
		ret4,
		ret5,
		ret6,
		ret7,
		ret8,
		ret9,
	}
	err := _Energyconsumption.contract.Call(opts, out, "getConsumer")
	return *ret0, *ret1, *ret2, *ret3, *ret4, *ret5, *ret6, *ret7, *ret8, *ret9, err
}

// GetConsumer is a free data retrieval call binding the contract method 0x6c07681f.
//
// Solidity: function getConsumer() constant returns(address, string, string, uint32, uint32, uint32, uint32, uint32, string, string)
func (_Energyconsumption *EnergyconsumptionSession) GetConsumer() (common.Address, string, string, uint32, uint32, uint32, uint32, uint32, string, string, error) {
	return _Energyconsumption.Contract.GetConsumer(&_Energyconsumption.CallOpts)
}

// GetConsumer is a free data retrieval call binding the contract method 0x6c07681f.
//
// Solidity: function getConsumer() constant returns(address, string, string, uint32, uint32, uint32, uint32, uint32, string, string)
func (_Energyconsumption *EnergyconsumptionCallerSession) GetConsumer() (common.Address, string, string, uint32, uint32, uint32, uint32, uint32, string, string, error) {
	return _Energyconsumption.Contract.GetConsumer(&_Energyconsumption.CallOpts)
}

// GetEnerConsumption is a free data retrieval call binding the contract method 0x203aec33.
//
// Solidity: function getEnerConsumption() constant returns(address, uint256, uint32)
func (_Energyconsumption *EnergyconsumptionCaller) GetEnerConsumption(opts *bind.CallOpts) (common.Address, *big.Int, uint32, error) {
	var (
		ret0 = new(common.Address)
		ret1 = new(*big.Int)
		ret2 = new(uint32)
	)
	out := &[]interface{}{
		ret0,
		ret1,
		ret2,
	}
	err := _Energyconsumption.contract.Call(opts, out, "getEnerConsumption")
	return *ret0, *ret1, *ret2, err
}

// GetEnerConsumption is a free data retrieval call binding the contract method 0x203aec33.
//
// Solidity: function getEnerConsumption() constant returns(address, uint256, uint32)
func (_Energyconsumption *EnergyconsumptionSession) GetEnerConsumption() (common.Address, *big.Int, uint32, error) {
	return _Energyconsumption.Contract.GetEnerConsumption(&_Energyconsumption.CallOpts)
}

// GetEnerConsumption is a free data retrieval call binding the contract method 0x203aec33.
//
// Solidity: function getEnerConsumption() constant returns(address, uint256, uint32)
func (_Energyconsumption *EnergyconsumptionCallerSession) GetEnerConsumption() (common.Address, *big.Int, uint32, error) {
	return _Energyconsumption.Contract.GetEnerConsumption(&_Energyconsumption.CallOpts)
}

// SetConsumer is a paid mutator transaction binding the contract method 0xa64062fc.
//
// Solidity: function setConsumer(_owner string, _deviceType string, _peakPowerPos uint32, _peakPowerNeg uint32, _latitude uint32, _longitude uint32, _voltageLevel uint32, _location string, _installDate string) returns()
func (_Energyconsumption *EnergyconsumptionTransactor) SetConsumer(opts *bind.TransactOpts, _owner string, _deviceType string, _peakPowerPos uint32, _peakPowerNeg uint32, _latitude uint32, _longitude uint32, _voltageLevel uint32, _location string, _installDate string) (*types.Transaction, error) {
	return _Energyconsumption.contract.Transact(opts, "setConsumer", _owner, _deviceType, _peakPowerPos, _peakPowerNeg, _latitude, _longitude, _voltageLevel, _location, _installDate)
}

// SetConsumer is a paid mutator transaction binding the contract method 0xa64062fc.
//
// Solidity: function setConsumer(_owner string, _deviceType string, _peakPowerPos uint32, _peakPowerNeg uint32, _latitude uint32, _longitude uint32, _voltageLevel uint32, _location string, _installDate string) returns()
func (_Energyconsumption *EnergyconsumptionSession) SetConsumer(_owner string, _deviceType string, _peakPowerPos uint32, _peakPowerNeg uint32, _latitude uint32, _longitude uint32, _voltageLevel uint32, _location string, _installDate string) (*types.Transaction, error) {
	return _Energyconsumption.Contract.SetConsumer(&_Energyconsumption.TransactOpts, _owner, _deviceType, _peakPowerPos, _peakPowerNeg, _latitude, _longitude, _voltageLevel, _location, _installDate)
}

// SetConsumer is a paid mutator transaction binding the contract method 0xa64062fc.
//
// Solidity: function setConsumer(_owner string, _deviceType string, _peakPowerPos uint32, _peakPowerNeg uint32, _latitude uint32, _longitude uint32, _voltageLevel uint32, _location string, _installDate string) returns()
func (_Energyconsumption *EnergyconsumptionTransactorSession) SetConsumer(_owner string, _deviceType string, _peakPowerPos uint32, _peakPowerNeg uint32, _latitude uint32, _longitude uint32, _voltageLevel uint32, _location string, _installDate string) (*types.Transaction, error) {
	return _Energyconsumption.Contract.SetConsumer(&_Energyconsumption.TransactOpts, _owner, _deviceType, _peakPowerPos, _peakPowerNeg, _latitude, _longitude, _voltageLevel, _location, _installDate)
}

// SetEnerConsumption is a paid mutator transaction binding the contract method 0x9926188f.
//
// Solidity: function setEnerConsumption(_enerValue uint32) returns()
func (_Energyconsumption *EnergyconsumptionTransactor) SetEnerConsumption(opts *bind.TransactOpts, _enerValue uint32) (*types.Transaction, error) {
	return _Energyconsumption.contract.Transact(opts, "setEnerConsumption", _enerValue)
}

// SetEnerConsumption is a paid mutator transaction binding the contract method 0x9926188f.
//
// Solidity: function setEnerConsumption(_enerValue uint32) returns()
func (_Energyconsumption *EnergyconsumptionSession) SetEnerConsumption(_enerValue uint32) (*types.Transaction, error) {
	return _Energyconsumption.Contract.SetEnerConsumption(&_Energyconsumption.TransactOpts, _enerValue)
}

// SetEnerConsumption is a paid mutator transaction binding the contract method 0x9926188f.
//
// Solidity: function setEnerConsumption(_enerValue uint32) returns()
func (_Energyconsumption *EnergyconsumptionTransactorSession) SetEnerConsumption(_enerValue uint32) (*types.Transaction, error) {
	return _Energyconsumption.Contract.SetEnerConsumption(&_Energyconsumption.TransactOpts, _enerValue)
}

// EnergyconsumptionConsTransactionEventIterator is returned from FilterConsTransactionEvent and is used to iterate over the raw logs and unpacked data for ConsTransactionEvent events raised by the Energyconsumption contract.
type EnergyconsumptionConsTransactionEventIterator struct {
	Event *EnergyconsumptionConsTransactionEvent // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *EnergyconsumptionConsTransactionEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EnergyconsumptionConsTransactionEvent)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(EnergyconsumptionConsTransactionEvent)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *EnergyconsumptionConsTransactionEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EnergyconsumptionConsTransactionEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EnergyconsumptionConsTransactionEvent represents a ConsTransactionEvent event raised by the Energyconsumption contract.
type EnergyconsumptionConsTransactionEvent struct {
	OliAddr     common.Address
	TxTime      []*big.Int
	TxValue     []uint32
	BlockNumber []*big.Int
	BlockHash   [][32]byte
	TxGasPrice  []*big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterConsTransactionEvent is a free log retrieval operation binding the contract event 0x359a3081bfbd4cd11fe54293feae203de45e9918596cc4f47d82962b52d57a2e.
//
// Solidity: e ConsTransactionEvent(oliAddr address, txTime uint256[], txValue uint32[], blockNumber uint256[], blockHash bytes32[], txGasPrice uint256[])
func (_Energyconsumption *EnergyconsumptionFilterer) FilterConsTransactionEvent(opts *bind.FilterOpts) (*EnergyconsumptionConsTransactionEventIterator, error) {

	logs, sub, err := _Energyconsumption.contract.FilterLogs(opts, "ConsTransactionEvent")
	if err != nil {
		return nil, err
	}
	return &EnergyconsumptionConsTransactionEventIterator{contract: _Energyconsumption.contract, event: "ConsTransactionEvent", logs: logs, sub: sub}, nil
}

// WatchConsTransactionEvent is a free log subscription operation binding the contract event 0x359a3081bfbd4cd11fe54293feae203de45e9918596cc4f47d82962b52d57a2e.
//
// Solidity: e ConsTransactionEvent(oliAddr address, txTime uint256[], txValue uint32[], blockNumber uint256[], blockHash bytes32[], txGasPrice uint256[])
func (_Energyconsumption *EnergyconsumptionFilterer) WatchConsTransactionEvent(opts *bind.WatchOpts, sink chan<- *EnergyconsumptionConsTransactionEvent) (event.Subscription, error) {

	logs, sub, err := _Energyconsumption.contract.WatchLogs(opts, "ConsTransactionEvent")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EnergyconsumptionConsTransactionEvent)
				if err := _Energyconsumption.contract.UnpackLog(event, "ConsTransactionEvent", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// EnergyconsumptionConsumerRegsIterator is returned from FilterConsumerRegs and is used to iterate over the raw logs and unpacked data for ConsumerRegs events raised by the Energyconsumption contract.
type EnergyconsumptionConsumerRegsIterator struct {
	Event *EnergyconsumptionConsumerRegs // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *EnergyconsumptionConsumerRegsIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EnergyconsumptionConsumerRegs)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(EnergyconsumptionConsumerRegs)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *EnergyconsumptionConsumerRegsIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EnergyconsumptionConsumerRegsIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EnergyconsumptionConsumerRegs represents a ConsumerRegs event raised by the Energyconsumption contract.
type EnergyconsumptionConsumerRegs struct {
	PvAddr       common.Address
	Owner        string
	DeviceType   string
	PeakPowerPos uint32
	PeakPowerNeg uint32
	Latitude     uint32
	Longitude    uint32
	VoltageLevel uint32
	Location     string
	InstallDate  string
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterConsumerRegs is a free log retrieval operation binding the contract event 0x71537b70bbd67fee09ad13f9c4280471a83004e2e7d0a3f0a9817677cd1cffd5.
//
// Solidity: e ConsumerRegs(pvAddr address, owner string, deviceType string, peakPowerPos uint32, peakPowerNeg uint32, latitude uint32, longitude uint32, voltageLevel uint32, location string, installDate string)
func (_Energyconsumption *EnergyconsumptionFilterer) FilterConsumerRegs(opts *bind.FilterOpts) (*EnergyconsumptionConsumerRegsIterator, error) {

	logs, sub, err := _Energyconsumption.contract.FilterLogs(opts, "ConsumerRegs")
	if err != nil {
		return nil, err
	}
	return &EnergyconsumptionConsumerRegsIterator{contract: _Energyconsumption.contract, event: "ConsumerRegs", logs: logs, sub: sub}, nil
}

// WatchConsumerRegs is a free log subscription operation binding the contract event 0x71537b70bbd67fee09ad13f9c4280471a83004e2e7d0a3f0a9817677cd1cffd5.
//
// Solidity: e ConsumerRegs(pvAddr address, owner string, deviceType string, peakPowerPos uint32, peakPowerNeg uint32, latitude uint32, longitude uint32, voltageLevel uint32, location string, installDate string)
func (_Energyconsumption *EnergyconsumptionFilterer) WatchConsumerRegs(opts *bind.WatchOpts, sink chan<- *EnergyconsumptionConsumerRegs) (event.Subscription, error) {

	logs, sub, err := _Energyconsumption.contract.WatchLogs(opts, "ConsumerRegs")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EnergyconsumptionConsumerRegs)
				if err := _Energyconsumption.contract.UnpackLog(event, "ConsumerRegs", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// EnergyconsumptionEnerConsumptionEventIterator is returned from FilterEnerConsumptionEvent and is used to iterate over the raw logs and unpacked data for EnerConsumptionEvent events raised by the Energyconsumption contract.
type EnergyconsumptionEnerConsumptionEventIterator struct {
	Event *EnergyconsumptionEnerConsumptionEvent // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *EnergyconsumptionEnerConsumptionEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EnergyconsumptionEnerConsumptionEvent)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(EnergyconsumptionEnerConsumptionEvent)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *EnergyconsumptionEnerConsumptionEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EnergyconsumptionEnerConsumptionEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EnergyconsumptionEnerConsumptionEvent represents a EnerConsumptionEvent event raised by the Energyconsumption contract.
type EnergyconsumptionEnerConsumptionEvent struct {
	OliAddr    common.Address
	ETime      *big.Int
	EnerAmount uint32
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterEnerConsumptionEvent is a free log retrieval operation binding the contract event 0x7357384fea585bd35b5329e6fe91e41971d0236bd93c74f999535200a0a93a30.
//
// Solidity: e EnerConsumptionEvent(oliAddr address, eTime uint256, enerAmount uint32)
func (_Energyconsumption *EnergyconsumptionFilterer) FilterEnerConsumptionEvent(opts *bind.FilterOpts) (*EnergyconsumptionEnerConsumptionEventIterator, error) {

	logs, sub, err := _Energyconsumption.contract.FilterLogs(opts, "EnerConsumptionEvent")
	if err != nil {
		return nil, err
	}
	return &EnergyconsumptionEnerConsumptionEventIterator{contract: _Energyconsumption.contract, event: "EnerConsumptionEvent", logs: logs, sub: sub}, nil
}

// WatchEnerConsumptionEvent is a free log subscription operation binding the contract event 0x7357384fea585bd35b5329e6fe91e41971d0236bd93c74f999535200a0a93a30.
//
// Solidity: e EnerConsumptionEvent(oliAddr address, eTime uint256, enerAmount uint32)
func (_Energyconsumption *EnergyconsumptionFilterer) WatchEnerConsumptionEvent(opts *bind.WatchOpts, sink chan<- *EnergyconsumptionEnerConsumptionEvent) (event.Subscription, error) {

	logs, sub, err := _Energyconsumption.contract.WatchLogs(opts, "EnerConsumptionEvent")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EnergyconsumptionEnerConsumptionEvent)
				if err := _Energyconsumption.contract.UnpackLog(event, "EnerConsumptionEvent", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}
