# syntax=docker/dockerfile:1.2

# Obraz zawierający język golang
FROM golang:1.20.3-alpine3.16 AS BuildStage

WORKDIR /TimeShower

COPY . .

WORKDIR src/

# Pobranie zależności dla aplikacji golang
RUN go mod download

# Budowa aplikacji golang
RUN go build -o ../app

# Drugi etap budowy obrazu, gdzie jedynie jest kompiowana aplikacja z porzedniego etapu. Co sprawia, że sam obraz
# jest bardzo lekki ponieważ zawiera tylko skompliowaną apliakcję.
FROM alpine:latest

LABEL org.opencontainers.image.authors="Paweł Czerwieniec <pawel.czerwieniec@pollub.edu.pl>"

WORKDIR /

RUN apk add --no-cache tzdata

COPY --from=BuildStage /TimeShower/app /app

EXPOSE 80

ENTRYPOINT ["/app"]