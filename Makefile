SHELL=/bin/bash

.PHONY: all

all: compile

clean: stop
	rm -rfv sqldash
	$(shell docker container rm $(shell docker ps -a --format "{{.ID}}"))

compile:
	@echo "Compiling..."
	go build -o sqldash

up:
	$(shell docker-compose up -d)

stop:
	$(shell docker-compose stop)


