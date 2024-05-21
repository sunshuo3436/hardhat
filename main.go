package main

import (
	"context"
	"fmt"
	"log"
	"math/big"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

const (
	infuraURL    = "https://holesky.infura.io/v3/c0fd902e8c32475f85909278c7352314"
	contractAddr = "0xEcF4c423a74C7cbc722F7201D0BD3625f8E12a7b"
	contractABI  = `[{"inputs":[{"internalType":"address","name":"to","type":"address"},{"internalType":"uint256","name":"value","type":"uint256"}],"name":"transfer","outputs":[{"internalType":"bool","name":"","type":"bool"}],"stateMutability":"nonpayable","type":"function"},{"anonymous":false,"inputs":[{"indexed":true,"internalType":"address","name":"from","type":"address"},{"indexed":true,"internalType":"address","name":"to","type":"address"},{"indexed":false,"internalType":"uint256","name":"value","type":"uint256"}],"name":"Transfer","type":"event"}]`
)

func main() {
	client, err := ethclient.Dial(infuraURL)
	if err != nil {
		log.Fatalf("Failed to connect to the Ethereum client: %v", err)
	}

	contractAddress := common.HexToAddress(contractAddr)
	query := ethereum.FilterQuery{
		FromBlock: big.NewInt(1584992),
		ToBlock:   big.NewInt(1585043),
		Addresses: []common.Address{contractAddress},
	}

	parsedABI, err := abi.JSON(strings.NewReader(contractABI))
	if err != nil {
		log.Fatalf("Failed to parse contract ABI: %v", err)
	}

	for {
		logs, err := client.FilterLogs(context.Background(), query)
		if err != nil {
			log.Fatalf("Failed to retrieve logs: %v", err)
		}

		for _, vLog := range logs {
			fmt.Printf("Log Block Number: %d\n", vLog.BlockNumber)
			fmt.Printf("Log Index: %d\n", vLog.Index)

			var transferEvent struct {
				From  common.Address
				To    common.Address
				Value *big.Int
			}

			// 解析未索引的参数
			err := parsedABI.UnpackIntoInterface(&transferEvent, "Transfer", vLog.Data)
			if err != nil {
				log.Fatalf("Failed to unpack log data: %v", err)
			}

			fmt.Printf("From: %s\n", transferEvent.From.Hex())
			fmt.Printf("To: %s\n", transferEvent.To.Hex())
			fmt.Printf("Value: %s\n", transferEvent.Value.String())

			// 访问索引的参数
			transferEvent.From = common.HexToAddress(vLog.Topics[1].Hex())
			transferEvent.To = common.HexToAddress(vLog.Topics[2].Hex())
			fmt.Printf("Indexed From: %s\n", transferEvent.From.Hex())
			fmt.Printf("Indexed To: %s\n", transferEvent.To.Hex())
		}

		// 每隔一段时间进行轮询
		time.Sleep(5 * time.Second)
	}
}
