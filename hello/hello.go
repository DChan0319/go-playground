package main

import "fmt"

const helloString = "Hello "

// Hello func returns: string
func Hello(name string) string {
	if name == "" {
		return "Hello, World"
	}
	return helloString + name + ", World"
}

func main() {
	fmt.Println(Hello("Darren"))
}
