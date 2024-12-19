package arraysslices

func Sum(numbers []int) int {
	var sum int
	for _, val := range numbers {
		sum += val
	}
	return sum
}

func SumAll(numbersToSum ...[]int) (sum []int) {
	for _, v := range numbersToSum {
		sum = append(sum, Sum(v))
	}
	return sum
}

func SumAllTails(numbersToSum ...[]int) []int {
	var tailedSlices [][]int
	for _, v := range numbersToSum {
		if len(v) == 0 {
			tailedSlices = append(tailedSlices, []int{0})
			continue
		}
		tailedSlices = append(tailedSlices, v[1:])
	}
	return SumAll(tailedSlices...)
}
