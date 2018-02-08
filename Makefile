linux:
	GOOS=linux GOARCH=amd64 go build -o="bin/easytorrent-linux-amd64"

osx:
	GOOS=darwin GOARCH=amd64 go build -o="bin/easytorrent-darwin-amd64"

windows:
	GOOS=windows GOARCH=amd64 go build -o="bin/easytorrent-windows-amd64.exe"

PHONY: all
all:
	$(MAKE) linux && $(MAKE) osx && $(MAKE) windows