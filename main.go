package main

import (
	"log"
	"os"
	"os/signal"

	"github.com/ezeev/notewarrior/config"
	"github.com/ezeev/notewarrior/fslisten"
)

type Server struct {
	FsListener fslisten.Listener
}

func (s *Server) Start(conf *config.Config) {
	s.FsListener = fslisten.NewListener()
	s.FsListener.AddWatchPath(conf.MarkDownPath)
	go s.FsListener.Listen()
	go s.FsListener.ProcessEvents(HandleFsChangeEvent)
}

func (s *Server) Stop() {
	s.FsListener.Shutdown()
}

func HandleFsChangeEvent(event fslisten.FsChangeEvent) {
	log.Printf("Received event %s with type %s", event.Path, event.Type)
}

func main() {
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt)

	dir, _ := os.Getwd()
	conf, err := config.NewConfig(dir + "/config.yml")
	if err != nil {
		log.Fatal(err)
	}

	s := Server{}
	s.Start(conf)
	<-stop
	s.Stop()
	/*
		// Index
		message := struct {
			Id   string
			From string
			Body string
		}{
			Id:   "example",
			From: "marty.schoch@gmail.com",
			Body: "bleve indexing is easy",
		}

		//mapping := bleve.NewIndexMapping()
		index, err := bleve.Open("example.bleve")
		if err != nil {
			panic(err)
		}
		index.Index(message.Id, message)

		// Query

		//index, _ = bleve.Open("example.bleve")
		query := bleve.NewQueryStringQuery("bleve")
		searchRequest := bleve.NewSearchRequest(query)
		searchResult, _ := index.Search(searchRequest)

		fmt.Println(searchResult)
	*/

}
