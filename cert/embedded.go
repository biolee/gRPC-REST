// Code generated by go-bindata.
// sources:
// server.key
// server.pem
// DO NOT EDIT!

package cert

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func bindataRead(data []byte, name string) ([]byte, error) {
	gz, err := gzip.NewReader(bytes.NewBuffer(data))
	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}

	var buf bytes.Buffer
	_, err = io.Copy(&buf, gz)
	clErr := gz.Close()

	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}
	if clErr != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

type asset struct {
	bytes []byte
	info  os.FileInfo
}

type bindataFileInfo struct {
	name    string
	size    int64
	mode    os.FileMode
	modTime time.Time
}

func (fi bindataFileInfo) Name() string {
	return fi.name
}
func (fi bindataFileInfo) Size() int64 {
	return fi.size
}
func (fi bindataFileInfo) Mode() os.FileMode {
	return fi.mode
}
func (fi bindataFileInfo) ModTime() time.Time {
	return fi.modTime
}
func (fi bindataFileInfo) IsDir() bool {
	return false
}
func (fi bindataFileInfo) Sys() interface{} {
	return nil
}

var _serverKey = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x6c\xd5\xb9\x16\xaa\x4a\x02\x85\xe1\x9c\xa7\x38\x39\xab\x17\x88\x20\x12\x74\x50\x40\xc9\x28\x50\x05\x32\x65\x20\xa3\xcc\xca\xfc\xf4\xbd\xfa\xdc\xf4\xee\x74\x27\x7f\xf6\xfd\xe7\xff\x13\xa1\xa2\x59\x7f\xb0\x0b\xfe\x38\x58\xf3\x81\x07\xff\x18\x30\xfa\xfb\x10\x4f\x4d\x83\x43\xa9\x89\x00\x18\x12\x40\x10\x30\x06\x54\x19\xc9\x7c\xfa\x9d\x63\x2a\xdb\x02\xf7\x50\x7e\x35\xac\x97\x21\xad\x6f\x64\x1a\x80\x9c\x61\x92\x4b\xbd\x1b\x7b\x87\x33\x37\x6c\x0a\xa2\x95\x8d\x85\x7d\xab\x26\x6d\x71\x43\xa5\x8b\x43\x6f\x52\x4b\xa6\x04\x8d\x76\x03\x1a\x77\x69\x57\x9d\x63\xcf\x56\x70\x77\x67\x88\xb3\xe0\x6c\xb2\x77\xc2\xa9\x82\xb4\x51\x3c\xfc\x5d\xdd\xe8\xc8\x88\xdc\x17\x1b\x38\x5c\xd9\x8b\xcc\x56\xf1\x34\x3e\xf4\x27\xb9\x3a\x38\x97\xd0\xe1\x49\xaa\xa1\xe0\x5b\xd8\x77\xd6\x1d\x0a\x3e\x83\x2b\xcf\x18\xeb\xa4\x9c\x0a\x77\xb4\xba\x0f\x7d\xc7\xe5\x75\xb8\x12\x60\xe4\xb5\x33\xc1\xb8\x9f\x3a\xe5\x30\xcc\x52\x87\x12\x53\xc1\xf0\xc3\xf7\xda\xaa\xce\x63\xc7\xba\xd2\xa7\x63\xc6\x9d\x4d\xb1\x81\xf0\x8e\xe1\xb2\x7c\x6c\xa5\x98\x44\xb9\xba\x2d\x11\x7a\x10\x26\x7b\x73\x34\x12\x19\x81\x6c\x71\xba\x40\x16\xd8\x7e\x09\x11\xad\xe2\xc7\xc1\xeb\xbd\xd7\x58\x87\xb7\xbc\xf6\xea\x23\x72\x5f\xc6\xc0\x39\xe2\xdf\x8a\x32\xbe\xf4\x13\x6b\xc7\xf1\x92\x47\x6a\x21\xae\xa4\xc5\xdd\xc1\x70\x20\xb5\xa0\xd2\x70\x0c\x20\x3d\xb8\x85\x5d\x6d\x16\x10\xfd\x04\x07\x41\xb9\xdb\x0e\xb3\x69\x32\x40\x40\x04\x83\x26\x02\x78\xac\xb3\xf9\x5e\xba\x2e\xbb\x7d\xe1\xaa\x10\x93\xd8\xdd\x05\xd6\x11\x63\xac\x78\xeb\xaf\xec\x11\xc0\xfb\xa4\x72\x03\xcf\x92\xe6\x4c\x8d\x53\xd5\x59\x48\xce\xea\x2c\xaa\xdb\xa5\x69\x1f\xa9\x97\xd0\xeb\xae\x0f\x2b\x8b\x7e\x7c\x4e\xce\x13\xa1\x46\x82\x46\xa7\x5f\xeb\xae\xca\xf8\x4a\x4d\xa6\x1a\x3d\x00\xed\x14\xd4\xb3\x56\x3d\x6a\xe4\xa6\xc3\xc2\x3a\x72\x1d\x14\x76\x21\x54\x9e\x7c\x61\x09\xdf\xcd\xe8\x3f\x3e\x99\xbc\x4d\x47\xda\x3a\xa2\xaa\x65\x30\xb6\x7c\x90\x44\x52\x97\xc0\x39\xad\xd2\x19\xe2\x8b\x40\xc9\x93\xf0\x13\x99\x75\x7b\x30\x55\xb0\x26\xee\xdd\x6b\xc6\x67\x7e\x1b\x66\xc7\xb7\xc9\x6b\x75\x7f\xc6\xc5\x57\x82\x48\x8a\x88\x9f\x31\xc6\x58\x87\x4a\xa2\x2a\x21\x69\xe7\x17\xd6\x6a\x3b\x9b\x0c\xaa\x51\xc2\x42\x14\x16\xf3\xc4\xc1\x9e\xa2\xce\xb8\xce\x39\x2f\xac\xdd\x27\x27\xdc\x34\xf9\x5b\xbf\xa2\x42\x18\x83\xaf\x20\x10\x61\x3e\x5c\xf4\xb2\x3c\xbe\x4a\x83\xd6\x4e\x29\x6a\x6e\xe1\xa6\xb1\xce\x43\x16\x89\x1f\xd1\xc8\xdc\xc9\xae\xea\xe4\x65\x24\x26\x12\xbc\x1b\xd9\xf5\xf5\x67\xbc\x4a\x45\x4b\xa9\xa0\xa9\x5c\x2e\x23\x30\x7a\x77\xfa\x06\xa5\x32\x82\xe0\x46\x96\x60\x05\xab\x93\x04\x7b\xee\x45\x4c\x55\x1f\x71\x61\xf2\xf3\x37\xdd\x74\xf0\xf9\x5a\xce\x2e\xd5\xdd\x2d\xd7\xf5\x38\xad\xc4\x23\xec\x15\x7e\xf0\x3d\x42\x8f\xd1\x40\x8f\x56\x02\x29\xb8\xff\xce\x95\x16\x87\x64\x7f\x61\xc3\x52\x9f\x5c\xf7\xba\x65\x30\xc6\xe5\x96\xd4\x11\xfb\x7a\x2d\xc6\x61\x30\x76\xe3\xaa\x70\xb9\xff\xec\x6d\x0c\xdf\x8a\xc6\x37\xc4\x92\x7d\x60\x8a\x93\x3c\x12\xdf\x16\xe8\x44\x86\x9b\xca\xfa\xb7\xa4\x9f\x74\xcd\x65\xbd\xde\x8c\x4b\x95\x2a\x82\x47\x8e\x6e\x7f\x85\xa4\x3f\xcc\xf8\xc5\xf8\xcd\xf5\xf5\x4f\xf1\xee\x2e\xc4\xb6\xbb\x13\x2f\xb4\xad\xcd\x70\x36\x27\xa2\xab\x20\xab\x46\xe1\x39\x79\x7a\x33\x44\x9c\xcd\x92\xf6\xbc\x98\xc2\x03\xa9\x9c\xf4\xdb\x7c\xf4\x5e\x2f\xe2\x36\xad\xf3\xea\xf0\x0c\xfb\x46\x66\x83\x09\xae\x70\xd1\xa1\x70\x8a\x30\xa9\x61\x57\x31\xd9\x47\x18\xc5\x1b\x57\xd1\xdc\xd6\xda\xe4\x83\xbe\xca\xce\x65\x44\x22\xc3\xa1\xa3\x1e\xca\xe3\x1b\x6c\x56\x17\x28\xfd\x37\xfb\xe9\x78\xb8\xb4\xa0\x22\xc6\x14\x06\x0f\x87\xa4\x98\x5b\x4f\xfa\xf3\x43\x7c\xa7\xf2\x5b\x6a\xd1\x6f\x75\x4c\xba\xb4\x4f\xbe\xaa\xc1\xb2\x82\xbb\x54\x46\xa2\xe0\x46\x9b\xf1\x5a\xac\x5e\x04\x67\x9c\x5d\x4e\x17\x49\x32\x26\xca\x45\x13\x38\x9d\x43\xd3\xe5\x76\x2e\xde\xfc\x2e\xf8\xbc\xcf\xa0\xfe\xb4\xb4\x84\x4d\x7e\x93\xc9\x9b\x2a\xd7\x8a\x2e\xa4\xe6\xf0\xfe\xb6\xce\x14\xf6\x6a\xf6\x9b\x0d\x43\xa0\xe6\x57\xd6\x58\x01\xb1\xc7\x00\xdc\xa5\xec\x7e\x09\x8b\x7d\x51\x7e\x82\x8a\x4a\x5f\x56\xbd\xb7\x7f\x24\xbc\x27\xd7\x55\x43\x63\x07\x5c\x35\x21\x69\xa4\xa8\xd4\xf8\xd6\x0f\x26\x0d\xbb\xb6\xc6\xc7\x77\xaf\xb6\x5c\x40\xe4\xe7\xee\x61\x72\xf8\x4a\xfc\xb1\x4a\xe1\x2c\x34\xc8\xcb\x72\x64\x88\xe5\x83\x39\xc9\xec\x21\x9d\xe6\xf6\x7e\xea\x72\x70\x7f\xf8\x5e\x95\x45\x73\x31\xa1\x30\xbe\xdb\xd5\x2e\x2c\x69\xe9\xb6\x03\x21\xdd\x18\xec\xb9\xd4\x41\x1f\x43\x15\x7c\xb6\xab\x52\xa6\xea\x16\xd9\xde\x18\xb4\xa9\xc2\x8f\x5a\x67\xb7\xfc\xc0\xe6\xf3\xc7\x65\x73\x27\xe7\x23\xeb\xbd\x93\x30\x79\x8a\x94\x30\xc7\x82\x9e\xf8\xc4\x5d\xa6\xab\x80\x8a\xdf\x2a\x9b\x55\x26\x4a\x3f\x02\x44\x26\xcd\x01\xdb\x10\x73\x41\xda\xf7\x6f\xe3\x1c\x14\x6d\xf0\x67\x61\x42\x2d\x29\xeb\xe3\x15\x5b\x20\x91\xe5\xe7\xb2\xdd\x1b\x8e\x7e\x44\x84\x6f\xd4\x29\x18\x14\xf0\x34\xe3\x36\x90\x67\x49\x05\xac\xce\x29\x95\x39\xe1\xc7\x69\x9f\x33\xa3\xdd\x38\x1a\xda\x0b\xd5\x58\x95\x32\xeb\xd4\x3d\x3a\xd2\x72\x7e\xfa\xc9\x60\x71\xb4\xa3\x14\x0d\x21\x7f\x79\xd2\x73\x29\x79\xd7\xcf\x88\xaf\xe8\xc3\x4a\x1a\xb9\xbc\x19\xb1\xe1\xb1\x2f\x76\x32\x2b\x9e\xf5\x1f\x89\x2a\xf5\x40\x3e\x50\x81\x7a\xdd\x20\x2f\x45\x51\x59\x55\x96\x0f\xee\x39\x92\x26\x71\x8e\xf2\xcb\x87\xa1\x7a\xcb\x61\x8e\x83\xee\xd8\x87\xe0\x13\xf4\xbf\x69\x59\xd2\x9c\x2e\xb1\x51\x30\x22\x22\x8f\xe8\xd3\x90\xd2\x6b\x33\x29\x30\x35\xec\x7f\x89\xbf\xa4\x40\x4b\xfe\x77\x6a\xfe\x17\x00\x00\xff\xff\x05\xe8\xe7\xb5\x8b\x06\x00\x00")

func serverKeyBytes() ([]byte, error) {
	return bindataRead(
		_serverKey,
		"server.key",
	)
}

func serverKey() (*asset, error) {
	bytes, err := serverKeyBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "server.key", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _serverPem = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x94\x94\xc9\x8e\xab\x38\x1b\x86\xf7\x5c\xc5\xbf\x47\xbf\xe2\x40\x20\xc5\xa2\x17\xb6\x31\x04\x12\x53\xcc\x01\x76\x84\x31\x0c\xa1\x08\x49\x0c\x5c\x7d\xab\x4a\xad\x1e\xea\x1c\xe9\xe8\x78\xf9\xfa\xb3\xf5\xd8\x7e\xfc\xfd\xff\x73\x20\xa2\x1b\xd6\xff\x30\x71\x7d\x43\x33\x30\xf4\xc9\x57\xca\x51\xc3\x20\xa8\xc1\x18\x3e\xe5\x0a\x32\x03\xc1\xca\x30\xe1\x69\x4d\xd5\x8c\xe4\x27\xa4\xee\x29\x04\x3a\xf6\x46\xdd\x33\x2e\xa2\xea\x10\x84\x9c\x00\x52\xed\x6d\xc6\x2b\x34\x51\x65\x85\x1c\x82\xb1\x0f\x5b\x2d\xa4\x2e\x65\xc4\x89\xd5\xd0\x71\x0c\xc2\x3e\xfc\x8b\xb0\xed\x4e\xa1\x05\xe2\xc8\xed\xa8\x47\xd8\x81\x7d\xcd\x1d\xc9\x5c\x9b\x97\x9b\xdb\x65\xbd\xd4\xe5\x18\x45\x5c\x7a\x76\x6f\x69\xe4\xae\x86\x86\x40\xe1\x21\x9a\xeb\xce\xac\xab\xf0\xfc\xb9\x3b\x82\xd4\x57\x85\xf9\x15\x0b\xda\x94\xea\xca\x9a\xab\xc3\x4c\x55\xc8\xa8\x0a\x0b\x8d\x81\xd9\x6a\xe0\xc2\x51\xbf\x62\x74\x0d\x24\x4b\x35\x53\x8d\x81\xe5\x33\xa4\x7e\xf0\x77\xf6\x6f\xdc\x9f\xd1\x72\xbf\x83\xfb\x33\x5a\xee\xd7\xb8\x18\x43\xcf\x60\xaa\x13\x9b\xc7\x21\x31\xea\x57\x66\xc1\xaf\xcb\x84\x6a\x55\x11\x9b\x83\x9f\x05\xce\x80\xab\x8a\x20\x68\x5d\x6b\xa4\x54\xd7\x55\x4b\x9a\x65\x9e\x4e\x35\x0d\xd9\xd6\x7c\x27\x62\x80\x4d\x09\x2a\x10\x1e\x1e\x8f\xfe\x11\x4f\xde\x94\xb4\x51\xd0\x49\x07\xc9\x59\xb8\xfb\xf3\x80\x16\x45\x2d\x52\xc3\x73\x52\x73\xd9\xec\x6f\x6e\xf7\x81\xdf\x2b\xfc\x9e\x27\x97\xa5\x28\xa8\x54\x06\x93\x2f\xeb\x51\x38\x49\x87\x48\x7f\xba\x1b\x67\xb2\xf9\xf9\x64\x81\x9e\xde\xa2\xce\xe1\x12\x74\xb4\x78\xcb\xe1\x8d\x73\x3a\xba\x9e\x6f\xdf\x57\x10\x55\x2d\x6d\x59\xbd\xd4\x6d\xd1\x49\x52\x53\x22\x3b\xcc\xdb\x38\x58\xc6\xb8\x1f\x8e\x37\x30\xfa\xc9\xa0\xd8\x24\xb6\x64\x8b\x1d\x0b\x6e\xa1\x42\x1b\x24\xf2\x47\x3d\x5d\xaf\x3b\xec\xc2\xc7\xe0\x6a\x3b\x5e\x5a\x4e\xf3\x3e\x4d\x9e\xa4\x1a\xa4\xc7\x98\x15\x3a\xf0\x5a\x92\x05\xee\xe9\x3e\xac\xf5\x4d\xae\x9c\x5d\x71\xef\x09\xf2\xf8\x77\xae\x59\xec\x16\x77\xd5\x5a\x78\x65\x79\x03\x7e\x67\x9f\x2d\x04\xdc\xec\xb9\x24\x40\x52\xb3\x76\xdf\xd1\x38\x66\x45\xaa\x5c\xdb\xa8\x7d\x17\x51\x3d\x86\x38\x03\x1e\x9d\x3a\x28\x97\xfb\xc7\xa6\xe1\x8a\x12\x1e\x69\x8b\xc4\x8d\xb0\x95\xc3\xda\x3a\x92\xdb\x3a\x64\xaa\x03\xc3\x73\x1b\x76\x03\xf5\x1b\x65\xc2\x90\x11\x08\xd3\x77\x34\xab\x18\x31\x1f\xe6\x9f\xaf\x76\x70\x76\x44\xab\x9c\x80\xdb\x9b\xe3\x71\xee\x5b\x61\xb3\x4b\x71\x97\x5d\x67\x51\xb8\xbc\xfc\x3b\xa0\x41\xeb\xb0\x2a\x21\x3a\xdc\x06\xb9\xc1\x3c\x74\xf5\x31\xaa\x07\x18\x7c\xaf\xe7\xfe\x59\xe0\xd5\xb1\xec\xd6\xbf\x32\xef\xbb\x78\xdc\xef\x9a\xf7\x5d\x3c\xee\x07\xf3\x8c\xef\xdf\x99\x7d\x1d\x83\x30\x47\xa3\x90\x22\x58\xbe\xfd\xc7\x4a\x0e\x3a\x44\xfb\x4b\x4b\x04\xf5\x01\xd0\x9c\xd8\x30\x94\xc9\x41\xe8\x6b\x23\x42\xe6\x45\x1e\x9b\x6a\x1f\x81\x56\x1f\xfb\x7a\x0b\x82\xb5\xb7\xb2\xcd\x7d\x67\x81\x2e\x2b\x07\x3b\x13\xb9\x51\xd9\x5a\xe2\x03\xb6\xc2\x3a\x43\xef\xbc\x95\x35\xfb\xd9\xe7\x62\xe1\x3e\x6e\x92\x96\x93\xf3\xc9\x3f\x0a\x1e\x34\x5b\x5b\xd8\x55\xf2\x56\x51\x3e\x82\x4b\x46\x5e\xeb\x21\x9d\xa4\x4d\xca\x5c\x43\xe4\x6c\x8d\x59\x26\x9f\x8d\x53\xdc\x46\xf3\x74\x06\xca\x26\x7b\xb9\x48\x1b\xe9\x78\x6f\xd3\xd7\x6b\xa0\x25\x66\x4e\x4d\x4d\xa6\xdf\xc6\x8e\x47\x85\xb5\xeb\x91\x07\x80\x59\x75\x87\xb3\x57\xaa\x3e\x57\xbc\xf9\xc2\x52\xb7\x76\x06\x78\xed\x09\x0f\xbb\xa4\xec\xf3\x24\xb8\xf0\x8b\x0d\x5f\x7c\x77\xf5\xf9\x0e\x33\xfe\x15\x90\xc9\x7a\xeb\x6f\x69\xff\xd6\x25\x47\xbc\xd6\xee\x7b\x58\xf6\xd5\x91\x1c\x28\x97\x14\x20\x6d\x94\xc7\x7a\x14\xbd\xe7\xf6\xd1\x40\xb2\x5a\xbb\xf4\x7e\x4f\x1b\x3c\x7c\xb4\x26\x14\x52\xd5\x10\xae\x6f\x24\xe1\x85\x64\xde\x8e\x97\xda\x8a\x98\x99\x8b\x44\xa1\xf1\xa4\x9c\xfb\x53\xc9\xb9\xd1\xfe\x0e\x52\xef\x2c\xc6\xca\x9d\xdf\x78\x7d\x13\x9b\xa7\x08\x89\x32\x66\x97\xac\x39\x1d\x7a\xc7\x02\x7f\x70\x5f\xfd\x97\x58\xea\x8f\x3d\xf9\xcf\x00\x00\x00\xff\xff\x4e\x2d\x12\x94\xb0\x05\x00\x00")

func serverPemBytes() ([]byte, error) {
	return bindataRead(
		_serverPem,
		"server.pem",
	)
}

func serverPem() (*asset, error) {
	bytes, err := serverPemBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "server.pem", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

// Asset loads and returns the asset for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func Asset(name string) ([]byte, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("Asset %s can't read by error: %v", name, err)
		}
		return a.bytes, nil
	}
	return nil, fmt.Errorf("Asset %s not found", name)
}

// MustAsset is like Asset but panics when Asset would return an error.
// It simplifies safe initialization of global variables.
func MustAsset(name string) []byte {
	a, err := Asset(name)
	if err != nil {
		panic("asset: Asset(" + name + "): " + err.Error())
	}

	return a
}

// AssetInfo loads and returns the asset info for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func AssetInfo(name string) (os.FileInfo, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("AssetInfo %s can't read by error: %v", name, err)
		}
		return a.info, nil
	}
	return nil, fmt.Errorf("AssetInfo %s not found", name)
}

// AssetNames returns the names of the assets.
func AssetNames() []string {
	names := make([]string, 0, len(_bindata))
	for name := range _bindata {
		names = append(names, name)
	}
	return names
}

// _bindata is a table, holding each asset generator, mapped to its name.
var _bindata = map[string]func() (*asset, error){
	"server.key": serverKey,
	"server.pem": serverPem,
}

// AssetDir returns the file names below a certain
// directory embedded in the file by go-bindata.
// For example if you run go-bindata on data/... and data contains the
// following hierarchy:
//     data/
//       foo.txt
//       img/
//         a.png
//         b.png
// then AssetDir("data") would return []string{"foo.txt", "img"}
// AssetDir("data/img") would return []string{"a.png", "b.png"}
// AssetDir("foo.txt") and AssetDir("notexist") would return an error
// AssetDir("") will return []string{"data"}.
func AssetDir(name string) ([]string, error) {
	node := _bintree
	if len(name) != 0 {
		cannonicalName := strings.Replace(name, "\\", "/", -1)
		pathList := strings.Split(cannonicalName, "/")
		for _, p := range pathList {
			node = node.Children[p]
			if node == nil {
				return nil, fmt.Errorf("Asset %s not found", name)
			}
		}
	}
	if node.Func != nil {
		return nil, fmt.Errorf("Asset %s not found", name)
	}
	rv := make([]string, 0, len(node.Children))
	for childName := range node.Children {
		rv = append(rv, childName)
	}
	return rv, nil
}

type bintree struct {
	Func     func() (*asset, error)
	Children map[string]*bintree
}
var _bintree = &bintree{nil, map[string]*bintree{
	"server.key": &bintree{serverKey, map[string]*bintree{}},
	"server.pem": &bintree{serverPem, map[string]*bintree{}},
}}

// RestoreAsset restores an asset under the given directory
func RestoreAsset(dir, name string) error {
	data, err := Asset(name)
	if err != nil {
		return err
	}
	info, err := AssetInfo(name)
	if err != nil {
		return err
	}
	err = os.MkdirAll(_filePath(dir, filepath.Dir(name)), os.FileMode(0755))
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(_filePath(dir, name), data, info.Mode())
	if err != nil {
		return err
	}
	err = os.Chtimes(_filePath(dir, name), info.ModTime(), info.ModTime())
	if err != nil {
		return err
	}
	return nil
}

// RestoreAssets restores an asset under the given directory recursively
func RestoreAssets(dir, name string) error {
	children, err := AssetDir(name)
	// File
	if err != nil {
		return RestoreAsset(dir, name)
	}
	// Dir
	for _, child := range children {
		err = RestoreAssets(dir, filepath.Join(name, child))
		if err != nil {
			return err
		}
	}
	return nil
}

func _filePath(dir, name string) string {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	return filepath.Join(append([]string{dir}, strings.Split(cannonicalName, "/")...)...)
}

