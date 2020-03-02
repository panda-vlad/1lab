TEMP = $(shell git describe)


default: out/example
clean:
	rm -rf out
test:
	go vet && go test
out/example: implementation.go cmd/example/main.go
	rm buildVersion.go
	echo "package main  \n\nconst(buildVersion=$(TEMP))" >> buildVersion.go 
	mkdir -p out
	go build -o out/example ./cmd/example 