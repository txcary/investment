package main

import (
	"github.com/txcary/investment/restful"
	"fmt"
	"net"
	"errors"
)

var (
	serverPort string = "8080"
)

func getIp() (string, error) {
	var ipLocal = net.IPv4(127,0,0,1)
	ifaces, err := net.Interfaces()
	if err!=nil {
		return "", err
	}
	for _, i := range ifaces {
	    addrs, err := i.Addrs()
	    if err!=nil {
	    	return "", err
	    }
	    for _, addr := range addrs {
	        var ip net.IP
	        switch v := addr.(type) {
	        case *net.IPNet:
	                ip = v.IP
	                break
	        case *net.IPAddr:
	                ip = v.IP
	                break
	        default:
	        	continue
	        }
	        if ip.Equal(ipLocal) {
	        	continue
	        }
	        return ip.String(), nil
	    }
	}
	return "", errors.New("No iface found!")
}

func main() {
	restful.StartServer(serverPort)
	ip,_ := getIp()	
	fmt.Println("Serving",ip,serverPort,"...")
	err := restful.Wait()
	fmt.Println(err)
}