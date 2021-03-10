# des.hex-encodefile
采用随机deskey和hex进行文件加密，常用于加密shellcode.
# 使用截图
![avatar](https://github.com/TRYblog/des.hex-encodefile/blob/main/9.png)
# 如何编译
Linux:
```
SET CGO_ENABLED=0
SET GOOS=linux
SET GOARCH=amd64
go build -ldflags "-s -w" main.go
```
Windows:
```
SET CGO_ENABLED=0
SET GOOS=windows
SET GOARCH=amd64
go build -ldflags "-s -w" main.go
```
# 时间
2021/03/10
