export GOPATH 		:= $(CURDIR)/_gopath
export GOBIN		:= $(CURDIR)/bin

CURRENT_GIT_GROUP 	:= github.com/litao44
CURRENT_GIT_REPO 	:= rpctest

all: build

dep_folders:
	mkdir -p $(CURDIR)/_gopath/src/$(CURRENT_GIT_GROUP)
	test -d $(CURDIR)/_gopath/src/$(CURRENT_GIT_GROUP)/$(CURRENT_GIT_REPO) || ln -s $(CURDIR) $(CURDIR)/_gopath/src/$(CURRENT_GIT_GROUP)

deps: dep_folders
	glide install

build:
	go build -o $(GOBIN)/server $(CURRENT_GIT_GROUP)/$(CURRENT_GIT_REPO)/server
	go build -o $(GOBIN)/client $(CURRENT_GIT_GROUP)/$(CURRENT_GIT_REPO)/client

pb:
	protoc --proto_path=./pb ./pb/*.proto --gogofast_out=plugins=grpc:./pb

clean:
	@rm -rf bin _gopath

.PHONY: deps test clean pb
