build:
	go build -o bin/rachio-next-run app/*.go

install:
	cp bin/rachio-next-run /usr/local/sbin/rachio-next-run

app-init:
	go generate
