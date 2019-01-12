package main

import (
	"os"
	"fmt"
	"strings"
	"net/http"
	"../utils"
	"../settings"
	"../actions/get"
	"../actions/send"
)

func main() {
	fmt.Println("| Server is listening ...")

	os.Mkdir(settings.SETTINGS_DIR, 0777)
	checkSettingDir()

	settings.Hostname = strings.Split(
		strings.Replace(utils.ReadFile(settings.CHOICE_DIR + "hostname"), "\n", "", -1),
		".",
	)[0]

	fmt.Println("| Hostname:", settings.Hostname)
	go send.Message()

	http.HandleFunc("/", get.Message)
	utils.CheckError(http.ListenAndServe(settings.HOST + settings.PORT, nil))
}

func checkSettingDir() {
	var files = [2]string{
		settings.CONFIG_FILE,
		settings.BLACK_LIST,
	}

	for _, name := range files {
		if utils.FileIsNotExist(name) {
			utils.WriteFile(name, "")
		}
	}
}
