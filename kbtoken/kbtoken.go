// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package kbtoken

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

// KbtokenABI is the input ABI used to generate the binding from.
const KbtokenABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"airdrop\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"remaining\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"burn\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"burnAll\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"burnFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"exchange\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"holder\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"exchange\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"launch\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"reserved\",\"type\":\"address\"}],\"name\":\"initCoin\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"seo\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"}],\"stateMutability\":\"payable\",\"type\":\"function\"}]"

// KbtokenBin is the compiled bytecode used for deploying new contracts.
var KbtokenBin = "0x60806040526002600260006101000a81548160ff021916908360ff16021790555063016e3600600360005090905566038d7ea4c68000600460005090905534801561004a5760006000fd5b50610050565b61286f8061005f6000396000f3fe6080604052600436106100f75760003560e01c806370a082311161008a578063a67c926411610059578063a67c92641461030d578063a9059cbb14610337578063d2f7265a14610367578063dd62ed3e14610385576100f7565b806370a082311461026157806379cc67901461029f57806395d89b41146102c95780639975038c146102f5576100f7565b8063313ce567116100c6578063313ce567146101c35780633884d635146101ef57806342966c681461020d57806370238ecd14610237576100f7565b806306fdde03146100fd578063095ea7b31461012957806318160ddd1461016757806323b872dd14610193576100f7565b60006000fd5b34801561010a5760006000fd5b506101136103c3565b604051610120919061232e565b60405180910390f35b3480156101365760006000fd5b50610151600480360381019061014c91906120d6565b61045d565b60405161015e9190612312565b60405180910390f35b3480156101745760006000fd5b5061017d6105d5565b60405161018a9190612351565b60405180910390f35b6101ad60048036038101906101a89190612084565b6105e7565b6040516101ba9190612312565b60405180910390f35b3480156101d05760006000fd5b506101d9610871565b6040516101e6919061236d565b60405180910390f35b6101f761088d565b6040516102049190612312565b60405180910390f35b34801561021a5760006000fd5b5061023560048036038101906102309190612115565b610d52565b005b3480156102445760006000fd5b5061025f600480360381019061025a91906120d6565b610d66565b005b34801561026e5760006000fd5b5061028960048036038101906102849190611fb3565b610f09565b6040516102969190612351565b60405180910390f35b3480156102ac5760006000fd5b506102c760048036038101906102c291906120d6565b610f5d565b005b3480156102d65760006000fd5b506102df6110a4565b6040516102ec919061232e565b60405180910390f35b3480156103025760006000fd5b5061030b61113e565b005b34801561031a5760006000fd5b506103356004803603810190610330919061201d565b61119c565b005b610351600480360381019061034c91906120d6565b611711565b60405161035e9190612312565b60405180910390f35b61036f61190c565b60405161037c9190612312565b60405180910390f35b3480156103925760006000fd5b506103ad60048036038101906103a89190611fde565b611d1f565b6040516103ba9190612351565b60405180910390f35b6060600060005080546103d59061271a565b80601f01602080910402602001604051908101604052809291908181526020018280546104019061271a565b801561044e5780601f106104235761010080835404028352916020019161044e565b820191906000526020600020905b81548152906001019060200180831161043157829003601f168201915b5050505050905061045a565b90565b600060008214806104f357506000600a60005060003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060005060008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060005054145b15156104ff5760006000fd5b81600a60005060003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060005060008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206000508190909055507f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b9253384846040516105be939291906122da565b60405180910390a1600190506105cf565b92915050565b600060036000505490506105e4565b90565b600081600960005060008673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060005054101580156106c3575081600a60005060008673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060005060003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206000505410155b15156106cf5760006000fd5b81600960005060008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206000828282505461072391906123a7565b92505081909090555081600960005060008673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008282825054610780919061260a565b92505081909090555081600a60005060008673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060005060003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206000828282505461081d919061260a565b9250508190909055507fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef848484604051610859939291906122da565b60405180910390a16001905061086a565b9392505050565b6000600260009054906101000a900460ff16905061088a565b90565b6000606460096000506000600860009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060005054101580156109d157506064600a6000506000600860009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206000506000600660009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206000505410155b15156109dd5760006000fd5b606460096000506000600860009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008282825054610a54919061260a565b9250508190909055506064600960005060003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008282825054610ab291906123a7565b9250508190909055507fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef600860009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16336064604051610b1193929190612232565b60405180910390a160006064600a6000506000600860009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206000506000600660009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060005054610bf1919061260a565b9050805080600a6000506000600860009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206000506000600660009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206000508190909055507f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925600860009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16600660009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1683604051610d3c939291906122da565b60405180910390a16001915050610d4f56505b90565b610d623382611db463ffffffff16565b5b50565b600860009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16141515610dc35760006000fd5b80600360008282825054610dd791906123a7565b92505081909090555080600960005060003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008282825054610e3491906123a7565b92505081909090555080600a60005060003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060005060008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206000508190909055507f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925338383604051610efc939291906122da565b60405180910390a15b5050565b6000600960005060008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600050549050610f58565b919050565b600a60005060008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060005060003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600050548111151515610ff25760006000fd5b80600a60005060008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060005060003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008282825054611086919061260a565b92505081909090555061109f8282611db463ffffffff16565b5b5050565b6060600160005080546110b69061271a565b80601f01602080910402602001604051908101604052809291908181526020018280546110e29061271a565b801561112f5780601f106111045761010080835404028352916020019161112f565b820191906000526020600020905b81548152906001019060200180831161111257829003601f168201915b5050505050905061113b565b90565b6000600960005060003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206000505490506111983382611db463ffffffff16565b505b565b6000600260009054906101000a900460ff1660ff16600a6111bd9190612488565b6003600050546111cd91906125af565b905080600960005060008773ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206000828282505461122391906123a7565b9250508190909055506040518060400160405280600481526020017f4b423234000000000000000000000000000000000000000000000000000000008152602001506000600050908051906020019061127d929190611ed5565b506040518060400160405280600481526020017f4b42323400000000000000000000000000000000000000000000000000000000815260200150600160005090805190602001906112cf929190611ed5565b5084600860006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555083600560006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555082600660006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555081600760006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506301312d00600a60005060003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206000506000600560009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206000508190909055507f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b92533600560009054906101000a900473ffffffffffffffffffffffffffffffffffffffff166301312d006040516114df9392919061226a565b60405180910390a1620f4240600a60005060003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206000506000600660009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206000508190909055507f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b92533600660009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16620f42406040516115f0939291906121fa565b60405180910390a1622dc6c0600a60005060003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206000506000600760009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206000508190909055507f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b92533600760009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16622dc6c0604051611701939291906122a2565b60405180910390a1505b50505050565b600081600960005060003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060005054101580156117fc5750600960005060008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206000505482600960005060008673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600050546117fa91906123a7565b115b15156118085760006000fd5b81600960005060003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206000828282505461185c919061260a565b92505081909090555081600960005060008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600082828250546118b991906123a7565b9250508190909055507fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef3384846040516118f5939291906122da565b60405180910390a160019050611906565b92915050565b600060003390506000349050808273ffffffffffffffffffffffffffffffffffffffff16311015151561193f5760006000fd5b60003373ffffffffffffffffffffffffffffffffffffffff166108fc839081150290604051600060405180830381858888f1935050505090506000151581151514156119915760009350505050611d1c565b6000600460005054836119a491906123fe565b90508060096000506000600860009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008282825054611a1c919061260a565b92505081909090555080600960005060003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008282825054611a7991906123a7565b9250508190909055507fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef600860009054906101000a900473ffffffffffffffffffffffffffffffffffffffff163383604051611ad7939291906122da565b60405180910390a1600081600a6000506000600860009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206000506000600560009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060005054611bb6919061260a565b9050805080600a6000506000600860009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206000506000600560009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206000508190909055507f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925600860009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16600560009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1683604051611d01939291906122da565b60405180910390a1600195505050505050611d1c5650505050505b90565b6000600a60005060008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060005060008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600050549050611dae565b92915050565b60008114151515611dc55760006000fd5b600960005060008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600050548111151515611e1a5760006000fd5b80600360008282825054611e2e919061260a565b92505081909090555080600960005060008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008282825054611e8b919061260a565b9250508190909055507fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef82600083604051611ec8939291906122da565b60405180910390a15b5050565b828054611ee19061271a565b90600052602060002090601f016020900481019282611f035760008555611f4f565b82601f10611f1c57805160ff1916838001178555611f4f565b82800160010185558215611f4f579182015b82811115611f4e5782518260005090905591602001919060010190611f2e565b5b509050611f5c9190611f60565b5090565b611f65565b80821115611f7f5760008181506000905550600101611f65565b509056612838565b600081359050611f9681612802565b5b92915050565b600081359050611fac8161281d565b5b92915050565b600060208284031215611fc65760006000fd5b6000611fd484828501611f87565b9150505b92915050565b6000600060408385031215611ff35760006000fd5b600061200185828601611f87565b925050602061201285828601611f87565b9150505b9250929050565b6000600060006000608085870312156120365760006000fd5b600061204487828801611f87565b945050602061205587828801611f87565b935050604061206687828801611f87565b925050606061207787828801611f87565b9150505b92959194509250565b6000600060006060848603121561209b5760006000fd5b60006120a986828701611f87565b93505060206120ba86828701611f87565b92505060406120cb86828701611f9d565b9150505b9250925092565b60006000604083850312156120eb5760006000fd5b60006120f985828601611f87565b925050602061210a85828601611f9d565b9150505b9250929050565b6000602082840312156121285760006000fd5b600061213684828501611f9d565b9150505b92915050565b6121498161263f565b82525b5050565b61215981612652565b82525b5050565b61216981612699565b82525b5050565b612179816126ac565b82525b5050565b612189816126bf565b82525b5050565b612199816126d2565b82525b5050565b60006121ab82612389565b6121b58185612395565b93506121c58185602086016126e5565b6121ce816127e2565b84019150505b92915050565b6121e381612680565b82525b5050565b6121f38161268b565b82525b5050565b600060608201905061220f6000830186612140565b61221c6020830185612140565b6122296040830184612160565b5b949350505050565b60006060820190506122476000830186612140565b6122546020830185612140565b6122616040830184612170565b5b949350505050565b600060608201905061227f6000830186612140565b61228c6020830185612140565b6122996040830184612180565b5b949350505050565b60006060820190506122b76000830186612140565b6122c46020830185612140565b6122d16040830184612190565b5b949350505050565b60006060820190506122ef6000830186612140565b6122fc6020830185612140565b61230960408301846121da565b5b949350505050565b60006020820190506123276000830184612150565b5b92915050565b6000602082019050818103600083015261234881846121a0565b90505b92915050565b600060208201905061236660008301846121da565b5b92915050565b600060208201905061238260008301846121ea565b5b92915050565b6000815190505b919050565b60008282526020820190505b92915050565b60006123b282612680565b91506123bd83612680565b9250827fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff038211156123f2576123f161274f565b5b82820190505b92915050565b600061240982612680565b915061241483612680565b925082151561242657612425612780565b5b82820490505b92915050565b600060008291508390505b600185111561247e578086048111156124595761245861274f565b5b60018516156124685780820291505b8081029050612476856127f4565b94505b61243d565b5b94509492505050565b600061249382612680565b915061249e83612680565b92506124cb7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff84846124d4565b90505b92915050565b60008215156124e657600190506125a8565b8115156124f657600090506125a8565b816001811461250c576002811461251a57612549565b60019150506125a856612549565b60ff84111561252c5761252b61274f565b5b8360020a9150848211156125435761254261274f565b5b506125a8565b5060208310610133831016604e8410600b841016171561257e5782820a9050838111156125795761257861274f565b5b6125a8565b61258b8484846001612432565b925090508184048111156125a2576125a161274f565b5b81810290505b9392505050565b60006125ba82612680565b91506125c583612680565b9250817fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff04831182151516156125fe576125fd61274f565b5b82820290505b92915050565b600061261582612680565b915061262083612680565b9250828210156126335761263261274f565b5b82820390505b92915050565b600061264a8261265f565b90505b919050565b600081151590505b919050565b600073ffffffffffffffffffffffffffffffffffffffff821690505b919050565b60008190505b919050565b600060ff821690505b919050565b60006126a482612680565b90505b919050565b60006126b782612680565b90505b919050565b60006126ca82612680565b90505b919050565b60006126dd82612680565b90505b919050565b60005b838110156127045780820151818401525b6020810190506126e8565b83811115612713576000848401525b505b505050565b60006002820490506001821680151561273457607f821691505b60208210811415612748576127476127b1565b5b505b919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601260045260246000fd5b565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b565b6000601f19601f83011690505b919050565b60008160011c90505b919050565b61280b8161263f565b811415156128195760006000fd5b5b50565b61282681612680565b811415156128345760006000fd5b5b50565bfea2646970667358221220d940b28a655b30388787f4eee1ead29a7fd32d7def7d4910517dec22f4c3fda564736f6c63430008030033"

// DeployKbtoken deploys a new Ethereum contract, binding an instance of Kbtoken to it.
func DeployKbtoken(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Kbtoken, error) {
	parsed, err := abi.JSON(strings.NewReader(KbtokenABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(KbtokenBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Kbtoken{KbtokenCaller: KbtokenCaller{contract: contract}, KbtokenTransactor: KbtokenTransactor{contract: contract}, KbtokenFilterer: KbtokenFilterer{contract: contract}}, nil
}

// Kbtoken is an auto generated Go binding around an Ethereum contract.
type Kbtoken struct {
	KbtokenCaller     // Read-only binding to the contract
	KbtokenTransactor // Write-only binding to the contract
	KbtokenFilterer   // Log filterer for contract events
}

// KbtokenCaller is an auto generated read-only Go binding around an Ethereum contract.
type KbtokenCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// KbtokenTransactor is an auto generated write-only Go binding around an Ethereum contract.
type KbtokenTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// KbtokenFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type KbtokenFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// KbtokenSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type KbtokenSession struct {
	Contract     *Kbtoken          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// KbtokenCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type KbtokenCallerSession struct {
	Contract *KbtokenCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// KbtokenTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type KbtokenTransactorSession struct {
	Contract     *KbtokenTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// KbtokenRaw is an auto generated low-level Go binding around an Ethereum contract.
type KbtokenRaw struct {
	Contract *Kbtoken // Generic contract binding to access the raw methods on
}

// KbtokenCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type KbtokenCallerRaw struct {
	Contract *KbtokenCaller // Generic read-only contract binding to access the raw methods on
}

// KbtokenTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type KbtokenTransactorRaw struct {
	Contract *KbtokenTransactor // Generic write-only contract binding to access the raw methods on
}

// NewKbtoken creates a new instance of Kbtoken, bound to a specific deployed contract.
func NewKbtoken(address common.Address, backend bind.ContractBackend) (*Kbtoken, error) {
	contract, err := bindKbtoken(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Kbtoken{KbtokenCaller: KbtokenCaller{contract: contract}, KbtokenTransactor: KbtokenTransactor{contract: contract}, KbtokenFilterer: KbtokenFilterer{contract: contract}}, nil
}

// NewKbtokenCaller creates a new read-only instance of Kbtoken, bound to a specific deployed contract.
func NewKbtokenCaller(address common.Address, caller bind.ContractCaller) (*KbtokenCaller, error) {
	contract, err := bindKbtoken(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &KbtokenCaller{contract: contract}, nil
}

// NewKbtokenTransactor creates a new write-only instance of Kbtoken, bound to a specific deployed contract.
func NewKbtokenTransactor(address common.Address, transactor bind.ContractTransactor) (*KbtokenTransactor, error) {
	contract, err := bindKbtoken(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &KbtokenTransactor{contract: contract}, nil
}

// NewKbtokenFilterer creates a new log filterer instance of Kbtoken, bound to a specific deployed contract.
func NewKbtokenFilterer(address common.Address, filterer bind.ContractFilterer) (*KbtokenFilterer, error) {
	contract, err := bindKbtoken(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &KbtokenFilterer{contract: contract}, nil
}

// bindKbtoken binds a generic wrapper to an already deployed contract.
func bindKbtoken(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(KbtokenABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Kbtoken *KbtokenRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Kbtoken.Contract.KbtokenCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Kbtoken *KbtokenRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Kbtoken.Contract.KbtokenTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Kbtoken *KbtokenRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Kbtoken.Contract.KbtokenTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Kbtoken *KbtokenCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Kbtoken.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Kbtoken *KbtokenTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Kbtoken.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Kbtoken *KbtokenTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Kbtoken.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address _owner, address _spender) view returns(uint256 remaining)
func (_Kbtoken *KbtokenCaller) Allowance(opts *bind.CallOpts, _owner common.Address, _spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Kbtoken.contract.Call(opts, &out, "allowance", _owner, _spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address _owner, address _spender) view returns(uint256 remaining)
func (_Kbtoken *KbtokenSession) Allowance(_owner common.Address, _spender common.Address) (*big.Int, error) {
	return _Kbtoken.Contract.Allowance(&_Kbtoken.CallOpts, _owner, _spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address _owner, address _spender) view returns(uint256 remaining)
func (_Kbtoken *KbtokenCallerSession) Allowance(_owner common.Address, _spender common.Address) (*big.Int, error) {
	return _Kbtoken.Contract.Allowance(&_Kbtoken.CallOpts, _owner, _spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address _owner) view returns(uint256 balance)
func (_Kbtoken *KbtokenCaller) BalanceOf(opts *bind.CallOpts, _owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Kbtoken.contract.Call(opts, &out, "balanceOf", _owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address _owner) view returns(uint256 balance)
func (_Kbtoken *KbtokenSession) BalanceOf(_owner common.Address) (*big.Int, error) {
	return _Kbtoken.Contract.BalanceOf(&_Kbtoken.CallOpts, _owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address _owner) view returns(uint256 balance)
func (_Kbtoken *KbtokenCallerSession) BalanceOf(_owner common.Address) (*big.Int, error) {
	return _Kbtoken.Contract.BalanceOf(&_Kbtoken.CallOpts, _owner)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_Kbtoken *KbtokenCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Kbtoken.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_Kbtoken *KbtokenSession) Decimals() (uint8, error) {
	return _Kbtoken.Contract.Decimals(&_Kbtoken.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_Kbtoken *KbtokenCallerSession) Decimals() (uint8, error) {
	return _Kbtoken.Contract.Decimals(&_Kbtoken.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Kbtoken *KbtokenCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Kbtoken.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Kbtoken *KbtokenSession) Name() (string, error) {
	return _Kbtoken.Contract.Name(&_Kbtoken.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Kbtoken *KbtokenCallerSession) Name() (string, error) {
	return _Kbtoken.Contract.Name(&_Kbtoken.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Kbtoken *KbtokenCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Kbtoken.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Kbtoken *KbtokenSession) Symbol() (string, error) {
	return _Kbtoken.Contract.Symbol(&_Kbtoken.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Kbtoken *KbtokenCallerSession) Symbol() (string, error) {
	return _Kbtoken.Contract.Symbol(&_Kbtoken.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Kbtoken *KbtokenCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Kbtoken.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Kbtoken *KbtokenSession) TotalSupply() (*big.Int, error) {
	return _Kbtoken.Contract.TotalSupply(&_Kbtoken.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Kbtoken *KbtokenCallerSession) TotalSupply() (*big.Int, error) {
	return _Kbtoken.Contract.TotalSupply(&_Kbtoken.CallOpts)
}

// Airdrop is a paid mutator transaction binding the contract method 0x3884d635.
//
// Solidity: function airdrop() payable returns(bool success)
func (_Kbtoken *KbtokenTransactor) Airdrop(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Kbtoken.contract.Transact(opts, "airdrop")
}

// Airdrop is a paid mutator transaction binding the contract method 0x3884d635.
//
// Solidity: function airdrop() payable returns(bool success)
func (_Kbtoken *KbtokenSession) Airdrop() (*types.Transaction, error) {
	return _Kbtoken.Contract.Airdrop(&_Kbtoken.TransactOpts)
}

// Airdrop is a paid mutator transaction binding the contract method 0x3884d635.
//
// Solidity: function airdrop() payable returns(bool success)
func (_Kbtoken *KbtokenTransactorSession) Airdrop() (*types.Transaction, error) {
	return _Kbtoken.Contract.Airdrop(&_Kbtoken.TransactOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address _spender, uint256 _value) returns(bool success)
func (_Kbtoken *KbtokenTransactor) Approve(opts *bind.TransactOpts, _spender common.Address, _value *big.Int) (*types.Transaction, error) {
	return _Kbtoken.contract.Transact(opts, "approve", _spender, _value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address _spender, uint256 _value) returns(bool success)
func (_Kbtoken *KbtokenSession) Approve(_spender common.Address, _value *big.Int) (*types.Transaction, error) {
	return _Kbtoken.Contract.Approve(&_Kbtoken.TransactOpts, _spender, _value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address _spender, uint256 _value) returns(bool success)
func (_Kbtoken *KbtokenTransactorSession) Approve(_spender common.Address, _value *big.Int) (*types.Transaction, error) {
	return _Kbtoken.Contract.Approve(&_Kbtoken.TransactOpts, _spender, _value)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 amount) returns()
func (_Kbtoken *KbtokenTransactor) Burn(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _Kbtoken.contract.Transact(opts, "burn", amount)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 amount) returns()
func (_Kbtoken *KbtokenSession) Burn(amount *big.Int) (*types.Transaction, error) {
	return _Kbtoken.Contract.Burn(&_Kbtoken.TransactOpts, amount)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 amount) returns()
func (_Kbtoken *KbtokenTransactorSession) Burn(amount *big.Int) (*types.Transaction, error) {
	return _Kbtoken.Contract.Burn(&_Kbtoken.TransactOpts, amount)
}

// BurnAll is a paid mutator transaction binding the contract method 0x9975038c.
//
// Solidity: function burnAll() returns()
func (_Kbtoken *KbtokenTransactor) BurnAll(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Kbtoken.contract.Transact(opts, "burnAll")
}

// BurnAll is a paid mutator transaction binding the contract method 0x9975038c.
//
// Solidity: function burnAll() returns()
func (_Kbtoken *KbtokenSession) BurnAll() (*types.Transaction, error) {
	return _Kbtoken.Contract.BurnAll(&_Kbtoken.TransactOpts)
}

// BurnAll is a paid mutator transaction binding the contract method 0x9975038c.
//
// Solidity: function burnAll() returns()
func (_Kbtoken *KbtokenTransactorSession) BurnAll() (*types.Transaction, error) {
	return _Kbtoken.Contract.BurnAll(&_Kbtoken.TransactOpts)
}

// BurnFrom is a paid mutator transaction binding the contract method 0x79cc6790.
//
// Solidity: function burnFrom(address account, uint256 amount) returns()
func (_Kbtoken *KbtokenTransactor) BurnFrom(opts *bind.TransactOpts, account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Kbtoken.contract.Transact(opts, "burnFrom", account, amount)
}

// BurnFrom is a paid mutator transaction binding the contract method 0x79cc6790.
//
// Solidity: function burnFrom(address account, uint256 amount) returns()
func (_Kbtoken *KbtokenSession) BurnFrom(account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Kbtoken.Contract.BurnFrom(&_Kbtoken.TransactOpts, account, amount)
}

// BurnFrom is a paid mutator transaction binding the contract method 0x79cc6790.
//
// Solidity: function burnFrom(address account, uint256 amount) returns()
func (_Kbtoken *KbtokenTransactorSession) BurnFrom(account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Kbtoken.Contract.BurnFrom(&_Kbtoken.TransactOpts, account, amount)
}

// Exchange is a paid mutator transaction binding the contract method 0xd2f7265a.
//
// Solidity: function exchange() payable returns(bool success)
func (_Kbtoken *KbtokenTransactor) Exchange(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Kbtoken.contract.Transact(opts, "exchange")
}

// Exchange is a paid mutator transaction binding the contract method 0xd2f7265a.
//
// Solidity: function exchange() payable returns(bool success)
func (_Kbtoken *KbtokenSession) Exchange() (*types.Transaction, error) {
	return _Kbtoken.Contract.Exchange(&_Kbtoken.TransactOpts)
}

// Exchange is a paid mutator transaction binding the contract method 0xd2f7265a.
//
// Solidity: function exchange() payable returns(bool success)
func (_Kbtoken *KbtokenTransactorSession) Exchange() (*types.Transaction, error) {
	return _Kbtoken.Contract.Exchange(&_Kbtoken.TransactOpts)
}

// InitCoin is a paid mutator transaction binding the contract method 0xa67c9264.
//
// Solidity: function initCoin(address holder, address exchange, address launch, address reserved) returns()
func (_Kbtoken *KbtokenTransactor) InitCoin(opts *bind.TransactOpts, holder common.Address, exchange common.Address, launch common.Address, reserved common.Address) (*types.Transaction, error) {
	return _Kbtoken.contract.Transact(opts, "initCoin", holder, exchange, launch, reserved)
}

// InitCoin is a paid mutator transaction binding the contract method 0xa67c9264.
//
// Solidity: function initCoin(address holder, address exchange, address launch, address reserved) returns()
func (_Kbtoken *KbtokenSession) InitCoin(holder common.Address, exchange common.Address, launch common.Address, reserved common.Address) (*types.Transaction, error) {
	return _Kbtoken.Contract.InitCoin(&_Kbtoken.TransactOpts, holder, exchange, launch, reserved)
}

// InitCoin is a paid mutator transaction binding the contract method 0xa67c9264.
//
// Solidity: function initCoin(address holder, address exchange, address launch, address reserved) returns()
func (_Kbtoken *KbtokenTransactorSession) InitCoin(holder common.Address, exchange common.Address, launch common.Address, reserved common.Address) (*types.Transaction, error) {
	return _Kbtoken.Contract.InitCoin(&_Kbtoken.TransactOpts, holder, exchange, launch, reserved)
}

// Seo is a paid mutator transaction binding the contract method 0x70238ecd.
//
// Solidity: function seo(address _to, uint256 _value) returns()
func (_Kbtoken *KbtokenTransactor) Seo(opts *bind.TransactOpts, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _Kbtoken.contract.Transact(opts, "seo", _to, _value)
}

// Seo is a paid mutator transaction binding the contract method 0x70238ecd.
//
// Solidity: function seo(address _to, uint256 _value) returns()
func (_Kbtoken *KbtokenSession) Seo(_to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _Kbtoken.Contract.Seo(&_Kbtoken.TransactOpts, _to, _value)
}

// Seo is a paid mutator transaction binding the contract method 0x70238ecd.
//
// Solidity: function seo(address _to, uint256 _value) returns()
func (_Kbtoken *KbtokenTransactorSession) Seo(_to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _Kbtoken.Contract.Seo(&_Kbtoken.TransactOpts, _to, _value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address _to, uint256 _value) payable returns(bool success)
func (_Kbtoken *KbtokenTransactor) Transfer(opts *bind.TransactOpts, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _Kbtoken.contract.Transact(opts, "transfer", _to, _value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address _to, uint256 _value) payable returns(bool success)
func (_Kbtoken *KbtokenSession) Transfer(_to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _Kbtoken.Contract.Transfer(&_Kbtoken.TransactOpts, _to, _value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address _to, uint256 _value) payable returns(bool success)
func (_Kbtoken *KbtokenTransactorSession) Transfer(_to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _Kbtoken.Contract.Transfer(&_Kbtoken.TransactOpts, _to, _value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address _from, address _to, uint256 _value) payable returns(bool success)
func (_Kbtoken *KbtokenTransactor) TransferFrom(opts *bind.TransactOpts, _from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _Kbtoken.contract.Transact(opts, "transferFrom", _from, _to, _value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address _from, address _to, uint256 _value) payable returns(bool success)
func (_Kbtoken *KbtokenSession) TransferFrom(_from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _Kbtoken.Contract.TransferFrom(&_Kbtoken.TransactOpts, _from, _to, _value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address _from, address _to, uint256 _value) payable returns(bool success)
func (_Kbtoken *KbtokenTransactorSession) TransferFrom(_from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _Kbtoken.Contract.TransferFrom(&_Kbtoken.TransactOpts, _from, _to, _value)
}

// KbtokenApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the Kbtoken contract.
type KbtokenApprovalIterator struct {
	Event *KbtokenApproval // Event containing the contract specifics and raw log

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
func (it *KbtokenApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(KbtokenApproval)
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
		it.Event = new(KbtokenApproval)
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
func (it *KbtokenApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *KbtokenApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// KbtokenApproval represents a Approval event raised by the Kbtoken contract.
type KbtokenApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address owner, address spender, uint256 value)
func (_Kbtoken *KbtokenFilterer) FilterApproval(opts *bind.FilterOpts) (*KbtokenApprovalIterator, error) {

	logs, sub, err := _Kbtoken.contract.FilterLogs(opts, "Approval")
	if err != nil {
		return nil, err
	}
	return &KbtokenApprovalIterator{contract: _Kbtoken.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address owner, address spender, uint256 value)
func (_Kbtoken *KbtokenFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *KbtokenApproval) (event.Subscription, error) {

	logs, sub, err := _Kbtoken.contract.WatchLogs(opts, "Approval")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(KbtokenApproval)
				if err := _Kbtoken.contract.UnpackLog(event, "Approval", log); err != nil {
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
func (_Kbtoken *KbtokenFilterer) ParseApproval(log types.Log) (*KbtokenApproval, error) {
	event := new(KbtokenApproval)
	if err := _Kbtoken.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// KbtokenTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the Kbtoken contract.
type KbtokenTransferIterator struct {
	Event *KbtokenTransfer // Event containing the contract specifics and raw log

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
func (it *KbtokenTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(KbtokenTransfer)
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
		it.Event = new(KbtokenTransfer)
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
func (it *KbtokenTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *KbtokenTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// KbtokenTransfer represents a Transfer event raised by the Kbtoken contract.
type KbtokenTransfer struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address owner, address spender, uint256 value)
func (_Kbtoken *KbtokenFilterer) FilterTransfer(opts *bind.FilterOpts) (*KbtokenTransferIterator, error) {

	logs, sub, err := _Kbtoken.contract.FilterLogs(opts, "Transfer")
	if err != nil {
		return nil, err
	}
	return &KbtokenTransferIterator{contract: _Kbtoken.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address owner, address spender, uint256 value)
func (_Kbtoken *KbtokenFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *KbtokenTransfer) (event.Subscription, error) {

	logs, sub, err := _Kbtoken.contract.WatchLogs(opts, "Transfer")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(KbtokenTransfer)
				if err := _Kbtoken.contract.UnpackLog(event, "Transfer", log); err != nil {
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
func (_Kbtoken *KbtokenFilterer) ParseTransfer(log types.Log) (*KbtokenTransfer, error) {
	event := new(KbtokenTransfer)
	if err := _Kbtoken.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
