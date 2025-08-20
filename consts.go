package erc20

import (
	"github.com/ethereum/go-ethereum/crypto"
)

var (
	LogTransferSig     = []byte("Transfer(address,address,uint256)")
	LogApprovalSig     = []byte("Approval(address,address,uint256)")
	LogTransferSigHash = crypto.Keccak256Hash(LogTransferSig)
	LogApprovalSigHash = crypto.Keccak256Hash(LogApprovalSig)
)
