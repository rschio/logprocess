.PHONY: all

all: docker-up build

.PHONY: build
build:
	rm -rf ./bin
	cd processor/cmd/ && go build -o ../../bin/processor

.PHONY: docker-up
docker-up:
	docker-compose up -d

.PHONY: clean
clean:
	rm -rf ./bin
	docker-compose down

.PHONY: destroy
destroy:
	sudo rm -rf ./storage/mysql/data/
