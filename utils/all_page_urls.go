package utils

import "fmt"

func (p *Pagination) Urls(baseURL string) map[int]string {
	urls := make(map[int]string)

	totalPages := TotalPages(p.PerPage, p.Total)
	for i := 1; i <= totalPages; i++ {
		url := fmt.Sprintf("%s?page=%d", baseURL, i)
		urls[i] = url
	}

	return urls
}
