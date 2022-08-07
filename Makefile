.PHONY: gqlgen up build

img_name = eu.gcr.io/vediagames/onlooker
version = latest
env_file = ./.env

include $(env_file)
export $(shell sed 's/=.*//' $(env_file))
PATH := $(PATH):$(GOPATH)/bin

setup:
	git config --global --add url."git@github.com:".insteadOf "https://github.com/"
	go env -w GOPRIVATE=github.com/vediagames/*

build:
	@docker build -f ./build/Dockerfile -t $(img_name):$(version) --build-arg GITHUB_TOKEN=$(GITHUB_TOKEN) .

swag/fmt:
	swag fmt

swag/init:
	swag init

migrate/new/%:
	@migrate create -ext sql -dir ./db/schema/ -seq $*.sql

migrate/up:
	migrate -database ${POSTGRES_CONN_STRING} -path db/schema  up