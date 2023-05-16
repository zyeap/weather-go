# weather-go
CLI app to display weather information, written in go
![demo_screenshot](https://github.com/zyeap/weather-go/assets/39923168/90f3ef41-035e-4802-aea1-12412dfa7b3c)

## Prerequisites
`weather-go` utilizes the free Weather API, more info can be found here: https://www.weatherapi.com/  

To run `weather-go`, please have a valid Weather API token.

Set up the token for usage with `weather-go` by running the following command:  
```
export WEATHER_API_TOKEN=<token value here>
```
## Building weather-go
Any of the following commands can be used to build weather-go
```
1. make
2. go build
```
You should see a weather-go binary file within the current directory
```
> ls
assets  go.mod  go.sum  main.go  Makefile  weather-go
```
## Running weather-go

Running the binary (default location is Houston)
```
> ./weather-go
───────────────────────────────────────────────────────────────────────────
Houston, Texas, United States of America: 75F, 24C
Today's condition: Partly cloudy
───────────────────────────────────────────────────────────────────────────
[17:00 - 79F, 26C] - [62% rain] - [Patchy rain possible]
[18:00 - 78F, 26C] - [0% rain] - [Partly cloudy]
[19:00 - 77F, 25C] - [0% rain] - [Clear]
[20:00 - 77F, 25C] - [0% rain] - [Partly cloudy]
[21:00 - 76F, 24C] - [76% rain] - [Patchy rain possible]
```

Running via `go run main.go`
```
> go run main.go -l san_jose
───────────────────────────────────────────────────────────────────────────
San Jose, California, United States of America: 80F, 27C
Today's condition: Partly cloudy
───────────────────────────────────────────────────────────────────────────
[17:00 - 87F, 31C] - [0% rain] - [Sunny]
[18:00 - 81F, 27C] - [0% rain] - [Sunny]
[19:00 - 78F, 26C] - [0% rain] - [Sunny]
[20:00 - 66F, 19C] - [0% rain] - [Sunny]
[21:00 - 64F, 18C] - [0% rain] - [Clear]
[22:00 - 64F, 18C] - [0% rain] - [Clear]
[23:00 - 64F, 18C] - [0% rain] - [Clear]
```
