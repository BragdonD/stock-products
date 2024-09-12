package stockproducts

import (
	"fmt"
	"net/http"

	simplerest "github.com/bragdond/simple-rest"
)

// router is the structure that holds all the datas relative to
// HTTP requests.
type router struct {
	server *simplerest.Server
}

// NewRouter initializes a new router for the given server and registers
// the necessary routes for handling product-related HTTP requests.
func NewRouter(server *simplerest.Server) (*router, error) {
	router := &router{
		server: server,
	}

	if err := router.server.HandleFunc("/products",
		nil,
		handleProducts, http.MethodGet, http.MethodPost); err != nil {
		return nil, fmt.Errorf("unable to register route %s on the server: %v", "/products", err)
	}

	if err := router.server.HandleFunc("/products/{uuid}",
		nil,
		handleProduct, http.MethodGet, http.MethodPut, http.MethodPatch, http.MethodDelete); err != nil {
		return nil, fmt.Errorf("unable to register route %s on the server: %v", "/products/{uuid}", err)
	}

	if err := router.server.HandleFunc("/products/search",
		nil,
		handleSearch, http.MethodGet); err != nil {
		return nil, fmt.Errorf("unable to register route %s on the server: %v", "/products/search", err)
	}

	return router, nil
}

// Serve enables the http server
func (r *router) Serve() error {
	return r.server.Serve()
}

// Close stops the http server
func (r *router) Close() error {
	return r.server.Close()
}

func handleProducts(w http.ResponseWriter, req *http.Request, params simplerest.Parameters) error {
	return nil
}

func handleProduct(w http.ResponseWriter, req *http.Request, params simplerest.Parameters) error {
	return nil
}

func handleSearch(w http.ResponseWriter, req *http.Request, params simplerest.Parameters) error {
	return nil
}
