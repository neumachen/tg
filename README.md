![main](https://github.com/neumachen/tg/actions/workflows/ci.yml/badge.svg?branch=main)
![ci](https://github.com/neumachen/tg/actions/workflows/ci.yml/badge.svg)
[![GoDoc](https://img.shields.io/badge/godoc-reference-blue.svg)](https://pkg.go.dev/github.com/neumachen/tg)

# Package tg

Package tg is a Go package that provides functionality for generating random strings.

## Installation

To install the package, use the following command:

```
go get github.com/your-username/your-module/tg
```

## Usage

Import the package in your Go code:

```go
import "github.com/neumachen/tg"
```

### RandStringBytes

The `RandStringBytes` function generates a random string of the specified length. It uses a predefined set of characters for generating the string.

```go
func RandStringBytes(n int) string
```

- `n` specifies the length of the generated string.
- Returns a random string of length `n`.

#### Example

```go
package main

import (
	"fmt"
	"github.com/your-username/your-module/tg"
)

func main() {
	randomString := tg.RandStringBytes(10)
	fmt.Println(randomString)
}
```

This example generates a random string of length 10 and prints it to the console.

## Contributing

Contributions are welcome! If you find any issues or have suggestions for improvements, please open an issue or submit a pull request.

## License

This package is open source and available under the [GNU Affero General Public License v3.0 (AGPL-3.0)](LICENSE).
