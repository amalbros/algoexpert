package program

type SpecialArray []interface{}

func ProductSum(array SpecialArray) int{
	return helper(array,1)
}

func helper(array SpecialArray, multiplier int) int{
	sum := 0
	for _,element := range array{
		if cast, ok := element.(SpecialArray); ok{
			sum+= helper(cast,multiplier+1)
		} else if cast, ok := element.(int); ok{
			sum+= cast
		}
	}
	return sum * multiplier
}