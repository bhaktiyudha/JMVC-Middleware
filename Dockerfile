FROM golang:1.15-alpine as build-env

WORKDIR /app

RUN apk update && apk add --no-cache git
RUN apk add --no-cache tzdata
RUN apk add build-base
# All these steps will be cached
COPY go.mod .
COPY go.sum .

# Get dependancies - will also be cached if we won't change mod/sum
RUN go mod download
# COPY the source code as the last step
COPY . .
ARG CGO_ENABLED=0
RUN go build -v -o ./bin/JMVC-Middleware .

FROM alpine:3.12

RUN apk add --no-cache ca-certificates openssl
RUN apk add --no-cache tzdata

WORKDIR /app
COPY --from=build-env /app/bin/JMVC-Middleware ./

CMD ["./JMVC-Middleware"]