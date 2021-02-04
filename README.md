# goliday_jp [![GoDoc](http://godoc.org/github.com/hlts2/goliday_jp?status.svg)](http://godoc.org/github.com/hlts2/goliday_jp)

Go library for japanese holiday

## Usage

```go

package main

import (
        "fmt"

        goliday "github.com/hlts2/goliday_jp"
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
