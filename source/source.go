package source

import "github.com/antitokens/priceindex/model"

type Source interface {
	GetPrice(token string) (model.Price, error)
}
