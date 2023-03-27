package main

import "github.com/digitwires/go-paginate/utils"

func Paginate(data []interface{}, perPage int, currentPage int, baseUrl string) *utils.Pagination {
	start := (currentPage - 1) * perPage
	end := start + perPage

	if end > len(data) {
		end = len(data)
	}

	pagination := &utils.Pagination{
		CurrentPage: currentPage,
		PerPage:     perPage,
		Total:       len(data),
		Data:        data[start:end],
	}

	if currentPage > 1 {
		pagination.PrevPage = utils.PrevPageURL(baseUrl, currentPage)
	}

	if currentPage < utils.TotalPages(perPage, len(data)) {
		pagination.NextPage = utils.NextPageURL(baseUrl, currentPage)
	}

	return pagination
}
