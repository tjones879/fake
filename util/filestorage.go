package util

import (
	"bufio"
	"compress/gzip"
	"fmt"
	"github.com/cespare/xxhash"
	"net/url"
	"os"
	"strings"
)

var rootDir = "/home/jones/fake"

// FileStorage TODO
type FileStorage struct {
	directory [5]string
	name      string
	contents  string
}

func getFileHash(contents string) (hash uint64) {
	hash = xxhash.Sum64String(contents)
	return
}

// CreateStorage TODO
func CreateStorage(hash uint64, uri, contents string) (file FileStorage) {
	file.directory = getPathByHash(hash)
	file.name = getFileNameByURL(uri)
	file.contents = contents
	return
}

/*
Take the hash of a file and get its correct file path.
Pad the front of the hash with 0's and split into 5
groups of 4 digits.
*/
func getPathByHash(hash uint64) (filePath [5]string) {
	j := 0
	hashStr := fmt.Sprintf("%020d", hash)
	res := ""

	for i, r := range hashStr {
		res = res + string(r)
		if i > 0 && (i+1)%4 == 0 {
			filePath[j] = res
			res = ""
			j++
		}
	}
	return
}

/*
Get the file name of a page by taking the host, path, and queries.
Limit the name to the last 30 characters.
*/
func getFileNameByURL(uri string) (name string) {
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}

	link, _ := url.Parse(uri)
	runes := []rune(link.Host + link.Path + link.RawQuery)
	chars := max(len(runes)-30, 0)
	name = string(runes[chars:])
	name = strings.Replace(name, "/", "-", -1)
	name = strings.Replace(name, ".", "-", -1)
	return
}

/*
Compress a file's contents into the passed bytes buffer.
*/
func compressFile(buf *bufio.Writer, contents string) {
	zw := gzip.NewWriter(buf)
	_, err := zw.Write([]byte(contents))
	if err != nil {
		fmt.Println("compressFile:", err)
	}
	if err = zw.Close(); err != nil {
		fmt.Println("compressFile:", err)
	}
	return
}

func fileExists(f FileStorage) bool {
	_, err := os.Stat(f.RealPath())
	return !os.IsNotExist(err)
}

// RealPath TODO
func (f FileStorage) RealPath() string {
	path := rootDir
	for _, s := range f.directory {
		path += "/" + s
	}
	return path
}

// SaveFile TODO
func (f FileStorage) SaveFile() {
	if !fileExists(f) {
		os.MkdirAll(f.RealPath(), 0777)
		fd, err := os.OpenFile(f.RealPath()+"/"+f.name, os.O_CREATE|os.O_WRONLY, 0666)
		if err != nil {
			fmt.Println(err)
		}
		writer := bufio.NewWriter(fd)
		compressFile(writer, f.contents)
		_ = writer.Flush()
		fd.Close()
	}
}
