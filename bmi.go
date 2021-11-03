package main

import (
	"github.com/suztopp/bmi/info"
)

func main() {

	info.PrintWelcome()

	weight, height := getUserMetrics()

	bmi := calculateBMI(weight, height)

	printBMI(bmi)

}

func calculateBMI(weight float64, height float64) float64 {
	return weight / (height * height)
}

// func main() {
// 	fmt.Println(info.MainTitle)
// 	fmt.Println(info.Separator)
// 	//output information

// 	fmt.Println(info.WeightPrompt)
// 	weightInput, _ := reader.ReadString('\n') //reads user input until certain character found
// 	//reads until the user hits enter key - ie \n for line space
// 	//underscore says we don't care what the second value is
// 	fmt.Println(info.HeightPrompt)
// 	heightInput, _ := reader.ReadString('\n')
// 	//prompt user for weight and height

// 	weightInput = strings.Replace(weightInput, "\n", "", -1)
// 	heightInput = strings.Replace(heightInput, "\n", "", -1)
// 	//this replaces a line break with nothing using the strings package
// 	//-1 means for all the line breaks you find
// 	weight, _ := strconv.ParseFloat(weightInput, 64)
// 	//second value is what type of float you want
// 	height, _ := strconv.ParseFloat(heightInput, 64)
// 	//save user input in variables

// 	bmi := weight / (height * height)
// 	//calculate the bmi (weight / (height * height))

// 	fmt.Printf("Your BMI is: %.2f", bmi)
// 	//output BMI

// }
