package processing

import (
	"fmt"
	"lld/dunzo/models/common"
	"reflect"
	"testing"
)

func TestConsumeIngredients(t *testing.T) {
	ConsumeIngredientsTests := []struct {
		Name      string
		Current   common.IngredientQuantity
		ToConsume common.IngredientQuantity
		Expected  common.IngredientQuantity
	}{
		{"simple",
			common.IngredientQuantity{"a": 20},
			common.IngredientQuantity{"a": 10},
			common.IngredientQuantity{"a": 10},
		},
		{"one gets fully used",
			common.IngredientQuantity{"a": 20, "b": 20},
			common.IngredientQuantity{"a": 10, "b": 20},
			common.IngredientQuantity{"a": 10, "b": 0},
		},
		{"does not care if it exceeds",
			common.IngredientQuantity{"a": 20},
			common.IngredientQuantity{"b": 10},
			common.IngredientQuantity{"a": 20, "b": -10},
		},
		{"does not care if it exceeds",
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
			}
		})
	}
}

func TestCanServe(t *testing.T) {
	ConsumeIngredientsTests := []struct {
		Name      string
		Current   common.IngredientQuantity
		ToConsume common.Beverage
		Expected  bool
	}{
		{"simple",
			common.IngredientQuantity{"a": 20},
			common.Beverage{Name: "drink1", IngredientQuantity: common.IngredientQuantity{"a": 10}},
			true,
		},
		{"simple",
			common.IngredientQuantity{"a": 20, "b": 10},
			common.Beverage{Name: "drink1", IngredientQuantity: common.IngredientQuantity{"a": 10, "b": 5}},
			true,
		},
		{"not enough quantity",
			common.IngredientQuantity{"a": 20},
			common.Beverage{Name: "drink1", IngredientQuantity: common.IngredientQuantity{"a": 30}},
			false,
		},
		{"ingredients missing",
			common.IngredientQuantity{"a": 20},
			common.Beverage{Name: "drink1", IngredientQuantity: common.IngredientQuantity{"b": 30}},
			false,
		},
	}

	for _, tc := range ConsumeIngredientsTests {
		t.Run(tc.Name, func(t *testing.T) {
			updated := CanServe(tc.Current, tc.ToConsume)
			if updated != tc.Expected {
				t.Errorf("%v != %v", updated, tc.Expected)
			}
		})
	}
}

func TestRemoveImpossible(t *testing.T) {
	ConsumeIngredientsTests := []struct {
		Name     string
		Current  common.IngredientQuantity
		Options  []common.Beverage
		Expected []common.Beverage
	}{
		{"simple",
			common.IngredientQuantity{"a": 20},
			[]common.Beverage{{Name: "drink1", IngredientQuantity: common.IngredientQuantity{"a": 21}}},
			[]common.Beverage{},
		},
		{"simple",
			common.IngredientQuantity{"a": 20},
			[]common.Beverage{{Name: "drink1", IngredientQuantity: common.IngredientQuantity{"a": 10}}},
			[]common.Beverage{{Name: "drink1", IngredientQuantity: common.IngredientQuantity{"a": 10}}},
		},
		{"leaves only possible drink",
			common.IngredientQuantity{"a": 20},
			[]common.Beverage{{Name: "drink1", IngredientQuantity: common.IngredientQuantity{"a": 20}}, {Name: "drink2", IngredientQuantity: common.IngredientQuantity{"a": 30}}},
			[]common.Beverage{{Name: "drink1", IngredientQuantity: common.IngredientQuantity{"a": 20}}},
		},
		{"leaves only possible drink",
			common.IngredientQuantity{"a": 40},
			[]common.Beverage{{Name: "drink1", IngredientQuantity: common.IngredientQuantity{"a": 20}}, {Name: "drink2", IngredientQuantity: common.IngredientQuantity{"a": 30}}},
			[]common.Beverage{{Name: "drink1", IngredientQuantity: common.IngredientQuantity{"a": 20}}, {Name: "drink2", IngredientQuantity: common.IngredientQuantity{"a": 30}}},
		},
	}

	for _, tc := range ConsumeIngredientsTests {
		t.Run(tc.Name, func(t *testing.T) {
			served := RemoveImpossible(tc.Current, tc.Options)
			fmt.Println(served, tc.Expected)

			if !reflect.DeepEqual(served, tc.Expected) {
				t.Errorf("%v != %v", served, tc.Expected)
			}
		})
	}
}
