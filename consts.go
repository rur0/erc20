package erc20

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

// useful constants
var (
	LogTransferSig     = []byte("Transfer(address,address,uint256)")
	LogApprovalSig     = []byte("Approval(address,address,uint256)")
	LogSwapSig         = []byte("Swap(address,uint256,uint256,uint256,uint256,address)")
	LogPairCreatedSig  = []byte("PairCreated(address,address,address,uint256)") // uniswap pool creation
	LogTransferSigHash = crypto.Keccak256Hash(LogTransferSig)
	LogApprovalSigHash = crypto.Keccak256Hash(LogApprovalSig)
	LogSwapSigHash     = crypto.Keccak256Hash(LogSwapSig)
	LogPairCreatedHash = crypto.Keccak256Hash(LogPairCreatedSig)

	WETHAddr           = common.HexToAddress("0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2")
	UniswapFactoryAddr = common.HexToAddress("0x5c69bee701ef814a2b6a3edd4b1652cb9cc5aa6f")
)
