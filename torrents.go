package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type torrentContext struct {
	torrents        *[]os.FileInfo
	choosenTorrents []*os.FileInfo
}

// Set the choosen torrent pointers
func (tc *torrentContext) setChoosenOnes(answers string) {
	var choosens []int
	stringNumbers := strings.Split(answers, ",")
	for _, stringNumber := range stringNumbers {
		choosen, err := strconv.Atoi(stringNumber)
		if err != nil {
			log.Fatal(err)
		}
		choosens = append(choosens, choosen)
	}
	torrents := *tc.torrents
	var choosenTorrents []*os.FileInfo
	for _, index := range choosens {
		choosenTorrents = append(choosenTorrents, &torrents[index])
	}
	tc.choosenTorrents = choosenTorrents
}

func (tc *torrentContext) getChoosenPaths() *[]string {
	var paths []string
	for _, pointer := range tc.choosenTorrents {
		value := *pointer
		paths = append(paths, fmt.Sprintf("%s/%s", getDefaultTorrentDir(), value.Name()))
	}
	return &paths
}

// Print out all torrents found
func (tc *torrentContext) printTorrents() {
	if tc.torrents == nil {
		log.Fatal("There are no torrents")
	}
	for k, v := range *tc.torrents {
		printTorrentName(k, v)
	}
}

// Print out only choosen torrents
func (tc *torrentContext) printChoosenTorrents() {
	for _, pointer := range tc.choosenTorrents {
		value := *pointer
		namer.Print(value.Name())
		fmt.Println()
	}
}
