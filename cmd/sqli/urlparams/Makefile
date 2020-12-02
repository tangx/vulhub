
docker: build
	docker build --tag doslab/vulhub-sqli . \
		&& docker push doslab/vulhub-sqli

build:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o sqli .
