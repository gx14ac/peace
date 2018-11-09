OS=linux
ARCH=amd64
ROOT=$(GOPATH)/src/github.com/OkumuraShintarou/peace

run:
	go get github.com/pilu/fresh
	fresh -c .fresh.conf
