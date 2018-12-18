FROM golang:alpine AS build-env

RUN apk update && apk add --no-cache git
RUN apk --no-cache add tzdata

ENV TZ=Asia/Bangkok
ENV GO111MODULE=on
ENV CGO_ENABLED=0
ADD . /src

RUN cd /src && go test ./...
RUN cd /src && go build -o goapp main.go

FROM alpine:latest

WORKDIR /app
COPY --from=build-env /src/goapp /app/

EXPOSE 1323
ENTRYPOINT [ "/app/goapp" ]
