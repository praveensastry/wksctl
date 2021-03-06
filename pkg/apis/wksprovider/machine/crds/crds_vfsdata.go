// Code generated by vfsgen; DO NOT EDIT.

// +build !dev

package crds

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	pathpkg "path"
	"time"
)

// CRDs statically implements the virtual filesystem provided to vfsgen.
var CRDs = func() http.FileSystem {
	fs := vfsgen۰FS{
		"/": &vfsgen۰DirInfo{
			name:    "/",
			modTime: time.Date(2019, 8, 23, 13, 41, 0, 103080382, time.UTC),
		},
		"/cluster_v1alpha1_cluster.yaml": &vfsgen۰CompressedFileInfo{
			name:             "cluster_v1alpha1_cluster.yaml",
			modTime:          time.Date(2019, 8, 23, 13, 41, 0, 103080382, time.UTC),
			uncompressedSize: 2469,

			compressedContent: []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xec\x56\x4d\x8f\x13\x31\x0c\xbd\xcf\xaf\xb0\x7a\x6f\xa1\x02\x21\x34\x37\xd8\x05\x89\x03\x2b\xb4\x8b\xf6\x82\x38\xb8\x19\x6f\x1b\x9a\x89\x83\xe3\x19\x58\x21\xfe\x3b\x9a\xaf\x6e\xb7\x33\xe9\x2c\x9c\xc9\xa9\xf3\xfc\xec\x38\xcf\x89\x6b\x0c\xf6\x96\x24\x5a\xf6\x39\x60\xb0\xf4\x53\xc9\x37\x5f\x71\xb5\x7f\x1d\x57\x96\x9f\xd5\xeb\x0d\x29\xae\xb3\xbd\xf5\x45\x0e\x17\x55\x54\x2e\xaf\x29\x72\x25\x86\x2e\xe9\xce\x7a\xab\x96\x7d\x56\x92\x62\x81\x8a\x79\x06\x60\x84\xb0\x01\x3f\xdb\x92\xa2\x62\x19\x72\xf0\x95\x73\x19\x80\xc3\x0d\xb9\xd8\x70\x00\x0c\x7b\x15\x76\x8e\x64\xa9\xcc\x6e\xd8\x30\x87\xc5\x7a\xf5\x7c\x91\x01\x78\x2c\x29\x07\xe3\xaa\xa8\x24\x71\xd5\xff\xe8\x69\x59\x0c\x64\x9a\x40\x5b\xe1\x2a\x1c\x68\x83\xb5\xf3\xee\x77\xea\x53\xef\x18\x2d\x12\x5c\x25\xe8\x1e\x82\x67\x00\xd1\x70\xa0\x1c\xae\x1a\xb7\x80\x86\x8a\x06\xab\x36\xd2\x1f\xb5\x0f\x15\x15\xb5\x8a\x39\xfc\xfa\x9d\x01\xd4\xe8\x6c\xd1\x9e\xb4\x33\x72\x20\xff\xe6\xd3\x87\xdb\x17\x37\x66\x47\x25\x76\x20\x40\x10\x0e\x24\x6a\x87\x18\xcd\x3a\x92\xfd\x80\x01\xe8\x7d\x93\x42\x54\xb1\x7e\x7b\x80\xdb\xec\xe7\x48\xc7\xf2\x3f\x26\xf2\xe6\x1b\x19\x3d\xc0\x83\x6c\xc3\x9a\x4a\xae\xad\x4e\xa7\xcc\x15\xe9\x0f\x96\xfd\x63\x5b\xda\xab\xb5\x71\x31\x81\x9e\xf7\x69\x77\xb4\x85\xbc\x75\x6c\xf6\x09\x3b\x80\x55\x2a\x93\xc6\x84\x30\x53\x14\x14\xc1\xfb\x09\x86\xd0\xf7\xca\x0a\x15\x53\x5b\x2c\x8f\xf2\x9b\x30\x4f\x6a\x3d\xac\x48\x52\x5b\x43\x97\x5c\xa2\xf5\x53\xc1\xcf\x66\xde\x7b\xff\xd7\xf4\x09\x51\x97\x07\xb5\x46\x86\xe6\x5a\xa6\xd8\x5d\x65\xb2\x27\x6e\x1f\x84\x6b\x5b\x90\xdc\x9c\x3c\x25\x98\x29\x48\x8d\xae\xa2\x74\xf9\x13\x97\xa7\xf5\x7a\x2f\x5c\xfe\x4b\xfd\x4b\x34\x3b\xeb\xe9\xc2\x61\x4c\x16\x79\x2e\xc6\xf1\x91\xd3\x8c\xbf\xb8\x2a\x89\x83\x9e\x35\x27\x4c\x53\xf7\x60\x79\xd2\xbb\x66\x3b\x62\xd7\xd2\x9f\xd0\x13\x31\xd8\x77\xbe\x08\x6c\xbd\x8e\xb4\x4a\xbc\xa3\xf3\xe2\xee\x38\xea\xb4\xa4\x33\x62\x06\x96\x84\xe3\x1d\x4b\x89\x9a\x83\xf5\xfa\xea\xe5\x99\xd0\xd6\x2b\x6d\xfb\x7f\xc3\xe3\x95\x7e\xaf\xcb\x36\xdd\x09\xb8\x49\x66\x04\xcf\xd6\x72\xdc\x30\x48\x84\xe5\x23\xc5\x88\xdb\xd1\x33\x49\xea\xd1\x3a\x5d\x13\x46\x1e\x75\xd6\xa4\xcf\xe1\x01\x8f\x2a\x7f\x26\xf3\x13\xb8\x1e\x86\xa6\x7a\x8d\x2e\xec\x70\x9d\x3d\xdc\x23\x34\x86\x82\x52\x71\x75\x3a\x82\x2c\x16\x8f\xa6\x8f\xf6\xd3\xb0\x2f\xda\x11\x2a\xe6\xf0\xe5\x6b\x33\x74\x28\x0b\x15\xfd\x74\xd0\x81\x7f\x02\x00\x00\xff\xff\x28\xd8\xc3\xb7\xa5\x09\x00\x00"),
		},
		"/cluster_v1alpha1_machine.yaml": &vfsgen۰CompressedFileInfo{
			name:             "cluster_v1alpha1_machine.yaml",
			modTime:          time.Date(2019, 8, 23, 13, 41, 0, 103080382, time.UTC),
			uncompressedSize: 2719,

			compressedContent: []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xe4\x56\xcd\x8e\xd3\x30\x10\xbe\xe7\x29\x46\xbd\xb7\x50\x71\x41\xb9\xa1\x45\x48\x1c\x76\x59\xed\xc2\x5e\x10\x87\xa9\x3d\x6d\xcd\x3a\xb6\x19\x4f\x22\x56\x88\x77\x47\x4e\x9a\xfe\x26\x6e\x2b\xc1\x09\xdf\x32\x33\xdf\x37\xbf\x9e\x18\x83\x79\x22\x8e\xc6\xbb\x12\x30\x18\xfa\x29\xe4\xd2\x57\x9c\x3d\xbf\x8d\x33\xe3\x5f\x35\xf3\x05\x09\xce\x8b\x67\xe3\x74\x09\x37\x75\x14\x5f\x3d\x50\xf4\x35\x2b\x7a\x4f\x4b\xe3\x8c\x18\xef\x8a\x8a\x04\x35\x0a\x96\x05\x80\x62\xc2\x24\xfc\x6c\x2a\x8a\x82\x55\x28\xc1\xd5\xd6\x16\x00\x16\x17\x64\x63\xb2\x01\x50\xde\x09\x7b\x6b\x89\xa7\xe2\xbd\xed\x1d\x96\x30\x99\xcf\x5e\x4f\x0a\x00\x87\x15\x95\x50\xa1\x5a\x1b\x47\x71\xa6\x6c\x1d\x85\x78\x63\x56\xc4\x40\x2a\x11\xad\xd8\xd7\xa1\x84\x23\x6d\x87\xde\x78\xea\x42\xbf\xed\x88\x5a\x49\xb0\x35\xa3\xdd\x91\x17\x00\x51\xf9\x40\x25\xdc\x25\x58\x40\x45\x3a\xc9\xea\x05\x6f\x52\xdd\x50\x45\x41\xa9\x63\x09\xbf\x7e\x17\x00\x0d\x5a\xa3\xdb\x4c\x3b\xa5\x0f\xe4\xde\xdd\x7f\x7c\x7a\xf3\xa8\xd6\x54\x61\x27\x04\x08\xec\x03\xb1\x98\x9e\x23\x9d\xbd\xb2\x6f\x65\x00\xf2\x92\x42\x88\xc2\xc6\xad\xb6\xe2\x36\xfa\x73\x46\xfb\xe5\x3f\x34\xf4\x8b\xef\xa4\x64\x2b\xee\xcb\xd6\x9f\xa1\xe0\x36\xdd\x59\x9a\xd5\x63\x9b\xfb\xa1\x66\x84\x78\x2c\x8a\x2c\x20\xb0\x6f\x8c\x26\x7e\x3c\x8a\x2a\x17\x59\x3a\x0d\xda\xfa\x24\xac\xac\xa7\x2d\xea\x03\xfb\x6a\x08\x99\x73\xd7\xe6\xd6\x8d\xca\x8d\xc5\x38\x62\x71\x9e\x63\x3f\xe5\x71\x8b\x91\x16\x5f\x91\x68\x56\x3d\xaa\x12\x34\x4e\x4e\x22\x37\x42\xd5\x40\x3a\x67\x1d\x20\x33\xbe\x1c\x68\x9a\x6e\xe0\x4f\xc8\x72\x55\xdb\x2c\x89\x7b\x8b\x2e\xd3\xef\x91\x4a\x3d\xd7\x0b\xb2\x24\x57\xe2\x98\x7e\xd4\x86\x49\x1f\xc3\xa6\x3d\xdf\x65\xf5\x1c\xa2\x99\x1e\x0c\xfc\xd9\xab\xda\xed\x9a\x0b\x2e\x2b\x6a\xcd\x14\xe3\x69\x01\xff\x5e\xf3\x94\x77\xba\xdd\xf4\xff\xd0\x07\x31\x7b\xbe\xa5\x18\x71\x35\xb2\x73\x06\x1a\xd6\x82\x1e\x08\xe3\xe1\x2a\xcd\x62\x2c\x46\xf9\x14\x88\xf7\x76\xf7\xee\xe4\xc6\x51\x53\x54\x6c\xc2\x10\x2c\xeb\xb0\x77\xfa\x25\x68\x94\xd3\xd1\x4a\x67\xe9\xb9\x42\x29\x21\x19\x4c\xc5\x54\x74\x2d\x7f\x1a\x97\xeb\xef\x48\xab\xbc\x0e\x34\xda\xd9\x4c\x86\xe7\xb2\x1b\x75\xe7\xbc\xa6\x07\x5a\x5e\xfe\x47\x59\x63\xbc\x7c\x78\xb6\xd7\xf1\xe4\xaa\x65\x9d\xfc\x5f\x9b\xec\x48\xdc\xf4\xcf\xc5\x66\x8e\x36\xac\x71\x5e\xec\x16\x15\x2a\x45\x41\x48\xdf\x1d\x3f\xbe\x26\x93\x83\x77\x57\xfb\xb9\xb7\x52\xe0\xeb\xb7\xf4\xdc\x12\xcf\xa4\x9f\xfa\xe2\x26\xe1\x9f\x00\x00\x00\xff\xff\x8b\x40\xfd\xc2\x9f\x0a\x00\x00"),
		},
		"/cluster_v1alpha1_machineclass.yaml": &vfsgen۰CompressedFileInfo{
			name:             "cluster_v1alpha1_machineclass.yaml",
			modTime:          time.Date(2019, 8, 23, 13, 41, 0, 103080382, time.UTC),
			uncompressedSize: 667,

			compressedContent: []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x84\x52\xb1\x8e\x14\x31\x0c\xed\xf3\x15\xd6\xf6\x37\x30\xa2\x41\xe9\xd0\xd1\x50\x70\x42\x1c\xba\x06\x51\x78\x13\x73\x1b\x2e\x13\x07\xdb\x19\xc1\xdf\xa3\xcc\xcc\x8e\x76\xb7\xe0\xd2\xf9\xe5\xe5\xbd\xbc\x27\x63\x4d\x4f\x24\x9a\xb8\x78\xc0\x9a\xe8\x8f\x51\xe9\x93\x0e\x2f\xef\x75\x48\xfc\x66\x1e\x8f\x64\x38\xba\x97\x54\xa2\x87\xfb\xa6\xc6\xd3\x57\x52\x6e\x12\xe8\x23\xfd\x4c\x25\x59\xe2\xe2\x26\x32\x8c\x68\xe8\x1d\x40\x10\xc2\x0e\x7e\x4b\x13\xa9\xe1\x54\x3d\x94\x96\xb3\x03\xc8\x78\xa4\xac\x9d\x03\x10\xb8\x98\x70\xce\x24\x77\xc6\x9c\xcf\x86\x1e\x0e\xe3\xf0\xf6\xe0\x00\x0a\x4e\xe4\x61\xc2\x70\x4a\x85\x42\x46\x55\xd2\x21\xe4\xa6\x46\xb2\x91\x9d\x56\x0a\x5d\xee\x59\xb8\x55\x0f\x37\xb7\xab\xc6\xe6\xb7\x06\xf8\xbc\xca\xdd\x77\xb9\x05\xae\xb9\x09\xe6\x5b\x1f\x07\xa0\x81\x2b\x79\x78\xe8\x0a\x15\x03\x45\x07\x30\x63\x4e\x71\xc9\xb6\x6a\x72\xa5\xf2\xe1\xcb\xa7\xa7\x77\x8f\xe1\x44\x13\xae\x20\x40\x15\xae\x24\x96\xce\xd6\xfd\x5c\x14\xbd\x63\x00\xf6\xb7\x7b\xa8\x49\x2a\xcf\x3b\xbc\xfc\xf4\x35\xd2\x65\xe1\xd7\x44\x3e\xfe\xa2\x60\x3b\x5c\x85\xe7\x14\x49\x1e\xb7\xaa\xfe\x43\x16\xfa\xdd\x92\xd0\x6e\x7e\x77\xf5\xb8\xe7\x3f\xaf\xca\x3c\x62\xae\x27\x1c\x9d\x1a\x5a\x5b\x62\x62\x08\x54\x8d\xe2\xc3\x6d\xe5\x87\xc3\x55\xd1\xcb\x18\xb8\xc4\x65\x71\xd4\xc3\xf7\x1f\xbd\x6d\x63\xa1\xb8\x35\xb4\x82\xff\x02\x00\x00\xff\xff\x7d\x29\x91\x1c\x9b\x02\x00\x00"),
		},
		"/cluster_v1alpha1_machinedeployment.yaml": &vfsgen۰CompressedFileInfo{
			name:             "cluster_v1alpha1_machinedeployment.yaml",
			modTime:          time.Date(2019, 8, 23, 13, 41, 0, 103080382, time.UTC),
			uncompressedSize: 3749,

			compressedContent: []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xc4\x57\x4b\x8f\xd3\x30\x10\xbe\xe7\x57\x58\xbd\xb7\x50\x16\x21\x94\x1b\xda\x15\x0f\x09\x96\x55\xcb\xee\x05\x71\x98\xd8\xd3\xd4\xac\x63\x1b\x7b\x12\x6d\x85\xf8\xef\xc8\x79\xd1\x64\x93\x26\x8b\x8a\xf0\xa9\x99\xf9\xfc\xcd\xb3\xf6\x18\xac\xbc\x43\xe7\xa5\xd1\x31\x03\x2b\xf1\x81\x50\x87\x2f\xbf\xba\x7f\xed\x57\xd2\x3c\x2b\xd6\x09\x12\xac\xa3\x7b\xa9\x45\xcc\x2e\x73\x4f\x26\xdb\xa0\x37\xb9\xe3\x78\x85\x3b\xa9\x25\x49\xa3\xa3\x0c\x09\x04\x10\xc4\x11\x63\xdc\x21\x04\xe1\x17\x99\xa1\x27\xc8\x6c\xcc\x74\xae\x54\xc4\x98\x82\x04\x95\x0f\x18\xc6\xb8\xd1\xe4\x8c\x52\xe8\x96\x64\x8c\x6a\x0c\xc6\x6c\xb1\x5e\x3d\x5f\x44\x8c\x69\xc8\x30\x66\x19\xf0\xbd\xd4\x28\xd0\x2a\x73\xc8\x50\x93\x5f\x71\x95\x7b\x42\x57\x6f\x88\xbc\x45\x1e\x28\x53\x67\x72\x1b\xb3\x9e\xb6\xe2\xa9\x6d\x56\x41\x7c\xaa\x28\xaf\x5a\xca\x52\x67\x55\xee\x40\x0d\x19\x8c\x18\xf3\xdc\x58\x8c\xd9\x75\xa0\xb2\xc0\x51\x04\x59\x9e\xb8\x3a\x11\x35\xbd\xe7\xa0\xb0\xfa\x59\xc7\xba\x45\x85\x9c\x8c\xbb\x01\xda\xc7\x6c\xe5\x09\x28\xf7\xab\x8e\xaa\x86\x87\x28\x36\x68\x95\xe4\xe0\x1b\xb4\x45\xbe\x72\xb5\xac\x81\x95\x0c\x7d\x60\x45\xdb\x81\x56\xb2\x98\xfd\xfc\x15\x31\x56\x80\x92\xa2\xac\x49\xe5\x9d\xb1\xa8\xdf\xdc\x7c\xb8\xbb\xd8\xf2\x3d\x66\xd0\xb8\x6c\x9d\xb1\xe8\x48\x36\xf1\x84\x75\xd4\x20\xad\x8c\x31\x3a\x84\x74\x78\x72\x52\xa7\xad\xb8\xcc\xee\x14\xe8\xb8\x51\xba\x40\x93\x7c\x47\x4e\xad\xb8\x29\x6b\xb3\x86\x9c\x2b\x09\xa5\xde\x20\x88\xc3\x16\xb9\xd1\xa2\xa7\x64\x6c\x67\x5c\x06\x14\x33\xa9\xe9\xe2\x45\x4f\x57\xd9\x95\x9a\x30\x45\xd7\xd1\x59\xc8\x3d\x8a\x3e\x57\x85\x4f\x8c\x51\x08\xba\x8b\x77\x26\x75\xe8\xfd\x15\x82\x50\x52\xe3\x59\x9d\x69\xea\x7a\x2e\xb6\x42\x86\x72\xbe\x97\x9e\x8c\x3b\x7c\x94\x99\xa4\xf3\x30\xfb\xba\xa3\x87\xd3\xd6\x2b\x6f\xb9\x81\x1c\x10\xa6\x87\xfe\x86\xb1\x52\x97\xee\x1b\xa5\xa4\x4e\x6f\xad\x00\xc2\xc7\xea\xd3\x9b\xc3\xca\xe0\x61\x9b\xbb\x74\x70\x6f\x58\x46\xe3\xe7\xdd\x98\x72\x39\xdc\xd4\xc3\xa0\xa1\x1c\x1d\x79\x71\xab\xa1\x00\xa9\x20\x51\xff\xd5\x97\xd1\xea\xb4\xca\xd1\x4d\x83\x96\x47\xf9\x08\x33\xab\x06\x8a\x76\xaa\x60\x43\xa7\xc5\x2c\xc7\xfb\xa7\xc7\x1c\x63\xac\xba\x95\x76\x32\xdd\x96\xa7\xfa\x58\xe2\x4f\x1a\x9e\xf2\x7a\x16\x81\x75\xa6\x90\x02\xdd\x76\x24\x8a\x39\x91\x84\x55\x80\xca\x47\xc3\x98\xe5\x49\xcb\xf2\xd6\x99\xec\x14\xd3\x1c\x77\x58\xd9\xf8\xe5\xf5\x7a\xa9\xc0\x4f\x20\xe7\x73\xb2\xa3\x94\x4d\x23\x4f\x76\xef\x30\x74\x22\x41\xb3\x60\x93\x10\x02\xa9\x69\x34\x52\x49\x98\x9d\x48\xc3\x6c\x07\xc0\x39\x38\x0c\x22\x8a\xea\x9a\x1f\x35\x32\xa7\x1a\xf5\x50\x77\xa3\x40\xcf\xe8\xbb\x89\x0a\xdc\xe7\x09\x2a\x7c\x74\x39\x3d\x91\xc7\xe1\x8f\x5c\xba\xc7\x97\x79\xb3\x96\x8d\x9d\xbf\xab\xdb\x29\xfa\x65\xe7\x8f\xfc\xb4\x23\x6c\x44\x35\x64\x6e\xd9\xde\xbc\x1d\x61\x73\xde\x4e\x0e\x5a\xd5\xa4\x38\x63\xd4\x6a\xef\xab\xcd\x59\x27\x12\x93\x78\x74\x05\x8a\x77\xa8\xd1\x1d\x4d\xa9\x83\xbc\xaf\x5e\x3e\x61\xd2\x01\x71\x38\xaf\xab\xe7\x1d\xc5\x72\xfd\x8f\x32\x9a\x97\xd3\x91\x38\x27\x69\xaf\x75\x8a\xe6\xd1\x58\xac\x41\xd9\x3d\xac\xa3\x3f\x6d\x04\x9c\xa3\x25\x14\xd7\xfd\x87\xd7\x62\xd1\x79\x69\x95\x9f\x61\x4e\x2e\x9f\x90\x3e\x66\x5f\xbf\x85\x67\x15\x19\x87\xe2\xae\x39\x8c\x82\xf0\x77\x00\x00\x00\xff\xff\xa6\xa2\xeb\xe1\xa5\x0e\x00\x00"),
		},
		"/cluster_v1alpha1_machineset.yaml": &vfsgen۰CompressedFileInfo{
			name:             "cluster_v1alpha1_machineset.yaml",
			modTime:          time.Date(2019, 8, 23, 13, 41, 0, 103080382, time.UTC),
			uncompressedSize: 3044,

			compressedContent: []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xac\x56\xcd\x8e\xdb\x36\x10\xbe\xeb\x29\x06\xbe\xdb\xad\xbb\x45\x51\xe8\x56\x6c\xd1\xa2\x40\x77\xb1\xb0\x83\xbd\x04\x39\x8c\xa9\xb1\xcd\x2c\x45\x32\xc3\x91\x10\x23\xc8\xbb\x07\xd4\x8f\x63\x69\x25\x4b\x9b\x98\x27\x69\xe6\xe3\x37\x3f\x1c\x72\x06\xbd\x7e\x26\x0e\xda\xd9\x14\xd0\x6b\xfa\x2c\x64\xe3\x5f\x58\xbd\xfc\x19\x56\xda\xfd\x52\xae\x77\x24\xb8\x4e\x5e\xb4\xcd\x52\xb8\x2f\x82\xb8\x7c\x43\xc1\x15\xac\xe8\x6f\xda\x6b\xab\x45\x3b\x9b\xe4\x24\x98\xa1\x60\x9a\x00\x28\x26\x8c\xc2\x77\x3a\xa7\x20\x98\xfb\x14\x6c\x61\x4c\x02\x60\x70\x47\x26\x44\x0c\x80\x72\x56\xd8\x19\x43\xbc\x14\xe7\x4c\x6b\x30\x85\xc5\x7a\xf5\xeb\x22\x01\xb0\x98\x53\x0a\x39\xaa\xa3\xb6\x14\x48\xc2\x4a\x99\x22\x08\x71\x83\x4c\x82\x27\x15\xb9\x0e\xec\x0a\x9f\x42\x4f\x5b\x13\x34\xc6\x6a\xef\x1f\x6a\xae\x2d\x49\x25\xf4\xa6\x60\x34\x1d\x13\x09\x40\x50\xce\x53\x0a\x8f\x71\xb3\x47\x45\x59\x94\x15\x3b\x6e\x62\x6e\x08\x83\x42\x43\xf5\x67\x13\xd6\x96\x0c\x29\x71\xfc\x84\x72\x4c\x61\x15\x04\xa5\x08\xab\x8e\xaa\x81\x47\xbf\x37\xe4\x8d\x56\x18\x5a\xb4\x27\xb5\xe2\x46\xd6\xc2\x2a\x86\x3e\xb0\xa6\xed\x40\x6b\x59\x0a\x5f\xbe\x26\x00\x25\x1a\x9d\x55\xe9\xaf\xbd\x73\x9e\xec\x5f\x4f\xff\x3d\xdf\x6d\xd5\x91\x72\x6c\x5d\xf6\xec\x3c\xb1\xe8\x36\x9e\xb8\x2e\x6a\xe1\x2c\x03\x90\x53\x4c\x47\x10\xd6\xf6\x70\x16\x57\xf9\x9c\x02\x5d\xd6\x44\x17\xe8\x76\x1f\x49\xc9\x59\xdc\x1e\x64\xbb\x86\x9c\xab\x08\xb5\xdd\x10\x66\xa7\x2d\x29\x67\xb3\x9e\x12\x60\xef\x38\x47\x49\x41\x5b\xb9\xfb\xad\xa7\xab\xed\x6a\x2b\x74\x20\xee\xe8\xda\x54\xde\x86\x2d\x34\x47\xdd\x67\x1b\x8c\xbb\x52\x50\xee\x0d\x0a\xf5\x37\x8c\xe5\x00\x46\x12\x3b\x69\x08\x06\x12\x3d\xc7\x18\xd4\x77\x75\xaf\x0f\xdb\xea\x02\x0c\x23\x26\x0c\x4f\x79\x3d\x8b\xc0\xb3\x2b\x75\x46\xbc\x1d\x89\x62\x4e\x24\x71\x95\x68\x8a\xd1\x30\x66\x79\x72\x66\xf9\x87\x5d\x7e\x8d\x69\x8e\x3b\x71\x35\x4f\xd0\xbd\xc1\x30\x81\x9c\xcf\x09\x17\x29\x9b\x46\x8e\xdc\xe1\x6b\xd0\x89\x04\xcd\x82\x4d\x42\x04\xb5\x95\xd1\x48\xb5\x50\x7e\x25\x0d\xb3\x1d\x40\x66\x3c\x0d\x22\xca\xfa\x45\x1c\x35\x32\xe7\x34\x9a\x56\xf7\x64\xd0\xce\xa8\xbb\x89\x13\x78\x29\x76\x64\x48\x7e\x92\x87\xe9\x53\xa1\x99\xb2\x31\x9a\x65\x6b\xe7\xc7\xce\xed\x1a\xfd\xb2\x73\x91\xdf\xf6\x84\x8d\xa8\x86\xcc\x2d\xcf\x6f\xf1\x64\xfb\xa9\xfb\xe7\x8c\x06\x84\x25\x6a\x83\x3b\x43\x9b\x9b\x36\x0d\x62\x76\xfc\x40\x21\xe0\xe1\x55\x7d\x8c\x9e\x65\xb5\x69\x43\x18\xba\xed\xfa\xea\x9e\x7d\x61\xcc\xe9\xff\x38\x92\x50\x76\xdb\x10\xdc\x2e\x10\x97\x94\xfd\x4b\x96\xf8\x62\xfc\x18\xe4\xfd\xe3\xf7\x37\x74\x67\xcc\x4e\xb7\x75\xf5\x96\x0d\x7f\xb8\xf0\x7a\x83\x1c\xbc\x2e\xbc\xb2\x9d\xb9\xcb\x35\x1a\x7f\xc4\x75\xf2\xbd\x08\x51\x29\xf2\x42\xd9\x63\x7f\x7c\x5d\x2c\x3a\x63\x6b\xf5\x1b\x07\xa1\x6a\x02\x0f\x29\xbc\xff\x10\x47\x55\x71\x4c\xd9\x73\xfb\x6a\x45\xe1\xb7\x00\x00\x00\xff\xff\x32\xcc\xbc\x93\xe4\x0b\x00\x00"),
		},
	}
	fs["/"].(*vfsgen۰DirInfo).entries = []os.FileInfo{
		fs["/cluster_v1alpha1_cluster.yaml"].(os.FileInfo),
		fs["/cluster_v1alpha1_machine.yaml"].(os.FileInfo),
		fs["/cluster_v1alpha1_machineclass.yaml"].(os.FileInfo),
		fs["/cluster_v1alpha1_machinedeployment.yaml"].(os.FileInfo),
		fs["/cluster_v1alpha1_machineset.yaml"].(os.FileInfo),
	}

	return fs
}()

type vfsgen۰FS map[string]interface{}

func (fs vfsgen۰FS) Open(path string) (http.File, error) {
	path = pathpkg.Clean("/" + path)
	f, ok := fs[path]
	if !ok {
		return nil, &os.PathError{Op: "open", Path: path, Err: os.ErrNotExist}
	}

	switch f := f.(type) {
	case *vfsgen۰CompressedFileInfo:
		gr, err := gzip.NewReader(bytes.NewReader(f.compressedContent))
		if err != nil {
			// This should never happen because we generate the gzip bytes such that they are always valid.
			panic("unexpected error reading own gzip compressed bytes: " + err.Error())
		}
		return &vfsgen۰CompressedFile{
			vfsgen۰CompressedFileInfo: f,
			gr:                        gr,
		}, nil
	case *vfsgen۰DirInfo:
		return &vfsgen۰Dir{
			vfsgen۰DirInfo: f,
		}, nil
	default:
		// This should never happen because we generate only the above types.
		panic(fmt.Sprintf("unexpected type %T", f))
	}
}

// vfsgen۰CompressedFileInfo is a static definition of a gzip compressed file.
type vfsgen۰CompressedFileInfo struct {
	name              string
	modTime           time.Time
	compressedContent []byte
	uncompressedSize  int64
}

func (f *vfsgen۰CompressedFileInfo) Readdir(count int) ([]os.FileInfo, error) {
	return nil, fmt.Errorf("cannot Readdir from file %s", f.name)
}
func (f *vfsgen۰CompressedFileInfo) Stat() (os.FileInfo, error) { return f, nil }

func (f *vfsgen۰CompressedFileInfo) GzipBytes() []byte {
	return f.compressedContent
}

func (f *vfsgen۰CompressedFileInfo) Name() string       { return f.name }
func (f *vfsgen۰CompressedFileInfo) Size() int64        { return f.uncompressedSize }
func (f *vfsgen۰CompressedFileInfo) Mode() os.FileMode  { return 0444 }
func (f *vfsgen۰CompressedFileInfo) ModTime() time.Time { return f.modTime }
func (f *vfsgen۰CompressedFileInfo) IsDir() bool        { return false }
func (f *vfsgen۰CompressedFileInfo) Sys() interface{}   { return nil }

// vfsgen۰CompressedFile is an opened compressedFile instance.
type vfsgen۰CompressedFile struct {
	*vfsgen۰CompressedFileInfo
	gr      *gzip.Reader
	grPos   int64 // Actual gr uncompressed position.
	seekPos int64 // Seek uncompressed position.
}

func (f *vfsgen۰CompressedFile) Read(p []byte) (n int, err error) {
	if f.grPos > f.seekPos {
		// Rewind to beginning.
		err = f.gr.Reset(bytes.NewReader(f.compressedContent))
		if err != nil {
			return 0, err
		}
		f.grPos = 0
	}
	if f.grPos < f.seekPos {
		// Fast-forward.
		_, err = io.CopyN(ioutil.Discard, f.gr, f.seekPos-f.grPos)
		if err != nil {
			return 0, err
		}
		f.grPos = f.seekPos
	}
	n, err = f.gr.Read(p)
	f.grPos += int64(n)
	f.seekPos = f.grPos
	return n, err
}
func (f *vfsgen۰CompressedFile) Seek(offset int64, whence int) (int64, error) {
	switch whence {
	case io.SeekStart:
		f.seekPos = 0 + offset
	case io.SeekCurrent:
		f.seekPos += offset
	case io.SeekEnd:
		f.seekPos = f.uncompressedSize + offset
	default:
		panic(fmt.Errorf("invalid whence value: %v", whence))
	}
	return f.seekPos, nil
}
func (f *vfsgen۰CompressedFile) Close() error {
	return f.gr.Close()
}

// vfsgen۰DirInfo is a static definition of a directory.
type vfsgen۰DirInfo struct {
	name    string
	modTime time.Time
	entries []os.FileInfo
}

func (d *vfsgen۰DirInfo) Read([]byte) (int, error) {
	return 0, fmt.Errorf("cannot Read from directory %s", d.name)
}
func (d *vfsgen۰DirInfo) Close() error               { return nil }
func (d *vfsgen۰DirInfo) Stat() (os.FileInfo, error) { return d, nil }

func (d *vfsgen۰DirInfo) Name() string       { return d.name }
func (d *vfsgen۰DirInfo) Size() int64        { return 0 }
func (d *vfsgen۰DirInfo) Mode() os.FileMode  { return 0755 | os.ModeDir }
func (d *vfsgen۰DirInfo) ModTime() time.Time { return d.modTime }
func (d *vfsgen۰DirInfo) IsDir() bool        { return true }
func (d *vfsgen۰DirInfo) Sys() interface{}   { return nil }

// vfsgen۰Dir is an opened dir instance.
type vfsgen۰Dir struct {
	*vfsgen۰DirInfo
	pos int // Position within entries for Seek and Readdir.
}

func (d *vfsgen۰Dir) Seek(offset int64, whence int) (int64, error) {
	if offset == 0 && whence == io.SeekStart {
		d.pos = 0
		return 0, nil
	}
	return 0, fmt.Errorf("unsupported Seek in directory %s", d.name)
}

func (d *vfsgen۰Dir) Readdir(count int) ([]os.FileInfo, error) {
	if d.pos >= len(d.entries) && count > 0 {
		return nil, io.EOF
	}
	if count <= 0 || count > len(d.entries)-d.pos {
		count = len(d.entries) - d.pos
	}
	e := d.entries[d.pos : d.pos+count]
	d.pos += count
	return e, nil
}
