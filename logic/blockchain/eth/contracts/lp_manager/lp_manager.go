// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package lp_manager

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

// LpManagerMetaData contains all meta data concerning the LpManager contract.
var LpManagerMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"EthLiquidityAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"EthLiquidityRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"EthLiquidityWithdraw\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"pool\",\"type\":\"address\"}],\"name\":\"EthPoolCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"EthRewardsAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"LiquidityAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"LiquidityRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"LiquidityWithdraw\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"pool\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"PoolCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"RewardsAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"addEthLiquidity\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"liquidity\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"addEthRewards\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"addLiquidity\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"liquidity\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"addRewards\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"allPools\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"backend\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"createEthPool\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"pool\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"createPool\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"pool\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"deposit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"depositEth\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ethPool\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getEthPoolBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getEthPoolReserves\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"reserves\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"getPool\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"getPoolBalanceViaToken\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"getPoolReservesViaToken\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"reserves\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pauseEth\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"pauseToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"poolsLength\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"liquidity\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"removeEthLiquidity\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"liquidity\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"removeLiquidity\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"backend_\",\"type\":\"address\"}],\"name\":\"setBackend\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unpause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unpauseEth\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"unpauseToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"withdrawEth\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
}

// LpManagerABI is the input ABI used to generate the binding from.
// Deprecated: Use LpManagerMetaData.ABI instead.
var LpManagerABI = LpManagerMetaData.ABI

// LpManager is an auto generated Go binding around an Ethereum contract.
type LpManager struct {
	LpManagerCaller     // Read-only binding to the contract
	LpManagerTransactor // Write-only binding to the contract
	LpManagerFilterer   // Log filterer for contract events
}

// LpManagerCaller is an auto generated read-only Go binding around an Ethereum contract.
type LpManagerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LpManagerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type LpManagerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LpManagerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type LpManagerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LpManagerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type LpManagerSession struct {
	Contract     *LpManager        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// LpManagerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type LpManagerCallerSession struct {
	Contract *LpManagerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// LpManagerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type LpManagerTransactorSession struct {
	Contract     *LpManagerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// LpManagerRaw is an auto generated low-level Go binding around an Ethereum contract.
type LpManagerRaw struct {
	Contract *LpManager // Generic contract binding to access the raw methods on
}

// LpManagerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type LpManagerCallerRaw struct {
	Contract *LpManagerCaller // Generic read-only contract binding to access the raw methods on
}

// LpManagerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type LpManagerTransactorRaw struct {
	Contract *LpManagerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewLpManager creates a new instance of LpManager, bound to a specific deployed contract.
func NewLpManager(address common.Address, backend bind.ContractBackend) (*LpManager, error) {
	contract, err := bindLpManager(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &LpManager{LpManagerCaller: LpManagerCaller{contract: contract}, LpManagerTransactor: LpManagerTransactor{contract: contract}, LpManagerFilterer: LpManagerFilterer{contract: contract}}, nil
}

// NewLpManagerCaller creates a new read-only instance of LpManager, bound to a specific deployed contract.
func NewLpManagerCaller(address common.Address, caller bind.ContractCaller) (*LpManagerCaller, error) {
	contract, err := bindLpManager(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &LpManagerCaller{contract: contract}, nil
}

// NewLpManagerTransactor creates a new write-only instance of LpManager, bound to a specific deployed contract.
func NewLpManagerTransactor(address common.Address, transactor bind.ContractTransactor) (*LpManagerTransactor, error) {
	contract, err := bindLpManager(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &LpManagerTransactor{contract: contract}, nil
}

// NewLpManagerFilterer creates a new log filterer instance of LpManager, bound to a specific deployed contract.
func NewLpManagerFilterer(address common.Address, filterer bind.ContractFilterer) (*LpManagerFilterer, error) {
	contract, err := bindLpManager(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &LpManagerFilterer{contract: contract}, nil
}

// bindLpManager binds a generic wrapper to an already deployed contract.
func bindLpManager(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := LpManagerMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_LpManager *LpManagerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _LpManager.Contract.LpManagerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_LpManager *LpManagerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LpManager.Contract.LpManagerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_LpManager *LpManagerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _LpManager.Contract.LpManagerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_LpManager *LpManagerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _LpManager.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_LpManager *LpManagerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LpManager.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_LpManager *LpManagerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _LpManager.Contract.contract.Transact(opts, method, params...)
}

// AllPools is a free data retrieval call binding the contract method 0x41d1de97.
//
// Solidity: function allPools(uint256 ) view returns(address)
func (_LpManager *LpManagerCaller) AllPools(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _LpManager.contract.Call(opts, &out, "allPools", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// AllPools is a free data retrieval call binding the contract method 0x41d1de97.
//
// Solidity: function allPools(uint256 ) view returns(address)
func (_LpManager *LpManagerSession) AllPools(arg0 *big.Int) (common.Address, error) {
	return _LpManager.Contract.AllPools(&_LpManager.CallOpts, arg0)
}

// AllPools is a free data retrieval call binding the contract method 0x41d1de97.
//
// Solidity: function allPools(uint256 ) view returns(address)
func (_LpManager *LpManagerCallerSession) AllPools(arg0 *big.Int) (common.Address, error) {
	return _LpManager.Contract.AllPools(&_LpManager.CallOpts, arg0)
}

// Backend is a free data retrieval call binding the contract method 0x099e4133.
//
// Solidity: function backend() view returns(address)
func (_LpManager *LpManagerCaller) Backend(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _LpManager.contract.Call(opts, &out, "backend")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Backend is a free data retrieval call binding the contract method 0x099e4133.
//
// Solidity: function backend() view returns(address)
func (_LpManager *LpManagerSession) Backend() (common.Address, error) {
	return _LpManager.Contract.Backend(&_LpManager.CallOpts)
}

// Backend is a free data retrieval call binding the contract method 0x099e4133.
//
// Solidity: function backend() view returns(address)
func (_LpManager *LpManagerCallerSession) Backend() (common.Address, error) {
	return _LpManager.Contract.Backend(&_LpManager.CallOpts)
}

// EthPool is a free data retrieval call binding the contract method 0xf16673a4.
//
// Solidity: function ethPool() view returns(address)
func (_LpManager *LpManagerCaller) EthPool(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _LpManager.contract.Call(opts, &out, "ethPool")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// EthPool is a free data retrieval call binding the contract method 0xf16673a4.
//
// Solidity: function ethPool() view returns(address)
func (_LpManager *LpManagerSession) EthPool() (common.Address, error) {
	return _LpManager.Contract.EthPool(&_LpManager.CallOpts)
}

// EthPool is a free data retrieval call binding the contract method 0xf16673a4.
//
// Solidity: function ethPool() view returns(address)
func (_LpManager *LpManagerCallerSession) EthPool() (common.Address, error) {
	return _LpManager.Contract.EthPool(&_LpManager.CallOpts)
}

// GetEthPoolBalance is a free data retrieval call binding the contract method 0x329a1cd1.
//
// Solidity: function getEthPoolBalance() view returns(uint256 balance)
func (_LpManager *LpManagerCaller) GetEthPoolBalance(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _LpManager.contract.Call(opts, &out, "getEthPoolBalance")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetEthPoolBalance is a free data retrieval call binding the contract method 0x329a1cd1.
//
// Solidity: function getEthPoolBalance() view returns(uint256 balance)
func (_LpManager *LpManagerSession) GetEthPoolBalance() (*big.Int, error) {
	return _LpManager.Contract.GetEthPoolBalance(&_LpManager.CallOpts)
}

// GetEthPoolBalance is a free data retrieval call binding the contract method 0x329a1cd1.
//
// Solidity: function getEthPoolBalance() view returns(uint256 balance)
func (_LpManager *LpManagerCallerSession) GetEthPoolBalance() (*big.Int, error) {
	return _LpManager.Contract.GetEthPoolBalance(&_LpManager.CallOpts)
}

// GetEthPoolReserves is a free data retrieval call binding the contract method 0x67f50e68.
//
// Solidity: function getEthPoolReserves() view returns(uint256 reserves)
func (_LpManager *LpManagerCaller) GetEthPoolReserves(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _LpManager.contract.Call(opts, &out, "getEthPoolReserves")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetEthPoolReserves is a free data retrieval call binding the contract method 0x67f50e68.
//
// Solidity: function getEthPoolReserves() view returns(uint256 reserves)
func (_LpManager *LpManagerSession) GetEthPoolReserves() (*big.Int, error) {
	return _LpManager.Contract.GetEthPoolReserves(&_LpManager.CallOpts)
}

// GetEthPoolReserves is a free data retrieval call binding the contract method 0x67f50e68.
//
// Solidity: function getEthPoolReserves() view returns(uint256 reserves)
func (_LpManager *LpManagerCallerSession) GetEthPoolReserves() (*big.Int, error) {
	return _LpManager.Contract.GetEthPoolReserves(&_LpManager.CallOpts)
}

// GetPool is a free data retrieval call binding the contract method 0xbbe4f6db.
//
// Solidity: function getPool(address ) view returns(address)
func (_LpManager *LpManagerCaller) GetPool(opts *bind.CallOpts, arg0 common.Address) (common.Address, error) {
	var out []interface{}
	err := _LpManager.contract.Call(opts, &out, "getPool", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetPool is a free data retrieval call binding the contract method 0xbbe4f6db.
//
// Solidity: function getPool(address ) view returns(address)
func (_LpManager *LpManagerSession) GetPool(arg0 common.Address) (common.Address, error) {
	return _LpManager.Contract.GetPool(&_LpManager.CallOpts, arg0)
}

// GetPool is a free data retrieval call binding the contract method 0xbbe4f6db.
//
// Solidity: function getPool(address ) view returns(address)
func (_LpManager *LpManagerCallerSession) GetPool(arg0 common.Address) (common.Address, error) {
	return _LpManager.Contract.GetPool(&_LpManager.CallOpts, arg0)
}

// GetPoolBalanceViaToken is a free data retrieval call binding the contract method 0x350c96e7.
//
// Solidity: function getPoolBalanceViaToken(address token) view returns(uint256 balance)
func (_LpManager *LpManagerCaller) GetPoolBalanceViaToken(opts *bind.CallOpts, token common.Address) (*big.Int, error) {
	var out []interface{}
	err := _LpManager.contract.Call(opts, &out, "getPoolBalanceViaToken", token)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetPoolBalanceViaToken is a free data retrieval call binding the contract method 0x350c96e7.
//
// Solidity: function getPoolBalanceViaToken(address token) view returns(uint256 balance)
func (_LpManager *LpManagerSession) GetPoolBalanceViaToken(token common.Address) (*big.Int, error) {
	return _LpManager.Contract.GetPoolBalanceViaToken(&_LpManager.CallOpts, token)
}

// GetPoolBalanceViaToken is a free data retrieval call binding the contract method 0x350c96e7.
//
// Solidity: function getPoolBalanceViaToken(address token) view returns(uint256 balance)
func (_LpManager *LpManagerCallerSession) GetPoolBalanceViaToken(token common.Address) (*big.Int, error) {
	return _LpManager.Contract.GetPoolBalanceViaToken(&_LpManager.CallOpts, token)
}

// GetPoolReservesViaToken is a free data retrieval call binding the contract method 0x431fb8e6.
//
// Solidity: function getPoolReservesViaToken(address token) view returns(uint256 reserves)
func (_LpManager *LpManagerCaller) GetPoolReservesViaToken(opts *bind.CallOpts, token common.Address) (*big.Int, error) {
	var out []interface{}
	err := _LpManager.contract.Call(opts, &out, "getPoolReservesViaToken", token)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetPoolReservesViaToken is a free data retrieval call binding the contract method 0x431fb8e6.
//
// Solidity: function getPoolReservesViaToken(address token) view returns(uint256 reserves)
func (_LpManager *LpManagerSession) GetPoolReservesViaToken(token common.Address) (*big.Int, error) {
	return _LpManager.Contract.GetPoolReservesViaToken(&_LpManager.CallOpts, token)
}

// GetPoolReservesViaToken is a free data retrieval call binding the contract method 0x431fb8e6.
//
// Solidity: function getPoolReservesViaToken(address token) view returns(uint256 reserves)
func (_LpManager *LpManagerCallerSession) GetPoolReservesViaToken(token common.Address) (*big.Int, error) {
	return _LpManager.Contract.GetPoolReservesViaToken(&_LpManager.CallOpts, token)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_LpManager *LpManagerCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _LpManager.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_LpManager *LpManagerSession) Owner() (common.Address, error) {
	return _LpManager.Contract.Owner(&_LpManager.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_LpManager *LpManagerCallerSession) Owner() (common.Address, error) {
	return _LpManager.Contract.Owner(&_LpManager.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_LpManager *LpManagerCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _LpManager.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_LpManager *LpManagerSession) Paused() (bool, error) {
	return _LpManager.Contract.Paused(&_LpManager.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_LpManager *LpManagerCallerSession) Paused() (bool, error) {
	return _LpManager.Contract.Paused(&_LpManager.CallOpts)
}

// PoolsLength is a free data retrieval call binding the contract method 0x2716ae66.
//
// Solidity: function poolsLength() view returns(uint256)
func (_LpManager *LpManagerCaller) PoolsLength(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _LpManager.contract.Call(opts, &out, "poolsLength")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PoolsLength is a free data retrieval call binding the contract method 0x2716ae66.
//
// Solidity: function poolsLength() view returns(uint256)
func (_LpManager *LpManagerSession) PoolsLength() (*big.Int, error) {
	return _LpManager.Contract.PoolsLength(&_LpManager.CallOpts)
}

// PoolsLength is a free data retrieval call binding the contract method 0x2716ae66.
//
// Solidity: function poolsLength() view returns(uint256)
func (_LpManager *LpManagerCallerSession) PoolsLength() (*big.Int, error) {
	return _LpManager.Contract.PoolsLength(&_LpManager.CallOpts)
}

// AddEthLiquidity is a paid mutator transaction binding the contract method 0x07a670bb.
//
// Solidity: function addEthLiquidity(address to) payable returns(uint256 liquidity)
func (_LpManager *LpManagerTransactor) AddEthLiquidity(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error) {
	return _LpManager.contract.Transact(opts, "addEthLiquidity", to)
}

// AddEthLiquidity is a paid mutator transaction binding the contract method 0x07a670bb.
//
// Solidity: function addEthLiquidity(address to) payable returns(uint256 liquidity)
func (_LpManager *LpManagerSession) AddEthLiquidity(to common.Address) (*types.Transaction, error) {
	return _LpManager.Contract.AddEthLiquidity(&_LpManager.TransactOpts, to)
}

// AddEthLiquidity is a paid mutator transaction binding the contract method 0x07a670bb.
//
// Solidity: function addEthLiquidity(address to) payable returns(uint256 liquidity)
func (_LpManager *LpManagerTransactorSession) AddEthLiquidity(to common.Address) (*types.Transaction, error) {
	return _LpManager.Contract.AddEthLiquidity(&_LpManager.TransactOpts, to)
}

// AddEthRewards is a paid mutator transaction binding the contract method 0x97bb441a.
//
// Solidity: function addEthRewards() payable returns()
func (_LpManager *LpManagerTransactor) AddEthRewards(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LpManager.contract.Transact(opts, "addEthRewards")
}

// AddEthRewards is a paid mutator transaction binding the contract method 0x97bb441a.
//
// Solidity: function addEthRewards() payable returns()
func (_LpManager *LpManagerSession) AddEthRewards() (*types.Transaction, error) {
	return _LpManager.Contract.AddEthRewards(&_LpManager.TransactOpts)
}

// AddEthRewards is a paid mutator transaction binding the contract method 0x97bb441a.
//
// Solidity: function addEthRewards() payable returns()
func (_LpManager *LpManagerTransactorSession) AddEthRewards() (*types.Transaction, error) {
	return _LpManager.Contract.AddEthRewards(&_LpManager.TransactOpts)
}

// AddLiquidity is a paid mutator transaction binding the contract method 0x1b879378.
//
// Solidity: function addLiquidity(address token, uint256 amount, address to) returns(uint256 liquidity)
func (_LpManager *LpManagerTransactor) AddLiquidity(opts *bind.TransactOpts, token common.Address, amount *big.Int, to common.Address) (*types.Transaction, error) {
	return _LpManager.contract.Transact(opts, "addLiquidity", token, amount, to)
}

// AddLiquidity is a paid mutator transaction binding the contract method 0x1b879378.
//
// Solidity: function addLiquidity(address token, uint256 amount, address to) returns(uint256 liquidity)
func (_LpManager *LpManagerSession) AddLiquidity(token common.Address, amount *big.Int, to common.Address) (*types.Transaction, error) {
	return _LpManager.Contract.AddLiquidity(&_LpManager.TransactOpts, token, amount, to)
}

// AddLiquidity is a paid mutator transaction binding the contract method 0x1b879378.
//
// Solidity: function addLiquidity(address token, uint256 amount, address to) returns(uint256 liquidity)
func (_LpManager *LpManagerTransactorSession) AddLiquidity(token common.Address, amount *big.Int, to common.Address) (*types.Transaction, error) {
	return _LpManager.Contract.AddLiquidity(&_LpManager.TransactOpts, token, amount, to)
}

// AddRewards is a paid mutator transaction binding the contract method 0xa9fc507b.
//
// Solidity: function addRewards(address token, uint256 amount) returns()
func (_LpManager *LpManagerTransactor) AddRewards(opts *bind.TransactOpts, token common.Address, amount *big.Int) (*types.Transaction, error) {
	return _LpManager.contract.Transact(opts, "addRewards", token, amount)
}

// AddRewards is a paid mutator transaction binding the contract method 0xa9fc507b.
//
// Solidity: function addRewards(address token, uint256 amount) returns()
func (_LpManager *LpManagerSession) AddRewards(token common.Address, amount *big.Int) (*types.Transaction, error) {
	return _LpManager.Contract.AddRewards(&_LpManager.TransactOpts, token, amount)
}

// AddRewards is a paid mutator transaction binding the contract method 0xa9fc507b.
//
// Solidity: function addRewards(address token, uint256 amount) returns()
func (_LpManager *LpManagerTransactorSession) AddRewards(token common.Address, amount *big.Int) (*types.Transaction, error) {
	return _LpManager.Contract.AddRewards(&_LpManager.TransactOpts, token, amount)
}

// CreateEthPool is a paid mutator transaction binding the contract method 0x6ad23d8c.
//
// Solidity: function createEthPool() returns(address pool)
func (_LpManager *LpManagerTransactor) CreateEthPool(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LpManager.contract.Transact(opts, "createEthPool")
}

// CreateEthPool is a paid mutator transaction binding the contract method 0x6ad23d8c.
//
// Solidity: function createEthPool() returns(address pool)
func (_LpManager *LpManagerSession) CreateEthPool() (*types.Transaction, error) {
	return _LpManager.Contract.CreateEthPool(&_LpManager.TransactOpts)
}

// CreateEthPool is a paid mutator transaction binding the contract method 0x6ad23d8c.
//
// Solidity: function createEthPool() returns(address pool)
func (_LpManager *LpManagerTransactorSession) CreateEthPool() (*types.Transaction, error) {
	return _LpManager.Contract.CreateEthPool(&_LpManager.TransactOpts)
}

// CreatePool is a paid mutator transaction binding the contract method 0x9049f9d2.
//
// Solidity: function createPool(address token) returns(address pool)
func (_LpManager *LpManagerTransactor) CreatePool(opts *bind.TransactOpts, token common.Address) (*types.Transaction, error) {
	return _LpManager.contract.Transact(opts, "createPool", token)
}

// CreatePool is a paid mutator transaction binding the contract method 0x9049f9d2.
//
// Solidity: function createPool(address token) returns(address pool)
func (_LpManager *LpManagerSession) CreatePool(token common.Address) (*types.Transaction, error) {
	return _LpManager.Contract.CreatePool(&_LpManager.TransactOpts, token)
}

// CreatePool is a paid mutator transaction binding the contract method 0x9049f9d2.
//
// Solidity: function createPool(address token) returns(address pool)
func (_LpManager *LpManagerTransactorSession) CreatePool(token common.Address) (*types.Transaction, error) {
	return _LpManager.Contract.CreatePool(&_LpManager.TransactOpts, token)
}

// Deposit is a paid mutator transaction binding the contract method 0x47e7ef24.
//
// Solidity: function deposit(address token, uint256 amount) returns()
func (_LpManager *LpManagerTransactor) Deposit(opts *bind.TransactOpts, token common.Address, amount *big.Int) (*types.Transaction, error) {
	return _LpManager.contract.Transact(opts, "deposit", token, amount)
}

// Deposit is a paid mutator transaction binding the contract method 0x47e7ef24.
//
// Solidity: function deposit(address token, uint256 amount) returns()
func (_LpManager *LpManagerSession) Deposit(token common.Address, amount *big.Int) (*types.Transaction, error) {
	return _LpManager.Contract.Deposit(&_LpManager.TransactOpts, token, amount)
}

// Deposit is a paid mutator transaction binding the contract method 0x47e7ef24.
//
// Solidity: function deposit(address token, uint256 amount) returns()
func (_LpManager *LpManagerTransactorSession) Deposit(token common.Address, amount *big.Int) (*types.Transaction, error) {
	return _LpManager.Contract.Deposit(&_LpManager.TransactOpts, token, amount)
}

// DepositEth is a paid mutator transaction binding the contract method 0x439370b1.
//
// Solidity: function depositEth() payable returns()
func (_LpManager *LpManagerTransactor) DepositEth(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LpManager.contract.Transact(opts, "depositEth")
}

// DepositEth is a paid mutator transaction binding the contract method 0x439370b1.
//
// Solidity: function depositEth() payable returns()
func (_LpManager *LpManagerSession) DepositEth() (*types.Transaction, error) {
	return _LpManager.Contract.DepositEth(&_LpManager.TransactOpts)
}

// DepositEth is a paid mutator transaction binding the contract method 0x439370b1.
//
// Solidity: function depositEth() payable returns()
func (_LpManager *LpManagerTransactorSession) DepositEth() (*types.Transaction, error) {
	return _LpManager.Contract.DepositEth(&_LpManager.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_LpManager *LpManagerTransactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LpManager.contract.Transact(opts, "pause")
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_LpManager *LpManagerSession) Pause() (*types.Transaction, error) {
	return _LpManager.Contract.Pause(&_LpManager.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_LpManager *LpManagerTransactorSession) Pause() (*types.Transaction, error) {
	return _LpManager.Contract.Pause(&_LpManager.TransactOpts)
}

// PauseEth is a paid mutator transaction binding the contract method 0x2998fc25.
//
// Solidity: function pauseEth() returns()
func (_LpManager *LpManagerTransactor) PauseEth(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LpManager.contract.Transact(opts, "pauseEth")
}

// PauseEth is a paid mutator transaction binding the contract method 0x2998fc25.
//
// Solidity: function pauseEth() returns()
func (_LpManager *LpManagerSession) PauseEth() (*types.Transaction, error) {
	return _LpManager.Contract.PauseEth(&_LpManager.TransactOpts)
}

// PauseEth is a paid mutator transaction binding the contract method 0x2998fc25.
//
// Solidity: function pauseEth() returns()
func (_LpManager *LpManagerTransactorSession) PauseEth() (*types.Transaction, error) {
	return _LpManager.Contract.PauseEth(&_LpManager.TransactOpts)
}

// PauseToken is a paid mutator transaction binding the contract method 0x7c41ad2c.
//
// Solidity: function pauseToken(address token) returns()
func (_LpManager *LpManagerTransactor) PauseToken(opts *bind.TransactOpts, token common.Address) (*types.Transaction, error) {
	return _LpManager.contract.Transact(opts, "pauseToken", token)
}

// PauseToken is a paid mutator transaction binding the contract method 0x7c41ad2c.
//
// Solidity: function pauseToken(address token) returns()
func (_LpManager *LpManagerSession) PauseToken(token common.Address) (*types.Transaction, error) {
	return _LpManager.Contract.PauseToken(&_LpManager.TransactOpts, token)
}

// PauseToken is a paid mutator transaction binding the contract method 0x7c41ad2c.
//
// Solidity: function pauseToken(address token) returns()
func (_LpManager *LpManagerTransactorSession) PauseToken(token common.Address) (*types.Transaction, error) {
	return _LpManager.Contract.PauseToken(&_LpManager.TransactOpts, token)
}

// RemoveEthLiquidity is a paid mutator transaction binding the contract method 0x04800ea8.
//
// Solidity: function removeEthLiquidity(uint256 liquidity, address to) returns(uint256 amount)
func (_LpManager *LpManagerTransactor) RemoveEthLiquidity(opts *bind.TransactOpts, liquidity *big.Int, to common.Address) (*types.Transaction, error) {
	return _LpManager.contract.Transact(opts, "removeEthLiquidity", liquidity, to)
}

// RemoveEthLiquidity is a paid mutator transaction binding the contract method 0x04800ea8.
//
// Solidity: function removeEthLiquidity(uint256 liquidity, address to) returns(uint256 amount)
func (_LpManager *LpManagerSession) RemoveEthLiquidity(liquidity *big.Int, to common.Address) (*types.Transaction, error) {
	return _LpManager.Contract.RemoveEthLiquidity(&_LpManager.TransactOpts, liquidity, to)
}

// RemoveEthLiquidity is a paid mutator transaction binding the contract method 0x04800ea8.
//
// Solidity: function removeEthLiquidity(uint256 liquidity, address to) returns(uint256 amount)
func (_LpManager *LpManagerTransactorSession) RemoveEthLiquidity(liquidity *big.Int, to common.Address) (*types.Transaction, error) {
	return _LpManager.Contract.RemoveEthLiquidity(&_LpManager.TransactOpts, liquidity, to)
}

// RemoveLiquidity is a paid mutator transaction binding the contract method 0x47fdbc8e.
//
// Solidity: function removeLiquidity(address token, uint256 liquidity, address to) returns(uint256 amount)
func (_LpManager *LpManagerTransactor) RemoveLiquidity(opts *bind.TransactOpts, token common.Address, liquidity *big.Int, to common.Address) (*types.Transaction, error) {
	return _LpManager.contract.Transact(opts, "removeLiquidity", token, liquidity, to)
}

// RemoveLiquidity is a paid mutator transaction binding the contract method 0x47fdbc8e.
//
// Solidity: function removeLiquidity(address token, uint256 liquidity, address to) returns(uint256 amount)
func (_LpManager *LpManagerSession) RemoveLiquidity(token common.Address, liquidity *big.Int, to common.Address) (*types.Transaction, error) {
	return _LpManager.Contract.RemoveLiquidity(&_LpManager.TransactOpts, token, liquidity, to)
}

// RemoveLiquidity is a paid mutator transaction binding the contract method 0x47fdbc8e.
//
// Solidity: function removeLiquidity(address token, uint256 liquidity, address to) returns(uint256 amount)
func (_LpManager *LpManagerTransactorSession) RemoveLiquidity(token common.Address, liquidity *big.Int, to common.Address) (*types.Transaction, error) {
	return _LpManager.Contract.RemoveLiquidity(&_LpManager.TransactOpts, token, liquidity, to)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_LpManager *LpManagerTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LpManager.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_LpManager *LpManagerSession) RenounceOwnership() (*types.Transaction, error) {
	return _LpManager.Contract.RenounceOwnership(&_LpManager.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_LpManager *LpManagerTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _LpManager.Contract.RenounceOwnership(&_LpManager.TransactOpts)
}

// SetBackend is a paid mutator transaction binding the contract method 0xda7fc24f.
//
// Solidity: function setBackend(address backend_) returns()
func (_LpManager *LpManagerTransactor) SetBackend(opts *bind.TransactOpts, backend_ common.Address) (*types.Transaction, error) {
	return _LpManager.contract.Transact(opts, "setBackend", backend_)
}

// SetBackend is a paid mutator transaction binding the contract method 0xda7fc24f.
//
// Solidity: function setBackend(address backend_) returns()
func (_LpManager *LpManagerSession) SetBackend(backend_ common.Address) (*types.Transaction, error) {
	return _LpManager.Contract.SetBackend(&_LpManager.TransactOpts, backend_)
}

// SetBackend is a paid mutator transaction binding the contract method 0xda7fc24f.
//
// Solidity: function setBackend(address backend_) returns()
func (_LpManager *LpManagerTransactorSession) SetBackend(backend_ common.Address) (*types.Transaction, error) {
	return _LpManager.Contract.SetBackend(&_LpManager.TransactOpts, backend_)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_LpManager *LpManagerTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _LpManager.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_LpManager *LpManagerSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _LpManager.Contract.TransferOwnership(&_LpManager.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_LpManager *LpManagerTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _LpManager.Contract.TransferOwnership(&_LpManager.TransactOpts, newOwner)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_LpManager *LpManagerTransactor) Unpause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LpManager.contract.Transact(opts, "unpause")
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_LpManager *LpManagerSession) Unpause() (*types.Transaction, error) {
	return _LpManager.Contract.Unpause(&_LpManager.TransactOpts)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_LpManager *LpManagerTransactorSession) Unpause() (*types.Transaction, error) {
	return _LpManager.Contract.Unpause(&_LpManager.TransactOpts)
}

// UnpauseEth is a paid mutator transaction binding the contract method 0x85b5bf9e.
//
// Solidity: function unpauseEth() returns()
func (_LpManager *LpManagerTransactor) UnpauseEth(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LpManager.contract.Transact(opts, "unpauseEth")
}

// UnpauseEth is a paid mutator transaction binding the contract method 0x85b5bf9e.
//
// Solidity: function unpauseEth() returns()
func (_LpManager *LpManagerSession) UnpauseEth() (*types.Transaction, error) {
	return _LpManager.Contract.UnpauseEth(&_LpManager.TransactOpts)
}

// UnpauseEth is a paid mutator transaction binding the contract method 0x85b5bf9e.
//
// Solidity: function unpauseEth() returns()
func (_LpManager *LpManagerTransactorSession) UnpauseEth() (*types.Transaction, error) {
	return _LpManager.Contract.UnpauseEth(&_LpManager.TransactOpts)
}

// UnpauseToken is a paid mutator transaction binding the contract method 0x3b3bff0f.
//
// Solidity: function unpauseToken(address token) returns()
func (_LpManager *LpManagerTransactor) UnpauseToken(opts *bind.TransactOpts, token common.Address) (*types.Transaction, error) {
	return _LpManager.contract.Transact(opts, "unpauseToken", token)
}

// UnpauseToken is a paid mutator transaction binding the contract method 0x3b3bff0f.
//
// Solidity: function unpauseToken(address token) returns()
func (_LpManager *LpManagerSession) UnpauseToken(token common.Address) (*types.Transaction, error) {
	return _LpManager.Contract.UnpauseToken(&_LpManager.TransactOpts, token)
}

// UnpauseToken is a paid mutator transaction binding the contract method 0x3b3bff0f.
//
// Solidity: function unpauseToken(address token) returns()
func (_LpManager *LpManagerTransactorSession) UnpauseToken(token common.Address) (*types.Transaction, error) {
	return _LpManager.Contract.UnpauseToken(&_LpManager.TransactOpts, token)
}

// Withdraw is a paid mutator transaction binding the contract method 0xd9caed12.
//
// Solidity: function withdraw(address token, address to, uint256 amount) returns()
func (_LpManager *LpManagerTransactor) Withdraw(opts *bind.TransactOpts, token common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _LpManager.contract.Transact(opts, "withdraw", token, to, amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0xd9caed12.
//
// Solidity: function withdraw(address token, address to, uint256 amount) returns()
func (_LpManager *LpManagerSession) Withdraw(token common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _LpManager.Contract.Withdraw(&_LpManager.TransactOpts, token, to, amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0xd9caed12.
//
// Solidity: function withdraw(address token, address to, uint256 amount) returns()
func (_LpManager *LpManagerTransactorSession) Withdraw(token common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _LpManager.Contract.Withdraw(&_LpManager.TransactOpts, token, to, amount)
}

// WithdrawEth is a paid mutator transaction binding the contract method 0x1b9a91a4.
//
// Solidity: function withdrawEth(address to, uint256 amount) returns()
func (_LpManager *LpManagerTransactor) WithdrawEth(opts *bind.TransactOpts, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _LpManager.contract.Transact(opts, "withdrawEth", to, amount)
}

// WithdrawEth is a paid mutator transaction binding the contract method 0x1b9a91a4.
//
// Solidity: function withdrawEth(address to, uint256 amount) returns()
func (_LpManager *LpManagerSession) WithdrawEth(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _LpManager.Contract.WithdrawEth(&_LpManager.TransactOpts, to, amount)
}

// WithdrawEth is a paid mutator transaction binding the contract method 0x1b9a91a4.
//
// Solidity: function withdrawEth(address to, uint256 amount) returns()
func (_LpManager *LpManagerTransactorSession) WithdrawEth(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _LpManager.Contract.WithdrawEth(&_LpManager.TransactOpts, to, amount)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_LpManager *LpManagerTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LpManager.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_LpManager *LpManagerSession) Receive() (*types.Transaction, error) {
	return _LpManager.Contract.Receive(&_LpManager.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_LpManager *LpManagerTransactorSession) Receive() (*types.Transaction, error) {
	return _LpManager.Contract.Receive(&_LpManager.TransactOpts)
}

// LpManagerEthLiquidityAddedIterator is returned from FilterEthLiquidityAdded and is used to iterate over the raw logs and unpacked data for EthLiquidityAdded events raised by the LpManager contract.
type LpManagerEthLiquidityAddedIterator struct {
	Event *LpManagerEthLiquidityAdded // Event containing the contract specifics and raw log

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
func (it *LpManagerEthLiquidityAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LpManagerEthLiquidityAdded)
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
		it.Event = new(LpManagerEthLiquidityAdded)
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
func (it *LpManagerEthLiquidityAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LpManagerEthLiquidityAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LpManagerEthLiquidityAdded represents a EthLiquidityAdded event raised by the LpManager contract.
type LpManagerEthLiquidityAdded struct {
	To   common.Address
	Arg1 *big.Int
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterEthLiquidityAdded is a free log retrieval operation binding the contract event 0x20da65354ff194c1eee93e0304f7a245c8e27a00c82810ac7f71954cdf3acbbd.
//
// Solidity: event EthLiquidityAdded(address to, uint256 arg1)
func (_LpManager *LpManagerFilterer) FilterEthLiquidityAdded(opts *bind.FilterOpts) (*LpManagerEthLiquidityAddedIterator, error) {

	logs, sub, err := _LpManager.contract.FilterLogs(opts, "EthLiquidityAdded")
	if err != nil {
		return nil, err
	}
	return &LpManagerEthLiquidityAddedIterator{contract: _LpManager.contract, event: "EthLiquidityAdded", logs: logs, sub: sub}, nil
}

// WatchEthLiquidityAdded is a free log subscription operation binding the contract event 0x20da65354ff194c1eee93e0304f7a245c8e27a00c82810ac7f71954cdf3acbbd.
//
// Solidity: event EthLiquidityAdded(address to, uint256 arg1)
func (_LpManager *LpManagerFilterer) WatchEthLiquidityAdded(opts *bind.WatchOpts, sink chan<- *LpManagerEthLiquidityAdded) (event.Subscription, error) {

	logs, sub, err := _LpManager.contract.WatchLogs(opts, "EthLiquidityAdded")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LpManagerEthLiquidityAdded)
				if err := _LpManager.contract.UnpackLog(event, "EthLiquidityAdded", log); err != nil {
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

// ParseEthLiquidityAdded is a log parse operation binding the contract event 0x20da65354ff194c1eee93e0304f7a245c8e27a00c82810ac7f71954cdf3acbbd.
//
// Solidity: event EthLiquidityAdded(address to, uint256 arg1)
func (_LpManager *LpManagerFilterer) ParseEthLiquidityAdded(log types.Log) (*LpManagerEthLiquidityAdded, error) {
	event := new(LpManagerEthLiquidityAdded)
	if err := _LpManager.contract.UnpackLog(event, "EthLiquidityAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LpManagerEthLiquidityRemovedIterator is returned from FilterEthLiquidityRemoved and is used to iterate over the raw logs and unpacked data for EthLiquidityRemoved events raised by the LpManager contract.
type LpManagerEthLiquidityRemovedIterator struct {
	Event *LpManagerEthLiquidityRemoved // Event containing the contract specifics and raw log

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
func (it *LpManagerEthLiquidityRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LpManagerEthLiquidityRemoved)
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
		it.Event = new(LpManagerEthLiquidityRemoved)
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
func (it *LpManagerEthLiquidityRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LpManagerEthLiquidityRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LpManagerEthLiquidityRemoved represents a EthLiquidityRemoved event raised by the LpManager contract.
type LpManagerEthLiquidityRemoved struct {
	To   common.Address
	Arg1 *big.Int
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterEthLiquidityRemoved is a free log retrieval operation binding the contract event 0x5f158a454ae2efda2e9bdb12fb0c9fbd5c394dd5e395880d1d9b03420cb8b456.
//
// Solidity: event EthLiquidityRemoved(address to, uint256 arg1)
func (_LpManager *LpManagerFilterer) FilterEthLiquidityRemoved(opts *bind.FilterOpts) (*LpManagerEthLiquidityRemovedIterator, error) {

	logs, sub, err := _LpManager.contract.FilterLogs(opts, "EthLiquidityRemoved")
	if err != nil {
		return nil, err
	}
	return &LpManagerEthLiquidityRemovedIterator{contract: _LpManager.contract, event: "EthLiquidityRemoved", logs: logs, sub: sub}, nil
}

// WatchEthLiquidityRemoved is a free log subscription operation binding the contract event 0x5f158a454ae2efda2e9bdb12fb0c9fbd5c394dd5e395880d1d9b03420cb8b456.
//
// Solidity: event EthLiquidityRemoved(address to, uint256 arg1)
func (_LpManager *LpManagerFilterer) WatchEthLiquidityRemoved(opts *bind.WatchOpts, sink chan<- *LpManagerEthLiquidityRemoved) (event.Subscription, error) {

	logs, sub, err := _LpManager.contract.WatchLogs(opts, "EthLiquidityRemoved")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LpManagerEthLiquidityRemoved)
				if err := _LpManager.contract.UnpackLog(event, "EthLiquidityRemoved", log); err != nil {
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

// ParseEthLiquidityRemoved is a log parse operation binding the contract event 0x5f158a454ae2efda2e9bdb12fb0c9fbd5c394dd5e395880d1d9b03420cb8b456.
//
// Solidity: event EthLiquidityRemoved(address to, uint256 arg1)
func (_LpManager *LpManagerFilterer) ParseEthLiquidityRemoved(log types.Log) (*LpManagerEthLiquidityRemoved, error) {
	event := new(LpManagerEthLiquidityRemoved)
	if err := _LpManager.contract.UnpackLog(event, "EthLiquidityRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LpManagerEthLiquidityWithdrawIterator is returned from FilterEthLiquidityWithdraw and is used to iterate over the raw logs and unpacked data for EthLiquidityWithdraw events raised by the LpManager contract.
type LpManagerEthLiquidityWithdrawIterator struct {
	Event *LpManagerEthLiquidityWithdraw // Event containing the contract specifics and raw log

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
func (it *LpManagerEthLiquidityWithdrawIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LpManagerEthLiquidityWithdraw)
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
		it.Event = new(LpManagerEthLiquidityWithdraw)
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
func (it *LpManagerEthLiquidityWithdrawIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LpManagerEthLiquidityWithdrawIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LpManagerEthLiquidityWithdraw represents a EthLiquidityWithdraw event raised by the LpManager contract.
type LpManagerEthLiquidityWithdraw struct {
	To     common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterEthLiquidityWithdraw is a free log retrieval operation binding the contract event 0xb2d40582c20f4fbc7f44a3358744b379dba51f3fcc1785b58471ac76eec2cd34.
//
// Solidity: event EthLiquidityWithdraw(address to, uint256 amount)
func (_LpManager *LpManagerFilterer) FilterEthLiquidityWithdraw(opts *bind.FilterOpts) (*LpManagerEthLiquidityWithdrawIterator, error) {

	logs, sub, err := _LpManager.contract.FilterLogs(opts, "EthLiquidityWithdraw")
	if err != nil {
		return nil, err
	}
	return &LpManagerEthLiquidityWithdrawIterator{contract: _LpManager.contract, event: "EthLiquidityWithdraw", logs: logs, sub: sub}, nil
}

// WatchEthLiquidityWithdraw is a free log subscription operation binding the contract event 0xb2d40582c20f4fbc7f44a3358744b379dba51f3fcc1785b58471ac76eec2cd34.
//
// Solidity: event EthLiquidityWithdraw(address to, uint256 amount)
func (_LpManager *LpManagerFilterer) WatchEthLiquidityWithdraw(opts *bind.WatchOpts, sink chan<- *LpManagerEthLiquidityWithdraw) (event.Subscription, error) {

	logs, sub, err := _LpManager.contract.WatchLogs(opts, "EthLiquidityWithdraw")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LpManagerEthLiquidityWithdraw)
				if err := _LpManager.contract.UnpackLog(event, "EthLiquidityWithdraw", log); err != nil {
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

// ParseEthLiquidityWithdraw is a log parse operation binding the contract event 0xb2d40582c20f4fbc7f44a3358744b379dba51f3fcc1785b58471ac76eec2cd34.
//
// Solidity: event EthLiquidityWithdraw(address to, uint256 amount)
func (_LpManager *LpManagerFilterer) ParseEthLiquidityWithdraw(log types.Log) (*LpManagerEthLiquidityWithdraw, error) {
	event := new(LpManagerEthLiquidityWithdraw)
	if err := _LpManager.contract.UnpackLog(event, "EthLiquidityWithdraw", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LpManagerEthPoolCreatedIterator is returned from FilterEthPoolCreated and is used to iterate over the raw logs and unpacked data for EthPoolCreated events raised by the LpManager contract.
type LpManagerEthPoolCreatedIterator struct {
	Event *LpManagerEthPoolCreated // Event containing the contract specifics and raw log

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
func (it *LpManagerEthPoolCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LpManagerEthPoolCreated)
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
		it.Event = new(LpManagerEthPoolCreated)
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
func (it *LpManagerEthPoolCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LpManagerEthPoolCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LpManagerEthPoolCreated represents a EthPoolCreated event raised by the LpManager contract.
type LpManagerEthPoolCreated struct {
	Pool common.Address
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterEthPoolCreated is a free log retrieval operation binding the contract event 0x4d0a5aaf2e8c528841acdb365f2a22216147bd5bc7dfacfa74ff362abc08fb52.
//
// Solidity: event EthPoolCreated(address pool)
func (_LpManager *LpManagerFilterer) FilterEthPoolCreated(opts *bind.FilterOpts) (*LpManagerEthPoolCreatedIterator, error) {

	logs, sub, err := _LpManager.contract.FilterLogs(opts, "EthPoolCreated")
	if err != nil {
		return nil, err
	}
	return &LpManagerEthPoolCreatedIterator{contract: _LpManager.contract, event: "EthPoolCreated", logs: logs, sub: sub}, nil
}

// WatchEthPoolCreated is a free log subscription operation binding the contract event 0x4d0a5aaf2e8c528841acdb365f2a22216147bd5bc7dfacfa74ff362abc08fb52.
//
// Solidity: event EthPoolCreated(address pool)
func (_LpManager *LpManagerFilterer) WatchEthPoolCreated(opts *bind.WatchOpts, sink chan<- *LpManagerEthPoolCreated) (event.Subscription, error) {

	logs, sub, err := _LpManager.contract.WatchLogs(opts, "EthPoolCreated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LpManagerEthPoolCreated)
				if err := _LpManager.contract.UnpackLog(event, "EthPoolCreated", log); err != nil {
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

// ParseEthPoolCreated is a log parse operation binding the contract event 0x4d0a5aaf2e8c528841acdb365f2a22216147bd5bc7dfacfa74ff362abc08fb52.
//
// Solidity: event EthPoolCreated(address pool)
func (_LpManager *LpManagerFilterer) ParseEthPoolCreated(log types.Log) (*LpManagerEthPoolCreated, error) {
	event := new(LpManagerEthPoolCreated)
	if err := _LpManager.contract.UnpackLog(event, "EthPoolCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LpManagerEthRewardsAddedIterator is returned from FilterEthRewardsAdded and is used to iterate over the raw logs and unpacked data for EthRewardsAdded events raised by the LpManager contract.
type LpManagerEthRewardsAddedIterator struct {
	Event *LpManagerEthRewardsAdded // Event containing the contract specifics and raw log

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
func (it *LpManagerEthRewardsAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LpManagerEthRewardsAdded)
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
		it.Event = new(LpManagerEthRewardsAdded)
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
func (it *LpManagerEthRewardsAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LpManagerEthRewardsAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LpManagerEthRewardsAdded represents a EthRewardsAdded event raised by the LpManager contract.
type LpManagerEthRewardsAdded struct {
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterEthRewardsAdded is a free log retrieval operation binding the contract event 0x66867bbbbfcb19b9f8b8b7141e1ffe8a0ef5c91c7feea74546eeef97942fdce7.
//
// Solidity: event EthRewardsAdded(uint256 amount)
func (_LpManager *LpManagerFilterer) FilterEthRewardsAdded(opts *bind.FilterOpts) (*LpManagerEthRewardsAddedIterator, error) {

	logs, sub, err := _LpManager.contract.FilterLogs(opts, "EthRewardsAdded")
	if err != nil {
		return nil, err
	}
	return &LpManagerEthRewardsAddedIterator{contract: _LpManager.contract, event: "EthRewardsAdded", logs: logs, sub: sub}, nil
}

// WatchEthRewardsAdded is a free log subscription operation binding the contract event 0x66867bbbbfcb19b9f8b8b7141e1ffe8a0ef5c91c7feea74546eeef97942fdce7.
//
// Solidity: event EthRewardsAdded(uint256 amount)
func (_LpManager *LpManagerFilterer) WatchEthRewardsAdded(opts *bind.WatchOpts, sink chan<- *LpManagerEthRewardsAdded) (event.Subscription, error) {

	logs, sub, err := _LpManager.contract.WatchLogs(opts, "EthRewardsAdded")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LpManagerEthRewardsAdded)
				if err := _LpManager.contract.UnpackLog(event, "EthRewardsAdded", log); err != nil {
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

// ParseEthRewardsAdded is a log parse operation binding the contract event 0x66867bbbbfcb19b9f8b8b7141e1ffe8a0ef5c91c7feea74546eeef97942fdce7.
//
// Solidity: event EthRewardsAdded(uint256 amount)
func (_LpManager *LpManagerFilterer) ParseEthRewardsAdded(log types.Log) (*LpManagerEthRewardsAdded, error) {
	event := new(LpManagerEthRewardsAdded)
	if err := _LpManager.contract.UnpackLog(event, "EthRewardsAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LpManagerLiquidityAddedIterator is returned from FilterLiquidityAdded and is used to iterate over the raw logs and unpacked data for LiquidityAdded events raised by the LpManager contract.
type LpManagerLiquidityAddedIterator struct {
	Event *LpManagerLiquidityAdded // Event containing the contract specifics and raw log

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
func (it *LpManagerLiquidityAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LpManagerLiquidityAdded)
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
		it.Event = new(LpManagerLiquidityAdded)
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
func (it *LpManagerLiquidityAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LpManagerLiquidityAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LpManagerLiquidityAdded represents a LiquidityAdded event raised by the LpManager contract.
type LpManagerLiquidityAdded struct {
	Token common.Address
	To    common.Address
	Arg2  *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterLiquidityAdded is a free log retrieval operation binding the contract event 0x9d278c56ba6dc86a12eefe6b43112bd6e06648eb4ec0b950ee2d783d40e2acb4.
//
// Solidity: event LiquidityAdded(address indexed token, address to, uint256 arg2)
func (_LpManager *LpManagerFilterer) FilterLiquidityAdded(opts *bind.FilterOpts, token []common.Address) (*LpManagerLiquidityAddedIterator, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _LpManager.contract.FilterLogs(opts, "LiquidityAdded", tokenRule)
	if err != nil {
		return nil, err
	}
	return &LpManagerLiquidityAddedIterator{contract: _LpManager.contract, event: "LiquidityAdded", logs: logs, sub: sub}, nil
}

// WatchLiquidityAdded is a free log subscription operation binding the contract event 0x9d278c56ba6dc86a12eefe6b43112bd6e06648eb4ec0b950ee2d783d40e2acb4.
//
// Solidity: event LiquidityAdded(address indexed token, address to, uint256 arg2)
func (_LpManager *LpManagerFilterer) WatchLiquidityAdded(opts *bind.WatchOpts, sink chan<- *LpManagerLiquidityAdded, token []common.Address) (event.Subscription, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _LpManager.contract.WatchLogs(opts, "LiquidityAdded", tokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LpManagerLiquidityAdded)
				if err := _LpManager.contract.UnpackLog(event, "LiquidityAdded", log); err != nil {
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

// ParseLiquidityAdded is a log parse operation binding the contract event 0x9d278c56ba6dc86a12eefe6b43112bd6e06648eb4ec0b950ee2d783d40e2acb4.
//
// Solidity: event LiquidityAdded(address indexed token, address to, uint256 arg2)
func (_LpManager *LpManagerFilterer) ParseLiquidityAdded(log types.Log) (*LpManagerLiquidityAdded, error) {
	event := new(LpManagerLiquidityAdded)
	if err := _LpManager.contract.UnpackLog(event, "LiquidityAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LpManagerLiquidityRemovedIterator is returned from FilterLiquidityRemoved and is used to iterate over the raw logs and unpacked data for LiquidityRemoved events raised by the LpManager contract.
type LpManagerLiquidityRemovedIterator struct {
	Event *LpManagerLiquidityRemoved // Event containing the contract specifics and raw log

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
func (it *LpManagerLiquidityRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LpManagerLiquidityRemoved)
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
		it.Event = new(LpManagerLiquidityRemoved)
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
func (it *LpManagerLiquidityRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LpManagerLiquidityRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LpManagerLiquidityRemoved represents a LiquidityRemoved event raised by the LpManager contract.
type LpManagerLiquidityRemoved struct {
	Token common.Address
	To    common.Address
	Arg2  *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterLiquidityRemoved is a free log retrieval operation binding the contract event 0x983e86fda8e7b1e2eae380201830eaf1cac55772e8e39583da349865e8178863.
//
// Solidity: event LiquidityRemoved(address indexed token, address to, uint256 arg2)
func (_LpManager *LpManagerFilterer) FilterLiquidityRemoved(opts *bind.FilterOpts, token []common.Address) (*LpManagerLiquidityRemovedIterator, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _LpManager.contract.FilterLogs(opts, "LiquidityRemoved", tokenRule)
	if err != nil {
		return nil, err
	}
	return &LpManagerLiquidityRemovedIterator{contract: _LpManager.contract, event: "LiquidityRemoved", logs: logs, sub: sub}, nil
}

// WatchLiquidityRemoved is a free log subscription operation binding the contract event 0x983e86fda8e7b1e2eae380201830eaf1cac55772e8e39583da349865e8178863.
//
// Solidity: event LiquidityRemoved(address indexed token, address to, uint256 arg2)
func (_LpManager *LpManagerFilterer) WatchLiquidityRemoved(opts *bind.WatchOpts, sink chan<- *LpManagerLiquidityRemoved, token []common.Address) (event.Subscription, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _LpManager.contract.WatchLogs(opts, "LiquidityRemoved", tokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LpManagerLiquidityRemoved)
				if err := _LpManager.contract.UnpackLog(event, "LiquidityRemoved", log); err != nil {
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

// ParseLiquidityRemoved is a log parse operation binding the contract event 0x983e86fda8e7b1e2eae380201830eaf1cac55772e8e39583da349865e8178863.
//
// Solidity: event LiquidityRemoved(address indexed token, address to, uint256 arg2)
func (_LpManager *LpManagerFilterer) ParseLiquidityRemoved(log types.Log) (*LpManagerLiquidityRemoved, error) {
	event := new(LpManagerLiquidityRemoved)
	if err := _LpManager.contract.UnpackLog(event, "LiquidityRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LpManagerLiquidityWithdrawIterator is returned from FilterLiquidityWithdraw and is used to iterate over the raw logs and unpacked data for LiquidityWithdraw events raised by the LpManager contract.
type LpManagerLiquidityWithdrawIterator struct {
	Event *LpManagerLiquidityWithdraw // Event containing the contract specifics and raw log

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
func (it *LpManagerLiquidityWithdrawIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LpManagerLiquidityWithdraw)
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
		it.Event = new(LpManagerLiquidityWithdraw)
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
func (it *LpManagerLiquidityWithdrawIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LpManagerLiquidityWithdrawIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LpManagerLiquidityWithdraw represents a LiquidityWithdraw event raised by the LpManager contract.
type LpManagerLiquidityWithdraw struct {
	Token  common.Address
	To     common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterLiquidityWithdraw is a free log retrieval operation binding the contract event 0x48fee4d6a8b00b4d37f992ec58e6026e0aabed6343ae86d580c1d1fa003962a1.
//
// Solidity: event LiquidityWithdraw(address indexed token, address to, uint256 amount)
func (_LpManager *LpManagerFilterer) FilterLiquidityWithdraw(opts *bind.FilterOpts, token []common.Address) (*LpManagerLiquidityWithdrawIterator, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _LpManager.contract.FilterLogs(opts, "LiquidityWithdraw", tokenRule)
	if err != nil {
		return nil, err
	}
	return &LpManagerLiquidityWithdrawIterator{contract: _LpManager.contract, event: "LiquidityWithdraw", logs: logs, sub: sub}, nil
}

// WatchLiquidityWithdraw is a free log subscription operation binding the contract event 0x48fee4d6a8b00b4d37f992ec58e6026e0aabed6343ae86d580c1d1fa003962a1.
//
// Solidity: event LiquidityWithdraw(address indexed token, address to, uint256 amount)
func (_LpManager *LpManagerFilterer) WatchLiquidityWithdraw(opts *bind.WatchOpts, sink chan<- *LpManagerLiquidityWithdraw, token []common.Address) (event.Subscription, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _LpManager.contract.WatchLogs(opts, "LiquidityWithdraw", tokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LpManagerLiquidityWithdraw)
				if err := _LpManager.contract.UnpackLog(event, "LiquidityWithdraw", log); err != nil {
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

// ParseLiquidityWithdraw is a log parse operation binding the contract event 0x48fee4d6a8b00b4d37f992ec58e6026e0aabed6343ae86d580c1d1fa003962a1.
//
// Solidity: event LiquidityWithdraw(address indexed token, address to, uint256 amount)
func (_LpManager *LpManagerFilterer) ParseLiquidityWithdraw(log types.Log) (*LpManagerLiquidityWithdraw, error) {
	event := new(LpManagerLiquidityWithdraw)
	if err := _LpManager.contract.UnpackLog(event, "LiquidityWithdraw", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LpManagerOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the LpManager contract.
type LpManagerOwnershipTransferredIterator struct {
	Event *LpManagerOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *LpManagerOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LpManagerOwnershipTransferred)
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
		it.Event = new(LpManagerOwnershipTransferred)
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
func (it *LpManagerOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LpManagerOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LpManagerOwnershipTransferred represents a OwnershipTransferred event raised by the LpManager contract.
type LpManagerOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_LpManager *LpManagerFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*LpManagerOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _LpManager.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &LpManagerOwnershipTransferredIterator{contract: _LpManager.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_LpManager *LpManagerFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *LpManagerOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _LpManager.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LpManagerOwnershipTransferred)
				if err := _LpManager.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_LpManager *LpManagerFilterer) ParseOwnershipTransferred(log types.Log) (*LpManagerOwnershipTransferred, error) {
	event := new(LpManagerOwnershipTransferred)
	if err := _LpManager.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LpManagerPausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the LpManager contract.
type LpManagerPausedIterator struct {
	Event *LpManagerPaused // Event containing the contract specifics and raw log

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
func (it *LpManagerPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LpManagerPaused)
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
		it.Event = new(LpManagerPaused)
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
func (it *LpManagerPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LpManagerPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LpManagerPaused represents a Paused event raised by the LpManager contract.
type LpManagerPaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_LpManager *LpManagerFilterer) FilterPaused(opts *bind.FilterOpts) (*LpManagerPausedIterator, error) {

	logs, sub, err := _LpManager.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &LpManagerPausedIterator{contract: _LpManager.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_LpManager *LpManagerFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *LpManagerPaused) (event.Subscription, error) {

	logs, sub, err := _LpManager.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LpManagerPaused)
				if err := _LpManager.contract.UnpackLog(event, "Paused", log); err != nil {
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

// ParsePaused is a log parse operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_LpManager *LpManagerFilterer) ParsePaused(log types.Log) (*LpManagerPaused, error) {
	event := new(LpManagerPaused)
	if err := _LpManager.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LpManagerPoolCreatedIterator is returned from FilterPoolCreated and is used to iterate over the raw logs and unpacked data for PoolCreated events raised by the LpManager contract.
type LpManagerPoolCreatedIterator struct {
	Event *LpManagerPoolCreated // Event containing the contract specifics and raw log

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
func (it *LpManagerPoolCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LpManagerPoolCreated)
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
		it.Event = new(LpManagerPoolCreated)
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
func (it *LpManagerPoolCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LpManagerPoolCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LpManagerPoolCreated represents a PoolCreated event raised by the LpManager contract.
type LpManagerPoolCreated struct {
	Token common.Address
	Pool  common.Address
	Arg2  *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterPoolCreated is a free log retrieval operation binding the contract event 0xf8a0462f666b427ea753848be7e91f9ce413975906f6f39950be296ca9a4d524.
//
// Solidity: event PoolCreated(address indexed token, address pool, uint256 arg2)
func (_LpManager *LpManagerFilterer) FilterPoolCreated(opts *bind.FilterOpts, token []common.Address) (*LpManagerPoolCreatedIterator, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _LpManager.contract.FilterLogs(opts, "PoolCreated", tokenRule)
	if err != nil {
		return nil, err
	}
	return &LpManagerPoolCreatedIterator{contract: _LpManager.contract, event: "PoolCreated", logs: logs, sub: sub}, nil
}

// WatchPoolCreated is a free log subscription operation binding the contract event 0xf8a0462f666b427ea753848be7e91f9ce413975906f6f39950be296ca9a4d524.
//
// Solidity: event PoolCreated(address indexed token, address pool, uint256 arg2)
func (_LpManager *LpManagerFilterer) WatchPoolCreated(opts *bind.WatchOpts, sink chan<- *LpManagerPoolCreated, token []common.Address) (event.Subscription, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _LpManager.contract.WatchLogs(opts, "PoolCreated", tokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LpManagerPoolCreated)
				if err := _LpManager.contract.UnpackLog(event, "PoolCreated", log); err != nil {
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

// ParsePoolCreated is a log parse operation binding the contract event 0xf8a0462f666b427ea753848be7e91f9ce413975906f6f39950be296ca9a4d524.
//
// Solidity: event PoolCreated(address indexed token, address pool, uint256 arg2)
func (_LpManager *LpManagerFilterer) ParsePoolCreated(log types.Log) (*LpManagerPoolCreated, error) {
	event := new(LpManagerPoolCreated)
	if err := _LpManager.contract.UnpackLog(event, "PoolCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LpManagerRewardsAddedIterator is returned from FilterRewardsAdded and is used to iterate over the raw logs and unpacked data for RewardsAdded events raised by the LpManager contract.
type LpManagerRewardsAddedIterator struct {
	Event *LpManagerRewardsAdded // Event containing the contract specifics and raw log

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
func (it *LpManagerRewardsAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LpManagerRewardsAdded)
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
		it.Event = new(LpManagerRewardsAdded)
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
func (it *LpManagerRewardsAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LpManagerRewardsAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LpManagerRewardsAdded represents a RewardsAdded event raised by the LpManager contract.
type LpManagerRewardsAdded struct {
	Token  common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterRewardsAdded is a free log retrieval operation binding the contract event 0xac140162a56e01aec90ea4ff6eea27f60ca53066fd39d57664700f1a8589de9a.
//
// Solidity: event RewardsAdded(address indexed token, uint256 amount)
func (_LpManager *LpManagerFilterer) FilterRewardsAdded(opts *bind.FilterOpts, token []common.Address) (*LpManagerRewardsAddedIterator, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _LpManager.contract.FilterLogs(opts, "RewardsAdded", tokenRule)
	if err != nil {
		return nil, err
	}
	return &LpManagerRewardsAddedIterator{contract: _LpManager.contract, event: "RewardsAdded", logs: logs, sub: sub}, nil
}

// WatchRewardsAdded is a free log subscription operation binding the contract event 0xac140162a56e01aec90ea4ff6eea27f60ca53066fd39d57664700f1a8589de9a.
//
// Solidity: event RewardsAdded(address indexed token, uint256 amount)
func (_LpManager *LpManagerFilterer) WatchRewardsAdded(opts *bind.WatchOpts, sink chan<- *LpManagerRewardsAdded, token []common.Address) (event.Subscription, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _LpManager.contract.WatchLogs(opts, "RewardsAdded", tokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LpManagerRewardsAdded)
				if err := _LpManager.contract.UnpackLog(event, "RewardsAdded", log); err != nil {
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

// ParseRewardsAdded is a log parse operation binding the contract event 0xac140162a56e01aec90ea4ff6eea27f60ca53066fd39d57664700f1a8589de9a.
//
// Solidity: event RewardsAdded(address indexed token, uint256 amount)
func (_LpManager *LpManagerFilterer) ParseRewardsAdded(log types.Log) (*LpManagerRewardsAdded, error) {
	event := new(LpManagerRewardsAdded)
	if err := _LpManager.contract.UnpackLog(event, "RewardsAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LpManagerUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the LpManager contract.
type LpManagerUnpausedIterator struct {
	Event *LpManagerUnpaused // Event containing the contract specifics and raw log

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
func (it *LpManagerUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LpManagerUnpaused)
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
		it.Event = new(LpManagerUnpaused)
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
func (it *LpManagerUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LpManagerUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LpManagerUnpaused represents a Unpaused event raised by the LpManager contract.
type LpManagerUnpaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_LpManager *LpManagerFilterer) FilterUnpaused(opts *bind.FilterOpts) (*LpManagerUnpausedIterator, error) {

	logs, sub, err := _LpManager.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &LpManagerUnpausedIterator{contract: _LpManager.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_LpManager *LpManagerFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *LpManagerUnpaused) (event.Subscription, error) {

	logs, sub, err := _LpManager.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LpManagerUnpaused)
				if err := _LpManager.contract.UnpackLog(event, "Unpaused", log); err != nil {
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

// ParseUnpaused is a log parse operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_LpManager *LpManagerFilterer) ParseUnpaused(log types.Log) (*LpManagerUnpaused, error) {
	event := new(LpManagerUnpaused)
	if err := _LpManager.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
