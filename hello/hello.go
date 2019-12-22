package main

import "fmt"

// Hello func returns: string
func Hello(name string) string {
	return "Hello " + name + ", World"
}

func main() {
	fmt.Println(Hello("Darren"))
}
