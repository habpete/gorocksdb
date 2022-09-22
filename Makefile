.PHONY: docker-build
docker-build:
	docker-compose up build

.PHONY: build
build:
	sh ./ext/rocksdb.sh &&