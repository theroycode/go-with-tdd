package main

func Sum(numbers []int) (sum int) {
	for _, num := range numbers {
		sum += num
	}
	return
}

func SumAll(numbersToSum ...[]int) []int {
	var sum []int
	for _, num := range numbersToSum {
		sum = append(sum, Sum(num))
	}
	return sum
}

func SumAllTails(tailsToSum ...[]int) []int {
	var sum []int
	for _, tail := range tailsToSum {
		if len(tail) == 0 {
			sum = append(sum, 0)
		} else {
			sum = append(sum, Sum(tail[1:]))
		}
	}
	return sum
}
