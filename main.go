package main

import (
	"bufio"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/fatih/color"
)

var question = color.New(color.FgCyan).Add(color.Underline)
var indexer = color.New(color.FgHiBlue)
var namer = color.New(color.FgHiYellow)

func printTorrentName(index int, file os.FileInfo) {
	fmt.Println()
	indexer.Print(index)
	fmt.Print(" for ")
	namer.Print(file.Name())
}

func main() {
	torrentDir := flag.String("torrentDir", "./", "Directory of the torrents")
	downloadDir := flag.String("downloadDir", "./downloads", "Directory of downloads")
	absTorrentDir, err := filepath.Abs(*torrentDir)
	if err != nil {
		log.Fatal(err)
	}
	absDownloadDir, err := filepath.Abs(*downloadDir)
	if err != nil {
		log.Fatal(err)
	}
	info, err := ioutil.ReadDir(*torrentDir)
	if err != nil {
		log.Fatal(err)
	}
	var torrents []os.FileInfo
	question.Println("\nWhich torrent would you like to download?")
	for _, v := range info {
		if strings.Contains(v.Name(), ".torrent") {
			torrents = append(torrents, v)
		}
	}
	for k, v := range torrents {
		printTorrentName(k, v)
	}
	fmt.Println()
	fmt.Println("Torrent I would like to download:")
	reader := bufio.NewReader(os.Stdin)
	text, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	replaced := strings.Replace(text, "\n", "", -1)
	choosen, err := strconv.Atoi(replaced)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(fmt.Sprintf("You have choosen %s", torrents[choosen].Name()))
	bashString := fmt.Sprintf(`#!/bin/bash -e
Taipei-Torrent -seedRatio 0 -fileDir %s %s/%s
`, absDownloadDir, absTorrentDir, torrents[choosen].Name())
	mode := os.FileMode(0770)
	if err := ioutil.WriteFile("start.sh", []byte(bashString), mode); err != nil {
		log.Fatal(err)
	}
}
