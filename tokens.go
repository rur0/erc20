package erc20

import (
	"bytes"
	"fmt"
	"math/big"
	"strings"
	"sync"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

var (
	//tokens = make(map[common.Address]*Token)
	tokens = sync.Map{}
	//tokenSpecs = make(map[common.Address]TokenSpec)
	tokenSpecs = sync.Map{}
)

type TokenSpec struct {
	Name        string         `json:"name"`
	Symbol      string         `json:"symbol"`
	Decimals    uint8          `json:"decimals"`
	TotalSupply *big.Int       `json:"total_supply"`
	Address     common.Address `json:"address"`
}

func (ts TokenSpec) String() string {
	return fmt.Sprintf("Name: %s \nSymbol: %s \nDecimals: %d \nTotalSupply: %s \nAddress: %s",
		ts.Name, ts.Symbol, ts.Decimals,
		new(big.Int).Div(ts.TotalSupply, new(big.Int).Exp(big.NewInt(10), big.NewInt(int64(ts.Decimals)), nil)),
		ts.Address,
	)
}

func (ts TokenSpec) Addr() string {
	return strings.ToLower(ts.Address.String())
}

func (ts TokenSpec) Equal(ts2 *TokenSpec) bool {
	return bytes.Equal(ts.Address.Bytes(), ts2.Address.Bytes())
}

type TokenSpecs []TokenSpec

func (tss TokenSpecs) Includes(ts TokenSpec) bool {
	return tss.Index(ts) >= 0
}

func (tss TokenSpecs) Index(ts TokenSpec) int {
	for i, Ts := range tss {
		if bytes.Equal(Ts.Address.Bytes(), ts.Address.Bytes()) {
			return i
		}
	}
	return -1
}

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
	//	tk, ok := tokens[contract]
	_tk, ok := tokens.Load(contract)
	if ok {
		return _tk.(*Token), nil
	}

	tk, err := NewToken(contract, client)
	if err != nil {
		return nil, err
	}

	tokens.Store(contract, tk)

	return tk, nil
}

func GetTokenSpec(contract common.Address, client *ethclient.Client) (*TokenSpec, error) {
	spec, ok := tokenSpecs.Load(contract)
	if ok {
		return spec.(*TokenSpec), nil
	}
	tkSpec := new(TokenSpec)

	tk, err := GetToken(contract, client)
	if err != nil {
		return nil, err
	}

	tkSpec.Name, err = tk.Name(nil)
	if err != nil {
		return nil, err
	}

	// todo allow parse bytes32 symbol
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

	tokenSpecs.Store(contract, tkSpec)

	return tkSpec, nil
}

func GetTokenSpecs(client *ethclient.Client, contracts ...common.Address) (TokenSpecs, error) {
	specs := TokenSpecs{}

	for _, addr := range contracts {
		spec, err := GetTokenSpec(addr, client)
		if err != nil {
			return nil, fmt.Errorf("%s: %s", addr, err)
		}
		specs = append(specs, *spec)
	}

	return specs, nil
}
