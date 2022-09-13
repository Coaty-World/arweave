package arweave

import (
	"context"
	"encoding/json"

	"github.com/everFinance/goar/types"

	"github.com/everFinance/goar"
)

type Client struct {
	Wallet []byte
	Store  string
}

func NewClient(wallet []byte, store string) *Client {
	return &Client{
		Wallet: wallet,
		Store:  store,
	}
}

type T struct {
	Media       string        `json:"media"`
	MediaHash   string        `json:"media_hash"`
	Tags        []interface{} `json:"tags"`
	Title       string        `json:"title"`
	Description string        `json:"description"`
	Extra       []interface{} `json:"extra"`
	Store       string        `json:"store"`
	Type        string        `json:"type"`
	Category    interface{}   `json:"category"`
}

type MintbaseData struct {
	Media       string   `json:"media"`
	MediaHash   string   `json:"media_hash"`
	Tags        []string `json:"tags"`
	Title       string   `json:"title"`
	Description string   `json:"description"`
	Extra       []string `json:"extra"`
	Store       string   `json:"store"`
	Type        string   `json:"type"`
	Category    *string  `json:"category"`
}

func (c *Client) UploadFile(ctx context.Context, data []byte, title string, description string) (string, error) {
	wallet, err := goar.NewWallet(c.Wallet, "https://arweave.net")
	if err != nil {
		return "", err
	}

	tx, err := wallet.SendData(data, []types.Tag{
		{
			Name:  "Content-Type",
			Value: "image/png",
		},
	})
	if err != nil {
		return "", err
	}

	mintbaseData := MintbaseData{
		Media:       "https://arweave.net/" + tx.ID,
		MediaHash:   tx.ID,
		Tags:        []string{},
		Title:       title,
		Description: description,
		Extra:       []string{},
		Store:       c.Store,
		Type:        "NEP171",
		Category:    nil,
	}
	jsonData, err := json.Marshal(mintbaseData)
	if err != nil {
		return "", err
	}

	tx, err = wallet.SendBundleTx(jsonData, []types.Tag{
		{
			Name:  "Content-Type",
			Value: "application/json",
		},
	})
	if err != nil {
		return "", err
	}

	return tx.ID, nil
}
