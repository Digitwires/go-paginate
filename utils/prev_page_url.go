package utils

import "fmt"

func PrevPageURL(baseUrl string, currentPage int) string {
	var prevPage int

	if currentPage > 1 {
		prevPage = currentPage - 1
	} else {
		prevPage = 1
	}

	prevPageURL := fmt.Sprintf("%s?page=%d", baseUrl, prevPage)

	return prevPageURL
}
