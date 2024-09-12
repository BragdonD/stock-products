package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	simplerest "github.com/bragdond/simple-rest"
	stockproducts "github.com/bragdond/stock-products"
)

func main() {
	server := simplerest.NewServer("localhost", 8080)
	router, err := stockproducts.NewRouter(server)
	if err != nil {
		panic(err)
	}
	go router.Serve()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	<-c
	router.Close()
	fmt.Println("Server has been closed...")
}
