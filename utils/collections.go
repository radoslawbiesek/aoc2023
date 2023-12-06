package utils

func Map[T any, U any](collection []T, callback func(T) U) (result []U) {
	for _, el := range collection {
		result = append(result, callback(el))
	}
	return
}
