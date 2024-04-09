FROM golang:1.20-alpine

WORKDIR /go/src/app
COPY . .
RUN go build -o kiriban ./cmd

ENTRYPOINT ["/go/src/app/kiriban"]