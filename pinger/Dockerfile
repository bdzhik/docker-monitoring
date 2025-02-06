FROM golang:1.23-alpine

WORKDIR /app
COPY . .

RUN go build -o pinger .

RUN apk add --no-cache docker-cli

CMD ["/app/pinger"]
