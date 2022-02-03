package main

import (
	"flag"
	"fmt"
	"log"
	"sync"
	"time"

	"github.com/go-ping/ping"
	"inet.af/netaddr"
)

var wg sync.WaitGroup
var BaseIP = flag.String("baseIP", "", "First IP address")
var Range = flag.Int("range", 10, "Incremental Range to the BaseIP")
var Timeout = flag.Int("timeout", 1000, "Ping Timeout in msec")
var PauseAfter = flag.Int("pauseAfter", 0, "Pause after nth icmp packet for 2*timeout sec")

func pingHost(host string, wg *sync.WaitGroup) {
	pinger, err := ping.NewPinger(host)
	if err != nil {
		fmt.Println("ERROR:", err)
		return
	}
	pinger.Count = 1
	pinger.Timeout = time.Duration(*Timeout) * time.Millisecond
	err = pinger.Run()
	if err != nil {
		fmt.Println("Failed to ping target host:", err)
	}
	stats := pinger.Statistics()
	if stats.PacketsRecv == 1 {
		fmt.Printf("%15s is alive\n", host)
	} else {
		fmt.Printf("%15s is dead\n", host)
	}
	wg.Done()
}

func main() {
	flag.Parse()
	version := "1.0"
	fmt.Printf("goPing version %s\n", version)
	if *BaseIP == "" {
		log.Fatal("baseIP is required to proceed")
	}
	ip, _ := netaddr.ParseIP(*BaseIP)
	for i := 1; i <= *Range; i++ {
		if *PauseAfter != 0 {
			if i%*PauseAfter == 0 {
				fmt.Printf("Pausing for %d msec\n", 2*(*Timeout))
				time.Sleep(2 * time.Duration(*Timeout) * time.Millisecond)
			}
		}
		wg.Add(1)
		ip = ip.Next()
		host := ip.String()
		go pingHost(host, &wg)
	}
	wg.Wait()
}
