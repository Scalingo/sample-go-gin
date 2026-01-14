FROM golang:1.25.5

RUN go install github.com/cespare/reflex@latest
ADD . /go/src/github.com/Scalingo/sample-go-gin
WORKDIR /go/src/github.com/Scalingo/sample-go-gin
EXPOSE 3000
RUN go install -buildvcs=false
CMD /go/bin/sample-go-gin
