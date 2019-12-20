FROM alpine

RUN go build trans_server.go

WORKDIR /app

EXPOSE 10011

ADD trans_server /app/

CMD ["/app/trans_server"]
