package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/fatih/color"
)

var question = color.New(color.FgCyan).Add(color.Underline)
var indexer = color.New(color.FgHiBlue)
var namer = color.New(color.FgHiYellow)
var answer = color.New(color.Bold).Add(color.Underline)

func printTorrentName(index int, file os.FileInfo) {
	fmt.Println()
	indexer.Print(index)
	fmt.Print(" for ")
	namer.Print(file.Name())
}

func getAnswers() string {
	answer.Println("\nTorrents I would like to download separated by comma:")
	reader := bufio.NewReader(os.Stdin)
	text, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	return strings.Replace(text, "\n", "", -1)
}

func printFinalAnswers(torrents *torrentContext) {
	answer.Println("\nYou have selected these torrents:")
	torrents.printChoosenTorrents()
}

func initiateUserCommunication(torrents *torrentContext) {
	question.Println("\nWhich torrents would you like to download?")
	torrents.printTorrents()
	answers := getAnswers()
	torrents.setChoosenOnes(answers)
	printFinalAnswers(torrents)
}
