FROM golang:1.21
WORKDIR /mnt/app
COPY ../.. .
RUN go build

COPY mailer-micro-service /usr/local/bin/mailer-micro-service
