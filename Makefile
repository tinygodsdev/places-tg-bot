ifneq ("$(wildcard env/.env.local)","")
    include env/.env.local
    export $(shell sed 's/=.*//' env/.env.local)
endif

.PHONY: run
run:
	go run cmd/main.go

.PHONY: build
build:
	go build -o bin/main cmd/main.go

.PHONY: test
test:
	go test -v ./...


.PHONY: deploy
deploy:
	git push dokku main

.PHONY: deps
deps:
	go get github.com/tinygodsdev/datasdk/pkg/server
	go mod tidy
