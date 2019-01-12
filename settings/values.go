package settings

const (
	HOST = "127.0.0.1"
	PORT = ":80"

	SEPARATOR = " "

	CHOICE_DIR   = "/www/onion/"
	SETTINGS_DIR = CHOICE_DIR + "settings/"

	CONFIG_FILE  = SETTINGS_DIR + "config.txt"
	BLACK_LIST	 = SETTINGS_DIR + "black.list"

	WAITING_TIME = 5
	BUFF_SIZE = 512

	TORS_PROXY   = "socks5://127.0.0.1:9050"
	CONTENT_TYPE = "text/plain"

	EXIT 		= ":exit"
	HELP		= ":help"
	WHOAMI		= ":whoami"
	CONFIG		= ":config"
	NETWORK 	= ":network"
	CONNECT 	= ":connect"
	DISCONNECT 	= ":disconnect"
)

var (
	Addresses []string
	Hostname string
)
