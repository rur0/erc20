package erc20

import (
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/sirupsen/logrus"
)

var (
	BSC_NODE = "ws://148.251.91.76:8546"
	ETH_NODE = "ws://157.90.5.252:3334"
)

func TestErc20(t *testing.T) {
	client, err := ethclient.Dial(ETH_NODE)
	if err != nil {
		logrus.Fatal(err)
	}
	addr := common.HexToAddress("0xfA5047c9c78B8877af97BDcb85Db743fD7313d4a")
	isErc20, err := IsErc20(client, addr)
	if err != nil {
		isErc20 = false
	}
	logrus.Info(isErc20)
}
