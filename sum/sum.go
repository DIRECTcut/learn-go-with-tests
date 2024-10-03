package sum

func Sum(input []int) int {
	result := 0
	for _, value := range input {
		result += value
	}

	return result
}

func SumAll(numbersToSum ...[]int) []int {
	var res []int

	for _, slice := range numbersToSum {
		res = append(res, Sum(slice))
	}

	return res
}

func SumAllTails(slicesToSum ...[]int) []int {
	res := []int{}

	for _, slice := range slicesToSum {
		if len(slice) == 0 {
			res = append(res, 0)
			continue
		}

		tail := slice[1:]
		res = append(res, Sum(tail))
	}

	return res
}
