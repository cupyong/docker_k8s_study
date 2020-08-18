FROM golang:1.14.7
WORKDIR /hellodocker
ADD . /hellodocker
RUN go build .
EXPOSE 6001
ENTRYPOINT ["./hellodocker"]