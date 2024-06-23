package main

import (
    "fmt"

)

func main() {

    var rpcProvider = "https://linea-goerli.infura.io/v3/9aa3d95b3bc440fa88ea12eaa4456161"
    web3, err := web3.NewWeb3(rpcProvider)
    if err != nil {
        panic(err)
    }

    blockNumber, err := web3.Eth.GetBlockNumber()
    if err != nil {
        panic(err)
    }

    fmt.Println("Current block number: ", blockNumber)
}
