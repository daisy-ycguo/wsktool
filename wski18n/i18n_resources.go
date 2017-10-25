/*
 * Licensed to the Apache Software Foundation (ASF) under one or more
 * contributor license agreements.  See the NOTICE file distributed with
 * this work for additional information regarding copyright ownership.
 * The ASF licenses this file to You under the Apache License, Version 2.0
 * (the "License"); you may not use this file except in compliance with
 * the License.  You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

// Code generated by go-bindata.
// sources:
// wski18n/resources/de_DE.all.json
// wski18n/resources/en_US.all.json
// wski18n/resources/es_ES.all.json
// wski18n/resources/fr_FR.all.json
// wski18n/resources/it_IT.all.json
// wski18n/resources/ja_JA.all.json
// wski18n/resources/ko_KR.all.json
// wski18n/resources/pt_BR.all.json
// wski18n/resources/zh_Hans.all.json
// wski18n/resources/zh_Hant.all.json
// DO NOT EDIT!

package wski18n

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

var _wski18nResourcesDe_deAllJson = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x01\x00\x00\xff\xff\x00\x00\x00\x00\x00\x00\x00\x00")

func wski18nResourcesDe_deAllJsonBytes() ([]byte, error) {
	return bindataRead(
		_wski18nResourcesDe_deAllJson,
		"wski18n/resources/de_DE.all.json",
	)
}

func wski18nResourcesDe_deAllJson() (*asset, error) {
	bytes, err := wski18nResourcesDe_deAllJsonBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "wski18n/resources/de_DE.all.json", size: 0, mode: os.FileMode(420), modTime: time.Unix(1501631495, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _wski18nResourcesEn_usAllJson = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xec\x5b\x5d\x6f\xdb\x3a\x12\x7d\xcf\xaf\x18\xe4\xc5\x2f\x81\xb6\xbd\x8b\x05\x16\x7d\x0b\x6e\xbf\x82\xb6\x69\x90\x64\x5b\x14\xdd\x02\x61\xc4\xb1\xc5\x9a\x22\x05\x92\x72\xe0\x1a\xfe\xef\x0b\x8a\x92\xed\x24\x14\x45\xc9\xb2\xb7\xbb\xb8\x79\x72\x64\xce\x39\x67\x86\x5f\x33\x14\xfd\xfd\x04\x60\x75\x02\x00\x70\xca\xe8\xe9\x2b\x38\x7d\x8f\x9c\xcb\xd3\x33\xf7\xc8\x28\x22\x34\x27\x86\x49\x61\xbf\x3b\x17\x70\x7e\x75\x01\x99\xd4\x06\xf2\x52\x1b\xb8\x47\x28\x94\x5c\x30\x8a\x34\x39\x3d\x01\x58\x9f\x3d\x85\xfb\xc4\xb4\x66\x62\x06\x69\x4e\x61\x8e\xcb\x16\xe0\xa6\xd5\x24\xcd\xe9\x04\x98\x28\x4a\x53\xb5\xf6\x42\xe6\x75\xe3\x9c\x08\x36\x45\x6d\x92\x25\xc9\x39\x4c\x19\xc7\x0e\x74\x8f\x81\x97\x80\x94\x26\x93\x8a\xfd\xaa\x00\xe0\xee\xc3\x9b\x6f\x77\x2d\xc8\xbe\x96\x5e\xc8\x87\x8c\xe9\x79\x15\xbc\xbb\xf7\x9f\x6f\x6e\xdb\xf0\x9e\x35\xf3\x82\x09\x92\xa3\x2e\x48\xda\xe6\xef\xf6\xfb\x2e\x2d\x5f\xde\x5c\xdf\x5c\x7c\xbe\x8c\x90\xb3\x69\xe9\xef\xe5\x3a\xb2\x55\x50\x41\x48\x03\x53\x59\x0a\x0a\xc4\x40\x41\x4c\x06\xab\x55\x52\x28\xf9\x13\x53\x73\x45\x4c\xb6\x5e\x27\xff\x16\x6d\x7d\x35\x00\x29\x38\xf0\x56\xab\xaa\xbb\xd7\xeb\xbf\xd9\x4f\xf6\x43\x05\x9d\xc0\x88\x9a\x0f\x41\x15\x11\x67\xa6\x1f\x63\x99\x0c\x1d\xde\xf7\xd5\x2a\xb1\x2d\x1c\xda\x8f\xd8\x68\xf7\xc1\xf3\xca\xfb\x57\x13\x86\x66\xa6\x39\x03\x98\x4a\x05\x14\x0b\x2e\x97\x39\x0a\xd3\x2e\x27\xde\xbe\x37\x7d\x29\xf6\x15\xf0\x14\xc1\x2b\xc1\x86\x4c\x95\xc2\xb0\x7c\x13\x4e\x5d\x16\x85\x54\x06\x29\xdc\x2f\xe1\x73\x81\xc2\xcd\xaa\x82\x13\x33\x95\x2a\x6f\x17\x33\x0c\xcb\x3f\xe5\xf5\xdc\x89\x87\x8c\x68\x48\x33\xa9\x51\x00\x81\x82\x28\xc3\xd2\x92\x13\xb5\x21\xb2\x9e\x5a\x62\x92\x5a\x19\xed\xe2\xf6\x41\xf4\x77\x9e\xd8\x3a\xd7\x98\x9a\x65\x81\x67\xa0\xd1\x80\x91\x20\x24\xc5\x9f\xba\xad\xe3\x22\xad\xbd\xd4\xb7\x56\x5d\x69\x32\x14\x86\xa5\x6e\x29\x9f\xe3\xb2\x89\x79\x2a\xc5\x94\xcd\x4a\x85\xb4\x3d\x1a\x7d\x10\x5a\x25\x6c\xb6\xd6\x9e\xc4\x61\xbb\x56\xba\xcd\x2e\xd1\x97\xaf\xc3\x30\xca\xbf\xd5\x2a\x21\x05\xb3\xff\xad\xd7\x67\x30\x55\x32\xaf\x1f\x69\x59\xaa\x14\x43\xab\xed\x20\xa8\x60\xbf\x37\x7d\xa5\xd1\xec\x00\x94\x26\x8b\x13\x13\x0d\x11\xd7\x15\xab\x55\xb2\xf9\x7f\xd7\xa3\xcd\xc3\x38\x55\xc3\x31\xbd\x32\xdf\x12\xc6\x91\xda\x99\x34\x43\xb7\x33\x3c\x9b\x70\xda\xc1\xda\x65\xe9\x6b\xb5\x2c\x69\x54\x0b\x96\xe2\x2b\xcb\x84\x4a\x85\x14\x8f\x06\xef\x15\x7f\x63\x88\xaa\x56\x81\x52\xe4\x44\xe9\x8c\xf0\x9d\xc5\x93\x89\xa9\x74\xd0\x5c\xa6\x84\xc3\x82\xf0\x12\x75\xbb\xd4\x81\x60\x2d\x8b\x5e\x08\x82\x09\x83\x4a\x60\x68\xbf\x8a\xb6\xf7\xd2\xbf\xde\x6c\x68\x90\xca\xbc\xe0\x68\xc3\xad\xcb\x34\x45\xad\xa7\x25\xe7\xcb\x76\xe6\x28\x53\x2f\xe9\x3b\x69\x00\x95\x92\x0a\x52\x85\xc4\xd8\x0d\xb7\x20\xe9\x9c\xcc\x10\x1e\x98\xc9\xea\xef\x72\xd4\x9a\xcc\x76\x3a\x17\x88\xa0\x8d\x9d\xa4\xee\x0b\xfb\x21\x34\xaa\x0e\x42\x15\xeb\x94\xdb\xef\xfe\x87\x7d\xda\xce\xb4\x3f\x39\xb3\x1d\xfd\xc6\x9a\xb7\xc8\x6a\x69\xec\x05\xbe\x10\x0b\xc2\x19\xad\xeb\x3c\x39\x85\x6f\x1d\x05\x5c\xc0\x20\xba\x37\x0a\x76\x8c\xae\xd8\x8f\x26\xd6\x19\x55\xf2\xa3\x4c\x96\x3d\x79\x3a\xdc\xd1\x68\x2a\x96\x6a\xb5\x37\xc4\x94\xda\x76\xed\x81\x7d\x3b\x08\x69\x6c\xbf\x19\xc5\x66\x33\x54\xc7\xe8\xba\xfd\xa9\xfa\x3a\x35\x45\xa4\xc7\xf4\x6c\x4f\xbe\xbe\x7b\xd3\x3d\x13\xd4\xfe\x7f\xc4\xf5\x7c\x7f\xca\xae\x5d\x5f\x4e\x6d\x5d\x8d\x82\xa2\x48\x97\xd6\x94\x62\x71\x49\x72\x5c\xaf\x81\x32\x5a\x27\xf7\x6e\x77\xb7\x9b\xfb\x66\x6f\x87\xeb\x52\xc0\xdd\xb6\x06\x6c\x6a\xe3\x3b\x9b\x17\x29\xcc\xe5\x02\x5d\x29\x48\x38\x5f\xd6\xa5\x3b\x52\x20\x5a\xa3\x09\xa4\x56\xbf\x83\xb2\x40\xc8\x76\x76\xf5\xd5\x2a\x91\xa5\x29\x4a\xb3\x5e\x43\x92\x24\x41\x7f\x02\x66\x1d\x64\xd5\xb2\xd4\x97\xca\x6b\xd4\x41\xf4\x68\x4e\xf5\x25\x0c\x1a\x77\x10\x37\x43\xbd\x2f\x67\x9b\x5d\x24\x5d\x33\xb3\x86\xd2\xb6\xd9\x77\xd0\x3f\x1e\xd1\xbd\x98\x03\xa6\xfe\xe4\xed\x43\x02\x7f\x12\x91\x22\xe7\xb5\x79\xe7\xe1\x57\xd0\xa4\x83\xc4\x1a\xc4\x1d\xb1\x85\x6d\x5a\xaa\xa4\x6d\xa3\xf0\xec\x0f\x14\x4a\x3d\x20\xba\x56\x4d\xb7\x8c\x0c\xa8\x94\xda\x0c\x7f\x5f\xaf\x1b\x84\xe7\x43\x70\xbb\x1e\xb7\x0f\xdf\x58\xeb\x6e\xff\x87\x6e\x07\x71\xa1\xd9\x03\xbd\x23\x8d\xa0\xc8\xf1\x38\xd5\xe0\x78\x4c\xb1\x2e\x1d\x38\x53\x1f\x8b\xe7\xaf\x2a\xe4\x49\x3c\x0f\x5f\x85\x8c\x48\xd5\xd7\xa9\x03\x57\x21\x63\xf3\xc5\xba\x77\xf8\x13\xb2\x11\xa9\xbc\x4e\x7d\x3d\xbf\xbe\xbc\xb8\x7c\xf7\x0a\x6e\x33\x84\x89\x3b\x6f\x9e\xc0\xb7\xf3\x4f\x1f\xdd\xd9\xb9\xd8\x84\x14\x85\x61\xa6\x3a\x4d\xa7\x58\x28\x4c\x89\x41\x9a\xc0\x15\x47\xa2\x11\x4a\x8d\x30\xb1\x41\x9f\x00\x13\xda\x20\xb1\xb9\x3b\x50\xd4\xa9\x62\xf7\x48\x2d\x8e\x2e\x30\x65\xd3\xfa\x3d\x4c\xa0\xda\xf8\x6f\x2a\x8a\xed\xf7\x26\xd1\x3c\x62\xff\x8f\x47\xd9\x96\x47\xf9\x12\xd9\x8c\x68\xb8\x47\x14\x8f\x72\xa3\x4d\x7d\x16\xcc\xae\x86\xc1\x79\xc5\x5d\x05\xf2\xfb\xa1\x0a\xf7\xc3\x0c\xca\x1c\x53\xde\x08\xb2\x6e\xeb\xe9\x32\x86\xac\x61\x58\x41\x59\xcf\xea\xd3\x7d\xb5\x0d\x00\xf4\x0a\xbc\x7e\x5a\xb1\x0f\x15\x36\x00\xc8\x2b\xe8\xfc\xf9\x69\xc5\x50\x49\x83\xa0\x22\x16\x8e\x6d\x86\xee\x87\x6b\x0a\xcb\xf8\xb5\xa3\x1f\x62\xd7\x04\xa8\xb7\x8e\x5d\xb4\xba\xfa\x8b\x1a\xf3\x61\xf3\xe0\x28\xb2\x59\x64\x2f\xda\x08\xc3\x8e\x51\xe2\x4a\x8e\x5e\xa4\x91\xc6\x1d\xef\xa2\x99\x58\xc8\x39\x56\x69\xb4\x4b\xc4\x32\x14\x47\xca\x09\x8f\x2c\x22\x9c\x53\x7d\x22\x82\x12\x23\xd5\x12\xa6\x0c\x39\x85\x66\x51\xff\x82\x4a\xdb\x30\x37\x97\x35\x75\xe8\x65\xf2\x50\xb4\xb0\xb4\xa7\xc6\xcd\x4d\x26\xb2\x70\xb9\x91\x0d\x5b\x5a\x2a\x65\xeb\xf0\xed\xb1\xed\xc2\xb5\x8e\x10\xbb\x37\xfe\xb0\xc8\x7e\x64\x29\x0a\x8d\x23\x45\xb6\x05\x2d\x2e\xb2\x8d\xf1\xa1\x22\x3b\x18\x3f\x2c\xbf\x81\x5d\xad\x12\xee\x3e\x5e\xbc\x5e\xaf\x1b\x16\x02\xee\x95\xaf\x14\x18\x21\xb5\x17\x96\xff\x52\x31\xe6\x52\x2d\x6f\xd8\x2f\xb4\x95\x38\x67\x39\x33\xda\x3a\xd9\x5c\x04\x04\x9d\xc9\x92\x53\xdb\x39\x44\x54\x37\x2b\xec\xa4\xbe\x47\xf3\x60\x57\xac\x97\x7f\xfc\xb3\x9a\xbe\xff\x78\xf9\x47\xbb\xda\x51\x29\xfc\xb7\x0f\x59\x8e\xd2\xbd\x21\xef\x0d\xff\xe2\x45\x05\xff\xf7\x17\xf6\x2f\x70\x29\x71\x4c\x0a\xaf\x13\x5c\xce\x86\xc6\xc8\xe1\xbf\x0c\xc8\x1f\x09\x3c\x3c\xb2\x9b\xeb\x0a\x15\x83\xbb\x0d\x38\xa9\xe3\x36\xb1\xcc\xf5\x09\xdd\x2e\x33\xd3\xc0\x66\x42\xaa\x9d\xba\x32\xcd\x30\x9d\xbb\x0d\x21\xa6\x88\x3d\x0c\x69\x7f\x47\xb7\xa3\xfc\xb8\xbe\x8e\xc2\xdb\xdf\xdd\x7a\x40\x1d\xd7\xd7\xfd\x49\xbb\x96\xe6\x6a\x6a\xb8\xe5\x34\x67\x46\xb8\xbc\x78\x73\xdf\x33\x23\x62\x46\xee\x39\x02\xd1\xee\x6e\xed\xc3\x19\x3c\x64\x2c\xcd\xe0\x81\x71\x6e\xe7\x4f\xc3\x1f\xb3\x74\x8f\xc8\xe5\x75\xab\xba\xc2\x2e\xa7\xd5\x7e\x95\xcc\x71\x19\xba\x69\xe4\x6f\xdb\x0d\x9b\xa2\x32\xd1\xb8\xdb\xc6\xc1\x1b\xa1\x4d\x7b\x52\x14\xbc\xb9\x55\x5c\xdd\x6c\xad\xea\x14\xf7\x5a\xc9\x6e\x72\x62\xe7\x55\x9d\xbb\xc8\x4f\x4c\x75\x65\x7f\xfb\xf8\x2d\xb3\xa9\xfd\x0f\xa0\x12\x5d\x58\x73\x62\xd2\xac\x42\xef\x66\x6a\x7e\x24\xb0\x76\x64\xf9\xa3\xdf\x0c\xd4\x54\xcd\xc3\x9a\xa8\xed\x85\xd0\xff\x85\x6b\xe1\xb9\xb3\x4b\xe4\xd1\x5f\x0d\x5a\x2d\xa5\xb0\x23\x77\x7b\x9a\x78\x06\xc5\xf6\x34\xb1\xfe\x2d\x48\x73\x9c\x18\x31\x89\x0e\x43\xda\xcb\xd1\xc7\xb1\x3b\x8a\x9b\xa3\x50\x5a\x27\x4f\x4e\x7e\x9c\xfc\x27\x00\x00\xff\xff\x3b\xec\xdb\xc6\x09\x37\x00\x00")

func wski18nResourcesEn_usAllJsonBytes() ([]byte, error) {
	return bindataRead(
		_wski18nResourcesEn_usAllJson,
		"wski18n/resources/en_US.all.json",
	)
}

func wski18nResourcesEn_usAllJson() (*asset, error) {
	bytes, err := wski18nResourcesEn_usAllJsonBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "wski18n/resources/en_US.all.json", size: 14089, mode: os.FileMode(420), modTime: time.Unix(1508886382, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _wski18nResourcesEs_esAllJson = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x01\x00\x00\xff\xff\x00\x00\x00\x00\x00\x00\x00\x00")

func wski18nResourcesEs_esAllJsonBytes() ([]byte, error) {
	return bindataRead(
		_wski18nResourcesEs_esAllJson,
		"wski18n/resources/es_ES.all.json",
	)
}

func wski18nResourcesEs_esAllJson() (*asset, error) {
	bytes, err := wski18nResourcesEs_esAllJsonBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "wski18n/resources/es_ES.all.json", size: 0, mode: os.FileMode(420), modTime: time.Unix(1501631495, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _wski18nResourcesFr_frAllJson = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x8a\xe6\x52\x50\xa8\xe6\x52\x50\x50\x50\x50\xca\x4c\x51\xb2\x52\x50\x4a\xaa\x2c\x48\x2c\x2e\x56\x48\x4e\x2d\x2a\xc9\x4c\xcb\x4c\x4e\x2c\x49\x55\x48\xce\x48\x4d\xce\xce\xcc\x4b\x57\xd2\x81\x28\x2c\x29\x4a\xcc\x2b\xce\x49\x2c\xc9\xcc\xcf\x03\xe9\x08\xce\xcf\x4d\x55\x40\x12\x53\xc8\xcc\x53\x70\x2b\x4a\xcd\x4b\xce\x50\xe2\x52\x50\xa8\xe5\x8a\xe5\x02\x04\x00\x00\xff\xff\x45\xa4\xe9\x62\x65\x00\x00\x00")

func wski18nResourcesFr_frAllJsonBytes() ([]byte, error) {
	return bindataRead(
		_wski18nResourcesFr_frAllJson,
		"wski18n/resources/fr_FR.all.json",
	)
}

func wski18nResourcesFr_frAllJson() (*asset, error) {
	bytes, err := wski18nResourcesFr_frAllJsonBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "wski18n/resources/fr_FR.all.json", size: 101, mode: os.FileMode(420), modTime: time.Unix(1501631495, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _wski18nResourcesIt_itAllJson = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x01\x00\x00\xff\xff\x00\x00\x00\x00\x00\x00\x00\x00")

func wski18nResourcesIt_itAllJsonBytes() ([]byte, error) {
	return bindataRead(
		_wski18nResourcesIt_itAllJson,
		"wski18n/resources/it_IT.all.json",
	)
}

func wski18nResourcesIt_itAllJson() (*asset, error) {
	bytes, err := wski18nResourcesIt_itAllJsonBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "wski18n/resources/it_IT.all.json", size: 0, mode: os.FileMode(420), modTime: time.Unix(1501631495, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _wski18nResourcesJa_jaAllJson = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x01\x00\x00\xff\xff\x00\x00\x00\x00\x00\x00\x00\x00")

func wski18nResourcesJa_jaAllJsonBytes() ([]byte, error) {
	return bindataRead(
		_wski18nResourcesJa_jaAllJson,
		"wski18n/resources/ja_JA.all.json",
	)
}

func wski18nResourcesJa_jaAllJson() (*asset, error) {
	bytes, err := wski18nResourcesJa_jaAllJsonBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "wski18n/resources/ja_JA.all.json", size: 0, mode: os.FileMode(420), modTime: time.Unix(1501631495, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _wski18nResourcesKo_krAllJson = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x01\x00\x00\xff\xff\x00\x00\x00\x00\x00\x00\x00\x00")

func wski18nResourcesKo_krAllJsonBytes() ([]byte, error) {
	return bindataRead(
		_wski18nResourcesKo_krAllJson,
		"wski18n/resources/ko_KR.all.json",
	)
}

func wski18nResourcesKo_krAllJson() (*asset, error) {
	bytes, err := wski18nResourcesKo_krAllJsonBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "wski18n/resources/ko_KR.all.json", size: 0, mode: os.FileMode(420), modTime: time.Unix(1501631495, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _wski18nResourcesPt_brAllJson = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x01\x00\x00\xff\xff\x00\x00\x00\x00\x00\x00\x00\x00")

func wski18nResourcesPt_brAllJsonBytes() ([]byte, error) {
	return bindataRead(
		_wski18nResourcesPt_brAllJson,
		"wski18n/resources/pt_BR.all.json",
	)
}

func wski18nResourcesPt_brAllJson() (*asset, error) {
	bytes, err := wski18nResourcesPt_brAllJsonBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "wski18n/resources/pt_BR.all.json", size: 0, mode: os.FileMode(420), modTime: time.Unix(1501631495, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _wski18nResourcesZh_hansAllJson = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x01\x00\x00\xff\xff\x00\x00\x00\x00\x00\x00\x00\x00")

func wski18nResourcesZh_hansAllJsonBytes() ([]byte, error) {
	return bindataRead(
		_wski18nResourcesZh_hansAllJson,
		"wski18n/resources/zh_Hans.all.json",
	)
}

func wski18nResourcesZh_hansAllJson() (*asset, error) {
	bytes, err := wski18nResourcesZh_hansAllJsonBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "wski18n/resources/zh_Hans.all.json", size: 0, mode: os.FileMode(420), modTime: time.Unix(1501631495, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _wski18nResourcesZh_hantAllJson = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x01\x00\x00\xff\xff\x00\x00\x00\x00\x00\x00\x00\x00")

func wski18nResourcesZh_hantAllJsonBytes() ([]byte, error) {
	return bindataRead(
		_wski18nResourcesZh_hantAllJson,
		"wski18n/resources/zh_Hant.all.json",
	)
}

func wski18nResourcesZh_hantAllJson() (*asset, error) {
	bytes, err := wski18nResourcesZh_hantAllJsonBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "wski18n/resources/zh_Hant.all.json", size: 0, mode: os.FileMode(420), modTime: time.Unix(1501631495, 0)}
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
	"wski18n/resources/de_DE.all.json": wski18nResourcesDe_deAllJson,
	"wski18n/resources/en_US.all.json": wski18nResourcesEn_usAllJson,
	"wski18n/resources/es_ES.all.json": wski18nResourcesEs_esAllJson,
	"wski18n/resources/fr_FR.all.json": wski18nResourcesFr_frAllJson,
	"wski18n/resources/it_IT.all.json": wski18nResourcesIt_itAllJson,
	"wski18n/resources/ja_JA.all.json": wski18nResourcesJa_jaAllJson,
	"wski18n/resources/ko_KR.all.json": wski18nResourcesKo_krAllJson,
	"wski18n/resources/pt_BR.all.json": wski18nResourcesPt_brAllJson,
	"wski18n/resources/zh_Hans.all.json": wski18nResourcesZh_hansAllJson,
	"wski18n/resources/zh_Hant.all.json": wski18nResourcesZh_hantAllJson,
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
	"wski18n": &bintree{nil, map[string]*bintree{
		"resources": &bintree{nil, map[string]*bintree{
			"de_DE.all.json": &bintree{wski18nResourcesDe_deAllJson, map[string]*bintree{}},
			"en_US.all.json": &bintree{wski18nResourcesEn_usAllJson, map[string]*bintree{}},
			"es_ES.all.json": &bintree{wski18nResourcesEs_esAllJson, map[string]*bintree{}},
			"fr_FR.all.json": &bintree{wski18nResourcesFr_frAllJson, map[string]*bintree{}},
			"it_IT.all.json": &bintree{wski18nResourcesIt_itAllJson, map[string]*bintree{}},
			"ja_JA.all.json": &bintree{wski18nResourcesJa_jaAllJson, map[string]*bintree{}},
			"ko_KR.all.json": &bintree{wski18nResourcesKo_krAllJson, map[string]*bintree{}},
			"pt_BR.all.json": &bintree{wski18nResourcesPt_brAllJson, map[string]*bintree{}},
			"zh_Hans.all.json": &bintree{wski18nResourcesZh_hansAllJson, map[string]*bintree{}},
			"zh_Hant.all.json": &bintree{wski18nResourcesZh_hantAllJson, map[string]*bintree{}},
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

