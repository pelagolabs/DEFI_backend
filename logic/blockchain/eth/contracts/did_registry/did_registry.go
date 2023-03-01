// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package did_registry

import (
	"errors"
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
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
	_ = abi.ConvertType
)

// DidRegistryMetaData contains all meta data concerning the DidRegistry contract.
var DidRegistryMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"did\",\"type\":\"string\"}],\"name\":\"getPublicKeyJwk\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"pubKeyJwk\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getType\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"did\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"pubkey\",\"type\":\"string\"}],\"name\":\"registerDid\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"type_\",\"type\":\"string\"}],\"name\":\"setType\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"did\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"pubkey\",\"type\":\"string\"}],\"name\":\"updateDid\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// DidRegistryABI is the input ABI used to generate the binding from.
// Deprecated: Use DidRegistryMetaData.ABI instead.
var DidRegistryABI = DidRegistryMetaData.ABI

// DidRegistry is an auto generated Go binding around an Ethereum contract.
type DidRegistry struct {
	DidRegistryCaller     // Read-only binding to the contract
	DidRegistryTransactor // Write-only binding to the contract
	DidRegistryFilterer   // Log filterer for contract events
}

// DidRegistryCaller is an auto generated read-only Go binding around an Ethereum contract.
type DidRegistryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DidRegistryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type DidRegistryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DidRegistryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type DidRegistryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DidRegistrySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type DidRegistrySession struct {
	Contract     *DidRegistry      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// DidRegistryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type DidRegistryCallerSession struct {
	Contract *DidRegistryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// DidRegistryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type DidRegistryTransactorSession struct {
	Contract     *DidRegistryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// DidRegistryRaw is an auto generated low-level Go binding around an Ethereum contract.
type DidRegistryRaw struct {
	Contract *DidRegistry // Generic contract binding to access the raw methods on
}

// DidRegistryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type DidRegistryCallerRaw struct {
	Contract *DidRegistryCaller // Generic read-only contract binding to access the raw methods on
}

// DidRegistryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type DidRegistryTransactorRaw struct {
	Contract *DidRegistryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewDidRegistry creates a new instance of DidRegistry, bound to a specific deployed contract.
func NewDidRegistry(address common.Address, backend bind.ContractBackend) (*DidRegistry, error) {
	contract, err := bindDidRegistry(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &DidRegistry{DidRegistryCaller: DidRegistryCaller{contract: contract}, DidRegistryTransactor: DidRegistryTransactor{contract: contract}, DidRegistryFilterer: DidRegistryFilterer{contract: contract}}, nil
}

// NewDidRegistryCaller creates a new read-only instance of DidRegistry, bound to a specific deployed contract.
func NewDidRegistryCaller(address common.Address, caller bind.ContractCaller) (*DidRegistryCaller, error) {
	contract, err := bindDidRegistry(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &DidRegistryCaller{contract: contract}, nil
}

// NewDidRegistryTransactor creates a new write-only instance of DidRegistry, bound to a specific deployed contract.
func NewDidRegistryTransactor(address common.Address, transactor bind.ContractTransactor) (*DidRegistryTransactor, error) {
	contract, err := bindDidRegistry(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &DidRegistryTransactor{contract: contract}, nil
}

// NewDidRegistryFilterer creates a new log filterer instance of DidRegistry, bound to a specific deployed contract.
func NewDidRegistryFilterer(address common.Address, filterer bind.ContractFilterer) (*DidRegistryFilterer, error) {
	contract, err := bindDidRegistry(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &DidRegistryFilterer{contract: contract}, nil
}

// bindDidRegistry binds a generic wrapper to an already deployed contract.
func bindDidRegistry(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := DidRegistryMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DidRegistry *DidRegistryRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _DidRegistry.Contract.DidRegistryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DidRegistry *DidRegistryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DidRegistry.Contract.DidRegistryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DidRegistry *DidRegistryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DidRegistry.Contract.DidRegistryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DidRegistry *DidRegistryCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _DidRegistry.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DidRegistry *DidRegistryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DidRegistry.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DidRegistry *DidRegistryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DidRegistry.Contract.contract.Transact(opts, method, params...)
}

// GetPublicKeyJwk is a free data retrieval call binding the contract method 0xa455db72.
//
// Solidity: function getPublicKeyJwk(string did) view returns(string pubKeyJwk)
func (_DidRegistry *DidRegistryCaller) GetPublicKeyJwk(opts *bind.CallOpts, did string) (string, error) {
	var out []interface{}
	err := _DidRegistry.contract.Call(opts, &out, "getPublicKeyJwk", did)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// GetPublicKeyJwk is a free data retrieval call binding the contract method 0xa455db72.
//
// Solidity: function getPublicKeyJwk(string did) view returns(string pubKeyJwk)
func (_DidRegistry *DidRegistrySession) GetPublicKeyJwk(did string) (string, error) {
	return _DidRegistry.Contract.GetPublicKeyJwk(&_DidRegistry.CallOpts, did)
}

// GetPublicKeyJwk is a free data retrieval call binding the contract method 0xa455db72.
//
// Solidity: function getPublicKeyJwk(string did) view returns(string pubKeyJwk)
func (_DidRegistry *DidRegistryCallerSession) GetPublicKeyJwk(did string) (string, error) {
	return _DidRegistry.Contract.GetPublicKeyJwk(&_DidRegistry.CallOpts, did)
}

// GetType is a free data retrieval call binding the contract method 0x15dae03e.
//
// Solidity: function getType() view returns(string)
func (_DidRegistry *DidRegistryCaller) GetType(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _DidRegistry.contract.Call(opts, &out, "getType")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// GetType is a free data retrieval call binding the contract method 0x15dae03e.
//
// Solidity: function getType() view returns(string)
func (_DidRegistry *DidRegistrySession) GetType() (string, error) {
	return _DidRegistry.Contract.GetType(&_DidRegistry.CallOpts)
}

// GetType is a free data retrieval call binding the contract method 0x15dae03e.
//
// Solidity: function getType() view returns(string)
func (_DidRegistry *DidRegistryCallerSession) GetType() (string, error) {
	return _DidRegistry.Contract.GetType(&_DidRegistry.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_DidRegistry *DidRegistryCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _DidRegistry.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_DidRegistry *DidRegistrySession) Owner() (common.Address, error) {
	return _DidRegistry.Contract.Owner(&_DidRegistry.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_DidRegistry *DidRegistryCallerSession) Owner() (common.Address, error) {
	return _DidRegistry.Contract.Owner(&_DidRegistry.CallOpts)
}

// RegisterDid is a paid mutator transaction binding the contract method 0x4b03e115.
//
// Solidity: function registerDid(string did, string pubkey) returns()
func (_DidRegistry *DidRegistryTransactor) RegisterDid(opts *bind.TransactOpts, did string, pubkey string) (*types.Transaction, error) {
	return _DidRegistry.contract.Transact(opts, "registerDid", did, pubkey)
}

// RegisterDid is a paid mutator transaction binding the contract method 0x4b03e115.
//
// Solidity: function registerDid(string did, string pubkey) returns()
func (_DidRegistry *DidRegistrySession) RegisterDid(did string, pubkey string) (*types.Transaction, error) {
	return _DidRegistry.Contract.RegisterDid(&_DidRegistry.TransactOpts, did, pubkey)
}

// RegisterDid is a paid mutator transaction binding the contract method 0x4b03e115.
//
// Solidity: function registerDid(string did, string pubkey) returns()
func (_DidRegistry *DidRegistryTransactorSession) RegisterDid(did string, pubkey string) (*types.Transaction, error) {
	return _DidRegistry.Contract.RegisterDid(&_DidRegistry.TransactOpts, did, pubkey)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_DidRegistry *DidRegistryTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DidRegistry.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_DidRegistry *DidRegistrySession) RenounceOwnership() (*types.Transaction, error) {
	return _DidRegistry.Contract.RenounceOwnership(&_DidRegistry.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_DidRegistry *DidRegistryTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _DidRegistry.Contract.RenounceOwnership(&_DidRegistry.TransactOpts)
}

// SetType is a paid mutator transaction binding the contract method 0x96282ba3.
//
// Solidity: function setType(string type_) returns()
func (_DidRegistry *DidRegistryTransactor) SetType(opts *bind.TransactOpts, type_ string) (*types.Transaction, error) {
	return _DidRegistry.contract.Transact(opts, "setType", type_)
}

// SetType is a paid mutator transaction binding the contract method 0x96282ba3.
//
// Solidity: function setType(string type_) returns()
func (_DidRegistry *DidRegistrySession) SetType(type_ string) (*types.Transaction, error) {
	return _DidRegistry.Contract.SetType(&_DidRegistry.TransactOpts, type_)
}

// SetType is a paid mutator transaction binding the contract method 0x96282ba3.
//
// Solidity: function setType(string type_) returns()
func (_DidRegistry *DidRegistryTransactorSession) SetType(type_ string) (*types.Transaction, error) {
	return _DidRegistry.Contract.SetType(&_DidRegistry.TransactOpts, type_)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_DidRegistry *DidRegistryTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _DidRegistry.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_DidRegistry *DidRegistrySession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _DidRegistry.Contract.TransferOwnership(&_DidRegistry.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_DidRegistry *DidRegistryTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _DidRegistry.Contract.TransferOwnership(&_DidRegistry.TransactOpts, newOwner)
}

// UpdateDid is a paid mutator transaction binding the contract method 0xdafffa7d.
//
// Solidity: function updateDid(string did, string pubkey) returns()
func (_DidRegistry *DidRegistryTransactor) UpdateDid(opts *bind.TransactOpts, did string, pubkey string) (*types.Transaction, error) {
	return _DidRegistry.contract.Transact(opts, "updateDid", did, pubkey)
}

// UpdateDid is a paid mutator transaction binding the contract method 0xdafffa7d.
//
// Solidity: function updateDid(string did, string pubkey) returns()
func (_DidRegistry *DidRegistrySession) UpdateDid(did string, pubkey string) (*types.Transaction, error) {
	return _DidRegistry.Contract.UpdateDid(&_DidRegistry.TransactOpts, did, pubkey)
}

// UpdateDid is a paid mutator transaction binding the contract method 0xdafffa7d.
//
// Solidity: function updateDid(string did, string pubkey) returns()
func (_DidRegistry *DidRegistryTransactorSession) UpdateDid(did string, pubkey string) (*types.Transaction, error) {
	return _DidRegistry.Contract.UpdateDid(&_DidRegistry.TransactOpts, did, pubkey)
}

// DidRegistryOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the DidRegistry contract.
type DidRegistryOwnershipTransferredIterator struct {
	Event *DidRegistryOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *DidRegistryOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DidRegistryOwnershipTransferred)
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
		it.Event = new(DidRegistryOwnershipTransferred)
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
func (it *DidRegistryOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DidRegistryOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DidRegistryOwnershipTransferred represents a OwnershipTransferred event raised by the DidRegistry contract.
type DidRegistryOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_DidRegistry *DidRegistryFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*DidRegistryOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _DidRegistry.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &DidRegistryOwnershipTransferredIterator{contract: _DidRegistry.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_DidRegistry *DidRegistryFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *DidRegistryOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _DidRegistry.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DidRegistryOwnershipTransferred)
				if err := _DidRegistry.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_DidRegistry *DidRegistryFilterer) ParseOwnershipTransferred(log types.Log) (*DidRegistryOwnershipTransferred, error) {
	event := new(DidRegistryOwnershipTransferred)
	if err := _DidRegistry.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
