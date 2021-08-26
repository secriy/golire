package cmd

import (
	"fmt"
	"sync"
	"sync/atomic"

	"github.com/secriy/golire/module"
	"github.com/secriy/golire/util"
	"github.com/spf13/cobra"
)

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

		ips := util.ParseHost(args[0])
		wg := &sync.WaitGroup{}

		var count int64 // count the alive host

		for _, v := range ips {
			wg.Add(1)

			go func(ip string) {
				defer wg.Done()
				if module.PingDefault(ip) {
					util.Print("PING", "Found a alive host: "+ip)
					atomic.AddInt64(&count, 1)
				}
			}(v)
		}
		wg.Wait()

		util.Print("PING", fmt.Sprintf("Total host: %d, surviving host: %d", len(ips), count))
	},
}

func init() {
	rootCmd.AddCommand(pingCmd)
}
