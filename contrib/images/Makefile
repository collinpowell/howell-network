all: howelld-env

howelld-env:
	docker build --build-arg UID=$(shell id -u) --build-arg GID=$(shell id -g) --tag howell-network/howelld-env howelld-env

.PHONY: all howelld-env
