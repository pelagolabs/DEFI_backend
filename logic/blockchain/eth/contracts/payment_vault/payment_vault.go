// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package payment_vault

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

// PaymentVaultMetaData contains all meta data concerning the PaymentVault contract.
var PaymentVaultMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_who\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"ReceivedEth\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_who\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_tokenAddress\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"Withdrawn\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_who\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"WithdrawnEth\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"token_\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipient_\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount_\",\"type\":\"uint256\"}],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipient_\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount_\",\"type\":\"uint256\"}],\"name\":\"withdrawEth\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
}

// PaymentVaultABI is the input ABI used to generate the binding from.
// Deprecated: Use PaymentVaultMetaData.ABI instead.
var PaymentVaultABI = PaymentVaultMetaData.ABI

// PaymentVault is an auto generated Go binding around an Ethereum contract.
type PaymentVault struct {
	PaymentVaultCaller     // Read-only binding to the contract
	PaymentVaultTransactor // Write-only binding to the contract
	PaymentVaultFilterer   // Log filterer for contract events
}

// PaymentVaultCaller is an auto generated read-only Go binding around an Ethereum contract.
type PaymentVaultCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PaymentVaultTransactor is an auto generated write-only Go binding around an Ethereum contract.
type PaymentVaultTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PaymentVaultFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type PaymentVaultFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PaymentVaultSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PaymentVaultSession struct {
	Contract     *PaymentVault     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// PaymentVaultCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PaymentVaultCallerSession struct {
	Contract *PaymentVaultCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// PaymentVaultTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PaymentVaultTransactorSession struct {
	Contract     *PaymentVaultTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// PaymentVaultRaw is an auto generated low-level Go binding around an Ethereum contract.
type PaymentVaultRaw struct {
	Contract *PaymentVault // Generic contract binding to access the raw methods on
}

// PaymentVaultCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PaymentVaultCallerRaw struct {
	Contract *PaymentVaultCaller // Generic read-only contract binding to access the raw methods on
}

// PaymentVaultTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PaymentVaultTransactorRaw struct {
	Contract *PaymentVaultTransactor // Generic write-only contract binding to access the raw methods on
}

// NewPaymentVault creates a new instance of PaymentVault, bound to a specific deployed contract.
func NewPaymentVault(address common.Address, backend bind.ContractBackend) (*PaymentVault, error) {
	contract, err := bindPaymentVault(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &PaymentVault{PaymentVaultCaller: PaymentVaultCaller{contract: contract}, PaymentVaultTransactor: PaymentVaultTransactor{contract: contract}, PaymentVaultFilterer: PaymentVaultFilterer{contract: contract}}, nil
}

// NewPaymentVaultCaller creates a new read-only instance of PaymentVault, bound to a specific deployed contract.
func NewPaymentVaultCaller(address common.Address, caller bind.ContractCaller) (*PaymentVaultCaller, error) {
	contract, err := bindPaymentVault(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &PaymentVaultCaller{contract: contract}, nil
}

// NewPaymentVaultTransactor creates a new write-only instance of PaymentVault, bound to a specific deployed contract.
func NewPaymentVaultTransactor(address common.Address, transactor bind.ContractTransactor) (*PaymentVaultTransactor, error) {
	contract, err := bindPaymentVault(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &PaymentVaultTransactor{contract: contract}, nil
}

// NewPaymentVaultFilterer creates a new log filterer instance of PaymentVault, bound to a specific deployed contract.
func NewPaymentVaultFilterer(address common.Address, filterer bind.ContractFilterer) (*PaymentVaultFilterer, error) {
	contract, err := bindPaymentVault(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &PaymentVaultFilterer{contract: contract}, nil
}

// bindPaymentVault binds a generic wrapper to an already deployed contract.
func bindPaymentVault(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := PaymentVaultMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PaymentVault *PaymentVaultRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _PaymentVault.Contract.PaymentVaultCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PaymentVault *PaymentVaultRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PaymentVault.Contract.PaymentVaultTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PaymentVault *PaymentVaultRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PaymentVault.Contract.PaymentVaultTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PaymentVault *PaymentVaultCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _PaymentVault.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PaymentVault *PaymentVaultTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PaymentVault.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PaymentVault *PaymentVaultTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PaymentVault.Contract.contract.Transact(opts, method, params...)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_PaymentVault *PaymentVaultCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _PaymentVault.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_PaymentVault *PaymentVaultSession) Owner() (common.Address, error) {
	return _PaymentVault.Contract.Owner(&_PaymentVault.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_PaymentVault *PaymentVaultCallerSession) Owner() (common.Address, error) {
	return _PaymentVault.Contract.Owner(&_PaymentVault.CallOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_PaymentVault *PaymentVaultTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PaymentVault.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_PaymentVault *PaymentVaultSession) RenounceOwnership() (*types.Transaction, error) {
	return _PaymentVault.Contract.RenounceOwnership(&_PaymentVault.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_PaymentVault *PaymentVaultTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _PaymentVault.Contract.RenounceOwnership(&_PaymentVault.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_PaymentVault *PaymentVaultTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _PaymentVault.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_PaymentVault *PaymentVaultSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _PaymentVault.Contract.TransferOwnership(&_PaymentVault.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_PaymentVault *PaymentVaultTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _PaymentVault.Contract.TransferOwnership(&_PaymentVault.TransactOpts, newOwner)
}

// Withdraw is a paid mutator transaction binding the contract method 0xd9caed12.
//
// Solidity: function withdraw(address token_, address recipient_, uint256 amount_) returns()
func (_PaymentVault *PaymentVaultTransactor) Withdraw(opts *bind.TransactOpts, token_ common.Address, recipient_ common.Address, amount_ *big.Int) (*types.Transaction, error) {
	return _PaymentVault.contract.Transact(opts, "withdraw", token_, recipient_, amount_)
}

// Withdraw is a paid mutator transaction binding the contract method 0xd9caed12.
//
// Solidity: function withdraw(address token_, address recipient_, uint256 amount_) returns()
func (_PaymentVault *PaymentVaultSession) Withdraw(token_ common.Address, recipient_ common.Address, amount_ *big.Int) (*types.Transaction, error) {
	return _PaymentVault.Contract.Withdraw(&_PaymentVault.TransactOpts, token_, recipient_, amount_)
}

// Withdraw is a paid mutator transaction binding the contract method 0xd9caed12.
//
// Solidity: function withdraw(address token_, address recipient_, uint256 amount_) returns()
func (_PaymentVault *PaymentVaultTransactorSession) Withdraw(token_ common.Address, recipient_ common.Address, amount_ *big.Int) (*types.Transaction, error) {
	return _PaymentVault.Contract.Withdraw(&_PaymentVault.TransactOpts, token_, recipient_, amount_)
}

// WithdrawEth is a paid mutator transaction binding the contract method 0x1b9a91a4.
//
// Solidity: function withdrawEth(address recipient_, uint256 amount_) returns()
func (_PaymentVault *PaymentVaultTransactor) WithdrawEth(opts *bind.TransactOpts, recipient_ common.Address, amount_ *big.Int) (*types.Transaction, error) {
	return _PaymentVault.contract.Transact(opts, "withdrawEth", recipient_, amount_)
}

// WithdrawEth is a paid mutator transaction binding the contract method 0x1b9a91a4.
//
// Solidity: function withdrawEth(address recipient_, uint256 amount_) returns()
func (_PaymentVault *PaymentVaultSession) WithdrawEth(recipient_ common.Address, amount_ *big.Int) (*types.Transaction, error) {
	return _PaymentVault.Contract.WithdrawEth(&_PaymentVault.TransactOpts, recipient_, amount_)
}

// WithdrawEth is a paid mutator transaction binding the contract method 0x1b9a91a4.
//
// Solidity: function withdrawEth(address recipient_, uint256 amount_) returns()
func (_PaymentVault *PaymentVaultTransactorSession) WithdrawEth(recipient_ common.Address, amount_ *big.Int) (*types.Transaction, error) {
	return _PaymentVault.Contract.WithdrawEth(&_PaymentVault.TransactOpts, recipient_, amount_)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_PaymentVault *PaymentVaultTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PaymentVault.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_PaymentVault *PaymentVaultSession) Receive() (*types.Transaction, error) {
	return _PaymentVault.Contract.Receive(&_PaymentVault.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_PaymentVault *PaymentVaultTransactorSession) Receive() (*types.Transaction, error) {
	return _PaymentVault.Contract.Receive(&_PaymentVault.TransactOpts)
}

// PaymentVaultOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the PaymentVault contract.
type PaymentVaultOwnershipTransferredIterator struct {
	Event *PaymentVaultOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *PaymentVaultOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PaymentVaultOwnershipTransferred)
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
		it.Event = new(PaymentVaultOwnershipTransferred)
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
func (it *PaymentVaultOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PaymentVaultOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PaymentVaultOwnershipTransferred represents a OwnershipTransferred event raised by the PaymentVault contract.
type PaymentVaultOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_PaymentVault *PaymentVaultFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*PaymentVaultOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _PaymentVault.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &PaymentVaultOwnershipTransferredIterator{contract: _PaymentVault.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_PaymentVault *PaymentVaultFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *PaymentVaultOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _PaymentVault.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PaymentVaultOwnershipTransferred)
				if err := _PaymentVault.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_PaymentVault *PaymentVaultFilterer) ParseOwnershipTransferred(log types.Log) (*PaymentVaultOwnershipTransferred, error) {
	event := new(PaymentVaultOwnershipTransferred)
	if err := _PaymentVault.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PaymentVaultReceivedEthIterator is returned from FilterReceivedEth and is used to iterate over the raw logs and unpacked data for ReceivedEth events raised by the PaymentVault contract.
type PaymentVaultReceivedEthIterator struct {
	Event *PaymentVaultReceivedEth // Event containing the contract specifics and raw log

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
func (it *PaymentVaultReceivedEthIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PaymentVaultReceivedEth)
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
		it.Event = new(PaymentVaultReceivedEth)
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
func (it *PaymentVaultReceivedEthIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PaymentVaultReceivedEthIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PaymentVaultReceivedEth represents a ReceivedEth event raised by the PaymentVault contract.
type PaymentVaultReceivedEth struct {
	Who    common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterReceivedEth is a free log retrieval operation binding the contract event 0x52a6cdf67c40ce333b3d846e4e143db87f71dd7935612a4cafcf6ba76047ca1f.
//
// Solidity: event ReceivedEth(address indexed _who, uint256 indexed _amount)
func (_PaymentVault *PaymentVaultFilterer) FilterReceivedEth(opts *bind.FilterOpts, _who []common.Address, _amount []*big.Int) (*PaymentVaultReceivedEthIterator, error) {

	var _whoRule []interface{}
	for _, _whoItem := range _who {
		_whoRule = append(_whoRule, _whoItem)
	}
	var _amountRule []interface{}
	for _, _amountItem := range _amount {
		_amountRule = append(_amountRule, _amountItem)
	}

	logs, sub, err := _PaymentVault.contract.FilterLogs(opts, "ReceivedEth", _whoRule, _amountRule)
	if err != nil {
		return nil, err
	}
	return &PaymentVaultReceivedEthIterator{contract: _PaymentVault.contract, event: "ReceivedEth", logs: logs, sub: sub}, nil
}

// WatchReceivedEth is a free log subscription operation binding the contract event 0x52a6cdf67c40ce333b3d846e4e143db87f71dd7935612a4cafcf6ba76047ca1f.
//
// Solidity: event ReceivedEth(address indexed _who, uint256 indexed _amount)
func (_PaymentVault *PaymentVaultFilterer) WatchReceivedEth(opts *bind.WatchOpts, sink chan<- *PaymentVaultReceivedEth, _who []common.Address, _amount []*big.Int) (event.Subscription, error) {

	var _whoRule []interface{}
	for _, _whoItem := range _who {
		_whoRule = append(_whoRule, _whoItem)
	}
	var _amountRule []interface{}
	for _, _amountItem := range _amount {
		_amountRule = append(_amountRule, _amountItem)
	}

	logs, sub, err := _PaymentVault.contract.WatchLogs(opts, "ReceivedEth", _whoRule, _amountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PaymentVaultReceivedEth)
				if err := _PaymentVault.contract.UnpackLog(event, "ReceivedEth", log); err != nil {
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

// ParseReceivedEth is a log parse operation binding the contract event 0x52a6cdf67c40ce333b3d846e4e143db87f71dd7935612a4cafcf6ba76047ca1f.
//
// Solidity: event ReceivedEth(address indexed _who, uint256 indexed _amount)
func (_PaymentVault *PaymentVaultFilterer) ParseReceivedEth(log types.Log) (*PaymentVaultReceivedEth, error) {
	event := new(PaymentVaultReceivedEth)
	if err := _PaymentVault.contract.UnpackLog(event, "ReceivedEth", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PaymentVaultWithdrawnIterator is returned from FilterWithdrawn and is used to iterate over the raw logs and unpacked data for Withdrawn events raised by the PaymentVault contract.
type PaymentVaultWithdrawnIterator struct {
	Event *PaymentVaultWithdrawn // Event containing the contract specifics and raw log

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
func (it *PaymentVaultWithdrawnIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PaymentVaultWithdrawn)
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
		it.Event = new(PaymentVaultWithdrawn)
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
func (it *PaymentVaultWithdrawnIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PaymentVaultWithdrawnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PaymentVaultWithdrawn represents a Withdrawn event raised by the PaymentVault contract.
type PaymentVaultWithdrawn struct {
	Who          common.Address
	TokenAddress common.Address
	Amount       *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterWithdrawn is a free log retrieval operation binding the contract event 0xd1c19fbcd4551a5edfb66d43d2e337c04837afda3482b42bdf569a8fccdae5fb.
//
// Solidity: event Withdrawn(address indexed _who, address indexed _tokenAddress, uint256 indexed _amount)
func (_PaymentVault *PaymentVaultFilterer) FilterWithdrawn(opts *bind.FilterOpts, _who []common.Address, _tokenAddress []common.Address, _amount []*big.Int) (*PaymentVaultWithdrawnIterator, error) {

	var _whoRule []interface{}
	for _, _whoItem := range _who {
		_whoRule = append(_whoRule, _whoItem)
	}
	var _tokenAddressRule []interface{}
	for _, _tokenAddressItem := range _tokenAddress {
		_tokenAddressRule = append(_tokenAddressRule, _tokenAddressItem)
	}
	var _amountRule []interface{}
	for _, _amountItem := range _amount {
		_amountRule = append(_amountRule, _amountItem)
	}

	logs, sub, err := _PaymentVault.contract.FilterLogs(opts, "Withdrawn", _whoRule, _tokenAddressRule, _amountRule)
	if err != nil {
		return nil, err
	}
	return &PaymentVaultWithdrawnIterator{contract: _PaymentVault.contract, event: "Withdrawn", logs: logs, sub: sub}, nil
}

// WatchWithdrawn is a free log subscription operation binding the contract event 0xd1c19fbcd4551a5edfb66d43d2e337c04837afda3482b42bdf569a8fccdae5fb.
//
// Solidity: event Withdrawn(address indexed _who, address indexed _tokenAddress, uint256 indexed _amount)
func (_PaymentVault *PaymentVaultFilterer) WatchWithdrawn(opts *bind.WatchOpts, sink chan<- *PaymentVaultWithdrawn, _who []common.Address, _tokenAddress []common.Address, _amount []*big.Int) (event.Subscription, error) {

	var _whoRule []interface{}
	for _, _whoItem := range _who {
		_whoRule = append(_whoRule, _whoItem)
	}
	var _tokenAddressRule []interface{}
	for _, _tokenAddressItem := range _tokenAddress {
		_tokenAddressRule = append(_tokenAddressRule, _tokenAddressItem)
	}
	var _amountRule []interface{}
	for _, _amountItem := range _amount {
		_amountRule = append(_amountRule, _amountItem)
	}

	logs, sub, err := _PaymentVault.contract.WatchLogs(opts, "Withdrawn", _whoRule, _tokenAddressRule, _amountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PaymentVaultWithdrawn)
				if err := _PaymentVault.contract.UnpackLog(event, "Withdrawn", log); err != nil {
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

// ParseWithdrawn is a log parse operation binding the contract event 0xd1c19fbcd4551a5edfb66d43d2e337c04837afda3482b42bdf569a8fccdae5fb.
//
// Solidity: event Withdrawn(address indexed _who, address indexed _tokenAddress, uint256 indexed _amount)
func (_PaymentVault *PaymentVaultFilterer) ParseWithdrawn(log types.Log) (*PaymentVaultWithdrawn, error) {
	event := new(PaymentVaultWithdrawn)
	if err := _PaymentVault.contract.UnpackLog(event, "Withdrawn", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PaymentVaultWithdrawnEthIterator is returned from FilterWithdrawnEth and is used to iterate over the raw logs and unpacked data for WithdrawnEth events raised by the PaymentVault contract.
type PaymentVaultWithdrawnEthIterator struct {
	Event *PaymentVaultWithdrawnEth // Event containing the contract specifics and raw log

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
func (it *PaymentVaultWithdrawnEthIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PaymentVaultWithdrawnEth)
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
		it.Event = new(PaymentVaultWithdrawnEth)
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
func (it *PaymentVaultWithdrawnEthIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PaymentVaultWithdrawnEthIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PaymentVaultWithdrawnEth represents a WithdrawnEth event raised by the PaymentVault contract.
type PaymentVaultWithdrawnEth struct {
	Who    common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterWithdrawnEth is a free log retrieval operation binding the contract event 0xa8841823f7591414a80deddb48af504373043ab5959ceddf527c2215d62cb29b.
//
// Solidity: event WithdrawnEth(address indexed _who, uint256 indexed _amount)
func (_PaymentVault *PaymentVaultFilterer) FilterWithdrawnEth(opts *bind.FilterOpts, _who []common.Address, _amount []*big.Int) (*PaymentVaultWithdrawnEthIterator, error) {

	var _whoRule []interface{}
	for _, _whoItem := range _who {
		_whoRule = append(_whoRule, _whoItem)
	}
	var _amountRule []interface{}
	for _, _amountItem := range _amount {
		_amountRule = append(_amountRule, _amountItem)
	}

	logs, sub, err := _PaymentVault.contract.FilterLogs(opts, "WithdrawnEth", _whoRule, _amountRule)
	if err != nil {
		return nil, err
	}
	return &PaymentVaultWithdrawnEthIterator{contract: _PaymentVault.contract, event: "WithdrawnEth", logs: logs, sub: sub}, nil
}

// WatchWithdrawnEth is a free log subscription operation binding the contract event 0xa8841823f7591414a80deddb48af504373043ab5959ceddf527c2215d62cb29b.
//
// Solidity: event WithdrawnEth(address indexed _who, uint256 indexed _amount)
func (_PaymentVault *PaymentVaultFilterer) WatchWithdrawnEth(opts *bind.WatchOpts, sink chan<- *PaymentVaultWithdrawnEth, _who []common.Address, _amount []*big.Int) (event.Subscription, error) {

	var _whoRule []interface{}
	for _, _whoItem := range _who {
		_whoRule = append(_whoRule, _whoItem)
	}
	var _amountRule []interface{}
	for _, _amountItem := range _amount {
		_amountRule = append(_amountRule, _amountItem)
	}

	logs, sub, err := _PaymentVault.contract.WatchLogs(opts, "WithdrawnEth", _whoRule, _amountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PaymentVaultWithdrawnEth)
				if err := _PaymentVault.contract.UnpackLog(event, "WithdrawnEth", log); err != nil {
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

// ParseWithdrawnEth is a log parse operation binding the contract event 0xa8841823f7591414a80deddb48af504373043ab5959ceddf527c2215d62cb29b.
//
// Solidity: event WithdrawnEth(address indexed _who, uint256 indexed _amount)
func (_PaymentVault *PaymentVaultFilterer) ParseWithdrawnEth(log types.Log) (*PaymentVaultWithdrawnEth, error) {
	event := new(PaymentVaultWithdrawnEth)
	if err := _PaymentVault.contract.UnpackLog(event, "WithdrawnEth", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
