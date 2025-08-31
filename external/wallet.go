package external

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"ewallet-ums/helpers"
	"fmt"
	"net/http"
)

type Wallet struct {
	Id      int     `json:"id"`
	UserId  int     `json:"user_Id"`
	Balance float64 `json:"balance"`
}

type ExtWallet struct {
}

func (*ExtWallet) CreateWallet(ctx context.Context, userId int) (*Wallet, error) {

	req := Wallet{UserId: userId}
	payload, err := json.Marshal(req)
	if err != nil {
		return nil, errors.New("failed marshal payload")
	}

	url := helpers.GetEnv("WALLET_HOST", "") + helpers.GetEnv("WALLET_ENDPOINT_CREATE", "")
	httpReq, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(payload))
	if err != nil {
		return nil, errors.New("failed create http request")
	}
	client := &http.Client{}
	resp, err := client.Do(httpReq)
	if err != nil {
		return nil, errors.New("failed to connect wallet service")
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to create wallet, status code: %d", resp.StatusCode)
	}

	result := Wallet{}
	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		return nil, errors.New("failed to decode response")
	}
	defer resp.Body.Close()
	return &result, nil
}
