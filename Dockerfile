# syntax=docker/dockerfile:1


FROM golang:1.16-alpine AS build

WORKDIR /app

COPY go.mod ./
COPY go.sum ./

RUN go mod download

COPY /src/. ./

RUN go build -o /val-exporter

FROM alpine:latest

WORKDIR /

COPY --from=build /val-exporter /val-exporter

EXPOSE 26661

ENTRYPOINT [ "./val-exporter" ]