FROM golang:1.11

WORKDIR /go/src/app
COPY . .

RUN go get -d -v ./...

CMD ["app"]
