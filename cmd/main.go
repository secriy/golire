package main

import (
	"flag"
	"fmt"
	"os"
	"time"

	"github.com/secriy/golire"
)

var (
	ips     = flag.String("h", "127.0.0.1/24", "Host, which represent by CIDR, like '192.168.1.0/24'")
	ports   = flag.String("p", "22,3389", "Port range, like '1-28', '22,53,3389' and '22,49-80'.")
	timeout = flag.Int("t", 200, "Timeout for one scan task, the default is 200 milliseconds.")
	number  = flag.Int("n", 300, "The number of goroutines to be created.")
)

func main() {
	flag.Parse()
	fmt.Println("golire is running, now detecting host...")
	tasks, err := golire.NewTasks(*ips, *ports)
	if err != nil {
		fmt.Println("Err: " + err.Error())
		os.Exit(1)
	}
	fmt.Println("Surviving host: ")
	for _, v := range tasks.IPPool {
		fmt.Println(v)
	}
	fmt.Println("Detection complete, now scanning TCP ports...")
	tasks.Run(time.Duration(*timeout)*time.Millisecond, *number)
}
