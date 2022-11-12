FROM golang:1.19.1-bullseye

RUN apt-get update && \
    apt-get install -y --no-install-recommends \
        protobuf-compiler && \
    apt-get -y clean && \
    rm -rf /var/lib/apt/lists/* && \
    go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.26 && \
    go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.1

WORKDIR /janken_go_grpc
CMD ["/bin/bash"]