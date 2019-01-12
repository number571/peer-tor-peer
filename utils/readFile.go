package utils

import "os"

func ReadFile(name string) string {
	file, err := os.Open(name)
	CheckError(err)
	defer file.Close()

	var buffer = make([]byte, 512)
	var data string

	for {
		length, err := file.Read(buffer)
		if length == 0 || err != nil { break }
		data += string(buffer[:length])
	}

	return data
}
