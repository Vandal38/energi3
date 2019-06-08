// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package abi

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

// StakerRewardV1ABI is the input ABI used to generate the binding from.
const StakerRewardV1ABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"_newImpl\",\"type\":\"address\"}],\"name\":\"destroy\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_blockNumber\",\"type\":\"uint256\"}],\"name\":\"getReward\",\"outputs\":[{\"name\":\"amount\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"reward\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"\",\"type\":\"address\"}],\"name\":\"migrate\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"proxy\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"_proxy\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"fallback\"}]"

// StakerRewardV1Bin is the compiled bytecode used for deploying new contracts.
const StakerRewardV1Bin = `608060405234801561001057600080fd5b506040516103cf3803806103cf8339818101604052602081101561003357600080fd5b5051600080546001600160a01b039092166001600160a01b031990921691909117905561036a806100656000396000f3fe6080604052600436106100595760003560e01c8063a9fb763c11610043578063a9fb763c1461013e578063ce5494bb1461015b578063ec5568891461019b57610059565b8062f55d9d146100c05780631c4b774b14610102575b604080517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152600d60248201527f4e6f7420737570706f7274656400000000000000000000000000000000000000604482015290519081900360640190fd5b3480156100cc57600080fd5b50610100600480360360208110156100e357600080fd5b503573ffffffffffffffffffffffffffffffffffffffff166101d9565b005b34801561010e57600080fd5b5061012c6004803603602081101561012557600080fd5b5035610278565b60408051918252519081900360200190f35b6101006004803603602081101561015457600080fd5b5035610290565b34801561016757600080fd5b506101006004803603602081101561017e57600080fd5b503573ffffffffffffffffffffffffffffffffffffffff16610293565b3480156101a757600080fd5b506101b0610319565b6040805173ffffffffffffffffffffffffffffffffffffffff9092168252519081900360200190f35b60005473ffffffffffffffffffffffffffffffffffffffff16331461025f57604080517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152600960248201527f4e6f742070726f78790000000000000000000000000000000000000000000000604482015290519081900360640190fd5b8073ffffffffffffffffffffffffffffffffffffffff16ff5b6000811561028b5750671fa42feb87e400005b919050565b50565b60005473ffffffffffffffffffffffffffffffffffffffff16331461029057604080517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152600960248201527f4e6f742070726f78790000000000000000000000000000000000000000000000604482015290519081900360640190fd5b60005473ffffffffffffffffffffffffffffffffffffffff168156fea265627a7a7230582082e4357dd3c131959516a9089e7d187d71a2c4aa614d49610dd8e66d00f2f7ec64736f6c63430005090032`

// DeployStakerRewardV1 deploys a new Ethereum contract, binding an instance of StakerRewardV1 to it.
func DeployStakerRewardV1(auth *bind.TransactOpts, backend bind.ContractBackend, _proxy common.Address) (common.Address, *types.Transaction, *StakerRewardV1, error) {
	parsed, err := abi.JSON(strings.NewReader(StakerRewardV1ABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(StakerRewardV1Bin), backend, _proxy)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &StakerRewardV1{StakerRewardV1Caller: StakerRewardV1Caller{contract: contract}, StakerRewardV1Transactor: StakerRewardV1Transactor{contract: contract}, StakerRewardV1Filterer: StakerRewardV1Filterer{contract: contract}}, nil
}

// StakerRewardV1 is an auto generated Go binding around an Ethereum contract.
type StakerRewardV1 struct {
	StakerRewardV1Caller     // Read-only binding to the contract
	StakerRewardV1Transactor // Write-only binding to the contract
	StakerRewardV1Filterer   // Log filterer for contract events
}

// StakerRewardV1Caller is an auto generated read-only Go binding around an Ethereum contract.
type StakerRewardV1Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StakerRewardV1Transactor is an auto generated write-only Go binding around an Ethereum contract.
type StakerRewardV1Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StakerRewardV1Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type StakerRewardV1Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StakerRewardV1Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type StakerRewardV1Session struct {
	Contract     *StakerRewardV1   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// StakerRewardV1CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type StakerRewardV1CallerSession struct {
	Contract *StakerRewardV1Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// StakerRewardV1TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type StakerRewardV1TransactorSession struct {
	Contract     *StakerRewardV1Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// StakerRewardV1Raw is an auto generated low-level Go binding around an Ethereum contract.
type StakerRewardV1Raw struct {
	Contract *StakerRewardV1 // Generic contract binding to access the raw methods on
}

// StakerRewardV1CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type StakerRewardV1CallerRaw struct {
	Contract *StakerRewardV1Caller // Generic read-only contract binding to access the raw methods on
}

// StakerRewardV1TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type StakerRewardV1TransactorRaw struct {
	Contract *StakerRewardV1Transactor // Generic write-only contract binding to access the raw methods on
}

// NewStakerRewardV1 creates a new instance of StakerRewardV1, bound to a specific deployed contract.
func NewStakerRewardV1(address common.Address, backend bind.ContractBackend) (*StakerRewardV1, error) {
	contract, err := bindStakerRewardV1(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &StakerRewardV1{StakerRewardV1Caller: StakerRewardV1Caller{contract: contract}, StakerRewardV1Transactor: StakerRewardV1Transactor{contract: contract}, StakerRewardV1Filterer: StakerRewardV1Filterer{contract: contract}}, nil
}

// NewStakerRewardV1Caller creates a new read-only instance of StakerRewardV1, bound to a specific deployed contract.
func NewStakerRewardV1Caller(address common.Address, caller bind.ContractCaller) (*StakerRewardV1Caller, error) {
	contract, err := bindStakerRewardV1(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &StakerRewardV1Caller{contract: contract}, nil
}

// NewStakerRewardV1Transactor creates a new write-only instance of StakerRewardV1, bound to a specific deployed contract.
func NewStakerRewardV1Transactor(address common.Address, transactor bind.ContractTransactor) (*StakerRewardV1Transactor, error) {
	contract, err := bindStakerRewardV1(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &StakerRewardV1Transactor{contract: contract}, nil
}

// NewStakerRewardV1Filterer creates a new log filterer instance of StakerRewardV1, bound to a specific deployed contract.
func NewStakerRewardV1Filterer(address common.Address, filterer bind.ContractFilterer) (*StakerRewardV1Filterer, error) {
	contract, err := bindStakerRewardV1(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &StakerRewardV1Filterer{contract: contract}, nil
}

// bindStakerRewardV1 binds a generic wrapper to an already deployed contract.
func bindStakerRewardV1(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(StakerRewardV1ABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_StakerRewardV1 *StakerRewardV1Raw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _StakerRewardV1.Contract.StakerRewardV1Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_StakerRewardV1 *StakerRewardV1Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StakerRewardV1.Contract.StakerRewardV1Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_StakerRewardV1 *StakerRewardV1Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _StakerRewardV1.Contract.StakerRewardV1Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_StakerRewardV1 *StakerRewardV1CallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _StakerRewardV1.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_StakerRewardV1 *StakerRewardV1TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StakerRewardV1.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_StakerRewardV1 *StakerRewardV1TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _StakerRewardV1.Contract.contract.Transact(opts, method, params...)
}

// GetReward is a free data retrieval call binding the contract method 0x1c4b774b.
//
// Solidity: function getReward(uint256 _blockNumber) constant returns(uint256 amount)
func (_StakerRewardV1 *StakerRewardV1Caller) GetReward(opts *bind.CallOpts, _blockNumber *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _StakerRewardV1.contract.Call(opts, out, "getReward", _blockNumber)
	return *ret0, err
}

// GetReward is a free data retrieval call binding the contract method 0x1c4b774b.
//
// Solidity: function getReward(uint256 _blockNumber) constant returns(uint256 amount)
func (_StakerRewardV1 *StakerRewardV1Session) GetReward(_blockNumber *big.Int) (*big.Int, error) {
	return _StakerRewardV1.Contract.GetReward(&_StakerRewardV1.CallOpts, _blockNumber)
}

// GetReward is a free data retrieval call binding the contract method 0x1c4b774b.
//
// Solidity: function getReward(uint256 _blockNumber) constant returns(uint256 amount)
func (_StakerRewardV1 *StakerRewardV1CallerSession) GetReward(_blockNumber *big.Int) (*big.Int, error) {
	return _StakerRewardV1.Contract.GetReward(&_StakerRewardV1.CallOpts, _blockNumber)
}

// Proxy is a free data retrieval call binding the contract method 0xec556889.
//
// Solidity: function proxy() constant returns(address)
func (_StakerRewardV1 *StakerRewardV1Caller) Proxy(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _StakerRewardV1.contract.Call(opts, out, "proxy")
	return *ret0, err
}

// Proxy is a free data retrieval call binding the contract method 0xec556889.
//
// Solidity: function proxy() constant returns(address)
func (_StakerRewardV1 *StakerRewardV1Session) Proxy() (common.Address, error) {
	return _StakerRewardV1.Contract.Proxy(&_StakerRewardV1.CallOpts)
}

// Proxy is a free data retrieval call binding the contract method 0xec556889.
//
// Solidity: function proxy() constant returns(address)
func (_StakerRewardV1 *StakerRewardV1CallerSession) Proxy() (common.Address, error) {
	return _StakerRewardV1.Contract.Proxy(&_StakerRewardV1.CallOpts)
}

// Destroy is a paid mutator transaction binding the contract method 0x00f55d9d.
//
// Solidity: function destroy(address _newImpl) returns()
func (_StakerRewardV1 *StakerRewardV1Transactor) Destroy(opts *bind.TransactOpts, _newImpl common.Address) (*types.Transaction, error) {
	return _StakerRewardV1.contract.Transact(opts, "destroy", _newImpl)
}

// Destroy is a paid mutator transaction binding the contract method 0x00f55d9d.
//
// Solidity: function destroy(address _newImpl) returns()
func (_StakerRewardV1 *StakerRewardV1Session) Destroy(_newImpl common.Address) (*types.Transaction, error) {
	return _StakerRewardV1.Contract.Destroy(&_StakerRewardV1.TransactOpts, _newImpl)
}

// Destroy is a paid mutator transaction binding the contract method 0x00f55d9d.
//
// Solidity: function destroy(address _newImpl) returns()
func (_StakerRewardV1 *StakerRewardV1TransactorSession) Destroy(_newImpl common.Address) (*types.Transaction, error) {
	return _StakerRewardV1.Contract.Destroy(&_StakerRewardV1.TransactOpts, _newImpl)
}

// Migrate is a paid mutator transaction binding the contract method 0xce5494bb.
//
// Solidity: function migrate(address ) returns()
func (_StakerRewardV1 *StakerRewardV1Transactor) Migrate(opts *bind.TransactOpts, arg0 common.Address) (*types.Transaction, error) {
	return _StakerRewardV1.contract.Transact(opts, "migrate", arg0)
}

// Migrate is a paid mutator transaction binding the contract method 0xce5494bb.
//
// Solidity: function migrate(address ) returns()
func (_StakerRewardV1 *StakerRewardV1Session) Migrate(arg0 common.Address) (*types.Transaction, error) {
	return _StakerRewardV1.Contract.Migrate(&_StakerRewardV1.TransactOpts, arg0)
}

// Migrate is a paid mutator transaction binding the contract method 0xce5494bb.
//
// Solidity: function migrate(address ) returns()
func (_StakerRewardV1 *StakerRewardV1TransactorSession) Migrate(arg0 common.Address) (*types.Transaction, error) {
	return _StakerRewardV1.Contract.Migrate(&_StakerRewardV1.TransactOpts, arg0)
}

// Reward is a paid mutator transaction binding the contract method 0xa9fb763c.
//
// Solidity: function reward(uint256 ) returns()
func (_StakerRewardV1 *StakerRewardV1Transactor) Reward(opts *bind.TransactOpts, arg0 *big.Int) (*types.Transaction, error) {
	return _StakerRewardV1.contract.Transact(opts, "reward", arg0)
}

// Reward is a paid mutator transaction binding the contract method 0xa9fb763c.
//
// Solidity: function reward(uint256 ) returns()
func (_StakerRewardV1 *StakerRewardV1Session) Reward(arg0 *big.Int) (*types.Transaction, error) {
	return _StakerRewardV1.Contract.Reward(&_StakerRewardV1.TransactOpts, arg0)
}

// Reward is a paid mutator transaction binding the contract method 0xa9fb763c.
//
// Solidity: function reward(uint256 ) returns()
func (_StakerRewardV1 *StakerRewardV1TransactorSession) Reward(arg0 *big.Int) (*types.Transaction, error) {
	return _StakerRewardV1.Contract.Reward(&_StakerRewardV1.TransactOpts, arg0)
}
