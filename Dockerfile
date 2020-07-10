FROM golang:1.13
MAINTAINER David Greco <david.greco@one.verizon.com>
COPY . /go/src
WORKDIR /go/src/
RUN go build -o PhoneBookServer
# Run Application
EXPOSE  8080
ENTRYPOINT ["./PhoneBookServer"]
