package source

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/antitokens/priceindex/model"
	"github.com/antitokens/priceindex/utils"
	"github.com/shopspring/decimal"
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

func (r Raydium) GetPrice(tokens ...string) ([]model.Price, error) {
	tokenAddresses := utils.GetTokenAddress(tokens...)

	resp, err := http.Get("https://api-v3.raydium.io/mint/price?mints=" + strings.Join(tokenAddresses, ","))
	if err != nil {
		return []model.Price{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return []model.Price{}, fmt.Errorf("failed to get price from raydium")
	}

	var raydiumResponse RaydiumResponse
	err = json.NewDecoder(resp.Body).Decode(&raydiumResponse)
	if err != nil {
		return []model.Price{}, err
	}

	prices := []model.Price{}
	for _, tokenAddress := range tokenAddresses {
		price := raydiumResponse.Data[tokenAddress]
		if price == "" {
			continue
		}

		priceDecimal, err := decimal.NewFromString(price)
		if err != nil {
			return []model.Price{}, err
		}

		prices = append(prices, model.Price{
			Price:   priceDecimal,
			Source:  "raydium",
			Address: tokenAddress,
		})
	}

	return prices, nil
}
