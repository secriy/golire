package main

import (
	"flag"
	"fmt"
	"os"
	"time"

	"github.com/secriy/golire"
)

var (
	ports   = flag.String("p", "22,3389", "Port range, such as '1-28', '22,53,3389' and '22,49-80'.")
	timeout = flag.Int("t", 200, "Timeout for one scan task, the default is 200 milliseconds.")
	number  = flag.Int("n", 300, "The number of goroutines to be created.")
)

func main() {
	flag.Usage = func() {
		fmt.Fprintf(flag.CommandLine.Output(), "Usage of golire: \n  golire [HOST] [OPTIONS], HOST is in CIDR format, such as '192.168.1.0/24'\n\n")
		flag.PrintDefaults()
	}
	flag.Parse()
	// host must be specified
	host := flag.Arg(0)
	if host == "" {
		fmt.Println("Error: No host specified")
		flag.Usage()
		os.Exit(1)
	}
	// running
	fmt.Println("golire is running, now detecting host...")
	tasks, err := golire.NewTasks(host, *ports)
	if err != nil {
		fmt.Println("Error: " + err.Error())
		os.Exit(1)
	}
	fmt.Println("Surviving host: ")
	for _, v := range tasks.IPPool {
		fmt.Println(v)
	}
	fmt.Println("Detection complete, now scanning TCP ports...")
	tasks.Run(time.Duration(*timeout)*time.Millisecond, *number)
}
