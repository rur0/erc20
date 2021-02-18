#!/bin/bash

# compile to abi
solc --abi erc20.sol --overwrite -o .

# generate golang bindings
abigen --abi=ERC20.abi --pkg=token --out=erc20.go

# rename the package
sed -i 's/package token/package erc20/g' erc20.go