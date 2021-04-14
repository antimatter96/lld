package processing

import (
	"fmt"
	"lld/dunzo/models/common.go"
	"testing"
)

func TestConsumeIngredients(t *testing.T) {
	ConsumeIngredientsTests := []struct {
		Name      string
		Current   common.IngredientQuantity
		ToConsume common.IngredientQuantity
		Expected  common.IngredientQuantity
	}{
		{"Simple",
			common.IngredientQuantity{"a": 20},
			common.IngredientQuantity{"a": 10},
			common.IngredientQuantity{"a": 10},
		},
		{"Simple",
			common.IngredientQuantity{"a": 20, "b": 20},
			common.IngredientQuantity{"a": 10, "b": 20},
			common.IngredientQuantity{"a": 10, "b": 0},
		},
		{"Simple",
			common.IngredientQuantity{"a": 20},
			common.IngredientQuantity{"b": 10},
			common.IngredientQuantity{"a": 20, "b": -10},
		},
		{"Simple",
			common.IngredientQuantity{"a": 20},
			common.IngredientQuantity{"a": 30},
			common.IngredientQuantity{"a": -10},
		},
	}

	for _, tc := range ConsumeIngredientsTests {
		t.Run(tc.Name, func(t *testing.T) {
			updated := ConsumeIngredients(tc.Current, tc.ToConsume)
			if fmt.Sprint(updated) != fmt.Sprint(tc.Expected) {
				t.Errorf("%v != %v", updated, tc.Expected)
			} else {
				t.Logf("%v == %v", updated, tc.Expected)
			}
		})
	}
}
