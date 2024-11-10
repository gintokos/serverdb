NAME_COMMIT ?= 

PROTO_DIR = ./protos/proto

OUT_DIR = ./protos

PROTOC = protoc
PROTOC_GEN_GO = $(GOPATH)/bin/protoc-gen-go
PROTOC_GEN_GO_GRPC = $(GOPATH)/bin/protoc-gen-go-grpc

PROTO_FILES := $(wildcard $(PROTO_DIR)/*.proto)

MAIN_FILE = ./cmd/main/main.go 

.PHONY: all
all: $(PROTO_FILES) generate

.PHONY: generate
generate:
	@mkdir -p $(OUT_DIR)
	$(PROTOC) -I=$(PROTO_DIR) --go_out=$(OUT_DIR) --go-grpc_out=$(OUT_DIR) $(PROTO_FILES)

.PHONY: clean
clean:
	rm -rf $(OUT_DIR)/gen

.PHONY: run
run: 
	go run $(MAIN_FILE)

.PHONY: push
push:
	@if [ -z "$(NAME_COMMIT)"]; then echo "Needed NAME_COMMIT"; exit 1;fi
	git add .
	git commit -m "$(NAME_COMMIT)"
	git push -u origin main