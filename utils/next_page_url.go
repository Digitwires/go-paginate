package utils

import "fmt"

func NextPageURL(baseURL string, currentPage int) string {
	var nextPage int

	if currentPage == 0 {
		nextPage = 2
	} else {
		nextPage = currentPage + 1
	}

	return fmt.Sprintf("%s?page=%d", baseURL, nextPage)
}
