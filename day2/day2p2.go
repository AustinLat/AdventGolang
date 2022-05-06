package main

import (
    "bufio"
    "os"
    "fmt"
    "strings"
    "strconv"
)

func main() {
    input, err := os.Open("input")
    if err != nil {
        fmt.Printf("Error opening file: %v\n", err)
    }
    defer input.Close()
    scanner := bufio.NewScanner(input)

    var distance int
    var depth int
    var aim int
    for scanner.Scan() {
        command := scanner.Text()
        commandArray := strings.Fields(command)
        value, _ := strconv.Atoi(commandArray[1])
        switch commandArray[0] {
        case "forward":
            distance = distance + value
            depth = depth + aim * value
        case "up":
            aim = aim - value
        case "down":
            aim = aim + value
        }
    }
    fmt.Println(distance*depth)
}
