// Code generated by go-bindata. DO NOT EDIT. @generated
// sources:
// crd/jenkins_v1alpha2_jenkinsinstance.yaml
package crddata

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

var _jenkins_v1alpha2_jenkinsinstanceYaml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xb4\x58\x4f\x73\xdb\xb6\x12\xbf\xeb\x53\xec\x38\x87\x24\x13\x4b\x7a\x4e\x2e\x6f\x74\xcb\x38\xef\xe0\xd7\x69\x26\x13\x67\xd2\x43\xa7\x07\x08\x58\x4a\x1b\x83\x00\x0a\x2c\x94\xa8\x9d\x7e\xf7\x0e\x00\x92\xfa\x63\x92\xa2\xd3\x54\x17\x9b\xc0\xfe\xe3\x6f\x17\xbb\x3f\x50\x38\xfa\x8c\x3e\x90\x35\x2b\x10\x8e\xf0\x1b\xa3\x49\x4f\x61\xf1\xf0\xdf\xb0\x20\xbb\xdc\xdd\xac\x91\xc5\xcd\xec\x81\x8c\x5a\xc1\x6d\x0c\x6c\xeb\x8f\x18\x6c\xf4\x12\xdf\x61\x45\x86\x98\xac\x99\xd5\xc8\x42\x09\x16\xab\x19\x80\xf4\x28\xd2\xe2\x27\xaa\x31\xb0\xa8\xdd\x0a\x4c\xd4\x7a\x06\xa0\xc5\x1a\x75\x48\x32\x00\xd2\x1a\xf6\x56\x6b\xf4\x73\xb6\x56\xb7\x0e\x57\x70\x75\xb3\xf8\xcf\xd5\x0c\xc0\x88\x1a\x57\xf0\x05\xcd\x03\x99\x40\x26\xb0\x30\x12\xc3\xa2\x59\x68\xff\x5a\x87\x5e\xb0\xf5\x8b\x20\xea\x10\xcd\x66\x2e\x8d\xe4\xc5\x86\x78\x1b\xd7\x0b\x69\xeb\x59\x70\x28\x93\xc7\x8d\xb7\xd1\x75\xf6\x26\xab\x97\x38\x9a\x98\x0b\x08\xff\x2f\xaa\x77\x4d\x48\x79\xc7\xe9\xe8\x85\x7e\x1c\xee\x0c\x20\x48\xeb\x70\x05\xef\x93\x19\x27\x24\xaa\xb4\x16\xd7\xbe\x01\xb1\x31\x1d\x58\x70\x0c\x2b\xf8\xf3\xaf\x19\xc0\x4e\x68\x52\x19\xc3\xb2\x69\x1d\x9a\xb7\x1f\xee\x3e\xbf\xb9\x97\x5b\xac\x45\x59\x04\x70\x3e\x85\xcf\xd4\xda\x48\xbf\xa3\x84\x76\x6b\x00\x0a\x83\xf4\xe4\xb2\x45\x78\x9e\x4c\x15\x19\x50\x29\x85\x18\x80\xb7\x08\xbb\xb2\x86\x0a\x42\x76\x03\xb6\x02\xde\x52\x00\x8f\xce\x63\x40\xc3\x39\xa4\x23\xb3\x90\x44\x84\x01\xbb\xfe\x82\x92\x17\x70\x8f\x3e\x19\x81\xb0\xb5\x51\xab\x94\xe2\x1d\x7a\x06\x8f\xd2\x6e\x0c\xfd\xd1\x59\x0e\xc0\x36\xbb\xd4\x82\x31\xf0\x89\x45\x32\x8c\xde\x08\x9d\x40\x88\x78\x0d\xc2\x28\xa8\xc5\x1e\x3c\x26\x1f\x10\xcd\x91\xb5\x2c\x12\x16\xf0\xb3\xf5\x08\x64\x2a\xbb\x82\x2d\xb3\x0b\xab\xe5\x72\x43\xdc\x96\xb0\xb4\x75\x1d\x0d\xf1\x7e\x99\x6b\x8e\xd6\x91\xad\x0f\x4b\x85\x3b\xd4\x4b\xe1\x68\x9e\xe3\x34\x9c\xcb\xbe\x56\xcf\xba\xcc\x3c\x3f\x0a\x8c\xf7\x29\x89\x81\x3d\x99\x4d\xb7\x9c\xeb\x61\x10\xe6\x9f\xc8\x28\xa0\x00\xa2\x51\x2b\xe1\x1e\xd0\x4c\x4b\x09\x84\x8f\xff\xbb\xff\x04\xad\xd3\x8c\xf8\x29\xc4\x19\xdc\x83\x5a\x38\xe0\x9c\x70\x21\x53\xa1\x2f\x79\xaa\xbc\xad\xb3\x45\x34\xca\x59\x32\x9c\x1f\xa4\x26\x34\xa7\x18\x87\xb8\xae\x89\x53\x62\x7f\x8f\x18\x38\xa5\x63\x01\xb7\xc2\x18\xcb\xb0\x46\x88\x4e\x09\x46\xb5\x80\x3b\x03\xb7\xa2\x46\x7d\x2b\x02\xfe\x68\x94\x13\xa0\x61\x9e\x10\xbc\x8c\xf3\x71\x77\x39\x15\x2c\xe0\x74\xcb\xed\x61\x6f\x7f\x7d\x27\x24\x9f\x12\x55\x93\x09\x28\x3d\xf2\xe9\xc6\x40\x00\x59\xa7\xca\xdd\x6e\x7f\xae\x70\x92\xf4\xb7\x8d\x10\x04\xe4\x94\xe0\xd0\x6b\xfd\x2c\xea\x6c\x3d\x81\x9f\xcf\x57\x18\x75\xd0\x34\x1f\x50\xe8\xb4\xdd\xd7\x68\xf8\x58\x75\xaa\x37\x29\x82\x94\xd6\x54\xb4\x19\x75\x56\x44\xa2\xcf\xc6\xe7\x22\xcc\xa5\x55\x98\x41\x3e\xd3\x1a\xc2\x19\x4a\xa3\xaf\x68\x53\x0b\xf7\x78\xeb\xcc\x9d\xf5\x39\xb0\x46\xe3\x45\x78\x09\xb5\x8d\x86\x51\x95\xca\x16\xb9\x15\xab\x66\x1b\x6a\xe1\x7a\x0c\x0e\xa6\xef\x10\x4a\xd9\xbc\x18\xcd\x51\x28\x20\x02\xd4\x51\x33\xcd\x35\x19\x84\xbd\xa8\xf5\x90\x8b\xd1\x08\x46\xf3\xd1\x5f\x8e\x53\xf2\x91\x15\x33\x38\x53\x6b\x59\x99\xe0\xac\x26\x39\x5e\xcc\xca\x04\x28\x62\x53\xed\xa2\xd9\x8d\x5a\x7c\x47\x32\xfd\x23\xfc\x3e\x4d\x0e\x34\x3b\xf2\xd6\xe4\x22\xde\x09\x4f\x62\xad\xb1\xe9\xe9\x53\x81\xdb\x78\x6b\x77\xfb\x09\xd0\x15\xc1\x53\x04\xbf\x07\x39\xaa\xc5\x06\x47\x3d\xfd\xb2\x15\x9c\xd9\x8d\x20\x83\xbe\x28\xa4\x79\x17\x03\x42\x65\x7d\x2a\x63\xfc\xda\x12\x05\xa0\x63\x16\x71\xf8\x39\xc1\x69\x0c\xae\x60\xf1\x6a\xb5\x78\xf5\xa4\xe0\x5c\xd4\x7a\x42\x72\x4b\x5c\x49\xf8\x89\x39\xee\x9c\x14\xf4\xc6\xdb\xd5\x91\x97\x46\xfc\x4c\x9a\x18\xeb\x9e\xa6\x31\x98\xf0\x76\x4b\x78\x2f\x4e\x43\x36\x56\x61\xe6\x8c\x63\xf1\xa4\xde\x45\x15\xc9\x2c\xfd\xa4\xc4\x27\x85\x80\x1a\x25\x5b\x3f\xea\xc2\x94\x33\x59\x24\xa7\xd6\xb1\xd3\x71\x43\x17\x5a\xff\xdb\xf4\xca\xe9\xdc\x34\xc2\xa9\xa8\xd6\x58\x2a\x48\xeb\xcc\x2b\x27\x20\x3b\xd6\xaa\x01\x48\xf5\xad\x9e\x05\x52\xfc\x03\x9d\x7b\xbc\x80\x60\xf9\xed\x1e\xb3\xd3\x71\x37\x8d\x42\x63\xf1\x1a\x2a\xab\xb5\xfd\x5a\x28\x6b\x65\x7d\x2d\x18\x04\x1f\x33\x92\x86\xba\x2f\x9b\x23\x26\x69\xa9\xac\x7c\x40\xff\xac\x18\x9c\x37\x06\xe7\x45\xf9\xe9\xef\x90\x58\x13\x79\xec\x41\x6a\xde\x87\xc9\xbc\x7d\x83\x1f\x51\xe4\x67\xb7\x86\x01\xe4\x0e\xdd\xa7\xa3\x95\x2d\xd5\x9b\x5a\x90\x01\xfd\x8e\xe4\xf8\x61\x6a\xb9\x48\x23\x0b\xd6\xf5\x91\x90\xb1\x82\x1b\x61\x3d\x17\xbd\x0d\xd3\x9e\x0b\xd8\x36\x57\xcb\xa7\xfa\xeb\x69\x16\x70\xa9\x54\x52\x2f\x70\xd6\x3f\x1a\x4d\x8f\x9c\xdd\x55\xd9\x50\xba\x2e\xe4\xfe\x91\x94\xae\xf3\xc8\xc8\xcc\xbe\x5b\x2b\xe3\xb1\xc7\x5a\x29\xe6\x55\xba\x3e\xbd\x79\x3d\x18\x66\xba\x5c\x6d\xf0\xbc\x2b\x75\xd9\xce\x42\x53\x71\x69\x27\x57\x07\x50\xd2\xfe\x41\x7c\xa8\x31\x29\xa4\x4c\x04\x70\xb4\x08\xef\xdb\x72\x28\xb2\x39\x4d\x79\xce\xb6\x13\x96\x2d\xf8\x68\x20\x1a\xf5\xe8\xc5\x07\x43\x0b\x6c\xfd\xa5\x39\xdf\xd5\x47\x91\xfd\x8e\xea\xff\x62\xd7\xc1\xed\xe4\x45\xc4\xdf\xa7\x57\x4a\x7d\xdf\xe3\x1c\xbf\x51\xc8\x37\xc7\x17\xd6\x83\xb1\xfc\x12\x3e\x7c\xbe\x2d\xef\x6b\xd7\xc3\x07\x61\xa0\x3e\x9b\x08\xce\xaf\x4d\xbd\x51\xdc\x55\xd9\x15\xb5\x73\x27\x7f\xe6\x41\x75\x0d\x5f\x13\xe1\xa1\x00\xe9\x46\xd9\x73\x37\x80\x29\x9d\xae\x67\x8b\xad\x46\x3f\xe1\x46\x74\x24\xf7\x2f\x71\x8b\xfe\xab\x66\xf9\x60\xf3\xcf\x2e\x9b\xa7\xdc\x04\x39\xba\x86\x26\x4d\x2d\x54\xb7\x15\xe1\x02\xe3\x61\xc1\x08\x54\x75\xe7\x21\xe4\x8f\x07\x43\xc4\x73\xc0\x53\xdf\xcc\x9b\x17\xef\xc3\x40\xb5\xa3\x1e\x76\x37\x42\xbb\xad\x78\x3d\x3b\x80\x26\xa4\x44\xc7\xa8\xde\x9f\x7f\x5d\xbb\xba\xca\x0f\xed\x07\xb5\xfc\x28\xad\x51\x54\x2a\x01\x7e\xfd\x6d\x56\xce\x27\xaa\xe6\x23\x56\x59\xfc\x3b\x00\x00\xff\xff\xe7\x60\xdb\xbc\xca\x14\x00\x00")

func jenkins_v1alpha2_jenkinsinstanceYamlBytes() ([]byte, error) {
	return bindataRead(
		_jenkins_v1alpha2_jenkinsinstanceYaml,
		"jenkins_v1alpha2_jenkinsinstance.yaml",
	)
}

func jenkins_v1alpha2_jenkinsinstanceYaml() (*asset, error) {
	bytes, err := jenkins_v1alpha2_jenkinsinstanceYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "jenkins_v1alpha2_jenkinsinstance.yaml", size: 0, mode: os.FileMode(511), modTime: time.Unix(0, 0)}
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
	"jenkins_v1alpha2_jenkinsinstance.yaml": jenkins_v1alpha2_jenkinsinstanceYaml,
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
	"jenkins_v1alpha2_jenkinsinstance.yaml": &bintree{jenkins_v1alpha2_jenkinsinstanceYaml, map[string]*bintree{}},
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
