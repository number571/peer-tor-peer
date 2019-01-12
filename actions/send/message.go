package send

import (
	"os"
	"fmt"
	"time"
	"strings"
	"net/url"
	"net/http"
	"golang.org/x/net/proxy"
	"../cmd"
	"../../utils"
	"../../settings"
)

func Message() {
	tbProxyURL, err := url.Parse(settings.TORS_PROXY)
	utils.CheckError(err)

    tbDialer, err := proxy.FromURL(tbProxyURL, proxy.Direct)
    utils.CheckError(err)

    tbTransport := &http.Transport{ Dial: tbDialer.Dial }

    var client = &http.Client {
        Transport: tbTransport,
        Timeout: time.Second * settings.WAITING_TIME,
    }

	for {
		var message = utils.InputString()
		var splited = strings.Split(message, settings.SEPARATOR)

		switch splited[0] {
			case settings.EXIT:
				Request(client, "| {Logged_out}: " + settings.Hostname)
				fmt.Println("| Good bye:", settings.Hostname)
				os.Exit(0)

			case settings.WHOAMI: 
				fmt.Println("| Hostname:", settings.Hostname)

			case settings.CONNECT:
				if len(splited) > 1 {
					cmd.Connect(splited[1:])
					Request(client, "| {Connected}: " + settings.Hostname)
				}

			case settings.DISCONNECT:
				if len(splited) > 1 {
					switch splited[1] {
						case "*": 
							cmd.DisconnectAll()
						default: 
							cmd.Disconnect(splited[1:])
					}
				}

			case settings.HELP:	
				switch len(splited) {
					case 1: cmd.Help()
					default:
						switch splited[1] {
							case "disconnect":
								cmd.HelpDisconnect()
							case "config":
								cmd.HelpConfig()
						}
				}

			case settings.NETWORK: 
				cmd.Network()

			case settings.CONFIG: 
				if len(splited) > 1 {
					switch splited[1] {
						case "--connect", "-c":
							cmd.ConnectFromConfig()
							Request(client, "| {Connected}: " + settings.Hostname)
						case "--insert", "-i":
							cmd.PushInConfig(splited[2:])
						case "--delete", "-d":
							cmd.DeleteFromConfig(splited[2:])
						case "--spaces", "-s":
							cmd.ClearConfig()
						case "--read", "-r":
							cmd.ReadConfig()
					}
				}

			default:
				Request(client, fmt.Sprintf("| [%s]: %s", settings.Hostname, message))
		}
	}
}
