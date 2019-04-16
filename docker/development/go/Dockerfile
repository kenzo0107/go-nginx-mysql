FROM golang:1.11-alpine

RUN apk add --no-cache --update \
    bash build-base mysql-client mysql-dev tzdata git \
    && go get github.com/pilu/fresh \
    && go get github.com/rubenv/sql-migrate/...
RUN cp /usr/share/zoneinfo/Asia/Tokyo /etc/localtime

WORKDIR /go/src/app
COPY . .

CMD ["docker/development/go/entrypoint.sh"]

EXPOSE 8080
