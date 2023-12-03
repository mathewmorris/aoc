package main

import (
	"os"
	"testing"
)

func TestOnlyDigits(t *testing.T) {
  data, err := os.ReadFile("./only_digits.txt")

  if err != nil {
    t.Fatal(err)
  }

  guess := addDigitsOnly(data)
  answer := 142

  if guess != answer {
    t.Fatalf(`Your guess, %v, does not equal %v`, guess, answer)
  }
}

func TestTextToDigits(t *testing.T) {
  data, err := os.ReadFile("./text_to_digits.txt")

  if err != nil {
    t.Fatal(err)
  }

  guess := addTextToDigits(data)
  answer := 281

  if guess != answer {
    t.Fatalf(`Your guess, %v, does not equal %v`, guess, answer)
  }
}

