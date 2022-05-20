package bapi

import (
	"encoding/json"
	"github.com/valyala/fasthttp"
	binance_struct "github.com/zxinuoke/binance-nft-go/core/domain/binance-api"
)

const (
	urlUserInfo = "https://www.binance.com/bapi/accounts/v1/private/account/user/base-detail"
)

func (api *Api) GetEmail() (string, error) {
	resp, err := api.post(urlUserInfo, nil)
	if err = handleError(resp, err); err != nil {
		return "", err
	}
	user, err := unmarshalUserInfo(resp)
	if err != nil {
		return "", err
	}
	return user.Data.Email, nil
}

func unmarshalUserInfo(resp *fasthttp.Response) (*binance_struct.UserInformationResponse, error) {
	var response binance_struct.UserInformationResponse
	err := json.Unmarshal(resp.Body(), &response)
	if err != nil {
		return nil, err
	}
	return &response, nil
}
