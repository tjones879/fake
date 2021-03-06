package util

import (
	"bufio"
	"compress/gzip"
	"fmt"
	"github.com/cespare/xxhash"
	"github.com/satori/go.uuid"
	"github.com/tjones879/fake/structs"
	"io"
	"io/ioutil"
	"net/url"
	"os"
	"strconv"
	"strings"
)

var rootDir = "/home/jones/fake"

func getFileHash(contents string) (hash uint64) {
	hash = xxhash.Sum64String(contents)
	return
}

/*
Take the hash of a file and get its correct file path.
Pad the front of the hash with 0's and split into 5
groups of 4 digits.
*/
func getPathByHash(hash uint64) (path string) {
	var filePath [5]string
	j := 0
	res := ""
	hashStr := fmt.Sprintf("%020d", hash)

	for i, r := range hashStr {
		res = res + string(r)
		if i > 0 && (i+1)%4 == 0 {
			filePath[j] = res
			res = ""
			j++
		}
	}

	path = rootDir
	for _, s := range filePath {
		path += "/" + s
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

/*
Decompress file
*/
func decompressFile(buf io.Reader) string {
	zr, err := gzip.NewReader(buf)
	if err != nil {
		fmt.Println("1)decompressFile:", err)
	}
	contents, err := ioutil.ReadAll(zr)
	if err != nil {
		fmt.Println("2)decompressFile:", err)
	}
	zr.Close()
	return string(contents)
}

// CreateStorage TODO
func CreateStorage(hash uint64, uri, contents string) (file structs.FileStorage) {
	file.Hash = strconv.FormatUint(hash, 10)
	file.Directory = getPathByHash(hash)
	file.Name = getFileNameByURL(uri)
	file.Contents = contents
	file.UID = uuid.NewV4().String()
	return
}

// SaveFile TODO
func SaveFile(f *structs.FileStorage) {
	if !f.FileExists() {
		os.MkdirAll(f.Directory, 0777)
		fd, err := os.OpenFile(f.Directory+"/"+f.Name, os.O_CREATE|os.O_WRONLY, 0666)
		if err != nil {
			fmt.Println("SaveFile:", err)
		}
		writer := bufio.NewWriter(fd)
		compressFile(writer, f.Contents)
		_ = writer.Flush()
		fd.Close()
	}
}

/*
LoadFile will load and decompress a file by name and hash
*/
func LoadFile(f *structs.FileStorage) {
	if f.FileExists() {
		fd, err := os.OpenFile(f.Directory+"/"+f.Name, os.O_RDONLY, 0666)
		if err != nil {
			fmt.Println("LoadFile:", err)
		}
		f.Contents = decompressFile(fd)
		fd.Close()
	}
}
