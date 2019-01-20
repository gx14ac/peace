OS=linux
ARCH=amd64

build:
	GOOS=$(OS) GOARCH=$(ARCH) CGO_ENABLED=1 go build -o bin/server -a -tags netgo -installsuffix netgo --ldflags '-linkmode external -extldflags -static' main.go


run:
	go get github.com/pilu/fresh
	fresh -c .fresh.conf
