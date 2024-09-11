package stockproducts

import simplerest "github.com/bragdond/simple-rest"

type Router struct {
	server *simplerest.Server
}

func NewRouter() (Router, error) {

	return Router{}, nil
}
