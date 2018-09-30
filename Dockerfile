FROM golang:1.9.2
ENV SRC_DIR=/go/src/github.com/kevinbyrom/echo-server
ENV GOBIN=/go/bin
ENV CONN_HOST=0.0.0.0
ENV CONN_PORT=8088

WORKDIR $GOBIN

#Add the source code
ADD . $SRC_DIR

RUN cd /go/src/github.com/kevinbyrom;

RUN go install github.com/kevinbyrom/echo-server
ENTRYPOINT ["./echo-server"]