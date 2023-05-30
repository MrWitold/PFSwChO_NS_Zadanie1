# syntax=docker/dockerfile:1

FROM golang:1.20.3-alpine3.16 AS buildstage

WORKDIR /TimeShower

COPY . .

RUN go mod download

EXPOSE 80

RUN go build -o /TimeShower main.go

# Deploy Stage
FROM alpine:latest

WORKDIR /

COPY --from-buildstage /TimeShower /TimeShower

EXPOSE 80

USER nonroot:nonroot

ENTRYPOINT ["/TimeShower"]