package cmd

import (
	"fmt"
	"sync"
	"sync/atomic"

	"github.com/panjf2000/ants/v2"
	"github.com/secriy/golire/module"
	"github.com/secriy/golire/util"
	"github.com/spf13/cobra"
)

var gors int // the number of goroutines

// pingCmd represents the ping command
var pingCmd = &cobra.Command{
	Use:   "ping [Host]",
	Short: "Find all surviving hosts using ICMP message.",
	Long: `Host specific as a CIDR notation,
like 192.168.0.0/24, 192.168.1.1/31.
`,
	Args: cobra.RangeArgs(1, 1),
	Run: func(cmd *cobra.Command, args []string) {
		module.Log(level) // set logger level
		var count int64   // count the alive host

		ips := util.ParseHost(args[0])

		// limit goroutine numbers
		if gors > len(ips) {
			gors = len(ips)
		}

		wg := &sync.WaitGroup{}
		p, _ := ants.NewPoolWithFunc(gors, func(i interface{}) {
			ip := i.(string)
			if module.PingDefault(ip) {
				util.Print("PING", "Found a alive host: "+ip)
				atomic.AddInt64(&count, 1)
			}
			wg.Done()
		}, ants.WithPreAlloc(true))
		defer p.Release()

		for _, v := range ips {
			wg.Add(1)
			_ = p.Invoke(v)
		}

		wg.Wait()

		util.Print("PING", fmt.Sprintf("Total host: %d, surviving host: %d, running goroutines: %d", len(ips), count, p.Running()))
	},
}

func init() {
	rootCmd.AddCommand(pingCmd)

	pingCmd.Flags().IntVarP(&gors, "threads", "t", 500, "-t [Number]")
}
