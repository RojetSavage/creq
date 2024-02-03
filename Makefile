clean: 
	rm bin/creq

build: 
	go build -o bin/creq pkg/main.go

run:
	go run pkg/main.go