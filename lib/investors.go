// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package investors

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

// InvestorsABI is the input ABI used to generate the binding from.
const InvestorsABI = "[{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_startover\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"_institutionflag\",\"type\":\"bool\"},{\"internalType\":\"address\",\"name\":\"_institution\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"firstBlock\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"hsfToken\",\"outputs\":[{\"internalType\":\"contractHillStoneFinance\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"hsfaddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tcaddress\",\"type\":\"address\"}],\"name\":\"initToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"institution\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"institutionflag\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"maneger\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"fromAdress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_betClass\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"betAmount\",\"type\":\"uint256\"}],\"name\":\"placeBet\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"toAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_betClass\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"removeAmount\",\"type\":\"uint256\"}],\"name\":\"removeBet\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"bounces\",\"type\":\"uint256\"}],\"name\":\"shareOut\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"startover\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"usToken\",\"outputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]"

// Investors is an auto generated Go binding around an Ethereum contract.
type Investors struct {
	InvestorsCaller     // Read-only binding to the contract
	InvestorsTransactor // Write-only binding to the contract
	InvestorsFilterer   // Log filterer for contract events
}

// InvestorsCaller is an auto generated read-only Go binding around an Ethereum contract.
type InvestorsCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// InvestorsTransactor is an auto generated write-only Go binding around an Ethereum contract.
type InvestorsTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// InvestorsFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type InvestorsFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// InvestorsSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type InvestorsSession struct {
	Contract     *Investors        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// InvestorsCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type InvestorsCallerSession struct {
	Contract *InvestorsCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// InvestorsTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type InvestorsTransactorSession struct {
	Contract     *InvestorsTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// InvestorsRaw is an auto generated low-level Go binding around an Ethereum contract.
type InvestorsRaw struct {
	Contract *Investors // Generic contract binding to access the raw methods on
}

// InvestorsCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type InvestorsCallerRaw struct {
	Contract *InvestorsCaller // Generic read-only contract binding to access the raw methods on
}

// InvestorsTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type InvestorsTransactorRaw struct {
	Contract *InvestorsTransactor // Generic write-only contract binding to access the raw methods on
}

// NewInvestors creates a new instance of Investors, bound to a specific deployed contract.
func NewInvestors(address common.Address, backend bind.ContractBackend) (*Investors, error) {
	contract, err := bindInvestors(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Investors{InvestorsCaller: InvestorsCaller{contract: contract}, InvestorsTransactor: InvestorsTransactor{contract: contract}, InvestorsFilterer: InvestorsFilterer{contract: contract}}, nil
}

// NewInvestorsCaller creates a new read-only instance of Investors, bound to a specific deployed contract.
func NewInvestorsCaller(address common.Address, caller bind.ContractCaller) (*InvestorsCaller, error) {
	contract, err := bindInvestors(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &InvestorsCaller{contract: contract}, nil
}

// NewInvestorsTransactor creates a new write-only instance of Investors, bound to a specific deployed contract.
func NewInvestorsTransactor(address common.Address, transactor bind.ContractTransactor) (*InvestorsTransactor, error) {
	contract, err := bindInvestors(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &InvestorsTransactor{contract: contract}, nil
}

// NewInvestorsFilterer creates a new log filterer instance of Investors, bound to a specific deployed contract.
func NewInvestorsFilterer(address common.Address, filterer bind.ContractFilterer) (*InvestorsFilterer, error) {
	contract, err := bindInvestors(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &InvestorsFilterer{contract: contract}, nil
}

// bindInvestors binds a generic wrapper to an already deployed contract.
func bindInvestors(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(InvestorsABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Investors *InvestorsRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Investors.Contract.InvestorsCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Investors *InvestorsRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Investors.Contract.InvestorsTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Investors *InvestorsRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Investors.Contract.InvestorsTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Investors *InvestorsCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Investors.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Investors *InvestorsTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Investors.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Investors *InvestorsTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Investors.Contract.contract.Transact(opts, method, params...)
}

// FirstBlock is a free data retrieval call binding the contract method 0x231b0268.
//
// Solidity: function firstBlock() view returns(uint256)
func (_Investors *InvestorsCaller) FirstBlock(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Investors.contract.Call(opts, &out, "firstBlock")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FirstBlock is a free data retrieval call binding the contract method 0x231b0268.
//
// Solidity: function firstBlock() view returns(uint256)
func (_Investors *InvestorsSession) FirstBlock() (*big.Int, error) {
	return _Investors.Contract.FirstBlock(&_Investors.CallOpts)
}

// FirstBlock is a free data retrieval call binding the contract method 0x231b0268.
//
// Solidity: function firstBlock() view returns(uint256)
func (_Investors *InvestorsCallerSession) FirstBlock() (*big.Int, error) {
	return _Investors.Contract.FirstBlock(&_Investors.CallOpts)
}

// HsfToken is a free data retrieval call binding the contract method 0xf1ce62b7.
//
// Solidity: function hsfToken() view returns(address)
func (_Investors *InvestorsCaller) HsfToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Investors.contract.Call(opts, &out, "hsfToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// HsfToken is a free data retrieval call binding the contract method 0xf1ce62b7.
//
// Solidity: function hsfToken() view returns(address)
func (_Investors *InvestorsSession) HsfToken() (common.Address, error) {
	return _Investors.Contract.HsfToken(&_Investors.CallOpts)
}

// HsfToken is a free data retrieval call binding the contract method 0xf1ce62b7.
//
// Solidity: function hsfToken() view returns(address)
func (_Investors *InvestorsCallerSession) HsfToken() (common.Address, error) {
	return _Investors.Contract.HsfToken(&_Investors.CallOpts)
}

// Institution is a free data retrieval call binding the contract method 0xd1384eb6.
//
// Solidity: function institution() view returns(address)
func (_Investors *InvestorsCaller) Institution(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Investors.contract.Call(opts, &out, "institution")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Institution is a free data retrieval call binding the contract method 0xd1384eb6.
//
// Solidity: function institution() view returns(address)
func (_Investors *InvestorsSession) Institution() (common.Address, error) {
	return _Investors.Contract.Institution(&_Investors.CallOpts)
}

// Institution is a free data retrieval call binding the contract method 0xd1384eb6.
//
// Solidity: function institution() view returns(address)
func (_Investors *InvestorsCallerSession) Institution() (common.Address, error) {
	return _Investors.Contract.Institution(&_Investors.CallOpts)
}

// Institutionflag is a free data retrieval call binding the contract method 0x7929ce1e.
//
// Solidity: function institutionflag() view returns(bool)
func (_Investors *InvestorsCaller) Institutionflag(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Investors.contract.Call(opts, &out, "institutionflag")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Institutionflag is a free data retrieval call binding the contract method 0x7929ce1e.
//
// Solidity: function institutionflag() view returns(bool)
func (_Investors *InvestorsSession) Institutionflag() (bool, error) {
	return _Investors.Contract.Institutionflag(&_Investors.CallOpts)
}

// Institutionflag is a free data retrieval call binding the contract method 0x7929ce1e.
//
// Solidity: function institutionflag() view returns(bool)
func (_Investors *InvestorsCallerSession) Institutionflag() (bool, error) {
	return _Investors.Contract.Institutionflag(&_Investors.CallOpts)
}

// Maneger is a free data retrieval call binding the contract method 0xd5091834.
//
// Solidity: function maneger() view returns(address)
func (_Investors *InvestorsCaller) Maneger(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Investors.contract.Call(opts, &out, "maneger")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Maneger is a free data retrieval call binding the contract method 0xd5091834.
//
// Solidity: function maneger() view returns(address)
func (_Investors *InvestorsSession) Maneger() (common.Address, error) {
	return _Investors.Contract.Maneger(&_Investors.CallOpts)
}

// Maneger is a free data retrieval call binding the contract method 0xd5091834.
//
// Solidity: function maneger() view returns(address)
func (_Investors *InvestorsCallerSession) Maneger() (common.Address, error) {
	return _Investors.Contract.Maneger(&_Investors.CallOpts)
}

// Startover is a free data retrieval call binding the contract method 0x38fe6a96.
//
// Solidity: function startover() view returns(uint256)
func (_Investors *InvestorsCaller) Startover(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Investors.contract.Call(opts, &out, "startover")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Startover is a free data retrieval call binding the contract method 0x38fe6a96.
//
// Solidity: function startover() view returns(uint256)
func (_Investors *InvestorsSession) Startover() (*big.Int, error) {
	return _Investors.Contract.Startover(&_Investors.CallOpts)
}

// Startover is a free data retrieval call binding the contract method 0x38fe6a96.
//
// Solidity: function startover() view returns(uint256)
func (_Investors *InvestorsCallerSession) Startover() (*big.Int, error) {
	return _Investors.Contract.Startover(&_Investors.CallOpts)
}

// UsToken is a free data retrieval call binding the contract method 0x23191418.
//
// Solidity: function usToken() view returns(address)
func (_Investors *InvestorsCaller) UsToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Investors.contract.Call(opts, &out, "usToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// UsToken is a free data retrieval call binding the contract method 0x23191418.
//
// Solidity: function usToken() view returns(address)
func (_Investors *InvestorsSession) UsToken() (common.Address, error) {
	return _Investors.Contract.UsToken(&_Investors.CallOpts)
}

// UsToken is a free data retrieval call binding the contract method 0x23191418.
//
// Solidity: function usToken() view returns(address)
func (_Investors *InvestorsCallerSession) UsToken() (common.Address, error) {
	return _Investors.Contract.UsToken(&_Investors.CallOpts)
}

// InitToken is a paid mutator transaction binding the contract method 0xba2b6bad.
//
// Solidity: function initToken(address hsfaddress, address tcaddress) returns()
func (_Investors *InvestorsTransactor) InitToken(opts *bind.TransactOpts, hsfaddress common.Address, tcaddress common.Address) (*types.Transaction, error) {
	return _Investors.contract.Transact(opts, "initToken", hsfaddress, tcaddress)
}

// InitToken is a paid mutator transaction binding the contract method 0xba2b6bad.
//
// Solidity: function initToken(address hsfaddress, address tcaddress) returns()
func (_Investors *InvestorsSession) InitToken(hsfaddress common.Address, tcaddress common.Address) (*types.Transaction, error) {
	return _Investors.Contract.InitToken(&_Investors.TransactOpts, hsfaddress, tcaddress)
}

// InitToken is a paid mutator transaction binding the contract method 0xba2b6bad.
//
// Solidity: function initToken(address hsfaddress, address tcaddress) returns()
func (_Investors *InvestorsTransactorSession) InitToken(hsfaddress common.Address, tcaddress common.Address) (*types.Transaction, error) {
	return _Investors.Contract.InitToken(&_Investors.TransactOpts, hsfaddress, tcaddress)
}

// PlaceBet is a paid mutator transaction binding the contract method 0xf231a4f6.
//
// Solidity: function placeBet(address fromAdress, uint256 _betClass, uint256 betAmount) payable returns(bool success)
func (_Investors *InvestorsTransactor) PlaceBet(opts *bind.TransactOpts, fromAdress common.Address, _betClass *big.Int, betAmount *big.Int) (*types.Transaction, error) {
	return _Investors.contract.Transact(opts, "placeBet", fromAdress, _betClass, betAmount)
}

// PlaceBet is a paid mutator transaction binding the contract method 0xf231a4f6.
//
// Solidity: function placeBet(address fromAdress, uint256 _betClass, uint256 betAmount) payable returns(bool success)
func (_Investors *InvestorsSession) PlaceBet(fromAdress common.Address, _betClass *big.Int, betAmount *big.Int) (*types.Transaction, error) {
	return _Investors.Contract.PlaceBet(&_Investors.TransactOpts, fromAdress, _betClass, betAmount)
}

// PlaceBet is a paid mutator transaction binding the contract method 0xf231a4f6.
//
// Solidity: function placeBet(address fromAdress, uint256 _betClass, uint256 betAmount) payable returns(bool success)
func (_Investors *InvestorsTransactorSession) PlaceBet(fromAdress common.Address, _betClass *big.Int, betAmount *big.Int) (*types.Transaction, error) {
	return _Investors.Contract.PlaceBet(&_Investors.TransactOpts, fromAdress, _betClass, betAmount)
}

// RemoveBet is a paid mutator transaction binding the contract method 0x8a9faad2.
//
// Solidity: function removeBet(address toAddress, uint256 _betClass, uint256 removeAmount) payable returns(bool success)
func (_Investors *InvestorsTransactor) RemoveBet(opts *bind.TransactOpts, toAddress common.Address, _betClass *big.Int, removeAmount *big.Int) (*types.Transaction, error) {
	return _Investors.contract.Transact(opts, "removeBet", toAddress, _betClass, removeAmount)
}

// RemoveBet is a paid mutator transaction binding the contract method 0x8a9faad2.
//
// Solidity: function removeBet(address toAddress, uint256 _betClass, uint256 removeAmount) payable returns(bool success)
func (_Investors *InvestorsSession) RemoveBet(toAddress common.Address, _betClass *big.Int, removeAmount *big.Int) (*types.Transaction, error) {
	return _Investors.Contract.RemoveBet(&_Investors.TransactOpts, toAddress, _betClass, removeAmount)
}

// RemoveBet is a paid mutator transaction binding the contract method 0x8a9faad2.
//
// Solidity: function removeBet(address toAddress, uint256 _betClass, uint256 removeAmount) payable returns(bool success)
func (_Investors *InvestorsTransactorSession) RemoveBet(toAddress common.Address, _betClass *big.Int, removeAmount *big.Int) (*types.Transaction, error) {
	return _Investors.Contract.RemoveBet(&_Investors.TransactOpts, toAddress, _betClass, removeAmount)
}

// ShareOut is a paid mutator transaction binding the contract method 0x35400d66.
//
// Solidity: function shareOut(uint256 bounces) payable returns()
func (_Investors *InvestorsTransactor) ShareOut(opts *bind.TransactOpts, bounces *big.Int) (*types.Transaction, error) {
	return _Investors.contract.Transact(opts, "shareOut", bounces)
}

// ShareOut is a paid mutator transaction binding the contract method 0x35400d66.
//
// Solidity: function shareOut(uint256 bounces) payable returns()
func (_Investors *InvestorsSession) ShareOut(bounces *big.Int) (*types.Transaction, error) {
	return _Investors.Contract.ShareOut(&_Investors.TransactOpts, bounces)
}

// ShareOut is a paid mutator transaction binding the contract method 0x35400d66.
//
// Solidity: function shareOut(uint256 bounces) payable returns()
func (_Investors *InvestorsTransactorSession) ShareOut(bounces *big.Int) (*types.Transaction, error) {
	return _Investors.Contract.ShareOut(&_Investors.TransactOpts, bounces)
}
