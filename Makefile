.PHONE:

build:
	# go build -o main cmd/main.go
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -tags netgo -ldflags '-w -extldflags "-static"' -o main ./cmd/main.go
run: build
	./main
