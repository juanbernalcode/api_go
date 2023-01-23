package server

import "net/http"

type Country struct {
	Name     string
	Languaje string
}

var countries []*Country = []*Country{}

func New(addr string) *http.Server {

	initRoutes()

	return &http.Server{
		Addr: addr,
	}
}
