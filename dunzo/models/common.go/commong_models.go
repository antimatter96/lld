package common

type Beverage struct {
	Name               string
	IngredientQuantity IngredientQuantity
}
type IngredientQuantity map[string]int

type BeveragesMap map[string]IngredientQuantity
