package api

import "testing"

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
