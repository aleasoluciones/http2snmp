all: clean build

update_dep:
	go get $(DEP)
	go mod tidy

update_all_deps:
	go get -u
	go mod tidy

format:
	go fmt ./...

build:
	go build http2snmp.go

clean:
	rm -f http2snmp


.PHONY: update_dep update_all_deps format build clean
