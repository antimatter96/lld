package main

import (
	"fmt"
	"strings"
)

type Machine struct {
	n           int
	ingredients Ingredients
}

type MachineOutput []string

func (m *Machine) Init(inp InputMachine) {
	m.n = inp.Outlets.CountN
	m.ingredients = inp.TotalItemsQuantity
}

func (m *Machine) CanServe(bev Beverage) bool {

	for ingredient, quantity := range bev.Ingredients {
		if ingredient == "_name_" {
			continue
		}
		if currentQuantity, present := m.ingredients[ingredient]; !present || currentQuantity < quantity {
			return false
		}
	}

	return true
}

func (m *Machine) TryServe(bevs Beverages) {
	bevsArray := make([]Beverage, 0, len(bevs))
	//nBev := len(bevs)

	for name, ingredients := range bevs {
		b := &Beverage{Name: name, Ingredients: ingredients}
		if m.CanServe(*b) {
			bevsArray = append(bevsArray, *b)
		}
	}

	fmt.Println(len(bevsArray), bevsArray)
}

// func (m *Machine) Serve(b Beverage) {

// }

func (m *Machine) String() string {
	str := strings.Builder{}

	str.WriteString(fmt.Sprintf("Outputs : %d\n", m.n))
	str.WriteString("Ingredients :\n")

	for k, v := range m.ingredients {
		str.WriteString(fmt.Sprintf("%-16s - %d\n", k, v))
	}

	return str.String()
}
