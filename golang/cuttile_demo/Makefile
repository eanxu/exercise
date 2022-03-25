.PHONY: build clean api statics
VERSION := $(shell git describe --always |sed -e "s/^v//")

linux:export GOOS=linux
linux:export GOARCH=amd64
linux:export GODEBUG=cgocheck=0

build: clean statics
	@echo "Compiling source"
	@rm -rf build
	@mkdir -p build
	go build -o build/drem main.go

linux: build

microservice: linux
#	@echo "docker images"
#	@docker image rm bjsh/drem:v1.0.0 | true
#	@docker image rm k8shub.com:1180/bjsh/drem:v1.0.0 | true
#	@docker build  -t bjsh/drem:v1.0.0 .
#	@docker images bjsh/drem:v1.0.0
	@echo "docker images"
	@docker image rm geovis/drem:v1.0.0 | true
	@docker build -t geovis/drem:v1.0.0 .

clean:
	@echo "Cleaning up workspace"
	@rm -rf build

saveDocker:
	@echo "save docker"
	@docker save -o /home/marshmallow/code/Docker/drem.tar geovis/drem:v1.0.0
