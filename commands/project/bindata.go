// Code generated by go-bindata.
// sources:
// commands/project/template/.github/CONTRIBUTING.md
// commands/project/template/.github/ISSUE_TEMPLATE.md
// commands/project/template/.github/PULL_REQUEST_TEMPLATE.md
// commands/project/template/.gitignore
// commands/project/template/LICENSE
// commands/project/template/README.md
// commands/project/template/mmo.yaml
// commands/project/template/wercker.yml
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

var _GithubContributingMd = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x84\x57\x4d\x8f\x1c\xb9\x0d\xbd\xeb\x57\x10\xf0\x61\x93\x41\xb9\x07\x4e\x7c\xf2\xcd\xb1\x67\x03\x03\xd9\xec\x60\xc6\x9b\x20\x18\xcc\x41\x2d\xb1\xba\x18\xab\xc4\x0a\xa5\xea\x76\xef\xaf\x0f\x48\xa9\x3f\xc6\xbb\xc6\x9e\xa6\xa7\x4a\xc5\x8f\xc7\xc7\x47\xea\x15\x7c\xe0\x5c\x85\xb6\x6b\xa5\xbc\x73\xee\xdf\x13\x66\x08\x57\x8f\xa0\x32\xd4\x89\x0a\x08\x2e\x5c\xa8\xb2\x1c\x07\x58\x12\xfa\x82\x30\x92\x94\x0a\x91\x4a\x58\x4b\x81\x3a\x21\x84\xc9\xe7\x1d\xc2\x91\x57\x38\x50\x99\xf4\xe3\xd9\x7f\x41\xd8\x93\x07\x2a\x65\xc5\xc1\xe1\xec\x29\x0d\xc0\x02\x3e\x1f\x81\xeb\x84\x02\x33\xd6\x89\x23\x1c\xa8\x4e\x66\x86\x0f\x19\xa5\x00\x8f\xdf\xba\x86\x2d\x8e\x2c\xa8\x46\x35\x36\xdf\x1d\x6e\xc0\xb9\xfb\x16\x53\xe6\x8a\x70\x40\x98\xfc\x1e\xf5\x3d\x47\x54\x3b\x81\x73\x5c\x43\xbd\x84\xce\x29\xf1\x01\xa8\x02\x65\xf0\x29\x69\xc8\x02\x94\x2b\x8a\x0f\x95\x38\x97\x4b\x34\x8b\xf0\x7f\x31\xd4\x8d\x73\xaf\x5e\xc1\xfd\x9a\x12\x3c\xe0\xff\x56\x2c\x15\xee\x85\x03\x96\xe2\xdc\x9b\x0d\xdc\xe5\xb2\x0a\x5a\x52\x94\x4b\x55\x9b\x2c\xb0\x5d\x29\x45\x88\xb8\x60\x8e\x98\x03\x61\x01\x2f\x08\x82\x33\xef\x31\x9e\xd2\x51\x2f\x98\x63\x4b\x18\x21\xf9\x23\x0a\x1c\xb4\x14\x91\x5b\x9e\x0e\xa0\xd9\xda\xb8\xbf\x6c\xe0\x97\x25\xfa\xda\x3e\x7b\xb8\x7b\xff\xf1\xa7\xbb\xcd\xdc\xd1\x8b\x58\x3d\x25\x83\xae\x41\x53\x5a\x01\xb1\xe5\x36\xfa\x80\x43\x03\x95\x72\x48\x6b\xc4\x02\x19\x0f\x80\x79\x4f\xc2\x79\xc6\x5c\xcd\xd5\xde\x0b\xf9\x6d\xc2\x32\x00\x7e\x5d\xb8\x60\x84\x85\xa5\x96\x01\xd6\x82\xe3\x9a\x60\xa4\x84\x90\x38\xf8\x86\x95\xcf\xd1\x58\xe3\x29\xa3\xc0\xe2\xc5\xcf\x58\x51\xca\xc6\xfd\x75\x03\x9f\x72\x10\x03\x5d\xc3\xd8\xa3\x14\xe2\x0c\x79\x9d\xb7\x5a\x63\x85\x3f\x1f\x01\xbf\xfa\x79\x49\x58\xcc\x70\xb3\xf7\x32\xbb\x9e\x85\x06\x7b\x32\x51\x27\x5f\x2d\x17\x8d\xf8\x45\x5d\x0e\xbc\xa6\xa8\xb4\x11\x2c\x98\xeb\x06\x3e\x5f\x3c\x2b\x9e\x25\x4c\x38\x1b\x4d\xd6\x82\x40\x05\x9e\x1e\x71\xfe\x17\xca\xf3\x9f\xa6\x5a\x97\x77\xb7\xb7\x05\xe7\x3d\xca\x86\x65\x77\xfb\xe7\x8d\x7b\xbb\x81\xff\xf0\x0a\xb3\x3f\xc2\x8c\xb2\x6b\x89\xbc\x70\x48\x19\x38\x87\x46\x7c\x63\x9e\x9e\x28\xb4\xcb\xaf\x79\x1c\xad\xac\x07\xee\x5c\x8f\xb8\xc7\xc4\x0b\x4a\xb1\x26\xa0\xd1\x3e\xd2\x14\x22\x2b\x7b\xdb\xf7\x0b\xca\x4c\xa5\xa5\xc9\xfa\x46\x93\x1d\xec\xa8\x86\x21\xdd\xaf\xb9\x41\x25\x37\x08\xee\x09\x0f\x28\xd6\x73\x16\x25\x55\x18\x59\xf4\x9b\x46\xdd\x0f\xbd\x17\x3e\xb4\x5e\xd0\x67\xaf\xe0\xe7\x55\xe0\x3e\x61\xdc\xa1\x73\x9f\xf2\x85\x29\x6a\x9d\x47\x18\xb9\x54\x14\xe3\x60\x06\x5e\x30\x5b\x69\x0e\x98\x02\xcf\xfa\xf4\x8a\x38\x83\xe2\xe9\x8b\x3b\xab\x07\x4b\x2b\xe4\xec\xa9\x33\xa3\x68\xfb\x45\x05\x90\x4f\x2d\xbc\x78\xa9\x14\x68\x31\x22\x19\x8e\xab\x9c\x3a\x4e\xbf\x76\xfa\x7f\xe0\x79\x5e\x33\xd5\x23\x78\x98\xbc\xf8\x52\xd4\xe1\xeb\x51\x10\x95\x9f\x28\x84\x8a\xbe\x66\x8b\x7b\x94\x23\x67\x1c\x40\x70\xe7\x25\x26\x2c\xd6\x0d\x7e\x87\x03\x6c\x39\x1e\x5d\xa1\x5f\x71\x50\xd5\xf2\x5b\x4a\x54\x8f\x03\x60\x9d\x32\x05\xfb\xb9\xd3\x3e\x15\xa0\x88\xb9\x9a\xbf\x1c\xd5\x83\xa0\xd5\x62\x80\xa4\xd5\x53\x7b\x17\xb7\x83\xcb\x16\xbc\x6f\xc6\xb4\xb2\xfa\x0f\xf8\x65\x41\x2f\x5e\x4f\x80\x58\xdf\x09\x26\xda\x99\x19\x16\x28\xf8\x75\xf5\xe9\x85\x27\xc7\x6a\xb1\x9a\xb5\xcd\xa5\x3c\x8f\xd5\xe7\xe8\x25\x16\xe7\xee\x4e\x6d\xc2\x23\x6c\x71\xf2\x7b\x62\x69\x7d\x70\x46\xbd\x75\xbc\x76\x5c\x6d\xd2\x61\xda\x49\x7b\xbc\xae\x95\xeb\xdd\xff\xce\xb9\x1b\xf8\xa5\xe8\xc1\x4b\x4d\x35\x67\x7b\x5f\xf4\xab\xe4\xf3\x6e\xf5\x3b\x74\x37\xf0\x37\xd4\xd7\x82\x65\xc1\x50\x55\x05\x78\x84\x48\xe3\xd8\x08\xa2\xf4\x5b\x98\x72\x2d\x27\xd4\x3a\x40\xc5\xdd\xc0\xdf\x15\x81\x71\x4d\xe9\x08\x3e\x04\x5c\x2c\xb6\xc0\xb9\x54\x59\x83\x45\x17\x84\x94\x08\x65\x76\x37\xf0\x23\x87\xd5\x82\xe2\x0c\x07\xcd\x8e\x0a\x6c\x95\x91\xa3\xa5\x8b\x17\x42\xb8\x1b\x78\x9c\xf8\x60\x54\x9c\x17\x5f\xa7\x23\x54\x3e\x28\x58\xbd\xdd\x2e\xd4\x99\xd1\x14\xe7\x25\x88\x6b\x6e\xe1\xa8\xd8\x5d\x10\xdd\x1e\x2f\xbc\xd4\x7c\xae\xd1\x52\x29\x51\xd5\xe0\xb1\x97\x90\x7e\xc5\x78\x46\xc9\x5a\x7a\xf6\x3b\x94\xc6\x9d\x35\x37\x5c\xf1\x54\x6f\x5f\xab\x16\x9c\x33\xb0\x38\x1f\xf7\xbe\x03\xf4\x59\x38\x25\xca\xbb\x41\xa7\xc7\x9a\x14\x9f\xdb\x88\xc2\x3b\x6f\x73\x4f\xd3\xc0\xac\x02\xac\x56\xcf\x1c\x63\x81\x85\x93\x02\xd7\x4c\xfb\xf0\x45\x8d\xdd\xaf\xdb\x44\xc1\xde\x0a\xed\x75\x54\x5c\xba\xe6\xf4\xba\x4c\x06\xb0\x82\x54\x7e\x38\x9f\xa3\x3c\xb2\xcc\xc6\xc0\x01\xca\x1a\x26\xf0\x45\x39\x34\x1d\x8b\xf9\xd0\x0e\x4b\x18\xaa\x70\xa6\xe0\x00\x7c\x8c\xda\x1b\x83\x4d\x1f\x5e\xab\x96\x3d\x69\x2f\x5d\x29\x98\xbb\x81\x9f\x7b\x2d\x4c\x78\xe0\x30\x51\x98\x20\x74\x91\xf6\x9a\xca\x36\xe9\x6c\x37\x46\x50\x44\x41\x65\xa0\x5f\x16\xe1\x45\xa8\xc5\x05\xde\x81\x0a\xc3\xd8\x5a\xd1\x27\x28\x58\xdb\xd2\x72\x6a\x94\x07\x2c\x8b\x1a\xb0\xbe\x26\x2c\xce\xdd\x77\x21\xb9\x96\xa0\x36\x7f\xfb\xc9\xd4\x54\x23\x24\x2f\x34\x1e\x6d\xdf\x51\x49\x3d\x75\x9c\x29\xc7\x99\x20\xee\x4c\x10\x2d\x82\xda\x51\x92\x87\x8a\x6d\x3c\xe9\xa2\x73\x1d\xb3\x1e\x1a\x3d\x69\xde\x22\xd8\x58\xde\x96\x0b\xa0\xec\x7a\x08\xa6\x85\xe7\xa5\x41\xc9\xf0\x5d\x5a\x6e\x7e\x3f\x9f\xf3\xbc\x11\xda\x4d\xa6\x99\x97\xf4\x4c\xe0\xd4\x43\x5b\x38\x06\xc0\x48\x55\xe5\xc7\x09\x9a\xa1\x0b\xaf\xf4\x17\xb5\x1f\x11\xb5\x9e\x5f\xc8\x4e\x97\xa1\x6d\x6e\x9d\x7a\xa7\xae\x3a\x6d\x88\x9c\x8b\x33\x01\x52\x38\x74\x7c\xf9\x44\xbb\xdc\x11\xd1\x3d\xe3\x9b\xa9\x63\xda\x57\x19\xb6\x3e\x43\xc5\x79\x61\xf1\x42\xe9\xa8\x21\x29\x63\x7c\xc6\x5c\x55\x24\xf2\x11\xae\x06\x89\x15\xa9\xb9\x3e\xa1\x51\x4e\xf3\x1f\x8f\x10\x11\xe7\x97\x84\x19\x5c\x9d\x54\x02\x31\x5b\x4f\xf1\x38\x62\x56\x39\x33\xf7\x93\x97\x79\x5c\x53\x97\xd8\xc7\xc0\x0b\x3a\xf7\xf9\x77\x82\xd5\x72\x26\xdd\xd9\xb6\x5c\x27\xa3\x38\xe5\xf3\x70\x2a\x8b\x0f\x7d\x55\xd1\xa7\xad\xe5\xda\x43\x67\xbb\x9b\xd7\x3a\x47\xda\x53\x34\xa1\x2f\x97\x85\xe4\xc4\xb3\x93\x29\x15\x8d\x5a\x2e\x5a\xb5\x81\x2b\x8d\x72\x2f\x3e\xf3\xd7\x1f\x5d\xc4\xad\x2b\x14\x34\xd9\xd4\x59\x3d\x8e\x14\xc8\xa7\xf3\x71\x7c\xad\x8b\xb7\x3b\x37\xec\xc2\xa5\x36\xd9\xf6\x2f\xce\x17\xb6\x3f\x33\x46\x7d\x11\x02\xaf\xb9\x15\x4d\xa9\xab\xb6\x35\x65\x05\x46\xa5\x1e\xe3\x25\x3a\xdf\x08\x5e\xcd\x5a\x4e\x94\x4d\x0c\x79\x1c\xed\x27\xee\x6d\x11\x7b\xb8\x3a\xad\x1a\x38\x5e\x25\xa4\x8b\xcd\x16\xdd\xb8\x4a\x5f\x94\x46\x52\x22\xd9\x6e\x69\xfd\x49\xba\x31\x1f\xaf\xce\x9f\x9b\xa0\x97\xf2\x4e\xd5\x2b\xa0\xc9\x9c\xfb\x74\xdd\x50\x7e\xbb\xb6\xf2\x37\x21\x6c\xa4\xe8\x94\x3a\x50\xc1\xef\x0c\x82\x1e\x92\x5e\x40\xa4\x36\xef\xb6\xe6\x86\xdf\x94\xb0\xa2\x9f\x35\xf9\xa7\x4f\xff\x7c\xbc\x7b\xf8\x0c\x77\x3f\xbd\xff\xf4\x0f\x78\xff\xf1\xe3\xc3\xdd\xe3\xe3\xf3\x06\xde\xa7\xe4\x02\xcf\x4b\xf2\x36\x21\x0f\x94\x92\x0a\x5e\xdf\xda\x62\xa7\xd1\x1e\x4b\xa5\x9d\xaf\xfd\x81\x9d\x12\xd4\x79\x60\x0a\x08\x17\xd1\x98\x7c\x75\x54\x8c\xf9\x18\x21\xa3\xde\x43\x7c\x1f\x3a\xd7\x22\xd4\xf7\xe6\x40\x12\xd6\xb9\x03\xd2\xf6\xe1\x17\x91\x53\x71\xbc\x4d\xdd\xb5\x6d\x67\x0d\x5b\x4d\x77\x6c\x3b\x8a\xed\x37\xed\x96\xd1\x56\xab\x93\xed\x8e\x8e\x18\xce\xca\xf9\x60\x1f\x6c\xdc\x8f\xe7\x4a\x9e\x2f\x25\xba\x3a\xd0\x48\x01\xf0\x52\x2a\x1b\x61\x76\x35\x6a\x70\x1b\x35\x31\x42\x41\xbd\x48\x54\x4c\xc7\xef\x28\xdf\x61\xe2\xd3\xd2\xdc\xef\x73\x3a\x9e\x9a\x61\x8b\xec\xdb\x76\xa6\x0c\x3b\xe6\xe8\x46\xaf\x59\xa8\x37\xbd\x10\x9d\x75\xe8\x68\x13\xf3\xa4\x42\x9a\x17\x8a\x5e\x6d\xdb\x15\x47\xc1\xae\x3a\xd4\x72\x23\x82\x71\xc7\xf5\xad\xe2\x74\x75\xeb\xa0\xfe\x50\x20\xa1\x8f\x28\x65\xa2\xa5\x93\xf3\x7d\x3d\x4b\xe6\x77\xd4\x86\x0a\xf8\xe8\x17\xcd\x7d\x14\x9e\xcd\xe0\xd3\x87\x2b\x15\xfc\xc0\x7b\xcc\x3e\xd7\xe7\xa7\x89\x67\x5c\xfc\x0e\x9f\x87\xf3\x75\xe8\xcd\xe6\xed\xe0\xfc\xde\x53\x32\x0e\x2b\x17\xfb\x6d\xe6\x4a\x48\x5f\x87\x6e\xc2\xee\x36\xfd\xd3\xdb\x37\xb7\x6f\x9f\x9f\xfa\x3f\xcf\xce\x5d\xac\xbf\x83\x3f\x30\xe1\xce\x9f\xfd\xe1\xd1\x6b\x6f\xb7\xff\x0f\x00\x00\xff\xff\x1b\xa1\x80\x79\x78\x10\x00\x00")

func GithubContributingMdBytes() ([]byte, error) {
	return bindataRead(
		_GithubContributingMd,
		".github/CONTRIBUTING.md",
	)
}

func GithubContributingMd() (*asset, error) {
	bytes, err := GithubContributingMdBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: ".github/CONTRIBUTING.md", size: 4216, mode: os.FileMode(436), modTime: time.Unix(1502117101, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _GithubIssue_templateMd = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x6c\xce\xb1\x0a\x02\x31\x10\x84\xe1\x7e\x9e\x62\x20\x85\xd5\x05\xd4\x07\x10\x4b\x6b\x0b\xeb\x90\x5d\xb8\x40\x30\x21\xbb\x01\x1f\x5f\x4e\xb9\xee\xea\xf9\x7e\x18\x84\x10\x02\x1f\x66\x53\x29\x6a\x79\x94\xee\xa5\xbd\x01\xfc\x97\xa7\x6b\x37\x7a\xe3\xd0\x3e\x9a\xcc\xac\xf4\x55\x59\xb6\x00\x38\x47\x12\x97\x48\x5c\x23\xf7\xe2\xb5\x26\x3f\xd9\x4f\xe9\xa7\x6b\x76\x15\x0e\xb5\x59\xfd\x06\x2c\x07\x2a\x65\x9f\xa9\x1e\x98\xbb\x48\xd9\xce\xa4\x4a\x51\x4f\xa5\x1a\x81\xe5\x1b\x00\x00\xff\xff\x26\xdf\x8a\x88\xb2\x00\x00\x00")

func GithubIssue_templateMdBytes() ([]byte, error) {
	return bindataRead(
		_GithubIssue_templateMd,
		".github/ISSUE_TEMPLATE.md",
	)
}

func GithubIssue_templateMd() (*asset, error) {
	bytes, err := GithubIssue_templateMdBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: ".github/ISSUE_TEMPLATE.md", size: 178, mode: os.FileMode(436), modTime: time.Unix(1502117101, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _GithubPull_request_templateMd = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xd2\x55\x88\x56\x88\x55\xf0\x48\x2c\x4b\x55\xa8\xcc\x2f\x55\x48\xc9\xcf\x4b\x55\xb0\xe7\x22\x41\x14\x10\x00\x00\xff\xff\xf6\x91\xed\x6c\x41\x00\x00\x00")

func GithubPull_request_templateMdBytes() ([]byte, error) {
	return bindataRead(
		_GithubPull_request_templateMd,
		".github/PULL_REQUEST_TEMPLATE.md",
	)
}

func GithubPull_request_templateMd() (*asset, error) {
	bytes, err := GithubPull_request_templateMdBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: ".github/PULL_REQUEST_TEMPLATE.md", size: 65, mode: os.FileMode(436), modTime: time.Unix(1502117101, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _Gitignore = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x1c\x8c\x41\x6e\x03\x21\x0c\x45\xf7\x3e\xc5\x97\xb2\x43\x0d\x73\x87\x5e\xa0\x5d\xf4\x00\x61\xc0\x61\x5c\xb9\x78\x04\x26\x6d\x6e\x5f\x91\xdd\xfb\xfa\x4f\xef\x82\x77\x69\xa9\x0b\x0f\xdc\xad\xe3\xec\x56\x7b\xfa\x19\x48\xad\xe0\xd4\x59\xa5\x0d\x0a\x91\xff\x98\x42\x2c\xaa\x14\xe2\xb0\x85\x4f\x95\x9d\xe8\x82\x2f\x1e\x8e\x7d\x25\x9e\x6f\xd8\xa7\x68\xc1\xaf\xf8\x81\x5b\x35\xf8\xfa\xae\xf9\x46\x21\x2e\x5c\xfa\xc7\xf4\x73\x3a\xec\x0e\x3f\x18\xd5\x90\xed\xc1\x3d\x55\x86\x9b\xad\xba\xcd\x97\xf7\xd9\xed\x9b\xb3\x5f\xd5\x72\x52\x54\x95\xc2\xc8\x29\x1f\x4c\xf1\x35\x36\x7a\x70\x2b\xd6\x37\x8a\x52\x38\x6d\xff\x01\x00\x00\xff\xff\xa8\x8b\x40\xaf\xc7\x00\x00\x00")

func GitignoreBytes() ([]byte, error) {
	return bindataRead(
		_Gitignore,
		".gitignore",
	)
}

func Gitignore() (*asset, error) {
	bytes, err := GitignoreBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: ".gitignore", size: 199, mode: os.FileMode(436), modTime: time.Unix(1502117101, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _license = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x54\x91\x4d\x6f\xeb\x28\x18\x85\xf7\xfc\x8a\xa3\xae\x5a\x09\x75\x3e\x34\xab\xd9\x51\x9b\xc4\xe8\xda\x10\x01\xb9\x99\x2c\x1d\x9b\xd4\x8c\x1c\x88\x0c\x99\xaa\xff\x7e\x84\x9b\xde\xde\xbb\xb2\xcc\xfb\x71\xce\x73\x5e\x3b\x39\x74\xc2\xa2\xf5\x83\x0b\xc9\xe1\xb1\x13\xf6\x89\x90\x2a\x5e\xdf\x17\xff\x3a\x65\x3c\x0e\x4f\xf8\xf3\xf7\x3f\xfe\xc2\x37\x17\x32\xaa\x67\xd4\x71\x1c\x13\x21\x3b\xb7\x5c\x7c\x4a\x3e\x06\xf8\x84\xc9\x2d\xee\xf4\x8e\xd7\xa5\x0f\xd9\x8d\x14\xe7\xc5\x39\xc4\x33\x86\xa9\x5f\x5e\x1d\x45\x8e\xe8\xc3\x3b\xae\x6e\x49\x31\x20\x9e\x72\xef\x83\x0f\xaf\xe8\x31\xc4\xeb\x3b\xe2\x99\xe4\xc9\x27\xa4\x78\xce\x6f\xfd\xe2\xd0\x87\x11\x7d\x4a\x71\xf0\x7d\x76\x23\xc6\x38\xdc\x2e\x2e\xe4\x3e\x17\xbd\xb3\x9f\x5d\xc2\x63\x9e\x1c\x1e\xcc\x7d\xe2\xe1\x69\x15\x19\x5d\x3f\xc3\x07\x52\x6a\x9f\x25\xbc\xf9\x3c\xc5\x5b\xc6\xe2\x52\x5e\xfc\x50\x76\x50\xf8\x30\xcc\xb7\xb1\x78\xf8\x2c\xcf\xfe\xe2\xef\x0a\x65\x7c\xc5\x4f\xc8\x91\xdc\x92\xa3\xab\x4f\x8a\x4b\x1c\xfd\xb9\x7c\xdd\x8a\x75\xbd\x9d\x66\x9f\x26\x8a\xd1\x97\xd5\xa7\x5b\x76\x14\xa9\x3c\xae\x69\xd2\xc2\xf1\x5b\x5c\x90\xdc\x3c\x97\x0d\xde\xa5\x0f\xd6\x2f\x77\x6b\x4f\xb1\x7e\x2d\x81\xe6\x7b\x44\x45\x17\x6f\x53\xbc\xe0\x17\x12\x9f\x70\xbe\x2d\xc1\xa7\xc9\xad\x33\x63\x44\x8a\x94\xa4\xdb\xe9\x5f\x37\xe4\xf2\x52\xda\xcf\x71\x9e\xe3\x5b\x41\x1b\x62\x18\x7d\x21\x4a\x7f\x13\x52\x4e\xdd\x9f\xe2\x7f\x6e\x65\xf9\xb8\x6e\x88\xd9\x0f\x1f\x71\xaf\x07\xb8\x7e\x5d\xf5\x5e\x4a\x53\x3f\xcf\x38\xb9\x7b\x60\x6e\x84\x0f\xe8\xe7\x99\x7c\xe2\x2c\x05\x38\xe5\x3e\x64\xdf\xcf\xb8\xc6\x65\xd5\x2b\xc7\xff\xd9\xfa\x33\x21\xb6\xe1\x30\x6a\x63\x0f\x4c\x73\x08\x83\x9d\x56\xdf\x45\xcd\x6b\x3c\x30\x03\x61\x1e\x28\x0e\xc2\x36\x6a\x6f\x71\x60\x5a\x33\x69\x8f\x50\x1b\x30\x79\xc4\x37\x21\x6b\x0a\xfe\xcf\x4e\x73\x63\xa0\x34\x11\xdd\xae\x15\xbc\xa6\x10\xb2\x6a\xf7\xb5\x90\x5b\xbc\xec\x2d\xa4\xb2\x68\x45\x27\x2c\xaf\x61\x15\x8a\xe0\x7d\x95\xe0\xa6\x2c\xeb\xb8\xae\x1a\x26\x2d\x7b\x11\xad\xb0\x47\x8a\x8d\xb0\x92\x1b\x43\x36\x4a\x83\x61\xc7\xb4\x15\xd5\xbe\x65\x1a\xbb\xbd\xde\x29\xc3\xc1\x64\x0d\xa9\xa4\x90\x1b\x2d\xe4\x96\x77\x5c\xda\x67\x08\x09\xa9\xc0\xbf\x73\x69\x61\x1a\xd6\xb6\xab\x14\xdb\xdb\x46\xe9\xd5\x5f\xa5\x76\x47\x2d\xb6\x8d\x45\xa3\xda\x9a\x6b\x83\x17\x8e\x56\xb0\x97\x96\x63\x95\x92\x47\x54\x2d\x13\x1d\x45\xcd\x3a\xb6\x2d\xee\x34\x94\x6d\xb8\x5e\xdb\xee\xee\x0e\x0d\x2f\x4f\x44\x48\x30\x09\x56\x59\xa1\x64\xc1\xa8\x94\xb4\x9a\x55\x96\xc2\x2a\x6d\x7f\x8c\x1e\x84\xe1\x14\x4c\x0b\x53\x02\xd9\x68\xd5\x51\x94\x38\xd5\xa6\xb4\x08\x49\x2a\x25\x25\xff\xd8\x52\xa2\xc6\x2f\x17\x51\x7a\xfd\xdf\x1b\xfe\xe5\xa5\xe6\xac\x15\x72\x6b\x0a\xf1\xcf\xcd\xcf\xff\x07\x00\x00\xff\xff\x6c\x73\xfa\xfc\x37\x04\x00\x00")

func licenseBytes() ([]byte, error) {
	return bindataRead(
		_license,
		"LICENSE",
	)
}

func license() (*asset, error) {
	bytes, err := licenseBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "LICENSE", size: 1079, mode: os.FileMode(436), modTime: time.Unix(1502117101, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _readmeMd = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x52\x56\xa8\xae\xd6\xf3\x4b\xcc\x4d\xad\xad\xe5\x02\x04\x00\x00\xff\xff\x03\x60\xee\x9e\x0c\x00\x00\x00")

func readmeMdBytes() ([]byte, error) {
	return bindataRead(
		_readmeMd,
		"README.md",
	)
}

func readmeMd() (*asset, error) {
	bytes, err := readmeMdBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "README.md", size: 12, mode: os.FileMode(436), modTime: time.Unix(1499416903, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _mmoYaml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xca\x4b\xcc\x4d\xb5\x52\xa8\xae\xd6\xf3\x4b\xcc\x4d\xad\xad\xe5\xca\x49\xcc\x4b\xb7\x52\x48\xcf\xe7\x4a\x49\x2d\x48\xcd\x4b\x49\xcd\x4b\xae\xf4\x4d\xcc\x4b\x4c\x4f\x2d\xb2\x52\x48\xcf\xc9\x4c\x49\x05\x04\x00\x00\xff\xff\x14\x69\x12\x72\x31\x00\x00\x00")

func mmoYamlBytes() ([]byte, error) {
	return bindataRead(
		_mmoYaml,
		"mmo.yaml",
	)
}

func mmoYaml() (*asset, error) {
	bytes, err := mmoYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "mmo.yaml", size: 49, mode: os.FileMode(436), modTime: time.Unix(1499854890, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _werckerYml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x6c\x93\x41\x6f\xdb\x38\x10\x85\xef\xfe\x15\x73\xc8\x95\xd6\x26\x6d\xdd\x2e\x81\x1c\x1c\x99\xf0\x7a\xdb\x44\x81\x64\x6f\xb1\x27\x81\x22\xc7\x12\xd7\x12\x49\x90\x54\x12\xfd\xfb\x05\x1d\x1a\x51\xda\x1c\xcd\xf7\x3c\xf3\xcd\xcc\x53\x63\x5e\x28\xb4\xa6\xe7\xba\xa5\xd7\xcb\x3f\x17\x8b\x66\x54\xbd\xa4\x0b\x00\x1f\xd0\x7a\x9a\x1e\x88\x47\xf7\xa4\x04\xce\x04\x00\x00\x02\x5c\x4a\x12\x0c\x39\x69\xf3\xac\xeb\xce\xf8\x90\x14\x80\xf8\x43\xf3\x01\x29\xb4\x2a\x74\x63\xb3\x14\x66\x48\xd2\x51\xe9\x16\x9d\x75\x4a\x07\x0a\xd7\x2b\x7a\xf3\x95\x72\x41\xf9\x17\xfa\x75\x45\x6f\xbe\xd1\x1b\x49\x3f\xad\xe8\xea\x13\xbd\x6e\xe8\x97\x15\xfd\x2c\x29\x36\x54\x1e\x29\x5f\xd1\xcf\xdf\x52\x89\x30\x59\xa4\xe0\x3c\x4f\x1c\x1e\xc3\x68\x49\x6b\xc8\xb3\x71\x27\x6f\xb9\xc0\x24\xb4\xbd\x92\x48\x94\xf6\x81\xf7\xfd\xc5\x2c\x9c\xb2\xe1\x02\xfa\x0a\x79\x1e\x13\xd2\x98\x49\x11\x46\x22\x85\x7c\x5b\xd4\xec\x61\x7d\xf7\x83\x6d\x6e\xff\x80\x6d\x51\x54\xb7\xbd\xd2\xe3\x0b\xb4\x26\xfd\x8b\x70\xb8\x74\xf0\xe3\xf1\xa8\x5e\x40\xb4\x06\x88\x81\xab\x9f\xac\xcc\xbf\xb3\xb2\x2e\x0e\xfb\xc7\xc3\xbe\xde\xec\xca\x6c\xe0\x4a\xc3\x32\x13\x83\xcc\xae\x2a\x56\xfe\xb3\xcb\xd9\xa5\xdd\xb3\xa4\xf0\xee\xed\x63\x54\x61\xec\x04\xc1\x40\xe8\x10\xbc\xe8\x50\x8e\xbd\xd2\x2d\x58\x65\xb1\x57\xfa\x3d\xbc\xb0\x40\xc8\x13\xba\xc6\x78\x04\xe2\x40\xa2\xed\xcd\x34\xa0\x0e\x1f\xc2\xbd\xc9\x1f\x32\x2d\x52\x3b\x24\xaf\xc6\xc8\x15\x13\x74\x36\x2b\x49\x81\xf7\xf6\x42\x20\x06\x49\x21\x6b\x94\xce\x7c\xf7\x6b\x6a\x94\x0e\xe8\x34\xef\x33\x69\xc4\x09\x1d\xb1\xa3\xef\x2e\x33\x8e\x3e\x4a\x71\xce\xfa\x3f\x6f\x74\x7d\xc2\x29\x29\x96\x7b\xff\x6c\x5c\x04\xda\xe6\x3f\x8a\xc3\xa6\xce\x59\xb9\x4f\xa2\x43\x6b\xbc\x0a\xc6\x4d\x14\xae\x36\xc5\x79\xb0\x92\x6d\x77\xd5\xbe\xfc\x37\xbb\x7a\x2c\x8b\xbf\x59\xbe\xaf\x1f\xd6\xf7\x8c\xfc\xb2\xf6\xc0\x5b\xfa\xb6\x8c\xed\x6e\x5f\xdf\x95\xeb\x87\xfc\x2f\xf2\xee\x2d\x2f\xee\xef\x77\x97\x66\xa8\x83\x9b\xac\x39\xe7\x77\x79\xbe\x68\x12\x62\xfa\x94\x6e\x89\x54\x6e\x56\xb3\x2c\x8a\x37\xcc\x56\xf9\x10\x21\xbb\x10\xac\xa7\x59\xd6\x0a\xb7\x54\x26\x7b\xba\x59\xa4\xe5\x34\xdc\x77\x24\xe0\x60\x7b\x1e\x90\xce\xef\x30\xbb\x4e\xf2\x9e\xc6\x06\x45\xe8\xdf\xe7\xe3\xd5\x16\x13\x12\x65\xa7\x31\xa0\x4f\x86\x98\x6f\x8c\x68\xdf\x0f\x77\xac\x7c\x60\x7b\x56\xd5\xf7\xeb\x6a\xcf\xca\xdf\xb6\x3f\xb7\x1c\x2a\x56\xc6\xd5\xfd\x7e\x88\x99\xe9\x71\x5d\x55\x3f\x8b\x72\x93\x4c\x4a\x7b\x14\xa3\x43\xe2\x4f\xca\x92\xd0\xfb\x18\x43\x75\x9c\x28\x04\x37\xbe\x85\x74\x18\xb8\x8e\xc1\xb1\xb6\x9f\x80\x90\xd8\xfb\xfc\xf9\xde\xce\x4b\xc7\xde\xd5\xe3\x3a\x67\x40\x8e\xb3\x2d\x64\x8b\xff\x03\x00\x00\xff\xff\x73\x62\x5b\x27\xbf\x04\x00\x00")

func werckerYmlBytes() ([]byte, error) {
	return bindataRead(
		_werckerYml,
		"wercker.yml",
	)
}

func werckerYml() (*asset, error) {
	bytes, err := werckerYmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "wercker.yml", size: 1215, mode: os.FileMode(436), modTime: time.Unix(1502117101, 0)}
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
	".github/CONTRIBUTING.md": GithubContributingMd,
	".github/ISSUE_TEMPLATE.md": GithubIssue_templateMd,
	".github/PULL_REQUEST_TEMPLATE.md": GithubPull_request_templateMd,
	".gitignore": Gitignore,
	"LICENSE": license,
	"README.md": readmeMd,
	"mmo.yaml": mmoYaml,
	"wercker.yml": werckerYml,
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
	".github": &bintree{nil, map[string]*bintree{
		"CONTRIBUTING.md": &bintree{GithubContributingMd, map[string]*bintree{}},
		"ISSUE_TEMPLATE.md": &bintree{GithubIssue_templateMd, map[string]*bintree{}},
		"PULL_REQUEST_TEMPLATE.md": &bintree{GithubPull_request_templateMd, map[string]*bintree{}},
	}},
	".gitignore": &bintree{Gitignore, map[string]*bintree{}},
	"LICENSE": &bintree{license, map[string]*bintree{}},
	"README.md": &bintree{readmeMd, map[string]*bintree{}},
	"mmo.yaml": &bintree{mmoYaml, map[string]*bintree{}},
	"wercker.yml": &bintree{werckerYml, map[string]*bintree{}},
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

