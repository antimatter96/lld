package main

import (
	"encoding/json"
	"fmt"
	"lld/dunzo/models"
	"lld/dunzo/models/common"
	"lld/dunzo/processing"
	"testing"
)

func TestFunctional(t *testing.T) {
	newjson := &models.JsonInput{}
	err := json.Unmarshal([]byte(input), &newjson)
	if err != nil {
		panic(err)
	}

	m := &models.Machine{}

	m.InitWithJson(newjson.InputMachine)

	t.Run("Inital Status", func(t *testing.T) {
		currentStatus := m.Status()
		expectedStatus := common.IngredientQuantity{"ginger_syrup": 100, "hot_milk": 500, "hot_water": 500, "sugar_syrup": 100, "tea_leaves_syrup": 100}

		if fmt.Sprint(currentStatus) != fmt.Sprint(expectedStatus) {
			t.Errorf("Ingredient status not same got %v != %v wanted", currentStatus, expectedStatus)
		}
	})

	t.Run("First Serving", func(t *testing.T) {
		served := m.ServeJson(processing.Common, newjson.Beverages, false)
		expectedServing := []common.Beverage{
			{Name: "black_tea", IngredientQuantity: common.IngredientQuantity{}},
			{Name: "hot_coffee", IngredientQuantity: common.IngredientQuantity{}},
		}

		if len(served) != len(expectedServing) {
			t.Errorf("Not served as expected %v != %v", served, expectedServing)
		}

		for i := 0; i < len(served); i++ {
			if expectedServing[i].Name != served[i].Name {
				t.Errorf("Output %d %v != %v", i+1, served[i].Name, expectedServing[i].Name)
			}
		}
	})

	t.Run("After first serving", func(t *testing.T) {
		currentStatus := m.Status()
		expectedStatus := common.IngredientQuantity{"ginger_syrup": 40, "hot_milk": 100, "hot_water": 100, "sugar_syrup": 0, "tea_leaves_syrup": 40}

		if fmt.Sprint(currentStatus) != fmt.Sprint(expectedStatus) {
			t.Errorf("Ingredient status not same got %v != %v wanted", currentStatus, expectedStatus)
		}
	})

	t.Run("After refill", func(t *testing.T) {
		m.Refill(common.IngredientQuantity{"green_mixture": 30, "sugar_syrup": 50})
		currentStatus := m.Status()
		expectedStatus := common.IngredientQuantity{"ginger_syrup": 40, "hot_milk": 100, "hot_water": 100, "sugar_syrup": 50, "tea_leaves_syrup": 40, "green_mixture": 30}

		if fmt.Sprint(currentStatus) != fmt.Sprint(expectedStatus) {
			t.Errorf("Ingredient status not same got %v != %v wanted", currentStatus, expectedStatus)
		}

	})

	t.Run("Second run", func(t *testing.T) {
		served := m.ServeJson(processing.Common, newjson.Beverages, false)
		expectedServing := []common.Beverage{{Name: "green_tea"}}

		if len(served) != len(expectedServing) {
			t.Errorf("Not served as expected %v != %v", served, expectedServing)
		}

		for i := 0; i < len(served); i++ {
			if expectedServing[i].Name != served[i].Name {
				t.Errorf("Output %d %v != %v", i+1, served[i].Name, expectedServing[i].Name)
			}
		}

	})
}
