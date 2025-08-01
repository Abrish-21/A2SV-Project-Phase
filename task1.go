package main

import "fmt"

type SubjectGradePair struct {
	Subject string
	Grade float32

}
func calculateGPA() float32 {
	var average float32
	var name string
	var totalGrades int
	var grade float32
	var subject string
	var grades []SubjectGradePair  
	

	fmt.Println("Hello There Please Enter Your Name: ")
	fmt.Scanln(&name)
	fmt.Printf("Welcome %s, Let's Calculate Your Grade\n", name)

	// enter total subjects 
	fmt.Println("Enter the number of subjects: ")
	fmt.Scanln(&totalGrades)
	
	// Loop over to accept the Grade, Subject Pair 
	
	for i:=1; i<= totalGrades;i++ {
		fmt.Println("Enter your subject grade pairs")
		
		// accept subject name 
		fmt.Println("Enter subject ", i)
		fmt.Scanln(&subject)

		// accept subject grade 
		fmt.Printf("Enter your %v grade:", subject)
		fmt.Scanln(&grade)

        // add subject grade as a pair 
		sg:= SubjectGradePair {
			Subject: subject,
			Grade: grade,
		}
	grades = append(grades, sg)
		
	}
	// print the separate grades 
	println("Grade Calculator:")
	println("Name: ", name)
	println("Here is the course break down:")
	for _,sg := range grades {
		fmt.Printf("%s: %.2f\n", sg.Subject, sg.Grade)
	}

	for _, sg := range grades {
		average += sg.Grade

	}
	 return average / float32(len(grades))
}