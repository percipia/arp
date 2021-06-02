package main

import (
	"fmt"

	"../../arp"
)

func main() {
	for ip, _ := range arp.Table() {
		fmt.Printf("%s : %s\n", ip, arp.Search(ip))
	}
}
