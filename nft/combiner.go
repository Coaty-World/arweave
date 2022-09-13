package nft

import (
	"bufio"
	"bytes"
	"fmt"
	"image/color"
	"image/png"
	"io"
	"log"
	"net/http"
	"os"
	"sort"

	"github.com/Coaty-World/coaty-api/domain/item"
	gim "github.com/ozankasikci/go-image-merge"
)

type ImageCombiner struct {
	CDNURL string
}

func NewImageCombiner(cdnURL string) *ImageCombiner {
	return &ImageCombiner{
		CDNURL: cdnURL,
	}
}

func (i *ImageCombiner) downloadImage(name string) (string, error) {
	resp, err := http.Get(i.CDNURL + "/" + name)
	if err != nil {
		return "", err
	}

	file, err := os.CreateTemp("./", "asset.*.png")
	if err != nil {
		return "", err
	}

	_, err = io.Copy(file, resp.Body)
	if err != nil {
		return "", err
	}

	return file.Name(), nil
}

func (i *ImageCombiner) createRaccoonAsset() (string, error) {
	file, err := os.CreateTemp("./", "raccoon.*.png")
	if err != nil {
		return "", err
	}

	_, err = io.Copy(file, bytes.NewBuffer(baseSkin))
	if err != nil {
		return "", err
	}

	return file.Name(), nil
}

func (i *ImageCombiner) CombineItems(items []item.Item) ([]byte, error) {
	backgroundItems := make([]item.Item, 0)
	foregroundItems := make([]item.Item, 0)
	for _, v := range items {
		if v.ZIndex < 100 {
			backgroundItems = append(backgroundItems, v)
		} else {
			foregroundItems = append(foregroundItems, v)
		}
	}
	sort.Slice(backgroundItems, func(i, j int) bool {
		return backgroundItems[i].ZIndex < backgroundItems[j].ZIndex
	})
	sort.Slice(foregroundItems, func(i, j int) bool {
		return foregroundItems[i].ZIndex < foregroundItems[j].ZIndex
	})

	assets := make([]string, 0, len(items)+1)
	for _, v := range backgroundItems {
		asset, err := i.downloadImage(v.CharacterAsset)
		if err != nil {
			return nil, err
		}

		assets = append(assets, asset)
	}
	raccoonAsset, err := i.createRaccoonAsset()
	if err != nil {
		return nil, err
	}
	assets = append(assets, raccoonAsset)
	for _, v := range foregroundItems {
		asset, err := i.downloadImage(v.CharacterAsset)
		if err != nil {
			return nil, err
		}

		assets = append(assets, asset)
	}

	defer func() {
		for _, v := range assets {
			if err := os.Remove(v); err != nil {
				log.Println(err)
			}
		}
	}()

	if len(assets) <= 1 {
		return nil, fmt.Errorf("no assets to combine")
	}

	grids := make([]*gim.Grid, 0, len(assets)-1)
	for i := 1; i < len(assets); i++ {
		grid := &gim.Grid{
			ImageFilePath:   assets[i],
			BackgroundColor: color.Transparent,
		}
		if err != nil {
			return nil, err
		}

		grids = append(grids, grid)
	}

	allGrids := []*gim.Grid{
		{
			ImageFilePath:   assets[0],
			BackgroundColor: color.Transparent,
			Grids:           grids,
		},
	}

	output, err := gim.New(allGrids, 1, 1).Merge()
	if err != nil {
		return nil, err
	}

	var result bytes.Buffer

	if err := png.Encode(bufio.NewWriter(&result), output); err != nil {
		return nil, err
	}

	return result.Bytes(), nil
}
