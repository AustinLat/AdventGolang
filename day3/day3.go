package main

import (
  "os"
  "bufio"
  "strconv"
  "fmt"
)

func main() {
  input, err :=  os.Open("input")
  if err != nil {
    fmt.Printf("Error opening file: %v\n", err)
  }
  defer input.Close()
  scanner := bufio.NewScanner(input)
  list := []string{}
  for scanner.Scan() {
    line := scanner.Text()
    list = append(list, line)
  }
  fmt.Println(powerConsumption(list))
}


func lifeSupport(data []string) int64 {
  one := 0
  zero := 0
  for i, line := range data {
    if string(line[i]) == "0" {
      zero++
    } else {
      one++
    }
  }
  for i, line := range data {
    if one > zero {
      if string(line[i]) == "0" {
        zero++
      } else {
        one++
      }
    }
  }

func powerConsumption(data []string) int64 {
  var gamma string
  var epsilon string
  for i, _ := range [12]int{} {
    one := 0
    zero := 0
    for _, line := range data {
      if string(line[i]) == "0" {
        zero++
      } else {
        one++
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

  gammaInt, err := strconv.ParseInt(gamma, 2, 32)
    if err != nil {
      fmt.Println(err)
    }
  epsilonInt, err := strconv.ParseInt(epsilon, 2, 32)
  if err != nil {
    fmt.Println(err)
  }
  return gammaInt * epsilonInt
}
