package main

import "fmt"

func main() {
	// turn decider
	// parse input
	// clculate score
	// decicde next step

	// fmt.Println(IsValidInput("/"))
	// fmt.Println(IsValidInput("9 /"))
	// fmt.Println(IsValidInput("9 /"))
	// fmt.Println(IsValidInput("-"))
	// fmt.Println(IsValidInput("9 -"))
	// fmt.Println(IsValidInput("9 / X"))
	// fmt.Println(IsValidInput("9 -"))
	// fmt.Println(IsValidInput("8 /"))
	// fmt.Println(IsValidInput("X"))
	// fmt.Println(IsValidInput("7 2"))
	// fmt.Println(IsValidInput("8 / 1"))

	// test_cases := [][]string{
	// 	{"X"},
	// 	{"1", "2", "X", "-"},
	// 	{"1", "/", "1", "-"},
	// 	{"X", "-", "1", "2"},
	// 	{"-", "-", "-", "-", "-", "-", "-", "-", "-", "-", "-", "-", "-", "-", "-", "-", "-", "-", "1", "/", "1"},
	// }

	test_cases := []struct {
		inp   []int
		score int
	}{
		{[]int{10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10}, 300},
		{[]int{3, 5, 1, 5, 7, 1, 10, 1, 6, 10, 6, 2, 1, 2, 0, 5, 8, 1}, 89},
		{[]int{9, 1, 5, 0, 3, 0, 8, 1, 6, 4, 7, 2, 7, 1, 6, 3, 10, 4, 4}, 101},
		{[]int{10, 8, 0, 4, 2, 6, 4, 10, 5, 3, 4, 5, 9, 1, 6, 4, 8, 2, 1}, 132},
		{[]int{3, 4, 4, 3, 7, 3, 8, 2, 0, 3, 5, 3, 7, 1, 8, 2, 3, 2, 5, 3}, 87},
		{[]int{2, 1, 5, 3, 8, 2, 8, 2, 8, 2, 2, 1, 3, 6, 0, 7, 0, 9, 9, 1, 10}, 107},
		{[]int{10, 10, 0, 1, 0, 2, 0, 3, 10, 10, 10, 0, 4, 0, 5}, 110},
		{[]int{10, 10, 5, 3, 10, 10, 7, 1, 10, 10, 10, 10, 10, 10}, 224},
		{[]int{0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0}, 10},
		{[]int{0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 10, 10, 10, 10, 10, 10, 10}, 155},
	}

	for _, tc := range test_cases {
		fmt.Println("<<= >>===<< =>>")
		fmt.Println(CalculateScore(tc.inp), tc.score)
	}

}
