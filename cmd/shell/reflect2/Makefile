
GOBUILD := CGO_ENABLED=0 GOOS=linux GOARCH=amd64

build:
	$(GOBUILD) go build -o reflect2 .

docker: build
	docker build -t doslab/vulhub-reflect2:static --build-arg from=gcr.io/distroless/static-debian10 .	\
		&& docker push doslab/vulhub-reflect2:static

	docker build -t doslab/vulhub-reflect2 --build-arg from=debian:buster . \
		&& docker push doslab/vulhub-reflect2

