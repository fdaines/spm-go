package utils

import (
	"fmt"
	"github.com/fdaines/spm-go/common"
)

func PrintMessage(message string) {
	fmt.Println(message)
}

func PrintVerboseMessage(message string) {
	if common.Verbose {
		fmt.Println(message)
	}
}
