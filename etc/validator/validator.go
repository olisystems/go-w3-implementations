// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package validator

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

// ValidatorABI is the input ABI used to generate the binding from.
const ValidatorABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"timestampParm\",\"type\":\"uint256\"},{\"name\":\"blockNumberStartParm\",\"type\":\"uint256\"},{\"name\":\"blockNumberEndParm\",\"type\":\"uint256\"},{\"name\":\"blockHashParm\",\"type\":\"string\"}],\"name\":\"setBackup\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"verifiedBackup\",\"outputs\":[{\"name\":\"timestamp\",\"type\":\"uint256\"},{\"name\":\"blockNumberStart\",\"type\":\"uint256\"},{\"name\":\"blockNumberEnd\",\"type\":\"uint256\"},{\"name\":\"blockHash\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"verifiedHash\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"verifiers\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"checkHashes\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"verifiersList\",\"type\":\"address[]\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"timestamp\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"blockNumberStart\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"blockNumberEnd\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"blockHash\",\"type\":\"string\"}],\"name\":\"VerifiedBackup\",\"type\":\"event\"}]"

// ValidatorBin is the compiled bytecode used for deploying new contracts.
const ValidatorBin = `608060405234801561001057600080fd5b50604051610a89380380610a8983398101806040528101908080518201929190505050600080600091505b82518210156100c357828281518110151561005257fe5b9060200190602002015190506001600760008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548160ff021916908315150217905550818060010192505061003b565b82600690805190602001906100d99291906100e2565b505050506101af565b82805482825590600052602060002090810192821561015b579160200282015b8281111561015a5782518260006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555091602001919060010190610102565b5b509050610168919061016c565b5090565b6101ac91905b808211156101a857600081816101000a81549073ffffffffffffffffffffffffffffffffffffffff021916905550600101610172565b5090565b90565b6108cb806101be6000396000f30060806040526004361061006d576000357c0100000000000000000000000000000000000000000000000000000000900463ffffffff1680633f4583691461007257806352bca26e146100f95780638c9b05ee1461019e578063ac1eff68146101d1578063d5f6780a1461023e575b600080fd5b34801561007e57600080fd5b506100f7600480360381019080803590602001909291908035906020019092919080359060200190929190803590602001908201803590602001908080601f016020809104026020016040519081016040528093929190818152602001838380828437820191505050505050919291929050505061026d565b005b34801561010557600080fd5b5061010e610502565b6040518085815260200184815260200183815260200180602001828103825283818151815260200191508051906020019080838360005b83811015610160578082015181840152602081019050610145565b50505050905090810190601f16801561018d5780820380516001836020036101000a031916815260200191505b509550505050505060405180910390f35b3480156101aa57600080fd5b506101b36105b8565b60405180826000191660001916815260200191505060405180910390f35b3480156101dd57600080fd5b506101fc600480360381019080803590602001909291905050506105be565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b34801561024a57600080fd5b506102536105fc565b604051808215151515815260200191505060405180910390f35b6102756107d1565b600060011515600760003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff161515141515610365576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260268152602001807f4f6e6c79207665726966696572732063616e2063616c6c20746869732066756e81526020017f6374696f6e2e000000000000000000000000000000000000000000000000000081525060400191505060405180910390fd5b60806040519081016040528087815260200186815260200185815260200184815250915061039282610738565b905080600560003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002081600019169055506103e46105fc565b156104fa57806000816000191690555081600160008201518160000155602082015181600101556040820151816002015560608201518160030190805190602001906104319291906107fa565b509050507fb43ce4d2c93d9e66cc96abea57b6ff99efce2ac3614305989fb45c4e492790c982600001518360200151846040015185606001516040518085815260200184815260200183815260200180602001828103825283818151815260200191508051906020019080838360005b838110156104bc5780820151818401526020810190506104a1565b50505050905090810190601f1680156104e95780820380516001836020036101000a031916815260200191505b509550505050505060405180910390a15b505050505050565b6001806000015490806001015490806002015490806003018054600181600116156101000203166002900480601f0160208091040260200160405190810160405280929190818152602001828054600181600116156101000203166002900480156105ae5780601f10610583576101008083540402835291602001916105ae565b820191906000526020600020905b81548152906001019060200180831161059157829003601f168201915b5050505050905084565b60005481565b6006818154811015156105cd57fe5b906000526020600020016000915054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b600080600080600560006006600081548110151561061657fe5b9060005260206000200160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054925060019150600090505b60068054905081101561072f57600560006006838154811015156106a757fe5b9060005260206000200160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000205460001916836000191614151561072257600091505b8080600101915050610687565b81935050505090565b60008082600001518360200151846040015185606001516040518085815260200184815260200183815260200182805190602001908083835b6020831015156107965780518252602082019150602081019050602083039250610771565b6001836020036101000a0380198251168184511680821785525050505050509050019450505050506040518091039020905080915050919050565b608060405190810160405280600081526020016000815260200160008152602001606081525090565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f1061083b57805160ff1916838001178555610869565b82800160010185558215610869579182015b8281111561086857825182559160200191906001019061084d565b5b509050610876919061087a565b5090565b61089c91905b80821115610898576000816000905550600101610880565b5090565b905600a165627a7a7230582068cec185bff0e2490c5a3ee90100131cf1d6596d33651966291685668b3f08140029`

// DeployValidator deploys a new Ethereum contract, binding an instance of Validator to it.
func DeployValidator(auth *bind.TransactOpts, backend bind.ContractBackend, verifiersList []common.Address) (common.Address, *types.Transaction, *Validator, error) {
	parsed, err := abi.JSON(strings.NewReader(ValidatorABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(ValidatorBin), backend, verifiersList)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Validator{ValidatorCaller: ValidatorCaller{contract: contract}, ValidatorTransactor: ValidatorTransactor{contract: contract}, ValidatorFilterer: ValidatorFilterer{contract: contract}}, nil
}

// Validator is an auto generated Go binding around an Ethereum contract.
type Validator struct {
	ValidatorCaller     // Read-only binding to the contract
	ValidatorTransactor // Write-only binding to the contract
	ValidatorFilterer   // Log filterer for contract events
}

// ValidatorCaller is an auto generated read-only Go binding around an Ethereum contract.
type ValidatorCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ValidatorTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ValidatorTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ValidatorFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ValidatorFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ValidatorSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ValidatorSession struct {
	Contract     *Validator        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ValidatorCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ValidatorCallerSession struct {
	Contract *ValidatorCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// ValidatorTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ValidatorTransactorSession struct {
	Contract     *ValidatorTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// ValidatorRaw is an auto generated low-level Go binding around an Ethereum contract.
type ValidatorRaw struct {
	Contract *Validator // Generic contract binding to access the raw methods on
}

// ValidatorCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ValidatorCallerRaw struct {
	Contract *ValidatorCaller // Generic read-only contract binding to access the raw methods on
}

// ValidatorTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ValidatorTransactorRaw struct {
	Contract *ValidatorTransactor // Generic write-only contract binding to access the raw methods on
}

// NewValidator creates a new instance of Validator, bound to a specific deployed contract.
func NewValidator(address common.Address, backend bind.ContractBackend) (*Validator, error) {
	contract, err := bindValidator(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Validator{ValidatorCaller: ValidatorCaller{contract: contract}, ValidatorTransactor: ValidatorTransactor{contract: contract}, ValidatorFilterer: ValidatorFilterer{contract: contract}}, nil
}

// NewValidatorCaller creates a new read-only instance of Validator, bound to a specific deployed contract.
func NewValidatorCaller(address common.Address, caller bind.ContractCaller) (*ValidatorCaller, error) {
	contract, err := bindValidator(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ValidatorCaller{contract: contract}, nil
}

// NewValidatorTransactor creates a new write-only instance of Validator, bound to a specific deployed contract.
func NewValidatorTransactor(address common.Address, transactor bind.ContractTransactor) (*ValidatorTransactor, error) {
	contract, err := bindValidator(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ValidatorTransactor{contract: contract}, nil
}

// NewValidatorFilterer creates a new log filterer instance of Validator, bound to a specific deployed contract.
func NewValidatorFilterer(address common.Address, filterer bind.ContractFilterer) (*ValidatorFilterer, error) {
	contract, err := bindValidator(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ValidatorFilterer{contract: contract}, nil
}

// bindValidator binds a generic wrapper to an already deployed contract.
func bindValidator(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ValidatorABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Validator *ValidatorRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Validator.Contract.ValidatorCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Validator *ValidatorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Validator.Contract.ValidatorTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Validator *ValidatorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Validator.Contract.ValidatorTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Validator *ValidatorCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Validator.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Validator *ValidatorTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Validator.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Validator *ValidatorTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Validator.Contract.contract.Transact(opts, method, params...)
}

// VerifiedBackup is a free data retrieval call binding the contract method 0x52bca26e.
//
// Solidity: function verifiedBackup() constant returns(timestamp uint256, blockNumberStart uint256, blockNumberEnd uint256, blockHash string)
func (_Validator *ValidatorCaller) VerifiedBackup(opts *bind.CallOpts) (struct {
	Timestamp        *big.Int
	BlockNumberStart *big.Int
	BlockNumberEnd   *big.Int
	BlockHash        string
}, error) {
	ret := new(struct {
		Timestamp        *big.Int
		BlockNumberStart *big.Int
		BlockNumberEnd   *big.Int
		BlockHash        string
	})
	out := ret
	err := _Validator.contract.Call(opts, out, "verifiedBackup")
	return *ret, err
}

// VerifiedBackup is a free data retrieval call binding the contract method 0x52bca26e.
//
// Solidity: function verifiedBackup() constant returns(timestamp uint256, blockNumberStart uint256, blockNumberEnd uint256, blockHash string)
func (_Validator *ValidatorSession) VerifiedBackup() (struct {
	Timestamp        *big.Int
	BlockNumberStart *big.Int
	BlockNumberEnd   *big.Int
	BlockHash        string
}, error) {
	return _Validator.Contract.VerifiedBackup(&_Validator.CallOpts)
}

// VerifiedBackup is a free data retrieval call binding the contract method 0x52bca26e.
//
// Solidity: function verifiedBackup() constant returns(timestamp uint256, blockNumberStart uint256, blockNumberEnd uint256, blockHash string)
func (_Validator *ValidatorCallerSession) VerifiedBackup() (struct {
	Timestamp        *big.Int
	BlockNumberStart *big.Int
	BlockNumberEnd   *big.Int
	BlockHash        string
}, error) {
	return _Validator.Contract.VerifiedBackup(&_Validator.CallOpts)
}

// VerifiedHash is a free data retrieval call binding the contract method 0x8c9b05ee.
//
// Solidity: function verifiedHash() constant returns(bytes32)
func (_Validator *ValidatorCaller) VerifiedHash(opts *bind.CallOpts) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _Validator.contract.Call(opts, out, "verifiedHash")
	return *ret0, err
}

// VerifiedHash is a free data retrieval call binding the contract method 0x8c9b05ee.
//
// Solidity: function verifiedHash() constant returns(bytes32)
func (_Validator *ValidatorSession) VerifiedHash() ([32]byte, error) {
	return _Validator.Contract.VerifiedHash(&_Validator.CallOpts)
}

// VerifiedHash is a free data retrieval call binding the contract method 0x8c9b05ee.
//
// Solidity: function verifiedHash() constant returns(bytes32)
func (_Validator *ValidatorCallerSession) VerifiedHash() ([32]byte, error) {
	return _Validator.Contract.VerifiedHash(&_Validator.CallOpts)
}

// Verifiers is a free data retrieval call binding the contract method 0xac1eff68.
//
// Solidity: function verifiers( uint256) constant returns(address)
func (_Validator *ValidatorCaller) Verifiers(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Validator.contract.Call(opts, out, "verifiers", arg0)
	return *ret0, err
}

// Verifiers is a free data retrieval call binding the contract method 0xac1eff68.
//
// Solidity: function verifiers( uint256) constant returns(address)
func (_Validator *ValidatorSession) Verifiers(arg0 *big.Int) (common.Address, error) {
	return _Validator.Contract.Verifiers(&_Validator.CallOpts, arg0)
}

// Verifiers is a free data retrieval call binding the contract method 0xac1eff68.
//
// Solidity: function verifiers( uint256) constant returns(address)
func (_Validator *ValidatorCallerSession) Verifiers(arg0 *big.Int) (common.Address, error) {
	return _Validator.Contract.Verifiers(&_Validator.CallOpts, arg0)
}

// CheckHashes is a paid mutator transaction binding the contract method 0xd5f6780a.
//
// Solidity: function checkHashes() returns(bool)
func (_Validator *ValidatorTransactor) CheckHashes(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Validator.contract.Transact(opts, "checkHashes")
}

// CheckHashes is a paid mutator transaction binding the contract method 0xd5f6780a.
//
// Solidity: function checkHashes() returns(bool)
func (_Validator *ValidatorSession) CheckHashes() (*types.Transaction, error) {
	return _Validator.Contract.CheckHashes(&_Validator.TransactOpts)
}

// CheckHashes is a paid mutator transaction binding the contract method 0xd5f6780a.
//
// Solidity: function checkHashes() returns(bool)
func (_Validator *ValidatorTransactorSession) CheckHashes() (*types.Transaction, error) {
	return _Validator.Contract.CheckHashes(&_Validator.TransactOpts)
}

// SetBackup is a paid mutator transaction binding the contract method 0x3f458369.
//
// Solidity: function setBackup(timestampParm uint256, blockNumberStartParm uint256, blockNumberEndParm uint256, blockHashParm string) returns()
func (_Validator *ValidatorTransactor) SetBackup(opts *bind.TransactOpts, timestampParm *big.Int, blockNumberStartParm *big.Int, blockNumberEndParm *big.Int, blockHashParm string) (*types.Transaction, error) {
	return _Validator.contract.Transact(opts, "setBackup", timestampParm, blockNumberStartParm, blockNumberEndParm, blockHashParm)
}

// SetBackup is a paid mutator transaction binding the contract method 0x3f458369.
//
// Solidity: function setBackup(timestampParm uint256, blockNumberStartParm uint256, blockNumberEndParm uint256, blockHashParm string) returns()
func (_Validator *ValidatorSession) SetBackup(timestampParm *big.Int, blockNumberStartParm *big.Int, blockNumberEndParm *big.Int, blockHashParm string) (*types.Transaction, error) {
	return _Validator.Contract.SetBackup(&_Validator.TransactOpts, timestampParm, blockNumberStartParm, blockNumberEndParm, blockHashParm)
}

// SetBackup is a paid mutator transaction binding the contract method 0x3f458369.
//
// Solidity: function setBackup(timestampParm uint256, blockNumberStartParm uint256, blockNumberEndParm uint256, blockHashParm string) returns()
func (_Validator *ValidatorTransactorSession) SetBackup(timestampParm *big.Int, blockNumberStartParm *big.Int, blockNumberEndParm *big.Int, blockHashParm string) (*types.Transaction, error) {
	return _Validator.Contract.SetBackup(&_Validator.TransactOpts, timestampParm, blockNumberStartParm, blockNumberEndParm, blockHashParm)
}

// ValidatorVerifiedBackupIterator is returned from FilterVerifiedBackup and is used to iterate over the raw logs and unpacked data for VerifiedBackup events raised by the Validator contract.
type ValidatorVerifiedBackupIterator struct {
	Event *ValidatorVerifiedBackup // Event containing the contract specifics and raw log

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
func (it *ValidatorVerifiedBackupIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ValidatorVerifiedBackup)
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
		it.Event = new(ValidatorVerifiedBackup)
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
func (it *ValidatorVerifiedBackupIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ValidatorVerifiedBackupIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ValidatorVerifiedBackup represents a VerifiedBackup event raised by the Validator contract.
type ValidatorVerifiedBackup struct {
	Timestamp        *big.Int
	BlockNumberStart *big.Int
	BlockNumberEnd   *big.Int
	BlockHash        string
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterVerifiedBackup is a free log retrieval operation binding the contract event 0xb43ce4d2c93d9e66cc96abea57b6ff99efce2ac3614305989fb45c4e492790c9.
//
// Solidity: e VerifiedBackup(timestamp uint256, blockNumberStart uint256, blockNumberEnd uint256, blockHash string)
func (_Validator *ValidatorFilterer) FilterVerifiedBackup(opts *bind.FilterOpts) (*ValidatorVerifiedBackupIterator, error) {

	logs, sub, err := _Validator.contract.FilterLogs(opts, "VerifiedBackup")
	if err != nil {
		return nil, err
	}
	return &ValidatorVerifiedBackupIterator{contract: _Validator.contract, event: "VerifiedBackup", logs: logs, sub: sub}, nil
}

// WatchVerifiedBackup is a free log subscription operation binding the contract event 0xb43ce4d2c93d9e66cc96abea57b6ff99efce2ac3614305989fb45c4e492790c9.
//
// Solidity: e VerifiedBackup(timestamp uint256, blockNumberStart uint256, blockNumberEnd uint256, blockHash string)
func (_Validator *ValidatorFilterer) WatchVerifiedBackup(opts *bind.WatchOpts, sink chan<- *ValidatorVerifiedBackup) (event.Subscription, error) {

	logs, sub, err := _Validator.contract.WatchLogs(opts, "VerifiedBackup")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ValidatorVerifiedBackup)
				if err := _Validator.contract.UnpackLog(event, "VerifiedBackup", log); err != nil {
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
