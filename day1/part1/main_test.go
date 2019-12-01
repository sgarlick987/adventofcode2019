package main

import "testing"

/**
Test against the examples from the problem statement

    For a mass of 12, divide by 3 and round down to get 4, then subtract 2 to get 2.
    For a mass of 14, dividing by 3 and rounding down still yields 4, so the fuel required is also 2.
    For a mass of 1969, the fuel required is 654.
    For a mass of 100756, the fuel required is 33583.
*/
func TestFuelCost(t *testing.T) {
	cost := fuelCost(12)
	if cost != 2 {
		t.Errorf("wanted 2 for mass of 12 but got %d", cost)
	}
	cost = fuelCost(14)
	if cost != 2 {
		t.Errorf("wanted 2 for mass of 14 but got %d", cost)
	}
	cost = fuelCost(1969)
	if cost != 654 {
		t.Errorf("wanted 654 for mass of 1969 but got %d", cost)
	}
	cost = fuelCost(100756)
	if cost != 33583 {
		t.Errorf("wanted 33583 for mass of 100756 but got %d", cost)
	}
}
