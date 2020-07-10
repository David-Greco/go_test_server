FROM golang:1.13
MAINTAINER David Greco <david.greco@one.verizon.com>
COPY . /go/src
WORKDIR /go/src/
RUN go run .\PhoneBookServer.go
# Run Application
EXPOSE  8080
ENTRYPOINT ["./phonebook"]
