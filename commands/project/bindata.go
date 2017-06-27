// Code generated by go-bindata.
// sources:
// template/CONTRIBUTING.md
// template/ISSUE_TEMPLATE.md
// template/README.md
// template/gitignore
// template/mmo.json
// template/wercker.yml
// DO NOT EDIT!

package project

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

var _templateContributingMd = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x01\x00\x00\xff\xff\x00\x00\x00\x00\x00\x00\x00\x00")

func templateContributingMdBytes() ([]byte, error) {
	return bindataRead(
		_templateContributingMd,
		"template/CONTRIBUTING.md",
	)
}

func templateContributingMd() (*asset, error) {
	bytes, err := templateContributingMdBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "template/CONTRIBUTING.md", size: 0, mode: os.FileMode(436), modTime: time.Unix(1498498506, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templateIssue_templateMd = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x01\x00\x00\xff\xff\x00\x00\x00\x00\x00\x00\x00\x00")

func templateIssue_templateMdBytes() ([]byte, error) {
	return bindataRead(
		_templateIssue_templateMd,
		"template/ISSUE_TEMPLATE.md",
	)
}

func templateIssue_templateMd() (*asset, error) {
	bytes, err := templateIssue_templateMdBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "template/ISSUE_TEMPLATE.md", size: 0, mode: os.FileMode(436), modTime: time.Unix(1498498506, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templateReadmeMd = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x52\x56\xa8\xae\xd6\xf3\x4b\xcc\x4d\xad\xad\xe5\x02\x04\x00\x00\xff\xff\x03\x60\xee\x9e\x0c\x00\x00\x00")

func templateReadmeMdBytes() ([]byte, error) {
	return bindataRead(
		_templateReadmeMd,
		"template/README.md",
	)
}

func templateReadmeMd() (*asset, error) {
	bytes, err := templateReadmeMdBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "template/README.md", size: 12, mode: os.FileMode(436), modTime: time.Unix(1498498506, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templateGitignore = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x01\x00\x00\xff\xff\x00\x00\x00\x00\x00\x00\x00\x00")

func templateGitignoreBytes() ([]byte, error) {
	return bindataRead(
		_templateGitignore,
		"template/gitignore",
	)
}

func templateGitignore() (*asset, error) {
	bytes, err := templateGitignoreBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "template/gitignore", size: 0, mode: os.FileMode(436), modTime: time.Unix(1498498506, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templateMmoJson = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xaa\xe6\x52\x50\x50\xca\x4b\xcc\x4d\x55\xb2\x52\x50\xaa\xae\xd6\xf3\x4b\xcc\x4d\xad\xad\x55\xd2\x01\x09\xe7\x24\xe6\xa5\x83\x84\xd3\xf3\x21\xfc\x94\xd4\x82\xd4\xbc\x94\xd4\xbc\xe4\x4a\xdf\xc4\xbc\xc4\xf4\xd4\x22\xb0\x64\x4e\x66\x4a\xaa\x12\x57\x2d\x20\x00\x00\xff\xff\x5e\x43\x43\xb9\x49\x00\x00\x00")

func templateMmoJsonBytes() ([]byte, error) {
	return bindataRead(
		_templateMmoJson,
		"template/mmo.json",
	)
}

func templateMmoJson() (*asset, error) {
	bytes, err := templateMmoJsonBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "template/mmo.json", size: 73, mode: os.FileMode(436), modTime: time.Unix(1498498506, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templateWerckerYml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x4a\xca\xaf\xb0\x52\x48\xcf\xcf\x49\xcc\x4b\xb7\x32\xd4\xb3\xe4\xe2\x4a\x2a\xcd\xcc\x49\xb1\x02\x04\x00\x00\xff\xff\x15\x05\x18\x27\x17\x00\x00\x00")

func templateWerckerYmlBytes() ([]byte, error) {
	return bindataRead(
		_templateWerckerYml,
		"template/wercker.yml",
	)
}

func templateWerckerYml() (*asset, error) {
	bytes, err := templateWerckerYmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "template/wercker.yml", size: 23, mode: os.FileMode(436), modTime: time.Unix(1498521298, 0)}
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
	"template/CONTRIBUTING.md": templateContributingMd,
	"template/ISSUE_TEMPLATE.md": templateIssue_templateMd,
	"template/README.md": templateReadmeMd,
	"template/gitignore": templateGitignore,
	"template/mmo.json": templateMmoJson,
	"template/wercker.yml": templateWerckerYml,
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
	"template": &bintree{nil, map[string]*bintree{
		"CONTRIBUTING.md": &bintree{templateContributingMd, map[string]*bintree{}},
		"ISSUE_TEMPLATE.md": &bintree{templateIssue_templateMd, map[string]*bintree{}},
		"README.md": &bintree{templateReadmeMd, map[string]*bintree{}},
		"gitignore": &bintree{templateGitignore, map[string]*bintree{}},
		"mmo.json": &bintree{templateMmoJson, map[string]*bintree{}},
		"wercker.yml": &bintree{templateWerckerYml, map[string]*bintree{}},
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

