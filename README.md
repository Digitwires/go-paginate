[![Go Reference](https://pkg.go.dev/badge/github.com/digitwires/go-paginate.svg)](https://pkg.go.dev/github.com/digitwires/go-paginate)
[![Go Report Card](https://goreportcard.com/badge/github.com/digitwires/go-paginate)](https://goreportcard.com/report/github.com/digitwires/go-paginate)

# go-paginate

`go-paginate` is a simple Go package that adds pagination functionality to slice of data.

## Installation

To install `go-paginate`, use `go get`:

```go
go get github.com/digitwires/go-paginate
```


## Usage

The `go-paginate` package provides a `Paginate` function that takes in a slice of data, the current page number, and the desired page size. It returns a new slice that contains only the data for the current page, as well as information about the total number of pages and the total number of items.

Here's an example:

```go
package main

import (
	"fmt"

	pageutil "github.com/digitwires/go-paginate"
)

func main() {
	data := []interface{}{"apple", "banana", "cherry", "date", "elderberry", "fig", "grape", "honeydew", "kiwi"}
	perPage := 3
	currentPage := 2
	baseUrl := "http://example.com"

	paginate, err := pageutil.Paginate(data, perPage, currentPage, baseUrl)
	if err != nil {
        fmt.Println(err)
    }

	fmt.Println("Current Page:", paginate.CurrentPage)
	fmt.Println("Total Pages:", paginate.Total)
	fmt.Println("Data:", paginate.Data)
	fmt.Println("Previous Page:", paginate.PrevPage)
	fmt.Println("Next Page:", paginate.NextPage)
}
```

Output:

```text
Current Page: 1
Total Pages: 3
Data: [apple banana cherry]
Previous Page: http://example.com?page=1
Next Page: http://example.com?page=3
```

## License

`go-paginate` is licensed under the MIT License. See [LICENSE](LICENSE) for the full license text.

## Contributing

Contributions are welcome! If you find a bug or have a feature request, please open an issue on GitHub. If you'd like to contribute code, please fork the repository and submit a pull request.

```text
Note that you can copy and paste this code into a `README.md` file in the root of your repository to use it as your package's documentation.
```

## Author

`go-paginate` was written by [Ahmed M. Ibrahim](https://github.com/2hmad).