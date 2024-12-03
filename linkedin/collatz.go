// collatz.go

package main

func CollatzStep(n int) int {
	if n%2 == 0 {
		return n / 2
	}
	return n*3 + 1
}
