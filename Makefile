PROJECT=github.com/olivere/greeter

all: proto build

.PHONY: proto build

proto:
	protoc -I ./proto/greeter proto/greeter/greeter.proto --go_out=plugins=micro:./proto/greeter

build:
	go build -o greeter $(PROJECT)
