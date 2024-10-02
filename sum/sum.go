package sum

func Sum(input []int) int {
	result := 0
	for _, value := range input {
		result += value
	}

	return result
}
