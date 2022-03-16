package itertools

// Reduce executes a user-supplied "reducer" callback function on each element of the array,
// in order, passing in the return value from the calculation on the preceding element.
// The final result of running the reducer across all elements of the array is a single value.
func Reduce[T any, R any](slice []T, iValue R, f func(prevValue R, currValue T, index int, slice []T) R) R {
	previousItem := iValue
	result := iValue
	for index := range slice {
		result = f(previousItem, slice[index], index, slice)
		previousItem = result
	}
	return result
}
