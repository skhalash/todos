FROM golang:latest as builder
ADD . /go/
WORKDIR /go/src/github.com/skhalash/todos
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o app .

FROM alpine:latest
WORKDIR /root/
COPY --from=builder /go/src/github.com/skhalash/todos/app .
CMD ["./app"]
