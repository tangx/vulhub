
GOBUILD := CGO_ENABLED=0 GOOS=linux GOARCH=amd64

build:
	$(GOBUILD) go build -o reflect2 .

docker: build
	docker build -t doslab/reflect2:static --build-arg from=gcr.io/distroless/static-debian10 .	\
		&& docker push doslab/reflect2:static

	docker build -t doslab/reflect2 --build-arg from=alpine:3.12 . \
		&& docker push doslab/reflect2

