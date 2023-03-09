FROM golang:latest AS builder
WORKDIR /src
ENV CGO_ENABLED=0
COPY go.mod /src/
RUN go mod download
COPY . .
RUN go build -a -o flag4ctf -trimpath

FROM alpine:latest

RUN apk add --no-cache ca-certificates \
    && rm -rf /var/cache/*

RUN mkdir -p /app \
    && adduser -D flag4ctf \
    && chown -R flag4ctf:flag4ctf /app

USER flag4ctf
WORKDIR /app

COPY --from=builder /src/flag4ctf .

ENTRYPOINT [ "./flag4ctf" ]