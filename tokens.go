package erc20

import (
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

var (
	tokens     = make(map[common.Address]*Token)
	tokenSpecs = make(map[common.Address]TokenSpec)
)

type TokenSpec struct {
	Symbol      string
	Decimals    uint8
	TotalSupply *big.Int
	Address     common.Address
}

func (ts TokenSpec) Addr() string {
	return strings.ToLower(ts.Address.String())
}

func GetToken(contract common.Address, client *ethclient.Client) (*Token, error) {
	var err error
	tk, ok := tokens[contract]
	if ok {
		return tk, nil
	}

	tk, err = NewToken(contract, client)
	if err != nil {
		return nil, err
	}
	tokens[contract] = tk

	return tk, nil
}

func GetTokenSpec(contract common.Address, client *ethclient.Client) (*TokenSpec, error) {
	tkSpec, ok := tokenSpecs[contract]
	if ok {
		return &tkSpec, nil
	}

	tk, err := GetToken(contract, client)
	if err != nil {
		return nil, err
	}

	tkSpec.Symbol, err = tk.Symbol(nil)
	if err != nil {
		return nil, err
	}

	tkSpec.Decimals, err = tk.Decimals(nil)
	if err != nil {
		return nil, err
	}

	tkSpec.TotalSupply, err = tk.TotalSupply(nil)
	if err != nil {
		return nil, err
	}

	tkSpec.Address = contract

	tokenSpecs[contract] = tkSpec

	return &tkSpec, nil
}
