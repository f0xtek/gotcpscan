package main

import (
	"flag"
	"fmt"
	"net"
	"os"
	"sort"

	"gitlab.com/f0xtek/gotcpscan/pkg/portformat"
)

func main() {

	hostPtr := flag.String("host", "127.0.0.1", "The host you wish to scan. Default: 127.0.0.1")
	portPtr := flag.String("ports", "1-65535", "The ports you wish to scan. Accepts nmap formatted port specifications. Default: 1-65535")
	flag.Parse()

	portRange, err := portformat.Parse(*portPtr)
	if err != nil {
		fmt.Printf("Problem parsing ports.\n%v", err)
		os.Exit(1)
	}

	ports := make(chan int, 1000)
	results := make(chan int)
	var openPorts []int

	for i := 0; i < cap(ports); i++ {
		go worker(*hostPtr, ports, results)
	}

	go func() {
		for _, p := range portRange {
			ports <- p
		}
	}()

	for i := 0; i < len(portRange); i++ {
		port := <-results
		if port != 0 {
			openPorts = append(openPorts, port)
		}
	}

	close(ports)
	close(results)
	sort.Ints(openPorts)

	for _, port := range openPorts {
		fmt.Printf("%d open\n", port)
	}
}

func worker(host string, ports, results chan int) {
	for p := range ports {
		address := fmt.Sprintf("%s:%d", host, p)
		conn, err := net.Dial("tcp", address)
		if err != nil {
			results <- 0
			continue
		}
		conn.Close()
		results <- p
	}
}
