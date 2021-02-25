package erc20

import (
	"errors"
	"fmt"
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

// Erc20Abi ...
var Erc20Abi abi.ABI

func init() {
	var err error
	Erc20Abi, err = abi.JSON(strings.NewReader(TokenABI))
	if err != nil {
		panic(err)
	}
}

// LogTransfer ...
type LogTransfer struct {
	From   common.Address
	To     common.Address
	Tokens *big.Int
}

// ParseLogTransfer parses a `Transfer` log
func ParseLogTransfer(log *types.Log) (*LogTransfer, error) {
	var logTransfer LogTransfer
	err := Erc20Abi.UnpackIntoInterface(&logTransfer, "Transfer", log.Data)
	if err != nil {
		return nil, err
	}

	if len(log.Topics) < 3 {
		fmt.Println(log.Topics)
		return nil, errors.New("no quantities in transfer")
	}

	logTransfer.From = common.BytesToAddress(log.Topics[1].Bytes())
	logTransfer.To = common.BytesToAddress(log.Topics[2].Bytes())
	return &logTransfer, nil
}

// FindFirstFromLog finds the first occurence of a log from a given address
func FindFirstFromLog(logs []*types.Log, topic common.Hash, from common.Address) *types.Log {
	for _, log := range logs {
		if log.Topics[0] == topic && log.Address == from {
			return log
		}
	}
	return nil
}

// IsAddrInSlice ...
func IsAddrInSlice(addr common.Address, addrs []common.Address) bool {
	for _, Addr := range addrs {
		if Addr == addr {
			return true
		}
	}
	return false
}
