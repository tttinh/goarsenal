# A REST Service Using Gin

Another example of building REST service in Go.

## Installation

```
$ go get github.com/tttinh/goarsenal
```

## Run the application using docker-compose

```
$ cd $GOPATH/src/goarsenal
$ docker-compose up -d
```

To run and rebuild the service image, please run:

```
$ docker-compose up -d --build
```

## Run the application from source

- Start your MySQL server.
- Modify `application.yml` with your database credentials.
- Run the application using the following commands:

```
$ cd $GOPATH/src/goarsenal

$ go run main.go
```

## Features

- Using GORM for database operation.
- Using Zap for logging.
- Using Gin for handling HTTP requests.
- Using Viper for application configuration.

## Notes

- The API to get a list of wagers will return an empty array `[]` if no data found. And the default `page`, `limit` are 0 and 10 if no query parameters provided.
