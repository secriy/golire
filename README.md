# golire
一个快速局域网扫描器，支持CIDR格式指定IP地址范围，使用 Go 并发编程处理大量任务。

## 功能

- 存活主机探测
- 主机端口扫描

## 使用

```shell
$ make build
$ golire.exe -h
Usage of H:\Repositories\golire\golire.exe:
  -i string
        CIDR, like '192.168.1.0/24' (default "127.0.0.1/24")
  -p string
        All legal ports, like '1-28', '22,53,3389' and '22,49-80'. (default "22,3389")
  -t int
        Scan timeout for one task, set 2 seconds by default. (default 2000)
```

## 模块

- 命令行参数处理
- IP和端口解析
- TCP端口扫描

## 目录结构

```
│
├───cmd                 // 命令行
│       main.go
│
├───scan                // 扫描
│       ping.go
│       ping_test.go
│       scan.go         // 扫描指定地址
│       scan_test.go
│       task.go         // 解析IP地址以及端口
│       task_test.go
│
└───utils               // 工具
        array.go
        check_sum.go
        gen_ip.go
```