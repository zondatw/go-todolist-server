GOCMD = go
GOBUILD = $(GOCMD) build
GOTEST = $(GOCMD) test
GOCLEAN = $(GOCMD) clean
DOCKERHUB = zondayang
NAMESPACE = todolist_backend
BINARY_NAME = go-todolist-server

all: test build

build:
	$(GOBUILD) -o $(BINARY_NAME) -v

test:
	$(GOTEST) -v ./...

docker:
	@docker build --tag docker.io/$(DOCKERHUB)/$(NAMESPACE):latest .
	@docker push docker.io/$(DOCKERHUB)/$(NAMESPACE):latest

.PHONY: clean
clean:
	$(GOCLEAN)
	rm $(BINARY_NAME)