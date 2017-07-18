// Code generated by go-bindata.
// sources:
// commands/service/template/cmd/{{.Name}}/main_go
// commands/service/template/protobuf/proto.proto
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

var _commandsServiceTemplateCmdNameMain_go = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x94\x55\x5d\xaf\xdb\x36\x0f\xbe\xb6\x7e\x05\x5f\x03\x6f\x27\x17\xa9\xb2\xed\x32\x43\x2e\x0e\xce\xc7\xce\x86\x20\x27\x48\x0e\x10\x0c\x45\x51\x28\x36\xe3\x08\x71\x24\x57\x92\x13\x14\x86\xff\xfb\x40\xc9\xce\xc7\x69\x37\xac\x57\x8e\x28\xf2\xe1\x43\xf2\xa1\x52\xcb\x7c\x2f\x4b\x84\x83\x54\x9a\x31\x75\xa8\x8d\xf5\xc0\x59\x92\x6a\xf4\x29\x4b\x52\xf7\x55\xe7\xf4\x2d\x8d\x29\x2b\x14\xa5\xa9\xa4\x2e\x85\xb1\xe5\xb8\xb4\xf5\xbf\xdc\x8c\x2d\x6e\x2b\xcc\xbd\x32\x3a\x6d\x5b\xb5\x05\xb1\xc6\xcd\x72\x71\xdf\x75\x14\xa2\xfc\xae\xd9\x88\xdc\x1c\xc6\xea\x50\x5b\xb3\x91\x9b\x0a\x3f\xa0\x8e\xa1\x1f\x4e\xb8\x19\x97\x26\xfc\x3e\xe1\x26\x65\x00\x00\xc4\x67\xbc\xf3\xbe\x4e\xdb\x16\x75\xd1\x75\x2c\x9a\xaf\xa0\x9c\xb2\x4d\xed\x50\x8f\x2b\x53\xda\xc6\xa5\xdf\x38\xe0\x51\x56\xf5\xce\x6c\x94\xec\x5d\x3e\x3b\xd4\xde\x7e\x4d\x59\xc6\xd8\xb6\xd1\x39\x28\xad\x3c\xcf\xa0\x65\xc9\xed\x91\x90\x92\xca\x94\x62\x85\xfe\xc9\xd8\x83\xf4\x1e\x2d\x7f\x47\x96\x3f\x57\x2f\xf3\xb3\xa9\xed\x32\x76\xe3\xfc\xd2\xf8\xba\xf1\xdc\x38\xb1\xf2\x85\x69\xfc\x9b\xeb\x19\x1e\xb1\xe2\x74\x78\xc0\x4d\x53\x86\xe3\xe0\x32\x1e\xc3\x1f\x5a\x79\x98\x05\xaa\x70\x52\x7e\x07\xab\xc0\x37\xde\x47\xee\x0f\x4e\x8f\xc0\xec\x61\x32\x05\xe3\xc4\xcc\x98\x7d\x53\x3f\xea\x23\x4f\x57\x8f\xf3\xd7\xe5\x5f\x9f\x1f\x56\xf3\x34\x8b\x01\x6a\x4b\x8e\x7d\x31\xc9\xce\x98\xfd\x08\xd0\x5a\x0a\xbd\x69\x87\x98\xe3\x29\x26\x7a\x36\x66\xcf\xaf\xf2\x7c\xfc\x44\x54\x03\xcb\x01\x26\x54\xb2\x90\x5a\xe5\xc1\x3c\xba\x36\x3f\x49\x2f\xab\x6f\xcd\x8f\xd6\x1a\x7b\x63\x3e\x77\x8d\x38\x12\xa5\xe9\x14\xb4\xaa\xce\x5c\x03\x59\xf1\xaa\x0e\x68\x1a\x0f\x53\xf8\xf5\x67\x78\x0f\x5e\x1d\x50\xac\x30\x37\xba\xb8\x71\x5b\x79\x99\xef\xbd\x95\x39\xde\x1b\xbd\x55\x65\x63\x25\xa9\x50\x3c\x6a\x52\x19\x4c\xc1\xdb\x06\xaf\xf9\xdc\x15\x45\x28\x94\xa2\xb3\x81\x51\xfc\x76\x80\x95\xc3\x33\x0f\xf2\x5e\x4b\xab\x2b\xcd\xd3\xb9\x81\x4b\x8b\xe1\x24\x1d\x6c\x4d\xa3\x8b\x11\xc4\x7e\x81\x45\x5a\x25\xa5\x4b\x38\x19\xfd\x93\x87\x93\xb1\xfb\x61\x12\x11\xbd\x63\x1d\xeb\x75\x47\xdb\xc7\xb3\x96\x25\x95\x72\xe7\xa1\x68\xf4\x62\xa6\x9c\x47\xcd\x53\x9f\xd7\xe9\x08\xd2\x34\x63\x43\x8b\xfe\x37\xb4\x28\x49\xda\x56\xcc\xe5\x01\xbb\x4e\xcc\x86\xae\x6f\x79\xba\x95\xaa\xc2\x02\xbc\x81\x2a\x80\x4c\xe0\xff\xc7\x34\x80\x67\x2c\xe9\x18\x4b\x1c\x25\xa1\x25\x8b\x03\xb7\x47\xb4\x3c\x63\x8c\x84\xb7\xc4\x92\x62\x2c\x5c\xb6\x18\x1c\xda\xa3\xca\x11\x8c\x86\x72\xb9\xb8\x0f\x67\xb4\x82\x25\x17\x1f\x31\xc4\x71\x47\x40\xb7\x7b\x6f\x6a\x72\x09\x49\x3f\x7e\xea\x77\x5b\xbc\x04\x63\xdb\xc5\xb4\x6b\x2b\xeb\xc0\xa8\x07\x27\xf2\xc3\x8b\xd0\x9b\x58\x72\xb2\xb2\xae\xb1\x88\x84\x87\x12\x08\x8b\xa2\xfb\x32\xdc\x08\xfa\x74\x42\x08\xa2\xb2\x93\xba\xa8\xa2\x3b\x35\x9c\x5b\x74\x35\xd0\x73\x22\x96\xe8\x6a\xa3\x1d\xae\xad\xf2\x68\x47\x60\xf1\x0b\xbc\xef\x6f\xbe\x34\xe8\x7c\x78\x0e\x6e\xb3\x8a\xf0\x79\x7e\x7d\x5d\x04\xa0\x10\xd4\x37\x95\x22\x2f\xd4\x02\x4e\x3c\x12\xc8\x5d\x51\xd8\x49\x78\x95\xd2\x11\x4b\x92\xe7\x48\x6a\x12\xdd\xfa\xd3\x13\xd1\xeb\xe9\x66\x23\x02\x4d\x86\x17\x2f\x39\x95\x04\xfa\x8e\xde\x64\xb1\x96\xca\xff\x6e\x4d\x53\xb7\xe1\x82\x54\xcc\x7f\xa1\x52\x4b\x13\x4b\xcc\xbe\xa3\x8d\x85\x55\xda\x93\x7a\x57\x5e\x46\x71\x5e\x8d\x12\x8c\x1e\x34\x36\x88\x6c\x32\x05\x17\xf9\xf3\x4a\xb9\xec\xb7\xb7\xca\xfb\x07\xe9\xd1\x44\x26\x70\x11\x60\xc0\xbf\xd1\x5f\xe8\x15\xd1\x7e\x30\x1a\x39\xf5\x8e\x67\xc9\xdb\xff\x89\xff\x54\x89\xbb\xa9\x64\x7d\x16\xca\xf7\xab\xb9\x8c\xa7\x5f\xad\x3b\x1d\x67\xca\x7f\xa8\x38\x12\xe4\x8f\x17\x78\xfe\xe7\x22\x3b\xcd\x8f\x67\xac\x63\x7f\x07\x00\x00\xff\xff\x3b\x4f\x63\x28\x82\x07\x00\x00")

func commandsServiceTemplateCmdNameMain_goBytes() ([]byte, error) {
	return bindataRead(
		_commandsServiceTemplateCmdNameMain_go,
		"commands/service/template/cmd/{{.Name}}/main_go",
	)
}

func commandsServiceTemplateCmdNameMain_go() (*asset, error) {
	bytes, err := commandsServiceTemplateCmdNameMain_goBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "commands/service/template/cmd/{{.Name}}/main_go", size: 1922, mode: os.FileMode(420), modTime: time.Unix(1500309498, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _commandsServiceTemplateProtobufProtoProto = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x2a\xae\xcc\x2b\x49\xac\x50\xb0\x55\x50\x2a\x28\xca\x2f\xc9\x37\x56\xb2\xe6\xe2\x2a\x48\x4c\xce\x4e\x4c\x4f\x55\xa8\xae\xd6\xf3\x4b\xcc\x4d\xad\xad\xb5\xe6\xe2\x2a\x4e\x2d\x2a\xcb\x4c\x46\x12\x53\xa8\xe6\xe2\xaa\x05\x04\x00\x00\xff\xff\x48\x81\x80\x06\x3e\x00\x00\x00")

func commandsServiceTemplateProtobufProtoProtoBytes() ([]byte, error) {
	return bindataRead(
		_commandsServiceTemplateProtobufProtoProto,
		"commands/service/template/protobuf/proto.proto",
	)
}

func commandsServiceTemplateProtobufProtoProto() (*asset, error) {
	bytes, err := commandsServiceTemplateProtobufProtoProtoBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "commands/service/template/protobuf/proto.proto", size: 62, mode: os.FileMode(420), modTime: time.Unix(1500286209, 0)}
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
	"commands/service/template/cmd/{{.Name}}/main_go": commandsServiceTemplateCmdNameMain_go,
	"commands/service/template/protobuf/proto.proto": commandsServiceTemplateProtobufProtoProto,
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
				"cmd": &bintree{nil, map[string]*bintree{
					"{{.Name}}": &bintree{nil, map[string]*bintree{
						"main_go": &bintree{commandsServiceTemplateCmdNameMain_go, map[string]*bintree{}},
					}},
				}},
				"protobuf": &bintree{nil, map[string]*bintree{
					"proto.proto": &bintree{commandsServiceTemplateProtobufProtoProto, map[string]*bintree{}},
				}},
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
