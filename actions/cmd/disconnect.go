package cmd

import (
	"fmt"
	"../../utils"
	"../../settings"
)

func Disconnect(splited []string) {
	for _, value := range splited {
		for index, address := range settings.Addresses {
			if value == address {
				settings.Addresses = utils.Remove(settings.Addresses, index)
				fmt.Println("| Disconnecting:", address)
				break
			}
		}
	}
}

func DisconnectAll() {
	if len(settings.Addresses) > 0 {
		for _, address := range settings.Addresses {
			fmt.Println("| Disconnecting:", address)
		}
	}
	settings.Addresses = []string{}
}
