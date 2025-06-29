package main

import "fmt"

func main() {
    grades := make(map[string]int)

    grades["Фирдафси"] = 4
    grades["Андрей"] = 5
    grades["Матвей"] = 2
    grades["Максим"] = 5
    grades["Никиата"] = 3

    fmt.Println("Оценка Фирдафси:", findGrade(grades, "Фирдафси"))
    fmt.Println("Оценка Петра:", findGrade(grades, "Петр")) //такого нет

    removeGrade(grades, "Никиата")
    fmt.Println("Оценка Никиаты после удаления:", findGrade(grades, "Никиата"))

    fmt.Println("Все оценки:", grades)
}

func addGrade(grades map[string]int, name string, grade int) {
    grades[name] = grade
}

func findGrade(grades map[string]int, name string) int {
    grade, exists := grades[name]
    if !exists {
        return -1
    }
    return grade
}

func removeGrade(grades map[string]int, name string) {
    delete(grades, name)
}
