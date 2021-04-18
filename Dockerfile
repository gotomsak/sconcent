FROM golang:1.15-alpine3.12

WORKDIR /go/src/app
COPY . .

RUN go get -d -v ./...
RUN go install -v ./...
RUN go build -v -o save-concentration

CMD ["save-concentration"]