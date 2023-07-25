#!/usr/bin/make

.PHONY: help
.DEFAULT_GOAL := help

env: ##Create environment file
	@if ! test -f "./.env"; then echo -e "Make: Copying env file.\n" && cp ./.env.example ./.env ; fi

build: ##Build containers
build: env
	@docker-compose build

help:
	@fgrep -h "##" $(MAKEFILE_LIST) | fgrep -v fgrep | sed -e 's/\\$$//' | sed -e 's/##//'