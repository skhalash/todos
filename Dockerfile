FROM golang:latest
ADD . /go/
WORKDIR /go/src/github.com/skhalash/todos
RUN go build -o main .
CMD ["/go/main"]
