FROM golang:1.18.3-alpine
RUN apk add --update --no-cache \
    vim \
    git \
    tar \
    postgresql-client

# migrate実行用にgolang-migrate install
RUN wget https://github.com/golang-migrate/migrate/releases/download/v4.14.1/migrate.linux-amd64.tar.gz && \
    tar -xvzf migrate.linux-amd64.tar.gz && \
    mv migrate.linux-amd64 bin/migrate && \
    rm migrate.linux-amd64.tar.gz

CMD ["go", "run", "/go/src/app/server.go"]
