package main

import (
    "fmt"
    "math"
)

func vars() {
    var age int = 20
    var name string = "Nikita"
    var isStudent bool = true

    height := 188.1

    fmt.Println("\n=== Задание 5 ===")
    fmt.Println("Имя:", name)
    fmt.Println("Возраст:", age)
    fmt.Println("Студент:", isStudent)
    fmt.Println("Рост (см):", height)
}

func consts() {
    const pi = math.Pi
    const e = math.E

    radius := 14.0
    area := pi * math.Pow(radius, 2)

    x := 3.0
    exp := math.Pow(e, x)

    fmt.Println("\n=== Задание 6 ===")
    fmt.Printf("π = %.4f\n", pi)
    fmt.Printf("e = %.4f\n", e)
    fmt.Printf("Площадь круга (r=8) = %.2f\n", area)
    fmt.Printf("e^%.1f = %.4f\n", x, exp)
}

func main() {
    vars()
    consts()
}
