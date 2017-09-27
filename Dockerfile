
# Build stage
FROM golang:1.8 as builder

COPY . /go/src/gobls3

WORKDIR /go/src/gobls3
RUN go get && CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o /go/bin/app


# Runtime stage
FROM golang:1.8-alpine

WORKDIR /root/
COPY --from=builder /go/bin/app .
COPY --from=builder /go/src/gobls3/config.json .

CMD ["./app"]

