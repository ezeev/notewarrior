package fslisten

import (
	"log"
	"time"

	fsnotify "gopkg.in/fsnotify.v1"
)

type Listener interface {
	Listen()
	ProcessEvents(fn func(FsChangeEvent))
	AddEvent(FsChangeEvent)
	AddWatchPath(path string) error
	Shutdown()
}

func NewListener() Listener {
	l := FsListener{}
	l.FsChangeChan = make(chan FsChangeEvent)
	var err error
	l.FsWatcher, err = fsnotify.NewWatcher()
	if err != nil {
		log.Fatal(err)
	}
	return &l
}

type FsChangeEvent struct {
	Path string
	Type string
}

type FsListener struct {
	FsChangeChan chan FsChangeEvent
	FsWatcher    *fsnotify.Watcher
}

func (fs *FsListener) ProcessEvents(fn func(FsChangeEvent)) {
	for event := range fs.FsChangeChan {
		fn(event)
	}
}

func (fs *FsListener) AddWatchPath(path string) error {
	return fs.FsWatcher.Add(path)
}

func (fs *FsListener) Listen() {

	// start listener
	/*watcher, err := fsnotify.NewWatcher()
	if err != nil {
		log.Fatal(err)
	}
	defer watcher.Close()*/

	/*if err := watcher.Add(path); err != nil {
		log.Fatal(err)
	}*/
	for {
		select {
		// watch for events
		case event := <-fs.FsWatcher.Events:
			fs.AddEvent(FsChangeEvent{Path: event.Name, Type: event.Op.String()})

		// watch for errors
		case err := <-fs.FsWatcher.Errors:
			log.Printf("ERROR: %s", err.Error())
		}
	}
}

func (fs *FsListener) AddEvent(event FsChangeEvent) {
	fs.FsChangeChan <- event
}

func (fs *FsListener) Shutdown() {
	// wait for new events to flush
	log.Print("Shutting down file system listener")
	time.Sleep(time.Second * 5)
	fs.FsWatcher.Close()
}
