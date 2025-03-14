// Code generated by go-bindata.
// sources:
// assets/cuda_master_functions.txt
// assets/cudart_functions.txt
// assets/cudnn_functions.txt
// DO NOT EDIT!

package cudnn_log_parser

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

var _cuda_master_functionsTxt = "\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x84\x58\xc1\x72\xa3\x38\x10\xbd\xcf\xd7\xac\x61\x2a\x95\x1c\x1d\x2b\x31\xa9\xb5\x27\x5e\x67\x93\x39\xa6\x34\xd0\x8e\xb5\x83\x05\x2b\x44\xca\x9e\xaf\xdf\x02\x59\x42\xdd\x6a\x76\x6e\xbc\xf7\xba\x9f\x5a\x8d\x84\x25\x97\xfd\xcd\xd7\xa5\x31\xf2\x92\x8b\x95\x01\x69\xe1\x4b\x60\x22\x2c\x72\x71\xb7\x95\xed\x1b\x18\x0b\xe7\xfb\xfe\x70\x00\x33\xf2\xeb\xcd\x56\xb6\x0e\x3f\xff\xf8\x07\x4a\xcb\xb3\xcb\xee\xa2\xcb\x51\xda\xc2\x69\x59\xd7\x0d\x06\x3b\x65\xcb\xa3\x67\x1e\x0d\x80\x7f\x5e\x83\x7d\xd2\x87\xc6\xc3\xa2\xe9\x2c\xca\x2e\xdb\x4b\x26\x10\x40\x23\x8d\xcc\xab\x96\xb5\xfa\xd0\x50\x45\x6c\x2e\x10\xa0\x49\x4b\xdb\xc4\x01\xc2\x36\x4b\x0c\x89\x9a\x18\x08\xdb\x14\x04\xd2\x90\x02\xbb\x14\xd4\xa5\x03\x2b\x16\x37\x18\x25\x01\x19\x09\xc9\xf8\xa0\x3c\xa3\x98\x09\xba\x25\x30\x09\xc1\x2e\x8c\x07\x72\x08\xf9\x74\x71\x21\xfc\xfe\x99\x79\x0a\x07\x24\xb2\x80\xce\x9a\xe6\xf2\xa5\xec\xef\x95\x96\xe6\x72\x5d\x26\xab\xe6\xd4\xaa\x1a\x76\xf6\x3c\x00\x7b\x0e\x2e\xe1\xd9\x79\xac\xec\x79\x72\x98\x40\x10\x5f\x2e\xba\x3c\x9a\x46\xab\x5f\x43\xb2\xc8\xc5\xe2\x8f\xd8\x0d\x13\xcf\x5a\xc0\xa7\x2a\x53\xc1\xd9\x8d\xdc\x56\xb6\x7b\xe8\x9a\xde\x94\xd0\x79\x6e\x0f\x1f\xaa\xb3\x60\xbc\xe0\xf9\x57\x7d\x62\xa2\x5f\xb5\xe1\xe3\x17\xb4\xb2\xc5\x5c\x65\x8b\xb4\xb2\x3b\x92\x7c\x37\x93\x7b\xc7\xa4\xa6\x73\x62\x3e\x0a\x0c\x3b\x39\xf0\x1d\x08\x34\x63\xc5\xf5\xc6\x91\x6c\x30\xdf\xb2\x48\x20\x59\x0f\xeb\xcd\x8b\x35\x20\x4f\xab\x46\x77\xfd\x09\xcc\xb2\xfc\xb7\x57\x06\x1e\x8d\x3c\x01\xa7\xaf\x1a\xad\xdd\x77\x6e\x4e\xfa\xae\xec\xf1\xb1\x96\x1f\x1d\x17\x23\x54\x57\xce\x3b\xec\xa1\x06\xd9\xa5\x83\xef\x4c\x53\xf5\x25\x3b\xb8\x97\x78\x63\xaf\xee\x0c\x74\xa0\xed\xac\xf1\x1e\x6c\x6f\x74\x90\x3f\x41\xdb\xb0\x48\x22\xf4\x68\x9a\xd3\x90\xeb\x76\x36\x11\xbe\xbd\x7d\x8b\x85\x69\xbb\xc5\xd0\x2d\x85\x91\xf9\xab\x07\x13\xe4\x3d\x94\x8d\xa9\x30\x7a\x6f\x6d\xf7\xcb\x53\x78\x83\xae\x37\xf1\x32\x8e\x90\xf3\x5f\x6f\x9e\xb4\xb2\xe3\x43\xfa\xfb\x34\xfb\xeb\xc4\x0b\xde\x71\x46\xf3\x45\x26\x01\x33\x79\x2e\xa5\xea\x46\xcd\xaf\xfb\xa4\xc2\x71\x81\xff\x9e\x9d\x2a\x9f\x16\x38\xcd\x32\xb2\x3d\xaa\xb2\x9b\xfb\x00\xc5\xfa\xe2\xff\x75\x6e\xf7\x7a\xf9\x61\x9a\xcc\xd3\x49\x7e\xc4\x12\x9d\x26\x2b\xd1\x24\xf2\xa9\xe1\xe8\xd0\xf9\xab\x96\x7c\x29\x78\x21\x4d\x63\xbe\x19\x5e\x7c\x13\xbb\xe5\xab\xaf\xf1\xb9\xb7\x6d\x6f\x5f\x7a\x73\x90\xf3\x41\x6f\xaa\x82\x66\x8a\xb9\xae\xc3\x8d\xec\xf5\x78\xc4\x71\x0f\xab\xa6\x69\xc1\x48\xab\x3e\xe1\x4f\x30\x1a\xea\x79\x65\xdb\xd7\x56\x85\x0f\xf3\x4c\x90\x9f\x93\x93\xd7\x46\x55\x08\xf8\x65\xe2\x18\x32\x22\x71\x50\xfa\xe7\xf0\x9b\x5a\xc3\xb8\xb3\x46\xe8\xb7\xd9\x04\xdc\xda\x1e\xf0\xb4\xcb\xa3\x93\x9d\x7f\x1c\xce\x6b\x04\xba\x4c\xcf\x6c\xa5\x96\x1f\xe3\xc9\x8c\x1e\x05\x11\xc6\x49\x01\x5d\x8f\x00\xd7\xa7\x69\xb0\x01\x85\xa0\xe9\x08\x39\x81\x20\xc6\x07\xca\x18\xa2\x00\xff\x66\x53\x06\x85\x4d\xeb\xc8\x71\x3b\x18\x56\x14\x65\x92\x28\x03\x07\xb0\xe5\xd1\xbf\x22\x4a\xf9\xf7\xe2\x4e\x89\xe1\x61\x3c\xf5\xd2\x33\x2f\xc2\xa1\x36\x44\x61\x3b\x7c\x38\x4e\x38\x6c\x11\xd3\xfe\x0b\xe6\x35\x1c\x99\x04\xe4\x22\x7a\xc4\xd5\xe6\x69\xb5\xf9\x4c\xb5\xb9\x18\x1a\x48\x20\x75\x0b\x1c\x97\x4b\xab\xc2\xa3\x26\x65\x63\x6f\xc6\x75\xe9\x2e\x06\x13\x88\x0d\xaf\x98\x58\xba\x53\xff\x04\x48\x06\x53\x84\xbb\x4a\x4c\x80\x54\xe5\x19\x62\x54\xf0\x3d\x1c\x14\x1a\x99\x0c\x29\xe2\x69\x09\x32\x2d\xc1\x4d\x4b\xc4\xd3\x8a\x2f\x44\x84\x21\x46\x33\x2f\x5a\x90\xbe\x08\xae\x2f\x22\xee\x8b\x48\xfa\x22\xb8\xbe\x88\xd9\xbe\x08\xd2\x17\xc1\xf5\xa5\x88\xfb\x32\x00\x3c\x64\x60\x62\xa3\x98\xc4\x43\x16\xa4\xb1\x05\xd7\xd8\x22\x6e\x6c\x91\x34\xb6\xe0\x1a\x5b\xcc\x36\xb6\x20\x8d\x2d\xb8\xc6\xa2\x4d\x96\x6e\xb1\x99\x0d\x96\x6e\x2f\x82\xfc\xb0\xe1\x5e\x9b\xdc\x6a\x31\x11\xd9\x3b\x96\x38\xe0\xba\xa3\xcb\x30\x73\x15\xa6\x14\xf1\xce\xa8\x7b\x36\xe3\x9f\x67\x18\x25\xfe\x57\x2a\xf1\xcf\x33\xe2\x3f\x12\x89\xff\x2d\x02\x89\xfb\x2d\x6f\x7e\x4b\xbc\x6f\x53\xeb\xb8\xf0\xa4\x6c\xbe\x68\x52\x32\x57\x70\x54\x2e\x2d\x96\x2d\x15\x17\x8a\xca\x54\xed\x49\xb6\x2d\x54\xf8\x9f\x00\x4c\x47\x47\x8d\xa6\xea\x6b\xd8\x34\xb2\x42\x40\x48\x2b\x13\xe2\xe1\x8c\xa8\x47\x69\xdd\xff\x07\x81\x7d\xd5\xb5\x33\x72\x37\xa2\x65\x55\xad\x64\x5d\xff\x90\xe5\x4f\x8e\xf3\x53\xba\x0a\xd6\xca\xf2\x38\x1c\x4c\xae\xd3\xe7\x68\x9c\x72\x2f\xed\x28\x3d\xb7\x0c\x85\x43\x43\x1f\x62\x38\xdc\x2e\x77\x46\x35\x46\xd9\x4b\x90\xa6\xd6\x20\xec\xfa\xed\x28\x7c\x7f\x4a\x38\x3c\xf2\x77\xa9\xec\x78\xeb\x4a\x99\x34\xf0\x4d\xd6\x3d\xe4\x19\xc7\xcd\x04\xdf\x7c\xe5\x38\x12\x6c\x94\x85\xc4\x3a\x22\xe7\xc2\x63\xf3\x88\x0c\xe1\xbd\x39\xb8\xcb\xd1\xd4\xde\x40\x45\x7d\xec\xcd\x61\x0f\x07\x14\xb3\x87\xc3\x14\xf0\x37\x9c\x89\x4d\x60\x50\x50\x6c\xe2\xe0\x24\x8f\x17\x87\xf8\x22\x8b\x89\xf7\xcf\xec\xbf\x00\x00\x00\xff\xff\xcb\xa4\x9f\xea\x9e\x15\x00\x00"

func cuda_master_functionsTxtBytes() ([]byte, error) {
	return bindataRead(
		_cuda_master_functionsTxt,
		"cuda_master_functions.txt",
	)
}

func cuda_master_functionsTxt() (*asset, error) {
	bytes, err := cuda_master_functionsTxtBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "cuda_master_functions.txt", size: 5534, mode: os.FileMode(420), modTime: time.Unix(1569725801, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _cudart_functionsTxt = "\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x8c\x57\xcd\x72\xdb\x20\x10\xbe\xf7\x69\x9a\xf8\xd4\xa3\x63\xc5\x6e\xa6\xd6\xc4\xf5\x4f\x72\xcc\x60\xb4\xb6\x69\x25\x50\x57\x28\x13\xe7\xe9\x3b\x32\x42\xa0\x05\xd4\xde\xd8\xef\xdb\x5f\x76\x41\x88\xb7\x05\x7b\x10\xb2\xd8\xb5\x78\x62\x1c\xf6\x6a\x8e\xc8\xae\x5f\x2c\xbc\x87\x0f\xdd\x22\x50\xf9\x3e\xa3\x48\xc2\x70\xaf\x72\x51\x57\xac\xae\xa1\x70\xfc\x42\xc9\x93\x38\xb7\x08\x0b\x56\x96\x06\x41\x60\x1a\xfa\x1c\x9e\x8f\xbf\x80\x6b\x0f\xef\x7d\x79\x78\x36\xcb\xee\xbe\xe6\xac\xde\x42\xa3\x5a\xe4\xd0\x38\x74\x0b\x67\xd1\x68\x40\x4b\x39\xe6\x20\xab\xa8\xc5\x41\x62\xc2\xe6\x5b\x2c\x44\x07\xbe\x00\x6a\xf8\x78\x68\x4f\x27\xc0\x01\x4f\x45\x1e\x88\xa8\x55\x3c\x2b\x03\x27\x0c\x92\xe9\x3a\x2a\xb4\x84\x46\xa3\xba\x86\x5b\xdc\x13\x91\x3d\x86\x77\xc1\x61\x0b\x0d\xf8\xf2\xee\x2a\xf9\x05\x95\x14\x9f\x26\xec\xe3\x6a\xbd\xd3\x08\xac\x5a\x28\xd9\xb4\x15\xe0\x9c\xff\x69\x05\xc2\x12\x59\x95\xd0\x58\x28\x29\x6d\x90\x14\xf9\x2a\xf4\x65\x59\xb2\x73\x13\xd7\xca\x44\xc3\xa7\xbc\x6c\xa1\x04\xd6\x4c\x25\x91\xd6\xd8\xa0\x2a\x5a\x9e\x48\xd3\x92\xa9\x04\x2c\xbf\x05\xdd\xa2\xf4\xbc\xbf\x83\xd4\x66\x9a\xa9\xbc\x44\x55\x75\xf6\x57\xc9\x29\x45\x76\xa1\x23\xfa\x76\x39\xe0\x67\x0b\xe8\x89\x5b\xe0\x0a\x0b\x2a\xbf\xd5\xba\xf9\x74\x20\x6d\xe2\x12\xc1\x2d\xdc\x39\xed\xa4\xef\xaa\xd1\x83\x10\x39\xcb\xab\x75\xce\x6a\x33\x67\xde\xec\x04\xe8\xbc\xb1\xe5\xad\xd6\xf6\x38\x44\xac\x6e\x53\xff\x3f\xb8\xef\xcf\x4d\x7d\x68\x89\xac\xbe\x08\xde\xa4\x6f\x06\x5f\xe3\xee\x5f\x1a\xf1\x23\x6e\x15\x1e\x5d\x69\x4f\x15\x3b\x8f\x49\x5a\x76\x82\x0c\x0d\x83\x2b\xc8\x12\x91\x8b\xc3\x51\xd1\x2b\xc2\xd2\x2f\xd9\x66\x7e\xb0\x11\x9f\x5b\x5d\xb7\xba\xbf\x17\xd2\x6a\x2f\xa2\x00\xe5\x6b\x75\x83\x31\x2f\x4b\xc5\x07\xc9\xaa\x0e\x80\x4b\xe3\x06\xad\x59\x2b\xf9\xc5\x5b\x2e\x94\xaa\x01\x99\x16\xef\xf0\x03\x50\x42\x39\xc5\xe5\x6d\xa9\x85\xb9\x83\xa6\xd4\xdc\xa4\x1b\x85\xc0\x71\x5c\xcd\xc9\x39\x1b\x8a\x32\xcb\x59\x36\x12\xdc\xe8\x1b\x84\xca\xc3\x81\x31\x62\xce\x24\x3b\x43\xe1\x23\xe1\x21\x32\xc4\x46\xe8\x7e\x7b\x72\xa8\x56\xa0\x9f\xe4\x49\x59\x71\x83\x70\x02\xcd\x2f\x6e\xee\x29\xe8\x55\x00\x15\xaf\xaf\xde\xb2\xff\x5a\x5b\xe1\x16\xd7\xff\x5e\xc7\x88\xb7\x5a\x17\xcd\x98\xf5\x43\x3b\x84\xc6\xbd\xcf\xba\x0b\x2d\x74\x3e\xa0\xa1\x9f\x31\x35\xe1\x30\x4c\x2a\x56\x46\x8f\x85\x71\x7c\x22\x8c\x92\x2c\x9c\x22\xb3\x6c\x24\xd0\x38\xb3\xf8\xbe\xcc\xb2\x0d\xf4\x07\xc1\x07\x42\xeb\x01\x8d\x7b\x08\xb3\xa1\x48\xa2\xbf\x53\xdd\xa5\x59\x44\xe3\xc7\xfa\x9a\xec\xea\x74\x4f\x93\x1d\xed\x88\xdd\xb5\x3a\xaa\x32\x0a\xc6\xa2\x78\x4c\x2c\x8c\xa1\x69\x1c\xd2\x8b\x58\x27\xc2\x1d\x4c\x0c\xd6\xd4\x58\x25\xf6\x7b\xaf\x82\x22\x2d\x14\x7a\x9f\x28\xd0\x92\xd4\xff\x48\xb6\x0f\x38\xb3\x74\x97\xc1\x4d\x18\x85\x73\xc8\x28\xd0\x0d\xa6\x2e\x67\xd9\x48\xa0\x7e\x62\xa7\xe0\x06\x53\x3f\xd4\x30\x6a\xe6\x8c\xba\x4e\x8d\x3e\x34\x1d\x40\x3e\x34\xe6\x29\x36\x2f\x8a\xee\xff\xe2\xc8\xf8\xef\x38\xea\x82\xf4\x94\xd6\x8c\x5f\x72\xa8\x5c\x4a\x31\x82\x9a\x79\xaf\x3a\x1f\x18\xbf\xdd\x28\xb3\x41\xa1\x50\xe8\xab\x47\xfa\x2f\x3b\x83\xb8\xa7\x9d\x27\xd3\xf0\xf4\x2d\x17\xa0\xd4\xe0\x95\x09\x7d\x7b\x05\xc6\x30\xa7\xbc\xbf\x20\xb0\x82\x7a\x3f\xc8\xa3\xfb\xb9\xfb\x1b\x00\x00\xff\xff\xfb\x3a\xb7\xa9\x3c\x0e\x00\x00"

func cudart_functionsTxtBytes() ([]byte, error) {
	return bindataRead(
		_cudart_functionsTxt,
		"cudart_functions.txt",
	)
}

func cudart_functionsTxt() (*asset, error) {
	bytes, err := cudart_functionsTxtBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "cudart_functions.txt", size: 3644, mode: os.FileMode(420), modTime: time.Unix(1569725782, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _cudnn_functionsTxt = "\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x8c\x54\xdb\x6e\xdb\x30\x0c\x7d\xdf\xe7\x6c\xfb\x81\xdc\x57\xa0\x75\x03\x3b\xc0\x1e\x07\x42\x62\x12\xa2\x8a\x64\x50\xb4\x97\xf5\xeb\x87\x22\x72\x2a\x29\x52\x93\x37\xfb\x9c\xc3\x23\xf1\x22\xaa\x41\x5b\x3b\x53\x42\x23\x08\x39\x3b\x07\xf5\xf6\x17\x58\x7f\xab\xe0\x7f\xc6\x1f\x75\xea\x67\x4e\xad\x1d\x97\xcc\x02\x5c\xf0\xba\x32\x93\x95\xd6\x3b\xb4\xde\xf1\xe5\x77\x0e\xa2\x8e\x8d\xe3\x13\x18\x7a\x2f\xdc\xb7\xce\xaf\xce\x35\x45\x38\xf2\xc9\xee\x91\xd1\x2a\xbc\xa3\xdb\x31\x90\x25\x7b\x78\x50\x36\x9d\xbb\xd8\x2d\x9e\x9d\xf7\xe1\xc7\xd9\xd1\x99\x21\xbe\xdf\x9c\xa0\x4e\x2e\x41\xa0\x4a\xae\xc9\x08\xf2\x2d\x4d\xe0\x2b\x8d\x88\x54\x29\xce\x08\x82\xf1\xf7\xcc\x1c\x1c\x93\x1c\x4f\x5b\xe4\xfd\x47\x96\xd7\xfa\x2c\xd1\x0b\xbb\x7f\xc9\xcf\x5d\xf5\x16\xd9\x93\x17\xb4\xd2\x36\xcd\xd6\x80\x0d\x2c\x8d\xe4\x69\xc4\x2f\x1a\x5b\x94\x24\x97\x5f\xb2\xeb\xdd\x20\x59\xd8\x05\x2c\x09\x37\x28\x2d\x7a\xe4\x11\xbb\x1e\x14\x76\xf4\x8e\x39\xdf\x09\x08\xfa\x4f\x66\x4d\x56\x57\xba\x73\x4d\xfd\x71\xe5\x34\x19\x15\xed\xa5\xad\x8f\xf9\x66\xda\x8a\x73\xa8\xc2\xd7\x96\xb9\x28\xf6\x6a\x9b\xe6\x6e\x26\x91\xe6\x37\xd2\xe1\x28\xbe\x26\xcb\xdf\xdd\x1d\xdd\xf4\xa0\x6e\x65\x83\x47\xfd\xda\xfb\xd5\x19\xd5\x30\x8d\xef\xd3\xe9\xfb\xc2\x99\xcb\xf7\x73\xdb\x2c\xd8\x79\xbf\x38\x82\xb5\x68\xd2\x09\xc9\xc8\x64\x52\x5e\xe0\x0d\x27\xf7\xcf\x69\x7d\x19\x8c\xd0\x2f\x04\x3d\x13\x29\xbc\xd0\x22\x1d\x4a\x51\x50\x24\x07\xbe\xf6\xf1\xb2\xdb\x3a\x67\xc8\x1e\xd2\xfb\x06\x30\x09\xcb\xfa\x52\x04\xa7\x6a\xdd\xf6\xa7\x86\x47\x11\xe5\x15\x59\x20\x6e\x63\xd2\x75\x79\x8b\x5f\x23\x50\x0f\x0a\xe3\xf4\x5b\x74\xac\x91\xc3\x6c\xdb\x68\x43\xb6\xe8\xc5\x31\x66\x83\x1c\xd0\xf0\x7c\x97\xe8\x15\x53\x2f\x93\x5b\x07\x63\x1e\xd0\x29\x30\xc9\x89\x9d\xdb\xcb\x09\xce\x69\xc1\x03\x98\x14\xbc\xeb\x41\x08\xcc\x6e\xbf\x61\xd2\x1b\xb4\xc8\x20\x8e\xb3\xb8\xa2\xa6\x6c\xd3\xc1\xa9\x37\x58\x33\x08\x6c\x12\xba\x63\xb0\xfe\x63\xc7\xc6\xdb\xff\x0a\xc6\x49\x65\xe0\xea\xfc\x3f\x00\x00\xff\xff\x52\x74\x7b\xd2\xf1\x07\x00\x00"

func cudnn_functionsTxtBytes() ([]byte, error) {
	return bindataRead(
		_cudnn_functionsTxt,
		"cudnn_functions.txt",
	)
}

func cudnn_functionsTxt() (*asset, error) {
	bytes, err := cudnn_functionsTxtBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "cudnn_functions.txt", size: 2033, mode: os.FileMode(420), modTime: time.Unix(1569725810, 0)}
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
	"cuda_master_functions.txt": cuda_master_functionsTxt,
	"cudart_functions.txt": cudart_functionsTxt,
	"cudnn_functions.txt": cudnn_functionsTxt,
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
	"cuda_master_functions.txt": &bintree{cuda_master_functionsTxt, map[string]*bintree{}},
	"cudart_functions.txt": &bintree{cudart_functionsTxt, map[string]*bintree{}},
	"cudnn_functions.txt": &bintree{cudnn_functionsTxt, map[string]*bintree{}},
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

