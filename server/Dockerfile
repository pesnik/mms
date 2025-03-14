FROM golang:1.23 AS builder

WORKDIR /go/src/github.com/pesnik/mms

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o mms-server cmd/api/main.go
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o mms-migrate cmd/migrate/main.go

FROM alpine:3.18
WORKDIR /app

RUN apk --no-cache add ca-certificates

COPY --from=builder /go/src/github.com/pesnik/mms/mms-server .
COPY --from=builder /go/src/github.com/pesnik/mms/mms-migrate .
COPY migrations/ /app/migrations/

RUN chmod +x /app/mms-server /app/mms-migrate
EXPOSE ${SERVER_PORT}
CMD ["./mms-server"]
