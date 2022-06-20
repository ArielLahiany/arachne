# ------------------------------------------------------------------------------------------------------------- Golang #
.PHONY go/install:
go/install:
	go mod tidy

.PHONY: go/run
go/run:
	go run server.go
