package main

import "fmt"

func FindLongestString(strings ...string) string {
    if len(strings) == 0 {
        return ""
    }

    longest := strings[0]
    for _, s := range strings[1:] {
        if len(s) > len(longest) {
            longest = s
        }
    }
    return longest
}

func main() {
    fmt.Println("Самая длинная строка:", FindLongestString("one", "two", "three", "four"))
    fmt.Println("Самая длинная строка:", FindLongestString("apple", "banana", "orange"))
    fmt.Println("Самая длинная строка:", FindLongestString("short", "very long str object", "average"))

    fmt.Println("Самая длинная строка:", FindLongestString())

    words := []string{"Go", "Python", "JavaScript", "Java"}
    fmt.Println("Самая длинная строка:", FindLongestString(words...))
}
