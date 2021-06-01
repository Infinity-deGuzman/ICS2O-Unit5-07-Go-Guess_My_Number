// Created by: Infinity de Guzman
// Created on: May 2021
//
// This program makes you guess a number

package main

import (
	
	"fmt"

	"math/rand"

	"time"
)
	
func main() {
  var userGuess int
  var answer int
  var counter = 0

	// process
	rand.Seed(time.Now().UnixNano())
	min := 0
	max := 9

	answer = rand.Intn(max-min) + min

	

  for userGuess != answer && counter < 10 {
		fmt.Print(" Enter a number from 0 to 9 (Integers only): ")
		fmt.Scanln(&userGuess)
		if userGuess < answer {
			fmt.Print(userGuess, " is too low")
			counter++
		} else if userGuess > answer {
			fmt.Print(userGuess, " is too high")
			counter++
		} else {
			fmt.Print("Congrats! You got it right! The hidden number is ", answer)
			counter++
			break
		}
	}
}
