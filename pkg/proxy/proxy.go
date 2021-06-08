package proxy

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"proxyscanner/pkg/models"
	"strings"
	"time"

	"h12.io/socks"
)

var (
	webURL  = "http://email.163.com"
	timeout = 2

	socksProxyProtocol = map[string]int{"SOCKS4": socks.SOCKS4, "SOCKS5": socks.SOCKS5, "SOCKS4A": socks.SOCKS4A, "SOCKS5A": socks.SOCKS5}
)

func checkHTTPProxy(ip string, port int, protocol string) (models.ProxyInfo, error) {
	proxyInfo := models.ProxyInfo{IP: ip, Port: port, Protocol: protocol}
	proxy := fmt.Sprintf("%v://%v:%v", protocol, ip, port)

	proxyURL, err := url.Parse(proxy)
	if err != nil {
		return proxyInfo, err
	}

	transport := &http.Transport{Proxy: http.ProxyURL(proxyURL)}
	client := &http.Client{Transport: transport, Timeout: time.Duration(timeout) * time.Second}

	resp, err := client.Get(webURL)
	if err != nil {
		return proxyInfo, err
	}

	if resp.StatusCode == http.StatusOK {
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return proxyInfo, err
		}

		if !strings.Contains(string(body), "网易免费邮箱") {
			return proxyInfo, fmt.Errorf("Not A Valid Proxy")
		}
	}
	return proxyInfo, nil
}

func checkSOCKSProxy(ip string, port int, protocol string) (models.ProxyInfo, error) {
	proxyInfo := models.ProxyInfo{IP: ip, Port: port, Protocol: protocol}
	proxy := fmt.Sprintf("%v:%v", ip, port)

	dialSocksProxy := socks.DialSocksProxy(socksProxyProtocol[protocol], proxy)

	transport := &http.Transport{Dial: dialSocksProxy}
	httpClient := &http.Client{Transport: transport, Timeout: time.Duration(timeout) * time.Second}

	resp, err := httpClient.Get(webURL)
	if err != nil {
		return proxyInfo, err
	}

	if resp.StatusCode == http.StatusOK {
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return proxyInfo, err
		}

		if !strings.Contains(string(body), "网易免费邮箱") {
			return proxyInfo, fmt.Errorf("Not A Valid Proxy")
		}
	}
	return proxyInfo, nil
}
