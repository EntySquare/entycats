package usdt

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
	"math/big"
	"testing"
)

func TestBalanceRead(t *testing.T) {
	client, err := ethclient.Dial("https://rinkeby.infura.io/v3/6c81fb1b66804f0698d49f2ec242afc9")
	if err != nil {
		log.Fatal(err)
	}

	address := common.HexToAddress("0xac3df53659f083e3002a76d06c649b0fe772035b")
	selfAddress := common.HexToAddress("0x5a07BC15761Ee0dCB3D3e2e61960D6CDbBAF1EF2")

	instance, err := NewUsdt(address, client)
	if err != nil {
		log.Fatal(err)
	}

	balance, err := instance.BalanceOf(nil, selfAddress)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(balance) // "951"
}

func TestTransfer(t *testing.T) {
	client, err := ethclient.Dial("https://rinkeby.infura.io/v3/6c81fb1b66804f0698d49f2ec242afc9")
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

	address := common.HexToAddress("0xac3df53659f083e3002a76d06c649b0fe772035b")
	toaddress := common.HexToAddress("0x4c6E666D6d63BB4eA74E0e6a0faF1ACcC7AdBf90")
	instance, err := NewUsdt(address, client)
	if err != nil {
		log.Fatal(err)
	}

	key := [32]byte{}
	value := [32]byte{}
	copy(key[:], []byte("foo"))
	copy(value[:], []byte("bar"))

	i := new(big.Int)
	_, err = fmt.Sscan("9518866985326561560", i)
	tx, err := instance.Transfer(auth, toaddress, i)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("tx sent: %s", tx.Hash().Hex()) // tx sent:
}
