// Code generated by go-bindata.
// sources:
// commands/service/template/logger_go
// commands/service/template/main_go
// commands/service/template/proto.proto
// DO NOT EDIT!

package service

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

var _commandsServiceTemplateLogger_go = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x5c\x91\x41\xab\xd4\x30\x14\x85\xd7\xcd\xaf\xb8\x74\x21\xed\x63\xc8\x88\x4b\xa1\x0b\xd1\xf7\x10\x79\xcc\x08\x75\x27\x22\x99\xf6\x36\x13\x26\xcd\x1d\x92\x9b\x11\x09\xf9\xef\x92\xb6\xa3\x8c\xbb\xe4\x3b\x27\x27\xe7\x26\x57\x35\x5c\x94\x46\x48\x49\x1e\xd4\x8c\x39\x0b\x61\xe6\x2b\x79\x86\x46\x54\xb5\x36\x7c\x8e\x27\x39\xd0\xbc\x0f\xc6\xc7\x6b\x40\xb7\xb7\xa4\x7d\x0c\xf5\xa3\x8a\x37\x65\xaf\x67\x3a\x19\xb5\xe9\x3f\x03\x3a\xf6\xbf\x8b\x8d\x16\x33\x9b\x19\x6b\xd1\x0a\x71\x53\xbe\x64\xef\xf7\xf0\x4a\x1a\xb4\xa5\x93\xb2\x10\x90\xd9\x38\x1d\x80\x26\xb0\xa4\x35\x7a\x51\x15\xf9\x69\x4d\x93\xaf\x1b\x4b\xc9\x4c\x20\x3f\x05\x97\xf3\x18\x1c\x74\x50\xa7\xb4\x6e\xeb\x94\xd0\x8d\x39\x97\x2b\xa6\xe8\x06\x30\xce\x70\xd3\x42\x12\x6b\x52\x07\x5b\xd4\x01\x7f\x35\xad\xb8\x17\xd0\xc6\x69\x98\xc8\xcf\x8a\xc1\x04\xf8\xd2\x1f\x0f\x8b\x5f\xbe\x2c\x8c\xd1\x43\x07\x6f\xb6\xa3\x45\xfd\xcb\x53\x5e\x93\xe5\x31\x32\x74\x40\x41\xf6\x3c\x52\x64\xf1\xd0\xd2\xe2\x0d\x6d\x80\xf7\x1d\x7c\xff\x71\x9f\xa5\xa0\x24\xaa\x6a\xdb\x7f\x55\xce\x0c\x0b\xdc\xfd\x83\x2f\x8a\x95\xfd\x1f\x3e\x7b\x4f\xfe\x0e\xb3\xa8\xce\x44\x97\x1d\xa0\xf7\x25\xff\xe1\xdd\xcb\x94\xfd\xb2\xfa\x4c\x74\x69\xc6\xe0\x76\xb0\x56\x69\xd7\x63\xf2\x9b\x99\x91\x96\xe6\xef\xde\xc2\x13\x94\xef\x91\x3d\x0e\xe4\xc6\xcd\xd0\xb3\x1a\x2e\xec\xd5\x80\x1f\xc9\x4d\x46\x47\xaf\xd8\x90\x93\xcf\x4e\x9d\x2c\x42\x07\xec\x23\x0a\x51\x99\x69\x69\xd0\x75\xe0\x8c\x85\x32\x57\x79\x94\x72\x6d\x90\x1f\xc6\xb1\x29\x61\xad\xa8\xf2\xfd\x7f\xb2\xf8\x13\x00\x00\xff\xff\x31\xad\x0c\x65\x74\x02\x00\x00")

func commandsServiceTemplateLogger_goBytes() ([]byte, error) {
	return bindataRead(
		_commandsServiceTemplateLogger_go,
		"commands/service/template/logger_go",
	)
}

func commandsServiceTemplateLogger_go() (*asset, error) {
	bytes, err := commandsServiceTemplateLogger_goBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "commands/service/template/logger_go", size: 628, mode: os.FileMode(436), modTime: time.Unix(1499856489, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _commandsServiceTemplateMain_go = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x94\x53\xc1\x6e\xdb\x38\x10\x3d\x93\x5f\x31\x4b\x60\x17\xd4\xc2\xa1\xb0\x57\x2f\x72\x08\x52\xa4\x39\x04\xa9\xe1\x04\xf0\xa1\xe8\x81\x96\xc6\x0c\x51\x99\x64\x48\x3a\x46\x21\xe8\xdf\x8b\xa1\xa4\xba\x0e\xd2\xa2\x39\x49\x14\x67\xde\xbc\xf7\xe6\x29\xe8\xe6\xab\x36\x08\x7b\x6d\x1d\xe7\x76\x1f\x7c\xcc\x20\x39\x13\x0e\xb3\xe0\x4c\xa4\x6f\xae\xa1\xa7\xf1\xde\x74\xa8\x8c\xef\xb4\x33\xca\x47\x53\x9b\x18\x7e\x73\x53\x47\xdc\x75\xd8\x64\xeb\x9d\xe8\x7b\xbb\x03\xb5\xc1\xed\x7a\x75\x3d\x0c\xd4\x62\xf3\xd3\x61\xab\x1a\xbf\xaf\xed\x3e\x44\xbf\xd5\xdb\x0e\x2f\xd0\x8d\xad\x17\x47\xdc\xd6\xc6\x97\xf7\x23\x6e\x05\x07\x00\x20\x3e\xf5\x53\xce\x41\xf4\x3d\xba\x76\x18\x78\xc5\xf9\xee\xe0\x9a\xc2\x5c\x56\x3d\x67\x9d\x4d\x0b\xc0\x18\x61\x79\x09\x0e\xb3\xba\xb3\x29\xa3\x93\x22\x37\x41\x2c\x40\x88\x8a\x33\xbb\x2b\x05\x7f\x5d\x82\xb3\x1d\xf4\x9c\xb1\xbe\x57\xf7\x7a\x8f\xc3\xa0\xee\xbc\x51\x37\x3a\xeb\x6e\x27\xc5\x4e\xdb\x0e\x5b\xc8\x1e\xba\x02\xb2\x84\xbf\x5f\x44\x01\xaf\x38\x1b\x38\x67\x89\x86\x10\x41\x75\x8f\xc7\x07\x8c\x2f\x18\x65\xc5\x39\xab\x6b\x58\xa3\xa1\x9e\x08\x27\x07\x20\x61\x7c\xb1\x0d\x82\x77\x60\xd6\xab\xeb\x72\xc6\xa8\x38\x3b\xd5\xa8\xb9\x4f\x26\x02\x3a\xf7\xcc\x07\x2a\x29\x43\x3f\x7f\x99\x7c\x51\x9f\xca\xc7\x7e\x18\xc7\x6e\xa2\x0e\x85\xd1\x04\x4e\xe4\x67\x37\xa7\x4f\x9c\x1d\xa3\x0e\x01\xdb\x91\xf0\x2c\x81\xb0\xa8\x7b\x92\x91\x16\x30\x8d\x53\x4a\x11\x95\x27\xed\xda\x6e\x2c\x27\xc3\x65\xc4\x14\x80\x56\xa1\xd6\x98\x82\x77\x09\x37\xd1\x66\x8c\x0b\x88\xf8\x0c\xff\x4e\x37\xcf\x07\x4c\xb9\x2a\x1e\x9f\x4d\x55\xe5\x71\xfb\xf8\xb8\x2a\x40\xa5\x69\x32\x95\x3a\x4f\xd4\x0a\xce\x78\x24\x90\xab\xb6\x8d\xcb\x92\x04\xb1\xe0\x8c\xdd\x8e\xa4\x96\x63\xd9\x74\xba\x21\x7a\x13\xdd\x6a\x41\xa0\x6c\x4e\x0b\x3b\x1a\x02\xfd\x87\xf2\xac\x36\xda\xe6\x8f\xd1\x1f\x42\x5f\x2e\xd4\x55\xdb\xca\xff\x48\xaa\xf1\xa3\xc4\xea\x8d\x6c\xac\xa2\x75\xb9\x73\x52\x3c\x64\x1d\xb3\x75\xe6\xe7\x55\x82\x77\x73\xc6\xe6\x90\x2d\x2f\x21\x8d\xfc\x65\x67\x53\xf5\xff\xeb\xe4\xfd\x22\x7a\xb4\x91\x25\x9c\x02\x58\xf0\xcf\xf2\x57\xbc\x22\xda\x1f\xbc\x43\x49\xde\xc9\x8a\xbd\xfe\xc7\xfe\x48\x49\x3a\x53\xb2\xf9\x11\x94\xb7\xd5\x9c\xd6\x33\xfd\x5a\x57\x6e\xdc\xa9\x7c\x97\x38\x0a\xe4\xfb\x05\xce\x7b\x2c\xfb\xa2\xfd\xc9\x8a\x0f\xfc\x7b\x00\x00\x00\xff\xff\xe6\x39\xa0\x24\xbe\x04\x00\x00")

func commandsServiceTemplateMain_goBytes() ([]byte, error) {
	return bindataRead(
		_commandsServiceTemplateMain_go,
		"commands/service/template/main_go",
	)
}

func commandsServiceTemplateMain_go() (*asset, error) {
	bytes, err := commandsServiceTemplateMain_goBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "commands/service/template/main_go", size: 1214, mode: os.FileMode(436), modTime: time.Unix(1499862304, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _commandsServiceTemplateProtoProto = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x2a\xae\xcc\x2b\x49\xac\x50\xb0\x55\x50\x2a\x28\xca\x2f\xc9\x37\x56\xb2\xe6\xe2\x2a\x48\x4c\xce\x4e\x4c\x4f\x55\xa8\xae\xd6\xf3\x4b\xcc\x4d\xad\xad\xb5\xe6\xe2\x2a\x4e\x2d\x2a\xcb\x4c\x46\x12\x53\xa8\xe6\xe2\xaa\x05\x04\x00\x00\xff\xff\x48\x81\x80\x06\x3e\x00\x00\x00")

func commandsServiceTemplateProtoProtoBytes() ([]byte, error) {
	return bindataRead(
		_commandsServiceTemplateProtoProto,
		"commands/service/template/proto.proto",
	)
}

func commandsServiceTemplateProtoProto() (*asset, error) {
	bytes, err := commandsServiceTemplateProtoProtoBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "commands/service/template/proto.proto", size: 62, mode: os.FileMode(436), modTime: time.Unix(1499854874, 0)}
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
	"commands/service/template/logger_go": commandsServiceTemplateLogger_go,
	"commands/service/template/main_go": commandsServiceTemplateMain_go,
	"commands/service/template/proto.proto": commandsServiceTemplateProtoProto,
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
	"commands": &bintree{nil, map[string]*bintree{
		"service": &bintree{nil, map[string]*bintree{
			"template": &bintree{nil, map[string]*bintree{
				"logger_go": &bintree{commandsServiceTemplateLogger_go, map[string]*bintree{}},
				"main_go": &bintree{commandsServiceTemplateMain_go, map[string]*bintree{}},
				"proto.proto": &bintree{commandsServiceTemplateProtoProto, map[string]*bintree{}},
			}},
		}},
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

