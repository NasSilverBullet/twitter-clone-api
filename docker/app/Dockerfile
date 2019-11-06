FROM golang:1.13.4-alpine3.10

RUN apk add --no-cache \
    git \
    binutils-gold \
    curl \
    g++ \
    gcc \
    gnupg \
    libgcc \
    linux-headers \
    make

RUN go get github.com/oxequa/realize
