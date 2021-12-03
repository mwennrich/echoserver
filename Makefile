GO111MODULE := on
DOCKER_TAG := $(or ${GIT_TAG_NAME}, latest)

all: echoserver

.PHONY: echoserver
echoserver:
	go build -tags netgo -o bin/echoserver *.go
	strip bin/echoserver

.PHONY: dockerimages
dockerimages:
	docker build -t mwennrich/echoserver:${DOCKER_TAG} .

.PHONY: dockerpush
dockerpush:
	docker push mwennrich/echoserver:${DOCKER_TAG}

.PHONY: clean
clean:
	rm -f bin/*
