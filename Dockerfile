FROM golang:latest

ENV GOPROXY https://goproxy.cn,direct
WORKDIR $GOPATH/src/github.com/kasra73/fidibo-api
COPY . $GOPATH/src/github.com/kasra73/fidibo-api
RUN go build .

EXPOSE 8000
ENTRYPOINT ["./fidibo-api"]
