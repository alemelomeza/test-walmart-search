FROM golang:1.14.0-alpine3.11
RUN mkdir /app
ADD . /app
WORKDIR /app
RUN go clean --modcache
RUN go build -o main .
CMD ["/app/main"]
EXPOSE 80