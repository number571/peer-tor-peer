#pragma once

#define DEBUG
#define BUFF_SIZE 512

#define MAIN_DIR   "www/"
#define CHOICE_DIR "onion/"
#define SETTINGS_DIR "settings/"

#define TORRC_FILE  "/etc/tor/torrc"
#define BLACK_LIST 	"black.list"
#define CONFIG_FILE "config.txt"
#define PRIVATE_DIR "/var/lib/tor/" CHOICE_DIR

#define HIDDEN_SERVICE_DIR  "HiddenServiceDir " PRIVATE_DIR
#define HIDDEN_SERVICE_PORT "HiddenServicePort 80 127.0.0.1:80"
