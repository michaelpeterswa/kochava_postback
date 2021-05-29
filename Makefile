all:
	./scripts/clean_containers.sh
	./scripts/build_containers.sh
	./scripts/start_containers.sh

help:
	@echo "build, clean, or start"

build:
	./scripts/build_containers.sh

clean:
	./scripts/clean_containers.sh

start:
	./scripts/start_containers.sh