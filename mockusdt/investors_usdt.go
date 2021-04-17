// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package investors_usdt

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
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// InvestorsUsdtABI is the input ABI used to generate the binding from.
const InvestorsUsdtABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"remaining\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"initialSupply\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"holder\",\"type\":\"address\"}],\"name\":\"initCoin\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"}],\"stateMutability\":\"payable\",\"type\":\"function\"}]"

// InvestorsUsdt is an auto generated Go binding around an Ethereum contract.
type InvestorsUsdt struct {
	InvestorsUsdtCaller     // Read-only binding to the contract
	InvestorsUsdtTransactor // Write-only binding to the contract
	InvestorsUsdtFilterer   // Log filterer for contract events
}

// InvestorsUsdtCaller is an auto generated read-only Go binding around an Ethereum contract.
type InvestorsUsdtCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// InvestorsUsdtTransactor is an auto generated write-only Go binding around an Ethereum contract.
type InvestorsUsdtTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// InvestorsUsdtFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type InvestorsUsdtFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// InvestorsUsdtSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type InvestorsUsdtSession struct {
	Contract     *InvestorsUsdt    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// InvestorsUsdtCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type InvestorsUsdtCallerSession struct {
	Contract *InvestorsUsdtCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// InvestorsUsdtTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type InvestorsUsdtTransactorSession struct {
	Contract     *InvestorsUsdtTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// InvestorsUsdtRaw is an auto generated low-level Go binding around an Ethereum contract.
type InvestorsUsdtRaw struct {
	Contract *InvestorsUsdt // Generic contract binding to access the raw methods on
}

// InvestorsUsdtCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type InvestorsUsdtCallerRaw struct {
	Contract *InvestorsUsdtCaller // Generic read-only contract binding to access the raw methods on
}

// InvestorsUsdtTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type InvestorsUsdtTransactorRaw struct {
	Contract *InvestorsUsdtTransactor // Generic write-only contract binding to access the raw methods on
}

// NewInvestorsUsdt creates a new instance of InvestorsUsdt, bound to a specific deployed contract.
func NewInvestorsUsdt(address common.Address, backend bind.ContractBackend) (*InvestorsUsdt, error) {
	contract, err := bindInvestorsUsdt(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &InvestorsUsdt{InvestorsUsdtCaller: InvestorsUsdtCaller{contract: contract}, InvestorsUsdtTransactor: InvestorsUsdtTransactor{contract: contract}, InvestorsUsdtFilterer: InvestorsUsdtFilterer{contract: contract}}, nil
}

// NewInvestorsUsdtCaller creates a new read-only instance of InvestorsUsdt, bound to a specific deployed contract.
func NewInvestorsUsdtCaller(address common.Address, caller bind.ContractCaller) (*InvestorsUsdtCaller, error) {
	contract, err := bindInvestorsUsdt(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &InvestorsUsdtCaller{contract: contract}, nil
}

// NewInvestorsUsdtTransactor creates a new write-only instance of InvestorsUsdt, bound to a specific deployed contract.
func NewInvestorsUsdtTransactor(address common.Address, transactor bind.ContractTransactor) (*InvestorsUsdtTransactor, error) {
	contract, err := bindInvestorsUsdt(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &InvestorsUsdtTransactor{contract: contract}, nil
}

// NewInvestorsUsdtFilterer creates a new log filterer instance of InvestorsUsdt, bound to a specific deployed contract.
func NewInvestorsUsdtFilterer(address common.Address, filterer bind.ContractFilterer) (*InvestorsUsdtFilterer, error) {
	contract, err := bindInvestorsUsdt(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &InvestorsUsdtFilterer{contract: contract}, nil
}

// bindInvestorsUsdt binds a generic wrapper to an already deployed contract.
func bindInvestorsUsdt(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(InvestorsUsdtABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_InvestorsUsdt *InvestorsUsdtRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _InvestorsUsdt.Contract.InvestorsUsdtCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_InvestorsUsdt *InvestorsUsdtRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _InvestorsUsdt.Contract.InvestorsUsdtTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_InvestorsUsdt *InvestorsUsdtRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _InvestorsUsdt.Contract.InvestorsUsdtTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_InvestorsUsdt *InvestorsUsdtCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _InvestorsUsdt.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_InvestorsUsdt *InvestorsUsdtTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _InvestorsUsdt.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_InvestorsUsdt *InvestorsUsdtTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _InvestorsUsdt.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address _owner, address _spender) view returns(uint256 remaining)
func (_InvestorsUsdt *InvestorsUsdtCaller) Allowance(opts *bind.CallOpts, _owner common.Address, _spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _InvestorsUsdt.contract.Call(opts, &out, "allowance", _owner, _spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address _owner, address _spender) view returns(uint256 remaining)
func (_InvestorsUsdt *InvestorsUsdtSession) Allowance(_owner common.Address, _spender common.Address) (*big.Int, error) {
	return _InvestorsUsdt.Contract.Allowance(&_InvestorsUsdt.CallOpts, _owner, _spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address _owner, address _spender) view returns(uint256 remaining)
func (_InvestorsUsdt *InvestorsUsdtCallerSession) Allowance(_owner common.Address, _spender common.Address) (*big.Int, error) {
	return _InvestorsUsdt.Contract.Allowance(&_InvestorsUsdt.CallOpts, _owner, _spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address _owner) view returns(uint256 balance)
func (_InvestorsUsdt *InvestorsUsdtCaller) BalanceOf(opts *bind.CallOpts, _owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _InvestorsUsdt.contract.Call(opts, &out, "balanceOf", _owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address _owner) view returns(uint256 balance)
func (_InvestorsUsdt *InvestorsUsdtSession) BalanceOf(_owner common.Address) (*big.Int, error) {
	return _InvestorsUsdt.Contract.BalanceOf(&_InvestorsUsdt.CallOpts, _owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address _owner) view returns(uint256 balance)
func (_InvestorsUsdt *InvestorsUsdtCallerSession) BalanceOf(_owner common.Address) (*big.Int, error) {
	return _InvestorsUsdt.Contract.BalanceOf(&_InvestorsUsdt.CallOpts, _owner)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_InvestorsUsdt *InvestorsUsdtCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _InvestorsUsdt.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_InvestorsUsdt *InvestorsUsdtSession) Decimals() (uint8, error) {
	return _InvestorsUsdt.Contract.Decimals(&_InvestorsUsdt.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_InvestorsUsdt *InvestorsUsdtCallerSession) Decimals() (uint8, error) {
	return _InvestorsUsdt.Contract.Decimals(&_InvestorsUsdt.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_InvestorsUsdt *InvestorsUsdtCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _InvestorsUsdt.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_InvestorsUsdt *InvestorsUsdtSession) Name() (string, error) {
	return _InvestorsUsdt.Contract.Name(&_InvestorsUsdt.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_InvestorsUsdt *InvestorsUsdtCallerSession) Name() (string, error) {
	return _InvestorsUsdt.Contract.Name(&_InvestorsUsdt.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_InvestorsUsdt *InvestorsUsdtCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _InvestorsUsdt.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_InvestorsUsdt *InvestorsUsdtSession) Symbol() (string, error) {
	return _InvestorsUsdt.Contract.Symbol(&_InvestorsUsdt.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_InvestorsUsdt *InvestorsUsdtCallerSession) Symbol() (string, error) {
	return _InvestorsUsdt.Contract.Symbol(&_InvestorsUsdt.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_InvestorsUsdt *InvestorsUsdtCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _InvestorsUsdt.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_InvestorsUsdt *InvestorsUsdtSession) TotalSupply() (*big.Int, error) {
	return _InvestorsUsdt.Contract.TotalSupply(&_InvestorsUsdt.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_InvestorsUsdt *InvestorsUsdtCallerSession) TotalSupply() (*big.Int, error) {
	return _InvestorsUsdt.Contract.TotalSupply(&_InvestorsUsdt.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address _spender, uint256 _value) returns(bool success)
func (_InvestorsUsdt *InvestorsUsdtTransactor) Approve(opts *bind.TransactOpts, _spender common.Address, _value *big.Int) (*types.Transaction, error) {
	return _InvestorsUsdt.contract.Transact(opts, "approve", _spender, _value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address _spender, uint256 _value) returns(bool success)
func (_InvestorsUsdt *InvestorsUsdtSession) Approve(_spender common.Address, _value *big.Int) (*types.Transaction, error) {
	return _InvestorsUsdt.Contract.Approve(&_InvestorsUsdt.TransactOpts, _spender, _value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address _spender, uint256 _value) returns(bool success)
func (_InvestorsUsdt *InvestorsUsdtTransactorSession) Approve(_spender common.Address, _value *big.Int) (*types.Transaction, error) {
	return _InvestorsUsdt.Contract.Approve(&_InvestorsUsdt.TransactOpts, _spender, _value)
}

// InitCoin is a paid mutator transaction binding the contract method 0xf50c3f4c.
//
// Solidity: function initCoin(uint256 initialSupply, address holder) returns()
func (_InvestorsUsdt *InvestorsUsdtTransactor) InitCoin(opts *bind.TransactOpts, initialSupply *big.Int, holder common.Address) (*types.Transaction, error) {
	return _InvestorsUsdt.contract.Transact(opts, "initCoin", initialSupply, holder)
}

// InitCoin is a paid mutator transaction binding the contract method 0xf50c3f4c.
//
// Solidity: function initCoin(uint256 initialSupply, address holder) returns()
func (_InvestorsUsdt *InvestorsUsdtSession) InitCoin(initialSupply *big.Int, holder common.Address) (*types.Transaction, error) {
	return _InvestorsUsdt.Contract.InitCoin(&_InvestorsUsdt.TransactOpts, initialSupply, holder)
}

// InitCoin is a paid mutator transaction binding the contract method 0xf50c3f4c.
//
// Solidity: function initCoin(uint256 initialSupply, address holder) returns()
func (_InvestorsUsdt *InvestorsUsdtTransactorSession) InitCoin(initialSupply *big.Int, holder common.Address) (*types.Transaction, error) {
	return _InvestorsUsdt.Contract.InitCoin(&_InvestorsUsdt.TransactOpts, initialSupply, holder)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address _to, uint256 _value) payable returns(bool success)
func (_InvestorsUsdt *InvestorsUsdtTransactor) Transfer(opts *bind.TransactOpts, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _InvestorsUsdt.contract.Transact(opts, "transfer", _to, _value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address _to, uint256 _value) payable returns(bool success)
func (_InvestorsUsdt *InvestorsUsdtSession) Transfer(_to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _InvestorsUsdt.Contract.Transfer(&_InvestorsUsdt.TransactOpts, _to, _value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address _to, uint256 _value) payable returns(bool success)
func (_InvestorsUsdt *InvestorsUsdtTransactorSession) Transfer(_to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _InvestorsUsdt.Contract.Transfer(&_InvestorsUsdt.TransactOpts, _to, _value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address _from, address _to, uint256 _value) payable returns(bool success)
func (_InvestorsUsdt *InvestorsUsdtTransactor) TransferFrom(opts *bind.TransactOpts, _from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _InvestorsUsdt.contract.Transact(opts, "transferFrom", _from, _to, _value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address _from, address _to, uint256 _value) payable returns(bool success)
func (_InvestorsUsdt *InvestorsUsdtSession) TransferFrom(_from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _InvestorsUsdt.Contract.TransferFrom(&_InvestorsUsdt.TransactOpts, _from, _to, _value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address _from, address _to, uint256 _value) payable returns(bool success)
func (_InvestorsUsdt *InvestorsUsdtTransactorSession) TransferFrom(_from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _InvestorsUsdt.Contract.TransferFrom(&_InvestorsUsdt.TransactOpts, _from, _to, _value)
}

// InvestorsUsdtApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the InvestorsUsdt contract.
type InvestorsUsdtApprovalIterator struct {
	Event *InvestorsUsdtApproval // Event containing the contract specifics and raw log

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
func (it *InvestorsUsdtApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(InvestorsUsdtApproval)
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
		it.Event = new(InvestorsUsdtApproval)
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
func (it *InvestorsUsdtApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *InvestorsUsdtApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// InvestorsUsdtApproval represents a Approval event raised by the InvestorsUsdt contract.
type InvestorsUsdtApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address owner, address spender, uint256 value)
func (_InvestorsUsdt *InvestorsUsdtFilterer) FilterApproval(opts *bind.FilterOpts) (*InvestorsUsdtApprovalIterator, error) {

	logs, sub, err := _InvestorsUsdt.contract.FilterLogs(opts, "Approval")
	if err != nil {
		return nil, err
	}
	return &InvestorsUsdtApprovalIterator{contract: _InvestorsUsdt.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address owner, address spender, uint256 value)
func (_InvestorsUsdt *InvestorsUsdtFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *InvestorsUsdtApproval) (event.Subscription, error) {

	logs, sub, err := _InvestorsUsdt.contract.WatchLogs(opts, "Approval")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(InvestorsUsdtApproval)
				if err := _InvestorsUsdt.contract.UnpackLog(event, "Approval", log); err != nil {
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

// ParseApproval is a log parse operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address owner, address spender, uint256 value)
func (_InvestorsUsdt *InvestorsUsdtFilterer) ParseApproval(log types.Log) (*InvestorsUsdtApproval, error) {
	event := new(InvestorsUsdtApproval)
	if err := _InvestorsUsdt.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	return event, nil
}

// InvestorsUsdtTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the InvestorsUsdt contract.
type InvestorsUsdtTransferIterator struct {
	Event *InvestorsUsdtTransfer // Event containing the contract specifics and raw log

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
func (it *InvestorsUsdtTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(InvestorsUsdtTransfer)
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
		it.Event = new(InvestorsUsdtTransfer)
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
func (it *InvestorsUsdtTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *InvestorsUsdtTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// InvestorsUsdtTransfer represents a Transfer event raised by the InvestorsUsdt contract.
type InvestorsUsdtTransfer struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address owner, address spender, uint256 value)
func (_InvestorsUsdt *InvestorsUsdtFilterer) FilterTransfer(opts *bind.FilterOpts) (*InvestorsUsdtTransferIterator, error) {

	logs, sub, err := _InvestorsUsdt.contract.FilterLogs(opts, "Transfer")
	if err != nil {
		return nil, err
	}
	return &InvestorsUsdtTransferIterator{contract: _InvestorsUsdt.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address owner, address spender, uint256 value)
func (_InvestorsUsdt *InvestorsUsdtFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *InvestorsUsdtTransfer) (event.Subscription, error) {

	logs, sub, err := _InvestorsUsdt.contract.WatchLogs(opts, "Transfer")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(InvestorsUsdtTransfer)
				if err := _InvestorsUsdt.contract.UnpackLog(event, "Transfer", log); err != nil {
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

// ParseTransfer is a log parse operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address owner, address spender, uint256 value)
func (_InvestorsUsdt *InvestorsUsdtFilterer) ParseTransfer(log types.Log) (*InvestorsUsdtTransfer, error) {
	event := new(InvestorsUsdtTransfer)
	if err := _InvestorsUsdt.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	return event, nil
}
