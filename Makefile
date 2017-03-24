all:
	go build -o demo

binary:
	docker run -ti --rm  -v `pwd`:/go/src/habor-template-demo -w /go/src/habor-template-demo iron/go:1.7-dev go build -o demo

docker-build: binary
	docker build -t demo .

.PHONY: binary
