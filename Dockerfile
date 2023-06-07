FROM golang:latest
RUN mkdir /weavestore
ADD go.mod /weavestore/
ADD go.sum /weavestore/
ADD . /weavestore
WORKDIR /weavestore
RUN go build -o main .
CMD ["/weavestore/main"]