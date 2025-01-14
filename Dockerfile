FROM golang:1.20 as build

ARG VERSION

ADD . /go/src/pluto

WORKDIR /go/src/pluto

RUN  export GO111MODULE=on GOPROXY=https://proxy.golang.org && \ 
  go build -ldflags="-X 'main.VERSION=${VERSION}'" -o pluto-server cmd/pluto-server/main.go && \
  go build -o pluto-migrate cmd/pluto-migrate/main.go

FROM ubuntu:22.04

RUN apt-get update && apt-get install -y wget

ENV DOCKERIZE_VERSION v0.6.1
RUN wget https://github.com/jwilder/dockerize/releases/download/$DOCKERIZE_VERSION/dockerize-linux-amd64-$DOCKERIZE_VERSION.tar.gz \
    && tar -C /usr/local/bin -xzvf dockerize-linux-amd64-$DOCKERIZE_VERSION.tar.gz \
    && rm dockerize-linux-amd64-$DOCKERIZE_VERSION.tar.gz

COPY --from=build /go/src/pluto/pluto-server /usr/bin/

COPY --from=build /go/src/pluto/pluto-migrate /usr/bin/

COPY --from=build /go/src/pluto/views views/

COPY --from=build /go/src/pluto/localization/*.toml localization/

ENTRYPOINT ["/usr/bin/pluto-server"]
