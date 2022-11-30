

## Hi, I`m CoderI421~

English | [中文版](https://github.com/CoderI421/pkg/blob/main/README_CN.md)

## Introduction

`CoderI421/pkg`  is a  useful common toolkit in daily development.

I will be glad if it can be of any help to you~

Limited ability, make it better~

- [Install](https://github.com/CoderI421/pkg/edit/main/README.md#install)
- [pkg/file package](https://github.com/CoderI421/pkg/edit/main/README.md#pkgfile-package)
  - [Example](https://github.com/CoderI421/pkg/edit/main/README.md#example)

## Install

With a [correctly configured](https://golang.org/doc/install#testing) Go toolchain:

```go
go get -u github.com/CoderI421/pkg
```

### `pkg/file` package

File related tool functions to reduce the difficulty of using native go package,  let you achieve target faster~

- func CheckExist(src string) bool

  Function: check if the file exists

- func CheckPermission(src string) bool

  Function: check whether it has permission to operate file

- func GetExt(fileName string) string

  Function: get the file final slash-separated element of path

- func GetLineNum() string

  Function: get executed code`s line num. 

  ex. a\b.go:6

- func GetLineNumWithTrace() string

  Function: get executed code`s line num and its trace link in its project

  ex. a\b.go:6 < c\d.go:10 < cmd\main.go

- func GetSize(f multipart.File) (int, error)

  Function: get the file size

- func IsNotExistMkDir(src string) error

  Function: check whether it has permission to operate file

- func MustOpen(fileName, filePath string) (os.File, error)

  Function: maximize trying to open the file. Create file if not exist

##### Example:

```go
// test/a/a.go
package a

import (
	"github.com/CoderI421/pkg/file"
)
func A() string {
    return file.GetLineNumWithTrace()
}
```

```go
//ex. GetLineNumWithTrace
package main

import (
	"fmt"

	"test/a"
)
func main()  {
    fmt.Println(a.A())
}

//output:
// a\a.go:8 < test\main.go:9
```
