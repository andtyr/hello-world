package main

import (
	"fmt"
)

func main() {
	fmt.Println("hello-world")

	greetinText := "Hello World!!!!!"
	luckyNumber := 69
	evenMoreLuckyNumber := luckyNumber - 10

	fmt.Println(greetinText)
	fmt.Println(luckyNumber)
	fmt.Println(evenMoreLuckyNumber)

	evenMoreLuckyNumber = luckyNumber * 3

	fmt.Println(evenMoreLuckyNumber)

	firstName := "Nunya"
	lastName := "Bidniss"
	var currentYear int = 2022
	currentYear = currentYear + 1
	birthYear := 1922
	nextYear := currentYear + 1
	age := currentYear - birthYear

	fmt.Println(firstName)
	fmt.Println(lastName)
	fmt.Println(currentYear)
	fmt.Println(birthYear)
	fmt.Println(nextYear)
	fmt.Println(age)

	var newNumber float64 = float64(luckyNumber) / 4

	fmt.Println(newNumber)

	var defaultFloat float64 = 1.23456789123456789
	var smallFloat float32 = 1.23456789123456789

	fmt.Println(defaultFloat)
	fmt.Println(smallFloat)

	var firstRune rune = '€'
	fmt.Println(firstRune)
	fmt.Println(string(firstRune))

	//Output as numeral '97', which corresponds to char value.
	var firstByte byte = 'a'
	fmt.Println(firstByte)

	//Example of combining strings. Whitespace must be specified with either '.' or [string] + " " + [string].
	fmt.Println(firstName, lastName)

	fullName := (firstName + " " + lastName)

	fmt.Printf("Hi, I am %v and I am %v years old\n", fullName, age)

	var pi float32 = 3.14
	fmt.Println(pi)

	var circleRadius int = 5
	fmt.Println(circleRadius)

	newFormula := (2 * pi * float32(circleRadius))
	fmt.Println(newFormula)

	fmt.Printf("For a radius of %v, the circle Circumference is %.2f", circleRadius, newFormula)

	//(Type %T) to come
}
