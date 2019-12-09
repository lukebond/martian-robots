NAME:=martian-robots
GITREPO:=github.com/lukebond/martian-robots
GITCOMMIT:=$(shell git describe --dirty --always)
VERSION:=0.1-dev
COMPILE_ARGS:=-buildmode pie

.PHONE: build
build:
	GOOS=linux go build -a -installsuffix cgo $(COMPILE_ARGS) -o martian-robots ./cmd

.PHONY: test
test:
	go test github.com/lukebond/martian-robots/pkg/{grid,instructions,robot}
