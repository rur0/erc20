package erc20

import (
	"fmt"
	"log"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

func TestTokenSpec(t *testing.T) {
	client, err := ethclient.Dial(ETH_NODE)
	if err != nil {
		log.Fatal(err)
	}

	tkSpec, err := GetTokenSpec(common.HexToAddress("0xb1e9157c2fdcc5a856c8da8b2d89b6c32b3c1229"), client)
	if err != nil {
		t.Fatal(err)
	}

	fmt.Println(tkSpec)
}

func TestCompareAddr(t *testing.T) {
	tkA, tkB := common.HexToAddress("0xb1e9157c2fdcc5a856c8da8b2d89b6c32b3c1229"), common.HexToAddress("0xb1e9157c2fdcc5a856c8da8b2d89b6c32b3c1229")
	fmt.Println(tkA == tkB)
}
