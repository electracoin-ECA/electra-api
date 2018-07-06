package api

import (
	"context"
	"errors"
	"testing"
)

func TestGetLatestBlock(t *testing.T) {
	resp, err := GetLatestBlock()

	if err != nil {
		t.Fatal(err)
	}

	// uninitialized
	if resp.Height == 0 {
		t.Error("malformed response on GetLatestBlock")
	}

}

func TestGetBlock(t *testing.T) {
	blk, err := GetLatestBlock()
	if err != nil {
		t.Fatal(err)
	}
	resp, err := GetBlock(blk.Hash)

	if err != nil {
		t.Fatal(err)
	}

	if resp.Block.Hash != blk.Hash {
		t.Fatal("malformed hash received on using GetBlock")
	}

	// t.Logf("%+v", resp)
}

func TestGetPreviousBlocks(t *testing.T) {
	blk, err := GetLatestBlock()
	if err != nil {
		t.Fatal(err)
	}

	ctx, cancel := context.WithCancel(context.TODO())

	blkChan, err := GetPreviousBlocks(ctx, blk.Hash)
	if err != nil {
		t.Fatal(err)
	}
	height := blk.Height
	hash := blk.Hash
	nblk := BlockResponse{}
	for i := 0; i < 5; i++ {
		nblk = <-blkChan
		if nblk.Block.Height+1 != height {
			t.Fatal(errors.New("Missing blocks"))
		}

		if nblk.Block.Nextblockhash != hash {
			t.Fatal(errors.New("Parent blockhash is not a match"))
		}

		hash = nblk.Block.Hash
		height = nblk.Block.Height
	}

	cancel()
}
