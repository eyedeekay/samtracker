
GO111MODULE=on

build: fmt
	go build -a -tags "netgo static" \
		-ldflags '-w -extldflags "-static"' \
		-o samtracker/samtracker ./samtracker

fmt:
	gofmt -w *.go

try: build
	mkdir -p tmp
	cd tmp && ../samtracker/samtracker