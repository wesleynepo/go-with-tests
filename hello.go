package main

import "fmt"

const spanish = "Spanish"
const englishHelloPrefix = "Hello, "
const spanishHelloPrefix = "Hola, "

func Hello(name string, languague string) string {
    if name == "" {
        name = "World"
    }

    if languague == spanish {
        return spanishHelloPrefix + name
    }

    return englishHelloPrefix + name
}

func main() {
    fmt.Println(Hello("world", ""))
}
