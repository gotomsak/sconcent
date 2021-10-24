FROM golang:1.17.2-alpine3.14

WORKDIR /go/src/app
COPY . .

RUN go get -d -v ./...
RUN go install -v ./...
RUN go build -v -o sconcent

CMD ["sconcent"]