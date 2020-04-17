FROM golang:1.13
WORKDIR /go/src/github.com/bkayranci/cron-entrypoint/

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o cron-entrypoint .

FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=0 /go/src/github.com/bkayranci/cron-entrypoint/cron-entrypoint .
CMD ["./cron-entrypoint", "-h"]
