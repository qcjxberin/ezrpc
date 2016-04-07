// Code generated by go-bindata.
// sources:
// tmpl/csharp/server.tmpl
// tmpl/golang/server.gogo
// DO NOT EDIT!

package tmpl

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

var _tmplCsharpServerTmpl = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x6c\x90\x3f\x4f\xf3\x30\x10\xc6\xe7\xe4\x53\xdc\xf6\xa6\xd5\xab\x74\x07\x31\x00\xad\xc4\x40\x11\x52\x2b\x21\xb1\x5d\x9d\x4b\x6a\xe4\x9c\xa3\xf3\xb9\x6d\xa8\xf2\xdd\x71\xd2\x22\x31\x74\x48\xee\xdf\xe3\x9f\x9f\xf3\x62\x3e\xcf\x61\x0e\xdb\xbd\x0d\x50\x5b\x47\x90\x22\x46\xf5\x0d\x31\x09\x2a\x55\xb0\xeb\x81\xbe\xa5\x33\xa3\x6c\xe9\xf9\x9f\x82\xd9\x23\x37\x04\x2d\x72\x44\xe7\xfa\x34\x58\xe4\x90\xc7\x60\xb9\x81\x4d\x1f\x94\xda\xfb\x3c\x67\x6c\x29\x74\x68\x08\xce\xe7\xf2\xed\xb7\x18\x06\x38\xe7\x59\x17\x77\xce\x1a\x08\x8a\x9a\xc2\xc1\xdb\x0a\x3c\xaf\x43\x53\xf8\xdd\x17\x19\x85\x40\x5c\x91\xfc\x87\xd4\x7a\x41\xae\x1c\xc9\xea\x40\xac\x8f\xd2\x04\xa0\xd9\x48\xc8\x9e\x3d\x07\xef\xa8\xfc\x10\xab\xf4\x6a\x99\x0a\x2a\xd7\x14\x02\x36\x54\x6e\xe2\x84\x99\x25\x17\x59\x76\x40\x01\x15\xe4\x00\x0f\xc0\x74\x4c\x9b\x8a\xad\xb5\xdc\x8e\xad\xce\x4b\xca\xd6\xd4\x7a\xe9\x9f\x62\x5d\x93\xfc\xa1\x2c\x51\x31\x21\x2e\x04\xdb\x89\xd7\x2b\x61\xf5\xf9\x9e\x0a\x5f\x4c\xd4\x49\xa1\xd2\x4f\xa6\xb2\x70\xb4\x6a\xf6\x70\xc3\xcb\x65\x3e\x7e\x15\xd5\x18\x9d\xde\x8d\xf9\x90\x7e\x03\x18\x9c\x0e\xad\x4e\x86\x3a\xb5\x9e\x81\x4e\x57\xfd\x8d\x2d\x4f\xe5\xd6\x6f\x54\xd2\x5b\x17\xb3\xe9\xf2\xc4\x18\xb2\x7c\xf8\x09\x00\x00\xff\xff\xb6\xbf\x39\xbb\xc9\x01\x00\x00")

func tmplCsharpServerTmplBytes() ([]byte, error) {
	return bindataRead(
		_tmplCsharpServerTmpl,
		"tmpl/csharp/server.tmpl",
	)
}

func tmplCsharpServerTmpl() (*asset, error) {
	bytes, err := tmplCsharpServerTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "tmpl/csharp/server.tmpl", size: 457, mode: os.FileMode(420), modTime: time.Unix(1460010100, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _tmplGolangServerGogo = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xbc\x54\x4f\x6f\xdb\x3e\x0c\x3d\xc7\x9f\x82\x3f\xa3\xf8\xc1\x09\x1c\xf5\x1e\xa0\x97\xb6\x3b\xec\xd0\xae\x6b\x06\xec\x2c\x2b\x74\xa2\xcd\x96\x3c\xfd\x59\x90\x19\xfe\xee\x23\x65\x77\x4b\xd3\x62\x0d\x76\xd8\xc5\x16\x29\xf2\xf1\xf1\x89\x52\xdf\x6f\xb0\xd6\x06\x21\xf7\xe8\xbe\xa3\xcb\x61\x39\x0c\x59\x27\xd5\x57\xb9\x45\xe8\x7b\x71\x2f\x5b\xf4\x64\x23\x90\xbf\xef\x2f\x38\x4c\x2b\x64\x37\xac\xae\x40\xac\x47\x3b\xc5\x51\x84\x6e\x3b\xeb\x02\x14\xd9\x2c\xaf\x0e\x01\x7d\x9e\xd1\x6a\xab\xc3\x2e\x56\x42\xd9\xf6\xd2\xc8\xe0\x97\xda\xa6\x7f\xfe\x7c\xcb\xcb\x36\x62\x73\xb9\xb5\xcb\xb0\x73\xba\x0e\x97\xe3\x2f\xcf\xe6\x59\x56\x47\xa3\xa0\xf0\xb0\xf8\x94\x7c\xf7\x94\xfd\x9c\xcb\x30\xac\x13\xff\x39\x58\x73\xe7\xb7\x45\xeb\xb7\xb0\xe0\x22\x82\xac\x39\xf4\xd9\xcc\x31\xdd\x11\x52\xdc\xe3\xfe\xc6\xb6\xd4\x55\x78\x70\x36\x58\x65\x9b\x47\x94\x1b\x74\x45\xe2\xcc\xdb\x93\x4d\x30\xe2\x56\x06\x39\x27\x0e\x33\xbf\xd7\x41\xed\x80\x7d\xeb\x58\x7d\x41\x15\x08\xb7\xef\x97\xe0\xa4\x21\xb1\x2e\x5a\x0c\x3b\xbb\x79\xa6\xca\x5d\x72\x79\x12\x66\xa6\xa4\x27\x95\x4f\x59\x0b\x72\x8c\x79\x93\x82\xf9\x2a\x41\xea\xfa\x09\x4f\x7c\x30\xb8\x97\x07\x86\x98\x75\x0c\xfe\xff\x29\xc6\x29\xc4\x23\x7e\x8b\xe8\x43\xcf\x19\xe8\x8e\xfb\xbe\x45\x65\x37\xb8\x0e\x2e\xaa\x50\xb8\x12\xba\x39\x85\x50\x29\x8e\xfa\xef\x0a\x8c\x6e\x58\x29\xaa\xe3\xb4\x09\x8d\x29\xc8\xcf\x11\x4f\x40\x57\xe0\xc5\x28\xf3\x0b\xda\xc5\xd9\x50\xdc\x1d\x36\x1e\xff\xb2\x21\x87\xfe\xcc\x1c\xdf\x59\xe3\xf1\x1f\xab\x50\x02\x11\x3c\x17\x8f\x3e\x55\xac\x53\x3b\xe3\xe0\x5d\xc7\xba\x46\x97\x28\xef\xff\x38\xae\x9f\x9d\x0e\x3c\xae\xb1\x66\xa4\x29\xec\x9d\x39\xea\x6b\xff\x8b\x89\x17\x37\xd6\x18\xf1\x10\xab\x46\xfb\x5d\x1a\xe9\x47\xec\x9a\x43\x09\x94\x2e\xae\xb9\x70\x41\xf3\x9d\xce\xc5\x6c\x86\xe1\x68\x45\x1c\x89\x65\x38\x74\x08\x6f\x5d\x3c\xf0\xa9\x2e\x77\x3a\x39\x16\xaf\x07\x66\x33\xa6\x03\x30\x5d\x4f\x36\xb8\x48\xba\xe1\xd4\xe8\xeb\x49\x05\xbd\x2b\xa4\xe2\xc9\x5e\x09\x8a\xa1\x7e\x03\xa5\x8b\xfe\xfa\x7c\x8c\x38\xfd\x7b\xc2\xc1\x16\x4d\x90\x41\x5b\xb3\x02\xc6\xe5\x83\x18\x5f\xbf\x94\xf9\x56\xa7\x7c\x96\xe3\x6a\x05\xbe\x24\x83\x2b\xaf\xa8\x21\x26\x53\xb2\x66\x13\xda\xa8\xfb\xc7\x88\x11\xe9\xbd\xf0\xca\xe9\x0a\x8b\x97\x2f\xc0\x22\x2f\x21\xc7\x1f\xae\x53\xb4\x98\x32\xd3\x23\x36\xcf\x8e\xcf\xe2\x67\x00\x00\x00\xff\xff\x00\xb5\xd6\x66\xaa\x05\x00\x00")

func tmplGolangServerGogoBytes() ([]byte, error) {
	return bindataRead(
		_tmplGolangServerGogo,
		"tmpl/golang/server.gogo",
	)
}

func tmplGolangServerGogo() (*asset, error) {
	bytes, err := tmplGolangServerGogoBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "tmpl/golang/server.gogo", size: 1450, mode: os.FileMode(420), modTime: time.Unix(1459936827, 0)}
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
	"tmpl/csharp/server.tmpl": tmplCsharpServerTmpl,
	"tmpl/golang/server.gogo": tmplGolangServerGogo,
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
	"tmpl": &bintree{nil, map[string]*bintree{
		"csharp": &bintree{nil, map[string]*bintree{
			"server.tmpl": &bintree{tmplCsharpServerTmpl, map[string]*bintree{}},
		}},
		"golang": &bintree{nil, map[string]*bintree{
			"server.gogo": &bintree{tmplGolangServerGogo, map[string]*bintree{}},
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

