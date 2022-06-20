# ---------------------------------------------------------------------------------------------- Variables declaration #
SERVICE = arachne

# ------------------------------------------------------------------------------------------------------------- Docker #
.PHONY: docker/build:
docker/build:
	docker build -f dep/docker/Dockerfile -t $(SERVICE):0.1.0 .

.PHONY: docker/prune:
docker/prune:
	docker system prune -a --volumes

# ----------------------------------------------------------------------------------------------------- Docker-Compose #
.PHONY docker-compose/down:
docker-compose/down:
	docker-compose -f dep/docker/docker-compose.yaml -p $(SERVICE) down --volumes

.PHONY docker-compose/up:
docker-compose/log:
	docker-compose -f dep/docker/docker-compose.yaml logs -f --tail 10

.PHONY docker-compose/pull:
docker-compose/pull:
	docker-compose -f dep/docker/docker-compose.yaml pull

.PHONY docker-compose/up:
docker-compose/up:
	docker-compose -f dep/docker/docker-compose.yaml -p $(SERVICE) up -d

# ------------------------------------------------------------------------------------------------------------- Golang #
.PHONY go/install:
go/install:
	go mod tidy

.PHONY: go/run
go/run:
	go run server.go
