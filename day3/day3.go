package main

import (
    "os"
    "bufio"
//    "strconv"
    "fmt"
)

func main() {
    input, err :=  os.Open("input")
    if err != nil {
        fmt.Printf("Error opening file: %v\n", err)
    }
    defer input.Close()
    scanner := bufio.NewScanner(input)

    var gamma string
    var epsilon string
    for i := range [12]int{} {
        one := 0
        zero := 0
        for scanner.Scan() {
            line := scanner.Text()
            if string(line[i]) == "0" {
                zero++
            } else {
                one ++
            }
        }
        if one > zero {
            gamma += "1"
            epsilon += "0"
        } else {
            gamma += "0"
            epsilon += "1"
        }
    }
    fmt.Println(gamma)
}
