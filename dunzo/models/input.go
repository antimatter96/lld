package models

import "lld/dunzo/models/common"

type JsonInput struct {
	InputMachine `json:"machine"`
}
type Outlets struct {
	CountN int `json:"count_n"`
}

type InputMachine struct {
	Outlets            Outlets                   `json:"outlets"`
	TotalItemsQuantity common.IngredientQuantity `json:"total_items_quantity"`
	Beverages          common.BeveragesMap       `json:"beverages"`
}
