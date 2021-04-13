package main

var input = `{
  "machine": {
    "outlets": {
      "count_n": 3
    },
    "total_items_quantity": {
      "hot_water": 500,
      "hot_milk": 500,
      "ginger_syrup": 100,
      "sugar_syrup": 100,
      "tea_leaves_syrup": 100
    },
    "beverages": {
      "hot_tea": {
        "hot_water": 200,
        "hot_milk": 100,
        "ginger_syrup": 10,
        "sugar_syrup": 10,
        "tea_leaves_syrup": 30
      },
      "hot_coffee": {
        "hot_water": 100,
        "ginger_syrup": 30,
        "hot_milk": 400,
        "sugar_syrup": 50,
        "tea_leaves_syrup": 30
      },
      "black_tea": {
        "hot_water": 300,
        "ginger_syrup": 30,
        "sugar_syrup": 50,
        "tea_leaves_syrup": 30
      },
      "green_tea": {
        "hot_water": 100,
        "ginger_syrup": 30,
        "sugar_syrup": 50,
        "green_mixture": 30
      }
    }
  }
}`

type Input struct {
	Machine InputMachine `json:"machine"`
}
type Outlets struct {
	CountN int `json:"count_n"`
}

type Ingredients map[string]int
type Beverage struct {
	Name        string
	Ingredients Ingredients
}
type Beverages map[string]Ingredients

type InputMachine struct {
	Outlets            Outlets     `json:"outlets"`
	TotalItemsQuantity Ingredients `json:"total_items_quantity"`
	Beverages          Beverages   `json:"beverages"`
}
