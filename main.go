package main

import (
	"log"
	"os"

	"github.com/armon/go-socks5"
)

var (
	port = "1080"
)

func main() {
	conf := &socks5.Config{
		Logger: log.New(os.Stdout, "", log.LstdFlags),
	}
	if os.Getenv("PROXY_USER") != "" {
		conf.AuthMethods = []socks5.Authenticator{
			socks5.UserPassAuthenticator{
				Credentials: socks5.StaticCredentials{
					os.Getenv("PROXY_USER"): os.Getenv("PROXY_PASSWORD"),
				},
			},
		}
	}

	if os.Getenv("PROXY_PORT") != "" {
		port = os.Getenv("PROXY_PORT")
	}

	srv, err := socks5.New(conf)
	if err != nil {
		panic(err)
	}

	if err := srv.ListenAndServe("tcp", "0.0.0.0:"+port); err != nil {
		panic(err)
	}
}
