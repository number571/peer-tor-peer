CC  = gcc
CGO = /usr/local/go/bin/go

DIR = /www/onion/
KDIR = /var/lib/tor/onion/

CDIR = compile/
CCDIR = actions/setup/
CCCDIR = actions/setup/append/

TORS = torSetting
P2RP = clientP2rP

.PHONY: default build install run clean delete rebuild reinstall reboot

default: build

build:
	mkdir -p bin
	$(CGO) get golang.org/x/net/proxy
	$(CGO) build -ldflags "-s -w" -o bin/$(P2RP) $(CDIR)main.go

	$(CC) -O2 -c -o $(CCDIR)configSetting.o $(CCDIR)configSetting.c
	$(CC) -O2 -c -o $(CCDIR)copyHostname.o $(CCDIR)copyHostname.c 

	$(CC) -O2 -c -o $(CCCDIR)dirwork.o $(CCCDIR)dirwork.c
	$(CC) -O2 -c -o $(CCCDIR)filework.o $(CCCDIR)filework.c
	$(CC) -O2 -c -o $(CCDIR)initialSetup.o $(CCDIR)initialSetup.c

	$(CC) -O2 -o bin/$(TORS) \
		$(CDIR)main.c \
		$(CCDIR)initialSetup.o \
		$(CCDIR)configSetting.o \
		$(CCDIR)copyHostname.o \
		$(CCCDIR)dirwork.o \
		$(CCCDIR)filework.o

install:
	sudo ./bin/$(TORS)
	sudo cp bin/$(P2RP) $(DIR)
	sudo mkdir -p $(DIR)settings/
	sudo touch $(DIR)settings/config.txt

run:
	sudo $(DIR)$(P2RP)

clean:
	rm -rf $(CCDIR)*.o
	rm -rf $(CCCDIR)*.o
	rm -rf bin

delete:
	sudo rm -rf $(DIR)
	sudo rm -rf $(KDIR)

rebuild: clean build
reinstall: rebuild install
reboot: delete reinstall
