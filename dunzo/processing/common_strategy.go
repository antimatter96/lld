package processing

import (
	"lld/dunzo/models/common.go"
	"sort"
)

// Common stratergy returns the beverages
// that are possible to serve in alphabatical order
func Common(ingredients common.IngredientQuantity, bevsArray []common.Beverage, n int) []common.Beverage {
	bevsArray = RemoveImpossible(ingredients, bevsArray)

	sort.Slice(bevsArray, func(i int, j int) bool {
		return bevsArray[i].Name < bevsArray[j].Name
	})

	serve := make([]common.Beverage, 0, len(bevsArray))

	for _, v := range bevsArray {
		if CanServe(ingredients, v) {
			ingredients = ConsumeIngredients(ingredients, v.IngredientQuantity)
			serve = append(serve, v)
		}
	}

	return serve
}
