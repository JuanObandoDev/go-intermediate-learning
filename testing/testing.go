package main

func Sum(a, b int) int {
	return a + b
}

// func without test, it can affect our code coverage
func GetMax(a, b int) int {
	if a > b {
		return a
	}
	return b
}
