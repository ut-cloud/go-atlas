GOHOSTOS:=$(shell go env GOHOSTOS)
GOPATH:=$(shell go env GOPATH)
VERSION=$(shell git describe --tags --always)
OUTPUT_FILE=api/core/v1/operation.gen.go  # operation 文件生成路径

ifeq ($(GOHOSTOS), windows)
	#the `find.exe` is different from `find` in bash/shell.
	#to see https://docs.microsoft.com/en-us/windows-server/administration/windows-commands/find.
	#changed to use git-bash.exe to run find cli or other cli friendly, caused of every developer has a Git.
	#Git_Bash= $(subst cmd\,bin\bash.exe,$(dir $(shell where git)))
	Git_Bash=$(subst \,/,$(subst cmd\,bin\bash.exe,$(dir $(shell where git))))
	INTERNAL_PROTO_FILES=$(shell $(Git_Bash) -c "find internal -name *.proto")
	API_PROTO_FILES=$(shell $(Git_Bash) -c "find api -name *.proto")
else
	INTERNAL_PROTO_FILES=$(shell find internal -name *.proto)
	API_PROTO_FILES=$(shell find api -name *.proto)
endif

.PHONY: init
# init env
init:
	go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
	go install github.com/go-kratos/kratos/cmd/kratos/v2@latest
	go install github.com/go-kratos/kratos/cmd/protoc-gen-go-http/v2@latest
	go install github.com/google/gnostic/cmd/protoc-gen-openapi@latest
	go install github.com/google/wire/cmd/wire@latest

.PHONY: config
# generate internal proto
config:
	protoc --proto_path=./internal \
	       --proto_path=./third_party \
 	       --go_out=paths=source_relative:./internal \
	       $(INTERNAL_PROTO_FILES)

.PHONY: api
# generate api proto
api:
	protoc --proto_path=./api \
	       --proto_path=./third_party \
 	       --go_out=paths=source_relative:./api \
 	       --go-http_out=paths=source_relative:./api \
 	       --go-grpc_out=paths=source_relative:./api \
	       --openapi_out=fq_schema_naming=true,default_response=false:. \
	       $(API_PROTO_FILES)
	./scandir.sh api | xargs -n1 -IX bash -c 'sed s/,omitempty// X > X.tmp'
	./scandir.sh api | xargs -n1 -IX bash -c 'mv X{.tmp,}'
	make scan_operations;

.PHONY: scan_operations
# 扫描以 Operation 开头的常量的目标
scan_operations:
	@echo "Scanning for Operation constants..."
	@echo "// Code generated by makefile. DO NOT EDIT." > $(OUTPUT_FILE)
	@echo "" >> $(OUTPUT_FILE)
	@echo "package v1" >> $(OUTPUT_FILE)
	@echo "" >> $(OUTPUT_FILE)
	@echo "func GetAllOperations() []string {" >> $(OUTPUT_FILE)  #添加函数声明
	@echo "    return []string{" >> $(OUTPUT_FILE)
	@find ./api -name "*.go" | while read -r file; do \
		grep -E 'const Operation[A-Za-z0-9]*' "$$file" | while read -r line; do \
			const_name=$$(echo "$$line" | awk '{print $$2}'); \
			const_value=$$(echo "$$line" | awk -F ' = ' '{print $$2}'); \
			echo "        $$const_name, //$$const_value" >> $(OUTPUT_FILE); \
		done; \
	done
	@echo "    }" >> $(OUTPUT_FILE)  # 结束 const 声明
	@echo "}" >> $(OUTPUT_FILE)  # 结束 const 声明
	@echo "Operation file gen success"

.PHONY: build
# build
build:
	mkdir -p bin/ && go build -ldflags "-X main.Version=$(VERSION)" -o ./bin/ ./...

.PHONY: generate
# generate
generate:
	go generate ./...
	go mod tidy

.PHONY: all
# generate all
all:
	make api;
	make config;
	make generate;

# show help
help:
	@echo ''
	@echo 'Usage:'
	@echo ' make [target]'
	@echo ''
	@echo 'Targets:'
	@awk '/^[a-zA-Z\-\_0-9]+:/ { \
	helpMessage = match(lastLine, /^# (.*)/); \
		if (helpMessage) { \
			helpCommand = substr($$1, 0, index($$1, ":")); \
			helpMessage = substr(lastLine, RSTART + 2, RLENGTH); \
			printf "\033[36m%-22s\033[0m %s\n", helpCommand,helpMessage; \
		} \
	} \
	{ lastLine = $$0 }' $(MAKEFILE_LIST)

.DEFAULT_GOAL := help

wire:
	cd cmd/atlas-core && wire