package main

import (
	//"bytes"
	"context"
	//"encoding/json"
	//"flag"
	"fmt"
	"os"
	"time"

	"github.com/SundaeSwap-finance/ogmigo/v6"
	"github.com/SundaeSwap-finance/ogmigo/v6/ouroboros/chainsync"
)

func main() {
	var callback ogmigo.MonitorMempoolFunc = func(ctx context.Context, snapshot []*chainsync.Tx, slot uint64) error {
		for _, tx := range snapshot {
			fmt.Printf("tx.ID: %s\n", tx.ID)
		}
		return nil
	}

	//debugPtr := flag.Bool("debug", false, "Print debug statements")
	//flag.Parse()
	ctx := context.Background()

	ogmiosAddr := os.Getenv("OGMIOS")
	client := ogmigo.New(ogmigo.WithEndpoint(ogmiosAddr))
	closer, err := client.MonitorMempool(ctx, callback)
	if err != nil {
		fmt.Printf("Failed MonitorMempool Open: %v\n", err)
		return
	}

	time.Sleep(60 * time.Second)

	if err := closer.Close(); err != nil {
		fmt.Printf("Failed MonitorMempool Close: %v\n", err)
	}
}
