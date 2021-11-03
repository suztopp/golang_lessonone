package info

import "fmt"

const mainTitle = "BMI Calculator"
const separator = "---------------"
const WeightPrompt = "Please enter your weight in kgs: "
const HeightPrompt = "Please now enter your height in metres: "

//these needed to change to capital first letters

func PrintWelcome() {
	fmt.Println(mainTitle)
	fmt.Println(separator)
}
