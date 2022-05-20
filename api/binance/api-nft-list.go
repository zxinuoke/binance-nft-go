package bapi

import (
	"encoding/json"
	"fmt"
	"github.com/zxinuoke/binance-nft-go/common"
	binance_struct "github.com/zxinuoke/binance-nft-go/core/domain/binance-api"
	"os"
)

const (
	urlNFTMysteryBoxList = "https://www.binance.com/bapi/nft/v1/public/nft/mystery-box/list?page=1&size=15"
)

func NFTMysteryBoxList() (*binance_struct.NftMysteryBoxesListResponse, error) {
	resp, err := common.RequestJsonBufByGet(urlNFTMysteryBoxList, fmt.Sprintf("http://%s", os.Getenv("PROXY")))
	if err != nil {
		return nil, err
	}
	b, err := unmarshalNFTMysteryBoxList(resp)
	if err != nil {
		return nil, err
	}
	return b, nil
}

func unmarshalNFTMysteryBoxList(resp []byte) (*binance_struct.NftMysteryBoxesListResponse, error) {
	var response binance_struct.NftMysteryBoxesListResponse
	err := json.Unmarshal(resp, &response)
	if err != nil {
		return nil, err
	}
	return &response, nil
}
