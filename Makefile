ifeq ($(OS),Windows_NT)
SHELL := powershell.exe
.SHELLFLAGS := -NoProfile -Command
endif

.DEFAULT_GOAL := docker-push-image

# local dev
test:
	go test ./...
.PHONY:test

build: test
	go build main.go
.PHONY:build

run-with-env: build
	$$env:SERVER_PORT='8080'; ./main
.PHONY:run-with-env

# clear local dev
clear:
	rm main.exe
.PHONY:clear

# docker
docker-build-image:
	docker build -t stakkato95/twitter-service-users:0.1.1 . -f Dockerfile
.PHONY:docker-build-image

docker-push-image: docker-build-image
	docker push stakkato95/twitter-service-users:0.1.1
.PHONY:docker-push-image

docker-run-tmp-container: docker-build-image
	docker run --rm -p 8080:8080 -d stakkato95/twitter-service-users
.PHONY:docker-local-container
