package utils

func CalculateRange(dataLength int, perPage int, currentPage int) (start int, end int) {
	start = (currentPage - 1) * perPage
	if start >= dataLength {
		start = dataLength
	}

	end = start + perPage
	if end > dataLength {
		end = dataLength
	}

	return start, end
}