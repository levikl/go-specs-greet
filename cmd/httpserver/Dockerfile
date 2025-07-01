# Make sure to specify the same Go version as the one in the go.mod file
FROM golang:1.24.4-alpine as builder

WORKDIR /app

COPY go.mod ./

RUN go mod download

COPY . .

RUN go build -o svr cmd/httpserver/*.go

FROM alpine:3.16.2
COPY --from=builder /app/svr .
# EXPOSE 8080
CMD [ "./svr" ]
