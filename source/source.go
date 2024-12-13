package source

import "github.com/antitokens/priceindex/model"

type Source interface {
	GetPrice(tokens ...string) ([]model.Price, error)
}
