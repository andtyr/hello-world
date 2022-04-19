package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	//output info
	fmt.Println("BMI Calculator")
	fmt.Println("-----------")

	//prompt for user input (w + h)
	fmt.Print("Please enter your weight (kg): ")
	weightInput, _ := reader.ReadString('\n')
	fmt.Print("Please enter your height (m): ")
	heightInput, _ := reader.ReadString('\n')

	//save user input as var
	weightInput = strings.Replace(weightInput, "\n", "", -1)
	heightInput = strings.Replace(heightInput, "\n", "", -1)

	weight, _ := strconv.ParseFloat(weightInput, 64)
	height, _ := strconv.ParseFloat(heightInput, 64)

	//calc BMI ( w / (h * h))
	bmi := weight / (height * height)

	//output calc BMI
	fmt.Printf("Your BMI: %.2f", bmi)

}
