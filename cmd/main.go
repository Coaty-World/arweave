package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/Coaty-World/arweave/arweave"
	"github.com/Coaty-World/arweave/config"
	"github.com/Coaty-World/arweave/nft"
	"github.com/Coaty-World/coaty-api/domain/item"
	"github.com/Coaty-World/coaty-api/server/codes"
	"math/rand"
	"net/http"
	"os"
)

var cfg config.Config

func init() {
	c, err := config.NewConfig()
	if err != nil {
		panic(err)
	}

	cfg = c
}

type Response struct {
	TransactionID string `json:"id"`
}

type Request struct {
	Items []item.Item `json:"items"`
}

func Mint(w http.ResponseWriter, r *http.Request) {
	var items Request

	if err := json.NewDecoder(r.Body).Decode(&items); err != nil {
		http.Error(w, codes.NewError(codes.InvalidArgument, "could not parse data "+err.Error()), http.StatusBadRequest)
		return
	}

	ar := arweave.NewClient(arweave.Wallet, cfg.ArweaveConfig.Store)
	combiner := nft.NewImageCombiner(cfg.CDNConfig.CharacterItems)
	data, err := combiner.CombineItems(items.Items)
	if err != nil {
		http.Error(w, codes.NewError(codes.InternalError, err.Error()), http.StatusInternalServerError)
		return
	}

	// generate random number
	num := rand.Intn(1000000000)

	tx, err := ar.UploadFile(context.Background(), data, fmt.Sprintf("Raccoon🦝 #%v", num), "This is a Coaty World Raccoon. It brings a lot of tokens 🪙")
	if err != nil {
		http.Error(w, codes.NewError(codes.InternalError, err.Error()), http.StatusInternalServerError)
		return
	}

	if err := json.NewEncoder(w).Encode(Response{TransactionID: tx}); err != nil {
		http.Error(w, codes.NewError(codes.InternalError, err.Error()), http.StatusInternalServerError)
		return
	}
}

func main() {
	port := "8888"
	if val, ok := os.LookupEnv("FUNCTIONS_CUSTOMHANDLER_PORT"); ok {
		port = val
	}

	http.HandleFunc("/api/mint", Mint)
	http.ListenAndServe(":"+port, nil)
}
