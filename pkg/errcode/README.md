## 错误码处理


#### 使用 go generate + stringer自动生成

```
## 安装stringer
go install golang.org/x/tools/cmd/stringer@latest

## 代码注释
//go:generate stringer -type ErrCode -linecomment -output code_string.go

## 执行
go generate

## windows安装make

1. cmd以管理员身份安装

@"%SystemRoot%\System32\WindowsPowerShell\v1.0\powershell.exe" -NoProfile -InputFormat None -ExecutionPolicy Bypass -Command "iex ((New-Object System.Net.WebClient).DownloadString('https://chocolatey.org/install.ps1'))" && SET "PATH=%PATH%;%ALLUSERSPROFILE%\chocolatey\bin"

2. choco install make

## make 文件
all:
    go generate && go build .

```

#### 参考

- [https://segmentfault.com/a/1190000020158429](https://segmentfault.com/a/1190000020158429)
- [https://www.jianshu.com/p/a1ba81746fe8](https://www.jianshu.com/p/a1ba81746fe8)
- [https://stackoverflow.com/questions/32127524/how-to-install-and-use-make-in-windows](https://stackoverflow.com/questions/32127524/how-to-install-and-use-make-in-windows)