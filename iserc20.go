package erc20

import (
	"context"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

func IsErc20(client *ethclient.Client, addr common.Address) (bool, error) {
	ctx := context.Background()
	bytecode, err := client.CodeAt(ctx, addr, nil) // nil is latest block
	if err != nil {
		return false, err
	}

	isContract := len(bytecode) > 0

	if !isContract {
		return false, nil
	}

	_, err = GetTokenSpec(addr, client)
	if err != nil {
		return false, err
	}

	return true, nil
}
