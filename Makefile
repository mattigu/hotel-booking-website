PROJECT_DIR := $(CURDIR)

all: docker_up

install_deps:
	sudo apt-get update && \
	sudo apt install -y docker docker-compose postgresql

docker_up:
	docker-compose up --build

docker_down:
	docker-compose down

clean_docker:
	docker-compose down --volumes --remove-orphans
	docker system prune --all --force --volumes