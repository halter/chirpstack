.PHONY: rust grpc-web go js python md java kotlin csharp php c fix-go-mod

all: rust grpc-web go js python md java kotlin csharp php c

fix-go-mod:
	grep -l -r 'github.com/halter/chirpstack' . | xargs -n1 gsed -i 's#github.com/halter/chirpstack#github.com/halter/chirpstack#g'

rust:
	cd rust && make

grpc-web:
	cd grpc-web && make

go: fix-go-mod
	cd go && make

js:
	cd js && make

python:
	docker compose run --rm chirpstack-api-python

md:
	docker compose run --rm chirpstack-api-md

java:
	docker compose run --rm chirpstack-api-java

kotlin:
	docker compose run --rm chirpstack-api-kotlin

csharp:
	docker compose run --rm chirpstack-csharp

php:
	docker compose run --rm chirpstack-api-php

c:
	docker compose run --rm chirpstack-api-c
