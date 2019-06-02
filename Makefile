image: build
	docker build -t concourse . 


build: ./bin/init ./bin/dadoo ./bin/gdn


run:
	docker-compose up --build -d


./bin/init:
	gcc -static -o $@ ./cmd/init/init.c ./cmd/init/ignore_sigchild.c

./bin/dadoo:
	go build -v -o $@ ./cmd/dadoo 

./bin/gdn:
	go build -i -v -o $@ ./cmd/gdn
.PHONY: ./bin/gdn

