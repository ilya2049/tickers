# Tickers

## Skip ticker

``` go
ticker := tickers.NewSkipTicker(time.Second)

var i int

for range ticker.C() {
	fmt.Println(i)
	i++

	if i > 5 && i < 15 && i != 10 {
		ticker.Skip()
	}

	if i == 20 {
		ticker.Stop()
	}
}
```