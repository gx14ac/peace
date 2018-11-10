OS=linux
ARCH=amd64
ROOT=$(GOPATH)/src/github.com/OkumuraShintarou/peace

build:
	GOOS=$(OS) GOARCH=$(ARCH) CG0_ENABLED=0 go build -o bin/server -a -tags netgo -installsuffix netgo --ldflags '-extldflags "-static"' main.go

run:
	go get github.com/pilu/fresh
	fresh -c .fresh.conf
