# 创建数据库
首先创建一个空的数据库 zaiyun
运行sql/zaiyun.sql

# 下载依赖
```sh
$ go mod tidy
```

# 运行
```sh
$ go run main.go
```

# 打包
Linux & MacOS
```sh
$ go build -o zaiyun main.go
```
Windows
```sh
$ go build -o zaiyun.exe main.go
```

Windows build Linux
```sh
$ $env:GOOS="linux"
$ go build -o zaiyun main.go
```

Windows build MacOS
```sh
$ $env:GOOS="darwin"
$ go build -o zaiyun main.go
```

Linux && MacOS build Windows
```sh
$ $env:GOOS="windows"
$ go build -o zaiyun.exe main.go
```
