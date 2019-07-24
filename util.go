package gql

import (
	"path/filepath"
	"io/ioutil"
	"log"
)

func MakeRequest(filename string) (req *Request) {
	req = NewRequest(readfile(filename))
	return
}

func readfile(filename string) (data string) {
	abspath, _ := filepath.Abs(filename)

	bytes, err := ioutil.ReadFile(abspath)
	if err != nil {
		log.Fatal(err)
	}

	data = string(bytes)
	return
}
