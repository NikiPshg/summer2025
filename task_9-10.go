package main

import (
    "fmt"
    "strings"
)

func task1() {
    fmt.Println("\n=== Задание 9 ===")

    str := "GoLang — интересный язык программирования !! Go !"

    fmt.Println("Верхний регистр:", strings.ToUpper(str))
    fmt.Println("Нижний регистр:", strings.ToLower(str))

    fmt.Println("Начинается с 'Go'?", strings.HasPrefix(str, "Go"))
    fmt.Println("Заканчивается на '!'?", strings.HasSuffix(str, "!"))

    substr := "Go"
    fmt.Println("Содержит 'язык'?", strings.Contains(str, "язык"))
    fmt.Printf("'%s' встречается %d раз(а)\n", substr, strings.Count(str, substr))

    newStr := strings.ReplaceAll(str, "Go", "GOLANG")
    fmt.Println("После замены:", newStr)
}

func task2() {
    fmt.Println("\n=== Задание 10 ===")

    sentence := "Ульяновск,Чебоксары,Москва,Севастополь,Казань"
    fmt.Println(sentence)

    words := strings.Split(sentence, ",")
    for i, word := range words {
        fmt.Printf("Слово %d: %s\n", i+1, word)
    }

    joined := strings.Join(words, " | ")
    fmt.Println("\nСоединённая строка:", joined)
}

func main() {
    task1()
    task2()
}
