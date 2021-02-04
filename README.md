# goliday [![GoDoc](http://godoc.org/github.com/hlts2/goliday?status.svg)](http://godoc.org/github.com/hlts2/goliday)

Go library for japanese holiday

## Usage

```go

package main

import (
        "fmt"

        "github.com/hlts2/goliday"
)

func main() {
        holidays := goliday.Holidays(
                goliday.WithYear(2050),
                goliday.WithDay(11),
        )

        for _, holiday := range holidays {
                fmt.Println(holiday)
        }

        ok := goliday.HasHoliday(
                goliday.WithYear(2050),
                goliday.WithMonth(11),
                goliday.WithDay(23),
        )

        fmt.Println(ok) // true
}

```

## License

MIT

## Author

@hlts2
