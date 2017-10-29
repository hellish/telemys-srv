run:
	go run common.go main.go

build:
	go build -o bin/telemys-server common.go main.go

win:
	GOOS=windows GOARCH=386 go build -o bin/telemys-server.exe common.go main.go