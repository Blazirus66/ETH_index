package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/big"
	"net/http"
	"time"
)

type Transaction struct {
	Hash      string `json:"hash"`
	Value     string `json:"value"`
	Timestamp int    `json:"timestamp"`
}

type Transactions struct {
	Operations []Transaction `json:"operations"`
}

func main() {
	address := "0x742d35Cc6634C0532925a3b844Bc454e4438f44e"
	url := fmt.Sprintf("https://api.ethplorer.io/getAddressHistory/%s?apiKey=freekey&limit=100", address)

	resp, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return
	}

	var txs Transactions
	err = json.Unmarshal(body, &txs)
	if err != nil {
		fmt.Println(err)
		return
	}

	for i, tx := range txs.Operations {
		value := new(big.Int)
		value, success := value.SetString(tx.Value[2:], 16)
		if !success {
			fmt.Println("Failed to parse value")
			return
		}

		timestamp := time.Unix(int64(tx.Timestamp), 0)
		fmt.Printf("Transaction %d: value %.8f ETH, timestamp %s\n", i, float64(value.Int64())/1000000000000000000, timestamp.Format("02/01/2006"))
	}
}
