FROM golang:1.16 as builder
RUN go get -u github.com/pressly/goose/cmd/goose

FROM debian
WORKDIR /app
COPY --from=builder /go/bin/goose .
COPY ./migrations .
COPY ./scripts/migrate.sh .
RUN chmod 744 migrate.sh
CMD ["/app/migrate.sh"]