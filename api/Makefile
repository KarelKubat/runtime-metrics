PROTO_SRC = $(wildcard *.proto)
PROTO_GEN = $(patsubst %.proto, %.pb.go, $(PROTO_SRC))

all: $(PROTO_GEN)

%.pb.go: %.proto
	protoc -I=/usr/local/include:. --go_out=plugins=grpc:. $<

