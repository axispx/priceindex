package source

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/antitokens/priceindex/model"
	"github.com/antitokens/priceindex/utils"
)

type RaydiumResponse struct {
	ID      string            `json:"id"`
	Success bool              `json:"success"`
	Data    map[string]string `json:"data"`
}

type Raydium struct {
}

func NewRaydium() Raydium {
	return Raydium{}
}

func (r Raydium) GetPrice(token string) (model.PriceResponse, error) {
	tokenAddress := utils.GetTokenAddress(token)

	resp, err := http.Get("https://api-v3.raydium.io/mint/price?mints=" + tokenAddress)
	if err != nil {
		return model.PriceResponse{}, err
	}
	defer resp.Body.Close()

	var raydiumResponse RaydiumResponse
	err = json.NewDecoder(resp.Body).Decode(&raydiumResponse)
	if err != nil {
		return model.PriceResponse{}, err
	}

	return model.PriceResponse{
		Price:     raydiumResponse.Data[tokenAddress],
		Timestamp: time.Now().Unix(),
		Source:    "raydium",
		Address:   tokenAddress,
	}, nil
}
