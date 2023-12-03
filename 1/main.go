package main

import (
  "fmt"
  "log"
  "os"
  "strings"
  "unicode"
)

func addDigitsOnly(data []byte) int {
  sum := 0
  var numbers []int

  for _, byte := range data {
    str := string(byte)

    if byte == '\n' {
      if len(numbers) >= 1 {
        combinedNumber := numbers[0]*10 + numbers[len(numbers)-1]
        sum = sum + combinedNumber
        numbers = numbers[:0]
      }
    }

    if unicode.IsDigit(rune(str[0])) {
      intValue := int(byte - '0')
      numbers = append(numbers, intValue) 
    }
  }

  return sum;
}


func addTextToDigits(data []byte) int {
  wordToDigit := map[string]string{
    "zero": "0",
    "one": "1",
    "two": "2",
    "three": "3",
    "four": "4",
    "five": "5",
    "six": "6",
    "seven": "7",
    "eight": "8",
    "nine": "9",
  }

  sum := 0
  slices := strings.Split(string(data), "\n")

  for _, slice := range slices {
    str := slice
    fmt.Println("before: ", str)
    // while this kind of works, it's not the best way to do it
    // 'eightwoone4' would be replaced to 'eigh214' instead of '8wo14'
    // need to "read" from left to right
    for word, digit := range wordToDigit {
      str = strings.ReplaceAll(str, word, digit)
    }
      fmt.Println("after: ", str)
  }

  return sum;
}

func main() {
  data, err := os.ReadFile("./input.txt")

  if err != nil {
    log.Fatal(err)
  }

  digitsOnly := addDigitsOnly(data)

  fmt.Println(digitsOnly)

  textToDigits := addTextToDigits(data)

  fmt.Println(textToDigits)
}

