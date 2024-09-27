package main

import (
	"errors"
	"fmt"
)

type Set struct {
	elements map[[4]int]struct{}
}

func NewSet() *Set {
	return &Set{
		elements: make(map[[4]int]struct{}),
	}
}

func (s *Set) Add(e [4]int) {
	s.elements[e] = struct{}{}
}

// makeChange calculates all the possible ways to represent n cents using quarters, dimes, nickels, and pennies, up until a 100 cents.
// It creates a set to store the combination of coins, then uses four nested loops to represent the number of coins of each type
// and then returns the result set containing all of the combination
func makeChange(n int) (*Set, error) {

	if n > 100 {
		return nil, errors.New("amount exceeds the limit of 100 cents") //limit of a 100 cents defined to not put a lot of stuff on the terminal
	}

	resultSet := NewSet()

	for quarters := 0; quarters <= n/25; quarters++ {
		for dimes := 0; dimes <= (n-quarters*25)/10; dimes++ {
			for nickels := 0; nickels <= (n-quarters*25-dimes*10)/5; nickels++ {
				pennies := n - quarters*25 - dimes*10 - nickels*5

				if pennies >= 0 {
					resultSet.Add([4]int{quarters, dimes, nickels, pennies})
				}
			}
		}
	}

	return resultSet, nil
}

func main() {
	var n int
	fmt.Print("Enter the amount in cents (up to 100 cents): ")
	fmt.Scan(&n)

	result, err := makeChange(n)
	if err != nil {
		fmt.Println(err)
		return
	}

	const maxCombinations = 1000
	var comboArray [maxCombinations][4]int
	index := 0

	for e := range result.elements {
		if index < maxCombinations {
			comboArray[index] = e
			index++
		} else {
			return
		}
	}

	fmt.Println(comboArray[:index])
}
