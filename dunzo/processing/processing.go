package processing

import "lld/dunzo/models/common.go"

func ConsumeIngredients(source, bev common.IngredientQuantity) common.IngredientQuantity {
	for k, v := range bev {
		source[k] -= v
	}
	return source
}

func CanServe(present common.IngredientQuantity, bev common.Beverage) bool {
	for ingredient, quantity := range bev.IngredientQuantity {
		if currentQuantity, present := present[ingredient]; !present || currentQuantity < quantity {
			return false
		}
	}

	return true
}

func RemoveImpossible(present common.IngredientQuantity, bevs []common.Beverage) []common.Beverage {
	bevsArray := make([]common.Beverage, 0, len(bevs))

	for _, bev := range bevs {
		if CanServe(present, bev) {
			bevsArray = append(bevsArray, bev)
		}
	}

	return bevsArray
}
