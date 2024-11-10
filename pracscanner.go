package main

import (
	"flag"
	"fmt"
	"net"
	"sync"
	"time"
)

const (
	default_ip     = "Localhost"
	PORT_RANGE int = 65535
)

type PortResult struct {
	port   int
	status bool
}

type PortScanner struct {
	commonPorts []int
	scanAll     bool
}

func ScanPort(ip string, port int, c chan PortResult, wg *sync.WaitGroup) {
	defer wg.Done()

	address := fmt.Sprintf("%s:%d", ip, port)
	conn, err := net.DialTimeout("tcp", address, 2*time.Second)
	if err != nil {
		c <- PortResult{port: port, status: false}
		return
	}
	c <- PortResult{port: port, status: true}
	conn.Close()
}

func (ps PortScanner) RunScan(ip string, c chan PortResult, wg *sync.WaitGroup) {
	portsToScan := ps.commonPorts
	if ps.scanAll {
		portsToScan = make([]int, PORT_RANGE)
		for i := 0; i < cap(portsToScan); i++ {
			portsToScan[i] = i
		}
	}

	for _, port := range portsToScan {
		wg.Add(1)
		go ScanPort(ip, port, c, wg)
	}

	func() {
		wg.Wait()
		close(c)
	}()
}

func (pr PortResult) display(showAll bool) {
	if pr.status || showAll {
		status := "CLOSED"
		if pr.status {
			status = "OPEN"
		}
		fmt.Printf("PORT: %d %s\n", pr.port, status)
	}
}
func main() {

	var (
		commonPorts = []int{
			20, 21, 22, 23, 25, 53, 80, 110, 143, 443, 445, 465, 587, 993, 995,
			3306, 3389, 5060, 5061, 8080,
		}
	)

	ip := flag.String("ip", default_ip, "IP address of host network")
	scanAll := flag.Bool("all", false, "Scan all ports")
	showPort := flag.Bool("show", false, "Show both open and closed port")

	flag.Parse()

	portchannel := make(chan PortResult)
	var wg sync.WaitGroup

	ports := PortScanner{
		commonPorts: commonPorts,
		scanAll:     *scanAll,
	}

	go ports.RunScan(*ip, portchannel, &wg)

	for p := range portchannel {
		p.display(*showPort)
	}
}
