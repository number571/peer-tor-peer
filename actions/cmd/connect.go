package cmd

import (
	"fmt"
	"strings"
	"../../utils"
	"../../settings"
)

func Connect(splited []string) {
	for _, value := range splited {
		settings.Addresses = append(settings.Addresses, value)
		fmt.Println("| Connecting:", value)
	}
}

func ConnectFromConfig() {
	var data = utils.ReadFile(settings.CONFIG_FILE)
	var splited = strings.Split(data, "\n")

	for _, value := range splited {
		var flag bool
		for _, address := range settings.Addresses {
			if value == address {
				flag = true
				break
			}
		}

		if !flag {
			settings.Addresses = append(settings.Addresses, value)
			fmt.Println("| Connecting:", value)
		}
	}
}
