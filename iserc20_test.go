package erc20

import (
	"log"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

var (
	BSC_NODE = "ws://148.251.91.76:8546"
	ETH_NODE = "ws://157.90.5.252:3334"
)

func TestErc20(t *testing.T) {
	client, err := ethclient.Dial(ETH_NODE)
	if err != nil {
		log.Fatal(err)
	}
	addr := common.HexToAddress("0xfA5047c9c78B8877af97BDcb85Db743fD7313d4a")
	isErc20, err := IsErc20(client, addr)
	if err != nil {
		isErc20 = false
	}
	log.Println(isErc20)
}
