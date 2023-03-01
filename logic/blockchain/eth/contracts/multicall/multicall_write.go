// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package multicall

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

// Multicall2Call is an auto generated low-level Go binding around an user-defined struct.
type Multicall2Call struct {
	Target   common.Address
	CallData []byte
}

// Multicall2Result is an auto generated low-level Go binding around an user-defined struct.
type Multicall2Result struct {
	Success    bool
	ReturnData []byte
}

// MulticallWriterMetaData contains all meta data concerning the MulticallWriter contract.
var MulticallWriterMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"callData\",\"type\":\"bytes\"}],\"internalType\":\"structMulticall2.Call[]\",\"name\":\"calls\",\"type\":\"tuple[]\"}],\"name\":\"aggregate\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"blockNumber\",\"type\":\"uint256\"},{\"internalType\":\"bytes[]\",\"name\":\"returnData\",\"type\":\"bytes[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"callData\",\"type\":\"bytes\"}],\"internalType\":\"structMulticall2.Call[]\",\"name\":\"calls\",\"type\":\"tuple[]\"}],\"name\":\"blockAndAggregate\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"blockNumber\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"blockHash\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"},{\"internalType\":\"bytes\",\"name\":\"returnData\",\"type\":\"bytes\"}],\"internalType\":\"structMulticall2.Result[]\",\"name\":\"returnData\",\"type\":\"tuple[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"blockNumber\",\"type\":\"uint256\"}],\"name\":\"getBlockHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"blockHash\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getBlockNumber\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"blockNumber\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getCurrentBlockCoinbase\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"coinbase\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getCurrentBlockDifficulty\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"difficulty\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getCurrentBlockGasLimit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"gaslimit\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getCurrentBlockTimestamp\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"getEthBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getLastBlockHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"blockHash\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"requireSuccess\",\"type\":\"bool\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"callData\",\"type\":\"bytes\"}],\"internalType\":\"structMulticall2.Call[]\",\"name\":\"calls\",\"type\":\"tuple[]\"}],\"name\":\"tryAggregate\",\"outputs\":[{\"components\":[{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"},{\"internalType\":\"bytes\",\"name\":\"returnData\",\"type\":\"bytes\"}],\"internalType\":\"structMulticall2.Result[]\",\"name\":\"returnData\",\"type\":\"tuple[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"requireSuccess\",\"type\":\"bool\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"callData\",\"type\":\"bytes\"}],\"internalType\":\"structMulticall2.Call[]\",\"name\":\"calls\",\"type\":\"tuple[]\"}],\"name\":\"tryBlockAndAggregate\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"blockNumber\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"blockHash\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"},{\"internalType\":\"bytes\",\"name\":\"returnData\",\"type\":\"bytes\"}],\"internalType\":\"structMulticall2.Result[]\",\"name\":\"returnData\",\"type\":\"tuple[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// MulticallWriterABI is the input ABI used to generate the binding from.
// Deprecated: Use MulticallWriterMetaData.ABI instead.
var MulticallWriterABI = MulticallWriterMetaData.ABI

// MulticallWriter is an auto generated Go binding around an Ethereum contract.
type MulticallWriter struct {
	MulticallWriterCaller     // Read-only binding to the contract
	MulticallWriterTransactor // Write-only binding to the contract
	MulticallWriterFilterer   // Log filterer for contract events
}

// MulticallWriterCaller is an auto generated read-only Go binding around an Ethereum contract.
type MulticallWriterCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MulticallWriterTransactor is an auto generated write-only Go binding around an Ethereum contract.
type MulticallWriterTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MulticallWriterFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MulticallWriterFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MulticallWriterSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MulticallWriterSession struct {
	Contract     *MulticallWriter  // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MulticallWriterCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MulticallWriterCallerSession struct {
	Contract *MulticallWriterCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts          // Call options to use throughout this session
}

// MulticallWriterTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MulticallWriterTransactorSession struct {
	Contract     *MulticallWriterTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// MulticallWriterRaw is an auto generated low-level Go binding around an Ethereum contract.
type MulticallWriterRaw struct {
	Contract *MulticallWriter // Generic contract binding to access the raw methods on
}

// MulticallWriterCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MulticallWriterCallerRaw struct {
	Contract *MulticallWriterCaller // Generic read-only contract binding to access the raw methods on
}

// MulticallWriterTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MulticallWriterTransactorRaw struct {
	Contract *MulticallWriterTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMulticallWriter creates a new instance of MulticallWriter, bound to a specific deployed contract.
func NewMulticallWriter(address common.Address, backend bind.ContractBackend) (*MulticallWriter, error) {
	contract, err := bindMulticallWriter(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &MulticallWriter{MulticallWriterCaller: MulticallWriterCaller{contract: contract}, MulticallWriterTransactor: MulticallWriterTransactor{contract: contract}, MulticallWriterFilterer: MulticallWriterFilterer{contract: contract}}, nil
}

// NewMulticallWriterCaller creates a new read-only instance of MulticallWriter, bound to a specific deployed contract.
func NewMulticallWriterCaller(address common.Address, caller bind.ContractCaller) (*MulticallWriterCaller, error) {
	contract, err := bindMulticallWriter(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MulticallWriterCaller{contract: contract}, nil
}

// NewMulticallWriterTransactor creates a new write-only instance of MulticallWriter, bound to a specific deployed contract.
func NewMulticallWriterTransactor(address common.Address, transactor bind.ContractTransactor) (*MulticallWriterTransactor, error) {
	contract, err := bindMulticallWriter(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MulticallWriterTransactor{contract: contract}, nil
}

// NewMulticallWriterFilterer creates a new log filterer instance of MulticallWriter, bound to a specific deployed contract.
func NewMulticallWriterFilterer(address common.Address, filterer bind.ContractFilterer) (*MulticallWriterFilterer, error) {
	contract, err := bindMulticallWriter(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MulticallWriterFilterer{contract: contract}, nil
}

// bindMulticallWriter binds a generic wrapper to an already deployed contract.
func bindMulticallWriter(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := MulticallWriterMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MulticallWriter *MulticallWriterRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MulticallWriter.Contract.MulticallWriterCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MulticallWriter *MulticallWriterRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MulticallWriter.Contract.MulticallWriterTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MulticallWriter *MulticallWriterRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MulticallWriter.Contract.MulticallWriterTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MulticallWriter *MulticallWriterCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MulticallWriter.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MulticallWriter *MulticallWriterTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MulticallWriter.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MulticallWriter *MulticallWriterTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MulticallWriter.Contract.contract.Transact(opts, method, params...)
}

// GetBlockHash is a free data retrieval call binding the contract method 0xee82ac5e.
//
// Solidity: function getBlockHash(uint256 blockNumber) view returns(bytes32 blockHash)
func (_MulticallWriter *MulticallWriterCaller) GetBlockHash(opts *bind.CallOpts, blockNumber *big.Int) ([32]byte, error) {
	var out []interface{}
	err := _MulticallWriter.contract.Call(opts, &out, "getBlockHash", blockNumber)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetBlockHash is a free data retrieval call binding the contract method 0xee82ac5e.
//
// Solidity: function getBlockHash(uint256 blockNumber) view returns(bytes32 blockHash)
func (_MulticallWriter *MulticallWriterSession) GetBlockHash(blockNumber *big.Int) ([32]byte, error) {
	return _MulticallWriter.Contract.GetBlockHash(&_MulticallWriter.CallOpts, blockNumber)
}

// GetBlockHash is a free data retrieval call binding the contract method 0xee82ac5e.
//
// Solidity: function getBlockHash(uint256 blockNumber) view returns(bytes32 blockHash)
func (_MulticallWriter *MulticallWriterCallerSession) GetBlockHash(blockNumber *big.Int) ([32]byte, error) {
	return _MulticallWriter.Contract.GetBlockHash(&_MulticallWriter.CallOpts, blockNumber)
}

// GetBlockNumber is a free data retrieval call binding the contract method 0x42cbb15c.
//
// Solidity: function getBlockNumber() view returns(uint256 blockNumber)
func (_MulticallWriter *MulticallWriterCaller) GetBlockNumber(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _MulticallWriter.contract.Call(opts, &out, "getBlockNumber")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetBlockNumber is a free data retrieval call binding the contract method 0x42cbb15c.
//
// Solidity: function getBlockNumber() view returns(uint256 blockNumber)
func (_MulticallWriter *MulticallWriterSession) GetBlockNumber() (*big.Int, error) {
	return _MulticallWriter.Contract.GetBlockNumber(&_MulticallWriter.CallOpts)
}

// GetBlockNumber is a free data retrieval call binding the contract method 0x42cbb15c.
//
// Solidity: function getBlockNumber() view returns(uint256 blockNumber)
func (_MulticallWriter *MulticallWriterCallerSession) GetBlockNumber() (*big.Int, error) {
	return _MulticallWriter.Contract.GetBlockNumber(&_MulticallWriter.CallOpts)
}

// GetCurrentBlockCoinbase is a free data retrieval call binding the contract method 0xa8b0574e.
//
// Solidity: function getCurrentBlockCoinbase() view returns(address coinbase)
func (_MulticallWriter *MulticallWriterCaller) GetCurrentBlockCoinbase(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _MulticallWriter.contract.Call(opts, &out, "getCurrentBlockCoinbase")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetCurrentBlockCoinbase is a free data retrieval call binding the contract method 0xa8b0574e.
//
// Solidity: function getCurrentBlockCoinbase() view returns(address coinbase)
func (_MulticallWriter *MulticallWriterSession) GetCurrentBlockCoinbase() (common.Address, error) {
	return _MulticallWriter.Contract.GetCurrentBlockCoinbase(&_MulticallWriter.CallOpts)
}

// GetCurrentBlockCoinbase is a free data retrieval call binding the contract method 0xa8b0574e.
//
// Solidity: function getCurrentBlockCoinbase() view returns(address coinbase)
func (_MulticallWriter *MulticallWriterCallerSession) GetCurrentBlockCoinbase() (common.Address, error) {
	return _MulticallWriter.Contract.GetCurrentBlockCoinbase(&_MulticallWriter.CallOpts)
}

// GetCurrentBlockDifficulty is a free data retrieval call binding the contract method 0x72425d9d.
//
// Solidity: function getCurrentBlockDifficulty() view returns(uint256 difficulty)
func (_MulticallWriter *MulticallWriterCaller) GetCurrentBlockDifficulty(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _MulticallWriter.contract.Call(opts, &out, "getCurrentBlockDifficulty")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetCurrentBlockDifficulty is a free data retrieval call binding the contract method 0x72425d9d.
//
// Solidity: function getCurrentBlockDifficulty() view returns(uint256 difficulty)
func (_MulticallWriter *MulticallWriterSession) GetCurrentBlockDifficulty() (*big.Int, error) {
	return _MulticallWriter.Contract.GetCurrentBlockDifficulty(&_MulticallWriter.CallOpts)
}

// GetCurrentBlockDifficulty is a free data retrieval call binding the contract method 0x72425d9d.
//
// Solidity: function getCurrentBlockDifficulty() view returns(uint256 difficulty)
func (_MulticallWriter *MulticallWriterCallerSession) GetCurrentBlockDifficulty() (*big.Int, error) {
	return _MulticallWriter.Contract.GetCurrentBlockDifficulty(&_MulticallWriter.CallOpts)
}

// GetCurrentBlockGasLimit is a free data retrieval call binding the contract method 0x86d516e8.
//
// Solidity: function getCurrentBlockGasLimit() view returns(uint256 gaslimit)
func (_MulticallWriter *MulticallWriterCaller) GetCurrentBlockGasLimit(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _MulticallWriter.contract.Call(opts, &out, "getCurrentBlockGasLimit")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetCurrentBlockGasLimit is a free data retrieval call binding the contract method 0x86d516e8.
//
// Solidity: function getCurrentBlockGasLimit() view returns(uint256 gaslimit)
func (_MulticallWriter *MulticallWriterSession) GetCurrentBlockGasLimit() (*big.Int, error) {
	return _MulticallWriter.Contract.GetCurrentBlockGasLimit(&_MulticallWriter.CallOpts)
}

// GetCurrentBlockGasLimit is a free data retrieval call binding the contract method 0x86d516e8.
//
// Solidity: function getCurrentBlockGasLimit() view returns(uint256 gaslimit)
func (_MulticallWriter *MulticallWriterCallerSession) GetCurrentBlockGasLimit() (*big.Int, error) {
	return _MulticallWriter.Contract.GetCurrentBlockGasLimit(&_MulticallWriter.CallOpts)
}

// GetCurrentBlockTimestamp is a free data retrieval call binding the contract method 0x0f28c97d.
//
// Solidity: function getCurrentBlockTimestamp() view returns(uint256 timestamp)
func (_MulticallWriter *MulticallWriterCaller) GetCurrentBlockTimestamp(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _MulticallWriter.contract.Call(opts, &out, "getCurrentBlockTimestamp")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetCurrentBlockTimestamp is a free data retrieval call binding the contract method 0x0f28c97d.
//
// Solidity: function getCurrentBlockTimestamp() view returns(uint256 timestamp)
func (_MulticallWriter *MulticallWriterSession) GetCurrentBlockTimestamp() (*big.Int, error) {
	return _MulticallWriter.Contract.GetCurrentBlockTimestamp(&_MulticallWriter.CallOpts)
}

// GetCurrentBlockTimestamp is a free data retrieval call binding the contract method 0x0f28c97d.
//
// Solidity: function getCurrentBlockTimestamp() view returns(uint256 timestamp)
func (_MulticallWriter *MulticallWriterCallerSession) GetCurrentBlockTimestamp() (*big.Int, error) {
	return _MulticallWriter.Contract.GetCurrentBlockTimestamp(&_MulticallWriter.CallOpts)
}

// GetEthBalance is a free data retrieval call binding the contract method 0x4d2301cc.
//
// Solidity: function getEthBalance(address addr) view returns(uint256 balance)
func (_MulticallWriter *MulticallWriterCaller) GetEthBalance(opts *bind.CallOpts, addr common.Address) (*big.Int, error) {
	var out []interface{}
	err := _MulticallWriter.contract.Call(opts, &out, "getEthBalance", addr)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetEthBalance is a free data retrieval call binding the contract method 0x4d2301cc.
//
// Solidity: function getEthBalance(address addr) view returns(uint256 balance)
func (_MulticallWriter *MulticallWriterSession) GetEthBalance(addr common.Address) (*big.Int, error) {
	return _MulticallWriter.Contract.GetEthBalance(&_MulticallWriter.CallOpts, addr)
}

// GetEthBalance is a free data retrieval call binding the contract method 0x4d2301cc.
//
// Solidity: function getEthBalance(address addr) view returns(uint256 balance)
func (_MulticallWriter *MulticallWriterCallerSession) GetEthBalance(addr common.Address) (*big.Int, error) {
	return _MulticallWriter.Contract.GetEthBalance(&_MulticallWriter.CallOpts, addr)
}

// GetLastBlockHash is a free data retrieval call binding the contract method 0x27e86d6e.
//
// Solidity: function getLastBlockHash() view returns(bytes32 blockHash)
func (_MulticallWriter *MulticallWriterCaller) GetLastBlockHash(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _MulticallWriter.contract.Call(opts, &out, "getLastBlockHash")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetLastBlockHash is a free data retrieval call binding the contract method 0x27e86d6e.
//
// Solidity: function getLastBlockHash() view returns(bytes32 blockHash)
func (_MulticallWriter *MulticallWriterSession) GetLastBlockHash() ([32]byte, error) {
	return _MulticallWriter.Contract.GetLastBlockHash(&_MulticallWriter.CallOpts)
}

// GetLastBlockHash is a free data retrieval call binding the contract method 0x27e86d6e.
//
// Solidity: function getLastBlockHash() view returns(bytes32 blockHash)
func (_MulticallWriter *MulticallWriterCallerSession) GetLastBlockHash() ([32]byte, error) {
	return _MulticallWriter.Contract.GetLastBlockHash(&_MulticallWriter.CallOpts)
}

// Aggregate is a paid mutator transaction binding the contract method 0x252dba42.
//
// Solidity: function aggregate((address,bytes)[] calls) returns(uint256 blockNumber, bytes[] returnData)
func (_MulticallWriter *MulticallWriterTransactor) Aggregate(opts *bind.TransactOpts, calls []Multicall2Call) (*types.Transaction, error) {
	return _MulticallWriter.contract.Transact(opts, "aggregate", calls)
}

// Aggregate is a paid mutator transaction binding the contract method 0x252dba42.
//
// Solidity: function aggregate((address,bytes)[] calls) returns(uint256 blockNumber, bytes[] returnData)
func (_MulticallWriter *MulticallWriterSession) Aggregate(calls []Multicall2Call) (*types.Transaction, error) {
	return _MulticallWriter.Contract.Aggregate(&_MulticallWriter.TransactOpts, calls)
}

// Aggregate is a paid mutator transaction binding the contract method 0x252dba42.
//
// Solidity: function aggregate((address,bytes)[] calls) returns(uint256 blockNumber, bytes[] returnData)
func (_MulticallWriter *MulticallWriterTransactorSession) Aggregate(calls []Multicall2Call) (*types.Transaction, error) {
	return _MulticallWriter.Contract.Aggregate(&_MulticallWriter.TransactOpts, calls)
}

// BlockAndAggregate is a paid mutator transaction binding the contract method 0xc3077fa9.
//
// Solidity: function blockAndAggregate((address,bytes)[] calls) returns(uint256 blockNumber, bytes32 blockHash, (bool,bytes)[] returnData)
func (_MulticallWriter *MulticallWriterTransactor) BlockAndAggregate(opts *bind.TransactOpts, calls []Multicall2Call) (*types.Transaction, error) {
	return _MulticallWriter.contract.Transact(opts, "blockAndAggregate", calls)
}

// BlockAndAggregate is a paid mutator transaction binding the contract method 0xc3077fa9.
//
// Solidity: function blockAndAggregate((address,bytes)[] calls) returns(uint256 blockNumber, bytes32 blockHash, (bool,bytes)[] returnData)
func (_MulticallWriter *MulticallWriterSession) BlockAndAggregate(calls []Multicall2Call) (*types.Transaction, error) {
	return _MulticallWriter.Contract.BlockAndAggregate(&_MulticallWriter.TransactOpts, calls)
}

// BlockAndAggregate is a paid mutator transaction binding the contract method 0xc3077fa9.
//
// Solidity: function blockAndAggregate((address,bytes)[] calls) returns(uint256 blockNumber, bytes32 blockHash, (bool,bytes)[] returnData)
func (_MulticallWriter *MulticallWriterTransactorSession) BlockAndAggregate(calls []Multicall2Call) (*types.Transaction, error) {
	return _MulticallWriter.Contract.BlockAndAggregate(&_MulticallWriter.TransactOpts, calls)
}

// TryAggregate is a paid mutator transaction binding the contract method 0xbce38bd7.
//
// Solidity: function tryAggregate(bool requireSuccess, (address,bytes)[] calls) returns((bool,bytes)[] returnData)
func (_MulticallWriter *MulticallWriterTransactor) TryAggregate(opts *bind.TransactOpts, requireSuccess bool, calls []Multicall2Call) (*types.Transaction, error) {
	return _MulticallWriter.contract.Transact(opts, "tryAggregate", requireSuccess, calls)
}

// TryAggregate is a paid mutator transaction binding the contract method 0xbce38bd7.
//
// Solidity: function tryAggregate(bool requireSuccess, (address,bytes)[] calls) returns((bool,bytes)[] returnData)
func (_MulticallWriter *MulticallWriterSession) TryAggregate(requireSuccess bool, calls []Multicall2Call) (*types.Transaction, error) {
	return _MulticallWriter.Contract.TryAggregate(&_MulticallWriter.TransactOpts, requireSuccess, calls)
}

// TryAggregate is a paid mutator transaction binding the contract method 0xbce38bd7.
//
// Solidity: function tryAggregate(bool requireSuccess, (address,bytes)[] calls) returns((bool,bytes)[] returnData)
func (_MulticallWriter *MulticallWriterTransactorSession) TryAggregate(requireSuccess bool, calls []Multicall2Call) (*types.Transaction, error) {
	return _MulticallWriter.Contract.TryAggregate(&_MulticallWriter.TransactOpts, requireSuccess, calls)
}

// TryBlockAndAggregate is a paid mutator transaction binding the contract method 0x399542e9.
//
// Solidity: function tryBlockAndAggregate(bool requireSuccess, (address,bytes)[] calls) returns(uint256 blockNumber, bytes32 blockHash, (bool,bytes)[] returnData)
func (_MulticallWriter *MulticallWriterTransactor) TryBlockAndAggregate(opts *bind.TransactOpts, requireSuccess bool, calls []Multicall2Call) (*types.Transaction, error) {
	return _MulticallWriter.contract.Transact(opts, "tryBlockAndAggregate", requireSuccess, calls)
}

// TryBlockAndAggregate is a paid mutator transaction binding the contract method 0x399542e9.
//
// Solidity: function tryBlockAndAggregate(bool requireSuccess, (address,bytes)[] calls) returns(uint256 blockNumber, bytes32 blockHash, (bool,bytes)[] returnData)
func (_MulticallWriter *MulticallWriterSession) TryBlockAndAggregate(requireSuccess bool, calls []Multicall2Call) (*types.Transaction, error) {
	return _MulticallWriter.Contract.TryBlockAndAggregate(&_MulticallWriter.TransactOpts, requireSuccess, calls)
}

// TryBlockAndAggregate is a paid mutator transaction binding the contract method 0x399542e9.
//
// Solidity: function tryBlockAndAggregate(bool requireSuccess, (address,bytes)[] calls) returns(uint256 blockNumber, bytes32 blockHash, (bool,bytes)[] returnData)
func (_MulticallWriter *MulticallWriterTransactorSession) TryBlockAndAggregate(requireSuccess bool, calls []Multicall2Call) (*types.Transaction, error) {
	return _MulticallWriter.Contract.TryBlockAndAggregate(&_MulticallWriter.TransactOpts, requireSuccess, calls)
}
