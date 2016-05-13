// Code generated by go-bindata.
// sources:
// persist.txt
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

var _persistTxt = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xd4\x57\x5d\x53\xe3\x36\x17\xbe\x8e\x7f\xc5\x59\xbf\x1b\xc6\x7e\xd7\x88\xe5\x96\x0e\x17\xdb\x85\x9d\x66\x66\x17\xba\x84\xb6\x17\x0c\xb3\x23\x1c\x25\xb8\x75\xac\x20\xc9\x10\xc6\xf5\x7f\xef\x39\xb2\x62\xe4\xf0\x11\x08\xcb\xd2\xe6\x26\xd6\xd7\x39\xcf\x39\x7a\xce\x23\x69\x6b\x0b\xf6\x0e\xe1\xe0\xf0\x18\xf6\xf7\x06\xc7\x70\xfc\xcb\x60\x08\x9f\x06\x9f\xf7\xe1\x4d\x80\x43\x03\x03\x99\x86\x89\x28\x84\xe2\x46\x8c\xe0\xec\x1a\x66\x42\x69\xec\x33\x52\xe6\x09\x68\x59\xaa\x54\xc0\x58\xc9\x29\x54\x15\x1b\xda\x66\x5d\xb3\x60\xc6\xd3\xbf\xf8\x44\x50\xe7\xaf\xcd\x67\x5d\x07\x41\x36\x9d\x49\x65\x20\x0a\x7a\xe1\x78\x6a\x42\xfc\x3b\xbb\x36\x42\xd3\x87\x28\x52\x39\xca\x8a\xc9\xd6\x9f\x5a\x16\xd4\x31\xc9\xcc\x79\x79\xc6\x52\x39\xdd\x9a\x48\x93\xcd\xf4\x56\x2e\x27\x38\x50\x55\x9b\xa0\x78\x81\xb6\xd9\xc0\x9a\xd3\x68\x19\x7b\x59\x5d\xd3\x90\x28\x46\x80\x1d\x71\x10\x5c\x72\x05\xdf\x60\x17\xc8\x20\xfb\xc2\x95\x3e\xe7\x79\xdb\x89\xee\xd9\x70\xa6\xb2\xc2\x8c\xdb\x3e\x7d\x91\xb3\xc3\x99\x28\x82\xc0\x5c\xcf\x08\x3a\xb0\x03\x3e\x15\x68\x0e\xb4\x51\x65\x6a\x2a\x0c\xc1\x73\xff\x45\x98\x73\x39\xd2\xe4\x2e\x18\x97\x45\x0a\xd1\xff\x71\xcd\xdb\xc5\xa2\xd8\xb7\x10\x55\x55\xb3\xea\x6d\x96\xbc\xbd\x34\xb0\xb3\x0b\x98\x18\xc5\xa7\x9a\x60\x67\x63\xec\x87\xbf\xa1\x10\xf0\xbe\xae\x13\x5c\x88\x61\xd0\x00\x4e\x65\xbf\x73\x85\x08\x9a\xef\x61\x9e\x51\x7e\x5d\xc3\xd0\x48\xf3\xdd\x26\xb9\xaa\x38\x66\xc0\xeb\x82\x90\x85\x8b\x59\xc7\x18\x17\x7d\xbb\x2c\xc5\x70\x17\xac\x23\x61\x4a\x55\xbc\x2e\xae\x2a\xe8\x5d\x94\x42\x5d\x27\xc0\xd5\x44\x13\x2c\x4b\x14\x76\x20\xae\x7e\x2e\xc7\x63\xa1\xa2\x93\x53\xea\xa9\xea\x38\x81\x93\x53\xdc\x46\xa1\xc6\x3c\xc5\xf6\xd2\x1e\x7d\x52\x7c\x32\x15\x85\x25\x89\x8d\x87\x7d\x94\x0d\x43\x7a\xd8\x40\xd6\x50\x93\xc2\x08\x1c\x79\x68\xa4\xf1\xcd\xfe\x50\x99\x11\x43\x83\x24\x99\x44\x21\xb1\xdb\x4c\x0d\x4e\x0d\x63\xcf\x01\xe6\x0d\x16\x89\xfb\x80\x50\x1b\x2e\x6e\x02\x1a\x17\x17\xc0\x86\x5f\x3f\x0f\x0a\x40\xf2\x08\x6b\xd7\x46\xe2\x51\x2f\x0a\xfb\x97\x61\x62\x2b\x27\x95\xb3\x36\x4b\x4d\x6b\x91\xa0\x26\xd1\x31\x2e\xb7\xb9\xd8\x05\x3e\x43\x92\x8e\x22\x6a\x25\x10\x56\xe1\x3b\x7d\xb2\xbd\x93\x8b\x22\xd2\xf1\xe6\xf6\xe9\xbb\xb0\x46\x84\x0d\x0a\x91\x6b\x41\x50\xda\x9c\x5b\x14\xf7\xb9\xfb\xe6\xb9\x4b\x40\x28\x45\x68\xfd\xea\x89\x1e\x0b\x94\xa2\xc7\xe5\x6f\x76\xa1\xc8\x72\xda\xcc\x5e\x0f\x8b\x97\xed\x2b\x25\x15\x06\x3d\x75\xe6\xfa\xff\xbb\x8c\x69\xa2\x54\x3b\xd0\xd7\x61\x72\x63\x9e\x75\x70\x90\xc9\xfa\x9e\xf8\x1d\xe5\x16\xa8\xa8\x81\x5b\x9a\x72\xd3\xa1\x69\x27\x21\xf5\x23\x6d\xdd\x8e\xd0\xb3\xd7\xf3\xc4\xc6\xfb\x5c\x6c\xbe\xe3\x55\xd0\xab\x5d\xcd\x78\xfc\x0a\x6c\x32\xf6\xc4\x59\x39\x89\x1a\xa6\x39\x92\xc5\xb1\x3f\x44\x98\x18\x63\x71\xc3\x67\xb2\x49\x65\x82\xd5\x88\xc4\x0a\xf9\x68\x14\x76\xfc\xfd\x56\xb8\xac\x5a\x0a\x92\xa4\x79\xa5\xdd\x52\xd4\x9f\x65\xcb\xa1\xa0\xf1\xa5\xe2\x96\x0a\xda\x3c\x40\x38\x0f\x1f\xc8\xab\x5b\x03\x4d\x35\xb6\x49\x21\x08\x23\xa1\xad\xcf\x6e\x75\xde\x89\x6a\x98\xf2\xe2\x7e\x40\x1b\x6b\x21\x42\x04\x44\xc1\x5d\x2a\xae\xe3\x39\x32\xea\x2b\x65\xfa\x48\x5e\x2d\xa5\xbc\x51\x18\xca\xb3\x85\x11\x11\x6c\x9b\xf5\x3b\x48\xec\x73\x38\x2b\xb4\x50\x26\xea\x23\x6d\x96\x69\x0c\x9e\x74\xb5\xfc\x55\x56\x55\xa9\x15\x10\x97\x5d\x13\x2d\x37\xdb\xdb\x56\xaa\xb7\xc7\x53\x3c\x10\xc7\xd7\x61\x6d\xa7\xeb\xb6\x22\x17\x01\xed\xcf\x45\x7a\x6f\x30\xab\xe0\x97\xb3\x11\x1e\xe8\x6b\xc3\xe7\x2d\x1a\x44\xc6\x30\xab\xfa\x03\xaa\x72\x8a\x57\x84\xe8\x65\x5d\xb7\x89\xe2\x64\x7e\x7b\xb5\xf1\xf9\xcc\xc2\x02\xee\xf0\xe1\x1a\x25\xaf\x12\x38\x2b\x0d\xf0\xd4\x94\x3c\xbf\x19\xea\x8f\x68\x0c\x71\x90\x60\x75\xa0\x70\x0f\x08\xe9\xf7\xc2\xdd\x7a\xe6\xad\xb9\x47\x92\x40\x89\xa9\xbc\x14\x2f\x43\x82\x91\xc8\xc5\x2b\x91\xe0\x59\xae\x57\x91\xe0\x96\xf1\xff\x3a\x09\x26\xc2\x84\x0b\x59\x9f\xd3\xde\x1f\x09\x5d\xe6\xb7\x6e\x57\x4b\xdd\xed\xb9\xec\xba\x9b\x5b\xd6\xeb\x1d\x0e\xaf\x7e\x22\x34\x47\xc2\xce\x8b\x9d\x09\xd6\xca\x93\x39\x8d\x96\x92\x45\x4d\x75\x9e\x37\xdd\x23\xbd\x07\xd0\xb3\x71\x77\x83\xb6\xfb\xde\x44\xdd\xbd\x87\xe1\xd1\x57\x88\xab\xc8\x7f\x80\xb9\xd9\xcb\xf4\x68\x78\x41\x88\xdc\x91\x69\xef\x7c\xad\xf7\xe8\x0e\xaf\x5e\x9e\x17\xf7\xb4\xa7\x60\xb3\xbe\xee\xb8\x24\x76\xa5\xbc\x05\xd0\xd7\x4b\xd7\xc4\xc7\xe0\x71\x49\xee\xd5\x9d\x5b\x89\xcb\xf9\xc6\x3c\x79\xb8\xe2\x52\x59\x16\x37\x35\x67\x5b\x80\xac\x5d\x97\x41\x1b\xd6\xc2\x0b\x11\xe8\x7d\xb2\x7c\xad\xb0\xde\x56\x44\x98\x67\xba\x09\x90\xe4\xe8\xd6\xb9\x62\xe3\x5a\xfb\x60\x79\x7e\x21\xa0\x58\xe0\x2b\xcf\x4a\x25\xfb\x98\x4b\x2d\xf0\x48\x09\x00\x7f\x56\x02\xb5\xa7\x81\xed\xd3\xb3\xed\x70\xcf\xcf\xe7\x89\xe4\x58\x3a\xef\x07\x62\x6e\x22\xfb\x22\xfd\x4e\xf2\xdb\x9b\x7b\x0f\x8f\x39\x66\x61\x63\x4e\xcf\x93\xd7\x92\x65\x70\xbf\x7f\x81\x3a\x3b\xf9\xb1\x69\x5f\x12\xde\x55\x4f\x4a\x8d\xd3\xed\x42\xa0\x8d\x7b\x22\xff\x6e\x13\xd0\x4a\xf1\xfd\x5a\xfc\x03\xa4\xf8\x07\x2a\xf1\x8a\xd4\x7e\x1f\x1d\xae\x5d\x3e\x9d\x0e\xd7\xad\x84\xb8\xfd\x46\x6f\x51\xfc\xd3\x43\xa2\xf2\xe0\x16\xe7\x1c\xd9\xbb\x8e\xce\xb8\x3e\xaa\xc3\x1b\xb9\xb4\x18\xfd\x27\xfb\x3f\x01\x00\x00\xff\xff\x0b\xbc\x96\x20\x26\x15\x00\x00")

func persistTxtBytes() ([]byte, error) {
	return bindataRead(
		_persistTxt,
		"persist.txt",
	)
}

func persistTxt() (*asset, error) {
	bytes, err := persistTxtBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "persist.txt", size: 5414, mode: os.FileMode(420), modTime: time.Unix(1463108373, 0)}
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
	"persist.txt": persistTxt,
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
	"persist.txt": &bintree{persistTxt, map[string]*bintree{}},
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

