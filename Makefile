
GO111MODULE=on
INSTALLDIR=/usr/local/bin

build: fmt
	go build -a -tags "netgo static" \
		-ldflags '-w -extldflags "-static"' \
		-o samtracker/samtracker ./samtracker

fmt:
	gofmt -w *.go samtracker/*.go

try: build
	mkdir -p tmp
	cd tmp && ../samtracker/samtracker

tag:
	cat changelog | gothub release -p -u eyedeekay -r samtracker -t $(VERSION) -n $(VERSION) -d -; true

upload:
	gothub upload -R -u eyedeekay -r samtracker -t $(VERSION) -n "samtracker.tar.gz" -f "./samtracker.tar.gz"

tar:
	tar --exclude=.git --exclude=samtracker.tar.gz -cvf ./samtracker.tar.gz .

install:
	install -m755 samtracker/samtracker $(INSTALLDIR)bin
