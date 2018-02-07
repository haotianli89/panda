FROM golang:latest
COPY . /go/src/github.com/haotianli89/panda
WORKDIR /go/src/github.com/haotianli89/panda

RUN go get github.com/golang/protobuf/proto
RUN go get github.com/micro/go-micro
RUN go get golang.org/x/net/context
RUN go get github.com/haotianli89/driversvc/pb

RUN go build -o main main.go

EXPOSE 8080

CMD ["./main", "--registry=mdns"]
