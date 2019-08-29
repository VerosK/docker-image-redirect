CONTAINER_NAME=wftech/http-redirect
CONTAINER_VERSION=latest

http-redirect: http-redirect.go
	CGO_ENABLED=0 GOOS=linux go build -a -tags netgo -ldflags '-w' http-redirect.go

	docker build -t ${CONTAINER_NAME}:${CONTAINER_VERSION}  .
