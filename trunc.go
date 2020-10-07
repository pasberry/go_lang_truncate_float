package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	fmt.Println("!!!!!!!! Welcome to trunc.go !!!!!")
	fmt.Println("Please enter a floating point number. Type exit to leave the program.")

	userResponse := ""

	for {

		fmt.Scan(&userResponse)

		if userResponse == "exit" {
			os.Exit(0)
		} else {

			userResponseAsFloat, err := strconv.ParseFloat(checkForPresenceOfDecimalPoint(userResponse), 64)

			if err != nil {

				for {

					fmt.Println("Please try again, enter a valid floating point number or type exit to leave. ")
					fmt.Scan(&userResponse)

					if userResponse == "exit" {
						os.Exit(0)
					}

					userSubsequentResponse, err := strconv.ParseFloat(checkForPresenceOfDecimalPoint(userResponse), 64)

					 if err == nil {

						userResponseAsFloat = userSubsequentResponse
					 	break
					 }
				}
			}


			fmt.Printf("You have entered the following truncated integer: %d\n", int(userResponseAsFloat))
			fmt.Printf("Enter another floating point number or type exit to leave the application.\n")

		}

	}

}

//The checkForPresenceOfDecimalPoint func will look at the user generated input and look for a period.
//The reason this func exists is to cover the edge case the user enters an integer.
//Package function strconv.ParseFloat will simple allow integers to be converted to float values.
// If this method determines the lack of a decimal point, we simply return a string stating
// "not a floating point number" which in turn will cause strconv.ParseFloat to throw an error.
func checkForPresenceOfDecimalPoint(userInput string) string {

	response := ""

	if strings.Contains(userInput, ".") {
		response = userInput
	} else {
		response = "not a floating point number"
	}

	return response
}