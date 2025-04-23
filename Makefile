build:
	mkdir -p bin
	go build -o ./bin/LD-8 ./src/main.go
run:
	go run src/main.go
clear:
	rm -rf bin
test:
	go test -v tests/*
build-windows:
	GOOS=windows GOARCH=amd64 go build -o LD-8.exe src/main.go