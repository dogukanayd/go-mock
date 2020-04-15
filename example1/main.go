package main

import (
	"fmt"
	"github.com/dogukanayd/go-mock/example1/greeter"
)

func main() {
	d := greeter.NewDB()

	g := greeter.NewGreeter(d, "en")
	fmt.Println(g.Greet()) // Message is: hello
	fmt.Println(g.GreetInDefaultMsg()) // Message is: default message

	g = greeter.NewGreeter(d, "es")
	fmt.Println(g.Greet()) // Message is: holla

	g = greeter.NewGreeter(d, "random")
	fmt.Println(g.Greet()) // Message is: bzzzz
}