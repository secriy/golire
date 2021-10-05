package cmd

import (
	"fmt"
	"strings"
	"sync"
	"time"

	"github.com/panjf2000/ants/v2"
	"github.com/secriy/golire/module"
	"github.com/secriy/golire/util"
	"github.com/spf13/cobra"
)

var portGors int
var portPool string
var timeout int

// portCmd represents the port command
var portCmd = &cobra.Command{
	Use:   "port [Host]",
	Short: "Scan all opened TCP or UDP ports.",
	Long:  ``,
	Args:  cobra.RangeArgs(1, 1),
	Run: func(cmd *cobra.Command, args []string) {
		module.Log(level) // set logger level

		ips := module.ParseHost(args[0])
		ports := module.ParsePort(portPool)

		// limit goroutine numbers
		if portGors > len(ports) {
			portGors = len(ports)
		}

		wg := &sync.WaitGroup{}

		p, _ := ants.NewPoolWithFunc(portGors, func(i interface{}) {
			if module.DetectTCP(i.(string), time.Duration(timeout)*time.Millisecond) {
				util.Print("TCP", "Found open port: "+strings.Split(i.(string), ":")[1])
			}
			wg.Done()
		}, ants.WithPreAlloc(true))
		defer p.Release()

		for _, ip := range ips {
			util.Print("SCAN", fmt.Sprintf("Detect Host %s", ip))
			for _, port := range ports {
				wg.Add(1)
				_ = p.Invoke(fmt.Sprintf("%s:%d", ip, port))
			}
		}

		wg.Wait()
	},
}

func init() {
	rootCmd.AddCommand(portCmd)

	portCmd.Flags().IntVarP(&portGors, "threads", "g", 500, "-g [Number]")
	portCmd.Flags().IntVarP(&timeout, "timeout", "t", 1000, "-t [Timeout]")
	portCmd.Flags().StringVarP(&portPool, "ports", "p", "22,80,3389", "-p [Ports], e.g. '-p 22-1024,3389,8080'")
}
