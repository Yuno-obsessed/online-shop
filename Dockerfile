# Initial stage: download modules
FROM golang:1.19 as modules

ADD ./go.mod ./go.sum /cmd/
RUN cd /cmd && go mod download


# Intermediate stage: Build the binary
FROM golang:1.19 as builder

COPY --from=modules /go/pkg /go/pkg

RUN mkdir -p /cmd
COPY . /cmd
WORKDIR /cmd

# Build the binary with go build
RUN GOOS=linux GOARCH=amd64 CGO_ENABLED=0 \
    go build -ldflags '-s -w -extldflags "-static"' \
    -o /bin/cmd ./cmd/online-shop/main.go

# Final stage: Run the binary
FROM alpine:latest as image

# and finally the binary

COPY --from=builder /bin/cmd .
EXPOSE 8080


ENTRYPOINT [ "./cmd" ]
