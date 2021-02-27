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
	LogTransferSigHash = crypto.Keccak256Hash(LogTransferSig)
	LogApprovalSigHash = crypto.Keccak256Hash(LogApprovalSig)
	LogSwapSigHash     = crypto.Keccak256Hash(LogSwapSig)

	WETHAddr = common.HexToAddress("0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2")
)
