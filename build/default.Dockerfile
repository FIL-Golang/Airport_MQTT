FROM golang:1.21 AS builder

ARG MAIN_PATH

WORKDIR /go/src/app

COPY . .

RUN go mod download

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build \
    -ldflags="-w -s" -o /go/bin/app "$MAIN_PATH"

FROM scratch AS runner

COPY --from=builder /go/bin/app /

VOLUME ["/config.yaml"]

ENTRYPOINT ["/app"]
