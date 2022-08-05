package main

import "fmt"

const spanish = "Spanish"
const french = "French"
const englishHelloPrefix = "Hello, "
const spanishHelloPrefix = "Hola, "
const frenchHelloPrefix = "Bonjour, "

func Hello(name string, languague string) string {
    if name == "" {
        name = "World"
    }

    prefix := englishHelloPrefix

    switch languague {
    case french:
        prefix = frenchHelloPrefix
    case spanish:
        prefix = spanishHelloPrefix
    }

    return prefix + name
}

func main() {
    fmt.Println(Hello("world", ""))
}
