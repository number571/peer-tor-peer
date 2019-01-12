package utils

import (
	"os"
	"fmt"
)

func PrintError(err string) {
	fmt.Println(err)
	os.Exit(1)
}
