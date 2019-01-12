package cmd

import (
	"fmt"
	"../../settings"
)

func Network() {
	if len(settings.Addresses) == 0 {
		fmt.Println("| Nothing")
	} else {
		for index, address := range settings.Addresses {
			fmt.Printf("| %d: %s\n", index, address)
		}
	}
}
