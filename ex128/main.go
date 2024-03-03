package main

import "fmt"

func main (){

	type Grades struct{
		grade int
		course string
	}

	/*
	Exercise 1
	1. Create your own struct type called person that stores the following data: name, age and a list with favorite colors.
	2. Declare and initialize two values of type person, one called me and another called you.
	3. Print out the struct values.
	*/ 

	type person struct {
		name string
		age int
		colors []string
		grades Grades
	}

	me := person{name: "tjasa", age : 35, colors:[]string{"green", "blue"}, grades: Grades{ grade: 32, course: "math"} }
	you := person{name: "vitor", age : 38, colors:[]string{"green", "yellow"}, grades: Grades{ grade: 78, course: "languages"} }

	fmt.Printf("%+v\n", me)
	fmt.Printf("%+v\n", you)

	/*
	Exercise 2
	Consider the code from the previous exercise and:
1. Change the name or the struct value called me to "Andrei"
2. Take in a new variable the favorites colors of struct value called you. Print out the type and the value of the new variable.
3. Add a new favorite color to the second struct value called you.
4. Print out the struct values.
*/

me.name = "Andrei"
you.colors = []string{"white", "pink"}

fmt.Println(you.colors)

you.colors = append(you.colors, "purple")
fmt.Printf("%+v\n", you)

/*
Exercise 3
Consider the code from Exercise #1.
Iterate and print out the favorite colors of the struct value called me. 
*/

for index, color := range me.colors{
	fmt.Printf("%v color: %v\n", index, color)
} 

/*
Exercise 4
Change the code from Exercise #1 and:
1. Create a new struct type called grades with  2 fields: grade and course
2. Add another field of type grades to person struct type (embedded struct).
3. Change the initialization of the struct values called me and you to contain information about the grades.
4. Change the grade and the course of one struct value to "Golang" and 98.
5. Print out the struct values.
*/

me.grades.course = "Golang"
me.grades.grade = 98

fmt.Printf("%+v\n", me)


}