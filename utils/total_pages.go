package utils

func TotalPages(perPage, total int) int {
	var totalPages int

	if total%perPage == 0 {
		totalPages = total / perPage
	} else {
		totalPages = (total / perPage) + 1
	}

	return totalPages
}
