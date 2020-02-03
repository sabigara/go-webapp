FROM golang:1.13-alpine

RUN ls
RUN apk add make
COPY entrypoint.sh /entrypoint.sh

ENTRYPOINT ["/entrypoint.sh"]
