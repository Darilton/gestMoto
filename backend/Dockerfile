FROM golang:1.25

WORKDIR /usr/src/app

COPY . .

RUN go build -v -o /usr/local/bin/app ./...

CMD ["app"]