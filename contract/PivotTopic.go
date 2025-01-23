// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contract

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

// PivotTopicMetaData contains all meta data concerning the PivotTopic contract.
var PivotTopicMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"a\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"promoter\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"topicId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"investment\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"position\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"}],\"name\":\"CreateTopic\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"investor\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"topicId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"position\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"}],\"name\":\"Invest\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"}],\"name\":\"Withdraw\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"}],\"name\":\"WithdrawCommission\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"topicId\",\"type\":\"uint256\"}],\"name\":\"_topicHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"topicId\",\"type\":\"uint256\"}],\"name\":\"_totalBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"topicId\",\"type\":\"uint256\"}],\"name\":\"_totalCommission\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"erc20Address\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"topicHash\",\"type\":\"bytes32\"}],\"name\":\"createTopic\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"topicId\",\"type\":\"uint256\"}],\"name\":\"getFixedInvestment\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"investor\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"topicId\",\"type\":\"uint256\"}],\"name\":\"getIncome\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"investor\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"topicId\",\"type\":\"uint256\"}],\"name\":\"getInvestment\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"topicId\",\"type\":\"uint256\"}],\"name\":\"getPromoter\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"topicId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"invest\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"sbtAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"sbtContract\",\"outputs\":[{\"internalType\":\"contractISBTContract\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"topicId\",\"type\":\"uint256\"}],\"name\":\"topicCoin\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"topicId\",\"type\":\"uint256\"}],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"topicId\",\"type\":\"uint256\"}],\"name\":\"withdrawCommission\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// PivotTopicABI is the input ABI used to generate the binding from.
// Deprecated: Use PivotTopicMetaData.ABI instead.
var PivotTopicABI = PivotTopicMetaData.ABI

// PivotTopic is an auto generated Go binding around an Ethereum contract.
type PivotTopic struct {
	PivotTopicCaller     // Read-only binding to the contract
	PivotTopicTransactor // Write-only binding to the contract
	PivotTopicFilterer   // Log filterer for contract events
}

// PivotTopicCaller is an auto generated read-only Go binding around an Ethereum contract.
type PivotTopicCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PivotTopicTransactor is an auto generated write-only Go binding around an Ethereum contract.
type PivotTopicTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PivotTopicFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type PivotTopicFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PivotTopicSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PivotTopicSession struct {
	Contract     *PivotTopic       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// PivotTopicCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PivotTopicCallerSession struct {
	Contract *PivotTopicCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// PivotTopicTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PivotTopicTransactorSession struct {
	Contract     *PivotTopicTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// PivotTopicRaw is an auto generated low-level Go binding around an Ethereum contract.
type PivotTopicRaw struct {
	Contract *PivotTopic // Generic contract binding to access the raw methods on
}

// PivotTopicCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PivotTopicCallerRaw struct {
	Contract *PivotTopicCaller // Generic read-only contract binding to access the raw methods on
}

// PivotTopicTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PivotTopicTransactorRaw struct {
	Contract *PivotTopicTransactor // Generic write-only contract binding to access the raw methods on
}

// NewPivotTopic creates a new instance of PivotTopic, bound to a specific deployed contract.
func NewPivotTopic(address common.Address, backend bind.ContractBackend) (*PivotTopic, error) {
	contract, err := bindPivotTopic(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &PivotTopic{PivotTopicCaller: PivotTopicCaller{contract: contract}, PivotTopicTransactor: PivotTopicTransactor{contract: contract}, PivotTopicFilterer: PivotTopicFilterer{contract: contract}}, nil
}

// NewPivotTopicCaller creates a new read-only instance of PivotTopic, bound to a specific deployed contract.
func NewPivotTopicCaller(address common.Address, caller bind.ContractCaller) (*PivotTopicCaller, error) {
	contract, err := bindPivotTopic(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &PivotTopicCaller{contract: contract}, nil
}

// NewPivotTopicTransactor creates a new write-only instance of PivotTopic, bound to a specific deployed contract.
func NewPivotTopicTransactor(address common.Address, transactor bind.ContractTransactor) (*PivotTopicTransactor, error) {
	contract, err := bindPivotTopic(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &PivotTopicTransactor{contract: contract}, nil
}

// NewPivotTopicFilterer creates a new log filterer instance of PivotTopic, bound to a specific deployed contract.
func NewPivotTopicFilterer(address common.Address, filterer bind.ContractFilterer) (*PivotTopicFilterer, error) {
	contract, err := bindPivotTopic(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &PivotTopicFilterer{contract: contract}, nil
}

// bindPivotTopic binds a generic wrapper to an already deployed contract.
func bindPivotTopic(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := PivotTopicMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PivotTopic *PivotTopicRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _PivotTopic.Contract.PivotTopicCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PivotTopic *PivotTopicRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PivotTopic.Contract.PivotTopicTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PivotTopic *PivotTopicRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PivotTopic.Contract.PivotTopicTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PivotTopic *PivotTopicCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _PivotTopic.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PivotTopic *PivotTopicTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PivotTopic.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PivotTopic *PivotTopicTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PivotTopic.Contract.contract.Transact(opts, method, params...)
}

// TopicHash is a free data retrieval call binding the contract method 0xb4d2e8be.
//
// Solidity: function _topicHash(uint256 topicId) view returns(bytes32)
func (_PivotTopic *PivotTopicCaller) TopicHash(opts *bind.CallOpts, topicId *big.Int) ([32]byte, error) {
	var out []interface{}
	err := _PivotTopic.contract.Call(opts, &out, "_topicHash", topicId)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// TopicHash is a free data retrieval call binding the contract method 0xb4d2e8be.
//
// Solidity: function _topicHash(uint256 topicId) view returns(bytes32)
func (_PivotTopic *PivotTopicSession) TopicHash(topicId *big.Int) ([32]byte, error) {
	return _PivotTopic.Contract.TopicHash(&_PivotTopic.CallOpts, topicId)
}

// TopicHash is a free data retrieval call binding the contract method 0xb4d2e8be.
//
// Solidity: function _topicHash(uint256 topicId) view returns(bytes32)
func (_PivotTopic *PivotTopicCallerSession) TopicHash(topicId *big.Int) ([32]byte, error) {
	return _PivotTopic.Contract.TopicHash(&_PivotTopic.CallOpts, topicId)
}

// TotalBalance is a free data retrieval call binding the contract method 0x20781e17.
//
// Solidity: function _totalBalance(uint256 topicId) view returns(uint256)
func (_PivotTopic *PivotTopicCaller) TotalBalance(opts *bind.CallOpts, topicId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _PivotTopic.contract.Call(opts, &out, "_totalBalance", topicId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalBalance is a free data retrieval call binding the contract method 0x20781e17.
//
// Solidity: function _totalBalance(uint256 topicId) view returns(uint256)
func (_PivotTopic *PivotTopicSession) TotalBalance(topicId *big.Int) (*big.Int, error) {
	return _PivotTopic.Contract.TotalBalance(&_PivotTopic.CallOpts, topicId)
}

// TotalBalance is a free data retrieval call binding the contract method 0x20781e17.
//
// Solidity: function _totalBalance(uint256 topicId) view returns(uint256)
func (_PivotTopic *PivotTopicCallerSession) TotalBalance(topicId *big.Int) (*big.Int, error) {
	return _PivotTopic.Contract.TotalBalance(&_PivotTopic.CallOpts, topicId)
}

// TotalCommission is a free data retrieval call binding the contract method 0x356b4ebb.
//
// Solidity: function _totalCommission(uint256 topicId) view returns(uint256)
func (_PivotTopic *PivotTopicCaller) TotalCommission(opts *bind.CallOpts, topicId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _PivotTopic.contract.Call(opts, &out, "_totalCommission", topicId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalCommission is a free data retrieval call binding the contract method 0x356b4ebb.
//
// Solidity: function _totalCommission(uint256 topicId) view returns(uint256)
func (_PivotTopic *PivotTopicSession) TotalCommission(topicId *big.Int) (*big.Int, error) {
	return _PivotTopic.Contract.TotalCommission(&_PivotTopic.CallOpts, topicId)
}

// TotalCommission is a free data retrieval call binding the contract method 0x356b4ebb.
//
// Solidity: function _totalCommission(uint256 topicId) view returns(uint256)
func (_PivotTopic *PivotTopicCallerSession) TotalCommission(topicId *big.Int) (*big.Int, error) {
	return _PivotTopic.Contract.TotalCommission(&_PivotTopic.CallOpts, topicId)
}

// GetFixedInvestment is a free data retrieval call binding the contract method 0x87ff5aa0.
//
// Solidity: function getFixedInvestment(uint256 topicId) view returns(uint256)
func (_PivotTopic *PivotTopicCaller) GetFixedInvestment(opts *bind.CallOpts, topicId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _PivotTopic.contract.Call(opts, &out, "getFixedInvestment", topicId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetFixedInvestment is a free data retrieval call binding the contract method 0x87ff5aa0.
//
// Solidity: function getFixedInvestment(uint256 topicId) view returns(uint256)
func (_PivotTopic *PivotTopicSession) GetFixedInvestment(topicId *big.Int) (*big.Int, error) {
	return _PivotTopic.Contract.GetFixedInvestment(&_PivotTopic.CallOpts, topicId)
}

// GetFixedInvestment is a free data retrieval call binding the contract method 0x87ff5aa0.
//
// Solidity: function getFixedInvestment(uint256 topicId) view returns(uint256)
func (_PivotTopic *PivotTopicCallerSession) GetFixedInvestment(topicId *big.Int) (*big.Int, error) {
	return _PivotTopic.Contract.GetFixedInvestment(&_PivotTopic.CallOpts, topicId)
}

// GetIncome is a free data retrieval call binding the contract method 0x6773f382.
//
// Solidity: function getIncome(address investor, uint256 topicId) view returns(uint256)
func (_PivotTopic *PivotTopicCaller) GetIncome(opts *bind.CallOpts, investor common.Address, topicId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _PivotTopic.contract.Call(opts, &out, "getIncome", investor, topicId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetIncome is a free data retrieval call binding the contract method 0x6773f382.
//
// Solidity: function getIncome(address investor, uint256 topicId) view returns(uint256)
func (_PivotTopic *PivotTopicSession) GetIncome(investor common.Address, topicId *big.Int) (*big.Int, error) {
	return _PivotTopic.Contract.GetIncome(&_PivotTopic.CallOpts, investor, topicId)
}

// GetIncome is a free data retrieval call binding the contract method 0x6773f382.
//
// Solidity: function getIncome(address investor, uint256 topicId) view returns(uint256)
func (_PivotTopic *PivotTopicCallerSession) GetIncome(investor common.Address, topicId *big.Int) (*big.Int, error) {
	return _PivotTopic.Contract.GetIncome(&_PivotTopic.CallOpts, investor, topicId)
}

// GetInvestment is a free data retrieval call binding the contract method 0xdf65dde3.
//
// Solidity: function getInvestment(address investor, uint256 topicId) view returns(uint256)
func (_PivotTopic *PivotTopicCaller) GetInvestment(opts *bind.CallOpts, investor common.Address, topicId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _PivotTopic.contract.Call(opts, &out, "getInvestment", investor, topicId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetInvestment is a free data retrieval call binding the contract method 0xdf65dde3.
//
// Solidity: function getInvestment(address investor, uint256 topicId) view returns(uint256)
func (_PivotTopic *PivotTopicSession) GetInvestment(investor common.Address, topicId *big.Int) (*big.Int, error) {
	return _PivotTopic.Contract.GetInvestment(&_PivotTopic.CallOpts, investor, topicId)
}

// GetInvestment is a free data retrieval call binding the contract method 0xdf65dde3.
//
// Solidity: function getInvestment(address investor, uint256 topicId) view returns(uint256)
func (_PivotTopic *PivotTopicCallerSession) GetInvestment(investor common.Address, topicId *big.Int) (*big.Int, error) {
	return _PivotTopic.Contract.GetInvestment(&_PivotTopic.CallOpts, investor, topicId)
}

// GetPromoter is a free data retrieval call binding the contract method 0x12426f30.
//
// Solidity: function getPromoter(uint256 topicId) view returns(address)
func (_PivotTopic *PivotTopicCaller) GetPromoter(opts *bind.CallOpts, topicId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _PivotTopic.contract.Call(opts, &out, "getPromoter", topicId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetPromoter is a free data retrieval call binding the contract method 0x12426f30.
//
// Solidity: function getPromoter(uint256 topicId) view returns(address)
func (_PivotTopic *PivotTopicSession) GetPromoter(topicId *big.Int) (common.Address, error) {
	return _PivotTopic.Contract.GetPromoter(&_PivotTopic.CallOpts, topicId)
}

// GetPromoter is a free data retrieval call binding the contract method 0x12426f30.
//
// Solidity: function getPromoter(uint256 topicId) view returns(address)
func (_PivotTopic *PivotTopicCallerSession) GetPromoter(topicId *big.Int) (common.Address, error) {
	return _PivotTopic.Contract.GetPromoter(&_PivotTopic.CallOpts, topicId)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_PivotTopic *PivotTopicCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _PivotTopic.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_PivotTopic *PivotTopicSession) Owner() (common.Address, error) {
	return _PivotTopic.Contract.Owner(&_PivotTopic.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_PivotTopic *PivotTopicCallerSession) Owner() (common.Address, error) {
	return _PivotTopic.Contract.Owner(&_PivotTopic.CallOpts)
}

// SbtAddress is a free data retrieval call binding the contract method 0x989147b3.
//
// Solidity: function sbtAddress() view returns(address)
func (_PivotTopic *PivotTopicCaller) SbtAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _PivotTopic.contract.Call(opts, &out, "sbtAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// SbtAddress is a free data retrieval call binding the contract method 0x989147b3.
//
// Solidity: function sbtAddress() view returns(address)
func (_PivotTopic *PivotTopicSession) SbtAddress() (common.Address, error) {
	return _PivotTopic.Contract.SbtAddress(&_PivotTopic.CallOpts)
}

// SbtAddress is a free data retrieval call binding the contract method 0x989147b3.
//
// Solidity: function sbtAddress() view returns(address)
func (_PivotTopic *PivotTopicCallerSession) SbtAddress() (common.Address, error) {
	return _PivotTopic.Contract.SbtAddress(&_PivotTopic.CallOpts)
}

// SbtContract is a free data retrieval call binding the contract method 0x06158360.
//
// Solidity: function sbtContract() view returns(address)
func (_PivotTopic *PivotTopicCaller) SbtContract(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _PivotTopic.contract.Call(opts, &out, "sbtContract")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// SbtContract is a free data retrieval call binding the contract method 0x06158360.
//
// Solidity: function sbtContract() view returns(address)
func (_PivotTopic *PivotTopicSession) SbtContract() (common.Address, error) {
	return _PivotTopic.Contract.SbtContract(&_PivotTopic.CallOpts)
}

// SbtContract is a free data retrieval call binding the contract method 0x06158360.
//
// Solidity: function sbtContract() view returns(address)
func (_PivotTopic *PivotTopicCallerSession) SbtContract() (common.Address, error) {
	return _PivotTopic.Contract.SbtContract(&_PivotTopic.CallOpts)
}

// TopicCoin is a free data retrieval call binding the contract method 0xee6c6f7c.
//
// Solidity: function topicCoin(uint256 topicId) view returns(address)
func (_PivotTopic *PivotTopicCaller) TopicCoin(opts *bind.CallOpts, topicId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _PivotTopic.contract.Call(opts, &out, "topicCoin", topicId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// TopicCoin is a free data retrieval call binding the contract method 0xee6c6f7c.
//
// Solidity: function topicCoin(uint256 topicId) view returns(address)
func (_PivotTopic *PivotTopicSession) TopicCoin(topicId *big.Int) (common.Address, error) {
	return _PivotTopic.Contract.TopicCoin(&_PivotTopic.CallOpts, topicId)
}

// TopicCoin is a free data retrieval call binding the contract method 0xee6c6f7c.
//
// Solidity: function topicCoin(uint256 topicId) view returns(address)
func (_PivotTopic *PivotTopicCallerSession) TopicCoin(topicId *big.Int) (common.Address, error) {
	return _PivotTopic.Contract.TopicCoin(&_PivotTopic.CallOpts, topicId)
}

// CreateTopic is a paid mutator transaction binding the contract method 0x6011dd7c.
//
// Solidity: function createTopic(uint256 amount, address erc20Address, bytes32 topicHash) returns()
func (_PivotTopic *PivotTopicTransactor) CreateTopic(opts *bind.TransactOpts, amount *big.Int, erc20Address common.Address, topicHash [32]byte) (*types.Transaction, error) {
	return _PivotTopic.contract.Transact(opts, "createTopic", amount, erc20Address, topicHash)
}

// CreateTopic is a paid mutator transaction binding the contract method 0x6011dd7c.
//
// Solidity: function createTopic(uint256 amount, address erc20Address, bytes32 topicHash) returns()
func (_PivotTopic *PivotTopicSession) CreateTopic(amount *big.Int, erc20Address common.Address, topicHash [32]byte) (*types.Transaction, error) {
	return _PivotTopic.Contract.CreateTopic(&_PivotTopic.TransactOpts, amount, erc20Address, topicHash)
}

// CreateTopic is a paid mutator transaction binding the contract method 0x6011dd7c.
//
// Solidity: function createTopic(uint256 amount, address erc20Address, bytes32 topicHash) returns()
func (_PivotTopic *PivotTopicTransactorSession) CreateTopic(amount *big.Int, erc20Address common.Address, topicHash [32]byte) (*types.Transaction, error) {
	return _PivotTopic.Contract.CreateTopic(&_PivotTopic.TransactOpts, amount, erc20Address, topicHash)
}

// Invest is a paid mutator transaction binding the contract method 0xd87aa643.
//
// Solidity: function invest(uint256 topicId, uint256 amount) returns()
func (_PivotTopic *PivotTopicTransactor) Invest(opts *bind.TransactOpts, topicId *big.Int, amount *big.Int) (*types.Transaction, error) {
	return _PivotTopic.contract.Transact(opts, "invest", topicId, amount)
}

// Invest is a paid mutator transaction binding the contract method 0xd87aa643.
//
// Solidity: function invest(uint256 topicId, uint256 amount) returns()
func (_PivotTopic *PivotTopicSession) Invest(topicId *big.Int, amount *big.Int) (*types.Transaction, error) {
	return _PivotTopic.Contract.Invest(&_PivotTopic.TransactOpts, topicId, amount)
}

// Invest is a paid mutator transaction binding the contract method 0xd87aa643.
//
// Solidity: function invest(uint256 topicId, uint256 amount) returns()
func (_PivotTopic *PivotTopicTransactorSession) Invest(topicId *big.Int, amount *big.Int) (*types.Transaction, error) {
	return _PivotTopic.Contract.Invest(&_PivotTopic.TransactOpts, topicId, amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 topicId) returns()
func (_PivotTopic *PivotTopicTransactor) Withdraw(opts *bind.TransactOpts, topicId *big.Int) (*types.Transaction, error) {
	return _PivotTopic.contract.Transact(opts, "withdraw", topicId)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 topicId) returns()
func (_PivotTopic *PivotTopicSession) Withdraw(topicId *big.Int) (*types.Transaction, error) {
	return _PivotTopic.Contract.Withdraw(&_PivotTopic.TransactOpts, topicId)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 topicId) returns()
func (_PivotTopic *PivotTopicTransactorSession) Withdraw(topicId *big.Int) (*types.Transaction, error) {
	return _PivotTopic.Contract.Withdraw(&_PivotTopic.TransactOpts, topicId)
}

// WithdrawCommission is a paid mutator transaction binding the contract method 0xbb2d3e0a.
//
// Solidity: function withdrawCommission(uint256 amount, uint256 topicId) returns()
func (_PivotTopic *PivotTopicTransactor) WithdrawCommission(opts *bind.TransactOpts, amount *big.Int, topicId *big.Int) (*types.Transaction, error) {
	return _PivotTopic.contract.Transact(opts, "withdrawCommission", amount, topicId)
}

// WithdrawCommission is a paid mutator transaction binding the contract method 0xbb2d3e0a.
//
// Solidity: function withdrawCommission(uint256 amount, uint256 topicId) returns()
func (_PivotTopic *PivotTopicSession) WithdrawCommission(amount *big.Int, topicId *big.Int) (*types.Transaction, error) {
	return _PivotTopic.Contract.WithdrawCommission(&_PivotTopic.TransactOpts, amount, topicId)
}

// WithdrawCommission is a paid mutator transaction binding the contract method 0xbb2d3e0a.
//
// Solidity: function withdrawCommission(uint256 amount, uint256 topicId) returns()
func (_PivotTopic *PivotTopicTransactorSession) WithdrawCommission(amount *big.Int, topicId *big.Int) (*types.Transaction, error) {
	return _PivotTopic.Contract.WithdrawCommission(&_PivotTopic.TransactOpts, amount, topicId)
}

// PivotTopicCreateTopicIterator is returned from FilterCreateTopic and is used to iterate over the raw logs and unpacked data for CreateTopic events raised by the PivotTopic contract.
type PivotTopicCreateTopicIterator struct {
	Event *PivotTopicCreateTopic // Event containing the contract specifics and raw log

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
func (it *PivotTopicCreateTopicIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PivotTopicCreateTopic)
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
		it.Event = new(PivotTopicCreateTopic)
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
func (it *PivotTopicCreateTopicIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PivotTopicCreateTopicIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PivotTopicCreateTopic represents a CreateTopic event raised by the PivotTopic contract.
type PivotTopicCreateTopic struct {
	Promoter   common.Address
	TopicId    *big.Int
	Investment *big.Int
	Position   *big.Int
	Nonce      *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterCreateTopic is a free log retrieval operation binding the contract event 0x51c53024b6cd374b1ecc00580030bfbe4930bec86dda0171d478cd003c76dfc1.
//
// Solidity: event CreateTopic(address indexed promoter, uint256 topicId, uint256 investment, uint256 position, uint256 nonce)
func (_PivotTopic *PivotTopicFilterer) FilterCreateTopic(opts *bind.FilterOpts, promoter []common.Address) (*PivotTopicCreateTopicIterator, error) {

	var promoterRule []interface{}
	for _, promoterItem := range promoter {
		promoterRule = append(promoterRule, promoterItem)
	}

	logs, sub, err := _PivotTopic.contract.FilterLogs(opts, "CreateTopic", promoterRule)
	if err != nil {
		return nil, err
	}
	return &PivotTopicCreateTopicIterator{contract: _PivotTopic.contract, event: "CreateTopic", logs: logs, sub: sub}, nil
}

// WatchCreateTopic is a free log subscription operation binding the contract event 0x51c53024b6cd374b1ecc00580030bfbe4930bec86dda0171d478cd003c76dfc1.
//
// Solidity: event CreateTopic(address indexed promoter, uint256 topicId, uint256 investment, uint256 position, uint256 nonce)
func (_PivotTopic *PivotTopicFilterer) WatchCreateTopic(opts *bind.WatchOpts, sink chan<- *PivotTopicCreateTopic, promoter []common.Address) (event.Subscription, error) {

	var promoterRule []interface{}
	for _, promoterItem := range promoter {
		promoterRule = append(promoterRule, promoterItem)
	}

	logs, sub, err := _PivotTopic.contract.WatchLogs(opts, "CreateTopic", promoterRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PivotTopicCreateTopic)
				if err := _PivotTopic.contract.UnpackLog(event, "CreateTopic", log); err != nil {
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

// ParseCreateTopic is a log parse operation binding the contract event 0x51c53024b6cd374b1ecc00580030bfbe4930bec86dda0171d478cd003c76dfc1.
//
// Solidity: event CreateTopic(address indexed promoter, uint256 topicId, uint256 investment, uint256 position, uint256 nonce)
func (_PivotTopic *PivotTopicFilterer) ParseCreateTopic(log types.Log) (*PivotTopicCreateTopic, error) {
	event := new(PivotTopicCreateTopic)
	if err := _PivotTopic.contract.UnpackLog(event, "CreateTopic", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PivotTopicInvestIterator is returned from FilterInvest and is used to iterate over the raw logs and unpacked data for Invest events raised by the PivotTopic contract.
type PivotTopicInvestIterator struct {
	Event *PivotTopicInvest // Event containing the contract specifics and raw log

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
func (it *PivotTopicInvestIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PivotTopicInvest)
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
		it.Event = new(PivotTopicInvest)
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
func (it *PivotTopicInvestIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PivotTopicInvestIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PivotTopicInvest represents a Invest event raised by the PivotTopic contract.
type PivotTopicInvest struct {
	Investor common.Address
	TopicId  *big.Int
	Amount   *big.Int
	Position *big.Int
	Nonce    *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterInvest is a free log retrieval operation binding the contract event 0x2f95f6eee5b8fb53c5446c9dede0db4adba4dd704a6f55fb19dffe3462f3eae8.
//
// Solidity: event Invest(address indexed investor, uint256 indexed topicId, uint256 amount, uint256 position, uint256 nonce)
func (_PivotTopic *PivotTopicFilterer) FilterInvest(opts *bind.FilterOpts, investor []common.Address, topicId []*big.Int) (*PivotTopicInvestIterator, error) {

	var investorRule []interface{}
	for _, investorItem := range investor {
		investorRule = append(investorRule, investorItem)
	}
	var topicIdRule []interface{}
	for _, topicIdItem := range topicId {
		topicIdRule = append(topicIdRule, topicIdItem)
	}

	logs, sub, err := _PivotTopic.contract.FilterLogs(opts, "Invest", investorRule, topicIdRule)
	if err != nil {
		return nil, err
	}
	return &PivotTopicInvestIterator{contract: _PivotTopic.contract, event: "Invest", logs: logs, sub: sub}, nil
}

// WatchInvest is a free log subscription operation binding the contract event 0x2f95f6eee5b8fb53c5446c9dede0db4adba4dd704a6f55fb19dffe3462f3eae8.
//
// Solidity: event Invest(address indexed investor, uint256 indexed topicId, uint256 amount, uint256 position, uint256 nonce)
func (_PivotTopic *PivotTopicFilterer) WatchInvest(opts *bind.WatchOpts, sink chan<- *PivotTopicInvest, investor []common.Address, topicId []*big.Int) (event.Subscription, error) {

	var investorRule []interface{}
	for _, investorItem := range investor {
		investorRule = append(investorRule, investorItem)
	}
	var topicIdRule []interface{}
	for _, topicIdItem := range topicId {
		topicIdRule = append(topicIdRule, topicIdItem)
	}

	logs, sub, err := _PivotTopic.contract.WatchLogs(opts, "Invest", investorRule, topicIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PivotTopicInvest)
				if err := _PivotTopic.contract.UnpackLog(event, "Invest", log); err != nil {
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

// ParseInvest is a log parse operation binding the contract event 0x2f95f6eee5b8fb53c5446c9dede0db4adba4dd704a6f55fb19dffe3462f3eae8.
//
// Solidity: event Invest(address indexed investor, uint256 indexed topicId, uint256 amount, uint256 position, uint256 nonce)
func (_PivotTopic *PivotTopicFilterer) ParseInvest(log types.Log) (*PivotTopicInvest, error) {
	event := new(PivotTopicInvest)
	if err := _PivotTopic.contract.UnpackLog(event, "Invest", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PivotTopicWithdrawIterator is returned from FilterWithdraw and is used to iterate over the raw logs and unpacked data for Withdraw events raised by the PivotTopic contract.
type PivotTopicWithdrawIterator struct {
	Event *PivotTopicWithdraw // Event containing the contract specifics and raw log

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
func (it *PivotTopicWithdrawIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PivotTopicWithdraw)
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
		it.Event = new(PivotTopicWithdraw)
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
func (it *PivotTopicWithdrawIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PivotTopicWithdrawIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PivotTopicWithdraw represents a Withdraw event raised by the PivotTopic contract.
type PivotTopicWithdraw struct {
	To     common.Address
	Amount *big.Int
	Nonce  *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterWithdraw is a free log retrieval operation binding the contract event 0xf279e6a1f5e320cca91135676d9cb6e44ca8a08c0b88342bcdb1144f6511b568.
//
// Solidity: event Withdraw(address indexed to, uint256 amount, uint256 nonce)
func (_PivotTopic *PivotTopicFilterer) FilterWithdraw(opts *bind.FilterOpts, to []common.Address) (*PivotTopicWithdrawIterator, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _PivotTopic.contract.FilterLogs(opts, "Withdraw", toRule)
	if err != nil {
		return nil, err
	}
	return &PivotTopicWithdrawIterator{contract: _PivotTopic.contract, event: "Withdraw", logs: logs, sub: sub}, nil
}

// WatchWithdraw is a free log subscription operation binding the contract event 0xf279e6a1f5e320cca91135676d9cb6e44ca8a08c0b88342bcdb1144f6511b568.
//
// Solidity: event Withdraw(address indexed to, uint256 amount, uint256 nonce)
func (_PivotTopic *PivotTopicFilterer) WatchWithdraw(opts *bind.WatchOpts, sink chan<- *PivotTopicWithdraw, to []common.Address) (event.Subscription, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _PivotTopic.contract.WatchLogs(opts, "Withdraw", toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PivotTopicWithdraw)
				if err := _PivotTopic.contract.UnpackLog(event, "Withdraw", log); err != nil {
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
// Solidity: event Withdraw(address indexed to, uint256 amount, uint256 nonce)
func (_PivotTopic *PivotTopicFilterer) ParseWithdraw(log types.Log) (*PivotTopicWithdraw, error) {
	event := new(PivotTopicWithdraw)
	if err := _PivotTopic.contract.UnpackLog(event, "Withdraw", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PivotTopicWithdrawCommissionIterator is returned from FilterWithdrawCommission and is used to iterate over the raw logs and unpacked data for WithdrawCommission events raised by the PivotTopic contract.
type PivotTopicWithdrawCommissionIterator struct {
	Event *PivotTopicWithdrawCommission // Event containing the contract specifics and raw log

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
func (it *PivotTopicWithdrawCommissionIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PivotTopicWithdrawCommission)
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
		it.Event = new(PivotTopicWithdrawCommission)
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
func (it *PivotTopicWithdrawCommissionIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PivotTopicWithdrawCommissionIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PivotTopicWithdrawCommission represents a WithdrawCommission event raised by the PivotTopic contract.
type PivotTopicWithdrawCommission struct {
	Owner  common.Address
	Amount *big.Int
	Nonce  *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterWithdrawCommission is a free log retrieval operation binding the contract event 0xf782b61788afc16b529174597fea43c57ebd5d24d8d8f49d42b39c28948550bd.
//
// Solidity: event WithdrawCommission(address indexed owner, uint256 amount, uint256 nonce)
func (_PivotTopic *PivotTopicFilterer) FilterWithdrawCommission(opts *bind.FilterOpts, owner []common.Address) (*PivotTopicWithdrawCommissionIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _PivotTopic.contract.FilterLogs(opts, "WithdrawCommission", ownerRule)
	if err != nil {
		return nil, err
	}
	return &PivotTopicWithdrawCommissionIterator{contract: _PivotTopic.contract, event: "WithdrawCommission", logs: logs, sub: sub}, nil
}

// WatchWithdrawCommission is a free log subscription operation binding the contract event 0xf782b61788afc16b529174597fea43c57ebd5d24d8d8f49d42b39c28948550bd.
//
// Solidity: event WithdrawCommission(address indexed owner, uint256 amount, uint256 nonce)
func (_PivotTopic *PivotTopicFilterer) WatchWithdrawCommission(opts *bind.WatchOpts, sink chan<- *PivotTopicWithdrawCommission, owner []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _PivotTopic.contract.WatchLogs(opts, "WithdrawCommission", ownerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PivotTopicWithdrawCommission)
				if err := _PivotTopic.contract.UnpackLog(event, "WithdrawCommission", log); err != nil {
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

// ParseWithdrawCommission is a log parse operation binding the contract event 0xf782b61788afc16b529174597fea43c57ebd5d24d8d8f49d42b39c28948550bd.
//
// Solidity: event WithdrawCommission(address indexed owner, uint256 amount, uint256 nonce)
func (_PivotTopic *PivotTopicFilterer) ParseWithdrawCommission(log types.Log) (*PivotTopicWithdrawCommission, error) {
	event := new(PivotTopicWithdrawCommission)
	if err := _PivotTopic.contract.UnpackLog(event, "WithdrawCommission", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
