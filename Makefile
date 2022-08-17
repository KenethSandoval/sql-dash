SHELL=/bin/bash

.PHONY: all

all: compile

clean: stop
	rm -rfv tuidb
	$(shell docker container rm $(shell docker ps -a --format "{{.ID}}"))

compile:
	@echo "Compiling..."
	go build -o tuidb

up:
	$(shell docker-compose up -d)

stop:
	$(shell docker-compose stop)


