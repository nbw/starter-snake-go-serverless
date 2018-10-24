build:
	dep ensure -v
	env GOOS=linux go build -ldflags="-s -w" -o bin/hello hello/main.go
	env GOOS=linux go build -ldflags="-s -w" -o bin/ping ping/main.go
	env GOOS=linux go build -ldflags="-s -w" -o bin/start start/main.go
	env GOOS=linux go build -ldflags="-s -w" -o bin/move 	move/main.go
	env GOOS=linux go build -ldflags="-s -w" -o bin/end   end/main.go

.PHONY: clean
clean:
	rm -rf ./bin ./vendor Gopkg.lock

.PHONY: deploy
deploy: clean build
	sls deploy --verbose

.PHONY: get
get:
	@echo ">> Getting any missing dependencies.."
	go get -t ./...

.PHONY: test
test:
	go test ./...
