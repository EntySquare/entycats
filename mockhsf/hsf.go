// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package hsf

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

// HsfABI is the input ABI used to generate the binding from.
const HsfABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"remaining\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"burn\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"burnAll\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"burnFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"initialSupply\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"holder\",\"type\":\"address\"}],\"name\":\"initCoin\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ivToken\",\"outputs\":[{\"internalType\":\"contractinvestors\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"manager\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_class\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"placeAndTransfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"manager\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_class\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"removeAndTransfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"seo\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"}],\"stateMutability\":\"payable\",\"type\":\"function\"}]"

// HsfBin is the compiled bytecode used for deploying new contracts.
var HsfBin = "0x60806040526002600360006101000a81548160ff021916908360ff16021790555034801561002c57600080fd5b50611f8c8061003c6000396000f3fe6080604052600436106101095760003560e01c806370a082311161009557806395ef9ca31161006457806395ef9ca3146103655780639975038c14610395578063a9059cbb146103ac578063dd62ed3e146103dc578063f50c3f4c1461041957610109565b806370a08231146102a957806379cc6790146102e65780638da5cb5b1461030f57806395d89b411461033a57610109565b806323b872dd116100dc57806323b872dd146101cc578063313ce567146101fc57806342966c6814610227578063528d0f491461025057806370238ecd1461028057610109565b806306fdde031461010e578063095ea7b3146101395780630e2bcdb81461017657806318160ddd146101a1575b600080fd5b34801561011a57600080fd5b50610123610442565b6040516101309190611aee565b60405180910390f35b34801561014557600080fd5b50610160600480360381019061015b9190611892565b6104d4565b60405161016d9190611ab8565b60405180910390f35b34801561018257600080fd5b5061018b61062f565b6040516101989190611ad3565b60405180910390f35b3480156101ad57600080fd5b506101b6610653565b6040516101c39190611b10565b60405180910390f35b6101e660048036038101906101e19190611843565b61065d565b6040516101f39190611ab8565b60405180910390f35b34801561020857600080fd5b506102116108b8565b60405161021e9190611b2b565b60405180910390f35b34801561023357600080fd5b5061024e60048036038101906102499190611946565b6108cf565b005b61026a600480360381019061026591906118ce565b6108dc565b6040516102779190611ab8565b60405180910390f35b34801561028c57600080fd5b506102a760048036038101906102a29190611892565b610ba0565b005b3480156102b557600080fd5b506102d060048036038101906102cb91906117de565b610c6d565b6040516102dd9190611b10565b60405180910390f35b3480156102f257600080fd5b5061030d60048036038101906103089190611892565b610cb6565b005b34801561031b57600080fd5b50610324610de0565b6040516103319190611a2f565b60405180910390f35b34801561034657600080fd5b5061034f610e06565b60405161035c9190611aee565b60405180910390f35b61037f600480360381019061037a91906118ce565b610e98565b60405161038c9190611ab8565b60405180910390f35b3480156103a157600080fd5b506103aa6111e3565b005b6103c660048036038101906103c19190611892565b611234565b6040516103d39190611ab8565b60405180910390f35b3480156103e857600080fd5b5061040360048036038101906103fe9190611807565b611407565b6040516104109190611b10565b60405180910390f35b34801561042557600080fd5b50610440600480360381019061043b919061196f565b61148e565b005b60606001805461045190611e63565b80601f016020809104026020016040519081016040528092919081815260200182805461047d90611e63565b80156104ca5780601f1061049f576101008083540402835291602001916104ca565b820191906000526020600020905b8154815290600101906020018083116104ad57829003601f168201915b5050505050905090565b6000808214806105605750600760003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020548210155b61056957600080fd5b81600760003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020819055507f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b92533848460405161061d93929190611a4a565b60405180910390a16001905092915050565b60008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b6000600454905090565b600081600660008673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020541015801561072a575081600760008673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000205410155b61073357600080fd5b81600660008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008282546107829190611b62565b9250508190555081600660008673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008282546107d89190611d83565b9250508190555081600760008673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600082825461086b9190611d83565b925050819055507fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef8484846040516108a593929190611a4a565b60405180910390a1600190509392505050565b6000600360009054906101000a900460ff16905090565b6108d933826115f3565b50565b600081600660003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054101580156109b55750600660008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000205482600660008773ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020546109b39190611b62565b115b6109be57600080fd5b81600660003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206000828254610a0d9190611d83565b9250508190555081600660008673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206000828254610a639190611b62565b92505081905550836000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663f231a4f63385856040518463ffffffff1660e01b8152600401610b0793929190611a81565b602060405180830381600087803b158015610b2157600080fd5b505af1158015610b35573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610b59919061191d565b507fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef338584604051610b8d93929190611a4a565b60405180910390a1600190509392505050565b600560009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614610bfa57600080fd5b8060046000828254610c0c9190611b62565b9250508190555080600660008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206000828254610c629190611b62565b925050819055505050565b6000600660008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020549050919050565b600760008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054811115610d3f57600080fd5b80600760008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206000828254610dcb9190611d83565b92505081905550610ddc82826115f3565b5050565b600560009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b606060028054610e1590611e63565b80601f0160208091040260200160405190810160405280929190818152602001828054610e4190611e63565b8015610e8e5780601f10610e6357610100808354040283529160200191610e8e565b820191906000526020600020905b815481529060010190602001808311610e7157829003601f168201915b5050505050905090565b600081600660008673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000205410158015610f65575081600760008673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000205410155b610f6e57600080fd5b81600660003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206000828254610fbd9190611b62565b9250508190555081600660008673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008282546110139190611d83565b9250508190555081600760008673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008282546110a69190611d83565b92505081905550836000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16638a9faad23385856040518463ffffffff1660e01b815260040161114a93929190611a81565b602060405180830381600087803b15801561116457600080fd5b505af1158015611178573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061119c919061191d565b507fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef8433846040516111d093929190611a4a565b60405180910390a1600190509392505050565b6000600660003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054905061123133826115f3565b50565b600081600660003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020541015801561130d5750600660008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000205482600660008673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000205461130b9190611b62565b115b61131657600080fd5b81600660003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008282546113659190611d83565b9250508190555081600660008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008282546113bb9190611b62565b925050819055507fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef3384846040516113f593929190611a4a565b60405180910390a16001905092915050565b6000600760008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054905092915050565b600360009054906101000a900460ff1660ff16600a6114ad9190611c0b565b826114b89190611d29565b600481905550600454600660008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600082825461150f9190611b62565b925050819055506040518060400160405280601081526020017f68696c6c73746f6e6566696e616e636500000000000000000000000000000000815250600190805190602001906115619291906116fc565b506040518060400160405280600381526020017f4853460000000000000000000000000000000000000000000000000000000000815250600290805190602001906115ad9291906116fc565b5080600560006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055505050565b600081141561160157600080fd5b600660008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000205481111561164d57600080fd5b806004600082825461165f9190611d83565b9250508190555080600660008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008282546116b59190611d83565b925050819055507fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef826000836040516116f093929190611a4a565b60405180910390a15050565b82805461170890611e63565b90600052602060002090601f01602090048101928261172a5760008555611771565b82601f1061174357805160ff1916838001178555611771565b82800160010185558215611771579182015b82811115611770578251825591602001919060010190611755565b5b50905061177e9190611782565b5090565b5b8082111561179b576000816000905550600101611783565b5090565b6000813590506117ae81611f11565b92915050565b6000815190506117c381611f28565b92915050565b6000813590506117d881611f3f565b92915050565b6000602082840312156117f057600080fd5b60006117fe8482850161179f565b91505092915050565b6000806040838503121561181a57600080fd5b60006118288582860161179f565b92505060206118398582860161179f565b9150509250929050565b60008060006060848603121561185857600080fd5b60006118668682870161179f565b93505060206118778682870161179f565b9250506040611888868287016117c9565b9150509250925092565b600080604083850312156118a557600080fd5b60006118b38582860161179f565b92505060206118c4858286016117c9565b9150509250929050565b6000806000606084860312156118e357600080fd5b60006118f18682870161179f565b9350506020611902868287016117c9565b9250506040611913868287016117c9565b9150509250925092565b60006020828403121561192f57600080fd5b600061193d848285016117b4565b91505092915050565b60006020828403121561195857600080fd5b6000611966848285016117c9565b91505092915050565b6000806040838503121561198257600080fd5b6000611990858286016117c9565b92505060206119a18582860161179f565b9150509250929050565b6119b481611db7565b82525050565b6119c381611dc9565b82525050565b6119d281611e0c565b82525050565b60006119e382611b46565b6119ed8185611b51565b93506119fd818560208601611e30565b611a0681611ef3565b840191505092915050565b611a1a81611df5565b82525050565b611a2981611dff565b82525050565b6000602082019050611a4460008301846119ab565b92915050565b6000606082019050611a5f60008301866119ab565b611a6c60208301856119ab565b611a796040830184611a11565b949350505050565b6000606082019050611a9660008301866119ab565b611aa36020830185611a11565b611ab06040830184611a11565b949350505050565b6000602082019050611acd60008301846119ba565b92915050565b6000602082019050611ae860008301846119c9565b92915050565b60006020820190508181036000830152611b0881846119d8565b905092915050565b6000602082019050611b256000830184611a11565b92915050565b6000602082019050611b406000830184611a20565b92915050565b600081519050919050565b600082825260208201905092915050565b6000611b6d82611df5565b9150611b7883611df5565b9250827fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff03821115611bad57611bac611e95565b5b828201905092915050565b6000808291508390505b6001851115611c0257808604811115611bde57611bdd611e95565b5b6001851615611bed5780820291505b8081029050611bfb85611f04565b9450611bc2565b94509492505050565b6000611c1682611df5565b9150611c2183611df5565b9250611c4e7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8484611c56565b905092915050565b600082611c665760019050611d22565b81611c745760009050611d22565b8160018114611c8a5760028114611c9457611cc3565b6001915050611d22565b60ff841115611ca657611ca5611e95565b5b8360020a915084821115611cbd57611cbc611e95565b5b50611d22565b5060208310610133831016604e8410600b8410161715611cf85782820a905083811115611cf357611cf2611e95565b5b611d22565b611d058484846001611bb8565b92509050818404811115611d1c57611d1b611e95565b5b81810290505b9392505050565b6000611d3482611df5565b9150611d3f83611df5565b9250817fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0483118215151615611d7857611d77611e95565b5b828202905092915050565b6000611d8e82611df5565b9150611d9983611df5565b925082821015611dac57611dab611e95565b5b828203905092915050565b6000611dc282611dd5565b9050919050565b60008115159050919050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000819050919050565b600060ff82169050919050565b6000611e1782611e1e565b9050919050565b6000611e2982611dd5565b9050919050565b60005b83811015611e4e578082015181840152602081019050611e33565b83811115611e5d576000848401525b50505050565b60006002820490506001821680611e7b57607f821691505b60208210811415611e8f57611e8e611ec4565b5b50919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b6000601f19601f8301169050919050565b60008160011c9050919050565b611f1a81611db7565b8114611f2557600080fd5b50565b611f3181611dc9565b8114611f3c57600080fd5b50565b611f4881611df5565b8114611f5357600080fd5b5056fea2646970667358221220af5f83b1bc8171adc0ec4d7a741725c9ee0dc6e90d1135fe3d517051b33a234364736f6c63430008030033"

// DeployHsf deploys a new Ethereum contract, binding an instance of Hsf to it.
func DeployHsf(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Hsf, error) {
	parsed, err := abi.JSON(strings.NewReader(HsfABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(HsfBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Hsf{HsfCaller: HsfCaller{contract: contract}, HsfTransactor: HsfTransactor{contract: contract}, HsfFilterer: HsfFilterer{contract: contract}}, nil
}

// Hsf is an auto generated Go binding around an Ethereum contract.
type Hsf struct {
	HsfCaller     // Read-only binding to the contract
	HsfTransactor // Write-only binding to the contract
	HsfFilterer   // Log filterer for contract events
}

// HsfCaller is an auto generated read-only Go binding around an Ethereum contract.
type HsfCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// HsfTransactor is an auto generated write-only Go binding around an Ethereum contract.
type HsfTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// HsfFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type HsfFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// HsfSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type HsfSession struct {
	Contract     *Hsf              // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// HsfCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type HsfCallerSession struct {
	Contract *HsfCaller    // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// HsfTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type HsfTransactorSession struct {
	Contract     *HsfTransactor    // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// HsfRaw is an auto generated low-level Go binding around an Ethereum contract.
type HsfRaw struct {
	Contract *Hsf // Generic contract binding to access the raw methods on
}

// HsfCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type HsfCallerRaw struct {
	Contract *HsfCaller // Generic read-only contract binding to access the raw methods on
}

// HsfTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type HsfTransactorRaw struct {
	Contract *HsfTransactor // Generic write-only contract binding to access the raw methods on
}

// NewHsf creates a new instance of Hsf, bound to a specific deployed contract.
func NewHsf(address common.Address, backend bind.ContractBackend) (*Hsf, error) {
	contract, err := bindHsf(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Hsf{HsfCaller: HsfCaller{contract: contract}, HsfTransactor: HsfTransactor{contract: contract}, HsfFilterer: HsfFilterer{contract: contract}}, nil
}

// NewHsfCaller creates a new read-only instance of Hsf, bound to a specific deployed contract.
func NewHsfCaller(address common.Address, caller bind.ContractCaller) (*HsfCaller, error) {
	contract, err := bindHsf(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &HsfCaller{contract: contract}, nil
}

// NewHsfTransactor creates a new write-only instance of Hsf, bound to a specific deployed contract.
func NewHsfTransactor(address common.Address, transactor bind.ContractTransactor) (*HsfTransactor, error) {
	contract, err := bindHsf(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &HsfTransactor{contract: contract}, nil
}

// NewHsfFilterer creates a new log filterer instance of Hsf, bound to a specific deployed contract.
func NewHsfFilterer(address common.Address, filterer bind.ContractFilterer) (*HsfFilterer, error) {
	contract, err := bindHsf(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &HsfFilterer{contract: contract}, nil
}

// bindHsf binds a generic wrapper to an already deployed contract.
func bindHsf(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(HsfABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Hsf *HsfRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Hsf.Contract.HsfCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Hsf *HsfRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Hsf.Contract.HsfTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Hsf *HsfRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Hsf.Contract.HsfTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Hsf *HsfCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Hsf.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Hsf *HsfTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Hsf.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Hsf *HsfTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Hsf.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address _owner, address _spender) view returns(uint256 remaining)
func (_Hsf *HsfCaller) Allowance(opts *bind.CallOpts, _owner common.Address, _spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Hsf.contract.Call(opts, &out, "allowance", _owner, _spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address _owner, address _spender) view returns(uint256 remaining)
func (_Hsf *HsfSession) Allowance(_owner common.Address, _spender common.Address) (*big.Int, error) {
	return _Hsf.Contract.Allowance(&_Hsf.CallOpts, _owner, _spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address _owner, address _spender) view returns(uint256 remaining)
func (_Hsf *HsfCallerSession) Allowance(_owner common.Address, _spender common.Address) (*big.Int, error) {
	return _Hsf.Contract.Allowance(&_Hsf.CallOpts, _owner, _spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address _owner) view returns(uint256 balance)
func (_Hsf *HsfCaller) BalanceOf(opts *bind.CallOpts, _owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Hsf.contract.Call(opts, &out, "balanceOf", _owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address _owner) view returns(uint256 balance)
func (_Hsf *HsfSession) BalanceOf(_owner common.Address) (*big.Int, error) {
	return _Hsf.Contract.BalanceOf(&_Hsf.CallOpts, _owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address _owner) view returns(uint256 balance)
func (_Hsf *HsfCallerSession) BalanceOf(_owner common.Address) (*big.Int, error) {
	return _Hsf.Contract.BalanceOf(&_Hsf.CallOpts, _owner)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_Hsf *HsfCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Hsf.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_Hsf *HsfSession) Decimals() (uint8, error) {
	return _Hsf.Contract.Decimals(&_Hsf.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_Hsf *HsfCallerSession) Decimals() (uint8, error) {
	return _Hsf.Contract.Decimals(&_Hsf.CallOpts)
}

// IvToken is a free data retrieval call binding the contract method 0x0e2bcdb8.
//
// Solidity: function ivToken() view returns(address)
func (_Hsf *HsfCaller) IvToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Hsf.contract.Call(opts, &out, "ivToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// IvToken is a free data retrieval call binding the contract method 0x0e2bcdb8.
//
// Solidity: function ivToken() view returns(address)
func (_Hsf *HsfSession) IvToken() (common.Address, error) {
	return _Hsf.Contract.IvToken(&_Hsf.CallOpts)
}

// IvToken is a free data retrieval call binding the contract method 0x0e2bcdb8.
//
// Solidity: function ivToken() view returns(address)
func (_Hsf *HsfCallerSession) IvToken() (common.Address, error) {
	return _Hsf.Contract.IvToken(&_Hsf.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Hsf *HsfCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Hsf.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Hsf *HsfSession) Name() (string, error) {
	return _Hsf.Contract.Name(&_Hsf.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Hsf *HsfCallerSession) Name() (string, error) {
	return _Hsf.Contract.Name(&_Hsf.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Hsf *HsfCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Hsf.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Hsf *HsfSession) Owner() (common.Address, error) {
	return _Hsf.Contract.Owner(&_Hsf.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Hsf *HsfCallerSession) Owner() (common.Address, error) {
	return _Hsf.Contract.Owner(&_Hsf.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Hsf *HsfCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Hsf.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Hsf *HsfSession) Symbol() (string, error) {
	return _Hsf.Contract.Symbol(&_Hsf.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Hsf *HsfCallerSession) Symbol() (string, error) {
	return _Hsf.Contract.Symbol(&_Hsf.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Hsf *HsfCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Hsf.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Hsf *HsfSession) TotalSupply() (*big.Int, error) {
	return _Hsf.Contract.TotalSupply(&_Hsf.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Hsf *HsfCallerSession) TotalSupply() (*big.Int, error) {
	return _Hsf.Contract.TotalSupply(&_Hsf.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address _spender, uint256 _value) returns(bool success)
func (_Hsf *HsfTransactor) Approve(opts *bind.TransactOpts, _spender common.Address, _value *big.Int) (*types.Transaction, error) {
	return _Hsf.contract.Transact(opts, "approve", _spender, _value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address _spender, uint256 _value) returns(bool success)
func (_Hsf *HsfSession) Approve(_spender common.Address, _value *big.Int) (*types.Transaction, error) {
	return _Hsf.Contract.Approve(&_Hsf.TransactOpts, _spender, _value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address _spender, uint256 _value) returns(bool success)
func (_Hsf *HsfTransactorSession) Approve(_spender common.Address, _value *big.Int) (*types.Transaction, error) {
	return _Hsf.Contract.Approve(&_Hsf.TransactOpts, _spender, _value)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 amount) returns()
func (_Hsf *HsfTransactor) Burn(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _Hsf.contract.Transact(opts, "burn", amount)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 amount) returns()
func (_Hsf *HsfSession) Burn(amount *big.Int) (*types.Transaction, error) {
	return _Hsf.Contract.Burn(&_Hsf.TransactOpts, amount)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 amount) returns()
func (_Hsf *HsfTransactorSession) Burn(amount *big.Int) (*types.Transaction, error) {
	return _Hsf.Contract.Burn(&_Hsf.TransactOpts, amount)
}

// BurnAll is a paid mutator transaction binding the contract method 0x9975038c.
//
// Solidity: function burnAll() returns()
func (_Hsf *HsfTransactor) BurnAll(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Hsf.contract.Transact(opts, "burnAll")
}

// BurnAll is a paid mutator transaction binding the contract method 0x9975038c.
//
// Solidity: function burnAll() returns()
func (_Hsf *HsfSession) BurnAll() (*types.Transaction, error) {
	return _Hsf.Contract.BurnAll(&_Hsf.TransactOpts)
}

// BurnAll is a paid mutator transaction binding the contract method 0x9975038c.
//
// Solidity: function burnAll() returns()
func (_Hsf *HsfTransactorSession) BurnAll() (*types.Transaction, error) {
	return _Hsf.Contract.BurnAll(&_Hsf.TransactOpts)
}

// BurnFrom is a paid mutator transaction binding the contract method 0x79cc6790.
//
// Solidity: function burnFrom(address account, uint256 amount) returns()
func (_Hsf *HsfTransactor) BurnFrom(opts *bind.TransactOpts, account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Hsf.contract.Transact(opts, "burnFrom", account, amount)
}

// BurnFrom is a paid mutator transaction binding the contract method 0x79cc6790.
//
// Solidity: function burnFrom(address account, uint256 amount) returns()
func (_Hsf *HsfSession) BurnFrom(account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Hsf.Contract.BurnFrom(&_Hsf.TransactOpts, account, amount)
}

// BurnFrom is a paid mutator transaction binding the contract method 0x79cc6790.
//
// Solidity: function burnFrom(address account, uint256 amount) returns()
func (_Hsf *HsfTransactorSession) BurnFrom(account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Hsf.Contract.BurnFrom(&_Hsf.TransactOpts, account, amount)
}

// InitCoin is a paid mutator transaction binding the contract method 0xf50c3f4c.
//
// Solidity: function initCoin(uint256 initialSupply, address holder) returns()
func (_Hsf *HsfTransactor) InitCoin(opts *bind.TransactOpts, initialSupply *big.Int, holder common.Address) (*types.Transaction, error) {
	return _Hsf.contract.Transact(opts, "initCoin", initialSupply, holder)
}

// InitCoin is a paid mutator transaction binding the contract method 0xf50c3f4c.
//
// Solidity: function initCoin(uint256 initialSupply, address holder) returns()
func (_Hsf *HsfSession) InitCoin(initialSupply *big.Int, holder common.Address) (*types.Transaction, error) {
	return _Hsf.Contract.InitCoin(&_Hsf.TransactOpts, initialSupply, holder)
}

// InitCoin is a paid mutator transaction binding the contract method 0xf50c3f4c.
//
// Solidity: function initCoin(uint256 initialSupply, address holder) returns()
func (_Hsf *HsfTransactorSession) InitCoin(initialSupply *big.Int, holder common.Address) (*types.Transaction, error) {
	return _Hsf.Contract.InitCoin(&_Hsf.TransactOpts, initialSupply, holder)
}

// PlaceAndTransfer is a paid mutator transaction binding the contract method 0x528d0f49.
//
// Solidity: function placeAndTransfer(address manager, uint256 _class, uint256 _value) payable returns(bool success)
func (_Hsf *HsfTransactor) PlaceAndTransfer(opts *bind.TransactOpts, manager common.Address, _class *big.Int, _value *big.Int) (*types.Transaction, error) {
	return _Hsf.contract.Transact(opts, "placeAndTransfer", manager, _class, _value)
}

// PlaceAndTransfer is a paid mutator transaction binding the contract method 0x528d0f49.
//
// Solidity: function placeAndTransfer(address manager, uint256 _class, uint256 _value) payable returns(bool success)
func (_Hsf *HsfSession) PlaceAndTransfer(manager common.Address, _class *big.Int, _value *big.Int) (*types.Transaction, error) {
	return _Hsf.Contract.PlaceAndTransfer(&_Hsf.TransactOpts, manager, _class, _value)
}

// PlaceAndTransfer is a paid mutator transaction binding the contract method 0x528d0f49.
//
// Solidity: function placeAndTransfer(address manager, uint256 _class, uint256 _value) payable returns(bool success)
func (_Hsf *HsfTransactorSession) PlaceAndTransfer(manager common.Address, _class *big.Int, _value *big.Int) (*types.Transaction, error) {
	return _Hsf.Contract.PlaceAndTransfer(&_Hsf.TransactOpts, manager, _class, _value)
}

// RemoveAndTransfer is a paid mutator transaction binding the contract method 0x95ef9ca3.
//
// Solidity: function removeAndTransfer(address manager, uint256 _class, uint256 _value) payable returns(bool success)
func (_Hsf *HsfTransactor) RemoveAndTransfer(opts *bind.TransactOpts, manager common.Address, _class *big.Int, _value *big.Int) (*types.Transaction, error) {
	return _Hsf.contract.Transact(opts, "removeAndTransfer", manager, _class, _value)
}

// RemoveAndTransfer is a paid mutator transaction binding the contract method 0x95ef9ca3.
//
// Solidity: function removeAndTransfer(address manager, uint256 _class, uint256 _value) payable returns(bool success)
func (_Hsf *HsfSession) RemoveAndTransfer(manager common.Address, _class *big.Int, _value *big.Int) (*types.Transaction, error) {
	return _Hsf.Contract.RemoveAndTransfer(&_Hsf.TransactOpts, manager, _class, _value)
}

// RemoveAndTransfer is a paid mutator transaction binding the contract method 0x95ef9ca3.
//
// Solidity: function removeAndTransfer(address manager, uint256 _class, uint256 _value) payable returns(bool success)
func (_Hsf *HsfTransactorSession) RemoveAndTransfer(manager common.Address, _class *big.Int, _value *big.Int) (*types.Transaction, error) {
	return _Hsf.Contract.RemoveAndTransfer(&_Hsf.TransactOpts, manager, _class, _value)
}

// Seo is a paid mutator transaction binding the contract method 0x70238ecd.
//
// Solidity: function seo(address _to, uint256 _value) returns()
func (_Hsf *HsfTransactor) Seo(opts *bind.TransactOpts, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _Hsf.contract.Transact(opts, "seo", _to, _value)
}

// Seo is a paid mutator transaction binding the contract method 0x70238ecd.
//
// Solidity: function seo(address _to, uint256 _value) returns()
func (_Hsf *HsfSession) Seo(_to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _Hsf.Contract.Seo(&_Hsf.TransactOpts, _to, _value)
}

// Seo is a paid mutator transaction binding the contract method 0x70238ecd.
//
// Solidity: function seo(address _to, uint256 _value) returns()
func (_Hsf *HsfTransactorSession) Seo(_to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _Hsf.Contract.Seo(&_Hsf.TransactOpts, _to, _value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address _to, uint256 _value) payable returns(bool success)
func (_Hsf *HsfTransactor) Transfer(opts *bind.TransactOpts, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _Hsf.contract.Transact(opts, "transfer", _to, _value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address _to, uint256 _value) payable returns(bool success)
func (_Hsf *HsfSession) Transfer(_to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _Hsf.Contract.Transfer(&_Hsf.TransactOpts, _to, _value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address _to, uint256 _value) payable returns(bool success)
func (_Hsf *HsfTransactorSession) Transfer(_to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _Hsf.Contract.Transfer(&_Hsf.TransactOpts, _to, _value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address _from, address _to, uint256 _value) payable returns(bool success)
func (_Hsf *HsfTransactor) TransferFrom(opts *bind.TransactOpts, _from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _Hsf.contract.Transact(opts, "transferFrom", _from, _to, _value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address _from, address _to, uint256 _value) payable returns(bool success)
func (_Hsf *HsfSession) TransferFrom(_from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _Hsf.Contract.TransferFrom(&_Hsf.TransactOpts, _from, _to, _value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address _from, address _to, uint256 _value) payable returns(bool success)
func (_Hsf *HsfTransactorSession) TransferFrom(_from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _Hsf.Contract.TransferFrom(&_Hsf.TransactOpts, _from, _to, _value)
}

// HsfApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the Hsf contract.
type HsfApprovalIterator struct {
	Event *HsfApproval // Event containing the contract specifics and raw log

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
func (it *HsfApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(HsfApproval)
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
		it.Event = new(HsfApproval)
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
func (it *HsfApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *HsfApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// HsfApproval represents a Approval event raised by the Hsf contract.
type HsfApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address owner, address spender, uint256 value)
func (_Hsf *HsfFilterer) FilterApproval(opts *bind.FilterOpts) (*HsfApprovalIterator, error) {

	logs, sub, err := _Hsf.contract.FilterLogs(opts, "Approval")
	if err != nil {
		return nil, err
	}
	return &HsfApprovalIterator{contract: _Hsf.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address owner, address spender, uint256 value)
func (_Hsf *HsfFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *HsfApproval) (event.Subscription, error) {

	logs, sub, err := _Hsf.contract.WatchLogs(opts, "Approval")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(HsfApproval)
				if err := _Hsf.contract.UnpackLog(event, "Approval", log); err != nil {
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
func (_Hsf *HsfFilterer) ParseApproval(log types.Log) (*HsfApproval, error) {
	event := new(HsfApproval)
	if err := _Hsf.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// HsfTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the Hsf contract.
type HsfTransferIterator struct {
	Event *HsfTransfer // Event containing the contract specifics and raw log

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
func (it *HsfTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(HsfTransfer)
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
		it.Event = new(HsfTransfer)
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
func (it *HsfTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *HsfTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// HsfTransfer represents a Transfer event raised by the Hsf contract.
type HsfTransfer struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address owner, address spender, uint256 value)
func (_Hsf *HsfFilterer) FilterTransfer(opts *bind.FilterOpts) (*HsfTransferIterator, error) {

	logs, sub, err := _Hsf.contract.FilterLogs(opts, "Transfer")
	if err != nil {
		return nil, err
	}
	return &HsfTransferIterator{contract: _Hsf.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address owner, address spender, uint256 value)
func (_Hsf *HsfFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *HsfTransfer) (event.Subscription, error) {

	logs, sub, err := _Hsf.contract.WatchLogs(opts, "Transfer")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(HsfTransfer)
				if err := _Hsf.contract.UnpackLog(event, "Transfer", log); err != nil {
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
func (_Hsf *HsfFilterer) ParseTransfer(log types.Log) (*HsfTransfer, error) {
	event := new(HsfTransfer)
	if err := _Hsf.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
