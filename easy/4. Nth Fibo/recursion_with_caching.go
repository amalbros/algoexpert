// in the last solution for recursion -
// https://github.com/PandaWhoCodes/algoexpert/blob/master/easy/4.%20Nth%20Fibo/recursion.go
// We see that we keep computing fib of a number again and again
// Let us now store a cache of the computed value so that only new values are computed
// And old values are loaded from cache

// This solution's time complexity will be O(2^n)
package program

func GetNthFib(n int) int {
	return ComputeFibo(n, map[int]int{
		1: 0,
		2: 1,
	})
}
func ComputeFibo(n int, cache map[int]int) int {
	if val, found := cache[n]; found {
		return val
	}
	cache[n] = ComputeFibo(n-1, cache) + ComputeFibo(n-2, cache)
	return cache[n]
	// we'll do
}
