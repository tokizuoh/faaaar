FROM golang:1.17-alpine
WORKDIR /go/src

COPY ./ ./
RUN apk add --no-cache gcc musl-dev
RUN go mod download
CMD go run main.go