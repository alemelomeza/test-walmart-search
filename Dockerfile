FROM golang:1.14.0-alpine3.11 as build
WORKDIR /app
ADD . /app
RUN go clean --modcache \
    && CGO_ENABLED=1 GOOS=linux go build -race -a -installsuffix cgo -o app .

FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /app
COPY --from=build /app .
CMD ["./app"]
EXPOSE 80
