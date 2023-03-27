package utils

type Pagination struct {
	CurrentPage int
	PerPage     int
	Total       int
	Data        interface{}
	PrevPage    string
	NextPage    string
}
