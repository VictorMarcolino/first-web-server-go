FROM golang:1.17-alpine as builder

MAINTAINER Victor Marcolino<marcolino.victor@gmail.com>
EXPOSE 8080
WORKDIR /src
COPY src /src
RUN go build -o main -v ./cmd/web/*.go

FROM alpine:latest

COPY --from=builder /src/main /src/main
COPY --from=builder /src/templates /src/templates
WORKDIR /src
CMD ./main