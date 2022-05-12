SHELL=/bin/bash

.PHONY: all

all: compile
clean:
	rm -rv -- sqldash

compile:
	@echo "Compiling..."
	go build -o sqldash

up:
	$(shell docker-compose up -d)

stop:
	$(shell docker-compose stop)


