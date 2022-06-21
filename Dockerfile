FROM golang:1.18.3-alpine
RUN apk add --update --no-cache \
    vim \
    git \
    tar \
    postgresql-client
