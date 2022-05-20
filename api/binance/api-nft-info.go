package bapi

import (
	"encoding/json"
	"fmt"
	"github.com/zxinuoke/binance-nft-go/common"
	binance_struct "github.com/zxinuoke/binance-nft-go/core/domain/binance-api"
	"github.com/zxinuoke/binance-nft-go/core/domain/mysterybox"
	"os"
	"strconv"
)

const (
	urlNFTMysteryBoxInfo = "https://www.binance.com/bapi/nft/v1/friendly/nft/mystery-box/detail?productId=%s"
)

func NFTMysteryBoxInfo(productID string) (*mysterybox.Information, error) {
	r, err := common.RequestJsonBufByGet(fmt.Sprintf(urlNFTMysteryBoxInfo, productID), fmt.Sprintf("http://%s", os.Getenv("PROXY")))
	if err != nil {
		return nil, err
	}
	b, err := unmarshalNFTMysteryBoxInfo(r)
	if err != nil {
		return nil, err
	}
	return b, nil
}

func unmarshalNFTMysteryBoxInfo(resp []byte) (*mysterybox.Information, error) {
	var response binance_struct.NftMysteryBoxesInfoResponse
	err := json.Unmarshal(resp, &response)
	if err != nil {
		return nil, err
	}
	price, err := strconv.ParseFloat(response.Data.Price, 64)
	if err != nil {
		return nil, err
	}
	userBalance, err := strconv.ParseFloat(response.Data.UserBalance, 64)
	if err != nil {
		return nil, err
	}
	return &mysterybox.Information{
		Price:       price,
		Balance:     userBalance,
		LimitPerBuy: response.Data.LimitPerTime,
		StartTime:   response.Data.StartTime / 1000,
	}, nil
}
