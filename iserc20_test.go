package erc20

import (
	"log"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

var (
	ETH_NODE = "https://eth.llamarpc.com"
)

func TestErc20(t *testing.T) {
	client, err := ethclient.Dial(ETH_NODE)
	if err != nil {
		log.Fatal(err)
	}
	addr := common.HexToAddress("0xfA5047c9c78B8877af97BDcb85Db743fD7313d4a")
	isErc20, err := IsErc20(client, addr)
	if err != nil {
		t.Fatal(err)
	}
	if !isErc20 {
		t.Fatal("expected erc20 is not")
	}
}
