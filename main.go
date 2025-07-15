package main

import "fmt"


func main() {
	// task1:= calculateGPA()
	
	// fmt.Printf("GPA: %.2f ", task1)
	task3 := CountWordFrequency()
	fmt.Println(task3)
	if PalindormeCheck() {
		fmt.Println(("It is palindrome"))
	} else {
		fmt.Println("It is not palindorme")
	}
	
}
