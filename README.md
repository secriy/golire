# golire
一个快速局域网扫描器，支持CIDR格式指定IP地址范围，使用指定数量的 Goroutine 并发处理大量任务。

## Usage
```shell
$ make build
$ golire.exe -h
```
```
Usage of golire.exe:
  -c string
        Host, which represent by CIDR, like '192.168.1.0/24' (default "127.0.0.1/24")
  -n int
        The number of goroutines to be created. (default 300)
  -p string
        Port range, like '1-28', '22,53,3389' and '22,49-80'. (default "22,3389")
  -t int
        Timeout for one scan task, the default is 200 milliseconds. (default 200)
```

## Module

- Arguments handler
- CIDR and port parser
- Ping to host
- TCP port scanner

## Directory

```
.
│   task.go
│
├───cmd
│       main.go
│
├───scan
│       ping.go
│       ping_test.go
│       scan.go
│       scan_test.go
│
└───utils
        array.go
        check_sum.go
```