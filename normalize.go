package arp

import (
	"fmt"
	"regexp"
	"strings"
)

func normalizeMACAddr(addr string) string {
	parts := regexp.MustCompile("[-:]").Split(addr, -1)
	for i, part := range parts {
		parts[i] = fmt.Sprintf("%02s", strings.ToLower(part))
	}
	return strings.Join(parts, ":")
}
