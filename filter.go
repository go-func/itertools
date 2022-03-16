package itertools

// Filter will return elements which return true based on the provided filtering functions.
func Filter[T any](slice []T, f func(item T) bool) []T {
	var result []T
	for _, item := range slice {
		if f(item) {
			result = append(result, item)
		}
	}
	return result
}
