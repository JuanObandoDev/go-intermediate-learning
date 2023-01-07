package main

import "testing"

func TestSum(t *testing.T) {
	total := Sum(5, 5)

	if total != 10 {
		t.Errorf("Sum was incorrect, got: %d, want: %d.", total, 10)
	}

	// other way to do it
	tables := []struct {
		x int
		y int
		z int
	}{
		{1, 1, 2},
		{2, 2, 4},
		{3, 3, 6},
	}

	for _, item := range tables {
		total := Sum(item.x, item.y)
		if total != item.z {
			t.Errorf("Sum of (%d+%d) was incorrect, got: %d, want: %d.", item.x, item.y, total, item.z)
		}
	}
}

func TestGetMax(t *testing.T) {
	tables := []struct {
		x int
		y int
		z int
	}{
		{1, 0, 1},
		{2, 1, 2},
		{1, 2, 2},
	}

	for _, item := range tables {
		max := GetMax(item.x, item.y)
		if max != item.z {
			t.Errorf("GetMax of (%d,%d) was incorrect, got: %d, want: %d.", item.x, item.y, max, item.z)
		}
	}
}

func TestFibonacci(t *testing.T) {
	tables := []struct {
		x int
		y int
	}{
		{1, 1},
		{8, 21},
		{50, 12586269025},
	}

	for _, item := range tables {
		fib := Fibonacci(item.x)
		if fib != item.y {
			t.Errorf("Fibonacci of (%d) was incorrect, got: %d, want: %d.", item.x, fib, item.y)
		}
	}
}
