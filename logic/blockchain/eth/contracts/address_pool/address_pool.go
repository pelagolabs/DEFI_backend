// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package address_pool

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

// AddressPoolMetaData contains all meta data concerning the AddressPool contract.
var AddressPoolMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"chain\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"vaultAddress\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"}],\"name\":\"AddedNewVault\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"string\",\"name\":\"requester\",\"type\":\"string\"},{\"indexed\":true,\"internalType\":\"string\",\"name\":\"token\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"txid\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"vault\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"requestHash\",\"type\":\"bytes32\"}],\"name\":\"PaymentRequestAdd\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"chain\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"numberOfvaults\",\"type\":\"uint256\"}],\"name\":\"ResetVaults\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"chain\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"vaultAddress\",\"type\":\"string\"}],\"name\":\"UsingVault\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"chain\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"vaultAddress\",\"type\":\"string\"}],\"name\":\"VaultReleased\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"chain_\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"vaultAddress_\",\"type\":\"string\"}],\"name\":\"addNewVault\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"chain_\",\"type\":\"string\"}],\"name\":\"getAvailableVault\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"vaultAddress\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"chain_\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"vaultAddress_\",\"type\":\"string\"}],\"name\":\"getVaultStatus\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"vaultStatus\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"paymentRequestNonce\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"name\":\"paymentRequestTxid\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"paymentRequests\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"requester\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"token\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"txid\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"vault\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"},{\"internalType\":\"enumAddressPool.RequestStatus\",\"name\":\"status\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"chain_\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"vaultAddress_\",\"type\":\"string\"}],\"name\":\"releaseVault\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"chain_\",\"type\":\"string\"},{\"internalType\":\"string[]\",\"name\":\"vaultAddresses_\",\"type\":\"string[]\"}],\"name\":\"resetVaults\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"supportedChains\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index_\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"chain_\",\"type\":\"string\"}],\"name\":\"useAvailableVault\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"vaultAddress\",\"type\":\"string\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index_\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"chain_\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"vaultAddress_\",\"type\":\"string\"}],\"name\":\"useVault\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"name\":\"vaultInfoNonce\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"name\":\"vaultTokenBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"vaults\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"vaultAddress\",\"type\":\"string\"},{\"internalType\":\"enumVaultStatus\",\"name\":\"status\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"withdrawRequestNonce\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"name\":\"withdrawRequestTxid\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"withdrawRequests\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"requester\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"token\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"txid\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"vault\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"},{\"internalType\":\"enumAddressPool.RequestStatus\",\"name\":\"status\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// AddressPoolABI is the input ABI used to generate the binding from.
// Deprecated: Use AddressPoolMetaData.ABI instead.
var AddressPoolABI = AddressPoolMetaData.ABI

// AddressPool is an auto generated Go binding around an Ethereum contract.
type AddressPool struct {
	AddressPoolCaller     // Read-only binding to the contract
	AddressPoolTransactor // Write-only binding to the contract
	AddressPoolFilterer   // Log filterer for contract events
}

// AddressPoolCaller is an auto generated read-only Go binding around an Ethereum contract.
type AddressPoolCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AddressPoolTransactor is an auto generated write-only Go binding around an Ethereum contract.
type AddressPoolTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AddressPoolFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type AddressPoolFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AddressPoolSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type AddressPoolSession struct {
	Contract     *AddressPool      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// AddressPoolCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type AddressPoolCallerSession struct {
	Contract *AddressPoolCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// AddressPoolTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type AddressPoolTransactorSession struct {
	Contract     *AddressPoolTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// AddressPoolRaw is an auto generated low-level Go binding around an Ethereum contract.
type AddressPoolRaw struct {
	Contract *AddressPool // Generic contract binding to access the raw methods on
}

// AddressPoolCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type AddressPoolCallerRaw struct {
	Contract *AddressPoolCaller // Generic read-only contract binding to access the raw methods on
}

// AddressPoolTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type AddressPoolTransactorRaw struct {
	Contract *AddressPoolTransactor // Generic write-only contract binding to access the raw methods on
}

// NewAddressPool creates a new instance of AddressPool, bound to a specific deployed contract.
func NewAddressPool(address common.Address, backend bind.ContractBackend) (*AddressPool, error) {
	contract, err := bindAddressPool(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &AddressPool{AddressPoolCaller: AddressPoolCaller{contract: contract}, AddressPoolTransactor: AddressPoolTransactor{contract: contract}, AddressPoolFilterer: AddressPoolFilterer{contract: contract}}, nil
}

// NewAddressPoolCaller creates a new read-only instance of AddressPool, bound to a specific deployed contract.
func NewAddressPoolCaller(address common.Address, caller bind.ContractCaller) (*AddressPoolCaller, error) {
	contract, err := bindAddressPool(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &AddressPoolCaller{contract: contract}, nil
}

// NewAddressPoolTransactor creates a new write-only instance of AddressPool, bound to a specific deployed contract.
func NewAddressPoolTransactor(address common.Address, transactor bind.ContractTransactor) (*AddressPoolTransactor, error) {
	contract, err := bindAddressPool(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &AddressPoolTransactor{contract: contract}, nil
}

// NewAddressPoolFilterer creates a new log filterer instance of AddressPool, bound to a specific deployed contract.
func NewAddressPoolFilterer(address common.Address, filterer bind.ContractFilterer) (*AddressPoolFilterer, error) {
	contract, err := bindAddressPool(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &AddressPoolFilterer{contract: contract}, nil
}

// bindAddressPool binds a generic wrapper to an already deployed contract.
func bindAddressPool(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := AddressPoolMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AddressPool *AddressPoolRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _AddressPool.Contract.AddressPoolCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AddressPool *AddressPoolRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AddressPool.Contract.AddressPoolTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AddressPool *AddressPoolRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AddressPool.Contract.AddressPoolTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AddressPool *AddressPoolCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _AddressPool.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AddressPool *AddressPoolTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AddressPool.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AddressPool *AddressPoolTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AddressPool.Contract.contract.Transact(opts, method, params...)
}

// GetAvailableVault is a free data retrieval call binding the contract method 0x4b5220bb.
//
// Solidity: function getAvailableVault(string chain_) view returns(string vaultAddress)
func (_AddressPool *AddressPoolCaller) GetAvailableVault(opts *bind.CallOpts, chain_ string) (string, error) {
	var out []interface{}
	err := _AddressPool.contract.Call(opts, &out, "getAvailableVault", chain_)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// GetAvailableVault is a free data retrieval call binding the contract method 0x4b5220bb.
//
// Solidity: function getAvailableVault(string chain_) view returns(string vaultAddress)
func (_AddressPool *AddressPoolSession) GetAvailableVault(chain_ string) (string, error) {
	return _AddressPool.Contract.GetAvailableVault(&_AddressPool.CallOpts, chain_)
}

// GetAvailableVault is a free data retrieval call binding the contract method 0x4b5220bb.
//
// Solidity: function getAvailableVault(string chain_) view returns(string vaultAddress)
func (_AddressPool *AddressPoolCallerSession) GetAvailableVault(chain_ string) (string, error) {
	return _AddressPool.Contract.GetAvailableVault(&_AddressPool.CallOpts, chain_)
}

// GetVaultStatus is a free data retrieval call binding the contract method 0x20022e40.
//
// Solidity: function getVaultStatus(string chain_, string vaultAddress_) view returns(string vaultStatus)
func (_AddressPool *AddressPoolCaller) GetVaultStatus(opts *bind.CallOpts, chain_ string, vaultAddress_ string) (string, error) {
	var out []interface{}
	err := _AddressPool.contract.Call(opts, &out, "getVaultStatus", chain_, vaultAddress_)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// GetVaultStatus is a free data retrieval call binding the contract method 0x20022e40.
//
// Solidity: function getVaultStatus(string chain_, string vaultAddress_) view returns(string vaultStatus)
func (_AddressPool *AddressPoolSession) GetVaultStatus(chain_ string, vaultAddress_ string) (string, error) {
	return _AddressPool.Contract.GetVaultStatus(&_AddressPool.CallOpts, chain_, vaultAddress_)
}

// GetVaultStatus is a free data retrieval call binding the contract method 0x20022e40.
//
// Solidity: function getVaultStatus(string chain_, string vaultAddress_) view returns(string vaultStatus)
func (_AddressPool *AddressPoolCallerSession) GetVaultStatus(chain_ string, vaultAddress_ string) (string, error) {
	return _AddressPool.Contract.GetVaultStatus(&_AddressPool.CallOpts, chain_, vaultAddress_)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_AddressPool *AddressPoolCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _AddressPool.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_AddressPool *AddressPoolSession) Owner() (common.Address, error) {
	return _AddressPool.Contract.Owner(&_AddressPool.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_AddressPool *AddressPoolCallerSession) Owner() (common.Address, error) {
	return _AddressPool.Contract.Owner(&_AddressPool.CallOpts)
}

// PaymentRequestNonce is a free data retrieval call binding the contract method 0xe8be1a43.
//
// Solidity: function paymentRequestNonce(bytes32 ) view returns(uint256)
func (_AddressPool *AddressPoolCaller) PaymentRequestNonce(opts *bind.CallOpts, arg0 [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _AddressPool.contract.Call(opts, &out, "paymentRequestNonce", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PaymentRequestNonce is a free data retrieval call binding the contract method 0xe8be1a43.
//
// Solidity: function paymentRequestNonce(bytes32 ) view returns(uint256)
func (_AddressPool *AddressPoolSession) PaymentRequestNonce(arg0 [32]byte) (*big.Int, error) {
	return _AddressPool.Contract.PaymentRequestNonce(&_AddressPool.CallOpts, arg0)
}

// PaymentRequestNonce is a free data retrieval call binding the contract method 0xe8be1a43.
//
// Solidity: function paymentRequestNonce(bytes32 ) view returns(uint256)
func (_AddressPool *AddressPoolCallerSession) PaymentRequestNonce(arg0 [32]byte) (*big.Int, error) {
	return _AddressPool.Contract.PaymentRequestNonce(&_AddressPool.CallOpts, arg0)
}

// PaymentRequestTxid is a free data retrieval call binding the contract method 0x2b26229c.
//
// Solidity: function paymentRequestTxid(string ) view returns(bytes32)
func (_AddressPool *AddressPoolCaller) PaymentRequestTxid(opts *bind.CallOpts, arg0 string) ([32]byte, error) {
	var out []interface{}
	err := _AddressPool.contract.Call(opts, &out, "paymentRequestTxid", arg0)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// PaymentRequestTxid is a free data retrieval call binding the contract method 0x2b26229c.
//
// Solidity: function paymentRequestTxid(string ) view returns(bytes32)
func (_AddressPool *AddressPoolSession) PaymentRequestTxid(arg0 string) ([32]byte, error) {
	return _AddressPool.Contract.PaymentRequestTxid(&_AddressPool.CallOpts, arg0)
}

// PaymentRequestTxid is a free data retrieval call binding the contract method 0x2b26229c.
//
// Solidity: function paymentRequestTxid(string ) view returns(bytes32)
func (_AddressPool *AddressPoolCallerSession) PaymentRequestTxid(arg0 string) ([32]byte, error) {
	return _AddressPool.Contract.PaymentRequestTxid(&_AddressPool.CallOpts, arg0)
}

// PaymentRequests is a free data retrieval call binding the contract method 0xd4358800.
//
// Solidity: function paymentRequests(uint256 ) view returns(string requester, string token, uint256 amount, string txid, string vault, uint256 nonce, uint256 timestamp, uint8 status)
func (_AddressPool *AddressPoolCaller) PaymentRequests(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Requester string
	Token     string
	Amount    *big.Int
	Txid      string
	Vault     string
	Nonce     *big.Int
	Timestamp *big.Int
	Status    uint8
}, error) {
	var out []interface{}
	err := _AddressPool.contract.Call(opts, &out, "paymentRequests", arg0)

	outstruct := new(struct {
		Requester string
		Token     string
		Amount    *big.Int
		Txid      string
		Vault     string
		Nonce     *big.Int
		Timestamp *big.Int
		Status    uint8
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Requester = *abi.ConvertType(out[0], new(string)).(*string)
	outstruct.Token = *abi.ConvertType(out[1], new(string)).(*string)
	outstruct.Amount = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.Txid = *abi.ConvertType(out[3], new(string)).(*string)
	outstruct.Vault = *abi.ConvertType(out[4], new(string)).(*string)
	outstruct.Nonce = *abi.ConvertType(out[5], new(*big.Int)).(**big.Int)
	outstruct.Timestamp = *abi.ConvertType(out[6], new(*big.Int)).(**big.Int)
	outstruct.Status = *abi.ConvertType(out[7], new(uint8)).(*uint8)

	return *outstruct, err

}

// PaymentRequests is a free data retrieval call binding the contract method 0xd4358800.
//
// Solidity: function paymentRequests(uint256 ) view returns(string requester, string token, uint256 amount, string txid, string vault, uint256 nonce, uint256 timestamp, uint8 status)
func (_AddressPool *AddressPoolSession) PaymentRequests(arg0 *big.Int) (struct {
	Requester string
	Token     string
	Amount    *big.Int
	Txid      string
	Vault     string
	Nonce     *big.Int
	Timestamp *big.Int
	Status    uint8
}, error) {
	return _AddressPool.Contract.PaymentRequests(&_AddressPool.CallOpts, arg0)
}

// PaymentRequests is a free data retrieval call binding the contract method 0xd4358800.
//
// Solidity: function paymentRequests(uint256 ) view returns(string requester, string token, uint256 amount, string txid, string vault, uint256 nonce, uint256 timestamp, uint8 status)
func (_AddressPool *AddressPoolCallerSession) PaymentRequests(arg0 *big.Int) (struct {
	Requester string
	Token     string
	Amount    *big.Int
	Txid      string
	Vault     string
	Nonce     *big.Int
	Timestamp *big.Int
	Status    uint8
}, error) {
	return _AddressPool.Contract.PaymentRequests(&_AddressPool.CallOpts, arg0)
}

// SupportedChains is a free data retrieval call binding the contract method 0x548d496f.
//
// Solidity: function supportedChains(uint256 ) view returns(string)
func (_AddressPool *AddressPoolCaller) SupportedChains(opts *bind.CallOpts, arg0 *big.Int) (string, error) {
	var out []interface{}
	err := _AddressPool.contract.Call(opts, &out, "supportedChains", arg0)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// SupportedChains is a free data retrieval call binding the contract method 0x548d496f.
//
// Solidity: function supportedChains(uint256 ) view returns(string)
func (_AddressPool *AddressPoolSession) SupportedChains(arg0 *big.Int) (string, error) {
	return _AddressPool.Contract.SupportedChains(&_AddressPool.CallOpts, arg0)
}

// SupportedChains is a free data retrieval call binding the contract method 0x548d496f.
//
// Solidity: function supportedChains(uint256 ) view returns(string)
func (_AddressPool *AddressPoolCallerSession) SupportedChains(arg0 *big.Int) (string, error) {
	return _AddressPool.Contract.SupportedChains(&_AddressPool.CallOpts, arg0)
}

// VaultInfoNonce is a free data retrieval call binding the contract method 0x23b0d78c.
//
// Solidity: function vaultInfoNonce(string , string ) view returns(uint256)
func (_AddressPool *AddressPoolCaller) VaultInfoNonce(opts *bind.CallOpts, arg0 string, arg1 string) (*big.Int, error) {
	var out []interface{}
	err := _AddressPool.contract.Call(opts, &out, "vaultInfoNonce", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// VaultInfoNonce is a free data retrieval call binding the contract method 0x23b0d78c.
//
// Solidity: function vaultInfoNonce(string , string ) view returns(uint256)
func (_AddressPool *AddressPoolSession) VaultInfoNonce(arg0 string, arg1 string) (*big.Int, error) {
	return _AddressPool.Contract.VaultInfoNonce(&_AddressPool.CallOpts, arg0, arg1)
}

// VaultInfoNonce is a free data retrieval call binding the contract method 0x23b0d78c.
//
// Solidity: function vaultInfoNonce(string , string ) view returns(uint256)
func (_AddressPool *AddressPoolCallerSession) VaultInfoNonce(arg0 string, arg1 string) (*big.Int, error) {
	return _AddressPool.Contract.VaultInfoNonce(&_AddressPool.CallOpts, arg0, arg1)
}

// VaultTokenBalance is a free data retrieval call binding the contract method 0xe4828e01.
//
// Solidity: function vaultTokenBalance(string , string ) view returns(uint256)
func (_AddressPool *AddressPoolCaller) VaultTokenBalance(opts *bind.CallOpts, arg0 string, arg1 string) (*big.Int, error) {
	var out []interface{}
	err := _AddressPool.contract.Call(opts, &out, "vaultTokenBalance", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// VaultTokenBalance is a free data retrieval call binding the contract method 0xe4828e01.
//
// Solidity: function vaultTokenBalance(string , string ) view returns(uint256)
func (_AddressPool *AddressPoolSession) VaultTokenBalance(arg0 string, arg1 string) (*big.Int, error) {
	return _AddressPool.Contract.VaultTokenBalance(&_AddressPool.CallOpts, arg0, arg1)
}

// VaultTokenBalance is a free data retrieval call binding the contract method 0xe4828e01.
//
// Solidity: function vaultTokenBalance(string , string ) view returns(uint256)
func (_AddressPool *AddressPoolCallerSession) VaultTokenBalance(arg0 string, arg1 string) (*big.Int, error) {
	return _AddressPool.Contract.VaultTokenBalance(&_AddressPool.CallOpts, arg0, arg1)
}

// Vaults is a free data retrieval call binding the contract method 0xe2aa0719.
//
// Solidity: function vaults(string , uint256 ) view returns(string vaultAddress, uint8 status)
func (_AddressPool *AddressPoolCaller) Vaults(opts *bind.CallOpts, arg0 string, arg1 *big.Int) (struct {
	VaultAddress string
	Status       uint8
}, error) {
	var out []interface{}
	err := _AddressPool.contract.Call(opts, &out, "vaults", arg0, arg1)

	outstruct := new(struct {
		VaultAddress string
		Status       uint8
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.VaultAddress = *abi.ConvertType(out[0], new(string)).(*string)
	outstruct.Status = *abi.ConvertType(out[1], new(uint8)).(*uint8)

	return *outstruct, err

}

// Vaults is a free data retrieval call binding the contract method 0xe2aa0719.
//
// Solidity: function vaults(string , uint256 ) view returns(string vaultAddress, uint8 status)
func (_AddressPool *AddressPoolSession) Vaults(arg0 string, arg1 *big.Int) (struct {
	VaultAddress string
	Status       uint8
}, error) {
	return _AddressPool.Contract.Vaults(&_AddressPool.CallOpts, arg0, arg1)
}

// Vaults is a free data retrieval call binding the contract method 0xe2aa0719.
//
// Solidity: function vaults(string , uint256 ) view returns(string vaultAddress, uint8 status)
func (_AddressPool *AddressPoolCallerSession) Vaults(arg0 string, arg1 *big.Int) (struct {
	VaultAddress string
	Status       uint8
}, error) {
	return _AddressPool.Contract.Vaults(&_AddressPool.CallOpts, arg0, arg1)
}

// WithdrawRequestNonce is a free data retrieval call binding the contract method 0x6cf175a3.
//
// Solidity: function withdrawRequestNonce(bytes32 ) view returns(uint256)
func (_AddressPool *AddressPoolCaller) WithdrawRequestNonce(opts *bind.CallOpts, arg0 [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _AddressPool.contract.Call(opts, &out, "withdrawRequestNonce", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// WithdrawRequestNonce is a free data retrieval call binding the contract method 0x6cf175a3.
//
// Solidity: function withdrawRequestNonce(bytes32 ) view returns(uint256)
func (_AddressPool *AddressPoolSession) WithdrawRequestNonce(arg0 [32]byte) (*big.Int, error) {
	return _AddressPool.Contract.WithdrawRequestNonce(&_AddressPool.CallOpts, arg0)
}

// WithdrawRequestNonce is a free data retrieval call binding the contract method 0x6cf175a3.
//
// Solidity: function withdrawRequestNonce(bytes32 ) view returns(uint256)
func (_AddressPool *AddressPoolCallerSession) WithdrawRequestNonce(arg0 [32]byte) (*big.Int, error) {
	return _AddressPool.Contract.WithdrawRequestNonce(&_AddressPool.CallOpts, arg0)
}

// WithdrawRequestTxid is a free data retrieval call binding the contract method 0x25d5b898.
//
// Solidity: function withdrawRequestTxid(string ) view returns(bytes32)
func (_AddressPool *AddressPoolCaller) WithdrawRequestTxid(opts *bind.CallOpts, arg0 string) ([32]byte, error) {
	var out []interface{}
	err := _AddressPool.contract.Call(opts, &out, "withdrawRequestTxid", arg0)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// WithdrawRequestTxid is a free data retrieval call binding the contract method 0x25d5b898.
//
// Solidity: function withdrawRequestTxid(string ) view returns(bytes32)
func (_AddressPool *AddressPoolSession) WithdrawRequestTxid(arg0 string) ([32]byte, error) {
	return _AddressPool.Contract.WithdrawRequestTxid(&_AddressPool.CallOpts, arg0)
}

// WithdrawRequestTxid is a free data retrieval call binding the contract method 0x25d5b898.
//
// Solidity: function withdrawRequestTxid(string ) view returns(bytes32)
func (_AddressPool *AddressPoolCallerSession) WithdrawRequestTxid(arg0 string) ([32]byte, error) {
	return _AddressPool.Contract.WithdrawRequestTxid(&_AddressPool.CallOpts, arg0)
}

// WithdrawRequests is a free data retrieval call binding the contract method 0x992a7dfb.
//
// Solidity: function withdrawRequests(uint256 ) view returns(string requester, string token, uint256 amount, string txid, string vault, uint256 nonce, uint256 timestamp, uint8 status)
func (_AddressPool *AddressPoolCaller) WithdrawRequests(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Requester string
	Token     string
	Amount    *big.Int
	Txid      string
	Vault     string
	Nonce     *big.Int
	Timestamp *big.Int
	Status    uint8
}, error) {
	var out []interface{}
	err := _AddressPool.contract.Call(opts, &out, "withdrawRequests", arg0)

	outstruct := new(struct {
		Requester string
		Token     string
		Amount    *big.Int
		Txid      string
		Vault     string
		Nonce     *big.Int
		Timestamp *big.Int
		Status    uint8
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Requester = *abi.ConvertType(out[0], new(string)).(*string)
	outstruct.Token = *abi.ConvertType(out[1], new(string)).(*string)
	outstruct.Amount = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.Txid = *abi.ConvertType(out[3], new(string)).(*string)
	outstruct.Vault = *abi.ConvertType(out[4], new(string)).(*string)
	outstruct.Nonce = *abi.ConvertType(out[5], new(*big.Int)).(**big.Int)
	outstruct.Timestamp = *abi.ConvertType(out[6], new(*big.Int)).(**big.Int)
	outstruct.Status = *abi.ConvertType(out[7], new(uint8)).(*uint8)

	return *outstruct, err

}

// WithdrawRequests is a free data retrieval call binding the contract method 0x992a7dfb.
//
// Solidity: function withdrawRequests(uint256 ) view returns(string requester, string token, uint256 amount, string txid, string vault, uint256 nonce, uint256 timestamp, uint8 status)
func (_AddressPool *AddressPoolSession) WithdrawRequests(arg0 *big.Int) (struct {
	Requester string
	Token     string
	Amount    *big.Int
	Txid      string
	Vault     string
	Nonce     *big.Int
	Timestamp *big.Int
	Status    uint8
}, error) {
	return _AddressPool.Contract.WithdrawRequests(&_AddressPool.CallOpts, arg0)
}

// WithdrawRequests is a free data retrieval call binding the contract method 0x992a7dfb.
//
// Solidity: function withdrawRequests(uint256 ) view returns(string requester, string token, uint256 amount, string txid, string vault, uint256 nonce, uint256 timestamp, uint8 status)
func (_AddressPool *AddressPoolCallerSession) WithdrawRequests(arg0 *big.Int) (struct {
	Requester string
	Token     string
	Amount    *big.Int
	Txid      string
	Vault     string
	Nonce     *big.Int
	Timestamp *big.Int
	Status    uint8
}, error) {
	return _AddressPool.Contract.WithdrawRequests(&_AddressPool.CallOpts, arg0)
}

// AddNewVault is a paid mutator transaction binding the contract method 0x086a2cd7.
//
// Solidity: function addNewVault(string chain_, string vaultAddress_) returns()
func (_AddressPool *AddressPoolTransactor) AddNewVault(opts *bind.TransactOpts, chain_ string, vaultAddress_ string) (*types.Transaction, error) {
	return _AddressPool.contract.Transact(opts, "addNewVault", chain_, vaultAddress_)
}

// AddNewVault is a paid mutator transaction binding the contract method 0x086a2cd7.
//
// Solidity: function addNewVault(string chain_, string vaultAddress_) returns()
func (_AddressPool *AddressPoolSession) AddNewVault(chain_ string, vaultAddress_ string) (*types.Transaction, error) {
	return _AddressPool.Contract.AddNewVault(&_AddressPool.TransactOpts, chain_, vaultAddress_)
}

// AddNewVault is a paid mutator transaction binding the contract method 0x086a2cd7.
//
// Solidity: function addNewVault(string chain_, string vaultAddress_) returns()
func (_AddressPool *AddressPoolTransactorSession) AddNewVault(chain_ string, vaultAddress_ string) (*types.Transaction, error) {
	return _AddressPool.Contract.AddNewVault(&_AddressPool.TransactOpts, chain_, vaultAddress_)
}

// ReleaseVault is a paid mutator transaction binding the contract method 0x8ca4a359.
//
// Solidity: function releaseVault(string chain_, string vaultAddress_) returns()
func (_AddressPool *AddressPoolTransactor) ReleaseVault(opts *bind.TransactOpts, chain_ string, vaultAddress_ string) (*types.Transaction, error) {
	return _AddressPool.contract.Transact(opts, "releaseVault", chain_, vaultAddress_)
}

// ReleaseVault is a paid mutator transaction binding the contract method 0x8ca4a359.
//
// Solidity: function releaseVault(string chain_, string vaultAddress_) returns()
func (_AddressPool *AddressPoolSession) ReleaseVault(chain_ string, vaultAddress_ string) (*types.Transaction, error) {
	return _AddressPool.Contract.ReleaseVault(&_AddressPool.TransactOpts, chain_, vaultAddress_)
}

// ReleaseVault is a paid mutator transaction binding the contract method 0x8ca4a359.
//
// Solidity: function releaseVault(string chain_, string vaultAddress_) returns()
func (_AddressPool *AddressPoolTransactorSession) ReleaseVault(chain_ string, vaultAddress_ string) (*types.Transaction, error) {
	return _AddressPool.Contract.ReleaseVault(&_AddressPool.TransactOpts, chain_, vaultAddress_)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_AddressPool *AddressPoolTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AddressPool.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_AddressPool *AddressPoolSession) RenounceOwnership() (*types.Transaction, error) {
	return _AddressPool.Contract.RenounceOwnership(&_AddressPool.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_AddressPool *AddressPoolTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _AddressPool.Contract.RenounceOwnership(&_AddressPool.TransactOpts)
}

// ResetVaults is a paid mutator transaction binding the contract method 0x79fb1092.
//
// Solidity: function resetVaults(string chain_, string[] vaultAddresses_) returns()
func (_AddressPool *AddressPoolTransactor) ResetVaults(opts *bind.TransactOpts, chain_ string, vaultAddresses_ []string) (*types.Transaction, error) {
	return _AddressPool.contract.Transact(opts, "resetVaults", chain_, vaultAddresses_)
}

// ResetVaults is a paid mutator transaction binding the contract method 0x79fb1092.
//
// Solidity: function resetVaults(string chain_, string[] vaultAddresses_) returns()
func (_AddressPool *AddressPoolSession) ResetVaults(chain_ string, vaultAddresses_ []string) (*types.Transaction, error) {
	return _AddressPool.Contract.ResetVaults(&_AddressPool.TransactOpts, chain_, vaultAddresses_)
}

// ResetVaults is a paid mutator transaction binding the contract method 0x79fb1092.
//
// Solidity: function resetVaults(string chain_, string[] vaultAddresses_) returns()
func (_AddressPool *AddressPoolTransactorSession) ResetVaults(chain_ string, vaultAddresses_ []string) (*types.Transaction, error) {
	return _AddressPool.Contract.ResetVaults(&_AddressPool.TransactOpts, chain_, vaultAddresses_)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_AddressPool *AddressPoolTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _AddressPool.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_AddressPool *AddressPoolSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _AddressPool.Contract.TransferOwnership(&_AddressPool.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_AddressPool *AddressPoolTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _AddressPool.Contract.TransferOwnership(&_AddressPool.TransactOpts, newOwner)
}

// UseAvailableVault is a paid mutator transaction binding the contract method 0xe9aab70a.
//
// Solidity: function useAvailableVault(uint256 index_, string chain_) returns(string vaultAddress)
func (_AddressPool *AddressPoolTransactor) UseAvailableVault(opts *bind.TransactOpts, index_ *big.Int, chain_ string) (*types.Transaction, error) {
	return _AddressPool.contract.Transact(opts, "useAvailableVault", index_, chain_)
}

// UseAvailableVault is a paid mutator transaction binding the contract method 0xe9aab70a.
//
// Solidity: function useAvailableVault(uint256 index_, string chain_) returns(string vaultAddress)
func (_AddressPool *AddressPoolSession) UseAvailableVault(index_ *big.Int, chain_ string) (*types.Transaction, error) {
	return _AddressPool.Contract.UseAvailableVault(&_AddressPool.TransactOpts, index_, chain_)
}

// UseAvailableVault is a paid mutator transaction binding the contract method 0xe9aab70a.
//
// Solidity: function useAvailableVault(uint256 index_, string chain_) returns(string vaultAddress)
func (_AddressPool *AddressPoolTransactorSession) UseAvailableVault(index_ *big.Int, chain_ string) (*types.Transaction, error) {
	return _AddressPool.Contract.UseAvailableVault(&_AddressPool.TransactOpts, index_, chain_)
}

// UseVault is a paid mutator transaction binding the contract method 0xead6d563.
//
// Solidity: function useVault(uint256 index_, string chain_, string vaultAddress_) returns()
func (_AddressPool *AddressPoolTransactor) UseVault(opts *bind.TransactOpts, index_ *big.Int, chain_ string, vaultAddress_ string) (*types.Transaction, error) {
	return _AddressPool.contract.Transact(opts, "useVault", index_, chain_, vaultAddress_)
}

// UseVault is a paid mutator transaction binding the contract method 0xead6d563.
//
// Solidity: function useVault(uint256 index_, string chain_, string vaultAddress_) returns()
func (_AddressPool *AddressPoolSession) UseVault(index_ *big.Int, chain_ string, vaultAddress_ string) (*types.Transaction, error) {
	return _AddressPool.Contract.UseVault(&_AddressPool.TransactOpts, index_, chain_, vaultAddress_)
}

// UseVault is a paid mutator transaction binding the contract method 0xead6d563.
//
// Solidity: function useVault(uint256 index_, string chain_, string vaultAddress_) returns()
func (_AddressPool *AddressPoolTransactorSession) UseVault(index_ *big.Int, chain_ string, vaultAddress_ string) (*types.Transaction, error) {
	return _AddressPool.Contract.UseVault(&_AddressPool.TransactOpts, index_, chain_, vaultAddress_)
}

// AddressPoolAddedNewVaultIterator is returned from FilterAddedNewVault and is used to iterate over the raw logs and unpacked data for AddedNewVault events raised by the AddressPool contract.
type AddressPoolAddedNewVaultIterator struct {
	Event *AddressPoolAddedNewVault // Event containing the contract specifics and raw log

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
func (it *AddressPoolAddedNewVaultIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AddressPoolAddedNewVault)
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
		it.Event = new(AddressPoolAddedNewVault)
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
func (it *AddressPoolAddedNewVaultIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AddressPoolAddedNewVaultIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AddressPoolAddedNewVault represents a AddedNewVault event raised by the AddressPool contract.
type AddressPoolAddedNewVault struct {
	Chain        string
	VaultAddress string
	Nonce        *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterAddedNewVault is a free log retrieval operation binding the contract event 0x032d76c05674b294893578f06241742277ab65ac9f8ca3d421971b2d38992083.
//
// Solidity: event AddedNewVault(string chain, string vaultAddress, uint256 nonce)
func (_AddressPool *AddressPoolFilterer) FilterAddedNewVault(opts *bind.FilterOpts) (*AddressPoolAddedNewVaultIterator, error) {

	logs, sub, err := _AddressPool.contract.FilterLogs(opts, "AddedNewVault")
	if err != nil {
		return nil, err
	}
	return &AddressPoolAddedNewVaultIterator{contract: _AddressPool.contract, event: "AddedNewVault", logs: logs, sub: sub}, nil
}

// WatchAddedNewVault is a free log subscription operation binding the contract event 0x032d76c05674b294893578f06241742277ab65ac9f8ca3d421971b2d38992083.
//
// Solidity: event AddedNewVault(string chain, string vaultAddress, uint256 nonce)
func (_AddressPool *AddressPoolFilterer) WatchAddedNewVault(opts *bind.WatchOpts, sink chan<- *AddressPoolAddedNewVault) (event.Subscription, error) {

	logs, sub, err := _AddressPool.contract.WatchLogs(opts, "AddedNewVault")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AddressPoolAddedNewVault)
				if err := _AddressPool.contract.UnpackLog(event, "AddedNewVault", log); err != nil {
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

// ParseAddedNewVault is a log parse operation binding the contract event 0x032d76c05674b294893578f06241742277ab65ac9f8ca3d421971b2d38992083.
//
// Solidity: event AddedNewVault(string chain, string vaultAddress, uint256 nonce)
func (_AddressPool *AddressPoolFilterer) ParseAddedNewVault(log types.Log) (*AddressPoolAddedNewVault, error) {
	event := new(AddressPoolAddedNewVault)
	if err := _AddressPool.contract.UnpackLog(event, "AddedNewVault", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AddressPoolOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the AddressPool contract.
type AddressPoolOwnershipTransferredIterator struct {
	Event *AddressPoolOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *AddressPoolOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AddressPoolOwnershipTransferred)
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
		it.Event = new(AddressPoolOwnershipTransferred)
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
func (it *AddressPoolOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AddressPoolOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AddressPoolOwnershipTransferred represents a OwnershipTransferred event raised by the AddressPool contract.
type AddressPoolOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_AddressPool *AddressPoolFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*AddressPoolOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _AddressPool.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &AddressPoolOwnershipTransferredIterator{contract: _AddressPool.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_AddressPool *AddressPoolFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *AddressPoolOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _AddressPool.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AddressPoolOwnershipTransferred)
				if err := _AddressPool.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_AddressPool *AddressPoolFilterer) ParseOwnershipTransferred(log types.Log) (*AddressPoolOwnershipTransferred, error) {
	event := new(AddressPoolOwnershipTransferred)
	if err := _AddressPool.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AddressPoolPaymentRequestAddIterator is returned from FilterPaymentRequestAdd and is used to iterate over the raw logs and unpacked data for PaymentRequestAdd events raised by the AddressPool contract.
type AddressPoolPaymentRequestAddIterator struct {
	Event *AddressPoolPaymentRequestAdd // Event containing the contract specifics and raw log

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
func (it *AddressPoolPaymentRequestAddIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AddressPoolPaymentRequestAdd)
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
		it.Event = new(AddressPoolPaymentRequestAdd)
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
func (it *AddressPoolPaymentRequestAddIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AddressPoolPaymentRequestAddIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AddressPoolPaymentRequestAdd represents a PaymentRequestAdd event raised by the AddressPool contract.
type AddressPoolPaymentRequestAdd struct {
	Nonce       *big.Int
	Requester   common.Hash
	Token       common.Hash
	Amount      *big.Int
	Txid        string
	Vault       string
	Timestamp   *big.Int
	RequestHash [32]byte
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterPaymentRequestAdd is a free log retrieval operation binding the contract event 0xd42b47b4f0b1df1765d3b6e4013946c28e8514a24f035cadca52a6030efb0c3e.
//
// Solidity: event PaymentRequestAdd(uint256 indexed nonce, string indexed requester, string indexed token, uint256 amount, string txid, string vault, uint256 timestamp, bytes32 requestHash)
func (_AddressPool *AddressPoolFilterer) FilterPaymentRequestAdd(opts *bind.FilterOpts, nonce []*big.Int, requester []string, token []string) (*AddressPoolPaymentRequestAddIterator, error) {

	var nonceRule []interface{}
	for _, nonceItem := range nonce {
		nonceRule = append(nonceRule, nonceItem)
	}
	var requesterRule []interface{}
	for _, requesterItem := range requester {
		requesterRule = append(requesterRule, requesterItem)
	}
	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _AddressPool.contract.FilterLogs(opts, "PaymentRequestAdd", nonceRule, requesterRule, tokenRule)
	if err != nil {
		return nil, err
	}
	return &AddressPoolPaymentRequestAddIterator{contract: _AddressPool.contract, event: "PaymentRequestAdd", logs: logs, sub: sub}, nil
}

// WatchPaymentRequestAdd is a free log subscription operation binding the contract event 0xd42b47b4f0b1df1765d3b6e4013946c28e8514a24f035cadca52a6030efb0c3e.
//
// Solidity: event PaymentRequestAdd(uint256 indexed nonce, string indexed requester, string indexed token, uint256 amount, string txid, string vault, uint256 timestamp, bytes32 requestHash)
func (_AddressPool *AddressPoolFilterer) WatchPaymentRequestAdd(opts *bind.WatchOpts, sink chan<- *AddressPoolPaymentRequestAdd, nonce []*big.Int, requester []string, token []string) (event.Subscription, error) {

	var nonceRule []interface{}
	for _, nonceItem := range nonce {
		nonceRule = append(nonceRule, nonceItem)
	}
	var requesterRule []interface{}
	for _, requesterItem := range requester {
		requesterRule = append(requesterRule, requesterItem)
	}
	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _AddressPool.contract.WatchLogs(opts, "PaymentRequestAdd", nonceRule, requesterRule, tokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AddressPoolPaymentRequestAdd)
				if err := _AddressPool.contract.UnpackLog(event, "PaymentRequestAdd", log); err != nil {
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

// ParsePaymentRequestAdd is a log parse operation binding the contract event 0xd42b47b4f0b1df1765d3b6e4013946c28e8514a24f035cadca52a6030efb0c3e.
//
// Solidity: event PaymentRequestAdd(uint256 indexed nonce, string indexed requester, string indexed token, uint256 amount, string txid, string vault, uint256 timestamp, bytes32 requestHash)
func (_AddressPool *AddressPoolFilterer) ParsePaymentRequestAdd(log types.Log) (*AddressPoolPaymentRequestAdd, error) {
	event := new(AddressPoolPaymentRequestAdd)
	if err := _AddressPool.contract.UnpackLog(event, "PaymentRequestAdd", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AddressPoolResetVaultsIterator is returned from FilterResetVaults and is used to iterate over the raw logs and unpacked data for ResetVaults events raised by the AddressPool contract.
type AddressPoolResetVaultsIterator struct {
	Event *AddressPoolResetVaults // Event containing the contract specifics and raw log

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
func (it *AddressPoolResetVaultsIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AddressPoolResetVaults)
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
		it.Event = new(AddressPoolResetVaults)
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
func (it *AddressPoolResetVaultsIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AddressPoolResetVaultsIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AddressPoolResetVaults represents a ResetVaults event raised by the AddressPool contract.
type AddressPoolResetVaults struct {
	Chain          string
	NumberOfvaults *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterResetVaults is a free log retrieval operation binding the contract event 0x688678ae71d66c5e0a1c117c27588f217a1b5b416ce5df6223f9e3028a59c89f.
//
// Solidity: event ResetVaults(string chain, uint256 numberOfvaults)
func (_AddressPool *AddressPoolFilterer) FilterResetVaults(opts *bind.FilterOpts) (*AddressPoolResetVaultsIterator, error) {

	logs, sub, err := _AddressPool.contract.FilterLogs(opts, "ResetVaults")
	if err != nil {
		return nil, err
	}
	return &AddressPoolResetVaultsIterator{contract: _AddressPool.contract, event: "ResetVaults", logs: logs, sub: sub}, nil
}

// WatchResetVaults is a free log subscription operation binding the contract event 0x688678ae71d66c5e0a1c117c27588f217a1b5b416ce5df6223f9e3028a59c89f.
//
// Solidity: event ResetVaults(string chain, uint256 numberOfvaults)
func (_AddressPool *AddressPoolFilterer) WatchResetVaults(opts *bind.WatchOpts, sink chan<- *AddressPoolResetVaults) (event.Subscription, error) {

	logs, sub, err := _AddressPool.contract.WatchLogs(opts, "ResetVaults")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AddressPoolResetVaults)
				if err := _AddressPool.contract.UnpackLog(event, "ResetVaults", log); err != nil {
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

// ParseResetVaults is a log parse operation binding the contract event 0x688678ae71d66c5e0a1c117c27588f217a1b5b416ce5df6223f9e3028a59c89f.
//
// Solidity: event ResetVaults(string chain, uint256 numberOfvaults)
func (_AddressPool *AddressPoolFilterer) ParseResetVaults(log types.Log) (*AddressPoolResetVaults, error) {
	event := new(AddressPoolResetVaults)
	if err := _AddressPool.contract.UnpackLog(event, "ResetVaults", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AddressPoolUsingVaultIterator is returned from FilterUsingVault and is used to iterate over the raw logs and unpacked data for UsingVault events raised by the AddressPool contract.
type AddressPoolUsingVaultIterator struct {
	Event *AddressPoolUsingVault // Event containing the contract specifics and raw log

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
func (it *AddressPoolUsingVaultIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AddressPoolUsingVault)
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
		it.Event = new(AddressPoolUsingVault)
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
func (it *AddressPoolUsingVaultIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AddressPoolUsingVaultIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AddressPoolUsingVault represents a UsingVault event raised by the AddressPool contract.
type AddressPoolUsingVault struct {
	Index        *big.Int
	Chain        string
	VaultAddress string
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterUsingVault is a free log retrieval operation binding the contract event 0x3243ca28802d2b9cfb4da2d678fb489d884177ee6746d39fd688b8a10af85cff.
//
// Solidity: event UsingVault(uint256 index, string chain, string vaultAddress)
func (_AddressPool *AddressPoolFilterer) FilterUsingVault(opts *bind.FilterOpts) (*AddressPoolUsingVaultIterator, error) {

	logs, sub, err := _AddressPool.contract.FilterLogs(opts, "UsingVault")
	if err != nil {
		return nil, err
	}
	return &AddressPoolUsingVaultIterator{contract: _AddressPool.contract, event: "UsingVault", logs: logs, sub: sub}, nil
}

// WatchUsingVault is a free log subscription operation binding the contract event 0x3243ca28802d2b9cfb4da2d678fb489d884177ee6746d39fd688b8a10af85cff.
//
// Solidity: event UsingVault(uint256 index, string chain, string vaultAddress)
func (_AddressPool *AddressPoolFilterer) WatchUsingVault(opts *bind.WatchOpts, sink chan<- *AddressPoolUsingVault) (event.Subscription, error) {

	logs, sub, err := _AddressPool.contract.WatchLogs(opts, "UsingVault")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AddressPoolUsingVault)
				if err := _AddressPool.contract.UnpackLog(event, "UsingVault", log); err != nil {
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

// ParseUsingVault is a log parse operation binding the contract event 0x3243ca28802d2b9cfb4da2d678fb489d884177ee6746d39fd688b8a10af85cff.
//
// Solidity: event UsingVault(uint256 index, string chain, string vaultAddress)
func (_AddressPool *AddressPoolFilterer) ParseUsingVault(log types.Log) (*AddressPoolUsingVault, error) {
	event := new(AddressPoolUsingVault)
	if err := _AddressPool.contract.UnpackLog(event, "UsingVault", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AddressPoolVaultReleasedIterator is returned from FilterVaultReleased and is used to iterate over the raw logs and unpacked data for VaultReleased events raised by the AddressPool contract.
type AddressPoolVaultReleasedIterator struct {
	Event *AddressPoolVaultReleased // Event containing the contract specifics and raw log

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
func (it *AddressPoolVaultReleasedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AddressPoolVaultReleased)
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
		it.Event = new(AddressPoolVaultReleased)
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
func (it *AddressPoolVaultReleasedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AddressPoolVaultReleasedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AddressPoolVaultReleased represents a VaultReleased event raised by the AddressPool contract.
type AddressPoolVaultReleased struct {
	Chain        string
	VaultAddress string
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterVaultReleased is a free log retrieval operation binding the contract event 0x4586d66964c4eafe1deea3c83f0a7167b1e591a4554bba96e38211c31e8c50fc.
//
// Solidity: event VaultReleased(string chain, string vaultAddress)
func (_AddressPool *AddressPoolFilterer) FilterVaultReleased(opts *bind.FilterOpts) (*AddressPoolVaultReleasedIterator, error) {

	logs, sub, err := _AddressPool.contract.FilterLogs(opts, "VaultReleased")
	if err != nil {
		return nil, err
	}
	return &AddressPoolVaultReleasedIterator{contract: _AddressPool.contract, event: "VaultReleased", logs: logs, sub: sub}, nil
}

// WatchVaultReleased is a free log subscription operation binding the contract event 0x4586d66964c4eafe1deea3c83f0a7167b1e591a4554bba96e38211c31e8c50fc.
//
// Solidity: event VaultReleased(string chain, string vaultAddress)
func (_AddressPool *AddressPoolFilterer) WatchVaultReleased(opts *bind.WatchOpts, sink chan<- *AddressPoolVaultReleased) (event.Subscription, error) {

	logs, sub, err := _AddressPool.contract.WatchLogs(opts, "VaultReleased")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AddressPoolVaultReleased)
				if err := _AddressPool.contract.UnpackLog(event, "VaultReleased", log); err != nil {
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

// ParseVaultReleased is a log parse operation binding the contract event 0x4586d66964c4eafe1deea3c83f0a7167b1e591a4554bba96e38211c31e8c50fc.
//
// Solidity: event VaultReleased(string chain, string vaultAddress)
func (_AddressPool *AddressPoolFilterer) ParseVaultReleased(log types.Log) (*AddressPoolVaultReleased, error) {
	event := new(AddressPoolVaultReleased)
	if err := _AddressPool.contract.UnpackLog(event, "VaultReleased", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
