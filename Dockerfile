FROM golang:1.22
WORKDIR /mnt/app
COPY ../.. .
RUN go build

FROM docker
COPY --from=0 /mnt/app/mailer-micro-service /usr/local/bin/mailer-micro-service
RUN apk add bash curl libc6-compat
