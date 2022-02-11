FROM golang:1.17-alpine
WORKDIR /go/src

COPY ./ ./
RUN apk add --no-cache gcc musl-dev
# TODO: [#46]