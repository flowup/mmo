// Code generated by go-bindata.
// sources:
// commands/service/template/logger.go
// commands/service/template/main.go
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

var _commandsServiceTemplateLoggerGo = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x5c\x91\x41\x8b\xdb\x30\x10\x85\xcf\xd6\xaf\x18\x72\x28\xf6\x12\x94\xd2\x63\xc1\x87\xd2\xee\x52\xca\x92\x2d\xb8\xb7\x52\x8a\x62\x8f\x95\x21\xb2\x26\x8c\xc6\x29\xc5\xf8\xbf\x17\xd9\x4e\xcb\xf6\x66\x7f\xf3\xde\xd3\xd3\xe8\xea\xda\x8b\xf3\x08\xd3\x64\x8f\x6e\xc0\x79\x36\x86\x86\x2b\x8b\x42\x69\x8a\x9d\x27\x3d\x8f\x27\xdb\xf2\x70\x48\x24\xe3\x35\x61\x3c\x04\xf6\x32\xa6\xdd\xeb\x29\xde\x5c\xb8\x9e\xf9\x44\x6e\x9b\xff\x4c\x18\x55\x7e\x67\x19\x2f\x62\xa5\x01\x77\xa6\x32\xe6\xe6\x24\x67\x1f\x0e\xf0\xcc\x1e\x7c\xe0\x93\x0b\x90\x50\x95\xa2\x4f\xc0\x3d\x04\xf6\x1e\xc5\x14\x79\xfc\xb0\xa6\xd9\xe7\x8d\x75\x29\x42\x0d\xbb\x69\xb2\x9f\x52\x9c\xe7\x25\xb0\x1f\x63\x0b\x14\x49\xcb\x0a\x26\xb3\xfa\x6a\xd8\x8c\x47\xfc\x55\x56\xe6\x7e\x9c\xa7\xe8\xa1\x67\x19\x9c\x02\x25\xf8\xd2\xbc\x1c\x17\xbd\x7d\x5a\x98\xa2\x40\x0d\x6f\x36\x6b\x9e\xfe\xe5\xd3\xbc\x26\xdb\x97\x51\xa1\x06\x4e\xb6\xd1\x8e\x47\x35\xa6\x08\x78\xc3\x90\xe0\x7d\x0d\xdf\x7f\xdc\xeb\x66\x34\x99\xa2\xd8\xfe\xbf\xba\x48\xed\x02\xf7\xff\xe0\x93\x53\x17\xfe\x87\x8f\x22\x2c\x77\x38\x9b\xe2\xcc\x7c\xd9\x03\x8a\xe4\xfc\x57\xab\xcd\x57\x6b\x96\xaf\xcf\xcc\x97\xb2\x4b\x71\x0f\x6b\x95\x6a\xb5\xd9\x6f\x34\x20\x2f\x75\xdf\xbd\x85\x07\xc8\x2f\x60\x1b\x6c\x39\x76\x9b\xa0\x51\xd7\x5e\x54\x5c\x8b\x1f\x39\xf6\xe4\x47\x71\x4a\x1c\xed\x63\x74\xa7\x80\x50\x83\xca\x88\xc6\x14\xd4\x2f\x0d\xea\x1a\x22\x05\xc8\xf7\xca\x9b\xc8\xc7\x26\xfb\xa1\xeb\xca\x1c\x56\xe5\xba\xb3\xf9\x13\x00\x00\xff\xff\x40\xd0\xa0\x5a\x50\x02\x00\x00")

func commandsServiceTemplateLoggerGoBytes() ([]byte, error) {
	return bindataRead(
		_commandsServiceTemplateLoggerGo,
		"commands/service/template/logger.go",
	)
}

func commandsServiceTemplateLoggerGo() (*asset, error) {
	bytes, err := commandsServiceTemplateLoggerGoBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "commands/service/template/logger.go", size: 592, mode: os.FileMode(436), modTime: time.Unix(1499788331, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _commandsServiceTemplateMainGo = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x94\x53\x4d\x6b\x1c\x39\x10\x3d\x4b\xbf\xa2\x56\xb0\x8b\x7a\x19\xab\xd9\xeb\x2c\x3e\x98\x04\xdb\x07\xe3\x0c\xb6\x61\x0e\x21\x07\x4d\x77\xb5\xac\x44\x2d\xc9\x92\xc6\x43\x68\xfa\xbf\x07\xa9\xbb\x3d\x19\x63\x43\x7c\xd2\x57\xd5\xab\xf7\xaa\x9e\xbc\x6c\x7e\x48\x85\xd0\x4b\x6d\x29\xd5\xbd\x77\x21\x01\xa7\x84\x59\x4c\x8c\x12\x16\x7f\xda\x26\xaf\xca\x39\x65\x50\x28\x67\xa4\x55\xc2\x05\x55\xab\xe0\xcb\x8b\xc5\x54\x3f\xa6\xe4\x4b\x94\x4e\x8f\xfb\x9d\x68\x5c\x5f\x77\xc6\x1d\xf6\xbe\xee\x7b\x57\x0f\x83\xd8\x04\xf7\x1d\x9b\x74\x2b\x7b\x1c\xc7\x7c\x31\xed\xde\x47\xae\x03\x76\x06\x9b\xa4\x9d\x65\x74\x18\x74\x07\x62\x8b\xbb\xab\xe0\x9b\x71\x3c\xa9\xa3\x7b\x1f\xdc\x4e\xee\x0c\x9e\xa1\x9d\x72\xcf\x0e\xb8\xab\x95\x2b\xfb\x03\xee\xd8\x30\xa0\x6d\xc7\x91\x56\x94\x76\x7b\xdb\x14\xad\xbc\x1a\x28\x31\x3a\xae\x00\x43\x80\xf5\x39\x58\x4c\xe2\x46\xc7\x84\x96\xb3\xd4\x78\xb6\x02\xc6\x2a\x4a\x74\x57\x02\xfe\x3a\x07\xab\x0d\x0c\x94\x90\x17\xf2\xe2\xc6\x29\x71\x29\x93\x34\x1d\x67\x9d\xd4\x06\x5b\x48\x0e\x4c\x01\x59\xc3\xdf\xcf\xac\x80\x57\x94\x8c\x94\x92\x98\x8b\x64\x46\xe2\x16\x0f\xf7\x18\x9e\x31\xf0\x8a\x52\x52\xd7\x70\x87\x2a\xe7\x04\x38\x6a\x86\x88\xe1\x59\x37\x08\xce\x82\xba\xdb\x7c\x2a\x67\x0c\x82\x92\x63\x8c\x58\xf2\x78\xcc\x40\xaf\x9a\xe4\x7c\x8e\x29\x55\xbf\x7e\x9b\x3b\x21\xbe\x94\xcb\x61\x9c\xea\x6e\x83\xf4\x85\xd2\x8c\x9e\xd9\x2f\xfd\x9b\xaf\x28\x39\x04\xe9\x3d\xb6\x13\xe3\x45\x43\xc6\xca\xd9\xb3\x8e\xb8\x82\xb9\x9c\x10\x22\x73\x79\x94\xb6\x35\x53\x78\xee\x38\x0f\x18\x3d\x64\x93\x88\x3b\x8c\xde\xd9\x88\xdb\xa0\x13\x86\x15\x04\x7c\x82\x7f\xe7\x97\xa7\x3d\xc6\x54\x95\x26\x9f\x54\x15\x65\xb9\x7e\x78\xd8\x14\xa0\x92\x34\x77\x35\x67\x1e\xa9\x15\x9c\xe9\x98\x41\x2e\xda\x36\xac\x01\x00\x18\x5b\x51\x42\xae\x27\x52\xeb\x29\x6c\x3e\x5d\x66\x7a\x33\xdd\x6a\x95\x41\xc9\x62\x17\x72\x50\x19\xf4\x9f\xfc\x05\xc4\x56\xea\x74\x15\xdc\xde\x0f\xe5\x41\x5c\xb4\x2d\xff\x2f\x4b\x55\x6e\x92\x58\xbd\x61\x8e\x4d\xd0\x36\x19\xcb\xd9\x7d\x92\x21\x69\xab\x7e\x9f\x25\x38\xbb\x98\x6c\x71\xd9\xfa\x1c\xe2\xc4\x9f\x1b\x1d\xab\xff\x5f\x5b\xef\x1d\xef\xe5\x89\xac\xe1\xe8\xc0\x82\x7f\x62\xc0\xd2\xab\x4c\xfb\xb3\xb3\xc8\x73\xef\xf8\x1b\x96\xf9\x23\x29\xf1\x44\xca\xf6\xc5\x29\x6f\xcb\x39\xce\x67\xfe\x5c\x17\x76\x1a\x2a\xff\x90\xba\xec\xc8\x8f\x2b\x5c\x06\x59\x06\x96\x07\xc8\x2b\x3a\xd2\x5f\x01\x00\x00\xff\xff\x81\x61\x80\xae\xf2\x04\x00\x00")

func commandsServiceTemplateMainGoBytes() ([]byte, error) {
	return bindataRead(
		_commandsServiceTemplateMainGo,
		"commands/service/template/main.go",
	)
}

func commandsServiceTemplateMainGo() (*asset, error) {
	bytes, err := commandsServiceTemplateMainGoBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "commands/service/template/main.go", size: 1266, mode: os.FileMode(436), modTime: time.Unix(1499788331, 0)}
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

	info := bindataFileInfo{name: "commands/service/template/proto.proto", size: 62, mode: os.FileMode(436), modTime: time.Unix(1499788331, 0)}
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
	"commands/service/template/logger.go": commandsServiceTemplateLoggerGo,
	"commands/service/template/main.go": commandsServiceTemplateMainGo,
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
				"logger.go": &bintree{commandsServiceTemplateLoggerGo, map[string]*bintree{}},
				"main.go": &bintree{commandsServiceTemplateMainGo, map[string]*bintree{}},
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

