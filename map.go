package itertools

// Map converts every item in the provided slice based on the provided function and returns the result.
func Map[T any, R any](slice []T, f func(item T) R) []R {
	var result []R
	for _, item := range slice {
		result = append(result, f(item))
	}
	return result
}
