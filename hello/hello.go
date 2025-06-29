package main

import (
    "fmt";
    "time"
)

func main() {
    name := "Nikita"
    currentTime := time.Now()

    fmt.Println("Hello, World!")
    fmt.Printf("Имя: %s\n", name)
    fmt.Printf("Текущая дата: %s\n", currentTime)

}
