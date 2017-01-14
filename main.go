package main

import (
	"io/ioutil"
	"log"
)

func main() {
	initDirectories()
	service := initService()
	defer service.Client.Close()
	info, err := ioutil.ReadDir(*defaultTorrentDir)
	if err != nil {
		log.Fatal(err)
	}
	torrents := selectTorrents(info)
	service.Tc = torrents
	initiateUserCommunication(service.Tc)
	if finish := service.DownloadChoosenTorrents(); !finish {
		log.Fatal("Something went wrong while downloading, couldn't wait all files.")
	}
}
