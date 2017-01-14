package main

import (
	"log"

	"github.com/anacrolix/torrent"
)

type service struct {
	Tc     *torrentContext
	Client *torrent.Client
	Config *torrent.Config
}

func initService() *service {
	config := torrent.Config{
		DataDir: *defaultDownloadDir,
		Seed:    false,
		Debug:   true,
	}
	client, err := torrent.NewClient(&config)
	if err != nil {
		log.Fatal(err)
	}
	serv := service{
		Tc:     &torrentContext{},
		Client: client,
		Config: &config,
	}
	return &serv
}

func (s *service) DownloadChoosenTorrents() bool {
	paths := s.Tc.getChoosenPaths()
	for _, v := range *paths {
		torrent, err := s.Client.AddTorrentFromFile(v)
		if err != nil {
			log.Fatal(err)
		}
		<-torrent.GotInfo()
		torrent.DownloadAll()
	}
	return s.Client.WaitAll()
}
