package main

import (
	"fmt"
	"proxyscanner/pkg/proxy"
)

// TODO: command Line
func main() {
	proxyInfo, err := proxy.CheckHTTPProxy("127.0.0.1", 1086, "HTTP")
	if err != nil {
		fmt.Printf("Not a HTTP Proxy: %v\n", err)
	} else {
		fmt.Printf("HTTP proxy IP: %s\n Port: %d\n", proxyInfo.IP, proxyInfo.Port)
	}

	proxyInfo, err = proxy.CheckSOCKSProxy("127.0.0.1", 1086, "SOCKS5")
	if err != nil {
		fmt.Printf("Not a SOCKS Proxy: %v\n", err)
	} else {
		fmt.Printf("SOCKS proxy IP: %s\n Port: %d\n", proxyInfo.IP, proxyInfo.Port)
	}
}
