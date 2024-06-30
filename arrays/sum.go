package arrays

func Sum(numbers []int) int {
	sum := 0

	for _, number := range numbers {

		sum += number
	}

	return sum
}

// func main() {
// 	var numbers = [5]int{1, 2, 3, 4, 5}
// 	Sum(numbers)
// }
