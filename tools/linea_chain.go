package main

import (
	"context"
	"fmt"
	"log"
	"math/big"
	"github.com/ethereum/go-ethereum/params"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/common"
)




func weiToEther(wei *big.Int) *big.Float {
    return new(big.Float).Quo(new(big.Float).SetInt(wei), big.NewFloat(params.Ether))
}
func main() {
	client, err := ethclient.Dial("https://linea-sepolia.infura.io/v3/9aa3d95b3bc440fa88ea12eaa4456161")
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println("Ja, alles gut!")
	}


	ctx := context.Background()

	/* Get the latest block option 1 */
	header, err := client.HeaderByNumber(ctx, nil)
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println(header.Number.String())
	}

	/* Get the latest block option 2 */
	block, err := client.BlockNumber(ctx)
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println("Block: ", block)
	}

	address := "0x4B965600D018C467D39bf3f36bab0a8fD9d5A0B1" // wallet wich i got in Linea hackathon :)


	balance, err := client.BalanceAt(ctx, common.HexToAddress(address), nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Main balance - ", address, weiToEther(balance).String(),
	"ETH")



	latest, err := client.BlockByNumber(ctx, big.NewInt(int64(block)))
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("blk num: %v\n", latest.Number())

	for _, tx := range latest.Transactions() {
		fmt.Println("Tx Hash: ", tx.Hash().Hex())
		fmt.Println("Tx Value: ", tx.Value().String())    // 10000000000000000
		fmt.Println("Tx Gas: ", tx.Gas())               // 105000
		fmt.Println("Tx Gas Price: ", tx.GasPrice().Uint64()) // 102000000000
		fmt.Println("Tx Nonce: ", tx.Nonce())             // 110644
		fmt.Println("Tx Data: ", tx.Data())              // []
		if tx.To() != nil {
			fmt.Println("Tx To: ", tx.To().Hex())
		} else {
			fmt.Println("Tx To: nil")
		}
	}
}
