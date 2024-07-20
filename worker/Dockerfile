FROM golang:latest AS builder
WORKDIR /app
COPY ./worker/go.mod ./worker/go.sum ./
COPY ./worker/*.go .
RUN go mod download \
    && GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o binary main.go

FROM alpine:3.11
ENV TZ=Asia/Jakarta

RUN apk add --no-cache --upgrade \
    bash \
    tzdata\
    && cp /usr/share/zoneinfo/$TZ /etc/localtime && echo $TZ > /etc/timezone
RUN mkdir /app
COPY --from=builder /app/binary /app
WORKDIR /app
RUN chmod +x binary
CMD ["./binary"]