// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package usdt

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

// UsdtABI is the input ABI used to generate the binding from.
const UsdtABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"remaining\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"initialSupply\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"holder\",\"type\":\"address\"}],\"name\":\"initCoin\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"}],\"stateMutability\":\"payable\",\"type\":\"function\"}]"

// UsdtBin is the compiled bytecode used for deploying new contracts.
var UsdtBin = "0x60806040526012600260006101000a81548160ff021916908360ff16021790555034801561002c57600080fd5b506112a88061003c6000396000f3fe6080604052600436106100915760003560e01c806370a082311161005957806370a082311461018457806395d89b41146101c1578063a9059cbb146101ec578063dd62ed3e1461021c578063f50c3f4c1461025957610091565b806306fdde0314610096578063095ea7b3146100c157806318160ddd146100fe57806323b872dd14610129578063313ce56714610159575b600080fd5b3480156100a257600080fd5b506100ab610282565b6040516100b89190610e45565b60405180910390f35b3480156100cd57600080fd5b506100e860048036038101906100e39190610d06565b610314565b6040516100f59190610e2a565b60405180910390f35b34801561010a57600080fd5b5061011361046f565b6040516101209190610e67565b60405180910390f35b610143600480360381019061013e9190610cb7565b610479565b6040516101509190610e2a565b60405180910390f35b34801561016557600080fd5b5061016e6106d4565b60405161017b9190610e82565b60405180910390f35b34801561019057600080fd5b506101ab60048036038101906101a69190610c52565b6106eb565b6040516101b89190610e67565b60405180910390f35b3480156101cd57600080fd5b506101d6610734565b6040516101e39190610e45565b60405180910390f35b61020660048036038101906102019190610d06565b6107c6565b6040516102139190610e2a565b60405180910390f35b34801561022857600080fd5b50610243600480360381019061023e9190610c7b565b610999565b6040516102509190610e67565b60405180910390f35b34801561026557600080fd5b50610280600480360381019061027b9190610d42565b610a20565b005b60606000805461029190611196565b80601f01602080910402602001604051908101604052809291908181526020018280546102bd90611196565b801561030a5780601f106102df5761010080835404028352916020019161030a565b820191906000526020600020905b8154815290600101906020018083116102ed57829003601f168201915b5050505050905090565b6000808214806103a057506000600660003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054145b6103a957600080fd5b81600660003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020819055507f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b92533848460405161045d93929190610df3565b60405180910390a16001905092915050565b6000600354905090565b600081600560008673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000205410158015610546575081600660008673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000205410155b61054f57600080fd5b81600560008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600082825461059e9190610eb9565b9250508190555081600560008673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008282546105f491906110da565b9250508190555081600660008673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600082825461068791906110da565b925050819055507fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef8484846040516106c193929190610df3565b60405180910390a1600190509392505050565b6000600260009054906101000a900460ff16905090565b6000600560008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020549050919050565b60606001805461074390611196565b80601f016020809104026020016040519081016040528092919081815260200182805461076f90611196565b80156107bc5780601f10610791576101008083540402835291602001916107bc565b820191906000526020600020905b81548152906001019060200180831161079f57829003601f168201915b5050505050905090565b600081600560003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020541015801561089f5750600560008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000205482600560008673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000205461089d9190610eb9565b115b6108a857600080fd5b81600560003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008282546108f791906110da565b9250508190555081600560008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600082825461094d9190610eb9565b925050819055507fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef33848460405161098793929190610df3565b60405180910390a16001905092915050565b6000600660008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054905092915050565b600260009054906101000a900460ff1660ff16600a610a3f9190610f62565b82610a4a9190611080565b600381905550600354600560008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206000828254610aa19190610eb9565b925050819055506040518060400160405280600481526020017f555344540000000000000000000000000000000000000000000000000000000081525060009080519060200190610af3929190610b85565b506040518060400160405280600481526020017f757364740000000000000000000000000000000000000000000000000000000081525060019080519060200190610b3f929190610b85565b5080600460006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055505050565b828054610b9190611196565b90600052602060002090601f016020900481019282610bb35760008555610bfa565b82601f10610bcc57805160ff1916838001178555610bfa565b82800160010185558215610bfa579182015b82811115610bf9578251825591602001919060010190610bde565b5b509050610c079190610c0b565b5090565b5b80821115610c24576000816000905550600101610c0c565b5090565b600081359050610c3781611244565b92915050565b600081359050610c4c8161125b565b92915050565b600060208284031215610c6457600080fd5b6000610c7284828501610c28565b91505092915050565b60008060408385031215610c8e57600080fd5b6000610c9c85828601610c28565b9250506020610cad85828601610c28565b9150509250929050565b600080600060608486031215610ccc57600080fd5b6000610cda86828701610c28565b9350506020610ceb86828701610c28565b9250506040610cfc86828701610c3d565b9150509250925092565b60008060408385031215610d1957600080fd5b6000610d2785828601610c28565b9250506020610d3885828601610c3d565b9150509250929050565b60008060408385031215610d5557600080fd5b6000610d6385828601610c3d565b9250506020610d7485828601610c28565b9150509250929050565b610d878161110e565b82525050565b610d9681611120565b82525050565b6000610da782610e9d565b610db18185610ea8565b9350610dc1818560208601611163565b610dca81611226565b840191505092915050565b610dde8161114c565b82525050565b610ded81611156565b82525050565b6000606082019050610e086000830186610d7e565b610e156020830185610d7e565b610e226040830184610dd5565b949350505050565b6000602082019050610e3f6000830184610d8d565b92915050565b60006020820190508181036000830152610e5f8184610d9c565b905092915050565b6000602082019050610e7c6000830184610dd5565b92915050565b6000602082019050610e976000830184610de4565b92915050565b600081519050919050565b600082825260208201905092915050565b6000610ec48261114c565b9150610ecf8361114c565b9250827fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff03821115610f0457610f036111c8565b5b828201905092915050565b6000808291508390505b6001851115610f5957808604811115610f3557610f346111c8565b5b6001851615610f445780820291505b8081029050610f5285611237565b9450610f19565b94509492505050565b6000610f6d8261114c565b9150610f788361114c565b9250610fa57fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8484610fad565b905092915050565b600082610fbd5760019050611079565b81610fcb5760009050611079565b8160018114610fe15760028114610feb5761101a565b6001915050611079565b60ff841115610ffd57610ffc6111c8565b5b8360020a915084821115611014576110136111c8565b5b50611079565b5060208310610133831016604e8410600b841016171561104f5782820a90508381111561104a576110496111c8565b5b611079565b61105c8484846001610f0f565b92509050818404811115611073576110726111c8565b5b81810290505b9392505050565b600061108b8261114c565b91506110968361114c565b9250817fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff04831182151516156110cf576110ce6111c8565b5b828202905092915050565b60006110e58261114c565b91506110f08361114c565b925082821015611103576111026111c8565b5b828203905092915050565b60006111198261112c565b9050919050565b60008115159050919050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000819050919050565b600060ff82169050919050565b60005b83811015611181578082015181840152602081019050611166565b83811115611190576000848401525b50505050565b600060028204905060018216806111ae57607f821691505b602082108114156111c2576111c16111f7565b5b50919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b6000601f19601f8301169050919050565b60008160011c9050919050565b61124d8161110e565b811461125857600080fd5b50565b6112648161114c565b811461126f57600080fd5b5056fea2646970667358221220e6e9c5545798b207cb8c6de5b4ff9d446b0f8dd3c2c6e52059234d9b5da9701d64736f6c63430008030033"

// DeployUsdt deploys a new Ethereum contract, binding an instance of Usdt to it.
func DeployUsdt(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Usdt, error) {
	parsed, err := abi.JSON(strings.NewReader(UsdtABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(UsdtBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Usdt{UsdtCaller: UsdtCaller{contract: contract}, UsdtTransactor: UsdtTransactor{contract: contract}, UsdtFilterer: UsdtFilterer{contract: contract}}, nil
}

// Usdt is an auto generated Go binding around an Ethereum contract.
type Usdt struct {
	UsdtCaller     // Read-only binding to the contract
	UsdtTransactor // Write-only binding to the contract
	UsdtFilterer   // Log filterer for contract events
}

// UsdtCaller is an auto generated read-only Go binding around an Ethereum contract.
type UsdtCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UsdtTransactor is an auto generated write-only Go binding around an Ethereum contract.
type UsdtTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UsdtFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type UsdtFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UsdtSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type UsdtSession struct {
	Contract     *Usdt             // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// UsdtCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type UsdtCallerSession struct {
	Contract *UsdtCaller   // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// UsdtTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type UsdtTransactorSession struct {
	Contract     *UsdtTransactor   // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// UsdtRaw is an auto generated low-level Go binding around an Ethereum contract.
type UsdtRaw struct {
	Contract *Usdt // Generic contract binding to access the raw methods on
}

// UsdtCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type UsdtCallerRaw struct {
	Contract *UsdtCaller // Generic read-only contract binding to access the raw methods on
}

// UsdtTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type UsdtTransactorRaw struct {
	Contract *UsdtTransactor // Generic write-only contract binding to access the raw methods on
}

// NewUsdt creates a new instance of Usdt, bound to a specific deployed contract.
func NewUsdt(address common.Address, backend bind.ContractBackend) (*Usdt, error) {
	contract, err := bindUsdt(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Usdt{UsdtCaller: UsdtCaller{contract: contract}, UsdtTransactor: UsdtTransactor{contract: contract}, UsdtFilterer: UsdtFilterer{contract: contract}}, nil
}

// NewUsdtCaller creates a new read-only instance of Usdt, bound to a specific deployed contract.
func NewUsdtCaller(address common.Address, caller bind.ContractCaller) (*UsdtCaller, error) {
	contract, err := bindUsdt(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &UsdtCaller{contract: contract}, nil
}

// NewUsdtTransactor creates a new write-only instance of Usdt, bound to a specific deployed contract.
func NewUsdtTransactor(address common.Address, transactor bind.ContractTransactor) (*UsdtTransactor, error) {
	contract, err := bindUsdt(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &UsdtTransactor{contract: contract}, nil
}

// NewUsdtFilterer creates a new log filterer instance of Usdt, bound to a specific deployed contract.
func NewUsdtFilterer(address common.Address, filterer bind.ContractFilterer) (*UsdtFilterer, error) {
	contract, err := bindUsdt(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &UsdtFilterer{contract: contract}, nil
}

// bindUsdt binds a generic wrapper to an already deployed contract.
func bindUsdt(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(UsdtABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Usdt *UsdtRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Usdt.Contract.UsdtCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Usdt *UsdtRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Usdt.Contract.UsdtTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Usdt *UsdtRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Usdt.Contract.UsdtTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Usdt *UsdtCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Usdt.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Usdt *UsdtTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Usdt.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Usdt *UsdtTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Usdt.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address _owner, address _spender) view returns(uint256 remaining)
func (_Usdt *UsdtCaller) Allowance(opts *bind.CallOpts, _owner common.Address, _spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Usdt.contract.Call(opts, &out, "allowance", _owner, _spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address _owner, address _spender) view returns(uint256 remaining)
func (_Usdt *UsdtSession) Allowance(_owner common.Address, _spender common.Address) (*big.Int, error) {
	return _Usdt.Contract.Allowance(&_Usdt.CallOpts, _owner, _spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address _owner, address _spender) view returns(uint256 remaining)
func (_Usdt *UsdtCallerSession) Allowance(_owner common.Address, _spender common.Address) (*big.Int, error) {
	return _Usdt.Contract.Allowance(&_Usdt.CallOpts, _owner, _spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address _owner) view returns(uint256 balance)
func (_Usdt *UsdtCaller) BalanceOf(opts *bind.CallOpts, _owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Usdt.contract.Call(opts, &out, "balanceOf", _owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address _owner) view returns(uint256 balance)
func (_Usdt *UsdtSession) BalanceOf(_owner common.Address) (*big.Int, error) {
	return _Usdt.Contract.BalanceOf(&_Usdt.CallOpts, _owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address _owner) view returns(uint256 balance)
func (_Usdt *UsdtCallerSession) BalanceOf(_owner common.Address) (*big.Int, error) {
	return _Usdt.Contract.BalanceOf(&_Usdt.CallOpts, _owner)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_Usdt *UsdtCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Usdt.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_Usdt *UsdtSession) Decimals() (uint8, error) {
	return _Usdt.Contract.Decimals(&_Usdt.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_Usdt *UsdtCallerSession) Decimals() (uint8, error) {
	return _Usdt.Contract.Decimals(&_Usdt.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Usdt *UsdtCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Usdt.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Usdt *UsdtSession) Name() (string, error) {
	return _Usdt.Contract.Name(&_Usdt.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Usdt *UsdtCallerSession) Name() (string, error) {
	return _Usdt.Contract.Name(&_Usdt.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Usdt *UsdtCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Usdt.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Usdt *UsdtSession) Symbol() (string, error) {
	return _Usdt.Contract.Symbol(&_Usdt.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Usdt *UsdtCallerSession) Symbol() (string, error) {
	return _Usdt.Contract.Symbol(&_Usdt.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Usdt *UsdtCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Usdt.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Usdt *UsdtSession) TotalSupply() (*big.Int, error) {
	return _Usdt.Contract.TotalSupply(&_Usdt.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Usdt *UsdtCallerSession) TotalSupply() (*big.Int, error) {
	return _Usdt.Contract.TotalSupply(&_Usdt.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address _spender, uint256 _value) returns(bool success)
func (_Usdt *UsdtTransactor) Approve(opts *bind.TransactOpts, _spender common.Address, _value *big.Int) (*types.Transaction, error) {
	return _Usdt.contract.Transact(opts, "approve", _spender, _value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address _spender, uint256 _value) returns(bool success)
func (_Usdt *UsdtSession) Approve(_spender common.Address, _value *big.Int) (*types.Transaction, error) {
	return _Usdt.Contract.Approve(&_Usdt.TransactOpts, _spender, _value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address _spender, uint256 _value) returns(bool success)
func (_Usdt *UsdtTransactorSession) Approve(_spender common.Address, _value *big.Int) (*types.Transaction, error) {
	return _Usdt.Contract.Approve(&_Usdt.TransactOpts, _spender, _value)
}

// InitCoin is a paid mutator transaction binding the contract method 0xf50c3f4c.
//
// Solidity: function initCoin(uint256 initialSupply, address holder) returns()
func (_Usdt *UsdtTransactor) InitCoin(opts *bind.TransactOpts, initialSupply *big.Int, holder common.Address) (*types.Transaction, error) {
	return _Usdt.contract.Transact(opts, "initCoin", initialSupply, holder)
}

// InitCoin is a paid mutator transaction binding the contract method 0xf50c3f4c.
//
// Solidity: function initCoin(uint256 initialSupply, address holder) returns()
func (_Usdt *UsdtSession) InitCoin(initialSupply *big.Int, holder common.Address) (*types.Transaction, error) {
	return _Usdt.Contract.InitCoin(&_Usdt.TransactOpts, initialSupply, holder)
}

// InitCoin is a paid mutator transaction binding the contract method 0xf50c3f4c.
//
// Solidity: function initCoin(uint256 initialSupply, address holder) returns()
func (_Usdt *UsdtTransactorSession) InitCoin(initialSupply *big.Int, holder common.Address) (*types.Transaction, error) {
	return _Usdt.Contract.InitCoin(&_Usdt.TransactOpts, initialSupply, holder)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address _to, uint256 _value) payable returns(bool success)
func (_Usdt *UsdtTransactor) Transfer(opts *bind.TransactOpts, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _Usdt.contract.Transact(opts, "transfer", _to, _value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address _to, uint256 _value) payable returns(bool success)
func (_Usdt *UsdtSession) Transfer(_to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _Usdt.Contract.Transfer(&_Usdt.TransactOpts, _to, _value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address _to, uint256 _value) payable returns(bool success)
func (_Usdt *UsdtTransactorSession) Transfer(_to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _Usdt.Contract.Transfer(&_Usdt.TransactOpts, _to, _value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address _from, address _to, uint256 _value) payable returns(bool success)
func (_Usdt *UsdtTransactor) TransferFrom(opts *bind.TransactOpts, _from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _Usdt.contract.Transact(opts, "transferFrom", _from, _to, _value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address _from, address _to, uint256 _value) payable returns(bool success)
func (_Usdt *UsdtSession) TransferFrom(_from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _Usdt.Contract.TransferFrom(&_Usdt.TransactOpts, _from, _to, _value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address _from, address _to, uint256 _value) payable returns(bool success)
func (_Usdt *UsdtTransactorSession) TransferFrom(_from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _Usdt.Contract.TransferFrom(&_Usdt.TransactOpts, _from, _to, _value)
}

// UsdtApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the Usdt contract.
type UsdtApprovalIterator struct {
	Event *UsdtApproval // Event containing the contract specifics and raw log

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
func (it *UsdtApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(UsdtApproval)
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
		it.Event = new(UsdtApproval)
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
func (it *UsdtApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *UsdtApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// UsdtApproval represents a Approval event raised by the Usdt contract.
type UsdtApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address owner, address spender, uint256 value)
func (_Usdt *UsdtFilterer) FilterApproval(opts *bind.FilterOpts) (*UsdtApprovalIterator, error) {

	logs, sub, err := _Usdt.contract.FilterLogs(opts, "Approval")
	if err != nil {
		return nil, err
	}
	return &UsdtApprovalIterator{contract: _Usdt.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address owner, address spender, uint256 value)
func (_Usdt *UsdtFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *UsdtApproval) (event.Subscription, error) {

	logs, sub, err := _Usdt.contract.WatchLogs(opts, "Approval")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(UsdtApproval)
				if err := _Usdt.contract.UnpackLog(event, "Approval", log); err != nil {
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
func (_Usdt *UsdtFilterer) ParseApproval(log types.Log) (*UsdtApproval, error) {
	event := new(UsdtApproval)
	if err := _Usdt.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// UsdtTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the Usdt contract.
type UsdtTransferIterator struct {
	Event *UsdtTransfer // Event containing the contract specifics and raw log

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
func (it *UsdtTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(UsdtTransfer)
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
		it.Event = new(UsdtTransfer)
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
func (it *UsdtTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *UsdtTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// UsdtTransfer represents a Transfer event raised by the Usdt contract.
type UsdtTransfer struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address owner, address spender, uint256 value)
func (_Usdt *UsdtFilterer) FilterTransfer(opts *bind.FilterOpts) (*UsdtTransferIterator, error) {

	logs, sub, err := _Usdt.contract.FilterLogs(opts, "Transfer")
	if err != nil {
		return nil, err
	}
	return &UsdtTransferIterator{contract: _Usdt.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address owner, address spender, uint256 value)
func (_Usdt *UsdtFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *UsdtTransfer) (event.Subscription, error) {

	logs, sub, err := _Usdt.contract.WatchLogs(opts, "Transfer")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(UsdtTransfer)
				if err := _Usdt.contract.UnpackLog(event, "Transfer", log); err != nil {
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
func (_Usdt *UsdtFilterer) ParseTransfer(log types.Log) (*UsdtTransfer, error) {
	event := new(UsdtTransfer)
	if err := _Usdt.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
