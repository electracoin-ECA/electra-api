package main

import (
	"context"
	"flag"
	"log"
	"os"
	"os/signal"

	api "github.com/Electra-project/electra-api/src/libs/electra-api"
	"github.com/Electra-project/electra-api/src/models"
)

var fullSync *bool

func isBlockPresent(hash string) bool {
	present, err := models.IsBlockPresent(hash)
	// possible case of failure is network disconnection
	if err != nil {
		log.Fatal(err)
	}

	if present {
		log.Println("Block ", hash, "already present in DB")
		if !*fullSync {
			log.Println("Fullsync disabled stopping the daemon")
			os.Exit(0)
		}
	}

	return present
}

// storeBlock is a helper which checks if the block is exists
func storeBlockIfNotExists(hash string) {

	present := isBlockPresent(hash)

	// we continue
	if present {
		return
	}

	blockResp, err := api.GetBlock(hash)

	err = models.StoreBlock(blockResp.Block)
	if err != nil {
		log.Fatal(err)
	}

}

func main() {

	// use this to sync the whole blockchain
	fullSync = flag.Bool("fullsync", false, "When set to true, the daemon will not stop when encountered a known block")
	flag.Parse()

	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, os.Interrupt)
	// use cancel with SIGKILL to handle cleanup
	ctx, cancel := context.WithCancel(context.Background())

	go func() {
		<-signalChan
		log.Println("Exiting")
		cancel()
		os.Exit(0)
	}()

	block, err := api.GetLatestBlock()
	if err != nil {
		log.Fatal(err)
	}

	storeBlockIfNotExists(block.Hash)

	blockChan, err := api.GetPreviousBlocks(ctx, block.Hash)
	if err != nil {
		cancel()
		log.Fatal(err)
	}
	for blockResp := range blockChan {
		present := isBlockPresent(blockResp.Block.Hash)
		if !present {

			log.Println("Storing block", blockResp.Block.Hash, "in DB")
			err = models.StoreBlock(blockResp.Block)

			if err != nil {
				log.Fatal(err)
			}
		} else { // if present
			if !*fullSync {
				log.Println("Fullsync disabled stopping the daemon")
				os.Exit(0)
			}
		}
		addresses := api.GetUniqueAddresses(blockResp.Txs...)
		addressChan := api.GetAddressDatas(addresses...)

		// this is to make sure if someone stops the daemon in the middle
		// we still add addresses which are left
		for address := range addressChan {
			log.Println("Storing Address", address.Addr, "in DB")

			err := models.StoreAddressIfNotPresent(&address)
			if err != nil {
				log.Fatal(err)
			}
		}
	}
}
