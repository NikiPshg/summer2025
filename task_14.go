package main

import (
    "fmt"
    "strings"
)

func main() {
    text := "hello world 1c 1c 1c more more"
    fmt.Println(text)

    wordFrequency := countWords(text)

    fmt.Println("Частота слов:")
    for word, count := range wordFrequency {
        fmt.Printf("%s: %d\n", word, count)
    }
}

func countWords(text string) map[string]int {
    words := strings.Fields(strings.ToLower(text))

    frequency := make(map[string]int)

    for _, word := range words {
        frequency[word]++
    }

    return frequency
}
