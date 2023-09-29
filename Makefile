PUBLIC = $(sort $(dir $(wildcard ./api/*/)))

SWAGGER = $(wildcard ./docs/*/*.json)

GOPATH = $(HOME)/go

.ONESHELL:

build:
	go build -v ./cmd/app

gen-pb:
	for public in $(PUBLIC) ; do \
		protoc -I ./api \
				-I$(GOPATH) \
				--go_out ./extra --go_opt paths=source_relative \
				--go-grpc_out=require_unimplemented_servers=false:./extra \
				--go-grpc_opt paths=source_relative \
				--grpc-gateway_out ./extra \
				--grpc-gateway_opt paths=source_relative \
				--openapiv2_out ./docs \
				--openapiv2_opt logtostderr=true \
				$$public/*.proto; \
	done \

gen-docs:
	for pb in $(SWAGGER) ; do \
  		redoc-cli bundle -o ./docs/generated/"$$(basename -s .json $$pb)".html $$pb; \
	done;