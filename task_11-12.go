package main

import "fmt"

func task1() {
    fmt.Println("\n=== Задание 11 ===")

    var slice []string

    slice = append(slice, "python")
    slice = append(slice, "rust")
    slice = append(slice, "go")

    fmt.Println("Срез после добавления:")
    for i, fruit := range slice {
        fmt.Printf("Индекс %d: %s\n", i, fruit)
    }
}

func task2() {
    fmt.Println("\n=== Задание 12 ===")

    slice := []string{"python", "rust", "go", "1c"}
    fmt.Println("Исходный срез:", slice)

    indexToRemove := 1

    slice = removeElement(slice, indexToRemove)

    fmt.Println("Срез после удаления:", slice)
}

func removeElement(slice []string, index int) []string {
    return append(slice[:index], slice[index+1:]...)
}

func main() {
    task1()
    task2()
}
