// Code generated by go-bindata.
// sources:
// service.swagger.json
// DO NOT EDIT!

package proto

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

var _serviceSwaggerJson = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xc4\x52\xbb\x8e\xdb\x30\x10\xec\xf5\x15\x8b\x4d\x4a\xc3\x72\x5c\xba\x4f\x91\x22\x55\xca\xc0\x05\x4d\xad\x25\x1a\x36\xc9\xdb\x5d\xe9\x60\x1c\xf4\xef\x07\xd2\x0f\xca\xf2\x1d\x70\xdd\xa9\xa2\xb8\x33\x83\xd9\xe1\xbc\x55\x00\x28\xaf\xa6\x6d\x89\x71\x03\xb8\x5e\xae\x70\x91\xee\x9c\xdf\x07\xdc\x40\x9a\x03\xa0\x3a\x3d\x52\x9a\x0b\xf1\xe0\x2c\x2d\x23\x07\x0d\x19\x09\x80\x03\xb1\xb8\xe0\xd3\xfc\x7a\x04\x1f\x14\x84\x14\x2b\x80\x31\xeb\x89\xed\xe8\x44\x82\x1b\xf8\x7f\x21\x75\xaa\xf1\x26\x90\xce\x92\xb0\xdb\x8c\xb5\xc1\x4b\xff\x00\x36\x31\x1e\x9d\x35\xea\x82\xaf\x0f\x12\x7c\xc1\x46\x0e\x4d\x6f\xbf\x88\x35\xda\x49\x59\xaa\x1e\x7e\xd5\x64\xbb\xb2\x66\x82\x04\xd1\xc9\x3f\x00\x86\x48\x9c\xc5\xfe\x34\x69\xc1\xdf\x89\xb0\x28\x63\x26\x89\xc1\x0b\xc9\x03\x0b\x00\xd7\xab\xd5\xec\x0a\x00\x1b\x12\xcb\x2e\xea\x35\xad\x89\x50\x1e\xe7\x90\xcc\x13\x0d\x00\x7f\x32\xed\x13\xe3\x47\xdd\xd0\xde\x79\x97\x14\xa4\xce\xaf\x90\x1c\xfd\x25\x11\xd3\x12\x3e\xd0\xc6\xea\xa3\xf3\x38\x31\x1f\x0d\x9b\x13\x29\x71\x89\xef\xf2\xcd\x6c\x7b\x73\xca\xaf\xbf\x0b\xcd\x79\xee\xd9\xf9\xcf\x26\x4c\x2f\xbd\x63\x4a\xb1\x29\xf7\xf4\x0d\xbb\x6e\x27\xbb\xaa\x69\xe7\x5b\xe6\xc7\xfc\x77\x69\x74\x91\xdb\x56\x53\x99\xf1\xde\xe0\x89\x99\xd2\xa1\x27\x57\x93\x2e\xe9\x39\xe6\xd4\xc2\xee\x40\x56\xef\xe9\x24\x4e\x24\x56\x37\xeb\x0c\x0e\xe6\xd8\xd3\xbc\x46\x37\x11\x51\x76\xbe\x2d\x2e\xc7\x67\x97\xd5\x58\xbd\x07\x00\x00\xff\xff\xae\xd0\x28\x93\xcf\x03\x00\x00")

func serviceSwaggerJsonBytes() ([]byte, error) {
	return bindataRead(
		_serviceSwaggerJson,
		"service.swagger.json",
	)
}

func serviceSwaggerJson() (*asset, error) {
	bytes, err := serviceSwaggerJsonBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "service.swagger.json", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
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
	"service.swagger.json": serviceSwaggerJson,
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
	"service.swagger.json": &bintree{serviceSwaggerJson, map[string]*bintree{}},
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

