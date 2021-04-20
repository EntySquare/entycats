package investors

import (
	"context"
	"crypto/ecdsa"
	investors_usdt "entysquare/entycats/mockusdt"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
	"math/big"
	"strings"
	"testing"
)

func TestRead(t *testing.T) {
	client, err := ethclient.Dial("https://rinkeby.infura.io")
	if err != nil {
		log.Fatal(err)
	}

	privateKey, err := crypto.HexToECDSA("c0feb08a997fc0c75c1a008a67a5209fb763ee53984c5126af35341570979213")
	if err != nil {
		log.Fatal(err)
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		log.Fatal(err)
	}

	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	auth := bind.NewKeyedTransactor(privateKey)
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)     // in wei
	auth.GasLimit = uint64(300000) // in units
	auth.GasPrice = gasPrice

	//input := "1.0"
	//address, tx, instance, err := store.DeployStore(auth, client, input)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//
	//fmt.Println(address.Hex())   // 0x147B8eb97fD247D06C4006D269c90C1908Fb5D54
	//fmt.Println(tx.Hash().Hex()) // 0xdae8ba5444eefdc99f4d45cd0c4f24056cba6a02cefbf78066ef9f4188ff7dc0

	//_ = instance
}

func Test5(t *testing.T) { //总高度
	url := "https://rinkeby.infura.io/v3/6c81fb1b66804f0698d49f2ec242afc9"
	//code:="ETH"
	client, err := ethclient.Dial(url)
	fmt.Println(err)
	//sync, err := client.SyncProgress(context.TODO())
	//fmt.Println(sync.CurrentBlock)
	tx, _, err := client.TransactionByHash(context.TODO(), common.HexToHash("0x1a944913aab5420504ead95438acfbb1274aa27f933581231816d3ffbc61e980"))

	if err != nil {
		t.Error(err)
	}

	d := tx.Data()

	method := d[:4]
	//addressMethod := d[4:]
	//address := addressMethod[:32]
	//amount := addressMethod[32:]

	parsed, err := abi.JSON(strings.NewReader(investors_usdt.InvestorsUsdtABI))
	if err != nil {
		panic(err)
	}

	m, err := parsed.MethodById(method)
	if err != nil {
		panic(err)
	}
	fmt.Println(m)

	i, err := parsed.Methods["transfer"].Inputs.Unpack(d[4:])
	addrInterfi := i[0]
	valueInterfi := i[1]
	address := (addrInterfi).(common.Address)
	value := (valueInterfi).(*big.Int)

	fmt.Println(address.Hex())
	fmt.Println(value.String())
	//common.HexToAddress(addrInterfi)

	//if err != nil {
	//	panic(err)
	//}
	//
	//// Value 参数解码
	//var addrwant common.Address
	//if err := parsed.Methods["Value"].Inputs.Unpack(&addrwant, valueInput[4:]); err != nil {
	//	panic(err)
	//}
	//fmt.Println("should equals", addrwant == address)
	//
	//// Value 返回值解码
	//var balance *big.Int
	//var returns = common.Hex2Bytes("0000000000000000000000000000000000000000000000000000000005f5e100")
	//if err := parsed.Unpack(&balance, "Value", returns); err != nil {
	//	panic(err)
	//}
	//fmt.Println("Value 返回值", balance)
	//
	//

	//i,err := parsed.Unpack("transfer",d)
	fmt.Println(i)

	//var balance *big.Int
	//if err := parsed.Unpack(&balance, "Value", returns); err != nil {
	//	panic(err)
	//}
	//
	//fmt.Println(method)
	//fmt.Println(addr.String())
	//fmt.Println(amountStr)
	//
	//
	//
	//
	//res := hexutil.Encode(d)
	//fmt.Println(res)

	//bytes := crypto.Keccak256([]byte("transfer(address,uint256)"))
	//bytes := crypto.Keccak256([]byte("placeAndTransfer(address,uint256,uint256)"))
	//bytes2 := bytes[:8]

	//fmt.Println(hexutil.Encode(bytes2))

}
