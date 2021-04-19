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
	investorsInstance, err := investors.NewInvestors(investorsAddress, client)
	if err != nil {
		log.Fatal(err)
	}
	investorsInstance.InitToken(getAuth(5000000, client), hsfAddress, usdtAddress)
}
func deploy(x int) (common.Address, *types.Transaction) {
	client, err := ethclient.Dial("https://rinkeby.infura.io/v3/6c81fb1b66804f0698d49f2ec242afc9") //测试网络
	if err != nil {
		log.Fatal(err)
	}

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
	auth.Value = big.NewInt(0)      // in wei
	auth.GasLimit = uint64(3000000) // in units 燃料限制
	auth.GasPrice = gasPrice
	if x == 1 {
		//		address, tx, instance, err := hsf.DeployHsf(auth, client)
		address, tx, _, err := hsf.DeployHsf(auth, client)
		if err != nil {
			log.Fatal(err)
		}
		//        instance.InitCoin(auth,big.NewInt(100),fromAddress)
		return address, tx
	}
	if x == 2 {
		//		address, tx, instance, err := usdt.DeployUsdt(auth, client)
		address, tx, _, err := usdt.DeployUsdt(auth, client)
		if err != nil {
			log.Fatal(err)
		}
		//		instance.InitCoin(auth,big.NewInt(100),fromAddress)
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
