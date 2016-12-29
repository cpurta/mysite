FROM golang:latest

MAINTAINER Chris Purta cpurta@gmail.com

RUN apt-get update && \
    apt-get install -y ca-certificates && \
    go get github.com/revel/revel && \
    go get github.com/revel/cmd/revel && \
    mkdir -p /app

ADD ./bin/mysite /app

RUN chmod +x /app/mysite

EXPOSE 9000

ENTRYPOINT ["/app/mysite"]
