package investors

import (
	"context"
	"crypto/ecdsa"
	hsf "entysquare/entycats/mockhsf"
	usdt "entysquare/entycats/mockusdt"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
	"math/big"
	"testing"
)

func TestDeploy(t *testing.T) {
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

	address, tx, instance, err := hsf.DeployHsf(auth, client)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(address.Hex())   // 0x147B8eb97fD247D06C4006D269c90C1908Fb5D54 部署的合约地址
	fmt.Println(tx.Hash().Hex()) // 0xdae8ba5444eefdc99f4d45cd0c4f24056cba6a02cefbf78066ef9f4188ff7dc0 交易hash

	_ = instance
}

func TestBalanceRead(t *testing.T) {
	client, err := ethclient.Dial("https://rinkeby.infura.io/v3/6c81fb1b66804f0698d49f2ec242afc9")
	if err != nil {
		log.Fatal(err)
	}

	address := common.HexToAddress("0xac3df53659f083e3002a76d06c649b0fe772035b")
	selfAddress := common.HexToAddress("0x5a07BC15761Ee0dCB3D3e2e61960D6CDbBAF1EF2")

	instance, err := usdt.NewUsdt(address, client)
	if err != nil {
		log.Fatal(err)
	}

	balance, err := instance.BalanceOf(nil, selfAddress)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(balance) // "951"
}
