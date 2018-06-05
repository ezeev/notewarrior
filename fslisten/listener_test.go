package fslisten

import (
	"io/ioutil"
	"os"
	"testing"
	"time"
)

func TestListen(t *testing.T) {

	l := NewListener()

	dir, _ := os.Getwd()
	t.Logf("Using working directory: %s", dir)

	l.AddWatchPath(dir)

	// start listener
	go l.Listen()

	removeEvents := 0
	// start event processor
	go l.ProcessEvents(func(event FsChangeEvent) {
		//fmt.Println(event.Path + " : " + event.Type)
		t.Log(event.Path + " : " + event.Type)
		if event.Type == "REMOVE" {
			removeEvents++
		}
	})

	// remove text files
	os.Remove(dir + "/test1.txt")
	os.Remove(dir + "/test2.txt")

	// write a couple of text files
	ioutil.WriteFile(dir+"/test1.txt", []byte("this is a test"), 0644)
	ioutil.WriteFile(dir+"/test2.txt", []byte("this is a test again"), 0644)
	time.Sleep(time.Second * 2)

	// remove text files
	os.Remove(dir + "/test1.txt")
	os.Remove(dir + "/test2.txt")
	time.Sleep(time.Second * 2)

	if removeEvents != 2 {
		t.Error("2 remove events were not detected!")
	}

}
