##  Hi, I`m CoderI421~

[English](https://github.com/CoderI421/pkg/README.md) | 中文版

## 介绍

`CoderI421/pkg` 是一个日常开发中，通用类的工具包，即拿即用。

很高兴，它能对你有帮助~

Limited ability, make it better~

### `pkg/file` 包

与 `文件操作` 相关的工具包，降低使用 go 原生库操作文件的难度，让你更快达成目标。

- func CheckExist(src string) bool

  功能：检查文件是否存在

- func CheckPermission(src string) bool

  功能：检查是否有操作文件的权限

- func GetExt(fileName string) string

  功能：从文件路径中获取文件的后缀

- func GetLineNum() string

  功能：获取被执行代码的行号

  例子： a\b.go:6

- func GetLineNumWithTrace() string

  功能：获取被执行代码的行号，以及项目中的所有链路调用的行号

  例子： a\b.go:6 < c\d.go:10 < cmd\main.go

- func GetSize(f multipart.File) (int, error)

  功能：获取文件的大小

- func IsNotExistMkDir(src string) error

  功能：如果文件夹不存在则创建

- func MustOpen(fileName, filePath string) (os.File, error)

  功能：最大努力打开文件，如果不存在，则创建