// Copyright 2022 The CoderI421. All rights reserved.

package file

import (
	"fmt"
	"io/ioutil"
	"mime/multipart"
	"os"
	"path"
	"path/filepath"
	"runtime"
	"strconv"
	"strings"
)

// GetSize get the file size
func GetSize(f multipart.File) (int, error) {
	content, err := ioutil.ReadAll(f)

	return len(content), err
}

// GetExt get the file final slash-separated element of path
func GetExt(fileName string) string {
	return path.Ext(fileName)
}

// getInfo get file info
func getInfo(src string) (os.FileInfo, error) {
	return os.Stat(src)
}

// CheckExist check if the file exists
func CheckExist(src string) bool {
	_, err := getInfo(src)
	return os.IsExist(err)
}

// CheckPermission check whether it has permission to operate file
func CheckPermission(src string) bool {
	_, err := getInfo(src)
	return !os.IsPermission(err)
}

// IsNotExistMkDir create a directory if it does not exist
func IsNotExistMkDir(src string) error {
	if exist := CheckExist(src); exist == false {
		if err := mkDir(src); err != nil {
			return err
		}
	}

	return nil
}

// mkDir create a directory
func mkDir(src string) error {
	return os.MkdirAll(src, os.ModePerm)
}

// MustOpen maximize trying to open the file. Create file if not exist
func MustOpen(fileName, filePath string) (*os.File, error) {
	dir, err := os.Getwd()
	if err != nil {
		return nil, fmt.Errorf("os.Getwd err: %v", err)
	}

	src := filepath.Join(dir, filePath)
	ok := CheckPermission(src)
	if !ok {
		return nil, fmt.Errorf("file.CheckPermission Permission denied src: %s", src)
	}

	err = IsNotExistMkDir(src)
	if err != nil {
		return nil, fmt.Errorf("file.IsNotExistMkDir src: %s, err: %v", src, err)
	}

	f, err := os.OpenFile(filepath.Join(src, fileName), os.O_APPEND|os.O_CREATE|os.O_RDWR, 0644)
	if err != nil {
		return nil, fmt.Errorf("fail to OpenFile :%v", err)
	}

	return f, nil
}

// GetLineNum get executed code`s line num. ex. a\b.go:6
func GetLineNum() string {
	for i := 1; i < 15; i++ {
		_, file, line, ok := runtime.Caller(i)

		// if ok && (!strings.HasPrefix(file, gormSourceDir) || strings.HasSuffix(file, "_test.go")) {

		if ok && !strings.HasSuffix(file, "_test.go") {
			dir, f := filepath.Split(file)

			return filepath.Join(filepath.Base(dir), f) + ":" + strconv.FormatInt(int64(line), 10)
		}
	}

	return ""
}

// GetLineNumWithTrace get executed code`s line num and its trace link in its project
// a\b.go:6 < c\d.go:10 < cmd\main.go
func GetLineNumWithTrace() string {
	var traceLine []string

	for i := 1; i < 15; i++ {
		_, fPath, line, ok := runtime.Caller(i)

		if ok && !strings.HasSuffix(fPath, "_test.go") {
			dir, f := filepath.Split(fPath)

			if strings.HasSuffix(dir, "runtime/") {
				return strings.Join(traceLine, " < ")
			}
			traceLine = append(traceLine, filepath.Join(filepath.Base(dir), f)+":"+strconv.FormatInt(int64(line), 10))
		}
	}

	return strings.Join(traceLine, " < ")
}
