package pageutil

import (
	"errors"
	"github.com/digitwires/go-paginate/utils"
)

func Paginate(data []interface{}, perPage int, currentPage int, baseUrl string) (*utils.Pagination, error) {
	if data == nil {
		return nil, errors.New("input data cannot be nil")
	}
	if perPage <= 0 {
		return nil, errors.New("items per page must be greater than zero")
	}
	if currentPage < 1 {
		return nil, errors.New("current page must be a positive integer")
	}

	start, end := utils.CalculateRange(len(data), perPage, currentPage)

	pagination := &utils.Pagination{
		CurrentPage: currentPage,
		PerPage:     perPage,
		Total:       len(data),
		TotalPages:  utils.TotalPages(perPage, len(data)),
		Data:        data[start:end],
		PrevPage:    utils.PrevPageURL(baseUrl, currentPage),
		NextPage:    utils.NextPageURL(baseUrl, currentPage),
	}

	if currentPage == 1 {
		pagination.PrevPage = ""
	}

	lastPage := utils.TotalPages(perPage, len(data))
	if currentPage == lastPage {
		pagination.NextPage = ""
	}

	return pagination, nil
}
