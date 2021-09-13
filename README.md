# golire

一个快速局域网扫描器，支持CIDR格式指定IP地址范围，使用指定数量的 Goroutine 并发处理大量任务。

## Usage

```shell
$ make build
$ golire.exe -h
```

```
┌─┐┌─┐┬  ┬┬─┐┌─┐
│ ┬│ ││  │├┬┘├┤
└─┘└─┘┴─┘┴┴└─└─┘

Author: Secriy
GitHub: https://github.com/secriy/golire

Usage:
  golire [Command] [flags]
  golire [command]

Available Commands:
  completion  generate the autocompletion script for the specified shell
  help        Help about any command
  ping        Find all surviving hosts using ICMP message.
  port        Scan all opened TCP or UDP ports.

Flags:
  -h, --help           help for golire
  -l, --level string   logger level, debug, info, warning, error and fatal
  -v, --version        version for golire

Use "golire [command] --help" for more information about a command.
```

## Example

### Ping

```shell
.\golire.exe ping 192.168.1.1/31
.\golire.exe ping 192.168.1.0/24
.\golire.exe ping 192.168.1.0/24 -l "debug"
.\golire.exe ping -l "info" 192.168.1.0/24
```

### PortScan

```shell
.\golire.exe port 192.168.1.1 -p 22,3389,8080,8081-8090
.\golire.exe port 192.168.1.1 -p 1-225 -t 5000
```