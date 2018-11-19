// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package energyproduction

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

// EnergyproductionABI is the input ABI used to generate the binding from.
const EnergyproductionABI = "[{\"constant\":true,\"inputs\":[{\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"getLocationForCoin\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"getEnerProductionForCoin\",\"outputs\":[{\"name\":\"\",\"type\":\"uint32[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_proAccntAddr\",\"type\":\"address\"}],\"name\":\"getProBalance\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getProducer\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"},{\"name\":\"\",\"type\":\"string\"},{\"name\":\"\",\"type\":\"string\"},{\"name\":\"\",\"type\":\"uint32\"},{\"name\":\"\",\"type\":\"uint32\"},{\"name\":\"\",\"type\":\"uint32\"},{\"name\":\"\",\"type\":\"uint32\"},{\"name\":\"\",\"type\":\"uint32\"},{\"name\":\"\",\"type\":\"string\"},{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_proAccntAddr\",\"type\":\"address\"}],\"name\":\"getProEnerProduction\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"},{\"name\":\"\",\"type\":\"uint256[]\"},{\"name\":\"\",\"type\":\"uint32[]\"},{\"name\":\"\",\"type\":\"uint256[]\"},{\"name\":\"\",\"type\":\"bytes32[]\"},{\"name\":\"\",\"type\":\"uint256[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getProAccntsList\",\"outputs\":[{\"name\":\"\",\"type\":\"address[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"producerAddr\",\"type\":\"address\"}],\"name\":\"proAccntsArr\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"countProducers\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getEnerProduction\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"},{\"name\":\"\",\"type\":\"uint256\"},{\"name\":\"\",\"type\":\"uint32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"proAccntList\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"getDeviceTypeForCoin\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_proAccntAddr\",\"type\":\"address\"}],\"name\":\"getProAccntDetails\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"},{\"name\":\"\",\"type\":\"string\"},{\"name\":\"\",\"type\":\"uint32\"},{\"name\":\"\",\"type\":\"string\"},{\"name\":\"\",\"type\":\"uint32\"},{\"name\":\"\",\"type\":\"uint32\"},{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_owner\",\"type\":\"string\"},{\"name\":\"_deviceType\",\"type\":\"string\"},{\"name\":\"_peakPowerPos\",\"type\":\"uint32\"},{\"name\":\"_peakPowerNeg\",\"type\":\"uint32\"},{\"name\":\"_latitude\",\"type\":\"uint32\"},{\"name\":\"_longitude\",\"type\":\"uint32\"},{\"name\":\"_voltageLevel\",\"type\":\"uint32\"},{\"name\":\"_location\",\"type\":\"string\"},{\"name\":\"_installDate\",\"type\":\"string\"}],\"name\":\"setProducer\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_enerValue\",\"type\":\"uint32\"}],\"name\":\"setEnerProduction\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"pvAddr\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"owner\",\"type\":\"string\"},{\"indexed\":false,\"name\":\"deviceType\",\"type\":\"string\"},{\"indexed\":false,\"name\":\"peakPowerPos\",\"type\":\"uint32\"},{\"indexed\":false,\"name\":\"peakPowerNeg\",\"type\":\"uint32\"},{\"indexed\":false,\"name\":\"latitude\",\"type\":\"uint32\"},{\"indexed\":false,\"name\":\"longitude\",\"type\":\"uint32\"},{\"indexed\":false,\"name\":\"voltageLevel\",\"type\":\"uint32\"},{\"indexed\":false,\"name\":\"location\",\"type\":\"string\"},{\"indexed\":false,\"name\":\"installDate\",\"type\":\"string\"}],\"name\":\"ProducerRegs\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"oliAddr\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"eTime\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"enerAmount\",\"type\":\"uint32\"}],\"name\":\"EnerProductionEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"oliAddr\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"txTime\",\"type\":\"uint256[]\"},{\"indexed\":false,\"name\":\"txValue\",\"type\":\"uint32[]\"},{\"indexed\":false,\"name\":\"blockNumber\",\"type\":\"uint256[]\"},{\"indexed\":false,\"name\":\"blockHash\",\"type\":\"bytes32[]\"},{\"indexed\":false,\"name\":\"txGasPrice\",\"type\":\"uint256[]\"}],\"name\":\"ProTransactionEvent\",\"type\":\"event\"}]"

// Energyproduction is an auto generated Go binding around an Ethereum contract.
type Energyproduction struct {
	EnergyproductionCaller     // Read-only binding to the contract
	EnergyproductionTransactor // Write-only binding to the contract
	EnergyproductionFilterer   // Log filterer for contract events
}

// EnergyproductionCaller is an auto generated read-only Go binding around an Ethereum contract.
type EnergyproductionCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EnergyproductionTransactor is an auto generated write-only Go binding around an Ethereum contract.
type EnergyproductionTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EnergyproductionFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type EnergyproductionFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EnergyproductionSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type EnergyproductionSession struct {
	Contract     *Energyproduction // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// EnergyproductionCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type EnergyproductionCallerSession struct {
	Contract *EnergyproductionCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts           // Call options to use throughout this session
}

// EnergyproductionTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type EnergyproductionTransactorSession struct {
	Contract     *EnergyproductionTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts           // Transaction auth options to use throughout this session
}

// EnergyproductionRaw is an auto generated low-level Go binding around an Ethereum contract.
type EnergyproductionRaw struct {
	Contract *Energyproduction // Generic contract binding to access the raw methods on
}

// EnergyproductionCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type EnergyproductionCallerRaw struct {
	Contract *EnergyproductionCaller // Generic read-only contract binding to access the raw methods on
}

// EnergyproductionTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type EnergyproductionTransactorRaw struct {
	Contract *EnergyproductionTransactor // Generic write-only contract binding to access the raw methods on
}

// NewEnergyproduction creates a new instance of Energyproduction, bound to a specific deployed contract.
func NewEnergyproduction(address common.Address, backend bind.ContractBackend) (*Energyproduction, error) {
	contract, err := bindEnergyproduction(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Energyproduction{EnergyproductionCaller: EnergyproductionCaller{contract: contract}, EnergyproductionTransactor: EnergyproductionTransactor{contract: contract}, EnergyproductionFilterer: EnergyproductionFilterer{contract: contract}}, nil
}

// NewEnergyproductionCaller creates a new read-only instance of Energyproduction, bound to a specific deployed contract.
func NewEnergyproductionCaller(address common.Address, caller bind.ContractCaller) (*EnergyproductionCaller, error) {
	contract, err := bindEnergyproduction(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &EnergyproductionCaller{contract: contract}, nil
}

// NewEnergyproductionTransactor creates a new write-only instance of Energyproduction, bound to a specific deployed contract.
func NewEnergyproductionTransactor(address common.Address, transactor bind.ContractTransactor) (*EnergyproductionTransactor, error) {
	contract, err := bindEnergyproduction(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &EnergyproductionTransactor{contract: contract}, nil
}

// NewEnergyproductionFilterer creates a new log filterer instance of Energyproduction, bound to a specific deployed contract.
func NewEnergyproductionFilterer(address common.Address, filterer bind.ContractFilterer) (*EnergyproductionFilterer, error) {
	contract, err := bindEnergyproduction(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &EnergyproductionFilterer{contract: contract}, nil
}

// bindEnergyproduction binds a generic wrapper to an already deployed contract.
func bindEnergyproduction(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(EnergyproductionABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Energyproduction *EnergyproductionRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Energyproduction.Contract.EnergyproductionCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Energyproduction *EnergyproductionRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Energyproduction.Contract.EnergyproductionTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Energyproduction *EnergyproductionRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Energyproduction.Contract.EnergyproductionTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Energyproduction *EnergyproductionCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Energyproduction.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Energyproduction *EnergyproductionTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Energyproduction.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Energyproduction *EnergyproductionTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Energyproduction.Contract.contract.Transact(opts, method, params...)
}

// CountProducers is a free data retrieval call binding the contract method 0xba0cfdce.
//
// Solidity: function countProducers() constant returns(uint256)
func (_Energyproduction *EnergyproductionCaller) CountProducers(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Energyproduction.contract.Call(opts, out, "countProducers")
	return *ret0, err
}

// CountProducers is a free data retrieval call binding the contract method 0xba0cfdce.
//
// Solidity: function countProducers() constant returns(uint256)
func (_Energyproduction *EnergyproductionSession) CountProducers() (*big.Int, error) {
	return _Energyproduction.Contract.CountProducers(&_Energyproduction.CallOpts)
}

// CountProducers is a free data retrieval call binding the contract method 0xba0cfdce.
//
// Solidity: function countProducers() constant returns(uint256)
func (_Energyproduction *EnergyproductionCallerSession) CountProducers() (*big.Int, error) {
	return _Energyproduction.Contract.CountProducers(&_Energyproduction.CallOpts)
}

// GetDeviceTypeForCoin is a free data retrieval call binding the contract method 0xdb9a89b5.
//
// Solidity: function getDeviceTypeForCoin(addr address) constant returns(string)
func (_Energyproduction *EnergyproductionCaller) GetDeviceTypeForCoin(opts *bind.CallOpts, addr common.Address) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _Energyproduction.contract.Call(opts, out, "getDeviceTypeForCoin", addr)
	return *ret0, err
}

// GetDeviceTypeForCoin is a free data retrieval call binding the contract method 0xdb9a89b5.
//
// Solidity: function getDeviceTypeForCoin(addr address) constant returns(string)
func (_Energyproduction *EnergyproductionSession) GetDeviceTypeForCoin(addr common.Address) (string, error) {
	return _Energyproduction.Contract.GetDeviceTypeForCoin(&_Energyproduction.CallOpts, addr)
}

// GetDeviceTypeForCoin is a free data retrieval call binding the contract method 0xdb9a89b5.
//
// Solidity: function getDeviceTypeForCoin(addr address) constant returns(string)
func (_Energyproduction *EnergyproductionCallerSession) GetDeviceTypeForCoin(addr common.Address) (string, error) {
	return _Energyproduction.Contract.GetDeviceTypeForCoin(&_Energyproduction.CallOpts, addr)
}

// GetEnerProduction is a free data retrieval call binding the contract method 0xc15ab975.
//
// Solidity: function getEnerProduction() constant returns(address, uint256, uint32)
func (_Energyproduction *EnergyproductionCaller) GetEnerProduction(opts *bind.CallOpts) (common.Address, *big.Int, uint32, error) {
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
	err := _Energyproduction.contract.Call(opts, out, "getEnerProduction")
	return *ret0, *ret1, *ret2, err
}

// GetEnerProduction is a free data retrieval call binding the contract method 0xc15ab975.
//
// Solidity: function getEnerProduction() constant returns(address, uint256, uint32)
func (_Energyproduction *EnergyproductionSession) GetEnerProduction() (common.Address, *big.Int, uint32, error) {
	return _Energyproduction.Contract.GetEnerProduction(&_Energyproduction.CallOpts)
}

// GetEnerProduction is a free data retrieval call binding the contract method 0xc15ab975.
//
// Solidity: function getEnerProduction() constant returns(address, uint256, uint32)
func (_Energyproduction *EnergyproductionCallerSession) GetEnerProduction() (common.Address, *big.Int, uint32, error) {
	return _Energyproduction.Contract.GetEnerProduction(&_Energyproduction.CallOpts)
}

// GetEnerProductionForCoin is a free data retrieval call binding the contract method 0x12a32118.
//
// Solidity: function getEnerProductionForCoin(addr address) constant returns(uint32[])
func (_Energyproduction *EnergyproductionCaller) GetEnerProductionForCoin(opts *bind.CallOpts, addr common.Address) ([]uint32, error) {
	var (
		ret0 = new([]uint32)
	)
	out := ret0
	err := _Energyproduction.contract.Call(opts, out, "getEnerProductionForCoin", addr)
	return *ret0, err
}

// GetEnerProductionForCoin is a free data retrieval call binding the contract method 0x12a32118.
//
// Solidity: function getEnerProductionForCoin(addr address) constant returns(uint32[])
func (_Energyproduction *EnergyproductionSession) GetEnerProductionForCoin(addr common.Address) ([]uint32, error) {
	return _Energyproduction.Contract.GetEnerProductionForCoin(&_Energyproduction.CallOpts, addr)
}

// GetEnerProductionForCoin is a free data retrieval call binding the contract method 0x12a32118.
//
// Solidity: function getEnerProductionForCoin(addr address) constant returns(uint32[])
func (_Energyproduction *EnergyproductionCallerSession) GetEnerProductionForCoin(addr common.Address) ([]uint32, error) {
	return _Energyproduction.Contract.GetEnerProductionForCoin(&_Energyproduction.CallOpts, addr)
}

// GetLocationForCoin is a free data retrieval call binding the contract method 0x03e85f52.
//
// Solidity: function getLocationForCoin(addr address) constant returns(string)
func (_Energyproduction *EnergyproductionCaller) GetLocationForCoin(opts *bind.CallOpts, addr common.Address) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _Energyproduction.contract.Call(opts, out, "getLocationForCoin", addr)
	return *ret0, err
}

// GetLocationForCoin is a free data retrieval call binding the contract method 0x03e85f52.
//
// Solidity: function getLocationForCoin(addr address) constant returns(string)
func (_Energyproduction *EnergyproductionSession) GetLocationForCoin(addr common.Address) (string, error) {
	return _Energyproduction.Contract.GetLocationForCoin(&_Energyproduction.CallOpts, addr)
}

// GetLocationForCoin is a free data retrieval call binding the contract method 0x03e85f52.
//
// Solidity: function getLocationForCoin(addr address) constant returns(string)
func (_Energyproduction *EnergyproductionCallerSession) GetLocationForCoin(addr common.Address) (string, error) {
	return _Energyproduction.Contract.GetLocationForCoin(&_Energyproduction.CallOpts, addr)
}

// GetProAccntDetails is a free data retrieval call binding the contract method 0xdd1063aa.
//
// Solidity: function getProAccntDetails(_proAccntAddr address) constant returns(string, string, uint32, string, uint32, uint32, string)
func (_Energyproduction *EnergyproductionCaller) GetProAccntDetails(opts *bind.CallOpts, _proAccntAddr common.Address) (string, string, uint32, string, uint32, uint32, string, error) {
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
	err := _Energyproduction.contract.Call(opts, out, "getProAccntDetails", _proAccntAddr)
	return *ret0, *ret1, *ret2, *ret3, *ret4, *ret5, *ret6, err
}

// GetProAccntDetails is a free data retrieval call binding the contract method 0xdd1063aa.
//
// Solidity: function getProAccntDetails(_proAccntAddr address) constant returns(string, string, uint32, string, uint32, uint32, string)
func (_Energyproduction *EnergyproductionSession) GetProAccntDetails(_proAccntAddr common.Address) (string, string, uint32, string, uint32, uint32, string, error) {
	return _Energyproduction.Contract.GetProAccntDetails(&_Energyproduction.CallOpts, _proAccntAddr)
}

// GetProAccntDetails is a free data retrieval call binding the contract method 0xdd1063aa.
//
// Solidity: function getProAccntDetails(_proAccntAddr address) constant returns(string, string, uint32, string, uint32, uint32, string)
func (_Energyproduction *EnergyproductionCallerSession) GetProAccntDetails(_proAccntAddr common.Address) (string, string, uint32, string, uint32, uint32, string, error) {
	return _Energyproduction.Contract.GetProAccntDetails(&_Energyproduction.CallOpts, _proAccntAddr)
}

// GetProAccntsList is a free data retrieval call binding the contract method 0x4bb4577f.
//
// Solidity: function getProAccntsList() constant returns(address[])
func (_Energyproduction *EnergyproductionCaller) GetProAccntsList(opts *bind.CallOpts) ([]common.Address, error) {
	var (
		ret0 = new([]common.Address)
	)
	out := ret0
	err := _Energyproduction.contract.Call(opts, out, "getProAccntsList")
	return *ret0, err
}

// GetProAccntsList is a free data retrieval call binding the contract method 0x4bb4577f.
//
// Solidity: function getProAccntsList() constant returns(address[])
func (_Energyproduction *EnergyproductionSession) GetProAccntsList() ([]common.Address, error) {
	return _Energyproduction.Contract.GetProAccntsList(&_Energyproduction.CallOpts)
}

// GetProAccntsList is a free data retrieval call binding the contract method 0x4bb4577f.
//
// Solidity: function getProAccntsList() constant returns(address[])
func (_Energyproduction *EnergyproductionCallerSession) GetProAccntsList() ([]common.Address, error) {
	return _Energyproduction.Contract.GetProAccntsList(&_Energyproduction.CallOpts)
}

// GetProBalance is a free data retrieval call binding the contract method 0x29b48069.
//
// Solidity: function getProBalance(_proAccntAddr address) constant returns(uint256)
func (_Energyproduction *EnergyproductionCaller) GetProBalance(opts *bind.CallOpts, _proAccntAddr common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Energyproduction.contract.Call(opts, out, "getProBalance", _proAccntAddr)
	return *ret0, err
}

// GetProBalance is a free data retrieval call binding the contract method 0x29b48069.
//
// Solidity: function getProBalance(_proAccntAddr address) constant returns(uint256)
func (_Energyproduction *EnergyproductionSession) GetProBalance(_proAccntAddr common.Address) (*big.Int, error) {
	return _Energyproduction.Contract.GetProBalance(&_Energyproduction.CallOpts, _proAccntAddr)
}

// GetProBalance is a free data retrieval call binding the contract method 0x29b48069.
//
// Solidity: function getProBalance(_proAccntAddr address) constant returns(uint256)
func (_Energyproduction *EnergyproductionCallerSession) GetProBalance(_proAccntAddr common.Address) (*big.Int, error) {
	return _Energyproduction.Contract.GetProBalance(&_Energyproduction.CallOpts, _proAccntAddr)
}

// GetProEnerProduction is a free data retrieval call binding the contract method 0x3ddc0835.
//
// Solidity: function getProEnerProduction(_proAccntAddr address) constant returns(address, uint256[], uint32[], uint256[], bytes32[], uint256[])
func (_Energyproduction *EnergyproductionCaller) GetProEnerProduction(opts *bind.CallOpts, _proAccntAddr common.Address) (common.Address, []*big.Int, []uint32, []*big.Int, [][32]byte, []*big.Int, error) {
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
	err := _Energyproduction.contract.Call(opts, out, "getProEnerProduction", _proAccntAddr)
	return *ret0, *ret1, *ret2, *ret3, *ret4, *ret5, err
}

// GetProEnerProduction is a free data retrieval call binding the contract method 0x3ddc0835.
//
// Solidity: function getProEnerProduction(_proAccntAddr address) constant returns(address, uint256[], uint32[], uint256[], bytes32[], uint256[])
func (_Energyproduction *EnergyproductionSession) GetProEnerProduction(_proAccntAddr common.Address) (common.Address, []*big.Int, []uint32, []*big.Int, [][32]byte, []*big.Int, error) {
	return _Energyproduction.Contract.GetProEnerProduction(&_Energyproduction.CallOpts, _proAccntAddr)
}

// GetProEnerProduction is a free data retrieval call binding the contract method 0x3ddc0835.
//
// Solidity: function getProEnerProduction(_proAccntAddr address) constant returns(address, uint256[], uint32[], uint256[], bytes32[], uint256[])
func (_Energyproduction *EnergyproductionCallerSession) GetProEnerProduction(_proAccntAddr common.Address) (common.Address, []*big.Int, []uint32, []*big.Int, [][32]byte, []*big.Int, error) {
	return _Energyproduction.Contract.GetProEnerProduction(&_Energyproduction.CallOpts, _proAccntAddr)
}

// GetProducer is a free data retrieval call binding the contract method 0x2f183cfd.
//
// Solidity: function getProducer() constant returns(address, string, string, uint32, uint32, uint32, uint32, uint32, string, string)
func (_Energyproduction *EnergyproductionCaller) GetProducer(opts *bind.CallOpts) (common.Address, string, string, uint32, uint32, uint32, uint32, uint32, string, string, error) {
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
	err := _Energyproduction.contract.Call(opts, out, "getProducer")
	return *ret0, *ret1, *ret2, *ret3, *ret4, *ret5, *ret6, *ret7, *ret8, *ret9, err
}

// GetProducer is a free data retrieval call binding the contract method 0x2f183cfd.
//
// Solidity: function getProducer() constant returns(address, string, string, uint32, uint32, uint32, uint32, uint32, string, string)
func (_Energyproduction *EnergyproductionSession) GetProducer() (common.Address, string, string, uint32, uint32, uint32, uint32, uint32, string, string, error) {
	return _Energyproduction.Contract.GetProducer(&_Energyproduction.CallOpts)
}

// GetProducer is a free data retrieval call binding the contract method 0x2f183cfd.
//
// Solidity: function getProducer() constant returns(address, string, string, uint32, uint32, uint32, uint32, uint32, string, string)
func (_Energyproduction *EnergyproductionCallerSession) GetProducer() (common.Address, string, string, uint32, uint32, uint32, uint32, uint32, string, string, error) {
	return _Energyproduction.Contract.GetProducer(&_Energyproduction.CallOpts)
}

// ProAccntList is a free data retrieval call binding the contract method 0xce0ccd09.
//
// Solidity: function proAccntList( uint256) constant returns(address)
func (_Energyproduction *EnergyproductionCaller) ProAccntList(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Energyproduction.contract.Call(opts, out, "proAccntList", arg0)
	return *ret0, err
}

// ProAccntList is a free data retrieval call binding the contract method 0xce0ccd09.
//
// Solidity: function proAccntList( uint256) constant returns(address)
func (_Energyproduction *EnergyproductionSession) ProAccntList(arg0 *big.Int) (common.Address, error) {
	return _Energyproduction.Contract.ProAccntList(&_Energyproduction.CallOpts, arg0)
}

// ProAccntList is a free data retrieval call binding the contract method 0xce0ccd09.
//
// Solidity: function proAccntList( uint256) constant returns(address)
func (_Energyproduction *EnergyproductionCallerSession) ProAccntList(arg0 *big.Int) (common.Address, error) {
	return _Energyproduction.Contract.ProAccntList(&_Energyproduction.CallOpts, arg0)
}

// ProAccntsArr is a free data retrieval call binding the contract method 0x7d0691fb.
//
// Solidity: function proAccntsArr(producerAddr address) constant returns(bool)
func (_Energyproduction *EnergyproductionCaller) ProAccntsArr(opts *bind.CallOpts, producerAddr common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Energyproduction.contract.Call(opts, out, "proAccntsArr", producerAddr)
	return *ret0, err
}

// ProAccntsArr is a free data retrieval call binding the contract method 0x7d0691fb.
//
// Solidity: function proAccntsArr(producerAddr address) constant returns(bool)
func (_Energyproduction *EnergyproductionSession) ProAccntsArr(producerAddr common.Address) (bool, error) {
	return _Energyproduction.Contract.ProAccntsArr(&_Energyproduction.CallOpts, producerAddr)
}

// ProAccntsArr is a free data retrieval call binding the contract method 0x7d0691fb.
//
// Solidity: function proAccntsArr(producerAddr address) constant returns(bool)
func (_Energyproduction *EnergyproductionCallerSession) ProAccntsArr(producerAddr common.Address) (bool, error) {
	return _Energyproduction.Contract.ProAccntsArr(&_Energyproduction.CallOpts, producerAddr)
}

// SetEnerProduction is a paid mutator transaction binding the contract method 0xf1f1ecb7.
//
// Solidity: function setEnerProduction(_enerValue uint32) returns()
func (_Energyproduction *EnergyproductionTransactor) SetEnerProduction(opts *bind.TransactOpts, _enerValue uint32) (*types.Transaction, error) {
	return _Energyproduction.contract.Transact(opts, "setEnerProduction", _enerValue)
}

// SetEnerProduction is a paid mutator transaction binding the contract method 0xf1f1ecb7.
//
// Solidity: function setEnerProduction(_enerValue uint32) returns()
func (_Energyproduction *EnergyproductionSession) SetEnerProduction(_enerValue uint32) (*types.Transaction, error) {
	return _Energyproduction.Contract.SetEnerProduction(&_Energyproduction.TransactOpts, _enerValue)
}

// SetEnerProduction is a paid mutator transaction binding the contract method 0xf1f1ecb7.
//
// Solidity: function setEnerProduction(_enerValue uint32) returns()
func (_Energyproduction *EnergyproductionTransactorSession) SetEnerProduction(_enerValue uint32) (*types.Transaction, error) {
	return _Energyproduction.Contract.SetEnerProduction(&_Energyproduction.TransactOpts, _enerValue)
}

// SetProducer is a paid mutator transaction binding the contract method 0xe1a9249c.
//
// Solidity: function setProducer(_owner string, _deviceType string, _peakPowerPos uint32, _peakPowerNeg uint32, _latitude uint32, _longitude uint32, _voltageLevel uint32, _location string, _installDate string) returns()
func (_Energyproduction *EnergyproductionTransactor) SetProducer(opts *bind.TransactOpts, _owner string, _deviceType string, _peakPowerPos uint32, _peakPowerNeg uint32, _latitude uint32, _longitude uint32, _voltageLevel uint32, _location string, _installDate string) (*types.Transaction, error) {
	return _Energyproduction.contract.Transact(opts, "setProducer", _owner, _deviceType, _peakPowerPos, _peakPowerNeg, _latitude, _longitude, _voltageLevel, _location, _installDate)
}

// SetProducer is a paid mutator transaction binding the contract method 0xe1a9249c.
//
// Solidity: function setProducer(_owner string, _deviceType string, _peakPowerPos uint32, _peakPowerNeg uint32, _latitude uint32, _longitude uint32, _voltageLevel uint32, _location string, _installDate string) returns()
func (_Energyproduction *EnergyproductionSession) SetProducer(_owner string, _deviceType string, _peakPowerPos uint32, _peakPowerNeg uint32, _latitude uint32, _longitude uint32, _voltageLevel uint32, _location string, _installDate string) (*types.Transaction, error) {
	return _Energyproduction.Contract.SetProducer(&_Energyproduction.TransactOpts, _owner, _deviceType, _peakPowerPos, _peakPowerNeg, _latitude, _longitude, _voltageLevel, _location, _installDate)
}

// SetProducer is a paid mutator transaction binding the contract method 0xe1a9249c.
//
// Solidity: function setProducer(_owner string, _deviceType string, _peakPowerPos uint32, _peakPowerNeg uint32, _latitude uint32, _longitude uint32, _voltageLevel uint32, _location string, _installDate string) returns()
func (_Energyproduction *EnergyproductionTransactorSession) SetProducer(_owner string, _deviceType string, _peakPowerPos uint32, _peakPowerNeg uint32, _latitude uint32, _longitude uint32, _voltageLevel uint32, _location string, _installDate string) (*types.Transaction, error) {
	return _Energyproduction.Contract.SetProducer(&_Energyproduction.TransactOpts, _owner, _deviceType, _peakPowerPos, _peakPowerNeg, _latitude, _longitude, _voltageLevel, _location, _installDate)
}

// EnergyproductionEnerProductionEventIterator is returned from FilterEnerProductionEvent and is used to iterate over the raw logs and unpacked data for EnerProductionEvent events raised by the Energyproduction contract.
type EnergyproductionEnerProductionEventIterator struct {
	Event *EnergyproductionEnerProductionEvent // Event containing the contract specifics and raw log

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
func (it *EnergyproductionEnerProductionEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EnergyproductionEnerProductionEvent)
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
		it.Event = new(EnergyproductionEnerProductionEvent)
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
func (it *EnergyproductionEnerProductionEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EnergyproductionEnerProductionEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EnergyproductionEnerProductionEvent represents a EnerProductionEvent event raised by the Energyproduction contract.
type EnergyproductionEnerProductionEvent struct {
	OliAddr    common.Address
	ETime      *big.Int
	EnerAmount uint32
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterEnerProductionEvent is a free log retrieval operation binding the contract event 0x745f59dd4d2101f37334b5e8def36fa361850e77d5ea97b24fde9c098425081f.
//
// Solidity: e EnerProductionEvent(oliAddr address, eTime uint256, enerAmount uint32)
func (_Energyproduction *EnergyproductionFilterer) FilterEnerProductionEvent(opts *bind.FilterOpts) (*EnergyproductionEnerProductionEventIterator, error) {

	logs, sub, err := _Energyproduction.contract.FilterLogs(opts, "EnerProductionEvent")
	if err != nil {
		return nil, err
	}
	return &EnergyproductionEnerProductionEventIterator{contract: _Energyproduction.contract, event: "EnerProductionEvent", logs: logs, sub: sub}, nil
}

// WatchEnerProductionEvent is a free log subscription operation binding the contract event 0x745f59dd4d2101f37334b5e8def36fa361850e77d5ea97b24fde9c098425081f.
//
// Solidity: e EnerProductionEvent(oliAddr address, eTime uint256, enerAmount uint32)
func (_Energyproduction *EnergyproductionFilterer) WatchEnerProductionEvent(opts *bind.WatchOpts, sink chan<- *EnergyproductionEnerProductionEvent) (event.Subscription, error) {

	logs, sub, err := _Energyproduction.contract.WatchLogs(opts, "EnerProductionEvent")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EnergyproductionEnerProductionEvent)
				if err := _Energyproduction.contract.UnpackLog(event, "EnerProductionEvent", log); err != nil {
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

// EnergyproductionProTransactionEventIterator is returned from FilterProTransactionEvent and is used to iterate over the raw logs and unpacked data for ProTransactionEvent events raised by the Energyproduction contract.
type EnergyproductionProTransactionEventIterator struct {
	Event *EnergyproductionProTransactionEvent // Event containing the contract specifics and raw log

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
func (it *EnergyproductionProTransactionEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EnergyproductionProTransactionEvent)
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
		it.Event = new(EnergyproductionProTransactionEvent)
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
func (it *EnergyproductionProTransactionEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EnergyproductionProTransactionEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EnergyproductionProTransactionEvent represents a ProTransactionEvent event raised by the Energyproduction contract.
type EnergyproductionProTransactionEvent struct {
	OliAddr     common.Address
	TxTime      []*big.Int
	TxValue     []uint32
	BlockNumber []*big.Int
	BlockHash   [][32]byte
	TxGasPrice  []*big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterProTransactionEvent is a free log retrieval operation binding the contract event 0x138c48eb84787286b8d6204838378ff1e3bad9ee92224577f0c005ac19ae9e59.
//
// Solidity: e ProTransactionEvent(oliAddr address, txTime uint256[], txValue uint32[], blockNumber uint256[], blockHash bytes32[], txGasPrice uint256[])
func (_Energyproduction *EnergyproductionFilterer) FilterProTransactionEvent(opts *bind.FilterOpts) (*EnergyproductionProTransactionEventIterator, error) {

	logs, sub, err := _Energyproduction.contract.FilterLogs(opts, "ProTransactionEvent")
	if err != nil {
		return nil, err
	}
	return &EnergyproductionProTransactionEventIterator{contract: _Energyproduction.contract, event: "ProTransactionEvent", logs: logs, sub: sub}, nil
}

// WatchProTransactionEvent is a free log subscription operation binding the contract event 0x138c48eb84787286b8d6204838378ff1e3bad9ee92224577f0c005ac19ae9e59.
//
// Solidity: e ProTransactionEvent(oliAddr address, txTime uint256[], txValue uint32[], blockNumber uint256[], blockHash bytes32[], txGasPrice uint256[])
func (_Energyproduction *EnergyproductionFilterer) WatchProTransactionEvent(opts *bind.WatchOpts, sink chan<- *EnergyproductionProTransactionEvent) (event.Subscription, error) {

	logs, sub, err := _Energyproduction.contract.WatchLogs(opts, "ProTransactionEvent")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EnergyproductionProTransactionEvent)
				if err := _Energyproduction.contract.UnpackLog(event, "ProTransactionEvent", log); err != nil {
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

// EnergyproductionProducerRegsIterator is returned from FilterProducerRegs and is used to iterate over the raw logs and unpacked data for ProducerRegs events raised by the Energyproduction contract.
type EnergyproductionProducerRegsIterator struct {
	Event *EnergyproductionProducerRegs // Event containing the contract specifics and raw log

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
func (it *EnergyproductionProducerRegsIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EnergyproductionProducerRegs)
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
		it.Event = new(EnergyproductionProducerRegs)
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
func (it *EnergyproductionProducerRegsIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EnergyproductionProducerRegsIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EnergyproductionProducerRegs represents a ProducerRegs event raised by the Energyproduction contract.
type EnergyproductionProducerRegs struct {
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

// FilterProducerRegs is a free log retrieval operation binding the contract event 0x5bba84d3568a0b013b9eac75d4ba0c2aa6affa2d3d931aff1565c817e7438c40.
//
// Solidity: e ProducerRegs(pvAddr address, owner string, deviceType string, peakPowerPos uint32, peakPowerNeg uint32, latitude uint32, longitude uint32, voltageLevel uint32, location string, installDate string)
func (_Energyproduction *EnergyproductionFilterer) FilterProducerRegs(opts *bind.FilterOpts) (*EnergyproductionProducerRegsIterator, error) {

	logs, sub, err := _Energyproduction.contract.FilterLogs(opts, "ProducerRegs")
	if err != nil {
		return nil, err
	}
	return &EnergyproductionProducerRegsIterator{contract: _Energyproduction.contract, event: "ProducerRegs", logs: logs, sub: sub}, nil
}

// WatchProducerRegs is a free log subscription operation binding the contract event 0x5bba84d3568a0b013b9eac75d4ba0c2aa6affa2d3d931aff1565c817e7438c40.
//
// Solidity: e ProducerRegs(pvAddr address, owner string, deviceType string, peakPowerPos uint32, peakPowerNeg uint32, latitude uint32, longitude uint32, voltageLevel uint32, location string, installDate string)
func (_Energyproduction *EnergyproductionFilterer) WatchProducerRegs(opts *bind.WatchOpts, sink chan<- *EnergyproductionProducerRegs) (event.Subscription, error) {

	logs, sub, err := _Energyproduction.contract.WatchLogs(opts, "ProducerRegs")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EnergyproductionProducerRegs)
				if err := _Energyproduction.contract.UnpackLog(event, "ProducerRegs", log); err != nil {
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
