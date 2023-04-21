// instance-gen: File auto generated -- DO NOT EDIT!!!
build:
	go fmt ./...
	go vet ./...
	go build -o bin/rachio-next-run app/*.go

install:
	cp bin/rachio-next-run /usr/local/sbin/rachio-next-run

golib-latest:
	go get -u github.com/skeletonkey/lib-instance-gen-go@latest

	go mod tidy

app-init:
	go generate
