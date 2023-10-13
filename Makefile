SHELL = /bin/bash
USER_ID ?= $(shell id -u)
GROUP_ID ?= $(shell id -g)
BUILD_IMAGE_SERVER  = golang:1.20
PROJECT_NAME        = "riskControl"
Image_NAME = "riskcontrol"
ifeq ($(TAGS_OPT),)
TAGS_OPT            = latest
else
endif

build:  build-server
	docker run -u $(USER_ID):$(GROUP_ID) --name build-local --rm -v $(shell pwd):/go/src/${PROJECT_NAME} -w /go/src/${PROJECT_NAME} ${BUILD_IMAGE_SERVER} make build-local

build-server:
	docker run  -e USER_ID=$(USER_ID)  -e GROUP_ID=$(GROUP_ID)  --name build-server-local --rm -v $(shell pwd):/go/src/${PROJECT_NAME} -w /go/src/${PROJECT_NAME} ${BUILD_IMAGE_SERVER} make build-server-local

build-local:
	if [ -d "build" ];then rm -rf build; else echo "build OK!"; fi \
	&& if [ -f "/.dockerenv" ];then echo "dockerenv OK!"; else  make build-server-local; fi \
	&& mkdir build && cp ./riskControl build/ && cp -r ./config.docker.yaml build/config.yaml


build-server-local:
	@if [ -f "riskControl" ];then rm -rf riskControl; else echo "riskControl OK!"; fi \
	&& go env -w GOPROXY=https://goproxy.cn,direct \
	&& go env -w CGO_ENABLED=0 && go env  && go mod tidy \
	&& git config --global --add safe.directory /go/src/riskControl\
	&& go build -ldflags "-B 0x$(shell head -c20 /dev/urandom|od -An -tx1|tr -d ' \n') -X main.Version=${TAGS_OPT}" -v  -o riskControl\
	&& chown -R $(USER_ID):$(GROUP_ID) ./riskControl

image: build
	docker build -t ${Image_NAME}:${TAGS_OPT} -f manifest/docker/Dockerfile .