/**
Finish the solution so that it sorts the passed in array of numbers.
If the function passes in an empty array or null/nil value then it should return an empty array.

For example:

solution(c(1, 2, 3, 10, 5)) # should return c(1, 2, 3, 5, 10)
solution(NULL)              # should return NULL
*/

package fundamentals

import "sort"

func SortNumbers(numbers []int) []int {
	if numbers == nil {
		return nil
	}
	// sort.Slice(numbers, func(i, j int) bool {
	// 	return numbers[i] < numbers[j]
	// })
	sort.Ints(numbers)
	return numbers
}
