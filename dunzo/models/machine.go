package models

import (
	"fmt"
	"lld/dunzo/models/common"
	"lld/dunzo/processing"
	"lld/dunzo/utils"
	"sort"
	"strings"
)

type Machine struct {
	n           int
	ingredients common.IngredientQuantity
}

type MachineOutput []string

func (m *Machine) InitWithJson(inp InputMachine) {
	m.n = inp.Outlets.CountN
	m.ingredients = inp.TotalItemsQuantity
}

func (m *Machine) ServeJson(stratergy processing.Strategy, bevs common.BeveragesMap) {
	bevsArray := make([]common.Beverage, 0, len(bevs))
	for k, v := range bevs {
		bevsArray = append(bevsArray, common.Beverage{Name: k, IngredientQuantity: v})
	}

	m.Serve(stratergy, bevsArray)

}

func (m *Machine) Serve(strategy processing.Strategy, bevs []common.Beverage) {
	fmt.Println("Serving ~~")

	serve := strategy(utils.CopyMap(m.ingredients), bevs, m.n)

	for index, bev := range serve {
		m.ingredients = processing.ConsumeIngredients(m.ingredients, bev.IngredientQuantity)

		m.ServeBeverage(bev, index+1)
	}
	fmt.Println("~~ Complete")
	fmt.Println()
}

func (m *Machine) ServeBeverage(bev common.Beverage, outlet int) {
	fmt.Printf("☕ Served %-15s from outlet %-2d\n", bev.Name, outlet)
}

func (m *Machine) Refill(ingredients common.IngredientQuantity) {
	for k, v := range ingredients {
		m.ingredients[k] += v
	}
}

func (m *Machine) String() string {
	str := strings.Builder{}

	str.WriteString("=========\n")
	str.WriteString(fmt.Sprintf("Outputs : %d\n", m.n))
	str.WriteString("---------\n")
	str.WriteString("Ingredients :\n")
	str.WriteString("----\n")

	ingredientNames := make([]string, 0, len(m.ingredients))
	for k := range m.ingredients {
		ingredientNames = append(ingredientNames, k)
	}
	sort.Strings(ingredientNames)
	for _, ingredient := range ingredientNames {
		str.WriteString(fmt.Sprintf("%-20s : %5d\n", ingredient, m.ingredients[ingredient]))
	}

	str.WriteString("===========\n")
	return str.String()
}
