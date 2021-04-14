package processing

import "lld/dunzo/models/common"

// Stratergy takes the current ingredients
// and an array of wanted beverages
// and returns the beverages that can be served
// beverages that can be served are based on implementation
type Strategy func(common.IngredientQuantity, []common.Beverage, int) []common.Beverage
