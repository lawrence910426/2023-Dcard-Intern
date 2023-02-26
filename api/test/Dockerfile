FROM golang:1.20

# Install build proto deps
ENV PROTOC_ZIP=protoc-3.13.0-linux-x86_64.zip
RUN apt-get update && apt-get install -y unzip
RUN curl -OL https://github.com/protocolbuffers/protobuf/releases/download/v3.13.0/$PROTOC_ZIP \
    && unzip -o $PROTOC_ZIP -d /usr/local bin/protoc \
    && unzip -o $PROTOC_ZIP -d /usr/local 'include/*' \ 
    && rm -f $PROTOC_ZIP

RUN go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28
RUN go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2
ENV PATH=$PATH:/go/bin

# Copy program
WORKDIR /app
COPY . .

# Generate the proto files
WORKDIR /app
RUN mkdir proto_gen
RUN protoc \
    --go_out=/app/proto_gen \
    --go_opt=paths=source_relative \
    --go-grpc_out=/app/proto_gen \
    --go-grpc_opt=paths=source_relative \
    set_list.proto

# Run the integrate test
WORKDIR /app
CMD ["go", "test"]