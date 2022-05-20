package bapi

import (
	"encoding/json"
	"fmt"
	"github.com/valyala/fasthttp"
	binance_struct "github.com/zxinuoke/binance-nft-go/core/domain/binance-api"
)

const (
	errorStatusCode = "statusCode: %d, body: %s"
	errorCode       = "code: %s, body: %s"

	validCode = "000000"
)

func handleError(resp *fasthttp.Response, err error) error {
	if err != nil {
		return err
	}
	if resp.StatusCode() != fasthttp.StatusOK {
		return fmt.Errorf(errorStatusCode, resp.StatusCode(), string(resp.Body()))
	}
	var c binance_struct.ErrorResponse
	if err = json.Unmarshal(resp.Body(), &c); err != nil {
		return err
	}
	if err != nil {
		return err
	}
	if c.Code != validCode {
		return fmt.Errorf(errorCode, c.Code, string(resp.Body()))
	}
	return nil
}
