FROM golang:1.8
WORKDIR /go/src/apigateway
COPY apigateway .
RUN go-wrapper download
RUN go-wrapper install
CMD ["go-wrapper", "run"]
