package utils

func Filter[T any](collection []T, test func(el T) bool) (filtered []T) {
	for _, el := range collection {
		if test(el) {
			filtered = append(filtered, el)
		}
	}
	return
}

func Map[T any, U any](collection []T, callback func(T) U) (result []U) {
	for _, el := range collection {
		result = append(result, callback(el))
	}
	return
}
