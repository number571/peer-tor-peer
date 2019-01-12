package utils

import "os"

func AppendFile(name, data string) {
	file, err := os.OpenFile(name, os.O_WRONLY|os.O_APPEND, 0777)
	CheckError(err)
	defer file.Close()

	file.WriteString(data)
}
