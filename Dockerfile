FROM golang:alpine AS build-env

RUN apk update && apk add --no-cache git
ENV GO111MODULE=on
ENV CGO_ENABLED=0
ADD . /src

RUN cd /src && go build -o goapp main.go

FROM scratch

WORKDIR /app
COPY --from=build-env /src/goapp /app/

EXPOSE 1323
ENTRYPOINT [ "/app/goapp" ]
