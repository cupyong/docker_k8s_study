FROM golang:1.14.7
WORKDIR /compose
ADD . /compose
RUN go build .
EXPOSE 6001
ENTRYPOINT ["./compose"]