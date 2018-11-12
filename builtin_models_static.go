// Code generated by go-bindata.
// sources:
// builtin_models/mock.yml
// DO NOT EDIT!

package mock

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

func bindataRead(data, name string) ([]byte, error) {
	gz, err := gzip.NewReader(strings.NewReader(data))
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

var _mockYml = "\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x6c\x91\x4b\x6e\xc3\x30\x0c\x44\xf7\x3a\xc5\x20\x46\x97\x2d\x92\xad\x17\x3d\x44\x6f\xa0\x38\x74\x42\xc4\xfa\x80\xa2\x1a\xf8\xf6\x85\x3e\x49\xe3\x36\x3b\x9b\xf3\xc8\x19\x52\xde\x3a\x1a\xe1\xc2\x74\xc5\x80\xf2\x83\x30\x63\x0d\x59\xe0\xc2\x89\x16\x33\x8b\x75\x74\x0b\x72\x1d\x0d\xb0\x81\x1f\x0a\xe6\x20\xd0\x0b\xf5\x0e\xe0\x9b\x24\x71\xf0\x23\x0e\x1f\xfb\x0d\xd8\x05\x4c\xc1\xab\x58\xf6\x6a\xfe\xa0\x77\x80\xfd\x1c\xc4\x59\x6d\xdf\x48\xe4\xac\x57\x9e\x1e\x7a\x53\xcd\x89\xd2\x24\x1c\xb5\x4e\xf8\x34\xc0\x17\x69\x16\x9f\x6a\x9a\xc5\x1e\x69\xc1\xae\xa4\xdd\xe1\xc6\x7a\xc1\x61\xbf\x7f\x43\x94\x70\xb4\x47\x5e\x58\xd7\x1a\xdc\xfa\x15\xec\x63\x56\xb0\xe2\x4c\x9a\xcc\x80\x85\x27\xf2\xa9\x9e\xe2\x77\xaf\x5e\x1c\x91\xbd\x50\x52\xe1\x49\xe9\x64\x86\xd6\x9c\xa0\xe1\x89\x6d\xb5\x72\xb2\x01\x33\x4b\xd2\x6e\xa1\x6b\xa4\x7f\xe7\x7a\xaf\xe5\x11\xec\xec\x99\x0c\x50\x9a\x9e\x36\xbb\xa7\x78\x9a\x53\xa1\xcd\xf2\x05\xe8\x5b\xd4\x29\x21\x6b\xcc\xda\x02\x14\xad\x1a\xf7\x41\x4d\x33\xe8\xb6\x33\x59\xcd\x42\x15\xb5\xaf\x8c\x1b\x8f\x68\xcb\x3b\x2a\x89\x79\xe1\xdd\x99\x7a\x73\xf3\x13\x00\x00\xff\xff\x61\xfd\x37\xf6\x54\x02\x00\x00"

func mockYmlBytes() ([]byte, error) {
	return bindataRead(
		_mockYml,
		"mock.yml",
	)
}

func mockYml() (*asset, error) {
	bytes, err := mockYmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "mock.yml", size: 596, mode: os.FileMode(420), modTime: time.Unix(1542041344, 0)}
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
	"mock.yml": mockYml,
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
	"mock.yml": &bintree{mockYml, map[string]*bintree{}},
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
