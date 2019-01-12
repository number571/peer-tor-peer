package utils

func Remove(splited []string, index int) []string {
	return append(splited[:index], splited[index+1:]...)
}
