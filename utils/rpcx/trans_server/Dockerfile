FROM scratch

RUN go build trans_server.go

WORKDIR /app

EXPOSE 10011
#所有的文件复制均使用COPY,仅在需要自动解压缩时使用ADD
COPY trans_server /app/

CMD ["/app/trans_server"]
