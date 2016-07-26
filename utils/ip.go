package utils

import (
	"fmt"
	"net"
	"strings"

	"github.com/astaxie/beego/context"
)

const (
	LOCAL_HOST_IPV4        = "127.0.0.1"
	LOCAL_HOST_IPV6_PREFIX = "[::1]"
)

func IsLocalHostIP(ip string) bool {
	//169.254.0.0/16 也是本机ip，但只有在网卡从dhcp分配不到ip地址时才会有这种地址
	return strings.HasPrefix(ip, LOCAL_HOST_IPV4) || strings.HasPrefix(ip, LOCAL_HOST_IPV6_PREFIX)
}

func IsPrivateIP(addr net.IP) bool {
	string_networks := []string{"192.168.0.0/16", "172.16.0.0/12", "10.0.0.0/8"}
	for _, s_network := range string_networks {
		_, network, _ := net.ParseCIDR(s_network)
		if network.Contains(addr) {
			return true
		}
	}

	return false
}

func IsPrivateIPStr(ip string) bool {
	if IsLocalHostIP(ip) {
		return true
	}

	addr := net.ParseIP(ip)
	if addr == nil {
		return false
	}

	return IsPrivateIP(addr)
}

func GetPrivateIPs() (*[]net.IP, error) {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return nil, err
	}

	IPs := make([]net.IP, 0, 10)

	for _, addr := range addrs {
		if strings.HasPrefix(addr.String(), LOCAL_HOST_IPV4) {
			continue
		}

		addr, _, _ := net.ParseCIDR(addr.String())
		if IsPrivateIP(addr) {
			IPs = append(IPs, addr)
		}
	}

	return &IPs, nil
}

func GetRealClientIPStr(ctx *context.Context) string {
	ip := ctx.Request.Header.Get("X-Forwarded-For")
	if ip != "" {
		ips := strings.Split(ip, ",")

		if len(ips) <= 0 {
			return ctx.Request.RemoteAddr
		}

		return ips[len(ips)-1]

	} else {
		return ctx.Request.RemoteAddr
	}
}

func GetRealClientIP(ctx *context.Context) (*net.IP, error) {
	ip_str := GetRealClientIPStr(ctx)
	addr := net.ParseIP(ip_str)
	if addr == nil {
		Logger.Error("Failed in paring ip %s\n", ip_str)
		return nil, fmt.Errorf("Failed in paring ip %s", ip_str)
	}

	return &addr, nil
}

func IsClientFromSamePrivateNetwork(ctx *context.Context) bool {
	ip_str := GetRealClientIPStr(ctx)
	if IsLocalHostIP(ip_str) {
		return true
	}

	addr := net.ParseIP(ip_str)
	if addr == nil {

		Logger.Error("Failed in paring ip %s\n", ip_str)
		return false
	}

	return IsSamePrivateNetworkIP(addr)
}

//检查这个ip是否和本机在一个IP段
func IsSamePrivateNetworkIP(ip net.IP) bool {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return false
	}

	for _, addr := range addrs {
		if strings.HasPrefix(addr.String(), LOCAL_HOST_IPV4) {
			continue
		}

		addr, network, _ := net.ParseCIDR(addr.String())
		if IsPrivateIP(addr) {
			if network.Contains(ip) {
				return true
			}
		}
	}

	return false
}

func IsAlipayIp(ip string) bool {
	addr := net.ParseIP(ip)
	if addr == nil {
		return false
	}

	string_networks := []string{"121.0.26.0/23", "110.75.128.0/19"}
	for _, s_network := range string_networks {
		_, network, _ := net.ParseCIDR(s_network)
		if network.Contains(addr) {
			return true
		}
	}

	return false
}
