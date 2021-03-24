package utils

import (
	"fmt"
	"github.com/fdaines/spm-go/common"
)

func PrintError(message string, err error) {
	fmt.Printf("Error: %s - %s\n", message, err.Error())
}

func PrintMessage(message string) {
	fmt.Println(message)
}

func PrintVerboseMessage(message string) {
	if common.Verbose {
		fmt.Println(message)
	}
}

func PrintStep() {
	if common.Verbose {
		fmt.Print(".")
	}
}
