package lib

import (
	"context"
	"crypto/ecdsa"
	hsf "entysquare/entycats/mockhsf"
	usdt "entysquare/entycats/mockusdt"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
	"math/big"
	"time"
)

//var ha common.Address
//var ua common.Address
//var ia common.Address
var rawURL string
var hexKey string

func contractSetter(ru string, pk string) {
	rawURL = ru
	hexKey = pk

}

func Init(ru string, pk string, so int64, itf bool, it string) {
	contractSetter(ru, pk)
	hsfAddress, _ := deploy(1, so, itf, it)
	//ha = hsfAddress
	time.Sleep(time.Second * 45)
	fmt.Println(hsfAddress.Hex())
	usdtAddress, _ := deploy(2, so, itf, it)
	//ua = usdtAddress
	fmt.Println(usdtAddress.Hex())
	time.Sleep(time.Second * 45)
	investorsAddress, _ := deploy(3, so, itf, it)
	//ia = investorsAddress
	fmt.Println(investorsAddress.Hex())
	client := getClient()
	//time.Sleep(time.Second * 45)
	//hsfInstance, hsferr := hsf.NewHsf(hsfAddress, client)
	//if hsferr != nil {
	//	log.Fatal(hsferr)
	//}
	//time.Sleep(time.Second * 30)
	//hsfInstance.InitCoin(getAuth(3000000, client), big.NewInt(10000), common.HexToAddress(ma))
	//time.Sleep(time.Second * 30)
	//usdtInstance, usdterr := usdt.NewUsdt(usdtAddress, client)
	//if usdterr != nil {
	//	log.Fatal(usdterr)
	//}
	//time.Sleep(time.Second * 30)
	//usdtInstance.InitCoin(getAuth(3000000, client), big.NewInt(10000), common.HexToAddress(ma))
	time.Sleep(time.Second * 30)
	investorsInstance, err := NewInvestors(investorsAddress, client)
	if err != nil {
		log.Fatal(err)
	}
	time.Sleep(time.Second * 30)
	investorsInstance.InitToken(getAuth(5000000, client), hsfAddress, usdtAddress)
	time.Sleep(time.Second * 30)
	//placeBet(10,1,getAuth(3000000,client),hsfAddress,hsfInstance)
	//placeBet(20,2,getAuth(3000000,client),hsfAddress,hsfInstance)
}

func PAR(class int64, value int64, tp string, ha string, ia string, pk string) {
	_ia := common.HexToAddress(ia)
	_ha := common.HexToAddress(ha)
	client := getClient()
	auth := getAuthByPrivateKey(3000000, client, pk)
	//	auth := getAuthByPrivateKey(3000000,client,"")
	hsfInstance, err := hsf.NewHsf(_ha, client)
	if err != nil {
		log.Fatal(err)
	}
	if tp == "place" {
		_, err = hsfInstance.PlaceAndTransfer(auth, _ia, big.NewInt(class), big.NewInt(value))
		if err != nil {
			log.Fatal(err)
		}
	}
	if tp == "remove" {
		_, err = hsfInstance.RemoveAndTransfer(auth, _ia, big.NewInt(class), big.NewInt(value))
		if err != nil {
			log.Fatal(err)
		}
	}
	//	placeBet(10,1,getAuth(3000000,client),hsfAddress,hsfInstance)
	//	placeBet(20,2,getAuth(3000000,client),hsfAddress,hsfInstance)
	//	removeBet(10,1,getAuth(3000000,client),hsfAddress,hsfInstance)
}
func Trans(_ia common.Address, _ua common.Address, _value int64) {
	client := getClient()
	auth := getAuth(300000, client)
	instance, err := usdt.NewUsdt(_ua, client)
	if err != nil {
		log.Fatal(err)
	}

	key := [32]byte{}
	value := [32]byte{}
	copy(key[:], []byte("foo"))
	copy(value[:], []byte("bar"))
	tx, err := instance.Transfer(auth, _ia, big.NewInt(_value))
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("tx sent: %s", tx.Hash().Hex()) // tx sent:
}
func ShareOut(bounces int64, ia string, pk string) {
	_ia := common.HexToAddress(ia)
	client := getClient()
	auth := getAuthByPrivateKey(3000000, client, pk)
	investorsInstance, err := NewInvestors(_ia, client)
	if err != nil {
		println(err)
	}
	_, err = investorsInstance.ShareOut(auth, big.NewInt(bounces))
	if err != nil {
		println(err)
	}
}
func Bal(_ia common.Address, _ha common.Address) {
	client := getClient()
	hsfInstance, err := hsf.NewHsf(_ha, client)
	if err != nil {
		println(err)
	}
	balance, err := hsfInstance.BalanceOf(nil, _ia)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(balance) // "951"
}
func deploy(x int, startOver int64, institutionflag bool, institution string) (common.Address, *types.Transaction) {
	client := getClient() //测试网络
	auth := getAuth(3000000, client)
	if x == 1 {
		//		address, tx, instance, err := hsf.DeployHsf(auth, client)
		address, tx, _, err := hsf.DeployHsf(auth, client)
		if err != nil {
			log.Fatal(err)
		}
		//		instance.InitCoin(nil,big.NewInt(100),mine)
		return address, tx
	}
	if x == 2 {
		//		address, tx, instance, err := usdt.DeployUsdt(auth, client)
		address, tx, _, err := usdt.DeployUsdt(auth, client)
		if err != nil {
			log.Fatal(err)
		}
		//		instance.InitCoin(nil,big.NewInt(100),mine)
		return address, tx
	}
	if x == 3 {
		address, tx, _, err := DeployInvestors(auth, client, big.NewInt(startOver), institutionflag, common.HexToAddress(institution))
		if err != nil {
			log.Fatal(err)
		}
		return address, tx
	}
	return common.HexToAddress("none"), nil
}

func getClient() *ethclient.Client {
	//client, err := ethclient.Dial("https://rinkeby.infura.io/v3/6c81fb1b66804f0698d49f2ec242afc9") //测试网络
	client, err := ethclient.Dial(rawURL) //测试网络
	if err != nil {
		log.Fatal(err)
	}
	return client
}
func getAuth(gaslimit int, client *ethclient.Client) *bind.TransactOpts {
	//privateKey, err := crypto.HexToECDSA("a30eb8e38d72915cd0dfd399f4f110838c4417f0a18b617d946b0d0636a5a5ae") //私钥
	privateKey, err := crypto.HexToECDSA(hexKey) //私钥
	if err != nil {
		log.Fatal(err)
	}
	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("error casting public key to ECDSA")
	}
	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		log.Fatal(err)
	}
	gasPrice, err := client.SuggestGasPrice(context.Background()) //燃料价格
	if err != nil {
		log.Fatal(err)
	}
	auth := bind.NewKeyedTransactor(privateKey)
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)       // in wei
	auth.GasLimit = uint64(gaslimit) // in units 燃料限制
	auth.GasPrice = gasPrice
	return auth
}
func getAuthByPrivateKey(gaslimit int, client *ethclient.Client, privateString string) *bind.TransactOpts {
	privateKey, err := crypto.HexToECDSA(privateString) //私钥
	if err != nil {
		log.Fatal(err)
	}
	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("error casting public key to ECDSA")
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		log.Fatal(err)
	}
	gasPrice, err := client.SuggestGasPrice(context.Background()) //燃料价格
	if err != nil {
		log.Fatal(err)
	}
	auth := bind.NewKeyedTransactor(privateKey)
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)       // in wei
	auth.GasLimit = uint64(gaslimit) // in units 燃料限制
	auth.GasPrice = gasPrice
	return auth
}
