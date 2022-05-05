package main


import (
    "os"
   "fmt"
    "bufio"
    "strconv"
)

func shiftingDepths(depths []int) int {
    count := 0
    lastNumber := depths[0]+depths[1]+depths[2]
    for i:=0; i<=len(depths)-3; i++ {
        total := depths[i]+depths[i+1]+depths[i+2]
        if total > lastNumber {
            count++
        }
        lastNumber = total
    }
    return count
}

func main() {
    input, err := os.Open("input")
    if err != nil{
        fmt.Printf("error opening file: %v\n", err)
    }
    defer input.Close()

    scanner := bufio.NewScanner(input)
    depths := []int{}
    for scanner.Scan() {
        number, _ := strconv.Atoi(scanner.Text())
        depths = append(depths, number)
    }
    fmt.Println(shiftingDepths(depths))
}
