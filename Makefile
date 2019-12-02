
build: fmt
	go build -o samtracker/samtracker ./samtracker

fmt:
	gofmt -w *.go

try: build
	mkdir -p tmp
	cd tmp && ../tracker