# goliday_jp

Go library for japanese holiday

## Usage

```go

holidays := Holidays(WithYear(2050), WithDay(11))

for _, holiday := range holidays {
    fmt.Println(holiday)
}

fmt.Println(HasHoliday(WithYear(2050), WithMonth(11), WithDay(23))) // true

```

## License

MIT

## Author

@hlts2
