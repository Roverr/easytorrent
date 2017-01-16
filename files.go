package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"os/user"
	"path/filepath"
	"strings"
)

var (
	defaultTorrentDir  *string
	defaultDownloadDir *string
)

func initDirectories() {
	torrentDir := flag.String("torrentDir", getDefaultTorrentDir(), "Directory of the torrents")
	downloadDir := flag.String("downloadDir", getDefaultDownloadDir(), "Directory of downloads")
	absTorrentDir, err := filepath.Abs(*torrentDir)
	if err != nil {
		log.Fatal(err)
	}
	absDownloadDir, err := filepath.Abs(*downloadDir)
	if err != nil {
		log.Fatal(err)
	}
	defaultDownloadDir = &absDownloadDir
	defaultTorrentDir = &absTorrentDir
}

func getDefaultTorrentDir() string {
	if defaultTorrentDir != nil {
		return *defaultTorrentDir
	}
	user, err := user.Current()
	if err != nil {
		log.Fatal(err)
	}
	defaultTorrentDir := fmt.Sprintf("%s/Work/torrents", user.HomeDir)
	return defaultTorrentDir
}

func getDefaultDownloadDir() string {
	if defaultDownloadDir != nil {
		return *defaultDownloadDir
	}
	defaultDownloadDir := fmt.Sprintf("%s/downloads", getDefaultTorrentDir())
	return defaultDownloadDir
}

func selectTorrents(info []os.FileInfo) *torrentContext {
	var torrents []os.FileInfo
	for _, v := range info {
		if strings.Contains(v.Name(), ".torrent") {
			torrents = append(torrents, v)
		}
	}
	context := torrentContext{torrents: &torrents}
	return &context
}
