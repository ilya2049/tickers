# Tickers

## Skip ticker

``` go
ticker := tickers.NewExtraTickTicker(time.Second)

var i int

for range ticker.C() {
	fmt.Println(i)
	i++

if i == 3 || i == 4 {
		ticker.AddExtraTick()
	}

if i == 6 {
		break
	}
}
```