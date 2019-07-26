package gql

import (
	"io/ioutil"
	"net/http"
	"path/filepath"
	"strings"
)

// ReadRequest creates a new request using a path, that can be am URL
// to download the file, or direct abstract path to file.
func ReadRequest(path string) (req *Request) {
	var content string
	var err error

	if strings.HasPrefix(path, "http") {
		content, err = readurl(path)
	} else {
		content, err = readfile(path)
	}

	req = NewRequest(content).Report(err)

	return
}

// readfile get content from a file abstract path.
func readfile(filename string) (data string, err error) {
	var abspath string
	abspath, err = filepath.Abs(filename)
	if err != nil {
		return
	}

	var bytes []byte
	bytes, err = ioutil.ReadFile(abspath)
	if err != nil {
		return
	}

	data = string(bytes)
	return
}

// readurl get content from a file url.
func readurl(url string) (data string, err error) {
	var resp *http.Response
	resp, err = http.Get(url)
	if err != nil {
		return
	}

	defer resp.Body.Close()

	// read data from url
	var bytes []byte
	bytes, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}

	data = string(bytes)
	return
}

// LoadRequest creates a new request using a flisystem, a path to the
// .gql file. It retrieves the file and use to build a Request.
func LoadRequest(fst http.FileSystem, path string) (req *Request) {
	var content string
	var err error

	content, err = readflsys(fst, path)

	req = NewRequest(content).Report(err)
	return
}

// readflsys get content from a file system.
func readflsys(fst http.FileSystem, path string) (data string, err error) {
	var file http.File
	if file, err = fst.Open(path); err != nil {
		return
	}

	var bytes []byte
	if bytes, err = ioutil.ReadAll(file); err != nil {
		return
	}

	data = string(bytes)
	return
}
