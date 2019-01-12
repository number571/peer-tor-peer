package cmd

import "fmt"

func Help() {
	fmt.Println(`| :exit = Exit from P2rP
| :config = Reload config
| :connect = Connect to client(s) P2rP
| :disconnect = Disconnect from client(s) P2rP
| :hostname	= Print hostname
| :network = Print connections`)
}

func HelpConfig() {
	fmt.Println(`| connect, -c = connect from config
| insert, -i = push address(es) in config
| delete, -d = delete address(es) from config
| spaces, -s = delete spaces in config
| read, -r = print addresses from config`)
}

func HelpDisconnect() {
	fmt.Println("| * = disconnect from all")
}
