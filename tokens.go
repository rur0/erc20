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

type TokenSpecs []TokenSpec

func (ts TokenSpecs) GetAddrs() []common.Address {
	addrs := []common.Address{}
	for _, tokenSpec := range ts {
		addrs = append(addrs, tokenSpec.Address)
	}

	return addrs
}

func (ts TokenSpecs) GetSymbols() []string {
	symbols := []string{}
	for _, tokenSpec := range ts {
		symbols = append(symbols, tokenSpec.Symbol)
	}

	return symbols
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

func GetTokenSpecs(client *ethclient.Client, contracts ...common.Address) (TokenSpecs, error) {
	specs := TokenSpecs{}

	for _, addr := range contracts {
		spec, err := GetTokenSpec(addr, client)
		if err != nil {
			return nil, err
		}
		specs = append(specs, *spec)
	}

	return specs, nil
}
