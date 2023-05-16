BINARY_NAME=weather-go

build:
	go build -o ${BINARY_NAME} main.go
run:
	go build -o ${BINARY_NAME} main.go
	./${BINARY_NAME}
clean:
	go clean
