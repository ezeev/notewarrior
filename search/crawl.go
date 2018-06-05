package search

import (
	"fmt"
	"io/ioutil"
)

func TraverseDir(path string, level int) error {

	files, err := ioutil.ReadDir(path)
	if err != nil {
		return err
	}

	for _, f := range files {
		if f.IsDir() {
			TraverseDir(path+"/"+f.Name(), level+1)
		} else {
			prefix := ""
			for i := 0; i < level; i++ {
				prefix = prefix + "-"
			}
			fmt.Println(prefix + " " + f.Name())
		}
	}

	return nil

}
