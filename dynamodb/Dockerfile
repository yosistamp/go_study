FROM golang:1.15.2 as builder

WORKDIR /go/pkg

COPY ./pkg/  ./

ARG CGO_ENABLED=0
ARG GOOS=linux
ARG GOARCH=amd64
RUN go mod tidy && \
    go build \
    -o /go/bin/main \
    -ldflags '-s -w'

FROM scratch as runner

COPY --from=builder /go/bin/main /app/main

ENTRYPOINT ["/app/main"]
