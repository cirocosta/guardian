image: build
	docker build -t a . 


build: ./bin/init ./bin/dadoo ./bin/gdn


start-jaeger:
	docker run -d --name jaeger \
	  -e COLLECTOR_ZIPKIN_HTTP_PORT=9411 \
	  -p 5775:5775/udp \
	  -p 6831:6831/udp \
	  -p 6832:6832/udp \
	  -p 5778:5778 \
	  -p 25000:16686 \
	  -p 14268:14268 \
	  -p 9411:9411 \
	  jaegertracing/all-in-one:1.12


./bin/init:
	gcc -static -o $@ ./cmd/init/init.c ./cmd/init/ignore_sigchild.c

./bin/dadoo:
	go build -v -o $@ ./cmd/dadoo 

./bin/gdn:
	go build -v -o $@ ./cmd/gdn
.PHONY: ./bin/gdn

