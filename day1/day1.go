package main


import (
    "os"
    "fmt"
    "bufio"
    "strconv"
)

func countIncreases(number, lastNumber *int) bool {
    if *number > *lastNumber {
        *lastNumber = *number
        return true
    }
    *lastNumber = *number
    return false
}

func main() {
    input, err := os.Open("input")
    if err != nil{
        fmt.Printf("error opening file: %v\n", err)
    }
    defer input.Close()

    scanner := bufio.NewScanner(input)
    var count int
    var lastNumber int
    for scanner.Scan() {
        number, _ := strconv.Atoi(scanner.Text())
        if countIncreases(&number, &lastNumber) { count++ }
    }
    fmt.Println(count)
}
