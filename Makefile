build:
	go build -o LD-8 src/main.go
run:
	go run src/main.go
clear:
	rm LD-8 && rm LD-8.exe
test:
	go test -v tests/*
build-windows:
	GOOS=windows GOARCH=amd64 go build -o LD-8.exe src/main.go