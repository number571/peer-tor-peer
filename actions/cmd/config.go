package cmd

import (
	"fmt"
	"strings"
	"../../utils"
	"../../settings"
)

func PushInConfig(addresses []string) {
	for _, address := range addresses {
		utils.AppendFile(settings.CONFIG_FILE, "\n" + address)
	}
}

func DeleteFromConfig(addresses []string) {
	var data = utils.ReadFile(settings.CONFIG_FILE)
	var splited = strings.Split(data, "\n")

	for _, value := range addresses {
		for index, address := range splited {
			if value == address {
				splited = utils.Remove(splited, index)
			}
		}
	}

	data = strings.Join(splited, "\n")
	utils.WriteFile(settings.CONFIG_FILE, data)
}

func ClearConfig() {
	var data = utils.ReadFile(settings.CONFIG_FILE)
	var splited = strings.Split(data, "\n")

	for range splited {
		for index, value := range splited {
			if value == "" {
				splited = utils.Remove(splited, index)
				break
			}
		}
	}

	data = strings.Join(splited, "\n")
	utils.WriteFile(settings.CONFIG_FILE, data)
}

func ReadConfig() {
	var data = utils.ReadFile(settings.CONFIG_FILE)
	var splited = strings.Split(data, "\n")
	for index, value := range splited {
		fmt.Printf("| %d: %s\n", index, value)
	}
}
