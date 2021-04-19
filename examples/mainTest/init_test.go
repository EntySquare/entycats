package mainTest

import (
	"context"
	"crypto/ecdsa"
	investors "entysquare/entycats/lib"
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
	"testing"
	"time"
)

const myAddress = "0x524467bCc0714a80702E1B6C607017d12799AB18"
const hsfoAddress = "0x2fb27c45Bc655757a3832DdaDD7ca3cbF492330a"
const usdtoAddress = "0xBdF42596CF3ef5ADadB9F735c1B3e9e9f5a342fF"
const inAddress = "0xAcF685c2350DB4CB3726E50e38399bD64b516476"

func TestInit(t *testing.T) {
	hsfAddress, _ := deploy(1)
	time.Sleep(time.Second * 45)
	fmt.Println(hsfAddress.Hex())
	usdtAddress, _ := deploy(2)
	fmt.Println(usdtAddress.Hex())
	time.Sleep(time.Second * 45)
	investorsAddress, _ := deploy(3)
	fmt.Println(investorsAddress.Hex())
	client := getClient()
	time.Sleep(time.Second * 45)
	hsfInstance, hsferr := hsf.NewHsf(hsfAddress, client)
	if hsferr != nil {
		log.Fatal(hsferr)
	}
	time.Sleep(time.Second * 30)
	hsfInstance.InitCoin(getAuth(3000000, client), big.NewInt(100), common.HexToAddress(myAddress))
	time.Sleep(time.Second * 30)
	usdtInstance, usdterr := usdt.NewUsdt(usdtAddress, client)
	if usdterr != nil {
		log.Fatal(usdterr)
	}
	time.Sleep(time.Second * 30)
	usdtInstance.InitCoin(getAuth(3000000, client), big.NewInt(100), common.HexToAddress(myAddress))
	time.Sleep(time.Second * 30)
	investorsInstance, err := investors.NewInvestors(investorsAddress, client)
	if err != nil {
		log.Fatal(err)
	}
	time.Sleep(time.Second * 30)
	investorsInstance.InitToken(getAuth(5000000, client), hsfAddress, usdtAddress)
	//placeBet(10,1,getAuth(3000000,client),hsfAddress,hsfInstance)
	//placeBet(20,2,getAuth(3000000,client),hsfAddress,hsfInstance)
}

func TestPAR(t *testing.T) {
	investorsAddress := common.HexToAddress(inAddress)
	client := getClient()
	hsfInstance, err := hsf.NewHsf(investorsAddress, client)
	if err != nil {
		println(err)
	}
	_, err = hsfInstance.PlaceAndTransfer(getAuth(3000000, client), investorsAddress, big.NewInt(1), big.NewInt(1))
	if err != nil {
		println(err)
	}
	//	placeBet(10,1,getAuth(3000000,client),hsfAddress,hsfInstance)
	//	placeBet(20,2,getAuth(3000000,client),hsfAddress,hsfInstance)
	//	removeBet(10,1,getAuth(3000000,client),hsfAddress,hsfInstance)
}

func placeBet(value int64, _class int64, opts *bind.TransactOpts, manager common.Address, instance *hsf.Hsf) {
	instance.PlaceAndTransfer(opts, manager, big.NewInt(value), big.NewInt(_class))
}
func removeBet(value int64, _class int64, opts *bind.TransactOpts, manager common.Address, instance *hsf.Hsf) {
	instance.RemoveAndTransfer(opts, manager, big.NewInt(value), big.NewInt(_class))
}
func deploy(x int) (common.Address, *types.Transaction) {
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
		address, tx, _, err := investors.DeployInvestors(auth, client, big.NewInt(10000), false, common.HexToAddress(""))
		if err != nil {
			log.Fatal(err)
		}
		return address, tx
	}
	return common.HexToAddress("none"), nil
}

func getClient() *ethclient.Client {
	client, err := ethclient.Dial("https://rinkeby.infura.io/v3/6c81fb1b66804f0698d49f2ec242afc9") //测试网络
	if err != nil {
		log.Fatal(err)
	}
	return client
}
func getAuth(gaslimit int, client *ethclient.Client) *bind.TransactOpts {
	privateKey, err := crypto.HexToECDSA("a30eb8e38d72915cd0dfd399f4f110838c4417f0a18b617d946b0d0636a5a5ae") //私钥
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
