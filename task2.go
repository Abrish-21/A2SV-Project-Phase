package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"
)

// workflow
// First step: accept string
// second step: clean string
// 3rd step: count frequency
// last step: print output



func CountWordFrequency() map[string] int{

	// accept strings with space 
    reader := bufio.NewReader(os.Stdin)
	fmt.Println(("ENter a string: "))
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)
	
	// in this step we clean the input 
	var b strings.Builder

	for _, val := range input {
		if unicode.IsDigit(val)  || unicode.IsLetter(val) {
			b.WriteRune(val)
		}

	}
	var CleanedInput string
	CleanedInput = b.String()
	// now change to lower case to make case insesitive 
	CleanedInput = strings.ToLower(CleanedInput)

	
 
	//create map and count frequency
	var frequency  map[string] int 
	frequency  = make(map[string] int )

	// loop through the input and count string  
	for _, val := range CleanedInput {
		ch:= string(val)
		_, exist:= frequency[ch]
		if exist { 
			frequency[ch] ++
		} else {
			frequency[ch] = 1
		}
	}
	
	return frequency 



}
func PalindormeCheck() bool {
	// accept strings with space 
    reader := bufio.NewReader(os.Stdin)
	fmt.Println(("ENter a string: "))
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)
	
	// in this step we clean the input 
	var b strings.Builder

	for _, val := range input {
		if unicode.IsDigit(val)  || unicode.IsLetter(val) {
			b.WriteRune(val)
		}

	}
	var CleanedInput string
	CleanedInput = b.String()
	// now change to lower case to make case insesitive 
	CleanedInput = strings.ToLower(CleanedInput)


	var first, last int  = 0, len(CleanedInput)-1

	for first < last {
	if CleanedInput[first] != CleanedInput[last] {
		return false
	} else {
		first++
		last--
	}

		
	}
	return true
}
