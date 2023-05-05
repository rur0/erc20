package erc20

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

func BalanceOf(client *ethclient.Client, token, holder common.Address) (*big.Int, error) {
	tk, err := NewToken(token, client)
	if err != nil {
		return nil, err
	}
	amount, err := tk.BalanceOf(nil, holder)
	if err != nil {
		return nil, err
	}

	return amount, nil
}
