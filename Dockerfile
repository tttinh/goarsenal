FROM golang:1.16-alpine AS build
WORKDIR /go/src/github.com/tttinh/goarsenal
COPY . .

RUN go build -o server .

FROM alpine:3.14
EXPOSE 8080
COPY --from=build /go/src/github.com/tttinh/goarsenal/application.yml /application.yml
COPY --from=build /go/src/github.com/tttinh/goarsenal/server /server
CMD ["/server"]