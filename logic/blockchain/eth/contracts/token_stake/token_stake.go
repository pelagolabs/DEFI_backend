// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package token_stake

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

// TokenStakeMetaData contains all meta data concerning the TokenStake contract.
var TokenStakeMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"who\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"pid\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Deposit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"who\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"pid\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Withdraw\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"contractIERC20Upgradeable\",\"name\":\"lpToken_\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"allocPoint_\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"lastRewardBlock_\",\"type\":\"uint256\"}],\"name\":\"addLPPool\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"arsw\",\"outputs\":[{\"internalType\":\"contractIERC20Upgradeable\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_pid\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"depositLP\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20Upgradeable\",\"name\":\"arsw_\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"rewardPerBlock_\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"startBlock_\",\"type\":\"uint256\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"lpPoolInfo\",\"outputs\":[{\"internalType\":\"contractIERC20Upgradeable\",\"name\":\"lpToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"allocPoint\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"lastRewardBlock\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"accRewardPerShare\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_pid\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_user\",\"type\":\"address\"}],\"name\":\"pendingLPReward\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"pending\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenAddress_\",\"type\":\"address\"}],\"name\":\"pullFunds\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"pid_\",\"type\":\"uint256\"},{\"internalType\":\"contractIERC20Upgradeable\",\"name\":\"lpToken_\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"allocPoint_\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"lastRewardBlock_\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"accRewardPerShare_\",\"type\":\"uint256\"}],\"name\":\"resetLPPool\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"rewardPerBlock\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20Upgradeable\",\"name\":\"token_\",\"type\":\"address\"}],\"name\":\"setArswAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_rewardPerBlock\",\"type\":\"uint256\"}],\"name\":\"setRewardPerBlock\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"startBlock\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalLPAllocPoint\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unpause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_pid\",\"type\":\"uint256\"}],\"name\":\"updateLPPool\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"userLPInfo\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"rewardDebt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"firstDepositedTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"averageDepositedTime\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_pid\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"withdrawLP\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// TokenStakeABI is the input ABI used to generate the binding from.
// Deprecated: Use TokenStakeMetaData.ABI instead.
var TokenStakeABI = TokenStakeMetaData.ABI

// TokenStake is an auto generated Go binding around an Ethereum contract.
type TokenStake struct {
	TokenStakeCaller     // Read-only binding to the contract
	TokenStakeTransactor // Write-only binding to the contract
	TokenStakeFilterer   // Log filterer for contract events
}

// TokenStakeCaller is an auto generated read-only Go binding around an Ethereum contract.
type TokenStakeCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TokenStakeTransactor is an auto generated write-only Go binding around an Ethereum contract.
type TokenStakeTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TokenStakeFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type TokenStakeFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TokenStakeSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TokenStakeSession struct {
	Contract     *TokenStake       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TokenStakeCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TokenStakeCallerSession struct {
	Contract *TokenStakeCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// TokenStakeTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TokenStakeTransactorSession struct {
	Contract     *TokenStakeTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// TokenStakeRaw is an auto generated low-level Go binding around an Ethereum contract.
type TokenStakeRaw struct {
	Contract *TokenStake // Generic contract binding to access the raw methods on
}

// TokenStakeCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TokenStakeCallerRaw struct {
	Contract *TokenStakeCaller // Generic read-only contract binding to access the raw methods on
}

// TokenStakeTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TokenStakeTransactorRaw struct {
	Contract *TokenStakeTransactor // Generic write-only contract binding to access the raw methods on
}

// NewTokenStake creates a new instance of TokenStake, bound to a specific deployed contract.
func NewTokenStake(address common.Address, backend bind.ContractBackend) (*TokenStake, error) {
	contract, err := bindTokenStake(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &TokenStake{TokenStakeCaller: TokenStakeCaller{contract: contract}, TokenStakeTransactor: TokenStakeTransactor{contract: contract}, TokenStakeFilterer: TokenStakeFilterer{contract: contract}}, nil
}

// NewTokenStakeCaller creates a new read-only instance of TokenStake, bound to a specific deployed contract.
func NewTokenStakeCaller(address common.Address, caller bind.ContractCaller) (*TokenStakeCaller, error) {
	contract, err := bindTokenStake(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TokenStakeCaller{contract: contract}, nil
}

// NewTokenStakeTransactor creates a new write-only instance of TokenStake, bound to a specific deployed contract.
func NewTokenStakeTransactor(address common.Address, transactor bind.ContractTransactor) (*TokenStakeTransactor, error) {
	contract, err := bindTokenStake(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TokenStakeTransactor{contract: contract}, nil
}

// NewTokenStakeFilterer creates a new log filterer instance of TokenStake, bound to a specific deployed contract.
func NewTokenStakeFilterer(address common.Address, filterer bind.ContractFilterer) (*TokenStakeFilterer, error) {
	contract, err := bindTokenStake(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TokenStakeFilterer{contract: contract}, nil
}

// bindTokenStake binds a generic wrapper to an already deployed contract.
func bindTokenStake(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := TokenStakeMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TokenStake *TokenStakeRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TokenStake.Contract.TokenStakeCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TokenStake *TokenStakeRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TokenStake.Contract.TokenStakeTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TokenStake *TokenStakeRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TokenStake.Contract.TokenStakeTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TokenStake *TokenStakeCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TokenStake.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TokenStake *TokenStakeTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TokenStake.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TokenStake *TokenStakeTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TokenStake.Contract.contract.Transact(opts, method, params...)
}

// Arsw is a free data retrieval call binding the contract method 0xa57fb803.
//
// Solidity: function arsw() view returns(address)
func (_TokenStake *TokenStakeCaller) Arsw(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _TokenStake.contract.Call(opts, &out, "arsw")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Arsw is a free data retrieval call binding the contract method 0xa57fb803.
//
// Solidity: function arsw() view returns(address)
func (_TokenStake *TokenStakeSession) Arsw() (common.Address, error) {
	return _TokenStake.Contract.Arsw(&_TokenStake.CallOpts)
}

// Arsw is a free data retrieval call binding the contract method 0xa57fb803.
//
// Solidity: function arsw() view returns(address)
func (_TokenStake *TokenStakeCallerSession) Arsw() (common.Address, error) {
	return _TokenStake.Contract.Arsw(&_TokenStake.CallOpts)
}

// LpPoolInfo is a free data retrieval call binding the contract method 0x2ba6b6f2.
//
// Solidity: function lpPoolInfo(uint256 ) view returns(address lpToken, uint256 allocPoint, uint256 lastRewardBlock, uint256 accRewardPerShare)
func (_TokenStake *TokenStakeCaller) LpPoolInfo(opts *bind.CallOpts, arg0 *big.Int) (struct {
	LpToken           common.Address
	AllocPoint        *big.Int
	LastRewardBlock   *big.Int
	AccRewardPerShare *big.Int
}, error) {
	var out []interface{}
	err := _TokenStake.contract.Call(opts, &out, "lpPoolInfo", arg0)

	outstruct := new(struct {
		LpToken           common.Address
		AllocPoint        *big.Int
		LastRewardBlock   *big.Int
		AccRewardPerShare *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.LpToken = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.AllocPoint = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.LastRewardBlock = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.AccRewardPerShare = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// LpPoolInfo is a free data retrieval call binding the contract method 0x2ba6b6f2.
//
// Solidity: function lpPoolInfo(uint256 ) view returns(address lpToken, uint256 allocPoint, uint256 lastRewardBlock, uint256 accRewardPerShare)
func (_TokenStake *TokenStakeSession) LpPoolInfo(arg0 *big.Int) (struct {
	LpToken           common.Address
	AllocPoint        *big.Int
	LastRewardBlock   *big.Int
	AccRewardPerShare *big.Int
}, error) {
	return _TokenStake.Contract.LpPoolInfo(&_TokenStake.CallOpts, arg0)
}

// LpPoolInfo is a free data retrieval call binding the contract method 0x2ba6b6f2.
//
// Solidity: function lpPoolInfo(uint256 ) view returns(address lpToken, uint256 allocPoint, uint256 lastRewardBlock, uint256 accRewardPerShare)
func (_TokenStake *TokenStakeCallerSession) LpPoolInfo(arg0 *big.Int) (struct {
	LpToken           common.Address
	AllocPoint        *big.Int
	LastRewardBlock   *big.Int
	AccRewardPerShare *big.Int
}, error) {
	return _TokenStake.Contract.LpPoolInfo(&_TokenStake.CallOpts, arg0)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_TokenStake *TokenStakeCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _TokenStake.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_TokenStake *TokenStakeSession) Owner() (common.Address, error) {
	return _TokenStake.Contract.Owner(&_TokenStake.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_TokenStake *TokenStakeCallerSession) Owner() (common.Address, error) {
	return _TokenStake.Contract.Owner(&_TokenStake.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_TokenStake *TokenStakeCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _TokenStake.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_TokenStake *TokenStakeSession) Paused() (bool, error) {
	return _TokenStake.Contract.Paused(&_TokenStake.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_TokenStake *TokenStakeCallerSession) Paused() (bool, error) {
	return _TokenStake.Contract.Paused(&_TokenStake.CallOpts)
}

// PendingLPReward is a free data retrieval call binding the contract method 0xdf7a899b.
//
// Solidity: function pendingLPReward(uint256 _pid, address _user) view returns(uint256[] pending)
func (_TokenStake *TokenStakeCaller) PendingLPReward(opts *bind.CallOpts, _pid *big.Int, _user common.Address) ([]*big.Int, error) {
	var out []interface{}
	err := _TokenStake.contract.Call(opts, &out, "pendingLPReward", _pid, _user)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// PendingLPReward is a free data retrieval call binding the contract method 0xdf7a899b.
//
// Solidity: function pendingLPReward(uint256 _pid, address _user) view returns(uint256[] pending)
func (_TokenStake *TokenStakeSession) PendingLPReward(_pid *big.Int, _user common.Address) ([]*big.Int, error) {
	return _TokenStake.Contract.PendingLPReward(&_TokenStake.CallOpts, _pid, _user)
}

// PendingLPReward is a free data retrieval call binding the contract method 0xdf7a899b.
//
// Solidity: function pendingLPReward(uint256 _pid, address _user) view returns(uint256[] pending)
func (_TokenStake *TokenStakeCallerSession) PendingLPReward(_pid *big.Int, _user common.Address) ([]*big.Int, error) {
	return _TokenStake.Contract.PendingLPReward(&_TokenStake.CallOpts, _pid, _user)
}

// RewardPerBlock is a free data retrieval call binding the contract method 0x8ae39cac.
//
// Solidity: function rewardPerBlock() view returns(uint256)
func (_TokenStake *TokenStakeCaller) RewardPerBlock(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _TokenStake.contract.Call(opts, &out, "rewardPerBlock")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RewardPerBlock is a free data retrieval call binding the contract method 0x8ae39cac.
//
// Solidity: function rewardPerBlock() view returns(uint256)
func (_TokenStake *TokenStakeSession) RewardPerBlock() (*big.Int, error) {
	return _TokenStake.Contract.RewardPerBlock(&_TokenStake.CallOpts)
}

// RewardPerBlock is a free data retrieval call binding the contract method 0x8ae39cac.
//
// Solidity: function rewardPerBlock() view returns(uint256)
func (_TokenStake *TokenStakeCallerSession) RewardPerBlock() (*big.Int, error) {
	return _TokenStake.Contract.RewardPerBlock(&_TokenStake.CallOpts)
}

// StartBlock is a free data retrieval call binding the contract method 0x48cd4cb1.
//
// Solidity: function startBlock() view returns(uint256)
func (_TokenStake *TokenStakeCaller) StartBlock(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _TokenStake.contract.Call(opts, &out, "startBlock")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// StartBlock is a free data retrieval call binding the contract method 0x48cd4cb1.
//
// Solidity: function startBlock() view returns(uint256)
func (_TokenStake *TokenStakeSession) StartBlock() (*big.Int, error) {
	return _TokenStake.Contract.StartBlock(&_TokenStake.CallOpts)
}

// StartBlock is a free data retrieval call binding the contract method 0x48cd4cb1.
//
// Solidity: function startBlock() view returns(uint256)
func (_TokenStake *TokenStakeCallerSession) StartBlock() (*big.Int, error) {
	return _TokenStake.Contract.StartBlock(&_TokenStake.CallOpts)
}

// TotalLPAllocPoint is a free data retrieval call binding the contract method 0x5a6e79bb.
//
// Solidity: function totalLPAllocPoint() view returns(uint256)
func (_TokenStake *TokenStakeCaller) TotalLPAllocPoint(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _TokenStake.contract.Call(opts, &out, "totalLPAllocPoint")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalLPAllocPoint is a free data retrieval call binding the contract method 0x5a6e79bb.
//
// Solidity: function totalLPAllocPoint() view returns(uint256)
func (_TokenStake *TokenStakeSession) TotalLPAllocPoint() (*big.Int, error) {
	return _TokenStake.Contract.TotalLPAllocPoint(&_TokenStake.CallOpts)
}

// TotalLPAllocPoint is a free data retrieval call binding the contract method 0x5a6e79bb.
//
// Solidity: function totalLPAllocPoint() view returns(uint256)
func (_TokenStake *TokenStakeCallerSession) TotalLPAllocPoint() (*big.Int, error) {
	return _TokenStake.Contract.TotalLPAllocPoint(&_TokenStake.CallOpts)
}

// UserLPInfo is a free data retrieval call binding the contract method 0xfaccc0a6.
//
// Solidity: function userLPInfo(uint256 , address ) view returns(uint256 amount, uint256 rewardDebt, uint256 firstDepositedTime, uint256 averageDepositedTime)
func (_TokenStake *TokenStakeCaller) UserLPInfo(opts *bind.CallOpts, arg0 *big.Int, arg1 common.Address) (struct {
	Amount               *big.Int
	RewardDebt           *big.Int
	FirstDepositedTime   *big.Int
	AverageDepositedTime *big.Int
}, error) {
	var out []interface{}
	err := _TokenStake.contract.Call(opts, &out, "userLPInfo", arg0, arg1)

	outstruct := new(struct {
		Amount               *big.Int
		RewardDebt           *big.Int
		FirstDepositedTime   *big.Int
		AverageDepositedTime *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Amount = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.RewardDebt = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.FirstDepositedTime = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.AverageDepositedTime = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// UserLPInfo is a free data retrieval call binding the contract method 0xfaccc0a6.
//
// Solidity: function userLPInfo(uint256 , address ) view returns(uint256 amount, uint256 rewardDebt, uint256 firstDepositedTime, uint256 averageDepositedTime)
func (_TokenStake *TokenStakeSession) UserLPInfo(arg0 *big.Int, arg1 common.Address) (struct {
	Amount               *big.Int
	RewardDebt           *big.Int
	FirstDepositedTime   *big.Int
	AverageDepositedTime *big.Int
}, error) {
	return _TokenStake.Contract.UserLPInfo(&_TokenStake.CallOpts, arg0, arg1)
}

// UserLPInfo is a free data retrieval call binding the contract method 0xfaccc0a6.
//
// Solidity: function userLPInfo(uint256 , address ) view returns(uint256 amount, uint256 rewardDebt, uint256 firstDepositedTime, uint256 averageDepositedTime)
func (_TokenStake *TokenStakeCallerSession) UserLPInfo(arg0 *big.Int, arg1 common.Address) (struct {
	Amount               *big.Int
	RewardDebt           *big.Int
	FirstDepositedTime   *big.Int
	AverageDepositedTime *big.Int
}, error) {
	return _TokenStake.Contract.UserLPInfo(&_TokenStake.CallOpts, arg0, arg1)
}

// AddLPPool is a paid mutator transaction binding the contract method 0x101cc6e5.
//
// Solidity: function addLPPool(address lpToken_, uint256 allocPoint_, uint256 lastRewardBlock_) returns()
func (_TokenStake *TokenStakeTransactor) AddLPPool(opts *bind.TransactOpts, lpToken_ common.Address, allocPoint_ *big.Int, lastRewardBlock_ *big.Int) (*types.Transaction, error) {
	return _TokenStake.contract.Transact(opts, "addLPPool", lpToken_, allocPoint_, lastRewardBlock_)
}

// AddLPPool is a paid mutator transaction binding the contract method 0x101cc6e5.
//
// Solidity: function addLPPool(address lpToken_, uint256 allocPoint_, uint256 lastRewardBlock_) returns()
func (_TokenStake *TokenStakeSession) AddLPPool(lpToken_ common.Address, allocPoint_ *big.Int, lastRewardBlock_ *big.Int) (*types.Transaction, error) {
	return _TokenStake.Contract.AddLPPool(&_TokenStake.TransactOpts, lpToken_, allocPoint_, lastRewardBlock_)
}

// AddLPPool is a paid mutator transaction binding the contract method 0x101cc6e5.
//
// Solidity: function addLPPool(address lpToken_, uint256 allocPoint_, uint256 lastRewardBlock_) returns()
func (_TokenStake *TokenStakeTransactorSession) AddLPPool(lpToken_ common.Address, allocPoint_ *big.Int, lastRewardBlock_ *big.Int) (*types.Transaction, error) {
	return _TokenStake.Contract.AddLPPool(&_TokenStake.TransactOpts, lpToken_, allocPoint_, lastRewardBlock_)
}

// DepositLP is a paid mutator transaction binding the contract method 0xbfe50f72.
//
// Solidity: function depositLP(uint256 _pid, uint256 _amount) returns()
func (_TokenStake *TokenStakeTransactor) DepositLP(opts *bind.TransactOpts, _pid *big.Int, _amount *big.Int) (*types.Transaction, error) {
	return _TokenStake.contract.Transact(opts, "depositLP", _pid, _amount)
}

// DepositLP is a paid mutator transaction binding the contract method 0xbfe50f72.
//
// Solidity: function depositLP(uint256 _pid, uint256 _amount) returns()
func (_TokenStake *TokenStakeSession) DepositLP(_pid *big.Int, _amount *big.Int) (*types.Transaction, error) {
	return _TokenStake.Contract.DepositLP(&_TokenStake.TransactOpts, _pid, _amount)
}

// DepositLP is a paid mutator transaction binding the contract method 0xbfe50f72.
//
// Solidity: function depositLP(uint256 _pid, uint256 _amount) returns()
func (_TokenStake *TokenStakeTransactorSession) DepositLP(_pid *big.Int, _amount *big.Int) (*types.Transaction, error) {
	return _TokenStake.Contract.DepositLP(&_TokenStake.TransactOpts, _pid, _amount)
}

// Initialize is a paid mutator transaction binding the contract method 0x7a1ac61e.
//
// Solidity: function initialize(address arsw_, uint256 rewardPerBlock_, uint256 startBlock_) returns()
func (_TokenStake *TokenStakeTransactor) Initialize(opts *bind.TransactOpts, arsw_ common.Address, rewardPerBlock_ *big.Int, startBlock_ *big.Int) (*types.Transaction, error) {
	return _TokenStake.contract.Transact(opts, "initialize", arsw_, rewardPerBlock_, startBlock_)
}

// Initialize is a paid mutator transaction binding the contract method 0x7a1ac61e.
//
// Solidity: function initialize(address arsw_, uint256 rewardPerBlock_, uint256 startBlock_) returns()
func (_TokenStake *TokenStakeSession) Initialize(arsw_ common.Address, rewardPerBlock_ *big.Int, startBlock_ *big.Int) (*types.Transaction, error) {
	return _TokenStake.Contract.Initialize(&_TokenStake.TransactOpts, arsw_, rewardPerBlock_, startBlock_)
}

// Initialize is a paid mutator transaction binding the contract method 0x7a1ac61e.
//
// Solidity: function initialize(address arsw_, uint256 rewardPerBlock_, uint256 startBlock_) returns()
func (_TokenStake *TokenStakeTransactorSession) Initialize(arsw_ common.Address, rewardPerBlock_ *big.Int, startBlock_ *big.Int) (*types.Transaction, error) {
	return _TokenStake.Contract.Initialize(&_TokenStake.TransactOpts, arsw_, rewardPerBlock_, startBlock_)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_TokenStake *TokenStakeTransactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TokenStake.contract.Transact(opts, "pause")
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_TokenStake *TokenStakeSession) Pause() (*types.Transaction, error) {
	return _TokenStake.Contract.Pause(&_TokenStake.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_TokenStake *TokenStakeTransactorSession) Pause() (*types.Transaction, error) {
	return _TokenStake.Contract.Pause(&_TokenStake.TransactOpts)
}

// PullFunds is a paid mutator transaction binding the contract method 0x5d5a92fb.
//
// Solidity: function pullFunds(address tokenAddress_) returns()
func (_TokenStake *TokenStakeTransactor) PullFunds(opts *bind.TransactOpts, tokenAddress_ common.Address) (*types.Transaction, error) {
	return _TokenStake.contract.Transact(opts, "pullFunds", tokenAddress_)
}

// PullFunds is a paid mutator transaction binding the contract method 0x5d5a92fb.
//
// Solidity: function pullFunds(address tokenAddress_) returns()
func (_TokenStake *TokenStakeSession) PullFunds(tokenAddress_ common.Address) (*types.Transaction, error) {
	return _TokenStake.Contract.PullFunds(&_TokenStake.TransactOpts, tokenAddress_)
}

// PullFunds is a paid mutator transaction binding the contract method 0x5d5a92fb.
//
// Solidity: function pullFunds(address tokenAddress_) returns()
func (_TokenStake *TokenStakeTransactorSession) PullFunds(tokenAddress_ common.Address) (*types.Transaction, error) {
	return _TokenStake.Contract.PullFunds(&_TokenStake.TransactOpts, tokenAddress_)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_TokenStake *TokenStakeTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TokenStake.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_TokenStake *TokenStakeSession) RenounceOwnership() (*types.Transaction, error) {
	return _TokenStake.Contract.RenounceOwnership(&_TokenStake.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_TokenStake *TokenStakeTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _TokenStake.Contract.RenounceOwnership(&_TokenStake.TransactOpts)
}

// ResetLPPool is a paid mutator transaction binding the contract method 0xed324ab0.
//
// Solidity: function resetLPPool(uint256 pid_, address lpToken_, uint256 allocPoint_, uint256 lastRewardBlock_, uint256 accRewardPerShare_) returns()
func (_TokenStake *TokenStakeTransactor) ResetLPPool(opts *bind.TransactOpts, pid_ *big.Int, lpToken_ common.Address, allocPoint_ *big.Int, lastRewardBlock_ *big.Int, accRewardPerShare_ *big.Int) (*types.Transaction, error) {
	return _TokenStake.contract.Transact(opts, "resetLPPool", pid_, lpToken_, allocPoint_, lastRewardBlock_, accRewardPerShare_)
}

// ResetLPPool is a paid mutator transaction binding the contract method 0xed324ab0.
//
// Solidity: function resetLPPool(uint256 pid_, address lpToken_, uint256 allocPoint_, uint256 lastRewardBlock_, uint256 accRewardPerShare_) returns()
func (_TokenStake *TokenStakeSession) ResetLPPool(pid_ *big.Int, lpToken_ common.Address, allocPoint_ *big.Int, lastRewardBlock_ *big.Int, accRewardPerShare_ *big.Int) (*types.Transaction, error) {
	return _TokenStake.Contract.ResetLPPool(&_TokenStake.TransactOpts, pid_, lpToken_, allocPoint_, lastRewardBlock_, accRewardPerShare_)
}

// ResetLPPool is a paid mutator transaction binding the contract method 0xed324ab0.
//
// Solidity: function resetLPPool(uint256 pid_, address lpToken_, uint256 allocPoint_, uint256 lastRewardBlock_, uint256 accRewardPerShare_) returns()
func (_TokenStake *TokenStakeTransactorSession) ResetLPPool(pid_ *big.Int, lpToken_ common.Address, allocPoint_ *big.Int, lastRewardBlock_ *big.Int, accRewardPerShare_ *big.Int) (*types.Transaction, error) {
	return _TokenStake.Contract.ResetLPPool(&_TokenStake.TransactOpts, pid_, lpToken_, allocPoint_, lastRewardBlock_, accRewardPerShare_)
}

// SetArswAddress is a paid mutator transaction binding the contract method 0x216c78ea.
//
// Solidity: function setArswAddress(address token_) returns()
func (_TokenStake *TokenStakeTransactor) SetArswAddress(opts *bind.TransactOpts, token_ common.Address) (*types.Transaction, error) {
	return _TokenStake.contract.Transact(opts, "setArswAddress", token_)
}

// SetArswAddress is a paid mutator transaction binding the contract method 0x216c78ea.
//
// Solidity: function setArswAddress(address token_) returns()
func (_TokenStake *TokenStakeSession) SetArswAddress(token_ common.Address) (*types.Transaction, error) {
	return _TokenStake.Contract.SetArswAddress(&_TokenStake.TransactOpts, token_)
}

// SetArswAddress is a paid mutator transaction binding the contract method 0x216c78ea.
//
// Solidity: function setArswAddress(address token_) returns()
func (_TokenStake *TokenStakeTransactorSession) SetArswAddress(token_ common.Address) (*types.Transaction, error) {
	return _TokenStake.Contract.SetArswAddress(&_TokenStake.TransactOpts, token_)
}

// SetRewardPerBlock is a paid mutator transaction binding the contract method 0xbb872b4a.
//
// Solidity: function setRewardPerBlock(uint256 _rewardPerBlock) returns()
func (_TokenStake *TokenStakeTransactor) SetRewardPerBlock(opts *bind.TransactOpts, _rewardPerBlock *big.Int) (*types.Transaction, error) {
	return _TokenStake.contract.Transact(opts, "setRewardPerBlock", _rewardPerBlock)
}

// SetRewardPerBlock is a paid mutator transaction binding the contract method 0xbb872b4a.
//
// Solidity: function setRewardPerBlock(uint256 _rewardPerBlock) returns()
func (_TokenStake *TokenStakeSession) SetRewardPerBlock(_rewardPerBlock *big.Int) (*types.Transaction, error) {
	return _TokenStake.Contract.SetRewardPerBlock(&_TokenStake.TransactOpts, _rewardPerBlock)
}

// SetRewardPerBlock is a paid mutator transaction binding the contract method 0xbb872b4a.
//
// Solidity: function setRewardPerBlock(uint256 _rewardPerBlock) returns()
func (_TokenStake *TokenStakeTransactorSession) SetRewardPerBlock(_rewardPerBlock *big.Int) (*types.Transaction, error) {
	return _TokenStake.Contract.SetRewardPerBlock(&_TokenStake.TransactOpts, _rewardPerBlock)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_TokenStake *TokenStakeTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _TokenStake.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_TokenStake *TokenStakeSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _TokenStake.Contract.TransferOwnership(&_TokenStake.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_TokenStake *TokenStakeTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _TokenStake.Contract.TransferOwnership(&_TokenStake.TransactOpts, newOwner)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_TokenStake *TokenStakeTransactor) Unpause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TokenStake.contract.Transact(opts, "unpause")
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_TokenStake *TokenStakeSession) Unpause() (*types.Transaction, error) {
	return _TokenStake.Contract.Unpause(&_TokenStake.TransactOpts)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_TokenStake *TokenStakeTransactorSession) Unpause() (*types.Transaction, error) {
	return _TokenStake.Contract.Unpause(&_TokenStake.TransactOpts)
}

// UpdateLPPool is a paid mutator transaction binding the contract method 0x2bddd6b1.
//
// Solidity: function updateLPPool(uint256 _pid) returns()
func (_TokenStake *TokenStakeTransactor) UpdateLPPool(opts *bind.TransactOpts, _pid *big.Int) (*types.Transaction, error) {
	return _TokenStake.contract.Transact(opts, "updateLPPool", _pid)
}

// UpdateLPPool is a paid mutator transaction binding the contract method 0x2bddd6b1.
//
// Solidity: function updateLPPool(uint256 _pid) returns()
func (_TokenStake *TokenStakeSession) UpdateLPPool(_pid *big.Int) (*types.Transaction, error) {
	return _TokenStake.Contract.UpdateLPPool(&_TokenStake.TransactOpts, _pid)
}

// UpdateLPPool is a paid mutator transaction binding the contract method 0x2bddd6b1.
//
// Solidity: function updateLPPool(uint256 _pid) returns()
func (_TokenStake *TokenStakeTransactorSession) UpdateLPPool(_pid *big.Int) (*types.Transaction, error) {
	return _TokenStake.Contract.UpdateLPPool(&_TokenStake.TransactOpts, _pid)
}

// WithdrawLP is a paid mutator transaction binding the contract method 0x91918d64.
//
// Solidity: function withdrawLP(uint256 _pid, uint256 _amount) returns()
func (_TokenStake *TokenStakeTransactor) WithdrawLP(opts *bind.TransactOpts, _pid *big.Int, _amount *big.Int) (*types.Transaction, error) {
	return _TokenStake.contract.Transact(opts, "withdrawLP", _pid, _amount)
}

// WithdrawLP is a paid mutator transaction binding the contract method 0x91918d64.
//
// Solidity: function withdrawLP(uint256 _pid, uint256 _amount) returns()
func (_TokenStake *TokenStakeSession) WithdrawLP(_pid *big.Int, _amount *big.Int) (*types.Transaction, error) {
	return _TokenStake.Contract.WithdrawLP(&_TokenStake.TransactOpts, _pid, _amount)
}

// WithdrawLP is a paid mutator transaction binding the contract method 0x91918d64.
//
// Solidity: function withdrawLP(uint256 _pid, uint256 _amount) returns()
func (_TokenStake *TokenStakeTransactorSession) WithdrawLP(_pid *big.Int, _amount *big.Int) (*types.Transaction, error) {
	return _TokenStake.Contract.WithdrawLP(&_TokenStake.TransactOpts, _pid, _amount)
}

// TokenStakeDepositIterator is returned from FilterDeposit and is used to iterate over the raw logs and unpacked data for Deposit events raised by the TokenStake contract.
type TokenStakeDepositIterator struct {
	Event *TokenStakeDeposit // Event containing the contract specifics and raw log

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
func (it *TokenStakeDepositIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenStakeDeposit)
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
		it.Event = new(TokenStakeDeposit)
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
func (it *TokenStakeDepositIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokenStakeDepositIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokenStakeDeposit represents a Deposit event raised by the TokenStake contract.
type TokenStakeDeposit struct {
	Who    common.Address
	Pid    *big.Int
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterDeposit is a free log retrieval operation binding the contract event 0x90890809c654f11d6e72a28fa60149770a0d11ec6c92319d6ceb2bb0a4ea1a15.
//
// Solidity: event Deposit(address who, uint256 pid, uint256 amount)
func (_TokenStake *TokenStakeFilterer) FilterDeposit(opts *bind.FilterOpts) (*TokenStakeDepositIterator, error) {

	logs, sub, err := _TokenStake.contract.FilterLogs(opts, "Deposit")
	if err != nil {
		return nil, err
	}
	return &TokenStakeDepositIterator{contract: _TokenStake.contract, event: "Deposit", logs: logs, sub: sub}, nil
}

// WatchDeposit is a free log subscription operation binding the contract event 0x90890809c654f11d6e72a28fa60149770a0d11ec6c92319d6ceb2bb0a4ea1a15.
//
// Solidity: event Deposit(address who, uint256 pid, uint256 amount)
func (_TokenStake *TokenStakeFilterer) WatchDeposit(opts *bind.WatchOpts, sink chan<- *TokenStakeDeposit) (event.Subscription, error) {

	logs, sub, err := _TokenStake.contract.WatchLogs(opts, "Deposit")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokenStakeDeposit)
				if err := _TokenStake.contract.UnpackLog(event, "Deposit", log); err != nil {
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

// ParseDeposit is a log parse operation binding the contract event 0x90890809c654f11d6e72a28fa60149770a0d11ec6c92319d6ceb2bb0a4ea1a15.
//
// Solidity: event Deposit(address who, uint256 pid, uint256 amount)
func (_TokenStake *TokenStakeFilterer) ParseDeposit(log types.Log) (*TokenStakeDeposit, error) {
	event := new(TokenStakeDeposit)
	if err := _TokenStake.contract.UnpackLog(event, "Deposit", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TokenStakeInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the TokenStake contract.
type TokenStakeInitializedIterator struct {
	Event *TokenStakeInitialized // Event containing the contract specifics and raw log

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
func (it *TokenStakeInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenStakeInitialized)
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
		it.Event = new(TokenStakeInitialized)
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
func (it *TokenStakeInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokenStakeInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokenStakeInitialized represents a Initialized event raised by the TokenStake contract.
type TokenStakeInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_TokenStake *TokenStakeFilterer) FilterInitialized(opts *bind.FilterOpts) (*TokenStakeInitializedIterator, error) {

	logs, sub, err := _TokenStake.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &TokenStakeInitializedIterator{contract: _TokenStake.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_TokenStake *TokenStakeFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *TokenStakeInitialized) (event.Subscription, error) {

	logs, sub, err := _TokenStake.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokenStakeInitialized)
				if err := _TokenStake.contract.UnpackLog(event, "Initialized", log); err != nil {
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

// ParseInitialized is a log parse operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_TokenStake *TokenStakeFilterer) ParseInitialized(log types.Log) (*TokenStakeInitialized, error) {
	event := new(TokenStakeInitialized)
	if err := _TokenStake.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TokenStakeOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the TokenStake contract.
type TokenStakeOwnershipTransferredIterator struct {
	Event *TokenStakeOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *TokenStakeOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenStakeOwnershipTransferred)
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
		it.Event = new(TokenStakeOwnershipTransferred)
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
func (it *TokenStakeOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokenStakeOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokenStakeOwnershipTransferred represents a OwnershipTransferred event raised by the TokenStake contract.
type TokenStakeOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_TokenStake *TokenStakeFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*TokenStakeOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _TokenStake.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &TokenStakeOwnershipTransferredIterator{contract: _TokenStake.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_TokenStake *TokenStakeFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *TokenStakeOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _TokenStake.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokenStakeOwnershipTransferred)
				if err := _TokenStake.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_TokenStake *TokenStakeFilterer) ParseOwnershipTransferred(log types.Log) (*TokenStakeOwnershipTransferred, error) {
	event := new(TokenStakeOwnershipTransferred)
	if err := _TokenStake.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TokenStakePausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the TokenStake contract.
type TokenStakePausedIterator struct {
	Event *TokenStakePaused // Event containing the contract specifics and raw log

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
func (it *TokenStakePausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenStakePaused)
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
		it.Event = new(TokenStakePaused)
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
func (it *TokenStakePausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokenStakePausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokenStakePaused represents a Paused event raised by the TokenStake contract.
type TokenStakePaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_TokenStake *TokenStakeFilterer) FilterPaused(opts *bind.FilterOpts) (*TokenStakePausedIterator, error) {

	logs, sub, err := _TokenStake.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &TokenStakePausedIterator{contract: _TokenStake.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_TokenStake *TokenStakeFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *TokenStakePaused) (event.Subscription, error) {

	logs, sub, err := _TokenStake.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokenStakePaused)
				if err := _TokenStake.contract.UnpackLog(event, "Paused", log); err != nil {
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
func (_TokenStake *TokenStakeFilterer) ParsePaused(log types.Log) (*TokenStakePaused, error) {
	event := new(TokenStakePaused)
	if err := _TokenStake.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TokenStakeUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the TokenStake contract.
type TokenStakeUnpausedIterator struct {
	Event *TokenStakeUnpaused // Event containing the contract specifics and raw log

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
func (it *TokenStakeUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenStakeUnpaused)
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
		it.Event = new(TokenStakeUnpaused)
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
func (it *TokenStakeUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokenStakeUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokenStakeUnpaused represents a Unpaused event raised by the TokenStake contract.
type TokenStakeUnpaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_TokenStake *TokenStakeFilterer) FilterUnpaused(opts *bind.FilterOpts) (*TokenStakeUnpausedIterator, error) {

	logs, sub, err := _TokenStake.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &TokenStakeUnpausedIterator{contract: _TokenStake.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_TokenStake *TokenStakeFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *TokenStakeUnpaused) (event.Subscription, error) {

	logs, sub, err := _TokenStake.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokenStakeUnpaused)
				if err := _TokenStake.contract.UnpackLog(event, "Unpaused", log); err != nil {
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
func (_TokenStake *TokenStakeFilterer) ParseUnpaused(log types.Log) (*TokenStakeUnpaused, error) {
	event := new(TokenStakeUnpaused)
	if err := _TokenStake.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TokenStakeWithdrawIterator is returned from FilterWithdraw and is used to iterate over the raw logs and unpacked data for Withdraw events raised by the TokenStake contract.
type TokenStakeWithdrawIterator struct {
	Event *TokenStakeWithdraw // Event containing the contract specifics and raw log

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
func (it *TokenStakeWithdrawIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenStakeWithdraw)
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
		it.Event = new(TokenStakeWithdraw)
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
func (it *TokenStakeWithdrawIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokenStakeWithdrawIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokenStakeWithdraw represents a Withdraw event raised by the TokenStake contract.
type TokenStakeWithdraw struct {
	Who    common.Address
	Pid    *big.Int
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterWithdraw is a free log retrieval operation binding the contract event 0xf279e6a1f5e320cca91135676d9cb6e44ca8a08c0b88342bcdb1144f6511b568.
//
// Solidity: event Withdraw(address who, uint256 pid, uint256 amount)
func (_TokenStake *TokenStakeFilterer) FilterWithdraw(opts *bind.FilterOpts) (*TokenStakeWithdrawIterator, error) {

	logs, sub, err := _TokenStake.contract.FilterLogs(opts, "Withdraw")
	if err != nil {
		return nil, err
	}
	return &TokenStakeWithdrawIterator{contract: _TokenStake.contract, event: "Withdraw", logs: logs, sub: sub}, nil
}

// WatchWithdraw is a free log subscription operation binding the contract event 0xf279e6a1f5e320cca91135676d9cb6e44ca8a08c0b88342bcdb1144f6511b568.
//
// Solidity: event Withdraw(address who, uint256 pid, uint256 amount)
func (_TokenStake *TokenStakeFilterer) WatchWithdraw(opts *bind.WatchOpts, sink chan<- *TokenStakeWithdraw) (event.Subscription, error) {

	logs, sub, err := _TokenStake.contract.WatchLogs(opts, "Withdraw")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokenStakeWithdraw)
				if err := _TokenStake.contract.UnpackLog(event, "Withdraw", log); err != nil {
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

// ParseWithdraw is a log parse operation binding the contract event 0xf279e6a1f5e320cca91135676d9cb6e44ca8a08c0b88342bcdb1144f6511b568.
//
// Solidity: event Withdraw(address who, uint256 pid, uint256 amount)
func (_TokenStake *TokenStakeFilterer) ParseWithdraw(log types.Log) (*TokenStakeWithdraw, error) {
	event := new(TokenStakeWithdraw)
	if err := _TokenStake.contract.UnpackLog(event, "Withdraw", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
