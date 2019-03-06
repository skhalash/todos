FROM golang:latest as builder
ADD . /go/
WORKDIR /go/src/services/todos
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o app .

FROM alpine:latest
WORKDIR /root/
COPY --from=builder /go/src/services/todos/app .
CMD ["./app"]
