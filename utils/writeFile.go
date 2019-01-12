package utils

import "os"

func WriteFile(name, data string) {
	file, err := os.Create(name)
	CheckError(err)
	defer file.Close()

	file.WriteString(data)
}
