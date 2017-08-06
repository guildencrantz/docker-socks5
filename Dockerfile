FROM golang:1.8 as build
WORKDIR /go/src/github.com/guildencrantz/docker-socks5
#RUN go get -d -v golang.org/x/net/html
COPY . .
RUN go get && CGO_ENABLED=0 GOOS=linux go build -o socks5

FROM scratch
COPY --from=build /go/src/github.com/guildencrantz/docker-socks5/socks5 /socks5
ENTRYPOINT ["/socks5"]
