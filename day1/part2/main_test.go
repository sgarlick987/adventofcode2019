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

/**
Needs to make sure negative fuel costs are returned as 0
*/
func TestFuelCostNegativeReturnsZero(t *testing.T) {
	cost := fuelCost(1)
	if cost != 0 {
		t.Errorf("wanted 0 for mass of 1 but got %d", cost)
	}
}

/**
Test against examples from the problem statement

    A module of mass 14 requires 2 fuel. This fuel requires no further fuel (2 divided by 3 and rounded down is 0, which would call for a negative fuel), so the total fuel required is still just 2.
    At first, a module of mass 1969 requires 654 fuel. Then, this fuel requires 216 more fuel (654 / 3 - 2). 216 then requires 70 more fuel, which requires 21 fuel, which requires 5 fuel, which requires no further fuel. So, the total fuel required for a module of mass 1969 is 654 + 216 + 70 + 21 + 5 = 966.
    The fuel required by a module of mass 100756 and its fuel is: 33583 + 11192 + 3728 + 1240 + 411 + 135 + 43 + 12 + 2 = 50346.
*/
func TestModuleFuelCost(t *testing.T) {
	cost := moduleFuelCost(14)
	if cost != 2 {
		t.Errorf("wanted 2 for mass of 14 but got %d", cost)
	}
	cost = moduleFuelCost(1969)
	if cost != 966 {
		t.Errorf("wanted 2 for mass of 966 but got %d", cost)
	}
	cost = moduleFuelCost(100756)
	if cost != 50346 {
		t.Errorf("wanted 50346 for mass of 100756 but got %d", cost)
	}
}
