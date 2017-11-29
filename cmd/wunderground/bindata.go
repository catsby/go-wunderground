// Code generated by go-bindata.
// sources:
// templates/.list.swp
// templates/list
// DO NOT EDIT!

package main

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

var _templatesListSwp = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xec\xd9\x31\x4e\xfb\x30\x14\x06\xf0\xaf\x7f\xfd\x57\x14\x44\x47\x96\x40\xe7\xc6\x05\x96\x9e\x00\x89\x81\xad\x74\x60\x4b\x63\x13\x22\x92\x38\x4a\x5e\x04\x92\x65\xb8\x09\x97\x61\x82\xb1\x77\xe8\x0d\x38\x00\x52\x1a\x90\x60\xc9\x46\x85\xf4\xfd\x96\x27\xbf\x67\x59\xdf\xe2\xc1\xf2\x6a\xb6\xbc\xb8\x0c\xe7\xd1\x0c\x00\xf6\x81\x97\xe5\xe1\xf5\x66\x02\xac\xc7\x40\x6e\xd3\xb8\xbc\xc3\xa0\x24\xb7\xad\x16\x5b\x0d\xed\x7b\xdc\x1e\xa8\x52\xab\x9a\x3a\x51\x69\x26\xb7\xed\x2a\x4a\x6c\xa1\xbe\x06\xd3\xfb\xb6\xd4\xa6\x4e\x6b\xdb\x96\x5a\x25\x85\x56\xdf\x1a\x62\x8a\x2a\x8f\xc5\x34\x2a\xcf\x1a\x19\x4e\x46\x44\x9d\x56\x6e\xa6\xf3\x3d\x9c\x9d\x9e\x74\x57\x7d\x72\x7c\x14\x8e\x0f\xae\x76\x9d\x8a\x88\x88\x88\x88\x88\x7e\x91\x54\x23\x3c\x01\xf8\xd7\xaf\xff\xf7\x75\xf4\xa3\x12\x11\x11\x11\x11\x11\x11\xd1\xdf\x15\x6b\xe0\x39\x00\xde\x82\xed\xff\xff\xe7\xfb\xff\x3d\x00\x36\x01\xb0\xee\x67\xaf\xc1\x8e\x83\x12\x11\x11\x11\x11\x11\x11\x11\x00\xe7\x4c\xa9\xbd\x47\xd8\x71\x2e\x3a\x4f\x64\x61\x1e\xa4\x6f\x39\x17\x2d\x32\xc9\x8d\xf7\x70\xae\x8e\xcb\xd4\x84\x91\xf7\xf8\x08\x00\x00\xff\xff\xe6\x63\x2a\x18\x00\x30\x00\x00")

func templatesListSwpBytes() ([]byte, error) {
	return bindataRead(
		_templatesListSwp,
		"templates/.list.swp",
	)
}

func templatesListSwp() (*asset, error) {
	bytes, err := templatesListSwpBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/.list.swp", size: 12288, mode: os.FileMode(420), modTime: time.Unix(1511765608, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templatesList = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xaa\xae\x2e\x4a\xcc\x4b\x4f\x55\xd0\xab\xad\xe5\x52\x50\x50\x50\xa8\xae\xd6\x0b\xc9\x2c\xc9\x49\x85\x72\xc1\x02\x6e\xc9\x25\x21\xa9\x15\x25\xb5\xb5\x5c\xd5\xd5\xa9\x79\x29\xb5\xb5\x5c\x80\x00\x00\x00\xff\xff\xe0\xe6\x4b\x29\x36\x00\x00\x00")

func templatesListBytes() ([]byte, error) {
	return bindataRead(
		_templatesList,
		"templates/list",
	)
}

func templatesList() (*asset, error) {
	bytes, err := templatesListBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/list", size: 54, mode: os.FileMode(436), modTime: time.Unix(1511765608, 0)}
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
	"templates/.list.swp": templatesListSwp,
	"templates/list":      templatesList,
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
	"templates": &bintree{nil, map[string]*bintree{
		".list.swp": &bintree{templatesListSwp, map[string]*bintree{}},
		"list":      &bintree{templatesList, map[string]*bintree{}},
	}},
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
