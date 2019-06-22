package program

// fibo series 0,1,1,2,3,5,8,13,21 ... 
func GetNthFib(n int) int {
	if n == 0 {
		return 0
	} else if n == 2 {
		return 1 // the second number in the fibonacci series is always 1
	} else if n == 1 {
		return 0 // the first number in the fibonacci series is always 0
	} else {
		return GetNthFib(n-1) + GetNthFib(n-2)
	}
}
