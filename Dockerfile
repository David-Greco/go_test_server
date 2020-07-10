FROM golang:1.13
MAINTAINER David Greco <david.greco@one.verizon.com>
RUN go get -u https://github.com/David-Greco/go_test_server
RUN go build -o webserver
# Run Application
EXPOSE  8080
ENTRYPOINT ["./webserver"]
RUN go run ./PhoneBookServer.go
