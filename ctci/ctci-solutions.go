package main

func main() {
	println(countSteps(1))
	println(countSteps(2))
	println(countSteps(3))
	println(countSteps(4))
	println(countSteps(5))
	println(countSteps(6))
	println(countSteps(7))
	println(countSteps(8))
	println(countSteps(9))
}

// countSteps
// triple step: a child is running up a  staircase with n steps and can hop either 1 2 or 3 steps.
// implement a method to count how many posible ways the child can run up the stairs
// 1
// 2
// 4
// 10
// 19
// 36
// 68
// 126
// 233

// start at the end. we can get there from three other steps + # of ways we can get to those 3 steps.
// transform to memoized version --> satisfy dependencies first (topologically sort the recursive dependency graph in your head).
// define base case: 1 step has 1 way, 2 steps has 2 ways, 3 steps has 4 ways (1 + 2 + a three step jump (1) )
// memoize and count.
func countSteps(n int) int {
	steps := []int{
		1, 2, 4,
	}
	for i := 3; i < n; i++ {
		steps = append(steps, 3+steps[i-3]+steps[i-2]+steps[i-1])
	}

	return steps[n-1]
}
