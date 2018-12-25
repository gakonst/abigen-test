// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package abigentest

import (
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/loomnetwork/go-loom/types"
)

// AbigentestABI is the input ABI used to generate the binding from.
const AbigentestABI = "[{\"constant\":false,\"inputs\":[],\"name\":\"x\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"b\",\"type\":\"uint256\"}],\"name\":\"x\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// Abigentest is an auto generated Go binding around an Ethereum contract.
type Abigentest struct {
	AbigentestCaller     // Read-only binding to the contract
	AbigentestTransactor // Write-only binding to the contract
	AbigentestFilterer   // Log filterer for contract events
}

// AbigentestCaller is an auto generated read-only Go binding around an Ethereum contract.
type AbigentestCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AbigentestTransactor is an auto generated write-only Go binding around an Ethereum contract.
type AbigentestTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AbigentestFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type AbigentestFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AbigentestSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type AbigentestSession struct {
	Contract     *Abigentest       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// AbigentestCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type AbigentestCallerSession struct {
	Contract *AbigentestCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// AbigentestTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type AbigentestTransactorSession struct {
	Contract     *AbigentestTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// AbigentestRaw is an auto generated low-level Go binding around an Ethereum contract.
type AbigentestRaw struct {
	Contract *Abigentest // Generic contract binding to access the raw methods on
}

// AbigentestCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type AbigentestCallerRaw struct {
	Contract *AbigentestCaller // Generic read-only contract binding to access the raw methods on
}

// AbigentestTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type AbigentestTransactorRaw struct {
	Contract *AbigentestTransactor // Generic write-only contract binding to access the raw methods on
}

// NewAbigentest creates a new instance of Abigentest, bound to a specific deployed contract.
func NewAbigentest(address common.Address, backend bind.ContractBackend) (*Abigentest, error) {
	contract, err := bindAbigentest(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Abigentest{AbigentestCaller: AbigentestCaller{contract: contract}, AbigentestTransactor: AbigentestTransactor{contract: contract}, AbigentestFilterer: AbigentestFilterer{contract: contract}}, nil
}

// NewAbigentestCaller creates a new read-only instance of Abigentest, bound to a specific deployed contract.
func NewAbigentestCaller(address common.Address, caller bind.ContractCaller) (*AbigentestCaller, error) {
	contract, err := bindAbigentest(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &AbigentestCaller{contract: contract}, nil
}

// NewAbigentestTransactor creates a new write-only instance of Abigentest, bound to a specific deployed contract.
func NewAbigentestTransactor(address common.Address, transactor bind.ContractTransactor) (*AbigentestTransactor, error) {
	contract, err := bindAbigentest(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &AbigentestTransactor{contract: contract}, nil
}

// NewAbigentestFilterer creates a new log filterer instance of Abigentest, bound to a specific deployed contract.
func NewAbigentestFilterer(address common.Address, filterer bind.ContractFilterer) (*AbigentestFilterer, error) {
	contract, err := bindAbigentest(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &AbigentestFilterer{contract: contract}, nil
}

// bindAbigentest binds a generic wrapper to an already deployed contract.
func bindAbigentest(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(AbigentestABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Abigentest *AbigentestRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Abigentest.Contract.AbigentestCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Abigentest *AbigentestRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Abigentest.Contract.AbigentestTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Abigentest *AbigentestRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Abigentest.Contract.AbigentestTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Abigentest *AbigentestCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Abigentest.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Abigentest *AbigentestTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Abigentest.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Abigentest *AbigentestTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Abigentest.Contract.contract.Transact(opts, method, params...)
}

// X is a paid mutator transaction binding the contract method 0x3f4263ef.
//
// Solidity: function x(b uint256) returns()
func (_Abigentest *AbigentestTransactor) X(opts *bind.TransactOpts, b *big.Int) (*types.Transaction, error) {
	return _Abigentest.contract.Transact(opts, "x", b)
}

// X is a paid mutator transaction binding the contract method 0x3f4263ef.
//
// Solidity: function x(b uint256) returns()
func (_Abigentest *AbigentestSession) X(b *big.Int) (*types.Transaction, error) {
	return _Abigentest.Contract.X(&_Abigentest.TransactOpts, b)
}

// X is a paid mutator transaction binding the contract method 0x3f4263ef.
//
// Solidity: function x(b uint256) returns()
func (_Abigentest *AbigentestTransactorSession) X(b *big.Int) (*types.Transaction, error) {
	return _Abigentest.Contract.X(&_Abigentest.TransactOpts, b)
}
